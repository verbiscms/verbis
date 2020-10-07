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

import { fieldMixin } from "@/util/fields"

export default {
	name: "FieldTextarea",
	mixins: [fieldMixin],
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
			this.typed = true;
			this.errors = [];
			this.validateMaxLength()
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getLayout() {
			return this.layout;
		},
		getResize() {
			return this.layout.options["resize"] ? '' : "none !important"
		},
		getRows() {
			const rows = this.layout.options['rows']
			return rows ? rows : 8
		},
		value: {
			get() {
				return this.setDefaultValue(this.replacePrependAppend());
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value))
			}
		},
	}
}

</script>
