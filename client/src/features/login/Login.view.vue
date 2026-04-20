<script setup lang="ts">
import LoginForm from "@/features/login/components/login.form.vue";
import LiveConsole from "@/shared/ui/LiveConsole.vue";
</script>

<template>
  <div class="auth">
    <!-- Colonne gauche : brand + hero + console live ─────────────────── -->
    <aside class="auth__aside">
      <header class="auth__brand">
        <img src="@/assets/logo.png" alt="Graefik" class="auth__brand-mark" />
      </header>

      <section class="auth__hero">
        <div class="auth__eyebrow">
          <span class="auth__eyebrow-bullet" aria-hidden="true"></span>
          TRAEFIK · OBSERVATORY
        </div>
        <h1 class="auth__headline">
          Traefik,<br />
          <span class="auth__headline-em">à découvert.</span>
        </h1>
        <p class="auth__lede">
          Télémétrie brute, latence au p99, routers et certificats en un
          regard. Un seul <code>docker compose up</code> — tout reste chez
          vous.
        </p>
      </section>

      <LiveConsole />

      <footer class="auth__foot-line mono" aria-hidden="true">
        vos données · vos serveurs · FOSS AGPL · built with Go
      </footer>
    </aside>

    <!-- Colonne droite : form ────────────────────────────────────────── -->
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
  grid-template-columns: minmax(0, 1.05fr) minmax(0, 1fr);
  min-height: 100vh;

  @media (max-width: 960px) {
    grid-template-columns: 1fr;
  }

  &__aside {
    position: relative;
    padding: 5rem 5rem 4rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    background: $bg-rail;
    border-right: 1px solid $border-subtle;
    overflow: hidden;

    &::before {
      content: "";
      position: absolute;
      inset: 0;
      background:
        radial-gradient(
          ellipse 70rem 50rem at 10% -10%,
          rgba(0, 194, 178, 0.08) 0%,
          transparent 55%
        ),
        radial-gradient(
          circle at 1px 1px,
          rgba(255, 255, 255, 0.035) 1px,
          transparent 0
        );
      background-size: 100% 100%, 2.4rem 2.4rem;
      pointer-events: none;
    }

    &::after {
      content: "";
      position: absolute;
      right: -1px;
      top: 0;
      bottom: 0;
      width: 1px;
      background: linear-gradient(
        180deg,
        transparent 0%,
        $border-accent 30%,
        $border-accent 70%,
        transparent 100%
      );
      opacity: 0.4;
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
    align-items: center;
  }

  &__brand-mark {
    width: 7rem;
    height: auto;
    border-radius: $radius-md;
    box-shadow: 0 6px 18px -8px rgba(0, 0, 0, 0.6);
  }

  &__hero {
    max-width: 46rem;
  }

  &__eyebrow {
    display: inline-flex;
    align-items: center;
    gap: 0.7rem;
    font-family: $font-mono;
    font-size: 1.05rem;
    font-weight: 500;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    color: $accent;
  }

  &__eyebrow-bullet {
    width: 6px;
    height: 6px;
    background: $accent;
    box-shadow: 0 0 6px rgba(0, 194, 178, 0.6);
    border-radius: 1px;
  }

  &__headline {
    margin-top: 1.6rem;
    font-size: 4rem;
    font-weight: 500;
    line-height: 1.05;
    letter-spacing: -0.03em;
    color: $text-primary;

    @media (max-width: 1200px) {
      font-size: 3.4rem;
    }

    &-em {
      color: $accent;
      font-weight: 500;
    }
  }

  &__lede {
    margin-top: 1.8rem;
    font-size: 1.5rem;
    line-height: 1.55;
    color: $text-secondary;
    max-width: 42rem;

    code {
      padding: 0.15rem 0.5rem;
      margin: 0 0.15rem;
      background: $bg-inset;
      border: 1px solid $border-subtle;
      border-radius: $radius-xs;
      font-family: $font-mono;
      font-size: 1.3rem;
      color: $accent-glow;
    }
  }

  &__foot-line {
    margin-top: auto;
    font-size: 1rem;
    color: $text-faint;
    letter-spacing: 0.1em;
    text-transform: uppercase;
  }

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
