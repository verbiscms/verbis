<!-- =====================
	Field - Password
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<div class="form-input-icon">
				<input class="form-input form-input-white" type="password"
					v-model="field"
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

import { fieldMixin } from "@/util/fields/fields"

export default {
	name: "FieldPassword",
	mixins: [fieldMixin],
	data: () => ({
		focused: false,
		defaultMaxLength: 36,
	}),
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			this.typed = true;
			this.validateMaxLength(this.defaultMaxLength);
			this.validateRequired();
		},
	},
	computed: {
		/*
		 * field()
		 * Replaces and sets the prepend and append values
		 * Fire's back up to the parent.
		 */
		field: {
			get() {
				return this.replacePrependAppend();
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(this.setPrependAppend(value)));
			}
		}
	}
}

</script>