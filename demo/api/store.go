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

// Theme is one seasonal colour set. Rows live in the Postgres `themes` table.
type Theme struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Foliage    string `json:"foliage"`
	Trunk      string `json:"trunk"`
	Background string `json:"background"`
	Accent     string `json:"accent"`
}

// Tree is a tree the user explicitly chose to share.
type Tree struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Theme      string    `json:"theme"`
	Foliage    string    `json:"foliage"`
	Trunk      string    `json:"trunk"`
	Background string    `json:"background"`
	CreatedAt  time.Time `json:"created_at"`
}

var (
	errNotFound = errors.New("not found")
	errNoDB     = errors.New("no database connected")
)

// Store is everything the handlers need from the database.
// Keeping it an interface lets the tests run without Postgres.
type Store interface {
	Ping(ctx context.Context) error
	Themes(ctx context.Context) ([]Theme, error)
	SaveTree(ctx context.Context, t Tree) (Tree, error)
	GetTree(ctx context.Context, id string) (Tree, error)
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

func (s *pgStore) Themes(ctx context.Context) ([]Theme, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT slug, name, foliage, trunk, background, accent FROM themes ORDER BY position`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Theme{}
	for rows.Next() {
		var t Theme
		if err := rows.Scan(&t.Slug, &t.Name, &t.Foliage, &t.Trunk, &t.Background, &t.Accent); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (s *pgStore) SaveTree(ctx context.Context, t Tree) (Tree, error) {
	t.ID = newID()
	err := s.pool.QueryRow(ctx,
		`INSERT INTO trees (id, url, theme, foliage, trunk, background)
		 VALUES ($1,$2,$3,$4,$5,$6) RETURNING created_at`,
		t.ID, t.URL, t.Theme, t.Foliage, t.Trunk, t.Background).Scan(&t.CreatedAt)
	return t, err
}

func (s *pgStore) GetTree(ctx context.Context, id string) (Tree, error) {
	var t Tree
	err := s.pool.QueryRow(ctx,
		`SELECT id, url, theme, foliage, trunk, background, created_at FROM trees WHERE id = $1`, id).
		Scan(&t.ID, &t.URL, &t.Theme, &t.Foliage, &t.Trunk, &t.Background, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return t, errNotFound
	}
	return t, err
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
