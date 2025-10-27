<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { SignOut24Filled } from '@vicons/fluent'
import AuthServices from '@/services/auth.service'
import { useToast } from 'vue-toastification'

const router = useRouter()
const toast = useToast()
const userInfo = ref<{ firstName: string; lastName: string } | null>(null)

const loadUserInfo = async () => {
  try {
    const response = await AuthServices.verify()
    if (typeof response.data.data === 'object') {
      userInfo.value = response.data.data as { firstName: string; lastName: string }
    }
  } catch {
    console.error('Failed to load user info')
  }
}

const handleLogout = async () => {
  try {
    await AuthServices.logout()
  } catch {
    console.error('Logout request failed')
  } finally {
    localStorage.removeItem('user')
    toast.success('Logged out successfully')
    router.push({ name: 'auth' })
  }
}

onMounted(() => {
  loadUserInfo()
})
</script>

<template>
  <div class="fixed inset-0 flex flex-col bg-white">
    <div class="border-b bg-white">
      <div class="px-6 py-4 flex justify-between items-center">
        <div>
          <h1 class="text-lg font-bold">Job Tracker</h1>
        </div>
        <div class="flex items-center gap-4">
          <div v-if="userInfo" class="text-sm">
            <p class="font-medium">{{ userInfo.firstName }} {{ userInfo.lastName }}</p>
          </div>
          <button
            @click="handleLogout"
            class="p-2 hover:bg-gray-100 rounded-md transition-colors"
            title="Logout"
          >
            <SignOut24Filled class="text-red-600 w-6 h-6" />
          </button>
        </div>
      </div>
    </div>

    <div class="flex-1 overflow-hidden lg:bg-gray-50">
      <RouterView />
    </div>
  </div>
</template>
