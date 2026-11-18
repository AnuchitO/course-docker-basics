<script setup lang="ts">
withDefaults(defineProps<{ height?: number }>(), { height: 76 })
</script>

<template>
  <!-- Dockerfile page: folded corner + four code lines in the same colours as the image layers.
       Shapes only (no <text>): slide CSS can resize SVG text. -->
  <svg :height="height" viewBox="0 0 64 80" xmlns="http://www.w3.org/2000/svg">
    <path d="M6 4 H40 L58 22 V76 a4 4 0 0 1 -4 4 H10 a4 4 0 0 1 -4 -4 Z" fill="#ffffff" stroke="#94a3b8" stroke-width="2" stroke-linejoin="round" />
    <path d="M40 4 V18 a4 4 0 0 0 4 4 H58 Z" fill="#e2e8f0" stroke="#94a3b8" stroke-width="2" stroke-linejoin="round" />
    <!-- small Docker-blue tag -->
    <rect x="13" y="13" width="14" height="6" rx="3" fill="#2496ed" />
    <!-- code lines: FROM / RUN / COPY / RUN -->
    <rect x="13" y="32" width="38" height="6" rx="3" fill="#0e7490" />
    <rect x="13" y="42" width="28" height="6" rx="3" fill="#2563eb" />
    <rect x="13" y="52" width="22" height="6" rx="3" fill="#7c3aed" />
    <rect x="13" y="62" width="33" height="6" rx="3" fill="#db2777" />
  </svg>
</template>
