<!-- =====================
	Field - Email
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="getOptions['prepend']">{{ getOptions['prepend'] }}</div>
			<input class="form-input form-input-white" type="text" value="The value"
				v-model="email"
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
	name: "FieldEmail",
	props: {
		layout: Object,
	},
	data: () => ({
		email: "",
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
			// eslint-disable-next-line no-useless-escape
			const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
			if (this.email !== "" && !re.test(String(this.email).toLowerCase())) {
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
		emit() {
			this.$emit("input", this.getOptions['prepend'] + this.email + this.getOptions['append'])
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		}
	}
}

</script>