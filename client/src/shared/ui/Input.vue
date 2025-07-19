<template>
  <input
    class="input"
    :type="type"
    :placeholder="placeholder"
    :disabled="disabled"
    :required="required"
  />
</template>

<script setup lang="ts">
interface Props {
  type?: "text" | "email" | "password";
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: "text",
  placeholder: "",
  disabled: false,
  required: false,
});

defineEmits<{
  input: [value: string];
}>();
</script>
<style scoped lang="scss">
.input {
  width: 100%;
  padding: 14px;
  border-radius: 12px;
  border: 2px solid rgba(255, 255, 255, 0.1);
  background-color: rgba(255, 255, 255, 0.05);
  color: $white;
  transition: all 0.3s ease;

  &:focus {
    border-color: $cta;
    outline: none;
    background-color: rgba(255, 255, 255, 0.1);
    box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.1);
  }
}
</style>
