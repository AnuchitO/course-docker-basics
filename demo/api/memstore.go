package main

import "context"

// memStore is used when DATABASE_URL is not set. It serves built-in colours,
// so the API container can run on its own (lesson step 2).
// With Postgres (step 4) the page gets six flower colours instead of three.
// Sharing needs a real database, so it says so.
type memStore struct{}

var defaultThemes = Themes{
	Seasons: []Season{
		{"spring", "Spring", "#f5efe6"},
		{"summer", "Summer", "#eef2e4"},
		{"autumn", "Autumn", "#f6ece0"},
	},
	Swatches: []string{"#e98aab", "#a57ad6", "#e0473c"},
}

func (memStore) Ping(context.Context) error                     { return errNoDB }
func (memStore) Themes(context.Context) (Themes, error)         { return defaultThemes, nil }
func (memStore) SaveTree(context.Context, Tree) (Tree, error)   { return Tree{}, errNoDB }
func (memStore) GetTree(context.Context, string) (Tree, error)  { return Tree{}, errNoDB }
func (memStore) CountTrees(context.Context) (int, error)        { return 0, errNoDB }
func (memStore) ListTrees(context.Context, int) ([]Tree, error) { return nil, errNoDB }
