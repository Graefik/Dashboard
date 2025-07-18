<template>
  <div class="main-layout">
    <header class="main-layout__header">Header</header>
    <main class="main-layout__main">
      <slot />
    </main>
    <footer class="main-layout__footer">Footer</footer>
  </div>
</template>

<style lang="scss" scoped>
.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;

  &__header {
    background-color: $primary-fg;
    padding: 1rem;
    color: $white;
  }

  &__main {
    flex: 1;
    padding: 2rem;
  }

  &__footer {
    background-color: $primary-fg;
    padding: 1rem;
    color: $white;
    text-align: center;
  }
}
</style>