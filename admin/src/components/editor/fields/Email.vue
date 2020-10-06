<!-- =====================
	Field - Email
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="value"
				@keyup="validate(false)"
				:placeholder="getOptions['placeholder']"
				:maxlength="getOptions['maxlength']"
				@focus="focused = true"
				@blur="handleBlur">
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

import common from './common'

export default {
	name: "FieldEmail",
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
		this.validate(true)
	},
	methods: {
		validate(timeout) {
			this.helpers.debounce(() => {
				this.errors = [];
				// eslint-disable-next-line no-useless-escape
				const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
				if (this.value !== "" && !re.test(String(this.value).toLowerCase())) {
					this.errors.push(`Enter a valid email address for the ${this.layout.label} field.`)
				}
			}, timeout).apply();
		},
		validateRequired() {
			this.errors.push(common.validateRequired(this.value, this.layout))
		},
		handleBlur() {
			this.focused = false;
			this.validateRequired()
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		value: {
			get() {
				return common.replacePrependAppend(this.fields, this.getOptions)
			},
			set(value) {
				this.$emit("update:fields", common.setPrependAppend(value, this.getOptions))
			}
		}
	}
}

</script>