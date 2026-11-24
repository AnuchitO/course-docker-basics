<script setup lang="ts">
defineProps<{ name?: string, stamped?: boolean }>()
</script>

<template>
  <div class="ctr flex flex-col items-center gap-1" :class="{ 'is-stamped': stamped }">
    <!-- layers, bottom to top: blue container, wax imprint, app window, stamp (while pressing) -->
    <div class="run-bob relative" style="width: 64px; height: 43px">
      <svg viewBox="0 0 72 48" width="64" height="43" class="layer base" style="z-index: 0" xmlns="http://www.w3.org/2000/svg">
        <rect x="2" y="2" width="68" height="44" rx="8" fill="#2496ed" />
      </svg>
      <img v-if="stamped" src="/img/stamp-wax.jpg" class="wax" />
      <svg viewBox="0 0 72 48" width="64" height="43" class="layer app" style="z-index: 2" xmlns="http://www.w3.org/2000/svg">
        <rect x="9" y="9" width="54" height="30" rx="4" fill="#ffffff" />
        <path d="M9 13a4 4 0 0 1 4-4h46a4 4 0 0 1 4 4v3H9Z" fill="#dbe4ee" />
        <circle cx="14.5" cy="12.5" r="1.4" fill="#ef4444" /><circle cx="19.5" cy="12.5" r="1.4" fill="#f59e0b" /><circle cx="24.5" cy="12.5" r="1.4" fill="#22c55e" />
        <circle cx="19" cy="28" r="5" fill="#2496ed" />
        <rect x="28" y="24" width="27" height="3.5" rx="1.75" fill="#94a3b8" />
        <rect x="28" y="30.5" width="17" height="3.5" rx="1.75" fill="#cbd5e1" />
      </svg>
      <img v-if="stamped" src="/img/stamp-upright.png" class="stamper" />
    </div>
    <div v-if="name" class="text-xs opacity-80">{{ name }}</div>
  </div>
</template>

<style scoped>
.layer { position: absolute; top: 0; left: 0; padding: 0; }
.wax { position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%) rotate(-12deg); width: 3.4rem; height: 3.4rem; border-radius: 50%; object-fit: cover; opacity: .75; pointer-events: none; z-index: 1; }
.stamper { position: absolute; top: 50%; left: 50%; width: 5.5rem; height: 5.5rem; opacity: 0; pointer-events: none; z-index: 3; transform: translate(-50%, -50%); }
/* when this container is revealed by v-click: the stamp comes straight down, presses, lifts away, and the wax appears */
.ctr.slidev-vclick-hidden .wax { opacity: 0; }
.ctr:not(.slidev-vclick-hidden) .stamper { animation: stamp-press 1.7s ease-in-out both; }
.ctr:not(.slidev-vclick-hidden) .wax { animation: wax-appear 1.7s ease-out both; }
@keyframes stamp-press {
  0%   { opacity: 0; transform: translate(-50%, -190%); }
  25%  { opacity: 1; transform: translate(-50%, -140%); }
  45%  { opacity: 1; transform: translate(-50%, -50%) scale(.94); }
  58%  { opacity: 1; transform: translate(-50%, -50%) scale(.94); }
  100% { opacity: 0; transform: translate(-50%, -210%); }
}
@keyframes wax-appear {
  0%, 44% { opacity: 0; transform: translate(-50%, -50%) rotate(-12deg) scale(.7); }
  54%     { opacity: .75; transform: translate(-50%, -50%) rotate(-12deg) scale(1.12); }
  66%, 100% { opacity: .75; transform: translate(-50%, -50%) rotate(-12deg) scale(1); }
}
/* order: stamp presses -> blue container + wax appear -> application icon pops in on top (and is drawn bigger) */
.layer.app { box-shadow: none !important; }
.ctr.is-stamped .app { transform: scale(1.12); }
.ctr.is-stamped:not(.slidev-vclick-hidden) .base { animation: base-appear 1.7s ease-out both; }
.ctr.is-stamped:not(.slidev-vclick-hidden) .app { animation: app-appear 1.7s ease-out both; }
@keyframes base-appear {
  0%, 43% { opacity: 0; transform: scale(.7); }
  54%     { opacity: 1; transform: scale(1.06); }
  66%, 100% { opacity: 1; transform: scale(1); }
}
@keyframes app-appear {
  0%, 68% { opacity: 0; transform: scale(.7); }
  82%     { opacity: 1; transform: scale(1.2); }
  92%, 100% { opacity: 1; transform: scale(1.12); }
}
.run-bob { animation: bob 2.4s ease-in-out infinite; }
@keyframes bob { 50% { transform: translateY(-4px); } }
</style>
