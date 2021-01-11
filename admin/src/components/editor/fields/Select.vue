<!-- =====================
	Field - Select
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="form-select-cont form-input">
			<select class="form-select" v-model="field" @blur="validate">
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
	data: () => ({
		focused: false,
	}),
	mounted() {
		this.setDefault();
	},
	methods: {
		setDefault() {
			if (this.getValue === "" && this.getOptions['default_value'] !== "") {
				this.field = this.getOptions['default_value'];
			}
		},
		/*
		* validate()
		* Fires when the publish button is clicked.
		*/
		validate() {
			this.errors = [];
			if (!this.getOptions["allow_null"]) {
				this.validateRequired();
			}
		},
		/*
		 * handleBlur()
		 * Inline validation when user has clicked off the field.
		 * And removes focus class.
		 */
		handleBlur() {
			this.focused = false;
			this.validateRequired();
		},
	},
	computed: {
		getPlaceholder() {
			const placeholder = this.getOptions['placeholder']
			if (!placeholder || placeholder === "") {
				return "Select " + this.getLayout['label'].toLowerCase()
			}
			return placeholder;
		},
		/*
		 * field()
		 * Replaces and sets the prepend and append values
		 * Fire's back up to the parent.
		 */
		field: {
			get() {

				return this.getValue;
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(this.setPrependAppend(value)));
			}
		}
	}
}

</script>