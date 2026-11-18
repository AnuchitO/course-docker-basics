package main

import "context"

// memStore is used when DATABASE_URL is not set. It serves built-in themes,
// so the API container can run on its own (lesson step 2).
// Sharing needs a real database, so it says so.
type memStore struct{}

var defaultThemes = []Theme{
	{"spring", "Spring", "#2f7d32", "#6d4c41", "#f6f1e7", "#f48fb1"},
	{"summer", "Summer", "#1b5e20", "#5d4037", "#fff8e1", "#ffb300"},
	{"autumn", "Autumn", "#bf360c", "#4e342e", "#fbeee0", "#ef6c00"},
	{"winter", "Winter", "#0d47a1", "#37474f", "#eef4fb", "#90caf9"},
	{"christmas", "Christmas", "#14532d", "#4e342e", "#fdf6ec", "#dc2626"},
}

func (memStore) Ping(context.Context) error                    { return errNoDB }
func (memStore) Themes(context.Context) ([]Theme, error)       { return defaultThemes, nil }
func (memStore) SaveTree(context.Context, Tree) (Tree, error)  { return Tree{}, errNoDB }
func (memStore) GetTree(context.Context, string) (Tree, error) { return Tree{}, errNoDB }
func (memStore) CountTrees(context.Context) (int, error)       { return 0, errNoDB }
