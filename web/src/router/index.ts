import {
  createRouter,
  createWebHistory,
} from 'vue-router'

import HomeView from '../views/HomeView.vue'
import CsView from '../views/CsView.vue'
import Cs1View from '../views/Cs1View.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/cs',
    name: 'Cs',
    component: CsView
  },
  {
    path: '/cs1',
    name: 'Cs1',
    component: Cs1View
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router