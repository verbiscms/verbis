<!-- =====================
	Field - Url
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="form-input-icon">
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="value"
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
	props: {
		layout: Object,
		fields: {
			type: String,
			default: ''
		},
	},
	data: () => ({
		errors: [],
	}),
	methods: {
		validate() {
			this.errors = [];
			let pattern = new RegExp('^(https?:\\/\\/)?'+ // protocol
				'((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|'+ // domain name
				'((\\d{1,3}\\.){3}\\d{1,3}))'+ // OR ip (v4) address
				'(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*'+ // port and path
				'(\\?[;&a-z\\d%_.~+=-]*)?'+ // query string
				'(\\#[-a-z\\d_]*)?$','i'); // fragment locator
			if (this.value !== "" && !pattern.test(this.value)) {
				this.errors.push(`Enter a valid url for the ${this.layout.label} field.`)
			}
		}
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
				return this.setDefaultValue(this.replacePrependAppend());
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value))
			}
		}
	}
}

</script>