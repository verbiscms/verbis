/**
 * fields.js
 * Common util functions for fields.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const fieldMixin = {
	data() {
		return {
			typed: false,
		};
	},
	methods: {

		/*
		 * setPrependAppend()
		 * Set the prepend and value options back for setting the field.
		 */
		setPrependAppend(value) {
			if (this.getOptions['prepend']) {
				value = this.getOptions['prepend'] + value;
			}
			if (this.getOptions['append']) {
				value = value + this.getOptions['append']
			}
			return value;
		},

		/*
		 * setDefaultValue()
		 * Set the default value if there is no field data.
		 */
		setDefaultValue(value) {
			if (value === "" && !this.typed && this.getOptions['default_value'] !== "") {
				this.typed = true;
				return this.getOptions['default_value']
			}
			this.typed = true;
			return value
		},

		/*
		 * setDefaultValueChoices()
		 * Set the default value for choices if there is no field data.
		 */
		setDefaultValueChoices() {
			if (this.fields === "" && this.getOptions['default_value'] !== "") {
				const opts = this.getOptions['default_value']
				let defaultVal = ""
				opts.forEach(opt => {
					defaultVal = this.getOptions['choices'][opt]
				});
				return defaultVal
			}
			return this.fields;
		},

		/*
		 * replacePrependAppend()
		 * Replace the field value with empty strings.
		 */
		replacePrependAppend() {
			return this.fields.toString().replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
		},

		/*
		 * validateRequired()
		 * Return a error message if the options are required & the value is nil.
		 */
		validateRequired() {
			if (this.value === "" && this.getLayout["required"]) {
				this.errors.push(`The ${this.getLayout.label.toLowerCase()} field is required.`);
			}
		},

		/*
		 * validateMaxLength()
		 * Return a error message the max length has been reached.
		 */
		validateMaxLength(length = false) {
			const maxLength = length ? length : this.getOptions['maxlength'];
			if (maxLength !== "" && (this.value.length === maxLength)) {
				this.errors.push(`The maximum length of the ${this.getLayout.label.toLowerCase()} can not exceed ${maxLength} characters.`);
			}
		},
	},
};