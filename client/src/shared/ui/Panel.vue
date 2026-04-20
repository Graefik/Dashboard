<script setup lang="ts">
interface Props {
  title?: string;
  eyebrow?: string;
  subtitle?: string;
  span?: 1 | 2 | 3 | 4 | 6 | 12;
  pad?: "tight" | "default" | "loose" | "none";
  live?: boolean;
}

withDefaults(defineProps<Props>(), {
  span: 6,
  pad: "default",
  live: false,
});
</script>

<template>
  <section
    class="panel"
    :class="[`panel--span-${span}`, `panel--pad-${pad}`]"
    :style="`--span: ${span}`"
  >
    <header v-if="title || $slots.header || eyebrow" class="panel__header">
      <div class="panel__titles">
        <div v-if="eyebrow || live" class="panel__eyebrow">
          <span v-if="live" class="panel__live" aria-label="Live">
            <span class="panel__live-dot" aria-hidden="true"></span>
            LIVE
          </span>
          <span v-if="eyebrow">{{ eyebrow }}</span>
        </div>
        <h3 v-if="title" class="panel__title">{{ title }}</h3>
        <p v-if="subtitle" class="panel__subtitle">{{ subtitle }}</p>
      </div>
      <div v-if="$slots.actions" class="panel__actions">
        <slot name="actions" />
      </div>
    </header>

    <div class="panel__body">
      <slot />
    </div>

    <footer v-if="$slots.footer" class="panel__footer">
      <slot name="footer" />
    </footer>
  </section>
</template>

<style lang="scss" scoped>
.panel {
  position: relative;
  grid-column: span var(--span);
  display: flex;
  flex-direction: column;
  background: linear-gradient(
    180deg,
    $bg-surface 0%,
    $bg-surface-2 100%
  );
  border: 1px solid $border-subtle;
  border-radius: $radius-lg;
  box-shadow: $shadow-md, $shadow-inset;
  transition:
    border-color 0.2s $ease-out,
    transform 0.2s $ease-out,
    box-shadow 0.2s $ease-out;
  overflow: hidden;

  &::before {
    content: "";
    position: absolute;
    inset: 0;
    border-radius: inherit;
    padding: 1px;
    background: linear-gradient(
      135deg,
      rgba(0, 194, 178, 0.18) 0%,
      transparent 30%,
      transparent 70%,
      rgba(0, 194, 178, 0.08) 100%
    );
    -webkit-mask:
      linear-gradient(#000 0 0) content-box,
      linear-gradient(#000 0 0);
    -webkit-mask-composite: xor;
    mask-composite: exclude;
    pointer-events: none;
    opacity: 0;
    transition: opacity 0.25s $ease-out;
  }

  &:hover::before {
    opacity: 1;
  }

  &__header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 1.6rem;
    padding: 1.6rem 1.8rem 1.2rem;
    border-bottom: 1px solid $border-subtle;
  }

  &__titles {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
    min-width: 0;
  }

  &__eyebrow {
    display: inline-flex;
    align-items: center;
    gap: 0.8rem;
    font-family: $font-mono;
    font-size: 1rem;
    font-weight: 500;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.12em;
  }

  &__live {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.2rem 0.55rem;
    background: $severity-danger-soft;
    color: $severity-danger;
    border-radius: $radius-sm;
    font-size: 0.95rem;
    font-weight: 600;
    letter-spacing: 0.15em;
  }

  &__live-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-danger;
    box-shadow: 0 0 6px $severity-danger;
    animation: graefik-pulse 1.2s $ease-in-out infinite;
  }

  &__title {
    font-size: 1.5rem;
    font-weight: 600;
    color: $text-primary;
    letter-spacing: -0.015em;
  }

  &__subtitle {
    font-size: 1.2rem;
    color: $text-muted;
    margin: 0;
  }

  &__actions {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    flex-shrink: 0;
  }

  &__footer {
    padding: 1.2rem 1.8rem;
    border-top: 1px solid $border-subtle;
    background: $bg-inset;
    font-family: $font-mono;
    font-size: 1.1rem;
    color: $text-muted;
  }

  // Padding variants for body
  &--pad-tight &__body {
    padding: 1rem;
  }
  &--pad-default &__body {
    padding: 1.6rem 1.8rem;
  }
  &--pad-loose &__body {
    padding: 2.4rem;
  }
  &--pad-none &__body {
    padding: 0;
  }

  // Responsive collapse
  @media (max-width: 1280px) {
    &--span-6 {
      grid-column: span 6;
    }
  }
  @media (max-width: 900px) {
    & {
      grid-column: span 12 !important;
    }
  }
}
</style>
