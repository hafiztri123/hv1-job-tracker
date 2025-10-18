<script setup lang="ts">
import { computed, provide, reactive } from 'vue'

const formValue = reactive<Record<string, string>>({})
const validationRegistry = reactive<Record<string, boolean>>({})

const isFormValid = computed<boolean>(() => {
  return Object.values(validationRegistry).every((isValid) => isValid)
})

const registerValidation = (field: string, isvalid: boolean) => {
  validationRegistry[field] = isvalid
}


provide('formValue', formValue)
provide('registerValidation', registerValidation)

defineExpose({
  formValue,
  isFormValid
})
</script>

<template>
  <slot name="default" />
  <slot name="footer" />
</template>
