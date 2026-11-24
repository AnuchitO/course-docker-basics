package main

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
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
	mux.HandleFunc("GET /api/trees", s.listTrees)
	mux.HandleFunc("POST /api/trees", s.saveTree)
	mux.HandleFunc("GET /api/r/{id}", s.scan)
	mux.HandleFunc("POST /api/trees/{id}/like", s.like(1))
	mux.HandleFunc("DELETE /api/trees/{id}/like", s.like(-1))
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
	writeJSON(w, 200, map[string]any{"host": s.host, "seasons": t.Seasons, "swatches": t.Swatches})
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
	in.Name = strings.TrimSpace(in.Name)
	if utf8.RuneCountInString(in.Name) > 40 {
		fail(w, 400, "name is too long (max 40 characters)")
		return
	}
	if in.Name == "" {
		in.Name = link
		if r := []rune(link); len(r) > 40 {
			in.Name = string(r[:40]) + "…"
		}
	}
	if !seasonRe.MatchString(in.Season) || in.Swatch < 0 || in.Swatch > 9 {
		fail(w, 400, "unknown season or colour")
		return
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
	saved.Link = trackingLink(r, saved.ID)
	writeJSON(w, 201, saved)
}

// GET /api/trees?sort=new|popular&season=spring  ->  the saved gardens, each with its QR matrix
// so the page can draw a thumbnail without a second request. Presets are seeded by db/init.sql.
func (s *server) listTrees(w http.ResponseWriter, r *http.Request) {
	sort, season := r.URL.Query().Get("sort"), r.URL.Query().Get("season")
	if season != "" && !seasonRe.MatchString(season) {
		fail(w, 400, "unknown season")
		return
	}
	trees, err := s.store.ListTrees(r.Context(), sort, season, 48)
	if errors.Is(err, errNoDB) {
		fail(w, http.StatusServiceUnavailable, "the gallery needs a database: start the db container first")
		return
	}
	if err != nil {
		fail(w, 500, "could not load the gallery")
		return
	}
	for i := range trees {
		trees[i].Link = trackingLink(r, trees[i].ID)
		if rows, err := cachedMatrix(trees[i].Link); err == nil {
			trees[i].Matrix = rows
		}
	}
	writeJSON(w, 200, map[string]any{"trees": trees, "host": s.host})
}

// trackingLink is what a saved garden's QR code encodes. It points back at this site, so a scan reaches
// the API first: that is the only moment we can count it.
func trackingLink(r *http.Request, id string) string {
	scheme := "http"
	if r.Header.Get("X-Forwarded-Proto") == "https" || r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host + "/api/r/" + id
}

var landing = template.Must(template.New("l").Parse(`<!doctype html><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>{{.Name}}</title><body style="font-family:system-ui;max-width:32rem;margin:15vh auto;padding:0 1.2rem;line-height:1.5">
<h1 style="font-size:1.3rem">{{.Name}}</h1><p style="white-space:pre-wrap;word-break:break-word;background:#f5efe6;padding:1rem;border-radius:12px">{{.URL}}</p>
<p style="color:#8c8274;font-size:.85rem">Scanned {{.Scans}} times · TreeQR</p>`))

// GET /api/r/{id}  <- what a phone opens after scanning. Counts the scan, then sends the visitor on:
// a web link gets a 302 redirect, anything else (text, Wi-Fi, mailto) gets a small page that shows it.
func (s *server) scan(w http.ResponseWriter, r *http.Request) {
	count := s.store.ScanTree
	if r.Method == http.MethodHead { // link checkers and previews often send HEAD first: look, don't count
		count = func(ctx context.Context, id string) (Tree, error) { return s.store.GetTree(ctx, id) }
	}
	t, err := count(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, errNotFound):
		fail(w, 404, "this garden does not exist")
		return
	case err != nil:
		fail(w, 503, "could not open this garden")
		return
	}
	w.Header().Set("Cache-Control", "no-store") // a cached redirect would hide later scans
	if u, err := url.Parse(t.URL); err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != "" {
		http.Redirect(w, r, t.URL, http.StatusFound)
		return
	}
	_ = landing.Execute(w, t)
}

// like returns a handler that moves a tree's like counter by delta.
func (s *server) like(delta int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n, err := s.store.LikeTree(r.Context(), r.PathValue("id"), delta)
		switch {
		case errors.Is(err, errNotFound):
			fail(w, 404, "tree not found")
		case errors.Is(err, errNoDB):
			fail(w, http.StatusServiceUnavailable, "likes need a database: start the db container first")
		case err != nil:
			fail(w, 500, "could not like")
		default:
			writeJSON(w, 200, map[string]any{"id": r.PathValue("id"), "likes": n})
		}
	}
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

var seasonRe = regexp.MustCompile(`^[a-z]{3,12}$`)

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
