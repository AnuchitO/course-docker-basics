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

ALTER TABLE trees ADD COLUMN IF NOT EXISTS name text NOT NULL DEFAULT '';
ALTER TABLE trees ADD COLUMN IF NOT EXISTS likes int NOT NULL DEFAULT 0;

INSERT INTO seasons (slug, name, background, position) VALUES
  ('spring', 'Spring', '#f5efe6', 1),
  ('summer', 'Summer', '#eef2e4', 2),
  ('autumn', 'Autumn', '#f6ece0', 3)
ON CONFLICT DO NOTHING;

INSERT INTO swatches (hex, position) VALUES
  ('#e98aab', 1), ('#a57ad6', 2), ('#e0473c', 3),
  ('#e9a426', 4), ('#3d8fe6', 5), ('#2fae9b', 6)
ON CONFLICT DO NOTHING;

INSERT INTO trees (id, name, url, season, swatch, likes) VALUES
  ('demo0001', 'Docker docs',      'https://docs.docker.com',                         'spring', 3, 42),
  ('demo0002', 'Go playground',    'https://go.dev/play',                             'summer', 5, 31),
  ('demo0003', 'Cherry blossom',   'https://example.com/hanami',                      'spring', 0, 57),
  ('demo0004', 'Autumn red',       'https://example.com/momiji',                      'autumn', 2, 18),
  ('demo0005', 'Wi-Fi: garden',    'WIFI:T:WPA;S:garden-guest;P:blossom2026;;',       'summer', 4, 12),
  ('demo0006', 'Say hello',        'Hello from the Docker Basics course!',            'spring', 1, 25),
  ('demo0007', 'Postgres docs',    'https://www.postgresql.org/docs/16/',             'autumn', 3, 9),
  ('demo0008', 'Compose reference','https://docs.docker.com/compose/',                'summer', 0, 36),
  ('demo0009', 'Slidev',           'https://sli.dev',                                 'autumn', 4, 14),
  ('demo0010', 'Mail me',          'mailto:hello@example.com',                        'spring', 5, 6)
ON CONFLICT DO NOTHING;
