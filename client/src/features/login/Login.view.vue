<script setup lang="ts">
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
    </div>
    <div class="login__form">
      <h1 class="login__title">Connexion</h1>
      <form>
        <div class="login__field">
          <label for="email" class="login__field-label">Email</label>
          <input type="email" id="email" class="login__field-input" required />
        </div>
        <div class="login__field">
          <label for="password" class="login__field-label">Mot de passe</label>
          <input
            type="password"
            id="password"
            class="login__field-input"
            required
          />
        </div>
        <button class="login__form-button" type="submit">Connexion</button>
      </form>
    </div>
  </Wrapper>
</template>

<style scoped lang="scss">
.login {
  &__header {
    text-align: center;
    margin-bottom: 3rem;

    &-logo {
      width: 200px;
      height: auto;
    }
  }

  &__form {
    max-width: 650px;
    background-color: $primary-fg;
    border-radius: 30px;
    padding: 3rem 6rem 5rem;
    margin: 0 auto;
    text-align: center;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(255, 255, 255, 0.1);

    @media (max-width: 768px) {
      padding: 2rem 2rem 3rem;
      margin: 0 1rem;
    }
  }

  &__title {
    margin-bottom: 2rem;
    color: $white;
    font-size: 2rem;
    font-weight: 600;
  }

  &__field {
    margin-bottom: 2rem;
    text-align: left;
  }

  &__field-label {
    font-weight: 600;
    margin-bottom: 0.5rem;
    display: block;
    color: rgba($white, 0.9);
    font-size: 0.95rem;
  }

  &__field-input {
    width: 100%;
    padding: 14px;
    border-radius: 12px;
    border: 2px solid rgba(255, 255, 255, 0.1);
    font-size: 1rem;
    background-color: rgba(255, 255, 255, 0.05);
    color: $white;
    transition: all 0.3s ease;

    &:focus {
      border-color: $cta;
      outline: none;
      background-color: rgba(255, 255, 255, 0.1);
      box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.1);
    }
  }

  &__form-button {
    background-color: $cta;
    border-radius: 12px;
    padding: 14px 32px;
    font-size: 1.1rem;
    font-weight: 600;
    border: none;
    width: 100%;
    color: $white;
    cursor: pointer;
    transition: all 0.3s ease;
    margin-top: 1rem;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);

    &:hover {
      background-color: darken($cta, 5%);
      transform: translateY(-2px);
      box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
    }

    &:active {
      transform: translateY(0);
    }
  }
}
</style>
