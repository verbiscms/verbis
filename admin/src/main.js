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
import titleMixin from './util/title';
import VueNoty from 'vuejs-noty'
import OnLoad from 'vue-onload'
import { PrismEditor } from 'vue-prism-editor';
import {globalMixin} from "@/util/global";
require('./functions.js');
//require('dotenv').config({path: '/Users/ainsley/Desktop/Reddico/apis/verbis/.env'})


// Local
Vue.config.productionTip = false;


/**
 * Variables
 *
 */
// Not working
// const url = process.env['VUE_APP_URL'];
// const port = process.env['VUE_APP_PORT'];
//const apiUrl = url + ":" + port;
const apiUrl = "https://verbiscms.com/api/v1"
console.log(apiUrl)

console.log(process.env)

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
axios.defaults.withCredentials = true
axios.defaults.baseURL = apiUrl

// Assign axios globally
Vue.prototype.axios = axios;

/**
 * Mixins
 * Global instances applied to every vue instance.
 * Mixins must be instantiated *before* your call to new Vue(...)
 */

// Global
Vue.mixin(globalMixin);

// Title
Vue.mixin(titleMixin);

/**
 * Plugins
 *
 */

// Noty
Vue.use(
	VueNoty, {
		timeout: 2500,
		theme: 'verbis',
		progressBar: true,
		layout: 'bottomRight'
	},
);

// Moment
Vue.use(require('vue-moment'));

// OnLoad
Vue.use(OnLoad)

/**
 * Components
 *
 */
Vue.component('PrismEditor', PrismEditor);

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
