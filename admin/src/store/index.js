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
		session: "",
		userInfo: {},
		site: {},
		users: [],
		theme: {},
		roles: [],
		profilePicture: false,
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
			state.site = {};
			state.users = [];
			state.theme = {};
			state.profilePicture = false;
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
		setRoles(state, roles) {
			state.roles = roles;
		},
		setUsers(state, users) {
			state.users = users;
		},
		setProfilePicture(state, picture) {
			state.profilePicture = picture;
		}
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
			});
		},
		/*
		 * getProfilePicture()
		 */
		getProfilePicture() {
			if (!this.state.userInfo['profile_picture_id']) return;
			return new Promise((resolve, reject) => {
				if (!this.state.profilePicture) {
					axios.get('/media/' + this.state.userInfo['profile_picture_id'])
						.then(res => {
							const picture = res.data.data;
							this.commit("setProfilePicture", picture)
							resolve(picture)
						})
						.catch(err => {
							reject(err)
						});
				} else {
					resolve(this.state.profilePicture)
				}
			});
		},
		/*
		 * getRoles()
		 */
		getRoles() {
			return new Promise((resolve, reject) => {
				if (!this.state.roles.length) {
					axios.get("/roles")
						.then(res => {
							const roles = res.data.data;
							this.commit("setRoles", roles);
							resolve(roles);
						})
						.catch(err => {
							reject(err);
						});
				} else {
					resolve(this.state.roles)
				}
			})

		},
	},
	modules: {},
	plugins: [createPersistedState()],
});

