import Vue from "vue";
import Vuex from "vuex";
import axios from 'axios'
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		auth: false,
		apiToken: "",
		site: false,
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
			state.site = {}
			state.theme = {}
			state.activeDomain = false
			state.activePage = false
			axios.defaults.headers.common = {
				"token": ''
			};
		},
		setResources(state, resources) {
			state.resources = resources;
		},
	},
	actions: {
		login(context) {
			context.commit("login")
		},
		getSiteConfig(context) {
			if (!context.state.site) {
				axios.get("/site")
					.then(res => {
						context.state.site = res.data.data
						this.commit('site', res.data.data);
					})
			}
		},
		getThemeConfig(context) {
			if (!context.state.theme) {
				axios.get("/theme")
					.then(res => {
						context.state.theme = res.data.data
						this.commit('theme', res.data.data);
					})
			}
		}
	},
	modules: {},
	plugins: [createPersistedState()],
});