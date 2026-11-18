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

# Where are you now?

Get the workshop code first:

```bash
git clone https://github.com/AnuchitO/docker-qadev.git
```

<div class="grid grid-cols-2 gap-6 mt-8">

<div class="card" v-click>
  <div class="text-xl font-700 mb-2">🐳 Docker</div>
  <div class="flex gap-3"><span class="chip">✅ I have Docker</span><span class="chip">❌ I don't have Docker</span></div>
</div>

<div class="card" v-click>
  <div class="text-xl font-700 mb-2">🐹 Golang</div>
  <div class="flex gap-3"><span class="chip">✅ I know Go</span><span class="chip">❌ I don't know Go</span></div>
</div>

</div>

<div v-click class="mt-8 text-center opacity-70">Either way — you will be fine. That's the whole point of Docker.</div>

---
layout: center
class: text-center
---

# Run the app <span class="gradient-text">inside</span> Docker

<div class="flex items-center justify-center gap-6 mt-10">
  <div v-click class="card text-4xl">📦 app</div>
  <div v-click class="arrow">→</div>
  <div v-click class="card glow pulse text-4xl">🐳 docker</div>
</div>

---

# Why do we need Docker?

<div class="grid grid-cols-2 gap-6 mt-8">

<div v-click class="card">
  <div class="text-3xl">🤷</div>
  <div class="font-700 text-lg mt-1">"It works on my machine"</div>
  <div class="opacity-60 text-sm">Different OS, versions and settings between Dev, QA and Prod.</div>
</div>

<div v-click class="card">
  <div class="text-3xl">📦</div>
  <div class="font-700 text-lg mt-1">Same environment everywhere</div>
  <div class="opacity-60 text-sm">Package the app and everything it needs, run it anywhere.</div>
</div>

<div v-click class="card">
  <div class="text-3xl">🧹</div>
  <div class="font-700 text-lg mt-1">Clean &amp; disposable</div>
  <div class="opacity-60 text-sm">Break it? Remove it and start again in seconds.</div>
</div>

<div v-click class="card">
  <div class="text-3xl">🚀</div>
  <div class="font-700 text-lg mt-1">Fast onboarding</div>
  <div class="opacity-60 text-sm">New teammate or tester runs the whole stack with one command.</div>
</div>

</div>

---
layout: center
---

# Docker life cycle

<div class="flex items-center justify-center gap-6 mt-12">
  <div v-click class="card text-center w-56">
    <div class="h-20 flex items-center justify-center"><DockerfileIcon /></div>
    <div class="text-xl font-700 mt-2">Dockerfile</div>
    <div class="opacity-60 text-sm">the recipe</div>
  </div>
  <div v-click class="arrow">→</div>
  <div v-click class="card text-center w-56">
    <div class="h-20 flex items-center justify-center text-5xl">💿</div>
    <div class="text-xl font-700 mt-2">Image</div>
    <div class="opacity-60 text-sm">the template</div>
  </div>
  <div v-click class="arrow">→</div>
  <div v-click class="card glow text-center w-56">
    <div class="h-20 flex items-center justify-center"><ContainerBig :open="true" :width="120" /></div>
    <div class="text-xl font-700 mt-2">Container</div>
    <div class="opacity-60 text-sm">the running thing</div>
  </div>
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
  <div v-click="5" class="image-tag"><span>💿</span> todo-app:1.0</div>
  <div class="image-stack" :class="{ sealed: $clicks >= 5 }">
    <div v-click="1" class="layer l1"><span class="who">base OS + Go</span><code>FROM golang:1.23-alpine</code></div>
    <div v-click="2" class="layer l2"><span class="who">dependencies</span><code>RUN go mod download</code></div>
    <div v-click="3" class="layer l3"><span class="who">your code</span><code>COPY . .</code></div>
    <div v-click="4" class="layer l4"><span class="who">build</span><code>RUN go build -o server .</code></div>
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
FROM golang:1.23-alpine
RUN <command>
```

```dockerfile
FROM golang:1.23-alpine
RUN go mod download
```

```dockerfile
FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o server .
CMD ["./server"]
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
  <ContainerBig name="todo-app" :open="$clicks >= 1" :width="400" />
</div>

<div v-click="1" class="mt-4 text-center"><span class="chip">docker run</span> — the doors open and the app starts</div>

<div v-click="2" class="flex items-center justify-center gap-4 mt-4">
  <span class="opacity-70">same image, more containers →</span>
  <Ctr v-for="n in 4" :key="n" />
</div>

<div v-click="3" class="mt-6 text-center text-xl">An image is the <b>class</b>, a container is an <b>instance</b>.</div>

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

# Build &amp; run

```bash {all|1|2}
docker build -t anuchito/docker-qadev .
docker run anuchito/docker-qadev
```

<div class="flex items-center justify-center gap-4 mt-10">
  <div class="card text-center">📄 Dockerfile</div>
  <div v-click="1" class="arrow">build →</div>
  <div v-click="1" class="card text-center">💿 image</div>
  <div v-click="2" class="arrow">run →</div>
  <div v-click="2" class="card glow text-center">📦 container</div>
</div>

---
layout: center
---

# Registry

<div class="flex items-center justify-center gap-6 mt-10">
  <div v-click class="card text-center"><div class="text-4xl">💻</div><div>you</div></div>
  <div v-click class="arrow">push →</div>
  <div v-click class="card glow text-center pulse"><div class="text-4xl">🗄️</div><div class="font-700">Registry</div><div class="opacity-60 text-xs">Docker Hub</div></div>
  <div v-click class="arrow">← pull</div>
  <div v-click class="card text-center"><div class="text-4xl">🧑‍💻</div><div>teammate</div></div>
</div>

<div v-click class="grid grid-cols-2 gap-6 mt-10">
  <div class="card text-center opacity-80">source code → <b>git</b> → GitHub</div>
  <div class="card text-center">image → <b>registry</b> → Docker Hub</div>
</div>

---
layout: center
class: text-center
---

<div class="chip mb-4">Breakout room</div>

# 🛠️ Workshop

<div class="text-left text-xl mt-6 inline-block">

<v-clicks>

1. `docker pull anuchito/docker-qadev`
2. `docker run` it on port **3000**
3. `docker exec` into the container, then `go test ./...`
4. Remove everything — and do it all once more

</v-clicks>

</div>

<div class="opacity-50 mt-6 text-sm">Use Docker Desktop to watch what happens</div>

---

# Two containers, one network

<div class="host mt-8">
  <div class="label">Host Machine</div>
  <div class="net">
    <div class="label">todo.network</div>
    <div class="grid grid-cols-[1fr_auto_1fr] items-center gap-4">
      <div v-click class="card glow text-center">
        <div class="text-3xl">🟩</div><div class="font-700">todo-app</div>
        <code class="text-xs">http://todo-app:3000</code>
      </div>
      <div v-click class="w-24"><div class="flow" /></div>
      <div v-click class="card text-center">
        <div class="text-3xl">🗃️</div><div class="font-700">todo-database</div>
        <code class="text-xs">http://mysql.todo.net:5432</code>
      </div>
    </div>
  </div>
</div>

---

# Config &amp; development loop

<div class="grid grid-cols-2 gap-6 mt-8">

<div v-click class="card">
  <div class="text-2xl">🧾 File DB → real DB</div>
  <div class="mt-2 opacity-70">The app finds the database through an environment variable:</div>

```bash
MYSQL_HOST=mysql.todo.net
```

</div>

<div v-click class="card">
  <div class="text-2xl">♻️ Auto reload</div>
  <div class="mt-2 opacity-70">Mount your source into the container — edit on the host, the app restarts by itself.</div>
</div>

</div>

---

# Many networks, many apps

<div class="host mt-6">
  <div class="label">Host Machine</div>
  <div class="grid grid-cols-2 gap-6">
    <div class="net">
      <div class="label">todo.network</div>
      <div class="flex gap-3 items-center justify-center">
        <div class="card text-center text-sm">todo-app<br><code>:3000</code></div>
        <div class="flow w-12" />
        <div class="card text-center text-sm">todo-database<br><code>:5432</code></div>
      </div>
    </div>
    <div v-click class="net alt">
      <div class="label">another.network</div>
      <div class="flex justify-center">
        <div class="card text-center text-sm">todo-app<br><code>:3000</code></div>
      </div>
    </div>
  </div>
  <div v-click class="flex justify-around mt-4 text-sm">
    <span class="chip">http://localhost:3000 ↑</span>
    <span class="chip">http://localhost:4000 ↑</span>
  </div>
</div>

<div v-click class="mt-4 text-center opacity-70">Same container port, different host ports — networks keep them apart.</div>

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
