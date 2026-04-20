<script setup lang="ts">
import LoginForm from "@/features/login/components/login.form.vue";
</script>

<template>
  <div class="auth">
    <!-- Colonne gauche : brand + télémétrie mockée ───────────── -->
    <aside class="auth__aside" aria-hidden="true">
      <div class="auth__brand">
        <img
          src="@/assets/logo.png"
          alt="Graefik"
          class="auth__logo"
        />
        <span class="auth__brand-meta mono">Traefik Observatory · on-prem</span>
      </div>

      <div class="auth__hero">
        <div class="eyebrow">
          <span class="auth__live">
            <span class="auth__live-dot"></span>
            OBSERVATORY ONLINE
          </span>
        </div>
        <h1 class="auth__headline">
          Votre plateforme Traefik,<br />
          <span class="auth__headline-accent">en vision directe.</span>
        </h1>
        <p class="auth__lede">
          Métriques, routers, services et certificats — tout votre reverse proxy
          en temps réel, sur vos propres infrastructures.
        </p>
      </div>

      <div class="auth__telemetry">
        <div class="auth__tile">
          <div class="auth__tile-label">Requests / sec</div>
          <div class="auth__tile-value mono">1 284</div>
          <svg class="auth__tile-spark" viewBox="0 0 120 28" preserveAspectRatio="none">
            <defs>
              <linearGradient id="spark1" x1="0" y1="0" x2="0" y2="1">
                <stop offset="0" stop-color="#00c2b2" stop-opacity="0.6" />
                <stop offset="1" stop-color="#00c2b2" stop-opacity="0" />
              </linearGradient>
            </defs>
            <path
              d="M0,20 L10,18 L20,22 L30,14 L40,16 L50,10 L60,12 L70,6 L80,10 L90,4 L100,8 L110,5 L120,9 L120,28 L0,28 Z"
              fill="url(#spark1)"
            />
            <path
              d="M0,20 L10,18 L20,22 L30,14 L40,16 L50,10 L60,12 L70,6 L80,10 L90,4 L100,8 L110,5 L120,9"
              fill="none"
              stroke="#2ce8d8"
              stroke-width="1.6"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
        <div class="auth__tile">
          <div class="auth__tile-label">Latency p95</div>
          <div class="auth__tile-value mono">42<span class="auth__tile-unit">ms</span></div>
          <svg class="auth__tile-spark" viewBox="0 0 120 28" preserveAspectRatio="none">
            <path
              d="M0,14 L10,12 L20,15 L30,13 L40,10 L50,14 L60,9 L70,11 L80,8 L90,12 L100,7 L110,10 L120,6"
              fill="none"
              stroke="#58a6ff"
              stroke-width="1.6"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
        <div class="auth__tile">
          <div class="auth__tile-label">Services</div>
          <div class="auth__tile-value mono">24<span class="auth__tile-unit">/24</span></div>
          <div class="auth__tile-bars">
            <span v-for="n in 24" :key="n" :style="`--h: ${40 + Math.sin(n) * 30 + Math.random() * 30}%`"></span>
          </div>
        </div>
      </div>
    </aside>

    <!-- Colonne droite : form ────────────────────────────────── -->
    <section class="auth__main">
      <div class="auth__form">
        <div class="auth__head">
          <div class="eyebrow">AUTHENTICATION · on-premise</div>
          <h2 class="auth__title">Bon retour</h2>
          <p class="auth__sub">
            Identifiez-vous pour accéder à votre console Graefik.
          </p>
        </div>

        <LoginForm />

        <div class="auth__foot">
          <a href="#" class="auth__link">Mot de passe oublié&nbsp;?</a>
          <span class="auth__foot-sep">·</span>
          <a href="#" class="auth__link">Mode SSO</a>
        </div>

        <div class="auth__legal mono">
          <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>L</kbd> pour se connecter
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
.auth {
  display: grid;
  grid-template-columns: minmax(0, 1.1fr) minmax(0, 1fr);
  min-height: 100vh;

  @media (max-width: 960px) {
    grid-template-columns: 1fr;
  }

  &__aside {
    position: relative;
    padding: 4rem 4rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    background:
      linear-gradient(135deg, $bg-rail 0%, $bg-base 100%);
    border-right: 1px solid $border-subtle;
    overflow: hidden;

    &::before {
      content: "";
      position: absolute;
      inset: 0;
      background:
        radial-gradient(circle at 20% 20%, rgba(0, 194, 178, 0.12) 0%, transparent 45%),
        radial-gradient(circle at 85% 70%, rgba(88, 166, 255, 0.08) 0%, transparent 45%);
      pointer-events: none;
    }

    &::after {
      content: "";
      position: absolute;
      inset: 0;
      background-image: radial-gradient(
        circle at 1px 1px,
        rgba(255, 255, 255, 0.04) 1px,
        transparent 0
      );
      background-size: 2.4rem 2.4rem;
      pointer-events: none;
      mask-image: linear-gradient(180deg, rgba(0, 0, 0, 0.8) 0%, transparent 100%);
    }

    > * {
      position: relative;
      z-index: 1;
    }

    @media (max-width: 960px) {
      display: none;
    }
  }

  &__brand {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  &__logo {
    width: 11rem;
    height: auto;
    border-radius: $radius-md;
    box-shadow: 0 10px 28px -12px rgba(0, 0, 0, 0.55);
  }

  &__brand-meta {
    font-size: 1.1rem;
    color: $text-muted;
    letter-spacing: 0.08em;
    text-transform: uppercase;
  }

  &__hero {
    max-width: 44rem;
  }

  &__live {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    color: $accent-glow;
  }

  &__live-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: $severity-success;
    box-shadow: 0 0 8px $severity-success;
    animation: graefik-pulse 1.6s $ease-in-out infinite;
  }

  &__headline {
    margin-top: 1.2rem;
    font-size: 3.4rem;
    font-weight: 500;
    line-height: 1.1;
    letter-spacing: -0.025em;
    color: $text-primary;

    &-accent {
      color: $text-secondary;
      font-weight: 400;
    }
  }

  &__lede {
    margin-top: 1.6rem;
    font-size: 1.55rem;
    line-height: 1.55;
    color: $text-secondary;
    max-width: 38rem;
  }

  &__telemetry {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.2rem;
    margin-top: auto;
  }

  &__tile {
    position: relative;
    padding: 1.4rem 1.4rem 1rem;
    background: linear-gradient(180deg, $bg-surface, $bg-surface-2);
    border: 1px solid $border-subtle;
    border-radius: $radius-md;
    overflow: hidden;

    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      height: 1px;
      background: linear-gradient(90deg, transparent, $accent, transparent);
      opacity: 0.4;
    }
  }

  &__tile-label {
    font-family: $font-mono;
    font-size: 0.95rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    margin-bottom: 0.6rem;
  }

  &__tile-value {
    font-size: 2rem;
    font-weight: 600;
    color: $text-primary;
    font-variant-numeric: tabular-nums;
    letter-spacing: -0.02em;
  }

  &__tile-unit {
    font-size: 1.2rem;
    color: $text-muted;
    margin-left: 0.2rem;
    font-weight: 500;
  }

  &__tile-spark {
    display: block;
    width: 100%;
    height: 2.4rem;
    margin-top: 0.6rem;
  }

  &__tile-bars {
    display: flex;
    align-items: flex-end;
    gap: 2px;
    height: 2.4rem;
    margin-top: 0.6rem;

    span {
      flex: 1;
      height: var(--h);
      min-height: 2px;
      background: linear-gradient(180deg, $severity-success 0%, $accent-deep 100%);
      border-radius: 1px;
      opacity: 0.75;
    }
  }

  // ── Main form column ─────────────────────────────
  &__main {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4rem 3rem;
    background: $bg-base;
    position: relative;

    &::before {
      content: "";
      position: absolute;
      inset: 0;
      background-image: radial-gradient(
        circle at 1px 1px,
        rgba(255, 255, 255, 0.03) 1px,
        transparent 0
      );
      background-size: 2.4rem 2.4rem;
      pointer-events: none;
      opacity: 0.6;
    }
  }

  &__form {
    width: 100%;
    max-width: 42rem;
    position: relative;
    z-index: 1;
  }

  &__head {
    margin-bottom: 2.8rem;
  }

  &__title {
    margin-top: 1.2rem;
    font-size: 3rem;
    font-weight: 600;
    letter-spacing: -0.025em;
    color: $text-primary;
  }

  &__sub {
    margin-top: 0.6rem;
    color: $text-muted;
    font-size: 1.4rem;
  }

  &__foot {
    margin-top: 2rem;
    display: flex;
    align-items: center;
    gap: 0.8rem;
    flex-wrap: wrap;
    font-size: 1.3rem;
  }

  &__foot-sep {
    color: $text-faint;
  }

  &__link {
    color: $accent;
    font-weight: 500;
    transition: color 0.15s $ease-out;

    &:hover {
      color: $accent-glow;
      text-decoration: underline;
      text-underline-offset: 3px;
    }
  }

  &__legal {
    margin-top: 3rem;
    padding-top: 2rem;
    border-top: 1px solid $border-subtle;
    color: $text-muted;
    font-size: 1.15rem;
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
  }
}
</style>
