<script setup lang="ts">
import Footer from "@/shared/components/Footer.vue";
import Header from "@/shared/components/Header.vue";
import Sidebar from "@/shared/components/Sidebar.vue";
</script>

<template>
  <div class="shell">
    <Sidebar />
    <div class="shell__body">
      <Header />
      <main class="shell__main">
        <slot />
      </main>
      <Footer />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.shell {
  display: flex;
  min-height: 100vh;
  width: 100%;

  &__body {
    flex: 1;
    min-width: 0;
    display: flex;
    flex-direction: column;
  }

  &__main {
    flex: 1;
    padding: 3rem 3rem 4rem;
    animation: graefik-fade-up 0.4s $ease-out both;
  }

  @media (max-width: 900px) {
    &__main {
      padding: 2rem 1.6rem 3rem;
    }
  }
}
</style>
