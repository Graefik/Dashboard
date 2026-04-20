<script setup lang="ts">
import type { Cmd } from "./commands.data";

defineProps<{
  cmd: Cmd;
  active: boolean;
}>();

defineEmits<{
  hover: [];
  select: [];
}>();
</script>

<template>
  <button
    type="button"
    class="item"
    :class="{ 'item--active': active }"
    role="option"
    :aria-selected="active"
    @mouseenter="$emit('hover')"
    @click="$emit('select')"
  >
    <span class="item__label">{{ cmd.label }}</span>
    <span v-if="cmd.hint" class="item__hint">{{ cmd.hint }}</span>
    <kbd v-if="cmd.shortcut" class="item__kbd">{{ cmd.shortcut }}</kbd>
  </button>
</template>

<style scoped lang="scss">
.item {
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  gap: 1rem;
  width: 100%;
  padding: 0.8rem 1rem;
  background: transparent;
  border: 0;
  border-radius: $radius-sm;
  color: $text-secondary;
  text-align: left;
  cursor: pointer;
  font-family: $font-sans;
  font-size: 1.35rem;
  transition:
    background 0.1s $ease-out,
    color 0.1s $ease-out;

  &--active {
    background: $accent-soft;
    color: $text-primary;

    .item__hint {
      color: $text-secondary;
    }

    .item__kbd {
      color: $accent-glow;
      border-color: $border-accent;
    }
  }

  &__label {
    color: inherit;
    font-weight: 500;
  }

  &__hint {
    color: $text-muted;
    font-size: 1.15rem;
    text-align: right;
    justify-self: end;
  }

  &__kbd {
    font-family: $font-mono;
    font-size: 1rem;
    padding: 0.15rem 0.5rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    color: $text-muted;
    border-radius: $radius-xs;
  }
}
</style>
