<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const breadcrumb = computed(() => {
  const name = (route.name as string) || "overview";
  const map: Record<string, string[]> = {
    home: ["Observatory", "Overview"],
    overview: ["Observatory", "Overview"],
  };
  return map[name] || ["Observatory", name];
});

const now = ref(new Date());
setInterval(() => (now.value = new Date()), 1000);

const clock = computed(() =>
  now.value.toLocaleTimeString("fr-FR", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  }),
);

const timeRanges = ["1h", "6h", "24h", "7d", "30d"];
const activeRange = ref("6h");
</script>

<template>
  <header class="topbar">
    <div class="topbar__left">
      <div class="topbar__crumbs">
        <template v-for="(crumb, i) in breadcrumb" :key="i">
          <span v-if="i > 0" class="topbar__crumbs-sep" aria-hidden="true">/</span>
          <span
            class="topbar__crumb"
            :class="{ 'topbar__crumb--current': i === breadcrumb.length - 1 }"
            >{{ crumb }}</span
          >
        </template>
      </div>
    </div>

    <button class="topbar__search" type="button" aria-label="Rechercher">
      <svg
        width="14"
        height="14"
        viewBox="0 0 14 14"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        aria-hidden="true"
      >
        <circle cx="6" cy="6" r="4" stroke="currentColor" stroke-width="1.4" />
        <path d="m9.5 9.5 3 3" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" />
      </svg>
      <span class="topbar__search-label">Rechercher un router, service, middleware…</span>
      <kbd>⌘</kbd><kbd>K</kbd>
    </button>

    <div class="topbar__right">
      <div class="topbar__ranges" role="group" aria-label="Plage temporelle">
        <button
          v-for="r in timeRanges"
          :key="r"
          type="button"
          class="topbar__range"
          :class="{ 'topbar__range--active': activeRange === r }"
          @click="activeRange = r"
        >
          {{ r }}
        </button>
      </div>

      <button class="topbar__refresh" type="button" aria-label="Rafraîchir">
        <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
          <path
            d="M13 8A5 5 0 1 1 8 3"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path d="M10 1.5 13 3l-1.5 3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
        <span class="mono topbar__refresh-text">5s</span>
      </button>

      <div class="topbar__clock mono" aria-label="Heure actuelle">
        <span class="topbar__clock-dot" aria-hidden="true"></span>
        {{ clock }}
      </div>

      <button class="topbar__user" type="button" aria-label="Compte">
        <div class="topbar__avatar" aria-hidden="true">CM</div>
      </button>
    </div>
  </header>
</template>

<style scoped lang="scss">
.topbar {
  position: sticky;
  top: 0;
  z-index: 50;
  display: flex;
  align-items: center;
  gap: 1.6rem;
  height: $topbar-height;
  padding: 0 2.4rem;
  background: rgba(6, 24, 42, 0.82);
  backdrop-filter: saturate(1.4) blur(14px);
  -webkit-backdrop-filter: saturate(1.4) blur(14px);
  border-bottom: 1px solid $border-subtle;

  &__left {
    display: flex;
    align-items: center;
    gap: 1.2rem;
  }

  &__crumbs {
    display: flex;
    align-items: center;
    gap: 0.7rem;
    font-size: 1.35rem;
  }

  &__crumb {
    color: $text-muted;
    font-weight: 500;

    &--current {
      color: $text-primary;
      font-weight: 600;
    }
  }

  &__crumbs-sep {
    color: $text-faint;
  }

  &__search {
    display: flex;
    align-items: center;
    gap: 1rem;
    flex: 1;
    max-width: 52rem;
    height: 3.6rem;
    padding: 0 1.2rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    border-radius: $radius-md;
    color: $text-muted;
    cursor: pointer;
    transition: all 0.15s $ease-out;
    font-family: $font-sans;

    &:hover {
      border-color: $border-default;
      color: $text-secondary;
      background: $bg-inset-hi;
    }

    &-label {
      flex: 1;
      text-align: left;
      font-size: 1.3rem;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }

    kbd {
      padding: 0.1rem 0.5rem;
      font-size: 1rem;
    }
  }

  &__right {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-left: auto;
  }

  &__ranges {
    display: inline-flex;
    align-items: center;
    padding: 0.3rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    border-radius: $radius-md;
    font-family: $font-mono;
  }

  &__range {
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

  &__refresh {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    height: 3.2rem;
    padding: 0 1rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    border-radius: $radius-md;
    color: $text-secondary;
    cursor: pointer;
    transition: all 0.15s $ease-out;

    &:hover {
      color: $accent-glow;
      border-color: $border-accent;
    }

    &-text {
      font-size: 1.15rem;
      font-weight: 500;
    }
  }

  &__clock {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    font-size: 1.2rem;
    color: $text-secondary;
    font-variant-numeric: tabular-nums;
    padding: 0 0.4rem;
  }

  &__clock-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 6px $severity-success;
    animation: graefik-pulse 1.4s $ease-in-out infinite;
  }

  &__user {
    background: transparent;
    border: 0;
    padding: 0;
    cursor: pointer;
  }

  &__avatar {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 3.2rem;
    height: 3.2rem;
    border-radius: $radius-pill;
    background: linear-gradient(135deg, $accent 0%, $accent-deep 100%);
    color: $bg-rail;
    font-family: $font-mono;
    font-size: 1.15rem;
    font-weight: 600;
    letter-spacing: 0.02em;
    box-shadow: 0 0 0 1px $border-accent, 0 4px 12px rgba(0, 194, 178, 0.2);
    transition: transform 0.18s $ease-spring;

    &:hover {
      transform: scale(1.05);
    }
  }

  @media (max-width: 1100px) {
    &__search-label {
      display: none;
    }
    &__ranges {
      display: none;
    }
  }
}
</style>
