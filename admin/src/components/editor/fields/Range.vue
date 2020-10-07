<!-- =====================
	Field - Range
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append field-prepend-append-range" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-range" type="range"
				v-model="value"
				:placeholder="getOptions['placeholder']"
				:step="getOptions['step']"
				:min="getOptions['min']"
				:max="getOptions['max']"
				@focus="focused = true" @blur="focused = false">
			<div class="field-append" v-if="layout.options['append'] !== ''">{{ layout.options['append'] }}</div>
		</div><!-- /Prepend Append -->
		<p class="range-value">Value: {{ value }}</p>
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
	name: "FieldRange",
	mixins: [fieldMixin],
	props: {
		layout: Object,
		fields: {
			type: String,
			default: "0",
		},
	},
	data: () => ({
		errors: [],
		focused: false,
	}),
	mounted() {
		this.setDefaultValue()
	},
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		value: {
			get() {
				// TODO: The range field is jumping when stripping default values.
				return this.replacePrependAppend()
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value));
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.range {

		// Title
		// =========================================================================

		&-value {
			text-align: right;
			margin-top: 10px;
			font-weight: 500;
			color: $primary
		}
	}

</style>