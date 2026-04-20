<script setup lang="ts">
defineEmits<{ open: [] }>();
</script>

<template>
  <button
    class="search"
    type="button"
    aria-label="Ouvrir la palette de commandes"
    @click="$emit('open')"
  >
    <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
      <circle cx="6" cy="6" r="4" stroke="currentColor" stroke-width="1.4" />
      <path d="m9.5 9.5 3 3" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" />
    </svg>
    <span class="search__label">Rechercher un router, service, middleware…</span>
    <kbd>⌘</kbd><kbd>K</kbd>
  </button>
</template>

<style scoped lang="scss">
.search {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
  max-width: 52rem;
  height: 3.6rem;
  padding: 0 1.2rem;
  background: $bg-inset;
  border: 1px solid $border-subtle;
  border-radius: $radius-md;
  color: $text-muted;
  cursor: pointer;
  transition: all 0.15s $ease-out;
  font-family: $font-sans;

  &:hover {
    border-color: $border-default;
    color: $text-secondary;
    background: $bg-inset-hi;
  }

  &__label {
    flex: 1;
    text-align: left;
    font-size: 1.3rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  kbd {
    padding: 0.1rem 0.5rem;
    font-size: 1rem;
  }

  @media (max-width: 1100px) {
    &__label {
      display: none;
    }
  }
}
</style>
