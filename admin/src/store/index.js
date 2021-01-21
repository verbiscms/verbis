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
		site: false,
		users: [],
		theme: false,
		roles: [],
		profilePicture: false,
	},
	mutations: {
		login(state, loginData) {
			state.apiToken = loginData.token;
			state.auth = true;
			state.userInfo =  loginData;
			axios.defaults.headers.common = {
				"token": loginData.token,
			};
		},
		logout(state) {
			state.apiToken = ''
			state.auth = false
			state.userInfo = {}
			state.site = false;
			state.users = [];
			state.theme = false;
			state.roles = [];
			state.profilePicture = false;
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
		// setSession(state, session) {
		// 	//Vue.prototype.helpers.setCookie("verbis-session", session, 1)
		// },
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
		},
		addUser(state, user) {
			state.users.push(user);
		}
	},
	actions: {
		/*
	 	 * getSiteConfig()
		 * Get site config from API & commit.
		 */
		getSiteConfig(context) {
			return new Promise((resolve, reject) => {
				if (!this.state.site) {
					axios.get("/site")
						.then(res => {
							const site = res.data.data
							context.state.site = site;
							this.commit('setSite', res.data.data);
							resolve(site)
						})
						.catch(err => {
							reject(err);
						});
				} else {
					resolve(this.state.site);
				}
			})
		},
		/*
		 * getTheme()
		 * Get them description, resources, assets paths from API & commit.
		 */
		getTheme() {
			return new Promise((resolve, reject) => {
				if (!this.state.theme) {
					axios.get(`/theme`)
						.then(res => {
							const theme = res.data.data;
							this.commit("setTheme", theme);
							resolve(theme);
						})
						.catch(err => {
							reject(err)
						})
				} else {
					resolve(this.state.theme)
				}
			});
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
							const users = res.data.data;
							this.commit("setUsers", users);
							resolve(users);
						})
						.catch(err => {
							reject(err);
						})
				} else {
					resolve(this.state.users);
				}
			});
		},
		/*
		 * getProfilePicture()
		 */
		getProfilePicture() {
			if (!this.state.userInfo['profile_picture_id']) return;
			return new Promise((resolve, reject) => {
				// NOTE: Don't check if the store already contains profile picture,
				// otherwise sidebar won't be updated.
				axios.get('/media/' + this.state.userInfo['profile_picture_id'])
					.then(res => {
						const picture = res.data.data;
						this.commit("setProfilePicture", picture);
						resolve(picture);
					})
					.catch(err => {
						this.commit("setProfilePicture", false);
						reject(err);
					});
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
					resolve(this.state.roles);
				}
			});
		},
	},
	modules: {},
	plugins: [createPersistedState()],
});

