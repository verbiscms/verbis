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
		resources: [],
	},
	mutations: {
		login(state, loginData) {
			state.apiToken = loginData.token
			state.auth = true
			state.userInfo = {
				id: loginData.id,
				firstName: loginData.first_name,
				lastName: loginData.last_name,
				email: loginData.email,
				accessLevel: loginData.access_level,
				email_verified_at: true,
			}
			axios.defaults.headers.common = {
				"token": loginData.token,
			};
		},
		logout(state) {
			state.apiToken = ''
			state.auth = false
			state.userInfo = {}
			state.activeDomain = false
			state.activePage = false
			Vue.prototype.helpers.deleteCookie("verbis-session")
			axios.defaults.headers.common = {
				"token": ''
			};
		},
		setSite(state, site) {
			state.site = site;
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
	},
	modules: {},
	plugins: [createPersistedState()],
});

