import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";
import axios from 'axios'

export default createStore({
	state: {
		auth: false,
		apiToken: "",
		siteInfo: false,
		userInfo: {},
		users: [],
		resources: [],
	},
	mutations: {
		siteInfo() {
			console.log("hello")
		},
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
			axios.defaults.headers.common = {
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
		getSiteConfig(context) {
			if (!context.state.siteInfo) {
				axios.get("/site")
					.then(res => {
						console.log(res)
						context.state.siteInfo = res.data.data
						this.commit('siteInfo', res.data.data);
					})
			}
		},
		logout({ commit }, payload) {
			this.axios.post("/logout", {})
				.then(() => {
					commit('logout', payload)
					location.reload()
				});
		}
	},
	modules: {},
	plugins: [createPersistedState()],
});
