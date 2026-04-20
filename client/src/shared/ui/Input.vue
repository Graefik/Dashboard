<template>
  <div class="input-wrap" :class="{ 'input-wrap--error': error }">
    <span v-if="icon" class="input-wrap__icon" aria-hidden="true">
      <slot name="icon" />
    </span>
    <input
      :id="id"
      class="input"
      :class="{ 'input--with-icon': icon }"
      :type="type"
      :placeholder="placeholder"
      :disabled="disabled"
      :value="modelValue"
      :autocomplete="autocomplete"
      :aria-invalid="ariaInvalid"
      :aria-describedby="ariaDescribedby"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement).value)
      "
    />
  </div>
</template>

<script setup lang="ts">
interface Props {
  id?: string;
  type?: "text" | "email" | "password" | "number" | "search";
  placeholder?: string;
  disabled?: boolean;
  modelValue?: string;
  autocomplete?: string;
  ariaInvalid?: boolean;
  ariaDescribedby?: string;
  error?: boolean;
  icon?: boolean;
}

withDefaults(defineProps<Props>(), {
  type: "text",
  placeholder: "",
  disabled: false,
  modelValue: "",
  ariaInvalid: false,
  error: false,
  icon: false,
});

defineEmits<{
  "update:modelValue": [value: string];
}>();
</script>

<style scoped lang="scss">
.input-wrap {
  position: relative;
  display: flex;
  align-items: center;

  &__icon {
    position: absolute;
    left: 1.2rem;
    color: $text-muted;
    pointer-events: none;
    display: inline-flex;
  }

  &--error .input {
    border-color: rgba(241, 105, 105, 0.5);

    &:focus {
      border-color: $severity-danger;
      box-shadow: 0 0 0 3px rgba(241, 105, 105, 0.15);
    }
  }
}

.input {
  width: 100%;
  height: 4.4rem;
  padding: 0 1.4rem;
  border-radius: $radius-md;
  border: 1px solid $border-default;
  background: $bg-inset;
  color: $text-primary;
  font-size: 1.4rem;
  font-family: $font-sans;
  transition:
    border-color 0.15s $ease-out,
    box-shadow 0.15s $ease-out,
    background 0.15s $ease-out;

  &::placeholder {
    color: $text-faint;
  }

  &--with-icon {
    padding-left: 3.6rem;
  }

  &:hover:not(:disabled) {
    border-color: $border-strong;
  }

  &:focus {
    outline: none;
    border-color: $border-accent;
    background: $bg-surface;
    box-shadow: 0 0 0 3px $accent-soft;
  }

  &:disabled {
    opacity: 0.55;
    cursor: not-allowed;
  }
}
</style>
