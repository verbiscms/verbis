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
import Meta from 'vue-meta'

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
	/**
	 * Auth
	 */
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
	/**
	 * Resources
	 *
	 */
	{
		path: "/resources/:resource",
		name: "resources",
		component: () => import("../views/resources/Archive.vue")
	},
	{
		path: "/editor/:id",
		name: "editor",
		component: () => import("../views/resources/Editor.vue")
	},
	{ 	path: '/editor',
		redirect: '/editor/new',
	},
	/**
	 * Content
	 *
	 */
	{
		path: '/categories',
		name: 'categories',
		component: () => import('../views/content/Categories.vue'),
	},
	{
		path: '/media',
		name: 'media',
		component: () => import('../views/content/Media.vue'),
	},
	{
		path: '/fields',
		name: 'fields',
		component: () => import('../views/content/Fields.vue'),
	},
	{
		path: '/users',
		name: 'users',
		component: () => import('../views/content/Users.vue'),
	},
	/**
	 * Account
	 *
	 */
	{
		path: '/profile',
		name: 'profile',
		component: () => import('../views/account/Profile.vue'),
	},
	/**
	 * Settings
	 *
	 */
	{
		path: '/settings/general',
		name: 'settings-general',
		component: () => import('../views/settings/General.vue'),
	},
	{
		path: '/settings/code-injection',
		name: 'settings-code-injection',
		component: () => import('../views/settings/CodeInjection.vue'),
	},
	{
		path: '/settings/performance',
		name: 'settings-performance',
		component: () => import('../views/settings/Performance.vue'),
	},
	{
		path: '/settings/seo-meta',
		name: 'settings-seo-meta',
		component: () => import('../views/settings/SEOMeta.vue'),
	},
	{
		path: '/settings/media',
		name: 'settings-media',
		component: () => import('../views/settings/Media.vue'),
	},
	/**
	 * Errors / Not Found
	 *
	 */
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
	{
		path: '*',
		redirect: '/404'
	},
];

/**
 * Setup Router
 *
 */
Vue.use(VueRouter);
Vue.use(Meta)
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
