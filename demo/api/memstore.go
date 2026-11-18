package main

import "context"

// memStore is used when DATABASE_URL is not set. It serves built-in themes,
// so the API container can run on its own (lesson step 2).
// Sharing needs a real database, so it says so.
type memStore struct{}

var defaultThemes = []Theme{
	{"spring", "Spring", "#f4a6c0", "#4a3426", "#e6f4d4", "#ffffff"},
	{"summer", "Summer", "#2e9e57", "#3b2d22", "#e3f6e6", "#a7f3d0"},
	{"autumn", "Autumn", "#e8761c", "#4a2c1c", "#f7ead3", "#c62828"},
	{"winter", "Winter", "#bcd8ee", "#2f3a45", "#eef4fb", "#ffffff"},
	{"christmas", "Christmas", "#1f7a3a", "#3b2418", "#fdf6ec", "#dc2626"},
}

func (memStore) Ping(context.Context) error                    { return errNoDB }
func (memStore) Themes(context.Context) ([]Theme, error)       { return defaultThemes, nil }
func (memStore) SaveTree(context.Context, Tree) (Tree, error)  { return Tree{}, errNoDB }
func (memStore) GetTree(context.Context, string) (Tree, error) { return Tree{}, errNoDB }
func (memStore) CountTrees(context.Context) (int, error)       { return 0, errNoDB }
