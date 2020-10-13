/**
 * store/index.js
 * Set up of Vuex
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
import Vuex from "vuex";
import axios from 'axios'
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		auth: false,
		apiToken: "",
		site: {},
		session: "",
		userInfo: {},
		users: [],
		theme: {},
	},
	mutations: {
		login(state, loginData) {
			state.apiToken = loginData.token
			state.auth = true
			state.userInfo =  loginData;
			axios.defaults.headers.common = {
				"token": loginData.token,
			};
		},
		logout(state) {
			state.apiToken = ''
			state.auth = false
			state.userInfo = {}
			Vue.prototype.helpers.deleteCookie("verbis-session")
			axios.defaults.headers.common = {
				"token": ''
			};
		},
		setSite(state, site) {
			state.site = site;
		},
		setUser(state, user) {
			state.userInfo = user;
		},
		setSession(state, session) {
			Vue.prototype.helpers.setCookie("verbis-session", session, 1)
			state.session = session;
		},
		setTheme(state, theme) {
			state.theme = theme;
		},
		setResources(state, resources) {
			state.resources = resources;
		},
		setUsers(state, users) {
			state.users = users;
		},
	},
	actions: {
		/*
	 	 * getSiteConfig()
		 * Get site config from API & commit.
		 */
		getSiteConfig(context) {
			return new Promise((resolve, reject) => {
				const site = context.state.site
				if (Object.keys(site).length === 0 && site.constructor === Object) {
					axios.get("/site")
						.then(res => {
							context.state.site = res.data.data
							this.commit('setSite', res.data.data);
							resolve()
						})
						.catch(() => {
							reject();
						});
				} else {
					resolve();
				}
			})
		},
		/*
	 	 * getUsers()
		 * Get users from API & commit.
		 */
		getUsers() {
			return new Promise((resolve, reject) => {
				if (!this.state.users.length) {
					axios.get(`/users`)
						.then(res => {
							const users = res.data.data
							this.commit("setUsers", users)
							resolve(users)
						})
						.catch(err => {
							reject(err)
						})
				} else {
					resolve(this.state.users)
				}
			})
		}
	},
	modules: {},
	plugins: [createPersistedState()],
});

