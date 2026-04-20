<script setup lang="ts">
defineProps<{ query: string }>();
</script>

<template>
  <div class="empty">
    <svg width="28" height="28" viewBox="0 0 28 28" fill="none" aria-hidden="true">
      <circle cx="12" cy="12" r="8" stroke="currentColor" stroke-width="1.3" />
      <path d="m18 18 5 5" stroke="currentColor" stroke-width="1.3" stroke-linecap="round" />
    </svg>
    <p>Aucun résultat pour <span class="mono">«&nbsp;{{ query }}&nbsp;»</span></p>
  </div>
</template>

<style scoped lang="scss">
.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.8rem;
  padding: 4rem 2rem;
  color: $text-muted;
  font-size: 1.3rem;
  text-align: center;

  p {
    margin: 0;
  }
}
</style>
