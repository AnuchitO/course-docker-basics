<script setup lang="ts">
withDefaults(defineProps<{
  skin?: string, hair?: string, shirt?: string, hairStyle?: 'short' | 'long' | 'bun' | 'curly' | 'cap'
}>(), { skin: '#f1c27d', hair: '#2b2118', shirt: '#2563eb', hairStyle: 'short' })
</script>

<template>
  <!-- head and shoulders, shapes only (no <text>) -->
  <svg viewBox="0 0 120 92" class="w-full block" xmlns="http://www.w3.org/2000/svg">
    <!-- hair behind the head -->
    <path v-if="hairStyle === 'long'" d="M38 30 C34 62 38 70 46 70 L74 70 C82 70 86 62 82 30 Z" :fill="hair" />
    <!-- shoulders + neck -->
    <path d="M6 92 C6 64 30 56 60 56 C90 56 114 64 114 92 Z" :fill="shirt" />
    <path d="M48 58 L60 70 L72 58 Z" fill="#fff" opacity="0.35" />
    <rect x="52" y="44" width="16" height="16" rx="6" :fill="skin" />
    <rect x="52" y="44" width="16" height="7" :fill="'#000'" opacity="0.1" />
    <!-- ears + head -->
    <circle cx="40.5" cy="35" r="4" :fill="skin" /><circle cx="79.5" cy="35" r="4" :fill="skin" />
    <ellipse cx="60" cy="33" rx="19" ry="21" :fill="skin" />
    <!-- hair -->
    <template v-if="hairStyle === 'short' || hairStyle === 'long' || hairStyle === 'bun'">
      <path d="M40.5 32 C38 9 82 9 79.5 32 C75 22 67 19 60 19 C53 19 45 22 40.5 32 Z" :fill="hair" />
    </template>
    <circle v-if="hairStyle === 'bun'" cx="60" cy="9" r="8" :fill="hair" />
    <template v-if="hairStyle === 'curly'">
      <circle cx="44" cy="19" r="9" :fill="hair" /><circle cx="60" cy="13" r="10" :fill="hair" /><circle cx="76" cy="19" r="9" :fill="hair" />
      <circle cx="41" cy="30" r="6" :fill="hair" /><circle cx="79" cy="30" r="6" :fill="hair" />
    </template>
    <template v-if="hairStyle === 'cap'">
      <path d="M40 27 C40 6 80 6 80 27 Z" :fill="hair" />
      <path d="M58 25 C70 22 92 24 94 29 C84 31 68 30 58 29 Z" :fill="hair" />
    </template>
    <!-- face -->
    <circle cx="52" cy="35" r="2.2" fill="#1e293b" /><circle cx="68" cy="35" r="2.2" fill="#1e293b" />
    <path d="M52 43 Q60 50 68 43" fill="none" stroke="#1e293b" stroke-width="2" stroke-linecap="round" />
  </svg>
</template>
