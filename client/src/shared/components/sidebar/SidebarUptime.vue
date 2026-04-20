<script setup lang="ts">
const bars = Array.from({ length: 24 }, (_, i) =>
  Math.round(20 + Math.sin(i / 2) * 40 + Math.random() * 40),
);
</script>

<template>
  <div class="uptime">
    <div class="uptime__row">
      <div class="uptime__label">
        <span class="uptime__dot" aria-hidden="true"></span>
        Uptime
      </div>
      <div class="uptime__value">12d 04:37</div>
    </div>
    <div class="uptime__bars" aria-hidden="true">
      <span v-for="(h, i) in bars" :key="i" :style="{ '--h': `${h}%` }"></span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.uptime {
  padding: 1.6rem 2rem;
  border-top: 1px solid $border-subtle;
  background: linear-gradient(180deg, transparent, rgba(0, 194, 178, 0.03));

  &__row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 0.8rem;
  }

  &__label {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
  }

  &__dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 6px $severity-success;
  }

  &__value {
    font-family: $font-mono;
    font-size: 1.15rem;
    font-weight: 600;
    color: $text-primary;
    font-variant-numeric: tabular-nums;
  }

  &__bars {
    display: flex;
    align-items: flex-end;
    gap: 2px;
    height: 2.4rem;

    span {
      flex: 1;
      height: var(--h);
      min-height: 2px;
      background: linear-gradient(180deg, $accent 0%, $accent-deep 100%);
      border-radius: 1px;
      opacity: 0.6;
      transition: opacity 0.3s $ease-out;
    }

    &:hover span {
      opacity: 1;
    }
  }
}
</style>
