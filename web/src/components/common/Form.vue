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

const clearInput = () => {
  Object.keys(formValue).forEach((key) => {
    formValue[key] = ''
  })

}

provide('formValue', formValue)
provide('registerValidation', registerValidation)

defineExpose({
  formValue,
  clearInput,
  isFormValid,
})
</script>

<template>
  <div class="flex flex-col gap-2">
    <slot name="default" />
  </div>
</template>
