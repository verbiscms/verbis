/**
 * fields.js
 * Common util functions for fields.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const fieldMixin = {
	props: {
		layout: Object,
		fields: {
			type: Object,
			default: () => {
				return {}
			}
		},
		index: {
			type: Number,
			default: null,
		},
		parent: {
			type: String,
			default: null,
		},
		parentLayout: {
			type: String,
			default: null,
		},
		errorTrigger: Boolean,
	},
	data: () => ({
		errors: [],
		typed: false,
	}),
	watch: {
		/*
		 * errorTrigger()
		 * Checks if the validation function is defined within the
		 * field and calls.
		 */
		errorTrigger: function() {
			if (typeof this.validate !== "undefined") {
				this.validate();
			}
		}
	},
	computed: {
		/*
		 * getOptions()
		 * Returns the fields layout.
		 */
		getLayout() {
			return this.layout;
		},
		/*
		 * getOptions()
		 * Returns the fields layout options.
		 */
		getOptions() {
			return this.layout.options;
		},
	},
	methods: {
		/*
		 * getFieldObject()
		 * Returns the field object for emitting back to the parent.
		 * Index, Parent & Layout are automatically set to null.
		 */
		getFieldObject(value) {
			return {
				uuid: this.getLayout.uuid,
				value: value,
				name: this.getLayout.name,
				type: this.getLayout.type,
				index: this.index,
				parent: this.parent,
				layout: this.parentLayout,
			};
		},
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
		 * setField()
		 * Set the default value if there is no field data.
		 */
		setDefaultValue() {
			const defaultVal = this.getOptions['default_value'];
			if (this.field === "" && defaultVal !== "" && defaultVal !== undefined) {
				this.field = this.getOptions['default_value'];
			}
		},
		/*
		 * setDefaultValueChoices()
		 * Set the default value for choices if there is no field data.
		 */
		setDefaultValueChoices() {
			if (this.fields === "" && this.getOptions['default_value'] !== "") {
				const opts = this.getOptions['default_value'];
				let defaultVal = ""
				opts.forEach(opt => {
					defaultVal = this.getOptions['choices'][opt]
				});
				if (defaultVal !== "") {
					this.field.value = defaultVal;
					return
				}
			}
			this.field = this.fields;
		},
		/*
		 * replacePrependAppend()
		 * Replace the field value with empty strings.
		 */
		replacePrependAppend() {
			if (!this.fields.value) {
				return "";
			}
			return this.fields.value.toString().replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
		},
		/*
		 * validateRequired()
		 * Return a error message if the options are required & the value is nil.
		 */
		validateRequired() {
			if (this.field === "" && this.getLayout["required"]) {
				this.errors.push(`The ${this.getLayout.label.toLowerCase()} field is required.`);
			}
		},
		/*
		 * validateMaxLength()
		 * Return a error message the max length has been reached.
		 */
		validateMaxLength(length = false) {
			const maxLength = length ? length : this.getOptions['maxlength'];
			if (maxLength !== "" && (this.field.length === maxLength)) {
				this.errors.push(`The maximum length of the ${this.getLayout.label.toLowerCase()} can not exceed ${maxLength} characters.`);
			}
		},
	},
};