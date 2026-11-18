<script setup lang="ts">
defineProps<{ name?: string }>()
</script>

<template>
  <div class="flex flex-col items-center gap-1">
    <!-- flat container badge for grids (shapes only, no <text>) -->
    <svg viewBox="0 0 72 48" width="64" height="43" class="run-bob" xmlns="http://www.w3.org/2000/svg">
      <rect x="2" y="2" width="68" height="44" rx="8" fill="#2496ed" />
      <rect x="9" y="9" width="54" height="30" rx="4" fill="#ffffff" />
      <path d="M9 13a4 4 0 0 1 4-4h46a4 4 0 0 1 4 4v3H9Z" fill="#dbe4ee" />
      <circle cx="14.5" cy="12.5" r="1.4" fill="#ef4444" /><circle cx="19.5" cy="12.5" r="1.4" fill="#f59e0b" /><circle cx="24.5" cy="12.5" r="1.4" fill="#22c55e" />
      <circle cx="19" cy="28" r="5" fill="#2496ed" />
      <rect x="28" y="24" width="27" height="3.5" rx="1.75" fill="#94a3b8" />
      <rect x="28" y="30.5" width="17" height="3.5" rx="1.75" fill="#cbd5e1" />
    </svg>
    <div v-if="name" class="text-xs opacity-80">{{ name }}</div>
  </div>
</template>

<style scoped>
.run-bob { animation: bob 2.4s ease-in-out infinite; }
@keyframes bob { 50% { transform: translateY(-4px); } }
</style>
