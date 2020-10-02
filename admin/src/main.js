/**
 * main.js
 * Set up Vue & Axios
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
import App from "./App.vue";
import router from "./router";
import store from "./store";
import axios from 'axios'
require('./functions.js');

// Local
Vue.config.productionTip = false;

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

// Set defaults
axios.defaults.headers.common = {
	"token": store.state.apiToken || "",
	'Access-Control-Allow-Origin': '*/*',
	'Content-Type': 'application/json',
};
// axios.defaults.withCredentials = true
// axios.default.credentials = "include"
axios.defaults.baseURL = apiUrl

// Add a 401 response interceptor
// axios.interceptors.response.use(function (response) {
// 	return response;
// }, function (err) {
// 	Vue.prototype.helpers.handleResponse(err)
// })
//
// Assign axios globally
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
 * Vue
 * Create App
 *
 */
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");

/**
 * Dispatch global store
 *
 */
store.dispatch('getSiteConfig').catch(err => {
	Vue.prototype.helpers.handleResponse(err)
})
