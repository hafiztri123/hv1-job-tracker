<script setup lang="ts">
import { computed, HTMLAttributes } from 'vue';

const props = defineProps<{
  disabled?: boolean
  severity?: 'dark'
  label: string
}>()

const getBehaviour = computed<HTMLAttributes['class']>(() => {
  if (props.disabled) {
    return 'bg-gray-400 hover:cursor-not-allowed text-white'
  }

  return  'bg-black hover:cursor-pointer text-white'


})
</script>

<template>
  <button
    class="w-full  p-2 rounded-md text-center font-semibold "
    :class="getBehaviour"
  >
    {{ label }}
  </button>
</template>
