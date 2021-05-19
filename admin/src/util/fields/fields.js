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
		fieldKey: {
			type: String,
			default: "",
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
		/*
		 * getValue()
		 * Returns the fields value.
		 */
		getValue() {
			return this.fields.value ? this.fields.value : "";
		},
		/*
		 * getButtonLabel()
		 * Retrieves the button label for the layout, if there
		 * is none set, 'Add Row' will be returned.
		 */
		getButtonLabel() {
			const label = this.getLayout['options']['button_label'];
			if (!label) {
				return "Add Row";
			}
			return label;
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
				key: this.fieldKey,
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
		 * replacePrependAppend()
		 * Replace the field value with empty strings.
		 */
		replacePrependAppend() {
			return this.getValue.toString().replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
		},
		/*
		 * validateRequired()
		 * Return a error message if the options are required & the value is nil.
		 */
		validateRequired() {
			if (this.field === "" && this.getLayout["required"] === true) {
				this.errors.push(`The ${this.getLayout.label.toLowerCase()} field is required.`);
			}
		},
		/*
		 * validateMaxLength()
		 * Return a error message the max length has been reached.
		 */
		validateMaxLength(length = false) {
			const maxLength = length ? length : this.getOptions['maxlength'];
			if (maxLength !== "" && (this.getValue.length === maxLength)) {
				this.errors.push(`The maximum length of the ${this.getLayout.label.toLowerCase()} can not exceed ${maxLength} characters.`);
			}
		},
	},
};