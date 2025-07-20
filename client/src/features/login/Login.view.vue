<script setup lang="ts">
import Wrapper from "@/shared/components/Wrapper.vue";
import Button from "@/shared/ui/Button.vue";
import Input from "@/shared/ui/Input.vue";

import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as zod from "zod";

const validationSchema = toTypedSchema(
  zod.object({
    email: zod
      .string({ message: "Email est requis" })
      .email({ message: "Doit être un email valide" }),
    password: zod
      .string({ message: "Mot de passe est requis" })
      .min(8, { message: "Trop court" }),
  })
);

const { handleSubmit, errors, defineField } = useForm({
  validationSchema,
  validateOnMount: false,
});

const [email, emailProps] = defineField("email", {
  validateOnModelUpdate: false,
});
const [password, passwordProps] = defineField("password", {
  validateOnModelUpdate: false,
});

const onSubmit = handleSubmit((values) => {
  alert(JSON.stringify(values, null, 2));
});
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
      <form @submit="onSubmit" novalidate>
        <div class="login__field">
          <label for="email" class="login__field-label">Email</label>
          <Input
            id="email"
            class="login__field-input"
            v-model="email"
            v-bind="emailProps"
          />
          <span class="login__field-error">{{ errors.email }}</span>
        </div>
        <div class="login__field">
          <label for="password" class="login__field-label">Mot de passe</label>
          <Input
            type="password"
            id="password"
            class="login__field-input"
            v-model="password"
            v-bind="passwordProps"
          />
          <span class="login__field-error">{{ errors.password }}</span>
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

    &-error {
      color: $error;
      font-size: 0.85rem;
      margin-top: 0.25rem;
    }
  }
}
</style>
