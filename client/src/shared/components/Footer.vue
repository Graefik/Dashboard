<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref } from "vue";
import { useSystem } from "@/shared/system/useSystem";

const { info, connected } = useSystem();

const now = ref(new Date());
let timer: number | undefined;
onMounted(() => {
  timer = window.setInterval(() => (now.value = new Date()), 1000);
});
onBeforeUnmount(() => {
  if (timer !== undefined) clearInterval(timer);
});

// URL Traefik sans le schéma, façon "host"
const traefikHost = computed(() => {
  const url = info.value?.traefikURL;
  if (!url) return "…";
  return url.replace(/^https?:\/\//, "").replace(/\/$/, "");
});

const statusLabel = computed(() => (connected.value ? "Connected" : "Offline"));
const statusTone = computed(() => (connected.value ? "success" : "danger"));
</script>

<template>
  <footer class="statusbar">
    <div class="statusbar__left">
      <span class="statusbar__item" :class="`statusbar__item--${statusTone}`">
        <span class="statusbar__dot" aria-hidden="true"></span>
        <span class="statusbar__label">{{ statusLabel }}</span>
        <span class="statusbar__sep">·</span>
        <span class="statusbar__value mono">{{ traefikHost }}</span>
      </span>
    </div>

    <div class="statusbar__right">
      <span class="statusbar__item mono">© {{ now.getFullYear() }} Graefik</span>
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
  &__right {
    display: flex;
    align-items: center;
    gap: 1.8rem;
  }

  &__item {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    white-space: nowrap;

    &--success {
      color: $severity-success;
    }

    &--danger {
      color: $severity-danger;

      .statusbar__dot {
        background: $severity-danger;
        box-shadow: 0 0 6px $severity-danger;
        animation-duration: 0.8s;
      }
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
  }
}
</style>
