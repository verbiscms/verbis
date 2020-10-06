<!-- =====================
	Field - Number
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-input form-input-white" type="number"
				v-model="value"
				@keyup="validate"
				@change="validateRequired"
				:placeholder="getOptions['placeholder']"
				:step="getOptions['step']"
				:min="getOptions['min']"
				:max="getOptions['max']"
				@focus="focused = true"
				@blur="handleBlur">
			<div class="field-append" v-if="layout.options['append'] !== ''">{{ layout.options['append'] }}</div>
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
	name: "FieldNumber",
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
		typed: false,
	}),
	methods: {
		validate() {
			this.errors = []
			const min = this.getOptions['min'],
				max = this.getOptions['max'];
			if (this.value !== "") {
				if (this.value > max && max !== "") {
					this.errors.push(`The maximum value of the ${this.layout.name} can not exceed ${max}.`)
				}
				if (this.value < min && min !== "") {
					this.errors.push(`The minimum value of the ${this.layout.name} can not be below ${min}.`)
				}
			}
		},
		validateRequired() {
			if (this.number === "" && this.layout["required"]) {
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
			return this.layout.options;
		},
		value: {
			get() {
				let value = this.fields.replace(this.getOptions['prepend'], "").replace(this.getOptions['append'], "");
				if (value === "" && !this.typed) { // eslint-disable-line
					this.typed = true; // eslint-disable-line
					value = this.getOptions['default_value']
				}
				this.typed = true // eslint-disable-line
				return value;
			},
			set(value) {
				this.$emit("update:fields", this.getOptions['prepend'] + value + this.getOptions['append'])
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


</style>