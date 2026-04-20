<script setup lang="ts">
import DashboardHead from "@/features/dashboard/components/DashboardHead.vue";
import KpiRow from "@/features/dashboard/components/KpiRow.vue";
import LatencyPanel from "@/features/dashboard/components/LatencyPanel.vue";
import ServicesTablePanel from "@/features/dashboard/components/ServicesTablePanel.vue";
import StatusCodesPanel from "@/features/dashboard/components/StatusCodesPanel.vue";
import TimelinePanel from "@/features/dashboard/components/TimelinePanel.vue";
import TopRoutersPanel from "@/features/dashboard/components/TopRoutersPanel.vue";
</script>

<template>
  <div class="dashboard">
    <DashboardHead />
    <KpiRow />
    <section class="dashboard__grid">
      <TimelinePanel />
      <StatusCodesPanel />
      <LatencyPanel />
      <TopRoutersPanel />
      <ServicesTablePanel />
    </section>
  </div>
</template>

<style scoped lang="scss">
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 2.4rem;

  &__grid {
    display: grid;
    grid-template-columns: repeat(12, minmax(0, 1fr));
    gap: 1.6rem;
  }
}
</style>
