<script setup lang="ts">
import Wrapper from "@/shared/components/Wrapper.vue";
import { computed } from "vue";

const currentYear = computed(() => new Date().getFullYear());
</script>

<template>
  <footer class="footer">
    <Wrapper>
      <div class="footer-content">
        <p>
          &copy; {{ currentYear }} Graefik - Votre tableau de bord analytique
        </p>
      </div>
    </Wrapper>
  </footer>
</template>

<style scoped lang="scss">
.footer {
  background-color: $primary-fg;
  border-top: 1px solid rgba($white, 0.1);
  padding: 2rem 0;
}

.footer-content {
  text-align: center;
  padding: 2rem 0;

  p {
    color: rgba($white, 0.6);
    font-size: 1.4rem;
    margin: 0;
  }
}
</style>
