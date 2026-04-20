<script setup lang="ts">
defineProps<{ name?: string }>();
</script>

<template>
  <svg viewBox="0 0 20 20" fill="none" aria-hidden="true">
    <template v-if="name === 'requests'">
      <path
        d="M3 10h4l2-5 2 10 2-5h4"
        stroke="currentColor"
        stroke-width="1.6"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </template>
    <template v-else-if="name === 'latency'">
      <circle cx="10" cy="10" r="7" stroke="currentColor" stroke-width="1.5" />
      <path d="M10 6v4l2.5 2.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
    </template>
    <template v-else-if="name === 'connections'">
      <circle cx="5" cy="10" r="2.2" stroke="currentColor" stroke-width="1.5" />
      <circle cx="15" cy="10" r="2.2" stroke="currentColor" stroke-width="1.5" />
      <path d="M7.2 10h5.6" stroke="currentColor" stroke-width="1.5" />
    </template>
    <template v-else-if="name === 'errors'">
      <path
        d="M10 2.5 2 16.5h16L10 2.5z"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linejoin="round"
      />
      <path d="M10 8v3.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
      <circle cx="10" cy="13.5" r="0.8" fill="currentColor" />
    </template>
    <template v-else-if="name === 'throughput'">
      <path
        d="M3 14l4-4 3 3 4-6 3 4"
        stroke="currentColor"
        stroke-width="1.6"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </template>
  </svg>
</template>
