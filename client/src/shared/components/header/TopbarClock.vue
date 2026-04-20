<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref } from "vue";

const now = ref(new Date());
let timer: number | undefined;

onMounted(() => {
  timer = window.setInterval(() => (now.value = new Date()), 1000);
});
onBeforeUnmount(() => {
  if (timer !== undefined) clearInterval(timer);
});

const text = computed(() =>
  now.value.toLocaleTimeString("fr-FR", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  }),
);
</script>

<template>
  <div class="clock mono" aria-label="Heure actuelle">
    <span class="clock__dot" aria-hidden="true"></span>
    {{ text }}
  </div>
</template>

<style scoped lang="scss">
.clock {
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  font-size: 1.2rem;
  color: $text-secondary;
  font-variant-numeric: tabular-nums;
  padding: 0 0.4rem;

  &__dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 6px $severity-success;
    animation: graefik-pulse 1.4s $ease-in-out infinite;
  }
}
</style>
