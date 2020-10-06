<!-- =====================
	Field - Textarea
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<textarea class="form-input form-input-white form-textarea"
			v-model="value"
			@keyup="validate"
			@blur="validateRequired"
			:rows="getRows"
			:maxlength="getOptions['maxlength']"
			:style="{ 'resize': getResize }">
		</textarea>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "FieldTextarea",
	props: {
		layout: Object,
		fields: {
			type: String,
			default: ''
		},
	},
	data: () => ({
		errors: [],
	}),
	methods: {
		validate() {
			this.errors = [];
			const maxLength = this.getOptions['maxlength']
			if (maxLength !== "" && this.value.length === maxLength - 1) {
				this.errors.push(`The maximum length of the ${this.layout.name} can not exceed ${this.getOptions["maxlength"]} characters.`)
			}
		},
		validateRequired() {
			if (this.text === "" && this.layout["required"]) {
				this.errors.push(`The ${this.layout.name} field is required.`)
			}
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		value: {
			get() {
				return this.fields.replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
			},
			set(value) {
				this.$emit("update:fields", this.getOptions['prepend'] + value + this.getOptions['append'])
			}
		},
		getResize() {
			return this.layout.options["resize"] ? '' : "none !important"
		},
		getRows() {
			const rows = this.layout.options['rows']
			return rows ? rows : 8
		},
	}
}

</script>
