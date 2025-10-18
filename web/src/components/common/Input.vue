<script setup lang="ts">
import { computed, inject, ref, shallowRef, watch } from 'vue'
import { Eye16Regular, EyeOff16Filled } from '@vicons/fluent'

const props = defineProps<{
  field: string
  label: string
  type?: string
  placeholder?: string
  min?: number
  max?: number
  required?: boolean
}>()

const formValue = inject<Record<string, string>>('formValue')
const registerValidation = inject<(field: string, isvalid: boolean) => void>('registerValidation')

const showWarning = ref<{
  max: boolean
  min: boolean
  required: boolean
}>({
  max: false,
  min: false,
  required: false,
})
const touched = shallowRef<boolean>(false)
const showPassword = shallowRef<boolean>(false)
let maxWarningTimeout: ReturnType<typeof setTimeout> | null = null

const isValid = computed<boolean>(() => {
  const length = value.value.length

  if (!length && props.required) return false
  if (props.min && length > 0 && length < props.min) return false
  if (props.max && length > props.max) return false

  return true
})

const value = computed<string>({
  get: () => formValue?.[props.field] || '',
  set: (val: string) => {
    if (formValue) {
      formValue[props.field] = val
    }
  },
})

const inputType = computed<string>(() => {
  if (showPassword.value) return 'text'
  return props.type || 'text'
})

const handleKeyDown = (e: KeyboardEvent) => {
  if (props.max && value.value.length >= props.max) {
    const allowedKeys = ['Backspace', 'ArrowLeft', 'ArrowRight', 'Delete', 'Tab']
    if (!allowedKeys.includes(e.key) && !e.ctrlKey && !e.metaKey) {
      showWarning.value.max = true

      if (maxWarningTimeout) clearTimeout(maxWarningTimeout)

      maxWarningTimeout = setTimeout(() => {
        showWarning.value.max = false
      }, 1000)
    }
  }
}

watch(
  () => value.value.length,
  () => {
    if (props.min) {
      showWarning.value.min = value.value.length < props.min
    }


    if (props.required) {
      showWarning.value.required = !value.value || value.value.length === 0
    }
  },
)


watch(
  isValid,
  (valid) => {
    if (registerValidation) {
      registerValidation(props.field, valid)
    }
  },
  { immediate: true },
)
</script>

<template>
  <div class="flex flex-col gap-1">
    <div class="relative">
      <label>{{ label }}</label>
      <input
        v-model="value"
        :type="inputType"
        @keydown="handleKeyDown"
        @blur="touched = true"
        :maxlength="max"
        :placeholder="placeholder ? placeholder : `Enter your ${label}`"
        class="relative border px-3 py-2 w-full rounded-md outline-none overflow-hidden transition-colors"
        :class="{
          'border-red-500 ring-2 ring-red-200': showWarning.max || showWarning.min || showWarning.required,

          'pr-10' : type === 'password',

        }

        "
      />

      <button v-if="type === 'password'" class="w-5 absolute inset-y-0 right-3 top-7 hover:cursor-pointer select-none" @click="showPassword = !showPassword">
        <Eye16Regular v-if="!showPassword"/>
        <EyeOff16Filled v-else/>
      </button>
    </div>

    <Transition
      enter-active-class="transition-opacity duration-200"
      leave-active-class="transition-opacity duration-200"
      enter-from-class="opacity-0"
      leave-to-class="opacity-0"
    >
      <span v-if="showWarning.max" class="text-sm text-red-500 font-medium">
        Maximum character limit is {{ max }}
      </span>

      <span v-else-if="showWarning.min" class="text-sm text-red-500 font-medium">
        Minimum character limit is {{ min }}
      </span>

      <span v-else-if="showWarning.required" class="text-sm text-red-500 font-medium">
        This field is required
      </span>
    </Transition>
  </div>
</template>
