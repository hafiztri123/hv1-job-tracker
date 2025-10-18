import router from '@/router'
import type { AxiosError, AxiosInstance, InternalAxiosRequestConfig } from 'axios'
import axios from 'axios'

export const createAxiosInstance = (path: string): AxiosInstance => {
  const url = import.meta.env.VITE_API_URL || 'http://localhost:3000/api/v1'
  const baseURL = `${url}/${path}`

  const instance = axios.create({
    baseURL,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
  })

  setAxiosInstanceRequestInterceptor(instance)
  setAxiosInstanceResponseInterceptor(instance)

  return instance
}

const setAxiosInstanceRequestInterceptor = (instance: AxiosInstance) => {
  instance.interceptors.request.use((config: InternalAxiosRequestConfig) => {
    const user = JSON.parse(localStorage.getItem('user') || '{}')

    if (user.token && config.headers) {
      config.headers.Authorization = `Bearer ${user.token}`
    }

    return config
  }, (error: AxiosError) => {
    return Promise.reject(error)
  })
}

const setAxiosInstanceResponseInterceptor = (instance: AxiosInstance) => {
  instance.interceptors.response.use((response) => response, (error: AxiosError) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('user')
      router.push({ name: 'auth' })
    }

    return Promise.reject(error)
  })
}

