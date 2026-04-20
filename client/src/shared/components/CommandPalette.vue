<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from "vue";
import {
  DialogContent,
  DialogOverlay,
  DialogPortal,
  DialogRoot,
  DialogTitle,
} from "reka-ui";

interface Props {
  open: boolean;
}
const props = defineProps<Props>();
const emit = defineEmits<{ "update:open": [v: boolean]; navigate: [to: string] }>();

interface Cmd {
  id: string;
  group: "Navigation" | "Infrastructure" | "Actions" | "Settings";
  label: string;
  hint?: string;
  shortcut?: string;
  to?: string;
}

const commands: Cmd[] = [
  { id: "nav-overview", group: "Navigation", label: "Overview", hint: "Dashboard principal", to: "/" },
  { id: "nav-logs", group: "Navigation", label: "Logs", hint: "Stream temps réel", to: "/logs" },
  { id: "nav-alerts", group: "Navigation", label: "Alerts", hint: "Alertes et règles", to: "/alerts" },
  { id: "nav-settings", group: "Navigation", label: "Settings", hint: "Configuration", to: "/settings" },

  { id: "infra-routers", group: "Infrastructure", label: "Routers" },
  { id: "infra-services", group: "Infrastructure", label: "Services" },
  { id: "infra-middlewares", group: "Infrastructure", label: "Middlewares" },
  { id: "infra-certs", group: "Infrastructure", label: "Certificates" },

  { id: "act-refresh", group: "Actions", label: "Actualiser les panneaux", shortcut: "R" },
  { id: "act-export", group: "Actions", label: "Exporter le dashboard", shortcut: "⌘ E" },
  { id: "act-panel", group: "Actions", label: "Ajouter un panneau", shortcut: "⌘ N" },
  { id: "act-theme", group: "Actions", label: "Basculer le thème" },

  { id: "set-switch-instance", group: "Settings", label: "Changer d'instance" },
  { id: "set-shortcuts", group: "Settings", label: "Raccourcis clavier", shortcut: "?" },
];

const query = ref("");
const activeIdx = ref(0);
const inputEl = ref<HTMLInputElement>();

const filtered = computed(() => {
  const q = query.value.trim().toLowerCase();
  if (!q) return commands;
  return commands.filter(
    (c) => c.label.toLowerCase().includes(q) || c.hint?.toLowerCase().includes(q),
  );
});

const grouped = computed(() => {
  const m = new Map<string, Cmd[]>();
  for (const c of filtered.value) {
    if (!m.has(c.group)) m.set(c.group, []);
    m.get(c.group)!.push(c);
  }
  return Array.from(m.entries());
});

const flatList = computed(() => filtered.value);

watch(query, () => (activeIdx.value = 0));
watch(
  () => props.open,
  (v) => {
    if (v) {
      query.value = "";
      activeIdx.value = 0;
      // focus géré par autofocus sur l'input à l'ouverture
      requestAnimationFrame(() => inputEl.value?.focus());
    }
  },
);

const close = () => emit("update:open", false);

const run = (cmd: Cmd) => {
  if (cmd.to) emit("navigate", cmd.to);
  close();
};

const onKey = (e: KeyboardEvent) => {
  const n = flatList.value.length;
  if (!n) return;
  if (e.key === "ArrowDown") {
    e.preventDefault();
    activeIdx.value = (activeIdx.value + 1) % n;
  } else if (e.key === "ArrowUp") {
    e.preventDefault();
    activeIdx.value = (activeIdx.value - 1 + n) % n;
  } else if (e.key === "Enter") {
    e.preventDefault();
    const cmd = flatList.value[activeIdx.value];
    if (cmd) run(cmd);
  }
};

// Raccourci global ⌘K / Ctrl+K
const onGlobalKey = (e: KeyboardEvent) => {
  if ((e.metaKey || e.ctrlKey) && e.key.toLowerCase() === "k") {
    e.preventDefault();
    emit("update:open", !props.open);
  }
};

onMounted(() => window.addEventListener("keydown", onGlobalKey));
onBeforeUnmount(() => window.removeEventListener("keydown", onGlobalKey));

// Index global au sein du filtre — sert à distinguer le surlignage
const flatIndex = (cmd: Cmd) => flatList.value.indexOf(cmd);
</script>

<template>
  <DialogRoot :open="open" @update:open="emit('update:open', $event)">
    <DialogPortal>
      <DialogOverlay class="cmdk__overlay" />
      <DialogContent class="cmdk__content" @keydown="onKey">
        <DialogTitle class="cmdk__sr">Command palette</DialogTitle>
        <div class="cmdk__search">
          <svg
            width="16"
            height="16"
            viewBox="0 0 16 16"
            fill="none"
            aria-hidden="true"
          >
            <circle cx="7" cy="7" r="4.5" stroke="currentColor" stroke-width="1.5" />
            <path
              d="m11 11 3.5 3.5"
              stroke="currentColor"
              stroke-width="1.5"
              stroke-linecap="round"
            />
          </svg>
          <input
            ref="inputEl"
            v-model="query"
            type="text"
            placeholder="Chercher une commande, une route, un service…"
            autocomplete="off"
            spellcheck="false"
          />
          <kbd>esc</kbd>
        </div>

        <div class="cmdk__list" role="listbox">
          <div v-if="!flatList.length" class="cmdk__empty">
            <svg width="28" height="28" viewBox="0 0 28 28" fill="none" aria-hidden="true">
              <circle cx="12" cy="12" r="8" stroke="currentColor" stroke-width="1.3" />
              <path
                d="m18 18 5 5"
                stroke="currentColor"
                stroke-width="1.3"
                stroke-linecap="round"
              />
            </svg>
            <p>Aucun résultat pour <span class="mono">«&nbsp;{{ query }}&nbsp;»</span></p>
          </div>

          <div v-for="[group, items] in grouped" :key="group" class="cmdk__group">
            <div class="cmdk__group-title">{{ group }}</div>
            <button
              v-for="cmd in items"
              :key="cmd.id"
              type="button"
              class="cmdk__item"
              :class="{ 'cmdk__item--active': flatIndex(cmd) === activeIdx }"
              @mouseenter="activeIdx = flatIndex(cmd)"
              @click="run(cmd)"
              role="option"
              :aria-selected="flatIndex(cmd) === activeIdx"
            >
              <span class="cmdk__item-label">{{ cmd.label }}</span>
              <span v-if="cmd.hint" class="cmdk__item-hint">{{ cmd.hint }}</span>
              <kbd v-if="cmd.shortcut" class="cmdk__item-kbd">{{ cmd.shortcut }}</kbd>
            </button>
          </div>
        </div>

        <footer class="cmdk__foot">
          <div class="cmdk__legend">
            <span><kbd>↑</kbd><kbd>↓</kbd> naviguer</span>
            <span><kbd>↵</kbd> exécuter</span>
            <span><kbd>esc</kbd> fermer</span>
          </div>
          <span class="cmdk__branding mono">Graefik · ⌘K</span>
        </footer>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<style lang="scss">
.cmdk {
  &__sr {
    position: absolute;
    width: 1px;
    height: 1px;
    overflow: hidden;
    clip: rect(0 0 0 0);
    white-space: nowrap;
  }

  &__overlay {
    position: fixed;
    inset: 0;
    background: rgba(4, 14, 25, 0.72);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    z-index: 200;
    animation: cmdk-fade 0.16s $ease-out;
  }

  &__content {
    position: fixed;
    z-index: 201;
    top: 18vh;
    left: 50%;
    transform: translateX(-50%);
    width: min(64rem, calc(100vw - 3rem));
    max-height: 64vh;
    display: flex;
    flex-direction: column;
    background: $bg-surface;
    border: 1px solid $border-default;
    border-radius: $radius-lg;
    box-shadow: $shadow-lg, $shadow-inset;
    overflow: hidden;
    animation: cmdk-in 0.2s $ease-spring;
  }

  &__search {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 1.2rem 1.4rem;
    border-bottom: 1px solid $border-subtle;
    color: $text-muted;

    input {
      flex: 1;
      background: transparent;
      border: 0;
      color: $text-primary;
      font-size: 1.4rem;
      font-family: $font-sans;
      outline: none;

      &::placeholder {
        color: $text-faint;
      }
    }

    kbd {
      background: $bg-inset;
      border: 1px solid $border-default;
      color: $text-muted;
      padding: 0.2rem 0.55rem;
      font-size: 1rem;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
  }

  &__list {
    flex: 1;
    overflow-y: auto;
    padding: 0.6rem 0.6rem 0.8rem;
  }

  &__empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.8rem;
    padding: 4rem 2rem;
    color: $text-muted;
    font-size: 1.3rem;
    text-align: center;

    p {
      margin: 0;
    }
  }

  &__group {
    & + & {
      margin-top: 0.6rem;
      padding-top: 0.6rem;
      border-top: 1px solid $border-subtle;
    }
  }

  &__group-title {
    padding: 0.6rem 0.8rem 0.4rem;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    font-weight: 500;
  }

  &__item {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    gap: 1rem;
    width: 100%;
    padding: 0.8rem 1rem;
    background: transparent;
    border: 0;
    border-radius: $radius-sm;
    color: $text-secondary;
    text-align: left;
    cursor: pointer;
    font-family: $font-sans;
    font-size: 1.35rem;
    transition: background 0.1s $ease-out, color 0.1s $ease-out;

    &--active {
      background: $accent-soft;
      color: $text-primary;

      .cmdk__item-hint {
        color: $text-secondary;
      }

      .cmdk__item-kbd {
        color: $accent-glow;
        border-color: $border-accent;
      }
    }
  }

  &__item-label {
    color: inherit;
    font-weight: 500;
  }

  &__item-hint {
    color: $text-muted;
    font-size: 1.15rem;
    text-align: right;
    justify-self: end;
  }

  &__item-kbd {
    font-family: $font-mono;
    font-size: 1rem;
    padding: 0.15rem 0.5rem;
    background: $bg-inset;
    border: 1px solid $border-subtle;
    color: $text-muted;
    border-radius: $radius-xs;
  }

  &__foot {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 1.6rem;
    padding: 0.8rem 1.2rem;
    background: $bg-inset;
    border-top: 1px solid $border-subtle;
    font-family: $font-mono;
    font-size: 1.05rem;
    color: $text-muted;
  }

  &__legend {
    display: inline-flex;
    align-items: center;
    gap: 1.2rem;

    span {
      display: inline-flex;
      align-items: center;
      gap: 0.4rem;
    }

    kbd {
      padding: 0.15rem 0.45rem;
      font-size: 0.95rem;
      background: $bg-surface;
      border: 1px solid $border-subtle;
      color: $text-secondary;
      border-radius: $radius-xs;
    }
  }

  &__branding {
    color: $text-faint;
    letter-spacing: 0.08em;
  }
}

@keyframes cmdk-fade {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes cmdk-in {
  from {
    opacity: 0;
    transform: translate(-50%, -8px) scale(0.98);
  }
  to {
    opacity: 1;
    transform: translate(-50%, 0) scale(1);
  }
}
</style>
