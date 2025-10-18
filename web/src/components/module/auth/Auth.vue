<script setup lang="ts">
import { Form, Input, Button } from '@/components/common'
import logo from '../../../assets/logo.svg'
import placeholder from '../../../assets/placeholder.svg'
import { ref, shallowRef, useTemplateRef } from 'vue'
import { AuthService } from '@/services'
import { useToast } from 'vue-toastification'
import { AxiosError } from 'axios'
import { camelToTitle } from '@/utils/camelCaseSplit'
import { useRouter } from 'vue-router'

const toast = useToast()
const router = useRouter()
const formRef = useTemplateRef<typeof Form>('formRef')
const isRegister = shallowRef<boolean>(false)
const fieldError = ref<{ field: string; message: string }[]>([])

const handleSubmit = async (): Promise<void> => {
  fieldError.value = []

  try {
    const payload = formRef?.value?.formValue
    if (isRegister.value) {
      await AuthService.register(payload)
    } else {
      const { data } = await AuthService.login(payload)
      localStorage.setItem(
        'user',
        JSON.stringify({
          token: data.data,
        }),
      )
    }

    toast.success(`Success, ${isRegister.value ? 'register' : 'login'} success`)

    if(isRegister.value) {
      handleFormSwitch()
    } else {
      router.push({ name: 'home' })
    }
  } catch (error: unknown) {
    toast.error(`Error, ${isRegister.value ? 'register' : 'login'} failed`)
    if (error instanceof AxiosError) {
      if (error.response?.data.error && Array.isArray(error.response?.data.error)) {
        error.response?.data.error.forEach((err: { field: string; message: string }) => {
          fieldError.value.push(err)
        })
      }
    }
  }
}

const handleFormSwitch = (): void => {
  isRegister.value = !isRegister.value
  formRef.value?.clearInput()
  fieldError.value = []
}
</script>

<template>
  <div class="h-full grid grid-cols-1 lg:grid-cols-2 bg-white lg:shadow-md">
    <div class="flex items-center justify-center p-6 lg:p-8">
      <div class="flex flex-col gap-4 w-full max-w-[400px]">
        <div class="flex flex-col gap-2 justify-center items-center">
          <img :src="logo" class="w-64" />
          <span class="text font-semibold">Start Managing Your Job Searching</span>
        </div>

        <div class="flex flex-col gap-3">
          <Form ref="formRef">
            <template v-if="!isRegister" #default>
              <Input :invalid="fieldError.length > 0" field="email" label="Email" required />
              <Input
                :invalid="fieldError.length > 0"
                field="password"
                label="Password"
                type="password"
                required
              />
            </template>

            <template v-else #default>
              <Input field="email" label="Email" required />
              <Input field="firstName" label="First Name" required />
              <Input field="lastName" label="Last Name" required />
              <Input
                :min="8"
                :max="64"
                field="password"
                label="Password"
                type="password"
                required
              />
            </template>
          </Form>

          <div v-if="fieldError.length" class="flex flex-col gap-1">
            <ul class="list-disc list-inside">
              <li class="text-xs text-red-500" :key="value.field" v-for="value in fieldError">
                {{ camelToTitle(value.field) }}: {{ value.message }}
              </li>
            </ul>
          </div>
        </div>

        <div class="flex flex-col mt-8">
          <Button @click="handleSubmit" label="Sign In" :disabled="!formRef?.isFormValid" />
          <span
            @click="handleFormSwitch"
            class="text-xs mt-2 text-blue-500 hover:underline hover:cursor-pointer"
          >
            {{
              isRegister ? 'Already have an account? Sign in' : "Didn't have an account? Sign up"
            }}
          </span>
        </div>
      </div>
    </div>

    <div
      class="hidden lg:flex items-center justify-center p-8 bg-gradient-to-br from-blue-50 to-indigo-100"
    >
      <img
        :src="placeholder"
        class="w-full h-full max-w-2xl object-contain drop-shadow-lg"
        alt="Job searching illustration"
      />
    </div>
  </div>
</template>
