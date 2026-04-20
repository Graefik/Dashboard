<template>
  <button class="refresh" type="button" aria-label="Rafraîchir">
    <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
      <path d="M13 8A5 5 0 1 1 8 3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
      <path
        d="M10 1.5 13 3l-1.5 3"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
    <span class="mono refresh__text">5s</span>
  </button>
</template>

<style scoped lang="scss">
.refresh {
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  height: 3.2rem;
  padding: 0 1rem;
  background: $bg-inset;
  border: 1px solid $border-subtle;
  border-radius: $radius-md;
  color: $text-secondary;
  cursor: pointer;
  transition: all 0.15s $ease-out;

  &:hover {
    color: $accent-glow;
    border-color: $border-accent;
  }

  &__text {
    font-size: 1.15rem;
    font-weight: 500;
  }
}
</style>
