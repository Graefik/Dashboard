<script setup lang="ts">
import MainLayout from "@/shared/layouts/MainLayout.vue";
</script>

<template>
  <MainLayout>
    <div>
      <h1>Dashboard</h1>
      <p>Welcome to the dashboard!</p>
      <router-link to="/login">Go to Login</router-link>
    </div>
  </MainLayout>
</template>
