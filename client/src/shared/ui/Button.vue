<template>
  <button
    :type="type"
    :disabled="disabled"
    :class="[
      'btn',
      `btn--${variant}`,
      `btn--${size}`,
      { 'btn--block': block, 'btn--loading': loading },
    ]"
    @click="$emit('click', $event)"
  >
    <span v-if="loading" class="btn__spinner" aria-hidden="true"></span>
    <slot />
  </button>
</template>

<script setup lang="ts">
interface Props {
  type?: "button" | "submit";
  disabled?: boolean;
  loading?: boolean;
  block?: boolean;
  variant?: "primary" | "ghost" | "subtle" | "danger";
  size?: "sm" | "md" | "lg";
}

withDefaults(defineProps<Props>(), {
  type: "button",
  disabled: false,
  loading: false,
  block: false,
  variant: "primary",
  size: "md",
});

defineEmits<{
  click: [event: MouseEvent];
}>();
</script>

<style scoped lang="scss">
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.6rem;
  border-radius: $radius-md;
  border: 1px solid transparent;
  font-weight: 500;
  font-family: $font-sans;
  letter-spacing: 0.005em;
  cursor: pointer;
  transition:
    transform 0.12s $ease-out,
    background 0.15s $ease-out,
    border-color 0.15s $ease-out,
    box-shadow 0.15s $ease-out,
    color 0.15s $ease-out;
  position: relative;
  line-height: 1;

  &:disabled,
  &--loading {
    opacity: 0.55;
    cursor: not-allowed;
  }

  &:not(:disabled):active {
    transform: translateY(1px);
  }

  &--block {
    width: 100%;
  }

  // Sizes
  &--sm {
    height: 3.2rem;
    padding: 0 1.2rem;
    font-size: 1.25rem;
  }
  &--md {
    height: 4rem;
    padding: 0 1.8rem;
    font-size: 1.35rem;
  }
  &--lg {
    height: 4.8rem;
    padding: 0 2.4rem;
    font-size: 1.45rem;
  }

  // Variants
  &--primary {
    background: $accent;
    color: $bg-rail;
    font-weight: 600;
    box-shadow:
      inset 0 1px 0 rgba(255, 255, 255, 0.18),
      inset 0 -1px 0 rgba(0, 0, 0, 0.12);

    &:not(:disabled):hover {
      background: $accent-glow;
    }
  }

  &--ghost {
    background: transparent;
    color: $text-secondary;
    border-color: $border-default;

    &:not(:disabled):hover {
      color: $text-primary;
      border-color: $border-accent;
      background: $accent-wash;
    }
  }

  &--subtle {
    background: $bg-inset;
    color: $text-secondary;
    border-color: $border-subtle;

    &:not(:disabled):hover {
      background: $bg-elevated;
      color: $text-primary;
      border-color: $border-default;
    }
  }

  &--danger {
    background: $severity-danger-soft;
    color: $severity-danger;
    border-color: rgba(241, 105, 105, 0.28);

    &:not(:disabled):hover {
      background: rgba(241, 105, 105, 0.2);
    }
  }

  &__spinner {
    width: 1.4rem;
    height: 1.4rem;
    border: 1.5px solid currentColor;
    border-top-color: transparent;
    border-radius: 50%;
    animation: btn-spin 0.8s linear infinite;
  }
}

@keyframes btn-spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
