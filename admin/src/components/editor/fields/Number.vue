<!-- =====================
	Field - Number
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-input form-input-white" type="number"
				v-model="number"
				@keyup="process"
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
	},
	data: () => ({
		number: 0,
		errors: [],
		focused: false,
	}),
	methods: {
		process() {
			this.validate()
			this.emit()
		},
		validate() {
			this.errors = []
			const min = this.getOptions['min'],
				max = this.getOptions['max'];

			if (this.number !== "") {
				if (this.number > max && max !== "") {
					this.errors.push(`The maximum value of the ${this.layout.name} can not exceed ${max}.`)
				}
				if (this.number < min && min !== "") {
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
		emit() {
			this.$emit("input", this.getOptions['prepend'] + this.number + this.getOptions['append'])
		}
	},
	computed: {
		getOptions() {
			return this.layout.options;
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


</style>