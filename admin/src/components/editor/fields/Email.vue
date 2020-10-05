<!-- =====================
	Field - Email
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="value"
				@keyup="validate"
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
	methods: {
		validate() {
			this.errors = [];
			// eslint-disable-next-line no-useless-escape
			const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
			if (this.value !== "" && !re.test(String(this.value).toLowerCase())) {
				this.errors.push(`Enter a valid email address for the ${this.layout.name} field.`)
			}

		},
		validateRequired() {
			if (this.email === "" && this.layout["required"]) {
				this.errors.push(`The ${this.layout.name} field is required.`)
			}
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
				return this.fields.replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
			},
			set(value) {
				this.$emit("update:fields", this.getOptions['prepend'] + value + this.getOptions['append'])
			}
		}
	}
}

</script>