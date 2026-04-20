<script setup lang="ts">
import { ref, computed } from "vue";

const now = ref(new Date());
setInterval(() => (now.value = new Date()), 1000);

const items = computed(() => [
  { label: "STATUS", value: "healthy", tone: "success" },
  { label: "API", value: "v0.1.0-rc.3", tone: "muted" },
  { label: "LATENCY", value: "p95 42ms", tone: "muted" },
  { label: "REGION", value: "eu-west-3", tone: "muted" },
  { label: "BUILD", value: "go1.22.4", tone: "accent" },
]);
</script>

<template>
  <footer class="statusbar">
    <div class="statusbar__left">
      <span class="statusbar__item statusbar__item--health">
        <span class="statusbar__dot" aria-hidden="true"></span>
        <span class="statusbar__label">Connected</span>
        <span class="statusbar__sep">·</span>
        <span class="statusbar__value mono">wss://graefik.local</span>
      </span>
    </div>

    <div class="statusbar__center">
      <span
        v-for="item in items.slice(1)"
        :key="item.label"
        class="statusbar__item"
      >
        <span class="statusbar__label">{{ item.label }}</span>
        <span class="statusbar__value mono" :class="`statusbar__value--${item.tone}`">{{
          item.value
        }}</span>
      </span>
    </div>

    <div class="statusbar__right">
      <span class="statusbar__item mono">
        © {{ now.getFullYear() }} Graefik
      </span>
    </div>
  </footer>
</template>

<style scoped lang="scss">
.statusbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 2rem;
  height: $statusbar-height;
  padding: 0 2rem;
  background: $bg-rail;
  border-top: 1px solid $border-subtle;
  font-family: $font-mono;
  font-size: 1.1rem;
  color: $text-muted;
  position: sticky;
  bottom: 0;

  &__left,
  &__center,
  &__right {
    display: flex;
    align-items: center;
    gap: 1.8rem;
  }

  &__center {
    @media (max-width: 900px) {
      display: none;
    }
  }

  &__item {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    white-space: nowrap;

    &--health {
      color: $severity-success;
    }
  }

  &__dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 6px $severity-success;
    animation: graefik-pulse 1.8s $ease-in-out infinite;
  }

  &__label {
    color: $text-faint;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    font-size: 1rem;
  }

  &__sep {
    color: $text-faint;
  }

  &__value {
    color: $text-secondary;
    font-weight: 500;

    &--success {
      color: $severity-success;
    }
    &--accent {
      color: $accent-glow;
    }
    &--muted {
      color: $text-secondary;
    }
  }
}
</style>
