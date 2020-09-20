import Vue from 'vue'
import Vuex from 'vuex'
//import axios from 'axios'
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex)

export default new Vuex.Store({
	state: {
		auth: false,
		apiToken: "",
		userInfo: {},
		resources: [],
		logo: false,
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

			Vue.prototype.axios.defaults.headers.common = {
				"token": loginData.token,
			};
		},
		logout(state) {
			state.apiToken = ''
			state.auth = false
			state.userInfo = {}
			state.activeDomain = false
			state.activePage = false

			Vue.prototype.axios.defaults.headers.common = {
				"token": ''
			};
		},
		setResources(state, resources) {
			state.resources = resources;
		},
		logo(state, logo) {
			state.logo = logo
		}
	},
	actions: {
		getLogo(context) {
			if (!context.state.logo) {
				Vue.prototype.axios.get("/theme/logo")
					.then(res => {
						context.commit('logo', res.data.data);
					})
			}
		}
	},
	modules: {},
	plugins: [createPersistedState()],
})
