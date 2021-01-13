<!-- =====================
	Field - Textarea
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<textarea class="form-input form-input-white form-textarea"
			v-model="field"
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

import { fieldMixin } from "@/util/fields/fields"

export default {
	name: "FieldTextarea",
	mixins: [fieldMixin],
	mounted() {
		this.setDefaultValue();
	},
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			this.validateMaxLength()
		},
	},
	computed: {
		/*
		 * getResize()
		 * Returns true if the textarea can be resized.
		 */
		getResize() {
			return this.layout.options["resize"] ? '' : "none !important"
		},
		/*
		 * getRows()
		 * Obtains the amount of textarea rows set in the options.
		 */
		getRows() {
			const rows = this.getOptions['rows'];
			return rows ? rows : 8;
		},
		/*
		 * field()
		 * Replaces and sets the prepend and append values,
		 * Fire's back up to the parent
		 */
		field: {
			get() {
				return this.getValue.replaceAll('</p>', '').replaceAll('<p>', '').replaceAll('<br />', '\n');
			},
			set(value) {
				if (this.getOptions.format === "paragraph") {
					value = value.split('\n').map(str => `<p>${str}</p>`).join('\n');
				} else if (this.getOptions.format === "line_break") {
					value = value.replace(/(?:\r\n|\r|\n)/g, '<br />')
				}
				this.$emit("update:fields", this.getFieldObject(value));
			}
		},
	}
}

</script>
