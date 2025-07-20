<script setup lang="ts">
import Wrapper from "@/shared/components/Wrapper.vue";
import LoginForm from "@/features/login/components/login.form.vue";
</script>

<template>
  <Wrapper class="login">
    <div class="login__header">
      <img
        src="@/assets/logo.png"
        alt="GraefikLogo"
        class="login__header-logo"
      />
    </div>
    <div class="login__form">
      <h1 class="login__title">Connexion</h1>
      <LoginForm />
    </div>
  </Wrapper>
</template>

<style scoped lang="scss">
.login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;

  &__header {
    text-align: center;
    margin-bottom: 3rem;

    &-logo {
      width: 20rem;
      height: auto;
    }
  }

  &__form {
    width: 100%;
    max-width: 60rem;
    background-color: $primary-fg;
    border-radius: 3rem;
    padding: 3rem 6rem 5rem;
    margin: 0 auto;
    text-align: center;
    box-shadow: 0 2rem 4rem rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(255, 255, 255, 0.1);

    @media (max-width: 768px) {
      padding: 2rem 2rem 3rem;
      margin: 0 1rem;
    }
  }
}
</style>
