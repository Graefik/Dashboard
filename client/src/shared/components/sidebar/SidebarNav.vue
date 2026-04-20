<script setup lang="ts">
import { useRoute } from "vue-router";
import NavIcon from "@/shared/components/sidebar/NavIcon.vue";

interface NavItem {
  label: string;
  to: string;
  icon: string;
  badge?: string;
}

const sections: { title?: string; items: NavItem[] }[] = [
  {
    items: [{ label: "Overview", to: "/", icon: "overview" }],
  },
  {
    title: "Console",
    items: [
      { label: "Logs", to: "/logs", icon: "logs" },
      { label: "Alerts", to: "/alerts", icon: "alerts" },
      { label: "Settings", to: "/settings", icon: "settings" },
    ],
  },
];

const route = useRoute();
const isActive = (to: string) => {
  if (to === "/") return route.path === "/";
  return route.path.startsWith(to);
};
</script>

<template>
  <nav class="nav">
    <div v-for="(section, sIdx) in sections" :key="sIdx" class="nav__section">
      <div v-if="section.title" class="nav__title">{{ section.title }}</div>
      <ul class="nav__list">
        <li v-for="item in section.items" :key="item.to">
          <a
            :href="item.to"
            class="nav__link"
            :class="{ 'nav__link--active': isActive(item.to) }"
            @click.prevent
          >
            <span class="nav__icon" aria-hidden="true">
              <NavIcon :name="item.icon" />
            </span>
            <span class="nav__label">{{ item.label }}</span>
            <span
              v-if="item.badge"
              class="nav__badge"
              :class="`nav__badge--${item.icon === 'errors' ? 'danger' : 'accent'}`"
              >{{ item.badge }}</span
            >
          </a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<style scoped lang="scss">
.nav {
  flex: 1;
  padding: 0 1.2rem;

  &__section {
    & + & {
      margin-top: 2.2rem;
      padding-top: 1.6rem;
      border-top: 1px solid $border-subtle;
    }
  }

  &__title {
    font-family: $font-mono;
    font-size: 1rem;
    font-weight: 500;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    padding: 0 0.8rem 0.8rem;
  }

  &__list {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
  }

  &__link {
    display: flex;
    align-items: center;
    gap: 1.1rem;
    padding: 0.85rem 0.9rem;
    border-radius: $radius-md;
    color: $text-secondary;
    font-size: 1.35rem;
    font-weight: 500;
    position: relative;
    transition: all 0.15s $ease-out;

    &::before {
      content: "";
      position: absolute;
      left: -1.2rem;
      top: 50%;
      transform: translateY(-50%);
      width: 3px;
      height: 0;
      background: $accent;
      border-radius: 0 2px 2px 0;
      transition: height 0.2s $ease-out;
    }

    &:hover {
      color: $text-primary;
      background: rgba(255, 255, 255, 0.03);
    }

    &--active {
      color: $text-primary;
      background: linear-gradient(90deg, $accent-soft 0%, transparent 100%);

      &::before {
        height: 60%;
      }

      .nav__icon {
        color: $accent;
      }
    }
  }

  &__icon {
    display: inline-flex;
    width: 1.8rem;
    height: 1.8rem;
    flex-shrink: 0;

    svg {
      width: 100%;
      height: 100%;
    }
  }

  &__label {
    flex: 1;
  }

  &__badge {
    font-family: $font-mono;
    font-size: 1rem;
    font-weight: 600;
    padding: 0.15rem 0.55rem;
    border-radius: $radius-pill;
    line-height: 1.4;

    &--accent {
      background: $accent-soft;
      color: $accent-glow;
    }

    &--danger {
      background: $severity-danger-soft;
      color: $severity-danger;
    }
  }
}
</style>
