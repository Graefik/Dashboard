<script setup lang="ts">
defineProps<{ name: string }>();
</script>

<template>
  <svg viewBox="0 0 20 20" fill="none" aria-hidden="true">
    <template v-if="name === 'overview'">
      <path d="M3 10.5l2.5-3 3 3.5 3.5-5 5 7" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
      <path d="M3 17h14" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" opacity=".4" />
    </template>
    <template v-else-if="name === 'traffic'">
      <path d="M3 6h14M3 10h10M3 14h6" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else-if="name === 'latency'">
      <circle cx="10" cy="10" r="7" stroke="currentColor" stroke-width="1.5" />
      <path d="M10 6v4l2.5 2.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else-if="name === 'errors'">
      <path d="M10 2.5 2 16.5h16L10 2.5z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round" />
      <path d="M10 8v3.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
      <circle cx="10" cy="13.5" r="0.8" fill="currentColor" />
    </template>
    <template v-else-if="name === 'router'">
      <rect x="3" y="8" width="14" height="8" rx="1.5" stroke="currentColor" stroke-width="1.5" />
      <path d="M7 8V5M10 8V4M13 8V5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
      <circle cx="6" cy="12" r="0.8" fill="currentColor" />
      <circle cx="9" cy="12" r="0.8" fill="currentColor" />
    </template>
    <template v-else-if="name === 'services'">
      <circle cx="5" cy="5" r="2" stroke="currentColor" stroke-width="1.5" />
      <circle cx="15" cy="5" r="2" stroke="currentColor" stroke-width="1.5" />
      <circle cx="5" cy="15" r="2" stroke="currentColor" stroke-width="1.5" />
      <circle cx="15" cy="15" r="2" stroke="currentColor" stroke-width="1.5" />
      <path d="M7 5h6M5 7v6M15 7v6M7 15h6" stroke="currentColor" stroke-width="1.5" />
    </template>
    <template v-else-if="name === 'middleware'">
      <rect x="3" y="4" width="14" height="4" rx="1" stroke="currentColor" stroke-width="1.5" />
      <rect x="3" y="12" width="14" height="4" rx="1" stroke="currentColor" stroke-width="1.5" />
      <path d="M10 8v4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else-if="name === 'cert'">
      <path d="M10 2l6 2.5v5c0 4-2.5 7-6 8-3.5-1-6-4-6-8v-5L10 2z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round" />
      <path d="M7 10l2 2 4-4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
    </template>
    <template v-else-if="name === 'logs'">
      <rect x="3" y="3" width="14" height="14" rx="1.5" stroke="currentColor" stroke-width="1.5" />
      <path d="M6 7h8M6 10h8M6 13h5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else-if="name === 'alerts'">
      <path d="M10 3a5 5 0 0 0-5 5v3l-1.5 2.5h13L15 11V8a5 5 0 0 0-5-5z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round" />
      <path d="M8 16a2 2 0 0 0 4 0" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else>
      <circle cx="10" cy="10" r="2.5" stroke="currentColor" stroke-width="1.5" />
      <path d="M10 2v2M10 16v2M4.5 4.5l1.4 1.4M14.1 14.1l1.4 1.4M2 10h2M16 10h2M4.5 15.5l1.4-1.4M14.1 5.9l1.4-1.4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
  </svg>
</template>
