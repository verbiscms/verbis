import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store/index'

Vue.use(VueRouter)

/*
 * Main Routes
 *
 */
const routes = [
    {
        path: '/',
        name: 'home',
        component: () => import('../views/Home.vue'),
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/auth/Login.vue'),
    },
    {
        path: '/password/reset',
        name: 'password-reset',
        component: () => import('../views/auth/SendResetPassword.vue'),
    },
    {
        path: '/password/reset/:token',
        name: 'password-reset',
        component: () => import('../views/auth/ResetPassword.vue'),
    },
    {
        path: '/content/:resource',
        name: 'archive',
        component: () => import('../views/pages/Archive.vue'),
    },
    {
        path: '/editor/:id',
        name: 'single',
        component: () => import('../views/pages/Single.vue'),
    },
    // {
    //     path: '/404',
    //     name: 'not-found',
    //     component: () => import('../views/errors/Error.vue'),
    // },
    // {
    //     path: '*',
    //     redirect: '/404'
    // },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

// Protect routes, return redirect if not logged in.
router.beforeEach((to, from, next) => {
   if (store.state.auth) {
       if (to.name === "login") {
            // Redirect to the page
            next({
                path: from.path
            })
       } else {
           next();
       }
   } else {
       if (to.name === "login" || to.name === "password-reset") {
           next();
       } else {
           next({
               name: "login"
           })
       }
   }
});

export default router
