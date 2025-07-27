<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from '@/shared/layouts/MainLayout.vue'

const route = useRoute()
const isLoginPage = computed(() => route.name === 'login')
</script>

<template>
  <div id="app">
    <MainLayout v-if="!isLoginPage">
      <router-view />
    </MainLayout>
    <router-view v-else />
  </div>
</template>

<style></style>
