/**
 * functions.js
 * Common functions between fields are stored here.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export default {

	/*
	 * replacePrependAppend()
	 * Replace the field value with empty strings.
	 */
	replacePrependAppend(value, options) {
		return value.replace(options['prepend'], "").replace(options['append'], "");
	},

	/*
	 * replacePrependAppend()
	 * Set the prepend and value options back for setting the field.
	 */
	setPrependAppend(value, options) {
		return options['prepend'] + value + options['append']
	},

	checkDefaultValue(value, options) {
		if (value === "" && options['default_value'] !== "") {
			return options['default_value']
		}
		return false
	},


	// checkDefaultPrependAppend(value, options) {
	//
	// },

	/*
	 * validateRequired()
	 * Return a error message if the options are required & the value is nil.
	 */
	validateRequired(value, layout) {
		if (value === "" && layout["required"]) {
			return [(`The ${layout.label.toLowerCase()} field is required.`)]
		}
		return []
	},

	/*
	 * validateMaxLength()
	 * Return a error message the max length has been reached.
	 */
	validateMaxLength(value, layout, options) {
		const maxLength = options['maxlength']
		if (maxLength !== "" && (value.length === maxLength)) {
			return [(`The maximum length of the ${layout.label.toLowerCase()} can not exceed ${options["maxlength"]} characters.`)]
		}
		return []
	}
}