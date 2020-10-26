/**
 * functions.js
 * Custom Vue functions stored here.
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
import store from "./store";
import axios from 'axios'
import router from './router';
import Noty from 'noty';

class helpers {

	constructor() {
		this.debounceTime = 1000;
	}

	// Test for an empty object
	isEmptyObject(obj) {
		return Object.keys(obj).length === 0 && obj.constructor === Object
	}

	// noty specific for helper
	noty(msg) {
		new Noty({
			text: msg,
			type: 'error',
			timeout: 3500,
			theme: 'verbis',
			progressBar: false,
			layout: 'bottomRight'
		}).show();
	}

	// Handle response data
	handleResponse(data) {

		this.checkServer(data);

		if (store.state.auth) {
			if (data) {
				if (data.response.status === 401 || data.response.status === 429) {
					axios.post("/logout", {})
						.finally(() => {
							store.commit("logout");
							store.dispatch("getSiteConfig");
							router.push("/login");

							const errors = data.response.data.data.errors;
							if (errors['session'] !== undefined) {
								this.noty(data.response.data.message);
							}
						})
				} else if (data.response.status === 429) {
					this.noty(data.response.data.message);
				} else {
					this.error("Error occurred, please refresh the page.");
				}
			}
		} else {
			return Promise.reject(data)
		}
	}

	checkServer(data) {
		if (!data.status && !data.response) {
			router.push({ "name" : "error" })
		}
	}

	// Debounce input
	debounce(fn, immediate = false) {
		let timeout = this.debounceTime
		if (immediate) {
			timeout = 0;
		}
		let timeoutID = null
		return function () {
			clearTimeout(timeoutID)
			let args = arguments,
				that = this
			timeoutID = setTimeout(function () {
				fn.apply(that, args)
			}, timeout)
		}
	}

	// Set Cookie
	setCookie(name, value, days) {
		let expires = "";
		if (days) {
			let date = new Date();
			date.setTime(date.getTime() + (days*24*60*60*1000));
			expires = "; expires=" + date.toUTCString();
		}
		document.cookie = name + "=" + (value || "")  + expires + "; path=/";
	}

	// Get Cookie
	getCookie(name) {
		let nameEQ = name + "=";
		let ca = document.cookie.split(';');
		for(let i=0; i < ca.length; i++) {
			let c = ca[i];
			while (c.charAt(0) === ' ') c = c.substring(1,c.length);
			if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length,c.length);
		}
		return null;
	}

	// Delete cookie
	deleteCookie(name) {
		document.cookie = name + '=; Max-Age=0; path=/; domain=' + location.host;
	}

	// Calculate height & children of element
	setHeight(el) {
		let children = el.children;
		let height = 0;
		for (let i = 0; i < children.length; i++) {
			height += children[i].clientHeight;
		}
		el.style.maxHeight = height + "px";
		console.log(el.style.maxHeight)
	}

	// Capitalize first character of string
	capitalize(s) {
		if (typeof s !== 'string') return '';
		return s.charAt(0).toUpperCase() + s.slice(1);
	}
}

Vue.prototype.helpers = new helpers();