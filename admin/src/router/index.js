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
import axios from 'axios'

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
		component: () => import("../views/auth/Login.vue"),
		meta: { transitionName : 'slide' },
	},
	{
		path: "/password/reset",
		name: "send-password-reset",
		component: () => import("../views/auth/SendResetPassword.vue"),
		meta: { transitionName : 'slide' },
	},
	{
		path: "/password/reset/:token",
		name: "password-reset",
		component: () => import("../views/auth/ResetPassword.vue"),
		meta: { transitionName : 'slide' },
	},
	/**
	 * Install
	 *
	 */
	{
		path: "/install",
		name: "install",
		component: () => import("../views/install/Install.vue"),
		meta: { transitionName : 'slide' },
	},
	/**
	 * Site
	 *
	 */
	{
		path: "/site",
		name: "site",
		component: () => import("../views/site/ViewSite.vue"),
		meta: { transitionName : 'fade' },
	},
	/**
	 * Content
	 *
	 */
	{
		path: "/resources/:resource",
		name: "resources",
		component: () => import("../views/resources/Archive.vue"),
		meta: { transitionName : 'fade' },
	},
	{
		path: "/editor/:id",
		name: "editor",
		component: () => import("../views/resources/Editor.vue"),
		meta: { transitionName : 'fade' },
	},
	{ 	path: '/editor',
		redirect: '/editor/new',
		meta: {
			transitionName : 'fade',
		},
	},
	/**
	 * Assets
	 *
	 */
	{
		path: '/categories',
		name: 'categories',
		component: () => import('../views/content/Categories.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/categories/:id',
		name: 'categories-single',
		component: () => import('../views/content/CategoriesSingle.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/media',
		name: 'media',
		component: () => import('../views/content/Media.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/fields',
		name: 'fields',
		component: () => import('../views/content/Fields.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/forms',
		name: 'forms',
		component: () => import('../views/content/Forms.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/forms/:id',
		name: 'forms-single',
		component: () => import('../views/content/FormsSingle.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/users',
		name: 'users',
		component: () => import('../views/content/Users.vue'),
		meta: { transitionName : 'fade' },
	},
	/**
	 * Integrations
	 *
	 */
	{
		path: '/integrations/storage',
		name: 'storage',
		component: () => import('../views/intergrations/Storage.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/integrations/console',
		name: 'console',
		component: () => import('../views/intergrations/Console.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/integrations/proxies',
		name: 'proxies',
		component: () => import('../views/intergrations/Proxies.vue'),
		meta: { transitionName : 'fade' },
	},
	/**
	 * Account
	 *
	 */
	{
		path: '/profile',
		name: 'profile',
		component: () => import('../views/account/Profile.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/users/:id',
		name: 'edit-user',
		component: () => import('../views/account/Profile.vue'),
		meta: { transitionName : 'fade' },
	},
	/**
	 * Settings
	 *
	 */
	{
		path: '/settings/general',
		name: 'settings-general',
		component: () => import('../views/settings/General.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/settings/code-injection',
		name: 'settings-code-injection',
		component: () => import('../views/settings/CodeInjection.vue'),
		meta: { transitionName : 'fade' },

	},
	{
		path: '/settings/performance',
		name: 'settings-performance',
		component: () => import('../views/settings/Performance.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/settings/seo-meta',
		name: 'settings-seo-meta',
		component: () => import('../views/settings/SEOMeta.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/settings/media',
		name: 'settings-media',
		component: () => import('../views/settings/Media.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/settings/themes',
		name: 'settings-themes',
		component: () => import('../views/settings/Themes.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '/settings/redirects',
		name: 'settings-redirects',
		component: () => import('../views/settings/Redirects.vue'),
		meta: { transitionName : 'fade' },
	},
	/**
	 * Errors / Not Found
	 *
	 */
	{
		path: '/404',
		name: 'not-found',
		component: () => import('../views/errors/Error.vue'),
		meta: { transitionName : 'fade' },
	},
	{
		path: '*',
		redirect: '/404',
		meta: { transitionName : 'fade' },
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
  //base: process.env.BASE_URL,
  base: '/admin',
  routes
});

/**
 * Check Session
 * Small request to check if the session is
 * still valid on every request.
 *
 */
const checkSession = () => {
	axios.get("/session")
		.then(() => {})
		.catch(err => {
			Vue.prototype.helpers.handleResponse(err)
		})
}

/**
 * Protect Routes
 * Return redirect if not logged in.
 *
 */
router.beforeEach((to, from, next) => {
	if (to.name !== "install") {
		checkSession();
	}
	if (store.state.auth) {
		if (to.name === "login") {
			// Redirect to the page
			next({
				path: from.path
			});
			return;
		}
		next();
	} else {
		const excluded = [
			"login",
			"password-reset",
			"send-password-reset",
			"error",
			"install"
		]
		if (excluded.includes(to.name)) {
			// console.log("in");
			// checkSession();
			next();
		} else {
			next({
				name: "login"
			});
		}
	}
});



export default router;
