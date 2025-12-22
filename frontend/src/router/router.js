import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'

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
        meta: { hideHeader: true }
    },
    {
        path: '/profile',
        name: 'Profile',
        component: ProfileView,
        meta: { requiresAuth: true },
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
        meta: { requiresAuth: true },
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginView,
        meta: { hideFooter: true, hideHeader: true }
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

router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore()
    if (!authStore.user && authStore.accessToken) {
        try {
            await authStore.fetchProfile()
        } catch {
            await authStore.logout()
            return next('/auth')
        }
    }
    if (to.meta.requiresAuth && !authStore.user) {
        return next('/auth')
    }
    if (
        (to.path === '/auth' || to.path === '/login') &&
        authStore.user
    ) {
        return next('/profile')
    }

    next()
})


export default router
