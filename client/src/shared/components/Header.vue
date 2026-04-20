<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import CommandPalette from "@/shared/components/CommandPalette.vue";
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "reka-ui";

const route = useRoute();
const router = useRouter();

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

const paletteOpen = ref(false);
const onNavigate = (to: string) => router.push(to);
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

    <button
      class="topbar__search"
      type="button"
      aria-label="Ouvrir la palette de commandes"
      @click="paletteOpen = true"
    >
      <svg
        width="14"
        height="14"
        viewBox="0 0 14 14"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        aria-hidden="true"
      >
        <circle cx="6" cy="6" r="4" stroke="currentColor" stroke-width="1.4" />
        <path
          d="m9.5 9.5 3 3"
          stroke="currentColor"
          stroke-width="1.4"
          stroke-linecap="round"
        />
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
          <path d="M13 8A5 5 0 1 1 8 3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
          <path
            d="M10 1.5 13 3l-1.5 3"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <span class="mono topbar__refresh-text">5s</span>
      </button>

      <div class="topbar__clock mono" aria-label="Heure actuelle">
        <span class="topbar__clock-dot" aria-hidden="true"></span>
        {{ clock }}
      </div>

      <DropdownMenuRoot>
        <DropdownMenuTrigger as-child>
          <button class="topbar__user" type="button" aria-label="Compte">
            <div class="topbar__avatar" aria-hidden="true">CM</div>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuPortal>
          <DropdownMenuContent
            class="dd"
            :side-offset="10"
            align="end"
          >
            <DropdownMenuLabel class="dd__label">
              <div class="dd__user">
                <div class="dd__user-name">Clément Maillet</div>
                <div class="dd__user-mail mono">cmaillet@inovera.fr</div>
              </div>
            </DropdownMenuLabel>
            <DropdownMenuSeparator class="dd__sep" />
            <DropdownMenuItem class="dd__item" @select="router.push('/')">
              <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
                <circle cx="8" cy="6" r="2.5" stroke="currentColor" stroke-width="1.4" />
                <path
                  d="M3 13.5c.8-2.5 2.8-4 5-4s4.2 1.5 5 4"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                />
              </svg>
              Profil
            </DropdownMenuItem>
            <DropdownMenuItem class="dd__item">
              <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
                <circle cx="8" cy="8" r="2" stroke="currentColor" stroke-width="1.4" />
                <path
                  d="M8 1.5v1.5M8 13v1.5M1.5 8H3M13 8h1.5M3.6 3.6l1 1M11.4 11.4l1 1M3.6 12.4l1-1M11.4 4.6l1-1"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                />
              </svg>
              Préférences
              <kbd class="dd__kbd">⌘ ,</kbd>
            </DropdownMenuItem>
            <DropdownMenuItem class="dd__item">
              <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
                <path
                  d="M6 5V3.5a1 1 0 0 1 1-1h5.5a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V11"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                />
                <path
                  d="M2 8h8m0 0L7 5m3 3-3 3"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
              Documentation
            </DropdownMenuItem>
            <DropdownMenuSeparator class="dd__sep" />
            <DropdownMenuItem class="dd__item dd__item--danger">
              <svg viewBox="0 0 16 16" width="14" height="14" fill="none" aria-hidden="true">
                <path
                  d="M6 13.5H3.5A1.5 1.5 0 0 1 2 12V4A1.5 1.5 0 0 1 3.5 2.5H6"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                />
                <path
                  d="M10 11 13 8 10 5M13 8H6"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
              Se déconnecter
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenuPortal>
      </DropdownMenuRoot>
    </div>

    <CommandPalette
      v-model:open="paletteOpen"
      @navigate="onNavigate"
    />
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
    background: $accent;
    color: $bg-rail;
    font-family: $font-mono;
    font-size: 1.15rem;
    font-weight: 600;
    letter-spacing: 0.02em;
    box-shadow: 0 0 0 1px $border-accent,
      inset 0 1px 0 rgba(255, 255, 255, 0.18);
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

