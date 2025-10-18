<script setup lang="ts">
import { Form, Input, Button } from '@/components/common'
import logo from '../../../assets/logo.svg'
import { useTemplateRef } from 'vue'
import { AxiosResponse } from 'axios'
import { FetchDetailResponse } from '@/services/type/response.type'
import { AuthService } from '@/services'

const formRef = useTemplateRef<typeof Form>('formRef')

const login = async (): Promise<void> => {
  try {
    const payload = formRef?.value?.formValue
    const { data } = await AuthService.login({
      email: payload.email,
      password: payload.password,
    })

    localStorage.setItem(
      'user',
      JSON.stringify({
        token: data.data,
      }),
    )
  } catch (error) {
    console.error(error)
  }
}
</script>

<template>
  <div class="h-full grid grid-cols-1 lg:grid-cols-2 place-items-center bg-white lg:shadow-md">
    <div class="flex flex-col gap-4 w-full max-w-[400px]">
      <div class="flex flex-col gap-2 justify-center items-center">
        <img :src="logo" class="w-64" />
        <span class="text font-semibold">Start Managing Your Job Searching</span>
      </div>

      <Form ref="formRef">
        <template #default>
          <Input field="email" label="Email" required />
          <Input
            :min="8"
            :max="64"
            field="password"
            label="Password"
            type="password"
            required
          />
        </template>

        <template #footer>
          <Button @click="login" label="Sign In" :disabled="!formRef?.isFormValid" />
        </template>
      </Form>
    </div>

    <div>
      <p>Auth</p>
    </div>
  </div>
</template>
