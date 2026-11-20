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
		{"javascript scheme", `{"url":"javascript:alert(1)"}`, 400},
		{"ftp scheme", `{"url":"ftp://example.com"}`, 400},
		{"no host", `{"url":"https://"}`, 400},
		{"too long", `{"url":"https://a.com/` + strings.Repeat("x", 600) + `"}`, 400},
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
