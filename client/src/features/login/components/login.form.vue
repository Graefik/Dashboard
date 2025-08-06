<script setup lang="ts">
import loginSchema from "@/features/login/schema/login.schema";
import Button from "@/shared/ui/Button.vue";
import Input from "@/shared/ui/Input.vue";
import { useForm } from "vee-validate";
import { ref } from "vue";

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

const isLoading = ref(false);
const loginError = ref("");

const onSubmit = handleSubmit(async (values) => {
  isLoading.value = true;
  loginError.value = "";
  
  try {
    // Simulation d'appel API - remplacer par votre logique d'authentification
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    // Ici vous devriez avoir votre logique de connexion réelle
    // Par exemple: await authService.login(values.email, values.password)
    
    alert(JSON.stringify(values, null, 2));
  } catch (error: any) {
    // Gestion des différents types d'erreurs
    if (error?.code === 'NETWORK_ERROR') {
      loginError.value = "Erreur de connexion. Vérifiez votre réseau.";
    } else if (error?.status === 401) {
      loginError.value = "Email ou mot de passe incorrect.";
    } else if (error?.status === 429) {
      loginError.value = "Trop de tentatives. Réessayez plus tard.";
    } else {
      loginError.value = "Une erreur est survenue. Réessayez plus tard.";
    }
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <form @submit="onSubmit" novalidate aria-label="Formulaire de connexion">
    <!-- Message d'erreur global -->
    <div v-if="loginError" class="login__error-message" role="alert" aria-live="polite">
      {{ loginError }}
    </div>

    <div class="login__field">
      <label for="email" class="login__field-label">Email</label>
      <Input
        id="email"
        class="login__field-input"
        v-model="email"
        v-bind="emailProps"
        autocomplete="username"
        :aria-describedby="errors.email ? 'email-error' : undefined"
        :aria-invalid="!!errors.email"
        :disabled="isLoading"
      />
      <span v-if="errors.email" id="email-error" class="login__field-error" role="alert">{{ errors.email }}</span>
    </div>

    <div class="login__field">
      <label for="password" class="login__field-label">Mot de passe</label>
      <Input
        type="password"
        id="password"
        class="login__field-input"
        v-model="password"
        v-bind="passwordProps"
        autocomplete="current-password"
        :aria-describedby="errors.password ? 'password-error' : undefined"
        :aria-invalid="!!errors.password"
        :disabled="isLoading"
      />
      <span v-if="errors.password" id="password-error" class="login__field-error" role="alert">{{ errors.password }}</span>
    </div>

    <div class="login__actions">
      <Button type="submit" :disabled="isLoading" class="login__submit-btn">
        <span v-if="!isLoading">Connexion</span>
        <span v-else aria-live="polite">Connexion en cours...</span>
      </Button>
    </div>
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
    }

    &-error {
      color: $error;
      font-size: 1.4rem;
      margin-top: 0.25rem;
    }
  }
}
</style>
