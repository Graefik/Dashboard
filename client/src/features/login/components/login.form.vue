<script setup lang="ts">
import loginSchema from "@/features/login/schema/login.schema";
import { authService } from "@/features/login/service/auth.service";
import Button from "@/shared/ui/Button.vue";
import Input from "@/shared/ui/Input.vue";
import { useAuth } from "@/shared/auth/useAuth";
import type { ApiError } from "@/shared/api/http";
import { useForm } from "vee-validate";
import { ref } from "vue";
import { useRouter } from "vue-router";

const { handleSubmit, errors, defineField } = useForm({
  validationSchema: loginSchema,
  validateOnMount: false,
});

const [username, usernameProps] = defineField("username", {
  validateOnModelUpdate: false,
});
const [password, passwordProps] = defineField("password", {
  validateOnModelUpdate: false,
});

const isLoading = ref(false);
const loginError = ref("");
const router = useRouter();
const auth = useAuth();

const onSubmit = handleSubmit(async (values) => {
  isLoading.value = true;
  loginError.value = "";

  try {
    const res = await authService.login({
      username: values.username,
      password: values.password,
    });
    auth.login(res.token, res.username, res.expiresAt);
    await router.push("/");
  } catch (error: unknown) {
    const err = error as ApiError;
    if (err.code === "NETWORK_ERROR") {
      loginError.value = "Impossible de joindre le serveur.";
    } else if (err.code === "INVALID_CREDENTIALS" || err.status === 401) {
      loginError.value = "Identifiants incorrects.";
    } else if (err.status === 429) {
      loginError.value = "Trop de tentatives. Réessayez plus tard.";
    } else {
      loginError.value = err.message ?? "Une erreur est survenue.";
    }
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <form
    class="lf"
    @submit="onSubmit"
    novalidate
    aria-label="Formulaire de connexion"
  >
    <div v-if="loginError" class="lf__alert" role="alert" aria-live="polite">
      <svg width="16" height="16" viewBox="0 0 16 16" fill="none" aria-hidden="true">
        <circle cx="8" cy="8" r="6.5" stroke="currentColor" stroke-width="1.3" />
        <path d="M8 5v4" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" />
        <circle cx="8" cy="11.2" r="0.7" fill="currentColor" />
      </svg>
      <span>{{ loginError }}</span>
    </div>

    <div class="lf__field">
      <label for="username">Identifiant</label>
      <Input
        id="username"
        type="text"
        v-model="username"
        v-bind="usernameProps"
        autocomplete="username"
        placeholder="admin"
        :aria-describedby="errors.username ? 'username-error' : undefined"
        :aria-invalid="!!errors.username"
        :error="!!errors.username"
        :disabled="isLoading"
      />
      <span v-if="errors.username" id="username-error" class="lf__error" role="alert">
        {{ errors.username }}
      </span>
    </div>

    <div class="lf__field">
      <div class="lf__field-head">
        <label for="password">Mot de passe</label>
        <a href="#" class="lf__hint">Oublié&nbsp;?</a>
      </div>
      <Input
        id="password"
        type="password"
        v-model="password"
        v-bind="passwordProps"
        autocomplete="current-password"
        placeholder="••••••••••••"
        :aria-describedby="errors.password ? 'password-error' : undefined"
        :aria-invalid="!!errors.password"
        :error="!!errors.password"
        :disabled="isLoading"
      />
      <span
        v-if="errors.password"
        id="password-error"
        class="lf__error"
        role="alert"
      >
        {{ errors.password }}
      </span>
    </div>

    <Button
      type="submit"
      variant="primary"
      size="lg"
      block
      :disabled="isLoading"
      :loading="isLoading"
    >
      <template v-if="!isLoading">
        Se connecter
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path
            d="M2 7h10m0 0L8 3m4 4-4 4"
            stroke="currentColor"
            stroke-width="1.8"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </template>
      <template v-else>Connexion en cours…</template>
    </Button>
  </form>
</template>

<style scoped lang="scss">
.lf {
  display: flex;
  flex-direction: column;
  gap: 1.6rem;

  &__alert {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    padding: 1rem 1.2rem;
    color: $severity-danger;
    background: $severity-danger-soft;
    border: 1px solid rgba(241, 105, 105, 0.22);
    border-radius: $radius-md;
    font-size: 1.3rem;
  }

  &__field {
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
  }

  &__field-head {
    display: flex;
    align-items: center;
    justify-content: space-between;

    label {
      margin-bottom: 0;
    }
  }

  &__hint {
    font-size: 1.15rem;
    color: $accent;
    font-weight: 500;

    &:hover {
      color: $accent-glow;
      text-decoration: underline;
      text-underline-offset: 3px;
    }
  }

  &__error {
    color: $severity-danger;
    font-size: 1.2rem;
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
  }
}
</style>
