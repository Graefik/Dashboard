<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const breadcrumb = computed(() => {
  const name = (route.name as string) || "overview";
  const map: Record<string, string[]> = {
    home: ["Observatory", "Overview"],
    overview: ["Observatory", "Overview"],
  };
  return map[name] || ["Observatory", name];
});
</script>

<template>
  <div class="crumbs">
    <template v-for="(crumb, i) in breadcrumb" :key="i">
      <span v-if="i > 0" class="crumbs__sep" aria-hidden="true">/</span>
      <span
        class="crumbs__item"
        :class="{ 'crumbs__item--current': i === breadcrumb.length - 1 }"
      >
        {{ crumb }}
      </span>
    </template>
  </div>
</template>

<style scoped lang="scss">
.crumbs {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  font-size: 1.35rem;

  &__item {
    color: $text-muted;
    font-weight: 500;

    &--current {
      color: $text-primary;
      font-weight: 600;
    }
  }

  &__sep {
    color: $text-faint;
  }
}
</style>
