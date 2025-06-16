import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/index/IndexView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'index',
      component: IndexView
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/register/RegisterView.vue')
    },
    {
      path: '/recovery',
      name: 'recovery',
      component: () => import('../views/recovery/RecoveryView.vue')
    },
    {
      path: '/link',
      name: 'link',
      component: () => import('../views/link/LinkView.vue')
    },
  ]
})

export default router
