<script setup lang="ts">
interface Chip {
  label: string;
  color: string;
  dash?: boolean;
}
defineProps<{ chips: Chip[] }>();
</script>

<template>
  <div class="chips">
    <span v-for="chip in chips" :key="chip.label" class="chips__chip">
      <span
        class="chips__dot"
        :class="{ 'chips__dot--dash': chip.dash }"
        :style="{ background: chip.dash ? 'transparent' : chip.color, borderColor: chip.color }"
      ></span>
      {{ chip.label }}
    </span>
  </div>
</template>

<style scoped lang="scss">
.chips {
  display: inline-flex;
  gap: 0.4rem;

  &__chip {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    padding: 0.25rem 0.6rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    border-radius: $radius-sm;
    font-family: $font-mono;
    font-size: 1.05rem;
    color: $text-secondary;
  }

  &__dot {
    width: 6px;
    height: 6px;
    border-radius: 2px;
    border: 1px solid transparent;

    &--dash {
      width: 8px;
      height: 0;
      border-style: dashed;
      border-width: 1px 0 0 0;
    }
  }
}
</style>
