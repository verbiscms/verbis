import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

Vue.config.productionTip = false

/**
 * Variables
 *
 */
// Not working
//const apiUrl = process.env.VUE_APP_APIURL
const baseUrl = "http://127.0.0.1:8080"
const apiUrl = "http://127.0.0.1:8080/api/v1"

/**
 * Axios
 *
 */
axios.defaults.headers.common = {
	"token": store.state.apiToken || "",
	'Access-Control-Allow-Origin': '*',
	'Content-Type': 'application/json',
};
//axios.defaults.withCredentials = true;
axios.defaults.baseURL = apiUrl
Vue.prototype.axios = axios;

/**
 * Mixins
 * Global instances applied to every vue instance.
 * Mixins must be instantiated *before* your call to new Vue(...)
 */

// Get global API Path
Vue.mixin({
	data: function () {
		return {
			get globalAPIPath() {
				return apiUrl
			},
			get globalBasePath() {
				return baseUrl
			}
		}
	}
})

/**
 * Set up new Vue instance.
 *
 */
new Vue({
	router,
	store,
	render: h => h(App)
}).$mount('#app')

/**
 * Dispatch global store
 *
 */
store.dispatch('getLogo')
