<!-- =====================
	Field - Email
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<div class="form-input-icon">
				<input class="form-input form-input-white" type="text" value="The value"
					v-model="value"
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
		focused: false,
	}),
	mounted() {
		this.setDefaultValue()
	},
	methods: {
		validate() {
			this.errors = [];
			// eslint-disable-next-line no-useless-escape
			const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
			if (this.value !== "" && !re.test(String(this.value).toLowerCase())) {
				this.errors.push(`Enter a valid email address for the ${this.layout.label} field.`)
			}
			// TODO: Come back to debounce, not working
			this.helpers.debounce(() => {
			}, false).apply();
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
				return this.replacePrependAppend();
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value))
			}
		}
	}
}

</script>