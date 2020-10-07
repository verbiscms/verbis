<!-- =====================
	Field - Password
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<div class="form-input-icon">
				<input class="form-input form-input-white" type="password"
					v-model="value"
					@keyup="validate"
					:placeholder="getOptions['placeholder']"
					@focus="focused = true"
					:maxlength="defaultMaxLength"
					@blur="validateRequired">
				<i class="fal fa-lock"></i>
			</div>
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
	name: "FieldPassword",
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
		focused: false,
		defaultMaxLength: 36,
	}),
	methods: {
		validate() {
			this.errors = [];
			this.typed = true;
			this.validateMaxLength(this.defaultMaxLength);
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getLayout() {
			return this.layout;
		},
		value: {
			get() {
				return this.replacePrependAppend();
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value))
			}
		}
	}
}

</script>