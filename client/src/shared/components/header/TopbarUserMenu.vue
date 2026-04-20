<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "reka-ui";
import { useAuth } from "@/shared/auth/useAuth";

const router = useRouter();
const auth = useAuth();

const initials = computed(() => {
  const name = auth.username.value ?? "?";
  return name.slice(0, 2).toUpperCase();
});

const onLogout = () => {
  auth.logout();
  router.push({ name: "login" });
};
</script>

<template>
  <DropdownMenuRoot>
    <DropdownMenuTrigger as-child>
      <button class="user" type="button" aria-label="Compte">
        <div class="user__avatar" aria-hidden="true">{{ initials }}</div>
      </button>
    </DropdownMenuTrigger>
    <DropdownMenuPortal>
      <DropdownMenuContent class="dd" :side-offset="10" align="end">
        <DropdownMenuLabel class="dd__label">
          <div class="dd__user">
            <div class="dd__user-name">{{ auth.username.value ?? "Utilisateur" }}</div>
            <div class="dd__user-mail mono">connecté</div>
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
        <DropdownMenuItem class="dd__item dd__item--danger" @select="onLogout">
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
</template>

<style scoped lang="scss">
.user {
  background: transparent;
  border: 0;
  padding: 0;
  cursor: pointer;

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
    box-shadow:
      0 0 0 1px $border-accent,
      inset 0 1px 0 rgba(255, 255, 255, 0.18);
    transition: transform 0.18s $ease-spring;

    &:hover {
      transform: scale(1.05);
    }
  }
}
</style>
