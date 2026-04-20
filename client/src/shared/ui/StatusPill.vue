<script setup lang="ts">
interface Props {
  tone?: "success" | "warning" | "danger" | "info" | "muted" | "accent";
  dot?: boolean;
  label: string;
}
withDefaults(defineProps<Props>(), {
  tone: "muted",
  dot: true,
});
</script>

<template>
  <span class="pill" :class="`pill--${tone}`">
    <span v-if="dot" class="pill__dot" aria-hidden="true"></span>
    {{ label }}
  </span>
</template>

<style lang="scss" scoped>
.pill {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.3rem 0.7rem;
  border-radius: $radius-sm;
  font-family: $font-mono;
  font-size: 1.05rem;
  font-weight: 500;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  border: 1px solid transparent;
  line-height: 1;

  &__dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: currentColor;
    box-shadow: 0 0 6px currentColor;
  }

  &--success {
    color: $severity-success;
    background: $severity-success-soft;
    border-color: rgba(45, 212, 167, 0.18);
  }
  &--warning {
    color: $severity-warning;
    background: $severity-warning-soft;
    border-color: rgba(245, 165, 36, 0.18);
  }
  &--danger {
    color: $severity-danger;
    background: $severity-danger-soft;
    border-color: rgba(241, 105, 105, 0.18);
  }
  &--info {
    color: $severity-info;
    background: $severity-info-soft;
    border-color: rgba(88, 166, 255, 0.18);
  }
  &--accent {
    color: $accent-glow;
    background: $accent-soft;
    border-color: $border-accent;
  }
  &--muted {
    color: $text-secondary;
    background: rgba(255, 255, 255, 0.04);
    border-color: $border-subtle;

    .pill__dot {
      background: $text-muted;
      box-shadow: none;
    }
  }
}
</style>
