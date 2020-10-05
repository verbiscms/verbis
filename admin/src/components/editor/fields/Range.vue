<!-- =====================
	Field - Range
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append field-prepend-append-range" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-range" type="range"
				v-model="value"
				:placeholder="getOptions['placeholder']"
				:step="getOptions['step']"
				:min="getOptions['min']"
				:max="getOptions['max']"
				@focus="focused = true" @blur="focused = false">
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
	name: "FieldRange",
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

	},
	computed: {
		getOptions() {
			return this.layout.options;
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

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


</style>