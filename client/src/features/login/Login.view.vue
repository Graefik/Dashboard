<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
    <router-link to="/">Back to Home</router-link>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const email = ref("");
const password = ref("");

const handleSubmit = () => {
  console.log("Login attempt:", {
    email: email.value,
    password: password.value,
  });
};
</script>
