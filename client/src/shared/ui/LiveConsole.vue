<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from "vue";

type Tone = "ok" | "warn" | "err";
interface LogEntry {
  id: number;
  time: string;
  method: string;
  path: string;
  status: number;
  latency: number;
  tone: Tone;
}
type LogSample = Omit<LogEntry, "id" | "time">;

interface Props {
  title?: string;
  maxLines?: number;
  intervalMs?: number;
  height?: string;
  pool?: LogSample[];
}

const props = withDefaults(defineProps<Props>(), {
  title: "graefik.trace",
  maxLines: 9,
  intervalMs: 1400,
  height: "26rem",
});

// Pool par défaut : flux réaliste de requêtes Traefik
const defaultPool: LogSample[] = [
  { method: "GET", path: "/api/users", status: 200, latency: 12, tone: "ok" },
  { method: "POST", path: "/api/auth/login", status: 201, latency: 34, tone: "ok" },
  { method: "GET", path: "/api/billing/invoices", status: 200, latency: 89, tone: "ok" },
  { method: "GET", path: "/api/users/42", status: 404, latency: 8, tone: "warn" },
  { method: "PUT", path: "/api/users/42", status: 204, latency: 21, tone: "ok" },
  { method: "GET", path: "/static/app.js", status: 200, latency: 4, tone: "ok" },
  { method: "DELETE", path: "/api/sessions/c4d", status: 200, latency: 14, tone: "ok" },
  { method: "POST", path: "/api/webhooks/stripe", status: 500, latency: 320, tone: "err" },
  { method: "GET", path: "/api/search?q=traefik", status: 200, latency: 47, tone: "ok" },
  { method: "HEAD", path: "/health", status: 200, latency: 2, tone: "ok" },
  { method: "GET", path: "/api/services", status: 200, latency: 18, tone: "ok" },
  { method: "POST", path: "/api/events", status: 202, latency: 26, tone: "ok" },
  { method: "GET", path: "/api/metrics", status: 200, latency: 9, tone: "ok" },
  { method: "PATCH", path: "/api/config", status: 403, latency: 11, tone: "warn" },
  { method: "GET", path: "/api/routers", status: 200, latency: 15, tone: "ok" },
];

const logs = ref<LogEntry[]>([]);
let cursor = 0;
let timer: number | undefined;

const now = () =>
  new Date().toLocaleTimeString("fr-FR", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });

const pushLog = () => {
  const source = props.pool ?? defaultPool;
  const sample = source[Math.floor(Math.random() * source.length)];
  logs.value.push({ ...sample, id: cursor++, time: now() });
  if (logs.value.length > props.maxLines) logs.value.shift();
};

onMounted(() => {
  for (let i = 0; i < props.maxLines; i++) pushLog();
  timer = window.setInterval(pushLog, props.intervalMs);
});

onBeforeUnmount(() => {
  if (timer !== undefined) clearInterval(timer);
});
</script>

<template>
  <section class="console" aria-label="Flux de requêtes en direct">
    <header class="console__head">
      <span class="console__dots" aria-hidden="true">
        <span></span><span></span><span></span>
      </span>
      <span class="console__title mono">{{ title }}</span>
      <span class="console__live">
        <span class="console__live-dot" aria-hidden="true"></span>
        LIVE
      </span>
    </header>
    <div class="console__body" :style="{ height }" aria-live="off">
      <TransitionGroup name="logln">
        <div
          v-for="entry in logs"
          :key="entry.id"
          class="logln"
          :class="`logln--${entry.tone}`"
        >
          <span class="logln__ts mono">{{ entry.time }}</span>
          <span class="logln__method mono">{{ entry.method }}</span>
          <span class="logln__path mono">{{ entry.path }}</span>
          <span class="logln__status mono">{{ entry.status }}</span>
          <span class="logln__lat mono">{{ entry.latency }}ms</span>
        </div>
      </TransitionGroup>
    </div>
  </section>
</template>

<style scoped lang="scss">
.console {
  background: linear-gradient(180deg, $bg-surface, $bg-surface-2);
  border: 1px solid $border-default;
  border-radius: $radius-lg;
  box-shadow:
    0 24px 48px -24px rgba(0, 0, 0, 0.6),
    inset 0 1px 0 rgba(255, 255, 255, 0.04);
  overflow: hidden;
  font-family: $font-mono;

  &__head {
    display: flex;
    align-items: center;
    gap: 1.2rem;
    padding: 1rem 1.4rem;
    background: $bg-inset;
    border-bottom: 1px solid $border-subtle;
  }

  &__dots {
    display: inline-flex;
    gap: 0.5rem;

    span {
      width: 10px;
      height: 10px;
      border-radius: 50%;
      background: rgba(255, 255, 255, 0.08);
    }
  }

  &__title {
    flex: 1;
    text-align: center;
    font-size: 1.15rem;
    color: $text-muted;
    letter-spacing: 0.04em;
  }

  &__live {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.2rem 0.6rem;
    background: $severity-danger-soft;
    color: $severity-danger;
    border-radius: $radius-sm;
    font-size: 0.95rem;
    font-weight: 600;
    letter-spacing: 0.14em;
    text-transform: uppercase;
  }

  &__live-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-danger;
    box-shadow: 0 0 6px $severity-danger;
    animation: graefik-pulse 1.2s $ease-in-out infinite;
  }

  &__body {
    position: relative;
    padding: 0.6rem 0.4rem;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;

    // fondu haut : effet tail -f
    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      height: 5rem;
      background: linear-gradient(180deg, $bg-surface, transparent);
      pointer-events: none;
      z-index: 2;
    }
  }
}

.logln {
  display: grid;
  grid-template-columns: auto auto 1fr auto auto;
  align-items: center;
  gap: 1.1rem;
  padding: 0.35rem 1.2rem;
  font-size: 1.2rem;
  line-height: 1.5;
  white-space: nowrap;
  font-variant-numeric: tabular-nums;

  &__ts {
    color: $text-faint;
  }

  &__method {
    color: $text-muted;
    font-weight: 500;
    min-width: 5rem;
  }

  &__path {
    color: $text-primary;
    overflow: hidden;
    text-overflow: ellipsis;
    font-size: 1.2rem;
  }

  &__status {
    font-weight: 600;
    min-width: 3.5rem;
    text-align: right;
  }

  &__lat {
    color: $text-muted;
    min-width: 5rem;
    text-align: right;
  }

  &--ok &__status {
    color: $severity-success;
  }

  &--warn &__status {
    color: $severity-warning;
  }

  &--err &__status {
    color: $severity-danger;
  }

  &--err {
    background: rgba(241, 105, 105, 0.04);
  }
}

// Transitions TransitionGroup
.logln-enter-active {
  transition:
    transform 0.35s $ease-out,
    opacity 0.35s $ease-out;
}
.logln-leave-active {
  transition:
    transform 0.25s $ease-out,
    opacity 0.25s $ease-out;
  position: absolute;
  left: 0;
  right: 0;
}
.logln-enter-from {
  opacity: 0;
  transform: translateY(8px);
}
.logln-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
.logln-move {
  transition: transform 0.35s $ease-out;
}
</style>
