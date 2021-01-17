<!-- =====================
	Field - Range
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append field-prepend-append-range" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-range" type="range"
				v-model="field"
				:placeholder="getOptions['placeholder']"
				:step="getOptions['step']"
				:min="getOptions['min']"
				:max="getOptions['max']"
				@focus="focused = true"
				@blur="focused = false">
			<div class="field-append" v-if="layout.options['append'] !== ''">{{ layout.options['append'] }}</div>
		</div><!-- /Prepend Append -->
		<p class="range-value">Value: {{ field }}</p>
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

import { fieldMixin } from "@/util/fields/fields"

export default {
	name: "FieldRange",
	mixins: [fieldMixin],
	data: () => ({
		focused: false,
	}),
	mounted() {
		this.setDefaultValue();
	},
	computed: {
		/*
		 * field()
		 * Replaces and sets the prepend and append values
		 * Fire's back up to the parent
		 */
		field: {
			get() {
				// TODO: The range field is jumping when stripping default values.
				return this.replacePrependAppend();
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(this.setPrependAppend(value.toString())));
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
			margin-bottom: 0;
			font-weight: 500;
			color: $primary
		}
	}

</style>