/**
 * fields.js
 * Common util functions for fields.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const choiceMixin = {
	computed: {
		/*
		 * getFormat()
		 * Retrieves the return format from the options.
		 * Options are: key, value or map.
		 */
		getFormat() {
			return "return_format" in this.getOptions ? this.getOptions['return_format'] : "value";
		},
		/*
		 * getChoices()
		 * Retrieves the choices from the options.
		 */
		getChoices() {
			return this.getOptions['choices'];
		}
	},
	methods: {
		/*
		 * setDefault()
		 * Sets default values on mounted.
		 */
		setDefault() {
			if (this.getValue === "" && this.getOptions['default_value'] !== "") {
				this.field = this.getOptions['default_value'];
			}
		},
		/*
		 * getKey()
		 * Gets the value by key in the choices.
		 */
		getKey(value) {
			return Object.keys(this.getChoices).find(key => this.getChoices[key] === value);
		},
		/*
		 * getMultipleFormat()
		 * Determines what format the current return format is,
		 * if there is an error an empty string will be returned
		 * to allow the placeholder to show.
		 */
		getMultipleFormat() {
			// Key format
			if (this.getFormat === "key") {
				return this.getValue in this.getChoices ? this.getChoices[this.getValue] : "";
				// Map format
			} else if (this.getFormat === "map") {
				try {
					JSON.parse(this.getValue)
				} catch(e) {
					return "";
				}
				return JSON.parse(this.getValue).value;
			}
			// Value format
			if (this.getKey(this.getValue) === undefined) {
				return "";
			}
			return this.getValue;
		},
		/*
		 * getMultipleFormat()
		 * Keys are obtained if the return format is key.
		 * Stringify's json if the format is a map.
		 */
		setMultipleFormat(value) {
			// Key format
			if (this.getFormat === "key") {
				value = this.getKey(value);
				// Map format
			} else if (this.getFormat === "map") {
				value = JSON.stringify({
					key: this.getKey(value),
					value: value,
				});
			}

			// Set the key for the API.
			let obj = this.getFieldObject(value);
			obj.key = this.getFormat;
			this.$emit("update:fields", obj);
		}
	},
};