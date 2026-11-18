<script setup lang="ts">
// "Step 3 of 7" progress dots shown at the top of each build-up step
defineProps<{ n: number; total?: number }>()
</script>

<template>
  <div class="stepper">
    <span v-for="i in (total ?? 7)" :key="i" class="dot" :class="{ done: i < n, now: i === n }">{{ i }}</span>
  </div>
</template>

<style scoped>
.stepper { display: flex; gap: 6px; align-items: center; }
.dot { width: 26px; height: 26px; border-radius: 50%; display: grid; place-items: center; font-size: .75rem; font-weight: 800;
  border: 2px solid var(--card-line); color: var(--muted); }
.dot.done { background: color-mix(in srgb, var(--m) 25%, transparent); border-color: var(--m); color: var(--m); }
.dot.now { background: var(--m); border-color: var(--m); color: #fff; box-shadow: 0 0 0 4px color-mix(in srgb, var(--m) 25%, transparent); }
</style>
