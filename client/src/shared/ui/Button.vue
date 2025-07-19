<template>
  <button
    :type="type"
    :disabled="disabled"
    :class="['button', variant]"
    @click="$emit('click', $event)"
  >
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
interface Props {
  type?: "button" | "submit";
  disabled?: boolean;
  variant?: "primary";
}

withDefaults(defineProps<Props>(), {
  type: "button",
  disabled: false,
});

defineEmits<{
  click: [event: MouseEvent];
}>();
</script>

<style scoped lang="scss">
.button {
  background-color: $cta;
  border-radius: 12px;
  padding: 14px 32px;
  font-size: 1.1rem;
  font-weight: 600;
  border: none;
  width: 100%;
  color: $white;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 1rem;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);

  &:hover {
    background-color: darken($cta, 5%);
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
  }

  &:active {
    transform: translateY(0);
  }
}
</style>
