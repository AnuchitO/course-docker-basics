CREATE TABLE IF NOT EXISTS seasons (
  slug       text PRIMARY KEY,
  name       text NOT NULL,
  background text NOT NULL,
  position   int  NOT NULL
);

CREATE TABLE IF NOT EXISTS swatches (
  hex      text PRIMARY KEY,
  position int NOT NULL
);

CREATE TABLE IF NOT EXISTS trees (
  id         text PRIMARY KEY,
  url        text NOT NULL,
  season     text NOT NULL,
  swatch     int  NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO seasons (slug, name, background, position) VALUES
  ('spring', 'Spring', '#f5efe6', 1),
  ('summer', 'Summer', '#eef2e4', 2),
  ('autumn', 'Autumn', '#f6ece0', 3)
ON CONFLICT DO NOTHING;

INSERT INTO swatches (hex, position) VALUES
  ('#e98aab', 1), ('#a57ad6', 2), ('#e0473c', 3),
  ('#e9a426', 4), ('#3d8fe6', 5), ('#2fae9b', 6)
ON CONFLICT DO NOTHING;
