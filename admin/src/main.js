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
import VueNestable from 'vue-nestable'
import { PrismEditor } from 'vue-prism-editor';
import {globalMixin} from "@/util/global";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

require('./functions.js');

// Local
Vue.config.productionTip = false;

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
axios.defaults.withCredentials = true;
axios.defaults.baseURL = "/api/v1";

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

// Nestable
Vue.use(VueNestable)

// Element
Vue.use(ElementUI);

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

// Element
Vue.use(ElementUI);

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
if (window.location.pathname !== "/admin/install") {
	store.dispatch('getSiteConfig').catch(err => {
		Vue.prototype.helpers.handleResponse(err)
	})
}
