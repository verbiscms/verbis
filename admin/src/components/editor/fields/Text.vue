<!-- =====================
	Field - Text
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<input class="form-input form-input-white" type="text"
				v-model="value"
				@keyup="validate"
				:placeholder="getOptions['placeholder']"
				:maxlength="getOptions['maxlength']"
				@focus="focused = true"
				@blur="handleBlur">
			<div class="field-append" v-if="getOptions['append']">{{ getOptions['append'] }}</div>
		</div><!-- /Prepend Append -->
		<!-- Message -->
		{{ hasError }}
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
	name: "FieldText",
	props: {
		layout: Object,
		fields: {
			type: String,
			default: ''
		},
	},
	data: () => ({
		errors: [],
		focused: false,
		typed: false,
		hasError: false,
	}),
	methods: {
		validate() {
			this.hasError = false;
			this.errors = common.validateMaxLength(this.value, this.layout, this.getOptions)
			this.$emit("update:error", this.hasError)
		},
		validateRequired() {
			this.errors = common.validateRequired(this.value, this.layout)
			this.hasError = true;
		},
		handleBlur() {
			this.focused = false;
			this.validateRequired()
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		value: {
			get() {
				let value = common.replacePrependAppend(this.fields, this.getOptions),
					defaultVal = common.checkDefaultValue(value, this.getOptions);
				if (defaultVal && !this.typed) {
					value = defaultVal
				}
				this.typed = true // eslint-disable-line
				return value;
			},
			set(value) {
				this.$emit("update:fields", common.setPrependAppend(value, this.getOptions))
			}
		}
	}
}

</script>