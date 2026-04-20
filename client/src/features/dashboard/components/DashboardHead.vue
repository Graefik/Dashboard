<script setup lang="ts">
import Button from "@/shared/ui/Button.vue";
</script>

<template>
  <header class="page-head">
    <div class="page-head__titles">
      <div class="eyebrow">
        <span class="page-head__live">
          <span class="page-head__live-dot" aria-hidden="true"></span>
          LIVE
        </span>
        <span>Observatory · Overview</span>
      </div>
      <h1 class="page-head__title">Traefik Telemetry</h1>
      <p class="page-head__subtitle">
        Vue d'ensemble en temps réel de vos routers, services et middlewares.
        <span class="mono page-head__last"
          >Mis à jour {{ new Date().toLocaleTimeString("fr-FR") }}</span
        >
      </p>
    </div>

    <div class="page-head__actions">
      <Button variant="subtle" size="sm">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path
            d="M7 1v8m0 0L4 6m3 3 3-3"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path d="M2 11h10" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
        </svg>
        Exporter
      </Button>
      <Button variant="ghost" size="sm">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path
            d="M12 7A5 5 0 1 1 7 2"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M9 1 12 2l-1 3"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        Actualiser
      </Button>
      <Button size="sm">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path d="M7 3v8M3 7h8" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
        </svg>
        Nouveau panneau
      </Button>
    </div>
  </header>
</template>

<style scoped lang="scss">
.page-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 2rem;
  flex-wrap: wrap;

  &__titles {
    min-width: 0;
  }

  &__live {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.2rem 0.55rem;
    background: $severity-danger-soft;
    color: $severity-danger;
    border-radius: $radius-sm;
    font-size: 0.95rem;
    font-weight: 600;
    letter-spacing: 0.15em;
    margin-right: 0.4rem;
  }

  &__live-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-danger;
    box-shadow: 0 0 6px $severity-danger;
    animation: graefik-pulse 1.2s $ease-in-out infinite;
  }

  &__title {
    margin-top: 0.8rem;
    font-size: 2.6rem;
    font-weight: 600;
    letter-spacing: -0.025em;
    line-height: 1.15;
    color: $text-primary;
  }

  &__subtitle {
    margin-top: 0.6rem;
    color: $text-muted;
    font-size: 1.4rem;
    display: flex;
    align-items: center;
    gap: 1rem;
    flex-wrap: wrap;
  }

  &__last {
    font-size: 1.15rem;
    color: $text-faint;
    padding: 0.2rem 0.6rem;
    border-radius: $radius-sm;
    background: $bg-inset;
    border: 1px solid $border-subtle;
  }

  &__actions {
    display: inline-flex;
    align-items: center;
    gap: 0.8rem;
  }
}
</style>
