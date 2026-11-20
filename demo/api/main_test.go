package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeStore struct{ trees map[string]Tree }

func (f *fakeStore) Ping(context.Context) error { return nil }
func (f *fakeStore) Themes(context.Context) (Themes, error) {
	return defaultThemes, nil
}
func (f *fakeStore) SaveTree(_ context.Context, t Tree) (Tree, error) {
	t.ID = "abc"
	f.trees[t.ID] = t
	return t, nil
}
func (f *fakeStore) GetTree(_ context.Context, id string) (Tree, error) {
	if t, ok := f.trees[id]; ok {
		return t, nil
	}
	return Tree{}, errNotFound
}
func (f *fakeStore) ScanTree(_ context.Context, id string) (Tree, error) {
	t, ok := f.trees[id]
	if !ok {
		return t, errNotFound
	}
	t.Scans++
	f.trees[id] = t
	return t, nil
}
func (f *fakeStore) LikeTree(_ context.Context, id string, d int) (int, error) {
	t, ok := f.trees[id]
	if !ok {
		return 0, errNotFound
	}
	t.Likes += d
	f.trees[id] = t
	return t.Likes, nil
}
func (f *fakeStore) ListTrees(_ context.Context, _, _ string, _ int) ([]Tree, error) {
	out := []Tree{}
	for _, t := range f.trees {
		out = append(out, t)
	}
	return out, nil
}
func (f *fakeStore) CountTrees(context.Context) (int, error) { return len(f.trees), nil }

func do(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestQR(t *testing.T) {
	h := (&server{store: &fakeStore{trees: map[string]Tree{}}, host: "t"}).routes()
	cases := []struct {
		name, body string
		code       int
	}{
		{"valid", `{"url":"https://example.com/hello"}`, 200},
		{"empty", `{"url":""}`, 400},
		{"plain text", `{"url":"hello garden"}`, 200},
		{"too long", `{"url":"https://a.com/` + strings.Repeat("x", 300) + `"}`, 400},
		{"bad json", `nope`, 400},
	}
	for _, c := range cases {
		if got := do(h, "POST", "/api/qr", c.body).Code; got != c.code {
			t.Errorf("%s: got %d want %d", c.name, got, c.code)
		}
	}
}

func TestQRMatrixIsSquareAndVerified(t *testing.T) {
	rec := do((&server{store: &fakeStore{trees: map[string]Tree{}}}).routes(), "POST", "/api/qr", `{"url":"https://example.com"}`)
	var out struct {
		Size     int      `json:"size"`
		Matrix   []string `json:"matrix"`
		Verified bool     `json:"verified"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatal(err)
	}
	if !out.Verified || out.Size == 0 || len(out.Matrix) != out.Size || len(out.Matrix[0]) != out.Size {
		t.Fatalf("bad matrix: %+v", out)
	}
}

func TestShareTree(t *testing.T) {
	st := &fakeStore{trees: map[string]Tree{}}
	h := (&server{store: st}).routes()
	ok := `{"url":"https://example.com","season":"spring","swatch":1}`
	if c := do(h, "POST", "/api/trees", ok).Code; c != 201 {
		t.Fatalf("share: %d", c)
	}
	bad := `{"url":"https://example.com","season":"spring","swatch":99}`
	if c := do(h, "POST", "/api/trees", bad).Code; c != 400 {
		t.Fatalf("unknown colour should be rejected, got %d", c)
	}
	if c := do(h, "GET", "/api/trees/abc", "").Code; c != 200 {
		t.Fatalf("get: %d", c)
	}
	if c := do(h, "GET", "/api/trees/nope", "").Code; c != 404 {
		t.Fatalf("missing: %d", c)
	}
}

func TestWithoutDatabase(t *testing.T) {
	h := (&server{store: memStore{}, host: "t"}).routes()
	if rec := do(h, "GET", "/api/health", ""); rec.Code != 200 || !strings.Contains(rec.Body.String(), "not connected") {
		t.Fatalf("health: %d %s", rec.Code, rec.Body)
	}
	if rec := do(h, "GET", "/api/themes", ""); rec.Code != 200 || !strings.Contains(rec.Body.String(), "autumn") {
		t.Fatalf("themes: %d", rec.Code)
	}
	body := `{"url":"https://example.com","season":"spring","swatch":0}`
	if c := do(h, "POST", "/api/trees", body).Code; c != 503 {
		t.Fatalf("share without db should be 503, got %d", c)
	}
}

func TestGallery(t *testing.T) {
	h := (&server{store: &fakeStore{trees: map[string]Tree{}}}).routes()
	rec := do(h, "POST", "/api/trees", `{"name":"Docs","url":"https://docs.docker.com","season":"summer","swatch":2}`)
	if rec.Code != 201 {
		t.Fatalf("save: %d %s", rec.Code, rec.Body)
	}
	rec = do(h, "GET", "/api/trees", "")
	var got struct{ Trees []Tree }
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	if rec.Code != 200 || len(got.Trees) != 1 || got.Trees[0].Name != "Docs" || got.Trees[0].Swatch != 2 {
		t.Fatalf("list: %d %s", rec.Code, rec.Body)
	}
	rec = do(h, "POST", "/api/trees/abc/like", "")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), `"likes":1`) {
		t.Fatalf("like: %d %s", rec.Code, rec.Body)
	}
	if do(h, "POST", "/api/trees/nope/like", "").Code != 404 {
		t.Fatal("like of a missing tree should be 404")
	}
	rec = do(h, "GET", "/api/trees", "")
	if !strings.Contains(rec.Body.String(), `"matrix"`) {
		t.Fatalf("list should carry matrices: %s", rec.Body)
	}
	// no name given: falls back to the text
	do(h, "POST", "/api/trees", `{"url":"hello","season":"spring","swatch":0}`)
	// without a database the gallery says so
	rec = do((&server{store: memStore{}}).routes(), "GET", "/api/trees", "")
	if rec.Code != 503 {
		t.Fatalf("no db: %d", rec.Code)
	}
}

func TestScan(t *testing.T) {
	h := (&server{store: &fakeStore{trees: map[string]Tree{}}}).routes()
	do(h, "POST", "/api/trees", `{"url":"https://docs.docker.com","season":"spring","swatch":0}`)
	rec := do(h, "GET", "/api/r/abc", "")
	if rec.Code != 302 || rec.Header().Get("Location") != "https://docs.docker.com" || rec.Header().Get("Cache-Control") != "no-store" {
		t.Fatalf("redirect: %d %v", rec.Code, rec.Header())
	}
	do(h, "GET", "/api/r/abc", "")
	rec = do(h, "GET", "/api/trees", "")
	var got struct{ Trees []Tree }
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	if got.Trees[0].Scans != 2 || !strings.HasSuffix(got.Trees[0].Link, "/api/r/abc") {
		t.Fatalf("scans/link: %+v", got.Trees[0])
	}
	// plain text is shown on a page, not redirected
	do(h, "POST", "/api/trees", `{"url":"hello garden","season":"spring","swatch":0}`)
	if do(h, "GET", "/api/r/nope", "").Code != 404 {
		t.Fatal("unknown id should be 404")
	}
}
