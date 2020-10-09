/**
 * router/router.js
 * Set up of Vue Router
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

/**
 * Require * Import
 *
 */

// Vendor
import Vue from "vue";
import VueRouter from "vue-router";
import store from '../store/index'



/*
 * Main Routes
 *
 */
const routes = [
	{
		path: "/",
		name: "home",
		component: () => import("../views/Home.vue")
	},
	{
		path: "/login",
		name: "login",
		component: () => import("../views/auth/Login.vue")
	},
	{
		path: "/password/reset",
		name: "password-reset",
		component: () => import("../views/auth/SendResetPassword.vue")
	},
	// {
	//     path: "/password/reset/:token",
	//     name: "password-reset",
	//     component: () => import("../views/auth/ResetPassword.vue")
	// },
	{
		path: "/content/:resource",
		name: "archive",
		component: () => import("../views/pages/Archive.vue")
	},
	{
		path: "/editor/:id",
		name: "editor",
		component: () => import("../views/pages/Editor.vue")
	},
	{
		path: "/editor/:id",
		name: "settings",
		component: () => import("../views/pages/Editor.vue")
	},
	{
		path: "/editor/:id",
		name: "media",
		component: () => import("../views/pages/Editor.vue")
	},
	{
		path: '/404',
		name: 'not-found',
		component: () => import('../views/errors/Error.vue'),
	},
	{
		path: '/error',
		name: 'error',
		component: () => import('../views/errors/ServerDown.vue'),
	},
	// {
	//     path: '*',
	//     redirect: '/404'
	// },
];

/**
 * Setup Router
 *
 */
Vue.use(VueRouter);
const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

/**
 * Protect Routes
 * Return redirect if not logged in.
 *
 */
router.beforeEach((to, from, next) => {
	if (store.state.auth) {
		if (to.name === "login") {
			// Redirect to the page
			next({
				path: from.path
			});
		} else {
			next();
		}
	}else {
		if (to.name === "login" || to.name === "password-reset" || to.name === "error") {
			next();
		} else {
			next({
				name: "login"
			});
		}
	}
});

export default router;
