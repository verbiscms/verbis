<!-- =====================
	Field - Email
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<div class="form-input-icon">
				<input class="form-input form-input-white" type="text" value="The value"
					v-model="field"
					@keyup="validate"
					:placeholder="getOptions['placeholder']"
					:maxlength="getOptions['maxlength']"
					@focus="focused = true"
					@blur="validateRequired">
				<i class="fal fa-envelope"></i>
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
	name: "FieldEmail",
	mixins: [fieldMixin],
	data: () => ({
		focused: false,
	}),
	mounted() {
		this.setDefaultValue();
	},
	methods: {
		/*
 	     * validate()
		 * Fires when the publish button is clicked.
		 * Checks if the email is valid.
		 */
		validate() {
			this.errors = [];
			// eslint-disable-next-line no-useless-escape
			if (this.getValue !== "" && !this.helpers.validateEmail(this.getValue)) {
				this.errors.push(`Enter a valid email address for the ${this.layout.label} field.`)
			}
			this.validateRequired();
		}
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