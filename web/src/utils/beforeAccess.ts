import { AuthService } from '@/services'
import { Router } from 'vue-router'

export const beforeAccess = (router: Router): void => {
  router.beforeEach(async (to, from, next) => {
    const requiresAuth = to.meta.requiresAuth

    if (!requiresAuth) {
      return next()
    }

    const user = JSON.parse(localStorage.getItem('user') ?? '{}')
    if (!user.token) {
      return next({ name: 'auth' })
    }

    try {
      const { data } = await AuthService.verify()

      if (data) {
        user.id = data.data
        localStorage.setItem('user', JSON.stringify(user))
        return next()
      } else {
        throw new Error('Invalid token')
      }
    } catch (err) {
      if (err) {
        localStorage.removeItem('user')
        return next({ name: 'auth' })
      }
    }
  })
}
