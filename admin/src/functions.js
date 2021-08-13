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

	// Rename object key
	renameKey(object, key, newKey) {
		const clonedObj =  Object.assign({}, object);
		const targetKey = clonedObj[key];
		delete clonedObj[key];
		clonedObj[newKey] = targetKey;
		return clonedObj;
	}

	// Handle response data
	handleResponse(data) {
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
					this.noty("Error occurred, please refresh the page.");
				}
			}
		} else {
			return Promise.reject(data)
		}
	}

	// Calculate & set height & children of element
	setHeight(el) {
		let children = el.children;
		let height = 0;
		for (let i = 0; i < children.length; i++) {
			height += children[i].clientHeight;
		}
		el.style.maxHeight = height + "px";
	}

	// Capitalize first character of string
	capitalize(s) {
		if (typeof s !== 'string') return '';
		return s.charAt(0).toUpperCase() + s.slice(1);
	}

	// Validate given email
	validateEmail(email) {
		// eslint-disable-next-line no-useless-escape
		const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
		return re.test(String(email).toLowerCase());
	}

	// Validate given url
	validateUrl(url) {
		let pattern = new RegExp('^(https?:\\/\\/)?'+ // protocol
			'((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|'+ // domain name
			'((\\d{1,3}\\.){3}\\d{1,3}))'+ // OR ip (v4) address
			'(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*'+ // port and path
			'(\\?[;&a-z\\d%_.~+=-]*)?'+ // query string
			'(\\#[-a-z\\d_]*)?$','i'); // fragment locator
		return pattern.test(url)
	}

	// Generate random alpha num string
	randomString(length) {
		const chars = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
		let result = '';
		for (let i = length; i > 0; --i) result += chars[Math.floor(Math.random() * chars.length)];
		return result;
	}
}

Vue.prototype.helpers = new helpers();
