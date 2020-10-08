/**
 * fields.js
 * Common util functions for arrays.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const arrayMixin = {
	methods: {
		moveUp(arr, index) {
			console.log(arr)
			return this.moveItem(arr, index, index - 1)
		},
		moveDown(arr, index) {
			return this.moveItem(arr, index, index + 1)
		},
		moveItem(arr, from, to) {
			return arr.splice(to, 0, arr.splice(from, 1)[0]);
		},
	}
}