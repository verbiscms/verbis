import {createApp} from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import axios from 'axios'

/**
 * Variables
 *
 */
// Not working
//const apiUrl = process.env.VUE_APP_APIURL
const baseUrl = "http://127.0.0.1:8080"
const apiUrl = "http://127.0.0.1:8080/api/v1"

/**
 * Vue
 * Create App
 *
 */
const app = createApp(App)

/**
 * Axios
 *
 */
axios.defaults.headers.common = {
	"token": store.state.apiToken || "",
	'Access-Control-Allow-Origin': '*',
	'Content-Type': 'application/json',
};
axios.defaults.baseURL = apiUrl
app.config.globalProperties.axios = axios

/**
 * Mixins
 * Global instances applied to every vue instance.
 * Mixins must be instantiated *before* your call to new Vue(...)
 */

// Get global API Path
app.mixin({
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
 * Mount Vue
 *
 */
app.use(store)
	.use(router)
	.mount("#app");

/**
 * Dispatch global store
 *
 */
store.dispatch('getSiteConfig')
