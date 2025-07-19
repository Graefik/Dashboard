<script setup lang="ts">
import Wrapper from "@/shared/components/Wrapper.vue";
import Button from "@/shared/ui/Button.vue";
import Input from "@/shared/ui/Input.vue";
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
          <Input type="email" id="email" class="login__field-input" required />
        </div>
        <div class="login__field">
          <label for="password" class="login__field-label">Mot de passe</label>
          <Input
            type="password"
            id="password"
            class="login__field-input"
            required
          />
        </div>
        <Button type="submit">Connexion</Button>
      </form>
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
      width: 200px;
      height: auto;
    }
  }

  &__form {
    width: 100%;
    max-width: 600px;
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
  &__field {
    margin-bottom: 2rem;
    text-align: left;

    &-label {
      font-weight: 600;
      margin-bottom: 0.5rem;
      display: block;
      color: $white;
      font-size: 0.95rem;
    }
  }
}
</style>
