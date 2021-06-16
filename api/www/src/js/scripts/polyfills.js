/**
 * polyfills.js
 * Polyfills for Internet Explorer & Older Browsers.
 * @author Ainsley Clark
 * @author URL:   https://www.ainsleyclark.com
 * @author Email: info@ainsleyclark.com
 */

 /**
 * Require * Import
 * 
 */

 //Vendor
import * as smoothscroll from 'smoothscroll-polyfill';

/**
 * Fix smoothscroll on Safari & IE
 *
 */
smoothscroll.polyfill();

/**
 * Detect Internet Explorer
 *
 */
function isInternetExplorer() {
	let ua = window.navigator.userAgent;
	let msie = ua.indexOf("MSIE ");

	return msie > 0;
}

/**
 * Add ie class to html if any type of Internet Explorer
 *
 */
if (isInternetExplorer()) {
	html.classList.add('ie');
}

/**
 * Polyfill for .forEach
 *
 */
if ('NodeList' in window && !NodeList.prototype.forEach) {
	console.info('polyfill for IE11');
	NodeList.prototype.forEach = function (callback, thisArg) {
		thisArg = thisArg || window;
		for (var i = 0; i < this.length; i++) {
			callback.call(thisArg, this[i], i, this);
		}
	};
}

/**
 * Polyfill for .closest()
 *
 */
if (!Element.prototype.matches) {
	Element.prototype.matches = Element.prototype.msMatchesSelector ||
		Element.prototype.webkitMatchesSelector;
}

if (!Element.prototype.closest) {
	Element.prototype.closest = function(s) {
		var el = this;

		do {
			if (Element.prototype.matches.call(el, s)) return el;
			el = el.parentElement || el.parentNode;
		} while (el !== null && el.nodeType === 1);
		return null;
	};
}