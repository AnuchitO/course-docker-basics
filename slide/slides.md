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
  <div class="font-700">🧩 For <code>docker compose</code> (Step 4)</div>
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

<div v-click class="mt-4 text-center text-lg">These two need to <b>remember things</b>, so they need the <span class="gradient-text font-700">database</span>. That is Step 5.</div>

---

# Meet the app: three containers

<div class="text-xl opacity-80">Paste a link, get a QR code that grows from a tree. <b>Each job runs in its own container.</b></div>

<div class="grid grid-cols-[auto_auto_1fr_auto_1fr_auto_1fr] items-center gap-2 mt-8">

  <div class="text-center" style="width: 6rem">
    <div class="text-5xl">🧑</div>
    <div class="font-700 mt-1">You</div>
    <div class="small">in a browser</div>
  </div>

  <div v-click class="arrow">→</div>

  <div v-click class="card text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">CONTAINER 1</div>
    <div class="text-4xl mt-1">🌐</div>
    <div class="text-2xl font-800">web</div>
    <div class="text-sm mt-1"><b>Shows the page</b> and the growing tree.</div>
    <div class="mt-auto pt-3"><span class="chip">nginx</span></div>
  </div>

  <div v-click class="arrow">→</div>

  <div v-click class="card text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">CONTAINER 2</div>
    <div class="text-4xl mt-1">⚙️</div>
    <div class="text-2xl font-800">api</div>
    <div class="text-sm mt-1"><b>Works out the QR code</b> from your link.</div>
    <div class="mt-auto pt-3"><span class="chip">Go</span></div>
  </div>

  <div v-click class="arrow">→</div>

  <div v-click class="card glow text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">CONTAINER 3</div>
    <div class="text-4xl mt-1">🐘</div>
    <div class="text-2xl font-800">db</div>
    <div class="text-sm mt-1"><b>Remembers</b> saved trees and the gallery.</div>
    <div class="mt-auto pt-3"><span class="chip">Postgres</span></div>
  </div>

</div>

<div v-click class="mt-8 text-center big-line">We will build it <span class="gradient-text">one container at a time</span>.</div>

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
    <div class="flex flex-col gap-2 shrink-0">
      <div v-click class="chip whitespace-nowrap">📄 Dockerfile</div>
      <div v-click class="chip !inline-flex flex-col items-center text-center" style="padding: .4rem .6rem; border-radius: 12px"><img src="/img/stamp-only.png" style="height: 6rem; width: auto" /><span class="mt-1">image</span></div>
    </div>
    <div v-click class="arrow">→</div>
    <div class="grid grid-cols-4 gap-x-4 gap-y-3">
      <Ctr v-for="n in 12" :key="n" v-click stamped />
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
    <div class="chip" style="display: inline-flex; align-items: center; gap: .45rem; padding: .25rem .8rem .25rem .45rem; border-radius: 12px"><img src="/img/stamp-only.png" style="height: 1.9rem; width: auto" />image</div><Ctr />
  </Laptop>
</div>

<div v-click class="mt-4 text-center text-xl">Six laptops · one image · <span class="gradient-text font-700">identical behaviour</span></div>

---
layout: center
---

# Docker life cycle

<div class="text-xl opacity-80">Three steps. <b>Write it, save it, start it.</b></div>

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
    <div class="h-24 flex items-center justify-center"><img src="/img/stamp-only.png" style="height: 6rem; width: auto" /></div>
    <div class="text-sm">Everything is saved as a <b>stamp</b>.</div>
    <div class="mt-auto pt-3"><span class="chip">Save</span></div>
  </div>

  <div v-click class="text-center self-center" style="width: 8.5rem">
    <div class="small">then</div>
    <div class="chip mt-1" style="font-size: 1rem; border-color: var(--m)">docker run</div>
    <div class="arrow">→</div>
  </div>

  <div v-click class="card glow text-center min-w-0 flex flex-col">
    <div class="text-sm font-700" style="color: var(--m)">STEP 3</div>
    <div class="text-3xl font-800 mt-1">Container</div>
    <div class="h-24 flex items-center justify-center"><img src="/img/photo-containers.jpg" style="border-radius:8px; height: 5.5rem; width: 100%; object-fit: cover" /></div>
    <div class="text-sm">Press it, and a container with your app <b>runs</b>. Press as many as you like.</div>
    <div class="mt-auto pt-3"><span class="chip">Start</span></div>
  </div>

</div>

<div v-click class="mt-8 text-center text-xl">
  <b>Dockerfile</b> = the instructions &nbsp;·&nbsp; <b>Image</b> = the stamp &nbsp;·&nbsp; <b>Container</b> = made from it, running
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

<div v-click class="mt-3 text-center opacity-70">each instruction → one layer of the <b>image</b></div>

<div v-click class="mt-4 pt-3" style="border-top: 1px solid var(--card-line)">
  <div class="text-center font-700 mb-2">Choose the right tool: pick the base image that fits your job</div>
  <div class="grid grid-cols-3 gap-4">
    <div class="card flex items-center gap-3" style="padding: .5rem .9rem"><img src="/img/logo-node.svg" style="height: 2.4rem; width: 2.4rem; object-fit: contain" /><div><div class="font-700">FROM node</div><div class="small">JavaScript</div></div></div>
    <div class="card flex items-center gap-3" style="padding: .5rem .9rem"><img src="/img/logo-python.svg" style="height: 2.4rem; width: 2.4rem; object-fit: contain" /><div><div class="font-700">FROM python</div><div class="small">Data / scripts</div></div></div>
    <div class="card flex items-center gap-3" style="padding: .5rem .9rem"><img src="/img/logo-go.svg" style="height: 2.4rem; width: 2.4rem; object-fit: contain" /><div><div class="font-700">FROM golang</div><div class="small">Services</div></div></div>
  </div>
  <div class="mt-2 text-center small">then add what you need on top <span class="opacity-60">· Go gopher by Renée French (CC BY 3.0)</span></div>
</div>

---

# Image <span class="opacity-50">?</span>

<div class="grid grid-cols-2 gap-12 mt-4 items-center">

<div>

<div class="text-2xl font-700 leading-snug">A <span class="gradient-text">stamp</span> of your app,<br>built from <b>layers</b>.</div>

<div class="mt-3 flex items-center gap-3"><img src="/img/stamp-only.png" style="height: 10rem; width: auto" /><div class="small" style="max-width: 9rem">one stamp,<br>press it as often<br>as you like</div></div>

<div class="mt-5 flex flex-col gap-3">
  <div v-click="5" class="flex items-center gap-3"><span class="text-3xl">🔒</span><div><b>Read-only</b><div class="small">a finished image never changes</div></div></div>
  <div v-click="6" class="flex items-center gap-3"><span class="text-3xl">📤</span><div><b>Shareable</b><div class="small">push it once, pull it anywhere</div></div></div>
  <div v-click="7" class="flex items-center gap-3"><span class="text-3xl">♻️</span><div><b>Reusable</b><div class="small">one image → as many containers as you like</div></div></div>
</div>

</div>

<div class="relative">
  <div v-click="5" class="image-tag"><span><img src="/img/stamp-only.png" style="height:1.7em;width:auto;display:inline-block;vertical-align:-0.45em" /></span> tree-api:1.0</div>
  <div class="image-stack" :class="{ sealed: $clicks >= 5 }">
    <div v-click="1" class="layer l1"><span class="who">base OS + Go</span><code>FROM golang:1.25-alpine</code></div>
    <div v-click="2" class="layer l2"><span class="who">dependencies</span><code>RUN go mod download</code></div>
    <div v-click="3" class="layer l3"><span class="who">your code</span><code>COPY . .</code></div>
    <div v-click="4" class="layer l4"><span class="who">build</span><code>RUN go build -o api .</code></div>
  </div>
</div>

</div>


---

# Container <span class="opacity-50">?</span>

<div class="flex items-center justify-center gap-8 mt-2">
  <div class="card glow text-center"><div class="text-5xl"><img src="/img/stamp-only.png" style="height:1.7em;width:auto;display:inline-block;vertical-align:-0.45em" /></div><div class="font-700">image</div></div>
  <div class="arrow">→</div>
  <ContainerBig name="tree-api" :open="$clicks >= 1" :width="400" />
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

<div class="text-lg opacity-80">Same <b>stamp</b> (image), a postcard and a ticket (two containers). <b>Writing on one does not touch the other.</b></div>

<div class="grid grid-cols-[1.1fr_1fr] gap-6 mt-3">

<div class="min-w-0 flex flex-col gap-2">

  <div v-click="1">
    <div class="small mb-1"><b>1</b> · Press the same stamp on a postcard and a ticket</div>

```bash
docker run -d --name postcard \
  alpine sleep 3600
docker run -d --name ticket \
  alpine sleep 3600
```

  </div>

  <div v-click="2">
    <div class="small mb-1"><b>2</b> · Write a note in <code>postcard</code> only</div>

```bash
docker exec postcard sh -c \
  "echo hi > /note.txt"
```

  </div>

  <div v-click="3">
    <div class="small mb-1"><b>3</b> · Look for the note in each one</div>

```bash
docker exec postcard cat /note.txt
docker exec ticket cat /note.txt
```

  </div>

</div>

<div class="min-w-0 flex flex-col items-center gap-2">

  <div class="card glow text-center" style="padding: .5rem 1.4rem"><span class="text-3xl"><img src="/img/stamp-only.png" style="height:1.7em;width:auto;display:inline-block;vertical-align:-0.45em" /></span> <b>stamp = image: alpine</b></div>
  <div class="arrow" style="transform: rotate(90deg); animation: none">→</div>

  <div class="grid grid-cols-2 gap-3 w-full">
    <div v-click="1" class="card text-center">
      <div class="font-700">container <code>postcard</code></div>
      <div v-click="2" class="mt-2 text-3xl">📝</div>
      <div v-click="2" class="small"><code>note.txt</code></div>
      <div v-click="3" class="label-good mt-1">✅ finds it</div>
    </div>
    <div v-click="1" class="card text-center">
      <div class="font-700">container <code>ticket</code></div>
      <div v-click="2" class="mt-2 text-3xl opacity-30">⬜</div>
      <div v-click="2" class="small">still blank</div>
      <div v-click="3" class="label-bad mt-1">❌ not found</div>
    </div>
  </div>

</div>

</div>

<div v-click="3" class="mt-3 text-center text-lg">Each container has its <b>own private space</b>, like a postcard and a ticket are separate papers. Nothing leaks between them.</div>

---

# Containers are stateless

<div class="text-lg opacity-80">Throw the postcard away and press a new one. <b>It starts blank, every time.</b></div>

<div class="grid grid-cols-[1.1fr_1fr] gap-6 mt-3">

<div class="min-w-0 flex flex-col gap-2">

  <div v-click="1">
    <div class="small mb-1"><b>4</b> · Throw <code>postcard</code> away</div>

```bash
docker rm -f postcard
```

  </div>

  <div v-click="2">
    <div class="small mb-1"><b>5</b> · Press the same stamp again: a new <code>postcard</code></div>

```bash
docker run -d --name postcard \
  alpine sleep 3600
```

  </div>

  <div v-click="3">
    <div class="small mb-1"><b>6</b> · Look for the note again</div>

```bash
docker exec postcard cat /note.txt
```

  </div>

</div>

<div class="min-w-0 flex flex-col items-center gap-2">

  <div class="card glow text-center" style="padding: .5rem 1.4rem"><span class="text-3xl"><img src="/img/stamp-only.png" style="height:1.7em;width:auto;display:inline-block;vertical-align:-0.45em" /></span> <b>stamp = image: alpine</b> <span class="small">(never changed)</span></div>

  <div class="grid grid-cols-[1fr_auto_1fr] items-center gap-2 w-full mt-2">
    <div class="card text-center">
      <div class="font-700">old <code>postcard</code></div>
      <div class="text-3xl mt-1">📝</div>
      <div class="small"><code>note.txt</code></div>
    </div>
    <div v-click="1" class="text-3xl">🗑️</div>
    <div v-click="2" class="card text-center">
      <div class="font-700">new <code>postcard</code></div>
      <div class="text-3xl mt-1 opacity-30">⬜</div>
      <div v-click="3" class="label-bad mt-1">❌ file gone</div>
    </div>
  </div>

</div>

</div>

<div v-click="3" class="mt-3 text-center text-lg">The stamp is the <b>same every time</b>, so every new postcard starts the <b>same</b>: blank. Nothing you wrote is saved.</div>

<div v-click="3" class="mt-1 text-center small">Need to keep data? That is what a <b>volume</b> is for (Step 6). Clean up: <code>docker rm -f postcard ticket</code></div>

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
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-1</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

<div class="small mb-1"><code>web/Dockerfile</code>: just the page, nothing else yet</div>

```dockerfile
FROM nginx:1.31.9-alpine3.24
COPY index.html og.svg /usr/share/nginx/html/
```

```bash {1|2}
docker build -t tree-web ./web
docker run -d --name web -p 3000:80 tree-web
```

<div class="mt-3 text-sm opacity-80 leading-relaxed">
<code>-d</code> run in the background · <code>--name</code> give it a name<br>
<code>-p 3000:80</code> <b>host</b> port 3000 → <b>container</b> port 80
</div>

<div class="mt-4"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  Open <b>http://localhost:3000</b>
  <div class="mt-2">A cherry-blossom tree appears. Type a link and the QR code and the tree work, with 3 flower colours. Now press <b>Save</b> or open the <b>Gallery</b>:</div>
  <div class="out mt-1">API not reachable</div>
  <div class="mt-3 text-sm opacity-70">There is no API in this image: only <code>index.html</code> and nginx. First, let's change the page.</div>
</div>

</div>

---

# Step 1 · Change the page, then rebuild

<Stepper :n="1" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-1b</code></div>

<div class="grid grid-cols-[1.25fr_1fr] gap-8 mt-3">

<div>

<div v-click="1">
<div class="small mb-1"><b>A</b> · Add a 4th flower colour in <code>web/index.html</code> (line 263). This is the <code>Step 1b</code> commit</div>

```diff
- let SWATCHES = ['#e98aab','#a57ad6','#e0473c'];
+ let SWATCHES = ['#e98aab','#a57ad6','#e0473c','#f5b301'];
```

</div>

<div v-click="2" class="mt-3">
<div class="small mb-1"><b>B</b> · Run it again</div>

```bash
docker restart web
```

</div>

<div v-click="3" class="mt-3">
<div class="small mb-1"><b>C</b> · Build a <b>new image</b> and run that one</div>

```bash
docker build -t tree-web:v2 ./web
docker rm -f web
docker run -d --name web -p 3000:80 tree-web:v2
docker images tree-web
```

</div>

</div>

<div>

<div v-click="2" class="see">
  <div class="h">B · YOU SHOULD SEE</div>
  Refresh <b>http://localhost:3000</b>: still <b>3</b> colours.
  <div class="mt-2 text-sm opacity-80">The image is a <b>snapshot</b>: <code>index.html</code> was copied in when we built it. Editing the file on your laptop, or restarting the container, changes nothing.</div>
</div>

<div v-click="3" class="see mt-3">
  <div class="h">C · YOU SHOULD SEE</div>
  Refresh again: <b>4</b> colours.
  <div class="out mt-1">REPOSITORY&nbsp;&nbsp;TAG<br>tree-web&nbsp;&nbsp;&nbsp;&nbsp;v2<br>tree-web&nbsp;&nbsp;&nbsp;&nbsp;latest</div>
  <div class="mt-2 text-sm opacity-80">New code means a <b>new image</b>, and a new container from it. The old image is untouched.</div>
</div>

</div>

</div>

<div v-click="4" class="mt-4 text-center text-lg">Save still says <b>API not reachable</b>. We need a second container: the API.</div>

---

# Step 2 · A second container: the API

<Stepper :n="2" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-2</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

```bash {1|2|3}
docker build -t tree-api ./api
docker run -d --name api -p 8080:8080 tree-api
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

# Step 3 · Point the web at the API

<Stepper :n="3" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-3</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

<div class="small mb-1">The API answers on <code>localhost:8080</code>, so give <code>web/Dockerfile</code> an nginx config that forwards <code>/api</code> there:</div>

```dockerfile {2-3}
FROM nginx:1.31.9-alpine3.24
ENV API_URL=http://localhost:8080
COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY index.html og.svg /usr/share/nginx/html/
```

```bash {1|2|3|4}
docker rm -f web
docker build -t tree-web ./web
docker run -d --name web -p 3000:80 tree-web
curl -i localhost:3000/api/health
```

<div class="mt-3"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out" style="color: #dc2626">HTTP/1.1 502 Bad Gateway</div>
  <div class="mt-3">The API is running and port 8080 is open on your laptop, yet the web cannot reach it through <code>localhost</code>.</div>
  <div class="mt-3 text-sm opacity-70"><code>docker logs web</code> says <code>connect() failed (111: Connection refused)</code>. Why?</div>
</div>

</div>

---

# Why? Every container has its own <code>localhost</code>

<div class="text-lg opacity-80"><code>localhost</code> means <b>"this machine"</b>. Inside a container, "this machine" is the container itself.</div>

<div class="grid grid-cols-3 gap-4 mt-6">

  <div v-click="1" class="card text-center">
    <div class="text-4xl">💻</div>
    <div class="font-800 mt-1">Your laptop</div>
    <div class="small mt-1"><code>localhost:8080</code></div>
    <div class="mt-2 label-good">✅ API answers</div>
    <div class="small">because of <code>-p 8080:8080</code></div>
  </div>

  <div v-click="2" class="card bad text-center">
    <div class="text-4xl">🌐</div>
    <div class="font-800 mt-1">web container</div>
    <div class="small mt-1"><code>localhost:8080</code></div>
    <div class="mt-2 label-bad">❌ nothing here</div>
    <div class="small">this is <b>web</b> itself, not the API</div>
  </div>

  <div v-click="3" class="card text-center">
    <div class="text-4xl">⚙️</div>
    <div class="font-800 mt-1">api container</div>
    <div class="small mt-1">listens on <code>:8080</code></div>
    <div class="mt-2 label-good">✅ it is here</div>
    <div class="small">but <code>localhost</code> never points here</div>
  </div>

</div>

<div v-click="4" class="mt-8 text-center big-line"><code>localhost</code> is the wrong address. <span class="gradient-text">So what is the right one?</span></div>

---

# Could web use the published port instead?

<div class="text-lg opacity-80">Yes, but not with <code>localhost</code>. Docker gives containers a name for <b>your laptop</b>: <code>host.docker.internal</code>.</div>

<div class="grid grid-cols-[1fr_1fr] gap-6 mt-3">

<div class="compact">

```bash {1-2|3-4}
# web → api   (api published: -p 8080:8080)
API_URL=http://host.docker.internal:8080
# api → db    (db published: -p 5432:5432)
DATABASE_URL=postgres://…@host.docker.internal:5432/…
```

<div v-click="1" class="grid grid-cols-2 gap-3 mt-4">
  <div class="card bad text-center"><div class="label-bad">❌ localhost</div><div class="small">the container itself</div></div>
  <div class="card good text-center"><div class="label-good">✅ host.docker.internal</div><div class="small">your laptop</div></div>
</div>

<div v-click="1" class="small mt-2">We tried both: only the second one answers.</div>

</div>

<div>

<div v-click="2" class="card" style="padding: .6rem .9rem">
  <div class="font-800">So why not do this?</div>
  <div class="small mt-1" style="color: inherit">🚪 <b>Every service publishes a port</b>, even the database: anyone who reaches your laptop reaches it.</div>
  <div class="small mt-1" style="color: inherit">🔁 The call leaves the container, goes to the host, and comes back.</div>
  <div class="small mt-1" style="color: inherit">💻 <b>Laptop-only.</b> Linux needs an extra <code>--add-host</code> flag, a server has no such host.</div>
  <div class="small mt-1" style="color: inherit">1️⃣ Only <b>one</b> container can publish <code>8080</code>: no <code>--scale api=3</code>.</div>
</div>

</div>

</div>

<div v-click="3" class="mt-3 text-center big-line">It works on your laptop. A <span class="gradient-text">network + alias</span> works everywhere.</div>

---

# Step 4 · Write it down: Docker Compose

<Stepper :n="4" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-4</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-6 mt-2">

<div class="compact">

```yaml {1-6|7-11|12-13}
services:
  web:
    build: ./web
    ports: ["3000:80"]
    networks: [tree]
    depends_on: [api]
  api:
    build: ./api
    networks:
      tree:
        aliases: [api]
networks:
  tree:
```

</div>

<div>

<div v-click="1" class="card">
  <b>web</b>: same image, same port <code>3000:80</code>, and it joins the private network <code>tree</code>.
</div>

<div v-click="2" class="card glow mt-3">
  <b>api</b> joins the same network and answers to the <b>alias</b> <code>api</code>. Any container on <code>tree</code> can call <code>http://api:8080</code>.
  <div class="small mt-1">No <code>ports:</code>: only <b>web</b> is open to your laptop.</div>
</div>

<div v-click="3" class="card mt-3">
  The network is declared once at the bottom.
</div>

</div>

</div>

---

# The same thing, written down

<div class="text-lg opacity-80">Every flag you typed has a line in the <b>compose file</b>. <span class="opacity-70">(Step 4)</span></div>

<div class="grid grid-cols-[1fr_auto_1fr] items-center gap-x-3 gap-y-2 mt-3">

  <div class="text-center font-700 opacity-70">On the command line</div><div></div><div class="text-center font-700 opacity-70">In <code>docker-compose.yml</code></div>

  <div v-click="1"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>docker build -t tree-web ./web</code><span class="small">make the image</span></div></div>
  <div v-click="1" class="arrow">→</div>
  <div v-click="1"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>build: ./web</code><span class="small">same folder, same Dockerfile</span></div></div>

  <div v-click="2"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>--name web</code><span class="small">name the container</span></div></div>
  <div v-click="2" class="arrow">→</div>
  <div v-click="2"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>web:</code><span class="small">the service name</span></div></div>

  <div v-click="3"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>-p 3000:80</code><span class="small">open a door to your browser</span></div></div>
  <div v-click="3" class="arrow">→</div>
  <div v-click="3"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>ports: ["3000:80"]</code><span class="small">same two numbers</span></div></div>

  <div v-click="4"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>docker run -d …</code><span class="small">once per container</span></div></div>
  <div v-click="4" class="arrow">→</div>
  <div v-click="4"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>docker compose up</code><span class="small">starts all of them together</span></div></div>

</div>

---

# Now the API, written down

<div class="text-lg opacity-80">The API gets its own service, <b>next to web</b>, on the <b>same network</b>. <span class="opacity-70">(Step 2 and Step 4)</span></div>

<div class="grid grid-cols-[1fr_auto_1fr] items-center gap-x-3 gap-y-2 mt-3">

  <div class="text-center font-700 opacity-70">On the command line</div><div></div><div class="text-center font-700 opacity-70">In <code>docker-compose.yml</code></div>

  <div v-click="1"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>docker build -t tree-api ./api</code><span class="small">make the image</span></div></div>
  <div v-click="1" class="arrow">→</div>
  <div v-click="1"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>build: ./api</code><span class="small">same folder, same Dockerfile</span></div></div>

  <div v-click="2"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>--name api</code><span class="small">name the container</span></div></div>
  <div v-click="2" class="arrow">→</div>
  <div v-click="2"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>api:</code><span class="small">the service name</span></div></div>

  <div v-click="3"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>-p 8080:8080</code><span class="small">open a door to your browser</span></div></div>
  <div v-click="3" class="arrow">→</div>
  <div v-click="3"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><b>nothing to write</b><span class="small">no <code>ports:</code>, the API stays private</span></div></div>

  <div v-click="4"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>http://localhost:8080</code><span class="small">only works on your laptop</span></div></div>
  <div v-click="4" class="arrow">→</div>
  <div v-click="4"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>aliases: [api]</code><span class="small">web calls <code>http://api:8080</code></span></div></div>

  <div v-click="5"><div class="card min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>start api before web</code><span class="small">order matters</span></div></div>
  <div v-click="5" class="arrow">→</div>
  <div v-click="5"><div class="card glow min-w-0 flex items-baseline gap-3" style="padding: .4rem .9rem"><code>depends_on: [api]</code><span class="small">written under <b>web</b></span></div></div>

</div>

---

# Step 4 · Fix the one wrong line

<Stepper :n="4" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

<div class="small mb-1"><code>web/Dockerfile</code></div>

```diff
- ENV API_URL=http://localhost:8080
+ ENV API_URL=http://api:8080
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
<code>api</code> is the <b>alias</b> from the compose file. <code>8080</code> is the port <b>inside</b> the API container, so nothing has to be published.
</div>

<div class="mt-3 text-sm opacity-80 leading-relaxed">
nginx also asks Docker's DNS for <code>api</code> <b>on every request</b>, so the API can be recreated without restarting the web.
</div>

</div>

<div v-click class="card glow">
  <div class="font-800">Review it</div>
  <div class="mt-2 text-sm">The web is the <b>only</b> file that changes. The API image is untouched: it does not know who calls it.</div>
</div>

</div>

---

# Run it, for real

<Stepper :n="4" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.2fr_1fr] gap-6 mt-2">

<div>

```bash {1|2|3|4}
docker rm -f web api
docker compose up -d --build
docker compose ps
curl localhost:3000/api/health
```

<div v-click="2" class="out mt-3" style="font-size: .8rem; line-height: 1.35">
Network demo_tree&nbsp;&nbsp;Created<br>
Container demo-api-1&nbsp;&nbsp;Started<br>
Container demo-web-1&nbsp;&nbsp;Started
</div>

<div v-click="3" class="out mt-2" style="font-size: .8rem; line-height: 1.35">
NAME&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SERVICE&nbsp;&nbsp;STATUS<br>
demo-api-1&nbsp;&nbsp;api&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Up 7 seconds<br>
demo-web-1&nbsp;&nbsp;web&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Up 7 seconds
</div>

<div v-click="4" class="out mt-2" style="font-size: .8rem">{"db":"not connected","host":"ec7dffe4c6bd","status":"ok"}</div>

<div v-click="4" class="mt-3 text-sm opacity-80">The call goes laptop → <b>web</b> → <code>api</code>. <code>curl localhost:8080</code> no longer works: the API is private now.</div>

</div>

<div v-click="4" class="text-center">
  <img src="/img/compose-result.jpg" style="height: 21rem; width: auto; border-radius: 10px; box-shadow: 0 6px 20px rgba(0,0,0,.18)" />
  <div class="small mt-1">the garden, started by Compose</div>
</div>

</div>

---

# Step 5 · Add the database

<Stepper :n="5" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-5</code></div>

<div class="grid grid-cols-[1.35fr_1fr] gap-6 mt-2">

<div class="compact">

```yaml {1-7|8-17}
  api:
    build: ./api
    environment:
      DATABASE_URL: postgres://magic:magictree@db:5432/treedb?sslmode=disable
    depends_on: { db: { condition: service_healthy } }
    networks: ...
  db:
    image: postgres:18-alpine3.24
    environment:
      POSTGRES_USER: magic
      POSTGRES_PASSWORD: magictree
      POSTGRES_DB: treedb
    volumes:
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro
    healthcheck: ...
    networks: [tree]
```

<div class="mt-3"><span class="try">⏸ Pause &amp; try</span> <code class="ml-2">docker compose up -d</code></div>

</div>

<div>

<div class="card">
  <div class="text-sm"><code>DATABASE_URL</code> tells the API where the database is. <code>db</code> is the <b>service name</b>: same idea as the alias.</div>
  <div class="text-sm mt-2"><code>init.sql</code> creates the tables and ten example gardens.</div>
</div>

<div v-click class="see mt-3">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out">{"db":"connected",<br>"shared_trees":10, ...}</div>
  <div class="mt-2 text-sm">The gallery has 10 gardens and <b>6 flower colours</b>. Press 🔖 to save yours: <code>shared_trees</code> becomes <b>11</b>.</div>
</div>

</div>

</div>

---

# Step 5 · …and then we lose it

<Stepper :n="5" class="absolute top-6 right-10" />

<div class="grid grid-cols-[1.1fr_1fr] gap-8 mt-6">

<div>

```bash {1|2|3}
docker compose down
docker compose up -d
curl localhost:3000/api/health
```

<div class="mt-4"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div v-click class="card bad">
  <div class="label-bad">❌ Without a volume</div>
  <div class="out mt-2">{"db":"connected",<br>"shared_trees":<b>10</b>, ...}</div>
  <div class="text-sm mt-2">Your garden is <b>gone</b>. <code>down</code> removed the db container, and the data lived <b>inside</b> it.</div>
</div>

</div>

<div v-click class="mt-8 text-center big-line">The data needs a home that is <span class="gradient-text">not the container</span>.</div>

---

# Why did the data disappear?

<Stepper :n="5" class="absolute top-6 right-10" />

<div class="text-lg opacity-80">A container is the image <b>plus a thin writable layer</b> on top. Everything the database saves goes into that layer.</div>

<div class="grid grid-cols-[1fr_auto_1fr] items-center gap-4 mt-5">

  <div class="card min-w-0">
    <div class="font-800 text-center">🐘 db container</div>
    <div class="card glow mt-2 text-center" style="padding: .5rem">✍️ <b>writable layer</b><div class="small">your saved trees live here</div></div>
    <div class="card mt-2 text-center" style="padding: .5rem">📦 <b>image layers</b><div class="small">postgres:18 · read-only</div></div>
  </div>

  <div v-click="1" class="text-center">
    <div class="text-3xl">🗑️</div>
    <div class="small"><code>docker compose down</code></div>
  </div>

  <div v-click="1" class="card min-w-0">
    <div class="font-800 text-center">🐘 new db container</div>
    <div class="card bad mt-2 text-center" style="padding: .5rem">🫙 <b>new, empty layer</b><div class="small label-bad">trees gone</div></div>
    <div class="card mt-2 text-center" style="padding: .5rem">📦 <b>same image layers</b><div class="small">postgres:18 · read-only</div></div>
  </div>

</div>

<div v-click="2" class="mt-6 text-center big-line">The writable layer <span class="label-bad">dies with the container</span>. Keep the data <span class="gradient-text">somewhere else</span>.</div>

---

# A volume lives outside the container

<Stepper :n="6" class="absolute top-6 right-10" />

<div class="text-lg opacity-80">A <b>volume</b> is a folder that Docker keeps on the host. A container <b>mounts</b> it at a path, like plugging in a USB drive.</div>

<div class="grid grid-cols-[1fr_auto_1fr] items-center gap-4 mt-5">

  <div class="card min-w-0 text-center">
    <div class="font-800">🐘 db container</div>
    <div class="small mt-1">Postgres writes to</div>
    <div class="chip mt-1"><code>/var/lib/postgresql</code></div>
  </div>

  <div class="text-center">
    <div class="arrow" style="font-size: 2rem">⇄</div>
    <div class="small">mounted</div>
  </div>

  <div class="card glow min-w-0 text-center">
    <div class="font-800">💾 volume <code>tree-data</code></div>
    <div class="small mt-1">managed by Docker</div>
    <div class="small">outside any container</div>
  </div>

</div>

<div class="grid grid-cols-3 gap-4 mt-6">

  <div v-click="1" class="card">
    <div class="font-800">1 · Save</div>
    <div class="text-sm mt-1">Postgres writes a tree. It lands <b>in the volume</b>, not in the container.</div>
  </div>

  <div v-click="2" class="card">
    <div class="font-800">2 · Remove</div>
    <div class="text-sm mt-1"><code>docker compose down</code> deletes the container. <b>The volume stays.</b> 🗑️ 💾</div>
  </div>

  <div v-click="3" class="card good">
    <div class="font-800">3 · Start again</div>
    <div class="text-sm mt-1">A new container mounts the <b>same volume</b>. The trees are back. ✅</div>
  </div>

</div>

<div v-click="3" class="mt-5 text-center big-line">Container = <span class="label-bad">disposable</span>. Volume = <span class="gradient-text">the data</span>.</div>

---

# Step 6 · Keep the data: a volume

<Stepper :n="6" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-6</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-6 mt-2">

<div class="compact">

```yaml {5|8-9}
  db:
    image: postgres:18-alpine3.24
    ...
    volumes:
      - tree-data:/var/lib/postgresql
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro
    ...
volumes:
  tree-data:
```

```bash
docker compose down -v      # start clean
docker compose up -d
# save a garden → 11
docker compose down         # no -v!
docker compose up -d
curl localhost:3000/api/health
```

<div class="mt-2"><span class="try">⏸ Pause &amp; try</span></div>

</div>

<div>

<div class="card">
  <div class="text-sm"><code>tree-data</code> is a <b>named volume</b>: storage managed by Docker, outside any container. Postgres 18 keeps its files in <code>/var/lib/postgresql</code>.</div>
</div>

<div v-click class="card good mt-3">
  <div class="label-good">✅ With a volume</div>
  <div class="out mt-2">{"db":"connected",<br>"shared_trees":<b>11</b>, ...}</div>
  <div class="text-sm mt-2">Delete and recreate the container: your garden is <b>still there</b>.</div>
</div>

<div v-click class="small mt-3">⚠️ <code>docker compose down <b>-v</b></code> also deletes the volume, and the data.</div>

</div>

</div>

---

# Look inside: volumes and mounts

<Stepper :n="6" class="absolute top-6 right-10" />

<div class="grid grid-cols-2 gap-6 mt-3">

<div>

```bash {1|2}
docker volume ls
docker volume inspect demo_tree-data
```

<div class="small mt-2">The name starts with the project folder (<code>demo_</code>). <code>Mountpoint</code> is where Docker keeps the files.</div>

<div v-click class="card bad mt-4">
  <div class="label-bad">⚠️ Gotcha</div>
  <div class="text-sm mt-1"><code>init.sql</code> only runs when the volume is <b>empty</b>, on the very first start. To start from scratch:</div>

```bash
docker compose down -v
```

</div>

</div>

<div>

<div class="card glow">
  <div class="font-800">💾 Named volume</div>
  <div class="mt-1"><code>tree-data:/var/lib/postgresql</code></div>
  <div class="small mt-1">Docker owns the files. For <b>data you keep</b>: databases, uploads.</div>
</div>

<div class="card mt-3">
  <div class="font-800">📁 Bind mount</div>
  <div class="mt-1" style="word-break: break-all"><code>./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro</code></div>
  <div class="small mt-1">A file or folder from <b>your laptop</b>. For <b>config and seed files</b>. <code>:ro</code> = read-only.</div>
</div>

<div class="small mt-3">Both are in our <code>db</code> service. The only difference is the left side of the colon: a <b>name</b> or a <b>path</b>.</div>

</div>

</div>

---

# Step 7 · Scale the API

<Stepper :n="7" class="absolute top-6 right-10" />
<div class="absolute top-[4.4rem] right-10 small opacity-70">get this step: <code>git checkout step-7</code></div>

<div class="grid grid-cols-[1.2fr_1fr] gap-8 mt-4">

<div>

```bash {1|2|3-4}
docker compose up -d --scale api=3
sleep 6
for i in 1 2 3 4 5 6; do
  curl -s localhost:3000/api/health; echo; done
```

<div class="mt-4 text-sm opacity-80 leading-relaxed">
Three copies of the <b>same image</b>, all answering to the alias <code>api</code>. Docker's DNS hands them out in turn.<br>
<span class="opacity-70">(<code>sleep 6</code>: nginx re-reads the DNS every 5 seconds.)</span>
</div>

<div class="mt-5"><span class="try">⏸ Pause &amp; try</span> <span class="small ml-2">finish with</span> <code>docker compose down</code></div>

</div>

<div v-click class="see">
  <div class="h">YOU SHOULD SEE</div>
  <div class="out">"host":"4d633404bd51"<br>"host":"a7a19b993e2b"<br>"host":"dc88b4d8a103"<br>...</div>
  <div class="flex gap-2 mt-3"><Ctr v-for="n in 3" :key="n" /></div>
  <div class="mt-2 text-sm opacity-70">Three different container IDs: one image, three containers. And the data is still there.</div>
</div>

</div>

---
layout: center
---

# What you built

<div class="grid grid-cols-2 gap-4 mt-6 text-left">
  <div v-click class="card"><b>1 · docker run</b><div class="small">starts a container from an image</div></div>
  <div v-click class="card"><b>2 · -p host:container</b><div class="small">opens a door to the outside</div></div>
  <div v-click class="card"><b>3 · localhost</b><div class="small">every container has its own</div></div>
  <div v-click class="card"><b>4 · compose + alias</b><div class="small">the stack in one file, found by name</div></div>
  <div v-click class="card"><b>5 · DATABASE_URL</b><div class="small">config without changing the image</div></div>
  <div v-click class="card"><b>6 · volume</b><div class="small">data outlives the container</div></div>
  <div v-click class="card col-span-2"><b>7 · --scale</b><div class="small">one image, many containers</div></div>
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
