<!-- =====================
	Field - Select TODO: Handle multiple!
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="form-select-cont form-input">
			<select class="form-select" v-model="value" @blur="validate">
				<option value="" disabled selected>{{ getPlaceholder }}</option>
				<option :value="choice" v-for="choice in getOptions['choices']" :key="choice">{{ choice }}</option>
			</select>
		</div>
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

import {fieldMixin} from "@/util/fields"

export default {
	name: "FieldSelect",
	mixins: [fieldMixin],
	props: {
		layout: Object,
		fields: {
			type: [String, Array],
			default: '',
		},
	},
	data: () => ({
		errors: [],
		focused: false,
	}),
	methods: {
		validate() {
			this.errors = [];
		},
		handleBlur() {
			this.focused = false;
			this.validateRequired();
		},
	},
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		getPlaceholder() {
			const placeholder = this.getOptions['placeholder']
			if (!placeholder || placeholder === "") {
				return "Select " + this.getLayout['label'].toLowerCase()
			}
			return placeholder;
		},
		value: {
			get() {
				return this.setDefaultValueChoices()
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value));
			}
		}
	}
}

</script>