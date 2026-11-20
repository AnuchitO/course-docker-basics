package main

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Season sets the page background for one season (rows of the `seasons` table).
type Season struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Background string `json:"background"`
}

// Themes is everything the page needs to dress the garden: seasons and flower colours.
type Themes struct {
	Seasons  []Season `json:"seasons"`
	Swatches []string `json:"swatches"`
}

// Tree is a garden the user explicitly chose to share.
type Tree struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Season    string    `json:"season"`
	Swatch    int       `json:"swatch"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	errNotFound = errors.New("not found")
	errNoDB     = errors.New("no database connected")
)

// Store is everything the handlers need from the database.
// Keeping it an interface lets the tests run without Postgres.
type Store interface {
	Ping(ctx context.Context) error
	Themes(ctx context.Context) (Themes, error)
	SaveTree(ctx context.Context, t Tree) (Tree, error)
	GetTree(ctx context.Context, id string) (Tree, error)
	ListTrees(ctx context.Context, limit int) ([]Tree, error)
	CountTrees(ctx context.Context) (int, error)
}

type pgStore struct{ pool *pgxpool.Pool }

// connect keeps retrying for a while: in Docker the database container may
// still be starting when the API container starts.
func connect(ctx context.Context, dsn string, wait time.Duration) (*pgStore, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	deadline := time.Now().Add(wait)
	for {
		pctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		err = pool.Ping(pctx)
		cancel()
		if err == nil {
			return &pgStore{pool}, nil
		}
		if time.Now().After(deadline) {
			pool.Close()
			return nil, err
		}
		time.Sleep(time.Second)
	}
}

func (s *pgStore) Ping(ctx context.Context) error { return s.pool.Ping(ctx) }

func (s *pgStore) Themes(ctx context.Context) (Themes, error) {
	out := Themes{Seasons: []Season{}, Swatches: []string{}}
	rows, err := s.pool.Query(ctx, `SELECT slug, name, background FROM seasons ORDER BY position`)
	if err != nil {
		return out, err
	}
	for rows.Next() {
		var x Season
		if err := rows.Scan(&x.Slug, &x.Name, &x.Background); err != nil {
			rows.Close()
			return out, err
		}
		out.Seasons = append(out.Seasons, x)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return out, err
	}
	rows, err = s.pool.Query(ctx, `SELECT hex FROM swatches ORDER BY position`)
	if err != nil {
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		var h string
		if err := rows.Scan(&h); err != nil {
			return out, err
		}
		out.Swatches = append(out.Swatches, h)
	}
	return out, rows.Err()
}

func (s *pgStore) SaveTree(ctx context.Context, t Tree) (Tree, error) {
	t.ID = newID()
	err := s.pool.QueryRow(ctx,
		`INSERT INTO trees (id, name, url, season, swatch)
		 VALUES ($1,$2,$3,$4,$5) RETURNING created_at`,
		t.ID, t.Name, t.URL, t.Season, t.Swatch).Scan(&t.CreatedAt)
	return t, err
}

func (s *pgStore) GetTree(ctx context.Context, id string) (Tree, error) {
	var t Tree
	err := s.pool.QueryRow(ctx,
		`SELECT id, name, url, season, swatch, created_at FROM trees WHERE id = $1`, id).
		Scan(&t.ID, &t.Name, &t.URL, &t.Season, &t.Swatch, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return t, errNotFound
	}
	return t, err
}

func (s *pgStore) ListTrees(ctx context.Context, limit int) ([]Tree, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, name, url, season, swatch, created_at FROM trees ORDER BY created_at DESC, id LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Tree{}
	for rows.Next() {
		var t Tree
		if err := rows.Scan(&t.ID, &t.Name, &t.URL, &t.Season, &t.Swatch, &t.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (s *pgStore) CountTrees(ctx context.Context) (int, error) {
	var n int
	err := s.pool.QueryRow(ctx, `SELECT count(*) FROM trees`).Scan(&n)
	return n, err
}

// newID returns a short, URL-safe, unguessable id such as "k3f9q2xa".
func newID() string {
	b := make([]byte, 5)
	_, _ = rand.Read(b)
	return strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b))
}
