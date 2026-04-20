<script setup lang="ts">
import { ref } from "vue";
import { useRoute } from "vue-router";
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "reka-ui";

interface NavItem {
  label: string;
  to: string;
  icon: string;
  badge?: string;
}

const route = useRoute();

const sections: { title?: string; items: NavItem[] }[] = [
  {
    items: [{ label: "Overview", to: "/", icon: "overview" }],
  },
  {
    title: "Console",
    items: [
      { label: "Logs", to: "/logs", icon: "logs" },
      { label: "Alerts", to: "/alerts", icon: "alerts" },
      { label: "Settings", to: "/settings", icon: "settings" },
    ],
  },
];

const isActive = (to: string) => {
  if (to === "/") return route.path === "/";
  return route.path.startsWith(to);
};

// ── Instances (mock) ────────────────────────────────────────────────────
interface Instance {
  id: string;
  name: string;
  endpoint: string;
  status: "success" | "warning" | "danger";
}
const instances: Instance[] = [
  { id: "prod-eu", name: "prod-eu-west", endpoint: "traefik.prod.eu-west.local", status: "success" },
  { id: "prod-us", name: "prod-us-east", endpoint: "traefik.prod.us-east.local", status: "success" },
  { id: "staging", name: "staging", endpoint: "traefik.staging.local", status: "warning" },
  { id: "edge", name: "edge-fallback", endpoint: "traefik.edge.local", status: "danger" },
];
const activeInstance = ref<Instance>(instances[0]);
const selectInstance = (i: Instance) => (activeInstance.value = i);
</script>

<template>
  <aside class="rail">
    <div class="rail__brand">
      <img
        src="@/assets/icon.png"
        alt="Graefik"
        class="rail__logo"
      />
      <div class="rail__brand-text">
        <span class="rail__brand-name">Graefik</span>
        <span class="rail__brand-meta">v0.1 · on-prem</span>
      </div>
    </div>

    <div class="rail__env">
      <div class="rail__env-label">Instance</div>
      <DropdownMenuRoot>
        <DropdownMenuTrigger as-child>
          <button class="rail__env-select" type="button">
            <span
              class="rail__env-dot"
              :class="`rail__env-dot--${activeInstance.status}`"
              aria-hidden="true"
            ></span>
            <span class="rail__env-name">{{ activeInstance.name }}</span>
            <svg
              class="rail__env-chevron"
              width="12"
              height="12"
              viewBox="0 0 12 12"
              fill="none"
              aria-hidden="true"
            >
              <path
                d="M3 4.5l3 3 3-3"
                stroke="currentColor"
                stroke-width="1.4"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuPortal>
          <DropdownMenuContent class="dd dd--env" :side-offset="8" align="start">
            <DropdownMenuLabel class="dd__label dd__label--env">
              Instances disponibles
            </DropdownMenuLabel>
            <DropdownMenuSeparator class="dd__sep" />
            <DropdownMenuItem
              v-for="i in instances"
              :key="i.id"
              class="dd__item dd__item--env"
              :data-active="i.id === activeInstance.id ? 'true' : undefined"
              @select="selectInstance(i)"
            >
              <span
                class="dd__env-dot"
                :class="`dd__env-dot--${i.status}`"
                aria-hidden="true"
              ></span>
              <div class="dd__env-info">
                <span class="dd__env-name">{{ i.name }}</span>
                <span class="dd__env-endpoint mono">{{ i.endpoint }}</span>
              </div>
              <svg
                v-if="i.id === activeInstance.id"
                class="dd__env-check"
                width="14"
                height="14"
                viewBox="0 0 14 14"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="m3 7.5 2.5 2.5L11 4.5"
                  stroke="currentColor"
                  stroke-width="1.8"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            </DropdownMenuItem>
            <DropdownMenuSeparator class="dd__sep" />
            <DropdownMenuItem class="dd__item dd__item--env dd__item--add">
              <svg
                width="14"
                height="14"
                viewBox="0 0 14 14"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="M7 3v8M3 7h8"
                  stroke="currentColor"
                  stroke-width="1.6"
                  stroke-linecap="round"
                />
              </svg>
              Ajouter une instance
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenuPortal>
      </DropdownMenuRoot>
    </div>

    <nav class="rail__nav">
      <div
        v-for="(section, sIdx) in sections"
        :key="sIdx"
        class="rail__section"
      >
        <div v-if="section.title" class="rail__section-title">
          {{ section.title }}
        </div>
        <ul class="rail__list">
          <li v-for="item in section.items" :key="item.to">
            <a
              :href="item.to"
              class="rail__link"
              :class="{ 'rail__link--active': isActive(item.to) }"
              @click.prevent
            >
              <span class="rail__icon" aria-hidden="true">
                <!-- Overview -->
                <svg
                  v-if="item.icon === 'overview'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <path
                    d="M3 10.5l2.5-3 3 3.5 3.5-5 5 7"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                  <path
                    d="M3 17h14"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    opacity=".4"
                  />
                </svg>
                <!-- Traffic -->
                <svg
                  v-else-if="item.icon === 'traffic'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <path
                    d="M3 6h14M3 10h10M3 14h6"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                  />
                </svg>
                <!-- Latency -->
                <svg
                  v-else-if="item.icon === 'latency'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <circle
                    cx="10"
                    cy="10"
                    r="7"
                    stroke="currentColor"
                    stroke-width="1.5"
                  />
                  <path
                    d="M10 6v4l2.5 2.5"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                  />
                </svg>
                <!-- Errors -->
                <svg
                  v-else-if="item.icon === 'errors'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <path
                    d="M10 2.5 2 16.5h16L10 2.5z"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linejoin="round"
                  />
                  <path
                    d="M10 8v3.5"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                  />
                  <circle cx="10" cy="13.5" r="0.8" fill="currentColor" />
                </svg>
                <!-- Router -->
                <svg
                  v-else-if="item.icon === 'router'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <rect
                    x="3"
                    y="8"
                    width="14"
                    height="8"
                    rx="1.5"
                    stroke="currentColor"
                    stroke-width="1.5"
                  />
                  <path d="M7 8V5M10 8V4M13 8V5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
                  <circle cx="6" cy="12" r="0.8" fill="currentColor" />
                  <circle cx="9" cy="12" r="0.8" fill="currentColor" />
                </svg>
                <!-- Services -->
                <svg
                  v-else-if="item.icon === 'services'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <circle cx="5" cy="5" r="2" stroke="currentColor" stroke-width="1.5" />
                  <circle cx="15" cy="5" r="2" stroke="currentColor" stroke-width="1.5" />
                  <circle cx="5" cy="15" r="2" stroke="currentColor" stroke-width="1.5" />
                  <circle cx="15" cy="15" r="2" stroke="currentColor" stroke-width="1.5" />
                  <path d="M7 5h6M5 7v6M15 7v6M7 15h6" stroke="currentColor" stroke-width="1.5" />
                </svg>
                <!-- Middleware -->
                <svg
                  v-else-if="item.icon === 'middleware'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <rect x="3" y="4" width="14" height="4" rx="1" stroke="currentColor" stroke-width="1.5" />
                  <rect x="3" y="12" width="14" height="4" rx="1" stroke="currentColor" stroke-width="1.5" />
                  <path d="M10 8v4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
                </svg>
                <!-- Cert -->
                <svg
                  v-else-if="item.icon === 'cert'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <path
                    d="M10 2l6 2.5v5c0 4-2.5 7-6 8-3.5-1-6-4-6-8v-5L10 2z"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linejoin="round"
                  />
                  <path d="M7 10l2 2 4-4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
                <!-- Logs -->
                <svg
                  v-else-if="item.icon === 'logs'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <rect x="3" y="3" width="14" height="14" rx="1.5" stroke="currentColor" stroke-width="1.5" />
                  <path d="M6 7h8M6 10h8M6 13h5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
                </svg>
                <!-- Alerts -->
                <svg
                  v-else-if="item.icon === 'alerts'"
                  viewBox="0 0 20 20"
                  fill="none"
                >
                  <path
                    d="M10 3a5 5 0 0 0-5 5v3l-1.5 2.5h13L15 11V8a5 5 0 0 0-5-5z"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linejoin="round"
                  />
                  <path d="M8 16a2 2 0 0 0 4 0" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
                </svg>
                <!-- Settings -->
                <svg v-else viewBox="0 0 20 20" fill="none">
                  <circle cx="10" cy="10" r="2.5" stroke="currentColor" stroke-width="1.5" />
                  <path
                    d="M10 2v2M10 16v2M4.5 4.5l1.4 1.4M14.1 14.1l1.4 1.4M2 10h2M16 10h2M4.5 15.5l1.4-1.4M14.1 5.9l1.4-1.4"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                  />
                </svg>
              </span>
              <span class="rail__label">{{ item.label }}</span>
              <span
                v-if="item.badge"
                class="rail__badge"
                :class="`rail__badge--${item.icon === 'errors' ? 'danger' : 'accent'}`"
                >{{ item.badge }}</span
              >
            </a>
          </li>
        </ul>
      </div>
    </nav>

    <div class="rail__footer">
      <div class="rail__uptime">
        <div class="rail__uptime-label">
          <span class="rail__uptime-dot" aria-hidden="true"></span>
          Uptime
        </div>
        <div class="rail__uptime-value">12d 04:37</div>
      </div>
      <div class="rail__bars" aria-hidden="true">
        <span v-for="n in 24" :key="n" :style="`--h: ${20 + Math.sin(n / 2) * 40 + Math.random() * 40}%`"></span>
      </div>
    </div>
  </aside>
</template>

<style scoped lang="scss">
.rail {
  position: sticky;
  top: 0;
  width: $rail-width;
  height: 100vh;
  flex-shrink: 0;
  background: linear-gradient(180deg, $bg-rail 0%, $bg-rail-2 100%);
  border-right: 1px solid $border-subtle;
  display: flex;
  flex-direction: column;
  padding: 2rem 0 0;
  overflow-y: auto;
  overflow-x: hidden;

  &__brand {
    display: flex;
    align-items: center;
    gap: 1.1rem;
    padding: 0 2rem 1.8rem;
  }

  &__logo {
    width: 4rem;
    height: 4rem;
    flex-shrink: 0;
    object-fit: contain;
    border-radius: $radius-sm;
    filter: drop-shadow(0 2px 6px rgba(0, 0, 0, 0.5));
  }

  &__brand-text {
    display: flex;
    flex-direction: column;
    line-height: 1.1;
  }

  &__brand-name {
    font-size: 1.7rem;
    font-weight: 600;
    color: $text-primary;
    letter-spacing: -0.02em;
  }

  &__brand-meta {
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    margin-top: 0.3rem;
    letter-spacing: 0.04em;
  }

  &__env {
    padding: 0 2rem 2rem;
    border-bottom: 1px solid $border-subtle;
    margin-bottom: 1.6rem;
  }

  &__env-label {
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    margin-bottom: 0.6rem;
  }

  &__env-select {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    width: 100%;
    padding: 0.8rem 1rem;
    background: $bg-surface;
    border: 1px solid $border-default;
    border-radius: $radius-md;
    color: $text-primary;
    cursor: pointer;
    transition: all 0.15s $ease-out;
    font-family: $font-mono;
    font-size: 1.2rem;

    &:hover {
      border-color: $border-accent;
      background: $bg-elevated;
    }
  }

  &__env-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    flex-shrink: 0;
    animation: graefik-pulse 2.2s $ease-in-out infinite;

    &--success {
      background: $severity-success;
      box-shadow: 0 0 8px $severity-success;
    }
    &--warning {
      background: $severity-warning;
      box-shadow: 0 0 8px $severity-warning;
    }
    &--danger {
      background: $severity-danger;
      box-shadow: 0 0 8px $severity-danger;
    }
  }

  &__env-name {
    flex: 1;
    text-align: left;
  }

  &__env-chevron {
    color: $text-muted;
  }

  &__nav {
    flex: 1;
    padding: 0 1.2rem;
  }

  &__section {
    & + & {
      margin-top: 2.2rem;
      padding-top: 1.6rem;
      border-top: 1px solid $border-subtle;
    }
  }

  &__section-title {
    font-family: $font-mono;
    font-size: 1rem;
    font-weight: 500;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    padding: 0 0.8rem 0.8rem;
  }

  &__list {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
  }

  &__link {
    display: flex;
    align-items: center;
    gap: 1.1rem;
    padding: 0.85rem 0.9rem;
    border-radius: $radius-md;
    color: $text-secondary;
    font-size: 1.35rem;
    font-weight: 500;
    position: relative;
    transition: all 0.15s $ease-out;

    &::before {
      content: "";
      position: absolute;
      left: -1.2rem;
      top: 50%;
      transform: translateY(-50%);
      width: 3px;
      height: 0;
      background: $accent;
      border-radius: 0 2px 2px 0;
      transition: height 0.2s $ease-out;
    }

    &:hover {
      color: $text-primary;
      background: rgba(255, 255, 255, 0.03);
    }

    &--active {
      color: $text-primary;
      background: linear-gradient(
        90deg,
        $accent-soft 0%,
        transparent 100%
      );

      &::before {
        height: 60%;
      }

      .rail__icon {
        color: $accent;
      }
    }
  }

  &__icon {
    display: inline-flex;
    width: 1.8rem;
    height: 1.8rem;
    flex-shrink: 0;

    svg {
      width: 100%;
      height: 100%;
    }
  }

  &__label {
    flex: 1;
  }

  &__badge {
    font-family: $font-mono;
    font-size: 1rem;
    font-weight: 600;
    padding: 0.15rem 0.55rem;
    border-radius: $radius-pill;
    line-height: 1.4;

    &--accent {
      background: $accent-soft;
      color: $accent-glow;
    }

    &--danger {
      background: $severity-danger-soft;
      color: $severity-danger;
    }
  }

  &__footer {
    padding: 1.6rem 2rem;
    border-top: 1px solid $border-subtle;
    background: linear-gradient(180deg, transparent, rgba(0, 194, 178, 0.03));
  }

  &__uptime {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 0.8rem;
  }

  &__uptime-label {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
  }

  &__uptime-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 6px $severity-success;
  }

  &__uptime-value {
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

<style lang="scss">
// ── Env dropdown overrides (global, partage les tokens .dd) ─────────────
.dd--env {
  min-width: 26rem;

  .dd__label--env {
    padding: 0.7rem 1rem 0.5rem;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
  }

  .dd__item--env {
    gap: 1rem;
  }

  .dd__item--env[data-active="true"] {
    background: $accent-soft;
    color: $text-primary;
  }

  .dd__env-dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    flex-shrink: 0;

    &--success {
      background: $severity-success;
      box-shadow: 0 0 6px $severity-success;
    }
    &--warning {
      background: $severity-warning;
      box-shadow: 0 0 6px $severity-warning;
    }
    &--danger {
      background: $severity-danger;
      box-shadow: 0 0 6px $severity-danger;
    }
  }

  .dd__env-info {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
    min-width: 0;
    flex: 1;
  }

  .dd__env-name {
    color: $text-primary;
    font-size: 1.3rem;
    font-weight: 500;
  }

  .dd__env-endpoint {
    color: $text-muted;
    font-size: 1.1rem;
  }

  .dd__env-check {
    color: $accent-glow;
    flex-shrink: 0;
  }

  .dd__item--add {
    color: $accent;
    font-weight: 500;

    svg {
      color: $accent;
    }
  }
}
</style>
