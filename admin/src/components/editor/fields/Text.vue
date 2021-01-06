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
				@blur="handleBlur"
				@focus="focused = true"
				:placeholder="getOptions['placeholder']"
				:maxlength="getOptions['maxlength']">
			<div class="field-append" v-if="getOptions['append']">{{ getOptions['append'] }}</div>
		</div><!-- /Prepend Append -->
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
	name: "FieldText",
	mixins: [fieldMixin],
	data: () => ({
		focused: false,
	}),
	methods: {
		validate() {
			this.errors = [];
			this.typed = true;
			this.validateMaxLength();
		},
		handleBlur() {
			this.focused = false;
			this.validateRequired();
		},
	},
	mounted() {
		this.setDefaultValue();
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
				return this.replacePrependAppend();
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value));
			}
		}
	}
}

</script>