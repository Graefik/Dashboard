<script setup lang="ts">
interface Props {
  bleed?: boolean;
}
withDefaults(defineProps<Props>(), { bleed: false });
</script>

<template>
  <div class="wrapper" :class="{ 'wrapper--bleed': bleed }">
    <slot />
  </div>
</template>

<style lang="scss" scoped>
.wrapper {
  width: 100%;
  margin: 0 auto;
  padding: 0 2.4rem;
  max-width: 144rem;

  &--bleed {
    max-width: none;
    padding: 0;
  }
}
</style>
