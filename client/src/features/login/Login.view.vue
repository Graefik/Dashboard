<script setup lang="ts">
import LoginForm from "@/features/login/components/login.form.vue";
import Wrapper from "@/shared/components/Wrapper.vue";
</script>

<template>
  <Wrapper class="login">
    <div class="login__header">
      <img
        src="@/assets/logo.png"
        alt="GraefikLogo"
        class="login__header-logo"
      />
      <p class="login__header-tagline">Votre tableau de bord analytique</p>
    </div>
    <div class="login__form">
      <h1 class="login__title">Bienvenue</h1>
      <p class="login__subtitle">
        Connectez-vous pour accéder à vos graphiques
      </p>
      <LoginForm />
      <div class="login__forgot-password">
        <a target="_blank" rel="noopener noreferrer" class="login__forgot-link">
          Mot de passe oublié ?
        </a>
      </div>
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
      margin-bottom: 1rem;
    }

    &-tagline {
      color: rgba($white, 0.8);
      font-size: 1.6rem;
      font-weight: 300;
      margin: 0;
      letter-spacing: 0.5px;
    }
  }

  &__form {
    width: 100%;
    max-width: 53rem;
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

  &__title {
    margin-bottom: 0.5rem;
    color: $white;
    font-size: 2.8rem;
    font-weight: 600;
  }

  &__subtitle {
    margin: 0 0 3rem 0;
    color: rgba($white, 0.7);
    font-size: 1.6rem;
    font-weight: 400;
  }

  &__forgot-password {
    margin-top: 2rem;
    text-align: center;
  }

  &__forgot-link {
    color: $cta;
    text-decoration: none;
    font-size: 1.4rem;
    transition: color 0.2s ease;

    &:hover {
      color: lighten($cta, 10%);
      text-decoration: underline;
    }
  }
}
</style>
