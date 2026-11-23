# Demo script: from one container to a full stack

Walk through it commit by commit (`git log --reverse --oneline`). One commit = one step = one slide (Step 1 to Step 7; Step 1b is a short extra step right after Step 1).
Run every command from this `demo/` folder.
Each step says what it is **for**, what to **run** and what you should **see**.
Only Step 3 fails on purpose, so that Step 4 has a reason to exist.

---

## Step 1: just serve index.html

**Intent:** the smallest possible container. nginx serves one static page. No API yet.

```bash
docker build -t magic-web ./web
docker run -d --name web -p 3000:80 magic-web
```

**See:** http://localhost:3000 shows the garden and the QR code works (the page builds it in the browser, because there is no API to ask).
`curl -i localhost:3000/api/health` is `404 Not Found`: nginx has no such file.

---

## Step 1b: change the page, then rebuild the image

**Intent:** show that an image is a snapshot. Editing `index.html` does not change a container that is already built and running. New code needs a **new image**.

The `Step 1b` commit adds a 4th flower colour in `web/index.html`. Check it out, or make the same edit by hand:

```diff
- let SWATCHES = ['#e98aab','#a57ad6','#e0473c'];
+ let SWATCHES = ['#e98aab','#a57ad6','#e0473c','#f5b301'];
```

```bash
docker restart web                         # run the old image again
docker build -t magic-web:v2 ./web          # build a NEW image
docker rm -f web
docker run -d --name web -p 3000:80 magic-web:v2
docker images magic-web
```

**See:** after `docker restart web`, http://localhost:3000 still has **3** colours (the file was copied into the image at build time). After the rebuild and `magic-web:v2`, it has **4**. `docker images magic-web` lists two images, `latest` (old) and `v2` (new). Save still says "API not reachable": that is why Step 2 adds the API.

---

---

## Step 2: a second container, the API

**Intent:** a Go service builds and verifies the QR code. Run it as another container and publish port 8080 to your laptop.

```bash
docker build -t magic-api ./api
docker run -d --name api -p 8080:8080 magic-api
curl localhost:8080/api/health
```

**See:** `{"db":"not connected","host":"<container id>","status":"ok"}`. The API runs without a database for now.
The web page does not use it: `curl -i localhost:3000/api/health` is still `404`. The two containers don't know each other.

---

## Step 3: point the web at the API port **(fails on purpose)**

**Intent:** the API answers on `localhost:8080`, so tell the web container to forward `/api` there. nginx reads `API_URL` from the Dockerfile.

```bash
docker rm -f web
docker build -t magic-web ./web
docker run -d --name web -p 3000:80 magic-web
curl -i localhost:3000/api/health
docker logs web | tail -3
```

**See:** `502 Bad Gateway`, and the log says `connect() failed (111: Connection refused)`.

**Why:** `localhost` inside the web container is the web container itself, not your laptop and not the API container. Every container has its own `localhost`. Port 8080 is open on your laptop, but nothing listens on 8080 inside the web container.

**Could the published port work?** Yes, but by another name. `API_URL=http://host.docker.internal:8080` reaches the port published with `-p 8080:8080` (tested on Docker Desktop; Linux needs `--add-host=host.docker.internal:host-gateway`). The same goes for the database: with `-p 5432:5432`, the API connects through `host.docker.internal:5432`, while `localhost:5432` is refused.
It works on a laptop, but every service has to publish a port (even the database), the call leaves the container and comes back through the host, and only one container can publish a given port, so it can't be scaled. The network and alias in Step 4 avoid all of that.

---

## Step 4: Docker Compose, a private network and an alias

**Intent:** stop chaining `docker run` flags. One file describes both containers and puts them on a private network `magic`. The API gets the **alias** `api`, so the web can reach it as `http://api:8080`.
Two matching edits in `web/`: `API_URL` changes from `http://localhost:8080` to `http://api:8080`, and nginx asks Docker's DNS for `api` on every request (so the API container can be recreated without restarting the web).
The API has no `ports:` any more: only the web is open to your laptop.

```bash
git diff HEAD~1 -- web/                # review what changed
docker rm -f web api
docker compose up -d --build
docker compose ps
curl localhost:3000/api/health
```

**See:** both services `running`, and `{"db":"not connected","host":"<api container id>","status":"ok"}` through port 3000.
`curl localhost:8080/api/health` now fails: the API is private. Open http://localhost:3000: the QR code is built and verified by the API.

---

## Step 5: add the database, then lose the data **(fails on purpose)**

**Intent:** add Postgres as a third service. The API gets `DATABASE_URL` (host `db` = the service name) and waits for the database healthcheck. `db/init.sql` is mounted so the first start creates the tables and ten example gardens.
The gallery and the save button start working. But the data lives inside the db container.

```bash
docker compose up -d
curl localhost:3000/api/health
# open http://localhost:3000, press the bookmark button to save a garden, open the gallery and find it
curl localhost:3000/api/health          # one more tree
docker compose down
docker compose up -d
curl localhost:3000/api/health
```

**See:** first `"db":"connected","shared_trees":10`, then `11` after saving.
After `down` + `up` the saved garden is **gone**: `shared_trees` is back to `10`.

**Why:** `docker compose down` removes the db container, and the data lived inside it. We need a place for the data that is not the container: a volume.

---

## Step 6: keep the data with a named volume

**Intent:** keep Postgres' files in a named volume `tree-data` (`/var/lib/postgresql` for the Postgres 18 image). The volume belongs to Docker, not to the container, so removing the container does not remove the data.

```bash
docker compose down -v          # start clean: also removes the volume left over from Step 5
docker compose up -d
curl localhost:3000/api/health  # shared_trees: 10
# save a garden in the browser (bookmark button), or:
curl -s -XPOST localhost:3000/api/trees -H 'content-type: application/json' \
  -d '{"url":"https://example.com","season":"spring","swatch":1,"name":"my tree"}'
curl localhost:3000/api/health  # shared_trees: 11
docker compose down             # no -v: keep the volume
docker compose up -d
curl localhost:3000/api/health  # still 11
docker volume ls | grep tree-data
```

**See:** the saved garden is **still there** after `down` + `up`: `shared_trees` stays `11`, and `docker volume ls` lists `<project>_tree-data`.
Careful: `docker compose down -v` deletes the volume, and the data with it.
