<script setup lang="ts">
withDefaults(defineProps<{ name?: string, variant?: 'classic' | 'emoji' | 'svg', width?: string }>(), { variant: 'classic', width: '270px' })
</script>

<template>
  <!-- real SVG laptop: lid with bezel + camera, screen holds the slot, wide base with thumb notch -->
  <div v-if="variant === 'svg'" class="laptop">
    <svg viewBox="0 0 300 186" class="w-full" xmlns="http://www.w3.org/2000/svg">
      <defs>
        <linearGradient id="lap-base" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0" stop-color="#e2e8f0" /><stop offset="1" stop-color="#94a3b8" />
        </linearGradient>
        <linearGradient id="lap-lid" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0" stop-color="#334155" /><stop offset="1" stop-color="#0f172a" />
        </linearGradient>
      </defs>
      <ellipse cx="150" cy="176" rx="132" ry="6" fill="#000" opacity="0.18" />
      <rect x="38" y="4" width="224" height="148" rx="11" fill="url(#lap-lid)" stroke="#64748b" stroke-width="2" />
      <circle cx="150" cy="9.5" r="1.8" fill="#64748b" />
      <rect x="47" y="15" width="206" height="129" rx="4" fill="var(--bg)" stroke="var(--card-line)" />
      <foreignObject x="47" y="15" width="206" height="129">
        <div xmlns="http://www.w3.org/1999/xhtml" class="svg-screen"><slot /></div>
      </foreignObject>
      <path d="M6 152 H294 L288 166 Q287 170 282 170 H18 Q13 170 12 166 Z" fill="url(#lap-base)" stroke="#64748b" stroke-width="2" stroke-linejoin="round" />
      <path d="M126 152 H174 L171 158 H129 Z" fill="#64748b" />
    </svg>
    <div v-if="name" class="name">{{ name }}</div>
  </div>

  <!-- emoji laptop: big 💻 with the content next to it -->
  <div v-else-if="variant === 'emoji'" class="laptop emoji">
    <div class="text-6xl leading-none">💻</div>
    <div class="flex items-center justify-center gap-2 mt-2"><slot /></div>
    <div v-if="name" class="name">{{ name }}</div>
  </div>

  <div v-else class="laptop" :style="{ maxWidth: width }">
    <!-- optional person peeking over the lid -->
    <div v-if="$slots.person" class="person"><slot name="person" /></div>
    <div class="lid">
      <div class="screen">
        <slot />
      </div>
    </div>
    <div class="base"><div class="notch" /></div>
    <div v-if="name" class="name">{{ name }}</div>
  </div>
</template>

<style scoped>
.laptop { display: flex; flex-direction: column; align-items: center; width: 100%; max-width: 270px; margin: 0 auto; }
/* 16:10 display inside a thin bezel, like a real MacBook */
.lid {
  width: 80%; padding: 7px 7px 9px; border-radius: 12px 12px 2px 2px;
  background: #1e293b; border: 2px solid #64748b; border-bottom: none;
}
.screen {
  aspect-ratio: 16 / 10; width: 100%; border-radius: 4px; padding: 0.4rem;
  display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 0.35rem;
  background: var(--bg); border: 1px solid var(--card-line);
}
/* base is ~12% wider than the lid, thin, with a thumb notch */
.base {
  width: 100%; height: 11px; border-radius: 1px 1px 14px 14px / 1px 1px 10px 10px; position: relative;
  background: linear-gradient(#e2e8f0, #94a3b8); border: 2px solid #64748b;
}
.notch { position: absolute; left: 50%; top: -2px; width: 16%; height: 5px; transform: translateX(-50%);
  background: #64748b; border-radius: 0 0 6px 6px; }
.person {
  width: 42%; margin-bottom: -12px; margin-left: 10%; align-self: flex-start; position: relative; z-index: 0;
  animation: peek 0.6s cubic-bezier(.2,1.3,.35,1) both; animation-delay: var(--d, 0.4s);
}
@keyframes peek { from { transform: translateY(70px); opacity: 0; } 30% { opacity: 1; } to { transform: translateY(0); opacity: 1; } }
.lid { position: relative; z-index: 1; }
.name { margin-top: 0.4rem; font-weight: 700; color: var(--m); }
.svg-screen { width: 100%; height: 100%; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 8px; }
.emoji { text-align: center; }
</style>
