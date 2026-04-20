<script setup lang="ts">
import { ref } from "vue";
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "reka-ui";

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

const active = ref<Instance>(instances[0]);
const select = (i: Instance) => (active.value = i);
</script>

<template>
  <div class="env">
    <div class="env__label">Instance</div>
    <DropdownMenuRoot>
      <DropdownMenuTrigger as-child>
        <button class="env__select" type="button">
          <span
            class="env__dot"
            :class="`env__dot--${active.status}`"
            aria-hidden="true"
          ></span>
          <span class="env__name">{{ active.name }}</span>
          <svg width="12" height="12" viewBox="0 0 12 12" fill="none" class="env__chevron" aria-hidden="true">
            <path d="M3 4.5l3 3 3-3" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuPortal>
        <DropdownMenuContent class="dd dd--env" :side-offset="8" align="start">
          <DropdownMenuLabel class="dd__label dd__label--env">Instances disponibles</DropdownMenuLabel>
          <DropdownMenuSeparator class="dd__sep" />
          <DropdownMenuItem
            v-for="i in instances"
            :key="i.id"
            class="dd__item dd__item--env"
            :data-active="i.id === active.id ? 'true' : undefined"
            @select="select(i)"
          >
            <span class="dd__env-dot" :class="`dd__env-dot--${i.status}`" aria-hidden="true"></span>
            <div class="dd__env-info">
              <span class="dd__env-name">{{ i.name }}</span>
              <span class="dd__env-endpoint mono">{{ i.endpoint }}</span>
            </div>
            <svg v-if="i.id === active.id" width="14" height="14" viewBox="0 0 14 14" fill="none" class="dd__env-check" aria-hidden="true">
              <path d="m3 7.5 2.5 2.5L11 4.5" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
          </DropdownMenuItem>
          <DropdownMenuSeparator class="dd__sep" />
          <DropdownMenuItem class="dd__item dd__item--env dd__item--add">
            <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
              <path d="M7 3v8M3 7h8" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" />
            </svg>
            Ajouter une instance
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenuPortal>
    </DropdownMenuRoot>
  </div>
</template>

<style scoped lang="scss">
.env {
  padding: 0 2rem 2rem;
  border-bottom: 1px solid $border-subtle;
  margin-bottom: 1.6rem;

  &__label {
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    margin-bottom: 0.6rem;
  }

  &__select {
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

  &__dot {
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

  &__name {
    flex: 1;
    text-align: left;
  }

  &__chevron {
    color: $text-muted;
  }
}
</style>

<style lang="scss">
// Styles globaux des items env (Reka téléporte le content hors du tree)
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
