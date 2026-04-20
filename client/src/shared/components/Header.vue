<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import CommandPalette from "@/shared/components/CommandPalette.vue";
import TopbarClock from "@/shared/components/header/TopbarClock.vue";
import TopbarCrumbs from "@/shared/components/header/TopbarCrumbs.vue";
import TopbarRefresh from "@/shared/components/header/TopbarRefresh.vue";
import TopbarSearch from "@/shared/components/header/TopbarSearch.vue";
import TopbarTimeRanges from "@/shared/components/header/TopbarTimeRanges.vue";
import TopbarUserMenu from "@/shared/components/header/TopbarUserMenu.vue";

const router = useRouter();
const paletteOpen = ref(false);
</script>

<template>
  <header class="topbar">
    <div class="topbar__left">
      <TopbarCrumbs />
    </div>

    <TopbarSearch @open="paletteOpen = true" />

    <div class="topbar__right">
      <TopbarTimeRanges />
      <TopbarRefresh />
      <TopbarClock />
      <TopbarUserMenu />
    </div>

    <CommandPalette
      v-model:open="paletteOpen"
      @navigate="(to: string) => router.push(to)"
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

  &__right {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-left: auto;
  }
}
</style>
