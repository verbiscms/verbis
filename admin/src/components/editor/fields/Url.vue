<!-- =====================
	Field - Url
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="form-input-icon">
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="field"
				@keyup="validate"
				:placeholder="getOptions['placeholder']"
				@blur="validateRequired">
			<i class="fal fa-globe"></i>
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

import { fieldMixin } from "@/util/fields"

export default {
	name: "FieldEmail",
	mixins: [fieldMixin],
	data: () => ({
		errors: [],
	}),
	mounted() {
		this.setDefaultValue();
	},
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 * Checks if the URL is valid.
		 */
		validate() {
			this.errors = [];
			let pattern = new RegExp('^(https?:\\/\\/)?'+ // Protocol
				'((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|'+ // Domain name
				'((\\d{1,3}\\.){3}\\d{1,3}))'+ // OR ip (v4) address
				'(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*'+ // Port and path
				'(\\?[;&a-z\\d%_.~+=-]*)?'+ // Query string
				'(\\#[-a-z\\d_]*)?$','i'); // Fragment locator
			if (this.getValue !== "" && !pattern.test(this.getValue)) {
				this.errors.push(`Enter a valid url for the ${this.layout.label} field.`)
			}
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