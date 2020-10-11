/**
 * title.js
 * Util mixin for title of page.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

function getTitle (vm) {
	const { title } = vm.$options
	if (title) {
		return typeof title === 'function'
			? title.call(vm)
			: title
	}
}

export default {
	created () {
		const title = getTitle(this)
		if (title) {
			document.title = "Verbis - " + title
		}
	}
}