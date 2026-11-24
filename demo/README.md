# MagicTree (demo for the Docker Basics course)

Paste a link, get a 3D cherry-blossom tree. Tap the scene: the blossoms fly into flower beds and form a flat, scannable QR code. Three containers:

| service | what | image |
|---|---|---|
| `web` | the page, nginx forwards `/api` to the API | `./web` |
| `api` | Go service: validates the link, builds and **verifies** the QR (level H) | `./api` |
| `db`  | Postgres: season colours, flower colours and shared gardens | `postgres:16-alpine` |

Run everything: `docker compose up --build` → http://localhost:3000

The course builds it step by step (run all commands from this `demo/` folder):

1. `docker build -t tree-web ./web && docker run -d --name web -p 3000:80 tree-web`
2. `docker build -t tree-api ./api && docker run -d --name api -p 8080:8080 tree-api`, then `curl localhost:8080/api/health`
3. `docker rm -f web api && docker network create tree`, then run `api` and `web` with `--network tree`
4. add `db` plus `-e DATABASE_URL=...` on `api`
5. add `-v tree-data:/var/lib/postgresql/data` to `db`
6. `docker compose up --build`
7. `docker compose up -d --scale api=3` and watch the `host` field change

Gallery: the bookmark button saves the current garden (text, season, colour) to Postgres; the grid button opens a Pinterest-style gallery (masonry cards with a mini garden thumbnail, heart counter, New/Popular and season filters; ten presets are seeded). Click a card to plant it. Nothing is stored unless you press bookmark. A saved garden’s QR encodes `<site>/api/r/<id>`: the API counts the scan, then redirects (web links) or shows the text; the gallery cards show the count. After upgrading, run `docker compose down -v` once so db/init.sql re-runs (it adds the `name` column).
Tests: `cd api && go test ./...`
