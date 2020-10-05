<!-- =====================
	Field - Text
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="text"
				@keyup="process"
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
	name: "FieldText",
	props: {
		layout: Object,
	},
	data: () => ({
		text: "",
		errors: [],
		focused: false,
	}),
	methods: {
		process() {
			this.validate()
			this.emit()
		},
		validate() {
			this.errors = [];
			const maxLength = this.getOptions['maxlength']

			if (maxLength !== "" && this.text.length === maxLength) {
				this.errors.push(`The maximum length of the ${this.layout.name} can not exceed ${this.getOptions["maxlength"]} characters.`)
			}
		},
		validateRequired() {
			console.log(this.text)
			if (this.text === "" && this.layout["required"]) {
				this.errors.push(`The ${this.layout.name} field is required.`)
			}
		},
		handleBlur() {
			this.focused = false;
			this.validateRequired()
		},
		emit() {
			this.$emit("input", this.getOptions['prepend'] + this.text + this.getOptions['append'])
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		}
	}
}

</script>