CREATE TABLE IF NOT EXISTS themes (
  slug       text PRIMARY KEY,
  name       text NOT NULL,
  foliage    text NOT NULL,
  trunk      text NOT NULL,
  background text NOT NULL,
  accent     text NOT NULL,
  position   int  NOT NULL
);

CREATE TABLE IF NOT EXISTS trees (
  id         text PRIMARY KEY,
  url        text NOT NULL,
  theme      text NOT NULL,
  foliage    text NOT NULL,
  trunk      text NOT NULL,
  background text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO themes (slug, name, foliage, trunk, background, accent, position) VALUES
  ('spring', 'Spring', '#f4a6c0', '#4a3426', '#e6f4d4', '#ffffff', 1),
  ('summer', 'Summer', '#2e9e57', '#3b2d22', '#e3f6e6', '#a7f3d0', 2),
  ('autumn', 'Autumn', '#e8761c', '#4a2c1c', '#f7ead3', '#c62828', 3),
  ('winter', 'Winter', '#bcd8ee', '#2f3a45', '#eef4fb', '#ffffff', 4),
  ('christmas', 'Christmas', '#1f7a3a', '#3b2418', '#fdf6ec', '#dc2626', 5)
ON CONFLICT DO NOTHING;
