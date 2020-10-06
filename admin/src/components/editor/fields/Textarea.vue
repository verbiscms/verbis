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

import common from "@/components/editor/fields/common";

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
		typed: false,
	}),
	methods: {
		validate() {
			this.errors = common.validateMaxLength(this.value, this.layout, this.getOptions)
		},
		validateRequired() {
			this.errors = common.validateRequired(this.value, this.layout)
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		value: {
			get() {
				let value = this.fields,
					defaultVal = common.checkDefaultValue(value, this.getOptions);
				if (defaultVal && !this.typed) {
					value = defaultVal
				}
				this.typed = true // eslint-disable-line
				return value;
			},
			set(value) {
				this.$emit("update:fields", value)
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
