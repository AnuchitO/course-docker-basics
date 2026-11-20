# ICQR Magic Tree (demo for the Docker Basics course)

Paste a link, get a 3D garden with a cherry-blossom tree. Tap the garden: the blossoms fly into flower beds and form a flat, scannable QR code. (Visual design ported from the "QR Garden" prototype.) Three containers:

| service | what | image |
|---|---|---|
| `web` | the page, nginx forwards `/api` to the API | `./web` |
| `api` | Go service: validates the link, builds and **verifies** the QR (level H) | `./api` |
| `db`  | Postgres: season colours, flower colours and shared gardens | `postgres:16-alpine` |

Run everything: `docker compose up --build` → http://localhost:3000

The course builds it step by step (run all commands from this `demo/` folder):

1. `docker build -t magic-web ./web && docker run -d --name web -p 3000:80 magic-web`
2. `docker build -t magic-api ./api && docker run -d --name api -p 8080:8080 magic-api`, then `curl localhost:8080/api/health`
3. `docker rm -f web api && docker network create magic`, then run `api` and `web` with `--network magic`
4. add `db` plus `-e DATABASE_URL=...` on `api`
5. add `-v tree-data:/var/lib/postgresql/data` to `db`
6. `docker compose up --build`
7. `docker compose up -d --scale api=3` and watch the `host` field change

Privacy: a link is stored only when the user presses "Copy share link".
Tests: `cd api && go test ./...`
