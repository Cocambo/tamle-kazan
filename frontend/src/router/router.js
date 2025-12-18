import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import RestaurantsView from '@/views/RestaurantsView.vue'
import FavouriteView from '@/views/FavouriteView.vue'
import LoginView from '@/views/LoginView.vue'
import AuthenticationView from '@/views/AuthenticationView.vue'
import ProfileView from '@/views/ProfileView.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    meta: { hideHeader: true  }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: ProfileView,
  },
  {
    path: '/restaurants',
    name: 'Restaurants',
    component: RestaurantsView,
  },
  {
    path: '/favourite',
    name: 'Favourite',
    component: FavouriteView,
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginView,
    meta: { hideFooter: true, hideHeader: true  }
  },
  {
    path: '/auth',
    name: 'Authentication',
    component: AuthenticationView,
    meta: { hideFooter: true, hideHeader: true }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 } 
  },
})

export default router
