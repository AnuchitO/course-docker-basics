const $ = (id) => document.getElementById(id);
const cv = $('cv'), ctx = cv.getContext('2d');
let themes = [], current = null, matrix = null, verified = false, timer;

// ---- contrast (WCAG ratio) -------------------------------------------------
const lum = (hex) => {
  const [r, g, b] = [1, 3, 5].map((i) => parseInt(hex.slice(i, i + 2), 16) / 255)
    .map((v) => (v <= 0.03928 ? v / 12.92 : ((v + 0.055) / 1.055) ** 2.4));
  return 0.2126 * r + 0.7152 * g + 0.0722 * b;
};
const contrast = (a, b) => { const [x, y] = [lum(a), lum(b)].sort((p, q) => q - p); return (x + 0.05) / (y + 0.05); };

// ---- API -------------------------------------------------------------------
async function api(path, opts) {
  const r = await fetch('/api' + path, opts);
  const j = await r.json().catch(() => ({}));
  if (!r.ok) throw new Error(j.error || 'request failed');
  return j;
}

async function loadThemes() {
  const d = await api('/themes');
  themes = d.themes;
  $('host').textContent = d.host;
  $('themes').innerHTML = '';
  themes.forEach((t) => {
    const b = document.createElement('button');
    b.innerHTML = `<span class="dot" style="background:${t.foliage}"></span>${t.name}`;
    b.onclick = () => pick(t);
    b.dataset.slug = t.slug;
    $('themes').append(b);
  });
  const fromLink = new URLSearchParams(location.search).get('t');
  if (fromLink) { try { const t = await api('/trees/' + fromLink); $('url').value = t.url; setColors(t); return refresh(); } catch (e) {} }
  pick(themes[0]);
}

function setColors(c) { $('cFoliage').value = c.foliage; $('cTrunk').value = c.trunk; $('cBg').value = c.background; }
function pick(t) {
  current = t; setColors(t);
  document.querySelectorAll('#themes button').forEach((b) => b.setAttribute('aria-pressed', b.dataset.slug === t.slug));
  refresh();
}

// ---- generate (debounced) --------------------------------------------------
function refresh() { clearTimeout(timer); timer = setTimeout(generate, 250); }

async function generate() {
  verified = false; setButtons();
  const url = $('url').value.trim();
  try {
    const d = await api('/qr', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ url }) });
    $('err').textContent = ''; matrix = d.matrix; verified = d.verified; $('host').textContent = d.host;
    $('status').textContent = '✔ verified: the code was decoded back to your link'; $('status').className = 'status ok';
  } catch (e) {
    $('err').textContent = e.message; $('status').textContent = ''; matrix = null;
  }
  draw(); checkContrast(); setButtons();
}

function checkContrast() {
  const ratio = contrast($('cFoliage').value, $('cBg').value);
  const bad = ratio < 3;
  $('warn').style.display = bad ? 'block' : 'none';
  if (bad) $('warnText').textContent = `Low contrast (${ratio.toFixed(1)}:1). Cameras may not scan this.`;
  return !bad;
}
$('fix').onclick = () => { $('cFoliage').value = '#14301a'; $('cBg').value = '#f6f1e7'; refresh(); };

// ---- draw: QR modules become leaves of a tree -------------------------------
function shade(hex, k) { // k<0 darker, k>0 lighter
  const c = [1, 3, 5].map((i) => parseInt(hex.slice(i, i + 2), 16));
  return '#' + c.map((v) => Math.max(0, Math.min(255, Math.round(v + (k < 0 ? v * k : (255 - v) * k)))).toString(16).padStart(2, '0')).join('');
}

function draw() {
  const W = cv.width, H = cv.height, bg = $('cBg').value, leaf = $('cFoliage').value, trunk = $('cTrunk').value;
  ctx.fillStyle = bg; ctx.fillRect(0, 0, W, H);
  if (!matrix) return;
  const n = matrix.length, quiet = 3, area = W - 40, cell = area / (n + quiet * 2);
  const ox = 20 + quiet * cell, oy = 20 + quiet * cell, r = cell * 0.5;
  // trunk under the code
  const codeBottom = oy + n * cell;
  ctx.fillStyle = trunk;
  ctx.fillRect(W / 2 - cell * 1.6, codeBottom + quiet * cell * 0.1, cell * 3.2, H - codeBottom - 20);
  ctx.fillStyle = shade(trunk, -0.25);
  ctx.fillRect(W / 2 + cell * 0.4, codeBottom + quiet * cell * 0.1, cell * 1.2, H - codeBottom - 20);

  const inFinder = (x, y) => (x < 8 && y < 8) || (x >= n - 8 && y < 8) || (x < 8 && y >= n - 8);
  for (let y = 0; y < n; y++) for (let x = 0; x < n; x++) {
    if (matrix[y][x] !== '1') continue;
    const px = ox + x * cell, py = oy + y * cell;
    if (inFinder(x, y)) { ctx.fillStyle = shade(leaf, -0.35); ctx.fillRect(px, py, cell + 0.5, cell + 0.5); continue; } // crisp finders
    // little 3D cube: shadow face, then lit top face
    ctx.fillStyle = shade(leaf, -0.4); ctx.beginPath(); ctx.roundRect(px + cell * 0.12, py + cell * 0.14, cell * 0.92, cell * 0.92, r * 0.5); ctx.fill();
    const g = ctx.createLinearGradient(px, py, px + cell, py + cell);
    g.addColorStop(0, shade(leaf, 0.18)); g.addColorStop(1, leaf);
    ctx.fillStyle = g; ctx.beginPath(); ctx.roundRect(px, py, cell * 0.95, cell * 0.95, r * 0.5); ctx.fill();
  }
}
['cFoliage', 'cTrunk', 'cBg'].forEach((id) => $(id).addEventListener('input', () => { draw(); checkContrast(); setButtons(); }));
$('url').addEventListener('input', refresh);

function setButtons() {
  const ok = verified && matrix && checkContrast();
  $('dl').disabled = !ok; $('share').disabled = !ok;
}

// ---- tilt with mouse / touch ------------------------------------------------
const tilt = $('tilt');
function move(e) {
  const p = e.touches ? e.touches[0] : e, b = tilt.getBoundingClientRect();
  const dx = (p.clientX - (b.left + b.width / 2)) / b.width, dy = (p.clientY - (b.top + b.height / 2)) / b.height;
  tilt.style.transform = `rotateY(${dx * 24}deg) rotateX(${-dy * 24}deg)`;
}
document.querySelector('.stage').addEventListener('mousemove', move);
document.querySelector('.stage').addEventListener('touchmove', move, { passive: true });
document.querySelector('.stage').addEventListener('mouseleave', () => (tilt.style.transform = ''));

// ---- download & share -------------------------------------------------------
$('dl').onclick = () => { const a = document.createElement('a'); a.download = 'magic-tree.png'; a.href = cv.toDataURL('image/png'); a.click(); };
$('share').onclick = async () => {
  try {
    const t = await api('/trees', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({
      url: $('url').value.trim(), theme: current ? current.slug : 'custom',
      foliage: $('cFoliage').value, trunk: $('cTrunk').value, background: $('cBg').value }) });
    const link = `${location.origin}/?t=${t.id}`;
    await navigator.clipboard.writeText(link).catch(() => {});
    $('status').textContent = '🔗 share link copied: ' + link; $('status').className = 'status ok';
  } catch (e) { $('status').textContent = e.message; $('status').className = 'status'; }
};

loadThemes().catch((e) => ($('err').textContent = 'API not reachable: ' + e.message));
