<script setup lang="ts">
import Footer from "@/shared/components/Footer.vue";
import Header from "@/shared/components/Header.vue";
import Wrapper from "@/shared/components/Wrapper.vue";
</script>

<template>
  <div class="main-layout">
    <Header />
    <main class="main-layout__main">
      <Wrapper>
        <slot />
      </Wrapper>
    </main>
    <Footer />
  </div>
</template>

<style lang="scss" scoped>
.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;

  &__main {
    flex: 1;
    background-color: $primary-bg;
    padding: 8rem 0;
  }
}
</style>
