<!-- =====================
	Field - Range
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-prepend-append field-prepend-append-range" :class="{ 'field-focused' : focused }">
			<div class="field-prepend" v-if="layout.options['prepend'] !== ''">{{ layout.options['prepend'] }}</div>
			<input class="form-range" type="range"
				v-model="number"
				@input="process"
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
	},
	data: () => ({
		number: 0,
		errors: [],
		focused: false,
	}),
	methods: {
		process() {
			//this.validate()
			this.emit()
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