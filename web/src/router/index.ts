import { beforeAccess } from '@/utils/beforeAccess'
import type { Component } from 'vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: Readonly<RouteRecordRaw[]> = [
  {
    path: '/',
    name: 'auth',
    component: (): Promise<Component> => import('@/layout/AuthLayout.vue'),
  },
  {
    path: '/home',
    component: (): Promise<Component> => import('@/layout/MainLayout.vue'),
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        name: 'home',
        component: (): Promise<Component> => import('@/view/MainView.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

beforeAccess(router)

export default router
