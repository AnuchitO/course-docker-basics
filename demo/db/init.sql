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
  ('spring',    'Spring',    '#2f7d32', '#6d4c41', '#f6f1e7', '#f48fb1', 1),
  ('summer',    'Summer',    '#1b5e20', '#5d4037', '#fff8e1', '#ffb300', 2),
  ('autumn',    'Autumn',    '#bf360c', '#4e342e', '#fbeee0', '#ef6c00', 3),
  ('winter',    'Winter',    '#0d47a1', '#37474f', '#eef4fb', '#90caf9', 4),
  ('christmas', 'Christmas', '#14532d', '#4e342e', '#fdf6ec', '#dc2626', 5)
ON CONFLICT DO NOTHING;
