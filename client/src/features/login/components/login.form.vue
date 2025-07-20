<script setup lang="ts">
import loginSchema from "@/features/login/schema/login.schema";
import Button from "@/shared/ui/Button.vue";
import Input from "@/shared/ui/Input.vue";
import { useForm } from "vee-validate";

const { handleSubmit, errors, defineField } = useForm({
  validationSchema: loginSchema,
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
</template>

<style scoped lang="scss">
.login {
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