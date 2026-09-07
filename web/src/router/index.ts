import {
  createRouter,
  createWebHistory,
} from 'vue-router'

import HomeView from '../views/HomeView.vue'
import CsView from '../views/CsView.vue'
import Cs1View from '../views/Cs1View.vue'
import TiaozhuanView from '../views/TiaozhuanView.vue'
import VModel from '../views/V-Model.vue'

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
  },
  {
    path: '/tiaozhuan',
    name: 'tiaozhuan',
    component: TiaozhuanView
  },
  {
    path: '/v-model',
    name: 'V-Model',
    component: VModel
  }

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router