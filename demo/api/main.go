package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type server struct {
	store Store
	host  string // container hostname: shows which copy answered
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

func (s *server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.health)
	mux.HandleFunc("GET /api/themes", s.themes)
	mux.HandleFunc("POST /api/qr", s.qr)
	mux.HandleFunc("POST /api/trees", s.saveTree)
	mux.HandleFunc("GET /api/trees/{id}", s.getTree)
	return mux
}

func (s *server) health(w http.ResponseWriter, r *http.Request) {
	// the API is healthy even without a database; "db" tells which mode it is in
	if err := s.store.Ping(r.Context()); err != nil {
		writeJSON(w, 200, map[string]any{"status": "ok", "host": s.host, "db": "not connected"})
		return
	}
	n, _ := s.store.CountTrees(r.Context())
	writeJSON(w, 200, map[string]any{"status": "ok", "host": s.host, "db": "connected", "shared_trees": n})
}

func (s *server) themes(w http.ResponseWriter, r *http.Request) {
	t, err := s.store.Themes(r.Context())
	if err != nil {
		fail(w, 500, "could not load themes")
		return
	}
	writeJSON(w, 200, map[string]any{"host": s.host, "themes": t})
}

type qrReq struct {
	URL string `json:"url"`
}

// POST /api/qr  {"url": "..."}  ->  verified QR matrix. The URL is never stored.
func (s *server) qr(w http.ResponseWriter, r *http.Request) {
	var in qrReq
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096)).Decode(&in); err != nil {
		fail(w, 400, "invalid request")
		return
	}
	link, err := validateURL(in.URL)
	if err != nil {
		fail(w, 400, err.Error())
		return
	}
	rows, err := qrMatrix(link)
	if err != nil {
		fail(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{"url": link, "size": len(rows), "matrix": rows, "verified": true, "host": s.host})
}

// POST /api/trees  -> stored only when the user clicks "share".
func (s *server) saveTree(w http.ResponseWriter, r *http.Request) {
	var in Tree
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096)).Decode(&in); err != nil {
		fail(w, 400, "invalid request")
		return
	}
	link, err := validateURL(in.URL)
	if err != nil {
		fail(w, 400, err.Error())
		return
	}
	in.URL = link
	for _, c := range []string{in.Foliage, in.Trunk, in.Background} {
		if !hexRe.MatchString(c) {
			fail(w, 400, "colours must look like #1a2b3c")
			return
		}
	}
	if contrast(in.Trunk, in.Background) < minContrast {
		fail(w, 400, "blocks and ground are too similar to scan")
		return
	}
	in.Theme = strings.TrimSpace(in.Theme)
	if len(in.Theme) > 40 {
		in.Theme = in.Theme[:40]
	}
	saved, err := s.store.SaveTree(r.Context(), in)
	if errors.Is(err, errNoDB) {
		fail(w, http.StatusServiceUnavailable, "sharing needs a database: start the db container first")
		return
	}
	if err != nil {
		fail(w, 500, "could not save")
		return
	}
	writeJSON(w, 201, saved)
}

func (s *server) getTree(w http.ResponseWriter, r *http.Request) {
	t, err := s.store.GetTree(r.Context(), r.PathValue("id"))
	if errors.Is(err, errNotFound) {
		fail(w, 404, "tree not found")
		return
	}
	if errors.Is(err, errNoDB) {
		fail(w, http.StatusServiceUnavailable, "no database connected")
		return
	}
	if err != nil {
		fail(w, 500, "could not load")
		return
	}
	writeJSON(w, 200, t)
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	// DATABASE_URL is the one setting that changes between steps of the lesson:
	// not set -> built-in themes, no sharing. Set -> Postgres.
	var store Store = memStore{}
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		pg, err := connect(context.Background(), dsn, 30*time.Second)
		if err != nil {
			log.Fatalf("database: %v", err)
		}
		store = pg
		log.Print("connected to the database")
	} else {
		log.Print("DATABASE_URL not set: running without a database")
	}
	host, _ := os.Hostname()
	s := &server{store: store, host: host}
	addr := ":" + env("PORT", "8080")
	log.Printf("api listening on %s (host %s)", addr, host)
	log.Fatal(http.ListenAndServe(addr, s.routes()))
}
