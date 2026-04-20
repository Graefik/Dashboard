<script setup lang="ts">
import SidebarBrand from "@/shared/components/sidebar/SidebarBrand.vue";
import SidebarEnvSelect from "@/shared/components/sidebar/SidebarEnvSelect.vue";
import SidebarNav from "@/shared/components/sidebar/SidebarNav.vue";
import SidebarUptime from "@/shared/components/sidebar/SidebarUptime.vue";
</script>

<template>
  <aside class="rail">
    <SidebarBrand />
    <SidebarEnvSelect />
    <SidebarNav />
    <SidebarUptime />
  </aside>
</template>

<style scoped lang="scss">
.rail {
  position: sticky;
  top: 0;
  width: $rail-width;
  height: 100vh;
  flex-shrink: 0;
  background: linear-gradient(180deg, $bg-rail 0%, $bg-rail-2 100%);
  border-right: 1px solid $border-subtle;
  display: flex;
  flex-direction: column;
  padding: 2rem 0 0;
  overflow-y: auto;
  overflow-x: hidden;

  > * {
    flex-shrink: 0;
  }

  > :nth-child(3) {
    flex: 1;
  }
}
</style>
