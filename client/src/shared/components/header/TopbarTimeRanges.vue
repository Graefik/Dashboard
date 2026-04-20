<script setup lang="ts">
import { ref } from "vue";

const ranges = ["1h", "6h", "24h", "7d", "30d"];
const active = ref("6h");
</script>

<template>
  <div class="ranges" role="group" aria-label="Plage temporelle">
    <button
      v-for="r in ranges"
      :key="r"
      type="button"
      class="ranges__btn"
      :class="{ 'ranges__btn--active': active === r }"
      @click="active = r"
    >
      {{ r }}
    </button>
  </div>
</template>

<style scoped lang="scss">
.ranges {
  display: inline-flex;
  align-items: center;
  padding: 0.3rem;
  background: $bg-inset;
  border: 1px solid $border-subtle;
  border-radius: $radius-md;
  font-family: $font-mono;

  &__btn {
    background: transparent;
    border: 0;
    padding: 0.45rem 0.9rem;
    color: $text-muted;
    font-size: 1.15rem;
    font-weight: 500;
    cursor: pointer;
    border-radius: $radius-sm;
    transition: all 0.15s $ease-out;
    letter-spacing: 0.02em;

    &:hover {
      color: $text-primary;
    }

    &--active {
      background: $accent-soft;
      color: $accent-glow;
      box-shadow: inset 0 0 0 1px $border-accent;
    }
  }

  @media (max-width: 1100px) {
    display: none;
  }
}
</style>
