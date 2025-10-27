<script setup lang="ts">
import { computed, provide, reactive } from 'vue'

interface Props {
  formValue: Record<string, string>
  onRegisterValidation?: (field: string, isValid: boolean) => void
}

const props = defineProps<Props>()
const emit = defineEmits<{ clear: [] }>()

const validationRegistry = reactive<Record<string, boolean>>({})

const isFormValid = computed<boolean>(() => {
  return Object.values(validationRegistry).every((isValid) => isValid)
})

const registerValidation = (field: string, isvalid: boolean) => {
  validationRegistry[field] = isvalid
  if (props.onRegisterValidation) {
    props.onRegisterValidation(field, isvalid)
  }
}

const clearInput = () => {
  emit('clear')
}

provide('formValue', props.formValue)
provide('registerValidation', registerValidation)

defineExpose({
  formValue: props.formValue,
  clearInput,
  isFormValid,
})
</script>

<template>
  <div class="flex flex-col gap-2">
    <slot name="default" />
  </div>
</template>
