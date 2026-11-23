# Demo script: from one container to a full stack

Walk through it commit by commit (`git log --reverse --oneline`). One commit = one step = one slide (Step 1 to Step 7).
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
