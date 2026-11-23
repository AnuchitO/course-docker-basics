---
theme: default
title: Docker Basics workshop
info: |
  Docker Basics workshop — by AnuchitO
colorSchema: auto
aspectRatio: 16/9
canvasWidth: 1100
class: text-center
transition: none
mdc: true
drawings:
  persist: false
fonts:
  provider: none
  sans: Inter, -apple-system, Segoe UI, Sarabun, Noto Sans Thai, sans-serif
  mono: Fira Code, JetBrains Mono, Menlo, monospace
highlighter: shiki
lineNumbers: false
codeCopy: true
duration: 90min
---

<div class="flex justify-center"><div style="position: relative">
  <div style="position: absolute; left: 50%; top: 50%; width: 26rem; height: 26rem; transform: translate(-50%, -50%); border-radius: 50%; background: radial-gradient(circle, rgba(36,150,237,.38), rgba(36,150,237,.12) 45%, transparent 70%); filter: blur(6px)"></div>
  <div class="float" style="position: relative; font-size: 9rem; line-height: 1">🐳</div>
</div></div>

<div class="flex justify-center mt-4"><span class="chip">hands-on workshop</span></div>

<h1 style="border-left: none; padding-left: 0; font-size: 4.6rem; line-height: 1.05; margin-top: 0.6rem; letter-spacing: -0.03em">Docker Basics</h1>

<div class="text-2xl font-700 mt-3">Build it once. <span class="gradient-text">Run it anywhere.</span></div>

<div class="mt-8 opacity-70">by AnuchitO</div>

<div class="abs-br m-6 text-sm opacity-50">Press <kbd>Space</kbd> to begin →</div>

---

# Before you start

<div class="grid grid-cols-[1.35fr_1fr] gap-8 mt-3">

<div class="min-w-0 flex flex-col gap-3">

<div v-click class="flex items-start gap-3">
  <span class="chip">1</span>
  <div class="min-w-0 flex-1"><b>Install Colima</b> <span class="small">(runs Docker for you, no Docker Desktop)</span>

```bash
brew install colima docker docker-compose
```

  </div>
</div>

<div v-click class="flex items-start gap-3">
  <span class="chip">2</span>
  <div class="min-w-0 flex-1"><b>Start it</b> <span class="small">(stop later with <code>colima stop</code>)</span>

```bash
colima start
```

  </div>
</div>

<div v-click class="flex items-start gap-3">
  <span class="chip">3</span>
  <div class="min-w-0 flex-1"><b>Point Docker at it</b> <span class="small">(the star <b>*</b> marks the one in use)</span>

```bash
docker context use colima
docker context ls
```

  </div>
</div>

<div v-click class="flex items-start gap-3">
  <span class="chip">4</span>
  <div class="min-w-0 flex-1"><b>Check it works</b>

```bash
docker --version
docker ps
```

  </div>
</div>

</div>

<div class="min-w-0 flex flex-col gap-4">

<div v-click class="card">
  <div class="font-700 text-lg">📁 Get the course code</div>
  <div class="opacity-60 text-sm mt-1">Download it (link under the video) and open a terminal in the <code>demo/</code> folder.</div>
</div>

<div v-click class="card">
  <div class="font-700 text-lg">⌨️ No Go needed</div>
  <div class="opacity-60 text-sm mt-1">Docker builds and runs everything. You only type commands.</div>
</div>

<div v-click class="card">
  <div class="font-700">🧩 For <code>docker compose</code> (Step 6)</div>
  <div class="opacity-60 text-sm mt-1">Link the plugin once:</div>

```bash
mkdir -p ~/.docker/cli-plugins
cd ~/.docker/cli-plugins
ln -sfn $(which docker-compose) .
```

</div>

</div>

</div>

<div v-click class="mt-5 text-center"><span class="try">⏸ Pause &amp; try</span> <span class="opacity-70 ml-2">appears on every hands-on step. Pause the video, run it, then press play.</span></div>

---
layout: center
class: text-center
---

<h1 style="border-left: none; padding-left: 0">Meet <span class="gradient-text">Magic Tree</span> 🌳</h1>

<div class="opacity-70 mt-2">Paste a link, get a QR code that grows from a tree. We will build it <b>one container at a time</b>.</div>

<div class="flex items-center justify-center gap-5 mt-10">
  <div v-click class="card text-center w-44"><div class="text-4xl">🌐</div><div class="font-700">web</div><div class="opacity-60 text-xs">the page (nginx)</div></div>
  <div v-click class="arrow">→</div>
  <div v-click class="card text-center w-44"><div class="text-4xl">⚙️</div><div class="font-700">api</div><div class="opacity-60 text-xs">makes the QR (Go)</div></div>
  <div v-click class="arrow">→</div>
  <div v-click class="card glow text-center w-44"><div class="text-4xl">🐘</div><div class="font-700">db</div><div class="opacity-60 text-xs">colours &amp; shares (Postgres)</div></div>
</div>

---

# Tour · The tree is the code

<div class="grid grid-cols-[1fr_1.1fr] gap-8 mt-2 items-center">

<div>
  <div class="text-2xl font-700 leading-snug">Type a link.<br>Watch it <span class="gradient-text">grow into a tree</span>.</div>
  <div class="mt-5 flex flex-col gap-3">
    <div v-click class="flex items-center gap-3"><span class="text-3xl">🌳</span><div><b>The crown is the code</b><div class="small">scan the canopy with your phone</div></div></div>
    <div v-click class="flex items-center gap-3"><span class="text-3xl">🦋</span><div><b>Butterflies fly around it</b><div class="small">just for fun</div></div></div>
    <div v-click class="flex items-center gap-3"><span class="text-3xl">👆</span><div><b>Drag to spin and tilt</b><div class="small">tap to change the view</div></div></div>
  </div>
</div>

<div class="flex justify-center"><img src="/img/site-butterflies.gif" style="border-radius:12px; border:1px solid var(--card-line); width: 100%; max-width: 520px" /></div>

</div>

---

# Tour · Tap to see the code

<div class="grid grid-cols-3 gap-5 mt-2 items-start">
  <div><img src="/img/site-tree.jpg" style="border-radius:12px; border:1px solid var(--card-line)" /><div class="text-center mt-2"><span class="chip">1</span> <b>Tree</b></div><div class="small text-center">where it starts</div></div>
  <div v-click><img src="/img/site-canopy.jpg" style="border-radius:12px; border:1px solid var(--card-line)" /><div class="text-center mt-2"><span class="chip">2</span> <b>Tap</b></div><div class="small text-center">the crown shows the code</div></div>
  <div v-click><img src="/img/site-qr.jpg" style="border-radius:12px; border:1px solid var(--card-line)" /><div class="text-center mt-2"><span class="chip">3</span> <b>Tap again</b></div><div class="small text-center">flat and ready to scan</div></div>
</div>

<div v-click class="mt-5 text-center text-xl">Same code in all three. <span class="gradient-text font-700">Only the look changes.</span></div>

---

# Tour · Seasons and colours

<div class="grid grid-cols-[1.3fr_1fr] gap-8 mt-2 items-center">

<img src="/img/site-autumn.jpg" style="border-radius:12px; border:1px solid var(--card-line)" />

<div class="flex flex-col gap-4">
  <div v-click class="card"><div class="font-700">🍂 Pick a season</div><div class="small">Spring, Summer or Autumn changes the whole garden.</div></div>
  <div v-click class="card"><div class="font-700">🎨 Pick a flower colour</div><div class="small">Six colours. The code stays readable.</div></div>
  <div v-click class="card"><div class="font-700">✅ No server needed</div><div class="small">The web container does all of this by itself.</div></div>
</div>

</div>

---

# Tour · Save and gallery

<div class="grid grid-cols-2 gap-6 mt-2 items-start">
  <div><img src="/img/site-saved.jpg" style="border-radius:12px; border:1px solid var(--card-line)" /><div class="text-center mt-2"><b>Save</b> 🔖 keeps your garden</div></div>
  <div v-click><img src="/img/site-gallery.jpg" style="border-radius:12px; border:1px solid var(--card-line)" /><div class="text-center mt-2"><b>Gallery</b> shows everyone's gardens, with scan counts</div></div>
</div>

<div v-click class="mt-4 text-center text-lg">These two need to <b>remember things</b>, so they need the <span class="gradient-text font-700">database</span>. That is Step 4.</div>

---
layout: center
class: text-center
---

<h1 style="border-left: none; padding-left: 0; font-size: 4.2rem; line-height: 1.1; margin: 0">Why we need Docker?</h1>

---

# The problem, in plain words

<div class="text-xl opacity-80">Same program. Different computers. <b>Different results.</b></div>

<div class="flex items-center justify-center gap-6 mt-8">

<div class="card good text-center w-52">
  <img src="/img/photo-laptop.jpg" style="border-radius:8px; height: 6.5rem; width: 100%; object-fit: cover" />
  <div class="font-700 mt-2">My laptop</div>
  <div class="label-good mt-1">✅ works</div>
</div>

<div class="arrow">→</div>

<div class="flex flex-col gap-3">
  <div v-click class="card bad flex items-center gap-4 w-96">
    <img src="/img/photo-desk.jpg" style="border-radius:8px; height: 3.2rem; width: 4.6rem; object-fit: cover; flex: none" />
    <div><b>Friend's laptop</b><div class="small">has an older version</div></div>
    <div class="label-bad ml-auto">❌ breaks</div>
  </div>
  <div v-click class="card bad flex items-center gap-4 w-96">
    <img src="/img/photo-rack.jpg" style="border-radius:8px; height: 3.2rem; width: 4.6rem; object-fit: cover; flex: none" />
    <div><b>The server</b><div class="small">is missing a piece</div></div>
    <div class="label-bad ml-auto">❌ breaks</div>
  </div>
  <div v-click class="card bad flex items-center gap-4 w-96">
    <img src="/img/photo-code.jpg" style="border-radius:8px; height: 3.2rem; width: 4.6rem; object-fit: cover; flex: none" />
    <div><b>The tester's PC</b><div class="small">has different settings</div></div>
    <div class="label-bad ml-auto">❌ breaks</div>
  </div>
</div>

</div>

<div v-click class="mt-8 text-center text-xl">Now imagine <b>100</b> computers. Nobody can keep them all the same by hand.</div>

---

# The idea: a shipping container

<div class="text-xl opacity-80">To ship anything overseas, <b>make it fit inside a standard container</b>.</div>

<div class="mt-4 flex flex-col gap-3">

<div v-click="1" class="flex items-center gap-4">
  <div class="w-24 text-right font-700 opacity-70">On the sea</div>
  <div class="card text-center w-64" style="padding: .6rem"><div class="grid grid-cols-3 gap-2"><img src="/img/photo-coffee.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img src="/img/photo-car.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img src="/img/photo-clothes.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /></div><div class="small mt-1">coffee, cars, clothes: anything</div></div>
  <div class="arrow">→</div>
  <div class="card glow w-56 text-center" style="padding: .6rem"><img src="/img/photo-containers.jpg" style="border-radius:8px; height: 3.8rem; width: 100%; object-fit: cover" /><div class="font-700 mt-1">Container</div><div class="small">same size, sealed</div></div>
  <div class="arrow">→</div>
  <div class="card text-center" style="width: 21rem; padding: .6rem">
    <div class="grid grid-cols-3 gap-2"><img v-click="2" src="/img/photo-ship.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img v-click="3" src="/img/photo-truck.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img v-click="4" src="/img/photo-train.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /></div>
    <div class="small mt-1">any ship, truck or train</div>
  </div>
</div>

<div v-click="5" class="flex items-center gap-4">
  <div class="w-24 text-right font-700" style="color: var(--blue)">Software</div>
  <div class="card text-center w-64" style="padding: .6rem"><img src="/img/photo-code.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><div class="small mt-1">your app, its tools, its settings</div></div>
  <div class="arrow">→</div>
  <div class="card glow w-56 text-center" style="padding: .6rem"><img src="/img/photo-containers.jpg" style="border-radius:8px; height: 3.8rem; width: 100%; object-fit: cover" /><div class="font-700 mt-1">Docker container</div><div class="small">same shape, sealed</div></div>
  <div class="arrow">→</div>
  <div class="card text-center" style="width: 21rem; padding: .6rem">
    <div class="grid grid-cols-3 gap-2"><img v-click="6" src="/img/photo-laptop.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img v-click="7" src="/img/photo-rack.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /><img v-click="8" src="/img/photo-dc.jpg" style="border-radius:8px; height: 5rem; width: 100%; object-fit: cover" /></div>
    <div class="small mt-1">any laptop, server or cloud</div>
  </div>
</div>

</div>

<div v-click="9" class="mt-4 grid grid-cols-2 gap-5">
  <div class="card"><b>The crane doesn't care what's inside.</b><div class="small">It only needs the box to be the standard size.</div></div>
  <div class="card"><b>🐳 Docker is the ship.</b><div class="small">Any computer with Docker can carry and run any container.</div></div>
</div>

<div v-click="10" class="mt-3 text-center big-line">Fit it in the box once. <span class="gradient-text">Ship it anywhere.</span></div>

<div class="abs-bl m-2 text-xs opacity-50">Photos via Wikimedia Commons: Petar Milošević (CC BY-SA 4.0), Virtual-Pano (CC BY 4.0), PiDatacenters (CC BY-SA 4.0); the rest CC0 (Wikideas1, Unsplash photographers)</div>

---

# On a host machine

<Laptop name="Host Machine" width="640px" class="mt-4">
  <div class="flex items-center gap-6">
    <div class="flex flex-col gap-2">
      <div v-click class="chip">📄 Dockerfile</div>
      <div v-click class="chip">💿 image</div>
    </div>
    <div v-click class="arrow">→</div>
    <div class="grid grid-cols-3 gap-3">
      <Ctr v-for="n in 12" :key="n" v-click />
    </div>
  </div>
</Laptop>

<div v-click class="mt-6 text-center opacity-70">The host stays clean — everything lives inside isolated containers.</div>

---

# Everyone, same container

<div class="grid grid-cols-3 gap-x-8 gap-y-12 mt-3">
  <Laptop width="225px" v-for="(p, i) in [
    { n: 'MeeN',     skin: '#f1c27d', hair: '#2b2118', shirt: '#2563eb', style: 'short' },
    { n: 'Ika',      skin: '#e0ac69', hair: '#1c1917', shirt: '#db2777', style: 'long' },
    { n: 'Pi',       skin: '#c68642', hair: '#292524', shirt: '#16a34a', style: 'cap' },
    { n: 'AnuchitO', skin: '#f1c27d', hair: '#1c1917', shirt: '#7c3aed', style: 'short' },
    { n: 'Tua',      skin: '#8d5524', hair: '#0c0a09', shirt: '#d97706', style: 'curly' },
    { n: 'Sonali',   skin: '#e0ac69', hair: '#3b2a1e', shirt: '#0891b2', style: 'bun' },
  ]" :key="p.n"
          v-motion :initial="{ y: 30, opacity: 0 }" :enter="{ y: 0, opacity: 1, transition: { delay: 120 * i } }"
          :style="{ '--d': (0.6 + 0.12 * i) + 's' }">
    <template #person><Person :skin="p.skin" :hair="p.hair" :shirt="p.shirt" :hair-style="p.style" /></template>
    <div class="chip">💿 image</div><Ctr />
  </Laptop>
</div>

<div v-click class="mt-4 text-center text-xl">Six laptops · one image · <span class="gradient-text font-700">identical behaviour</span></div>

---
layout: center
---

# Docker life cycle

<div class="text-xl opacity-80">Three steps. <b>Write it, pack it, start it.</b></div>

<div class="grid grid-cols-[1fr_auto_1fr_auto_1fr] items-stretch gap-2 mt-6">

  <div v-click class="card text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">STEP 1</div>
    <div class="text-3xl font-800 mt-1">Dockerfile</div>
    <div class="h-24 flex items-center justify-center"><DockerfileIcon /></div>
    <div class="text-sm">List what your app needs, in a <b>text file</b>.</div>
    <div class="mt-auto pt-3"><span class="chip">Write</span></div>
  </div>

  <div v-click class="text-center self-center" style="width: 8.5rem">
    <div class="small">then</div>
    <div class="chip mt-1" style="font-size: 1rem; border-color: var(--m)">docker build</div>
    <div class="arrow">→</div>
  </div>

  <div v-click class="card text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">STEP 2</div>
    <div class="text-3xl font-800 mt-1">Image</div>
    <div class="h-24 flex items-center justify-center"><img src="/img/photo-containers.jpg" style="border-radius:8px; height: 5.5rem; width: 100%; object-fit: cover" /></div>
    <div class="text-sm">Everything goes into <b>one sealed box</b>.</div>
    <div class="mt-auto pt-3"><span class="chip">Pack</span></div>
  </div>

  <div v-click class="text-center self-center" style="width: 8.5rem">
    <div class="small">then</div>
    <div class="chip mt-1" style="font-size: 1rem; border-color: var(--m)">docker run</div>
    <div class="arrow">→</div>
  </div>

  <div v-click class="card glow text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">STEP 3</div>
    <div class="text-3xl font-800 mt-1">Container</div>
    <div class="h-24 flex items-center justify-center gap-1" style="zoom: .7"><Ctr v-for="n in 3" :key="n" /></div>
    <div class="text-sm">Open the box and your app <b>runs</b>. Open as many as you like.</div>
    <div class="mt-auto pt-3"><span class="chip">Start</span></div>
  </div>

</div>

<div v-click class="mt-8 text-center text-xl">
  <b>Dockerfile</b> = the instructions &nbsp;·&nbsp; <b>Image</b> = the packed box &nbsp;·&nbsp; <b>Container</b> = the box, running
</div>

---

# Image <span class="opacity-50">?</span>

<div class="grid grid-cols-2 gap-12 mt-4 items-center">

<div>

<div class="text-2xl font-700 leading-snug">A <span class="gradient-text">recipe</span> your app needs,<br>baked into <b>layers</b>.</div>

<div class="mt-6 flex flex-col gap-4">
  <div v-click="5" class="flex items-center gap-3"><span class="text-3xl">🔒</span><div><b>Read-only</b><div class="small">a finished image never changes</div></div></div>
  <div v-click="6" class="flex items-center gap-3"><span class="text-3xl">📤</span><div><b>Shareable</b><div class="small">push it once, pull it anywhere</div></div></div>
  <div v-click="7" class="flex items-center gap-3"><span class="text-3xl">♻️</span><div><b>Reusable</b><div class="small">one image → as many containers as you like</div></div></div>
</div>

</div>

<div class="relative">
  <div v-click="5" class="image-tag"><span>💿</span> magic-api:1.0</div>
  <div class="image-stack" :class="{ sealed: $clicks >= 5 }">
    <div v-click="1" class="layer l1"><span class="who">base OS + Go</span><code>FROM golang:1.25-alpine</code></div>
    <div v-click="2" class="layer l2"><span class="who">dependencies</span><code>RUN go mod download</code></div>
    <div v-click="3" class="layer l3"><span class="who">your code</span><code>COPY . .</code></div>
    <div v-click="4" class="layer l4"><span class="who">build</span><code>RUN go build -o api .</code></div>
  </div>
</div>

</div>

---

# Dockerfile <span class="opacity-50">?</span>

A text file of instructions that **builds an image**.

````md magic-move {lines: true}
```dockerfile
FROM <image>
```

```dockerfile
FROM <image>
RUN <command>
```

```dockerfile
FROM golang:1.25-alpine
RUN <command>
```

```dockerfile
FROM golang:1.25-alpine
RUN go mod download
```

```dockerfile
FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go build -o api .
CMD ["./api"]
```
````

<div v-click class="mt-4 text-center opacity-70">each instruction → one layer of the <b>image</b></div>

---
layout: center
class: text-center
---

# Choose the right tool

<div class="grid grid-cols-3 gap-6 mt-10">
  <div v-click class="card"><div class="text-4xl">🟩</div><div class="font-700 mt-1">FROM node</div><div class="opacity-60 text-sm">JavaScript</div></div>
  <div v-click class="card"><div class="text-4xl">🐍</div><div class="font-700 mt-1">FROM python</div><div class="opacity-60 text-sm">Data / scripts</div></div>
  <div v-click class="card"><div class="text-4xl">🐹</div><div class="font-700 mt-1">FROM golang</div><div class="opacity-60 text-sm">Services</div></div>
</div>

<div v-click class="mt-8 opacity-70">Pick the base image that fits your job — then add what you need on top.</div>

---

# Container <span class="opacity-50">?</span>

<div class="flex items-center justify-center gap-8 mt-2">
  <div class="card glow text-center"><div class="text-5xl">💿</div><div class="font-700">image</div></div>
  <div class="arrow">→</div>
  <ContainerBig name="magic-api" :open="$clicks >= 1" :width="400" />
</div>

<div v-click="1" class="mt-4 text-center"><span class="chip">docker run</span> — the doors open and the app starts</div>

<div v-click="2" class="flex items-center justify-center gap-4 mt-4">
  <span class="opacity-70">same image, more containers →</span>
  <Ctr v-for="n in 4" :key="n" />
</div>

<div v-click="3" class="mt-6 text-center text-xl">An image is the <b>class</b>, a container is an <b>instance</b>.</div>

---

# Hello Docker

<div class="grid grid-cols-[1fr_1fr] gap-8 mt-4 items-start">

<div>

<div class="text-lg mb-2">Run your very first container:</div>

```bash
docker run hello-world
```

<div class="mt-4 small">Docker finds the image, starts a container, prints a message, and the container <b>stops</b>: it had nothing else to do.</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out" style="font-size: 1rem; word-break: normal">Hello from Docker!<br>This message shows that your installation appears to be working correctly.</div>
  <div class="mt-3 text-sm opacity-70">Your container ran and finished in under a second.</div>
</div>

</div>

---

# Containers are isolated

<div class="text-lg opacity-80">Same image, two containers. <b>Changing one does not touch the other.</b></div>

<div class="grid grid-cols-[1.1fr_1fr] gap-6 mt-3">

<div class="min-w-0 flex flex-col gap-2">

  <div v-click="1">
    <div class="small mb-1"><b>1</b> · Start two containers from the same image</div>

```bash
docker run -d --name tesla \
  alpine sleep 3600
docker run -d --name byd \
  alpine sleep 3600
```

  </div>

  <div v-click="2">
    <div class="small mb-1"><b>2</b> · Write a file inside <code>tesla</code> only</div>

```bash
docker exec tesla sh -c \
  "echo hi > /note.txt"
```

  </div>

  <div v-click="3">
    <div class="small mb-1"><b>3</b> · Look for the file in each one</div>

```bash
docker exec tesla cat /note.txt
docker exec byd cat /note.txt
```

  </div>

</div>

<div class="min-w-0 flex flex-col items-center gap-2">

  <div class="card glow text-center" style="padding: .5rem 1.4rem"><span class="text-3xl">💿</span> <b>image: alpine</b></div>
  <div class="arrow" style="transform: rotate(90deg); animation: none">→</div>

  <div class="grid grid-cols-2 gap-3 w-full">
    <div v-click="1" class="card text-center">
      <div class="font-700">container <code>tesla</code></div>
      <div v-click="2" class="mt-2 text-3xl">📄</div>
      <div v-click="2" class="small"><code>note.txt</code></div>
      <div v-click="3" class="label-good mt-1">✅ finds it</div>
    </div>
    <div v-click="1" class="card text-center">
      <div class="font-700">container <code>byd</code></div>
      <div v-click="2" class="mt-2 text-3xl opacity-30">🫙</div>
      <div v-click="2" class="small">untouched</div>
      <div v-click="3" class="label-bad mt-1">❌ not found</div>
    </div>
  </div>

</div>

</div>

<div v-click="3" class="mt-3 text-center text-lg">Each container has its <b>own private space</b>. Nothing leaks between them.</div>

---

# Containers are stateless

<div class="text-lg opacity-80">Remove a container and start a new one. <b>It starts clean, every time.</b></div>

<div class="grid grid-cols-[1.1fr_1fr] gap-6 mt-3">

<div class="min-w-0 flex flex-col gap-2">

  <div v-click="1">
    <div class="small mb-1"><b>4</b> · Throw <code>tesla</code> away</div>

```bash
docker rm -f tesla
```

  </div>

  <div v-click="2">
    <div class="small mb-1"><b>5</b> · Start a new <code>tesla</code> from the same image</div>

```bash
docker run -d --name tesla \
  alpine sleep 3600
```

  </div>

  <div v-click="3">
    <div class="small mb-1"><b>6</b> · Look for the file again</div>

```bash
docker exec tesla cat /note.txt
```

  </div>

</div>

<div class="min-w-0 flex flex-col items-center gap-2">

  <div class="card glow text-center" style="padding: .5rem 1.4rem"><span class="text-3xl">💿</span> <b>image: alpine</b> <span class="small">(never changed)</span></div>

  <div class="grid grid-cols-[1fr_auto_1fr] items-center gap-2 w-full mt-2">
    <div class="card text-center">
      <div class="font-700">old <code>tesla</code></div>
      <div class="text-3xl mt-1">📄</div>
      <div class="small"><code>note.txt</code></div>
    </div>
    <div v-click="1" class="text-3xl">🗑️</div>
    <div v-click="2" class="card text-center">
      <div class="font-700">new <code>tesla</code></div>
      <div class="text-3xl mt-1 opacity-30">🫙</div>
      <div v-click="3" class="label-bad mt-1">❌ file gone</div>
    </div>
  </div>

</div>

</div>

<div v-click="3" class="mt-3 text-center text-lg">The image is the <b>same every time</b>, so the container starts the <b>same every time</b>. Nothing is saved.</div>

<div v-click="3" class="mt-1 text-center small">Need to keep data? That is what a <b>volume</b> is for (Step 5). Clean up: <code>docker rm -f tesla byd</code></div>

---


# Keep it running

<div class="grid grid-cols-2 gap-8 mt-2">

<div class="min-w-0">

<div class="small mb-2"><code>hello-world</code> stopped by itself: it had nothing left to do. A server keeps working, so run it <b>in the background</b>.</div>

```bash {1|2|3|4|5}
docker run -d --name web -p 8080:80 nginx
docker ps
docker logs web
docker stop web
docker start web
```

<div class="mt-3"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div class="min-w-0">

<div class="card">
  <div><code>-d</code> <b>detached</b>: runs in the background, you get your terminal back</div>
  <div class="mt-2"><code>docker ps</code> lists what is running</div>
  <div class="mt-2"><code>docker stop</code> / <code>start</code> pause and resume the same container</div>
</div>

<div v-click class="see mt-4">
  <div class="h">YOU SHOULD SEE</div>
  Open <b>http://localhost:8080</b>: "Welcome to nginx!"
</div>

<div v-click class="mt-4">
  <div class="small mb-1">Come back after a reboot or crash:</div>

```bash
docker run -d --restart unless-stopped \
  --name web2 -p 8081:80 nginx
```

</div>

</div>

</div>

---

# Step 1 · Run your first container

<Stepper :n="1" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

```bash {1|2}
docker build -t magic-web ./web
docker run -d --name web -p 3000:80 magic-web
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
<code>-d</code> run in the background · <code>--name</code> give it a name<br>
<code>-p 3000:80</code> <b>host</b> port 3000 → <b>container</b> port 80
</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  Open <b>http://localhost:3000</b>
  <div class="mt-2">A cherry-blossom tree appears. Type a link and the QR code and the tree work, with 3 flower colours. Now press <b>Save</b> or open the <b>Gallery</b>:</div>
  <div class="out mt-1">API not reachable</div>
  <div class="mt-3 text-sm opacity-70">One container works alone. Anything that needs to <b>remember</b> or <b>count</b> has nobody to ask yet.</div>
</div>

</div>

---

# Step 2 · A second container: the API

<Stepper :n="2" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

```bash {1|2|3}
docker build -t magic-api ./api
docker run -d --name api -p 8080:8080 magic-api
curl localhost:8080/api/health
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
The Go API starts on its own, with built-in themes and <b>no database</b>.
</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out">{"db":"not connected",<br>"host":"bd9a1a742cb4",<br>"status":"ok"}</div>
  <div class="mt-3 text-sm opacity-70"><code>host</code> is the container ID. Remember it: we use it in step 7.</div>
  <div class="mt-2 text-sm opacity-70">Save at :3000 still says “API not reachable”: containers are <b>isolated</b>.</div>
</div>

</div>

---

# Step 3 · Connect them with a network

<Stepper :n="3" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.25fr_1fr] gap-8 mt-4">

<div>

```bash {1|2|3|4}
docker rm -f web api
docker network create magic
docker run -d --name api --network magic magic-api
docker run -d --name web --network magic -p 3000:80 magic-web
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
On the same network, <b>web</b> reaches the API by its name: <code>http://api:8080</code>.<br>
The API has no <code>-p</code>: only <b>web</b> is open to the outside.
</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  Refresh <b>http://localhost:3000</b>
  <div class="mt-2">🌳 The garden now carries a <b>QR code</b> under the tree. <b>Tap the garden</b>: the blossoms fly into flower beds and the code appears, flat and scannable.</div>
  <div class="mt-3 text-sm opacity-70">The header shows <b>API &lt;container id&gt;</b>. Seasons work; the 🔗 link button says it needs a database.</div>
</div>

</div>

---

# Step 4 · Add the database

<Stepper :n="4" class="absolute top-6 right-10" />

<div class="grid grid-cols-1 gap-3 mt-2">

<div>

```bash {1-4|5-8}
docker run -d --name db --network magic \
  -e POSTGRES_USER=magic -e POSTGRES_PASSWORD=magic -e POSTGRES_DB=magictree \
  -v "$PWD/db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro" \
  postgres:16-alpine
docker rm -f api
docker run -d --name api --network magic \
  -e DATABASE_URL="postgres://magic:magic@db:5432/magictree?sslmode=disable" \
  magic-api
```

<div class="mt-3 text-sm opacity-80 leading-relaxed">
<code>-e</code> sets an <b>environment variable</b>: that is how the API learns where the database is.
</div>

<div class="mt-3"><span class="try">⏸ Pause &amp; try</span> <span class="small ml-2">wait ~5 seconds, then</span> <code>curl localhost:3000/api/health</code></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out">{"db":"connected",<br>"shared_trees":0, ...}</div>
  <div class="mt-3">The colour dots now come from <b>Postgres</b>: <b>6 flower colours</b> instead of 3. Press the 🔗 button: a share link works.</div>
</div>

</div>

---

# Step 5 · Keep the data: a volume

<Stepper :n="5" class="absolute top-6 right-10" />

<div class="grid grid-cols-2 gap-8 mt-4">

<div class="card bad">
  <div class="label-bad">❌ Without a volume</div>

```bash
docker rm -f db
# start db again (same command as step 4)
```

<div class="text-sm opacity-80">Delete the container, and the shared trees are <b>gone</b>: <code>shared_trees: 0</code>.</div>
</div>

<div v-click class="card good">
  <div class="label-good">✅ With a volume</div>

```bash {1|3}
docker rm -f db
docker run -d --name db --network magic \
  -v tree-data:/var/lib/postgresql/data \
  ...same as step 4
docker restart api
```

<div class="text-sm opacity-80">Delete and recreate: the trees are <b>still there</b>.</div>
</div>

</div>

<div v-click class="mt-5 text-center"><span class="try">⏸ Pause &amp; try</span> <span class="opacity-70 ml-2">share a garden, remove <code>db</code>, start it again, then check <code>/api/health</code>.</span></div>

---

# Step 6 · One command: Docker Compose

<Stepper :n="6" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.3fr_1fr] gap-6 mt-2">

<div class="compact">

```yaml {1-5|6-11|12-22}
services:
  web:
    build: ./web
    ports: ["3000:80"]
    depends_on: [api]
  api:
    build: ./api
    environment:
      DATABASE_URL: postgres://magic:magic@db:5432/magictree?sslmode=disable
    depends_on: { db: { condition: service_healthy } }
  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: magic
      POSTGRES_PASSWORD: magic
      POSTGRES_DB: magictree
    volumes:
      - tree-data:/var/lib/postgresql/data
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro
volumes:
  tree-data:
```

</div>

<div>

<div class="card">
  <div class="text-sm opacity-70">First, clean up the manual containers:</div>

```bash
docker rm -f web api db
docker network rm magic
```

  <div class="text-sm opacity-70 mt-2">Then everything, with one command:</div>

```bash
docker compose up --build
```

</div>

<div v-click class="mt-3 text-center"><span class="try">⏸ Pause &amp; try</span></div>

</div>

</div>

---

# Step 7 · Scale the API

<Stepper :n="7" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

```bash {1|2-3}
docker compose up -d --scale api=3
for i in 1 2 3 4 5 6; do
  curl -s localhost:3000/api/health; echo; done
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
Three copies of the <b>same image</b>. Compose gives them one name, <code>api</code>, and requests are spread between them.
</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span> <span class="small ml-2">finish with</span> <code>docker compose down</code></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out">"host":"4d633404bd51"<br>"host":"a7a19b993e2b"<br>"host":"dc88b4d8a103"<br>...</div>
  <div class="flex gap-2 mt-3"><Ctr v-for="n in 3" :key="n" /></div>
  <div class="mt-2 text-sm opacity-70">Three different container IDs: one image, three containers.</div>
</div>

</div>

---
layout: center
---

# What you built

<div class="grid grid-cols-2 gap-4 mt-6 text-left">
  <div v-click class="card"><b>1 · docker run</b><div class="small">starts a container from an image</div></div>
  <div v-click class="card"><b>2 · -p host:container</b><div class="small">opens a door to the outside</div></div>
  <div v-click class="card"><b>3 · network</b><div class="small">containers find each other by name</div></div>
  <div v-click class="card"><b>4 · -e variable</b><div class="small">config without changing the image</div></div>
  <div v-click class="card"><b>5 · volume</b><div class="small">data outlives the container</div></div>
  <div v-click class="card"><b>6 · compose</b><div class="small">the whole stack in one file</div></div>
</div>

---
layout: center
class: text-center
---

# Thank you 🐳

<div class="mt-8 grid grid-cols-2 gap-4 text-left text-lg">
  <a v-click href="https://github.com/AnuchitO" target="_blank" class="card">🐙 GitHub <div class="opacity-60 text-sm">AnuchitO</div></a>
  <a v-click href="https://www.facebook.com/AnuchitO.rice" target="_blank" class="card">📘 Facebook page <div class="opacity-60 text-sm">AnuchitO.rice</div></a>
  <a v-click href="https://www.youtube.com/AnuchitO" target="_blank" class="card">▶️ YouTube <div class="opacity-60 text-sm">AnuchitO</div></a>
  <a v-click href="https://dev.to/anuchito" target="_blank" class="card">✍️ Dev.to <div class="opacity-60 text-sm">anuchito</div></a>
</div>
