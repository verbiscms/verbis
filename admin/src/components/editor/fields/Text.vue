<!-- =====================
	Field - Text
	===================== -->
<template>
	<div class="form-group" :class="{ 'form-group-error' : error }">
		{{ layout.options.maxlength }}
		{{ error }}
		<input class="form-input form-input-white" type="text" value="The value" @keyup="process" v-model="text">
		<span class="form-message">Too many charatcers</span>
	</div>
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
		error: false,
	}),
	methods: {
		process() {
			this.validate()
			this.emit()
		},
		validate() {
			const maxLength = this.layout.options.maxlength;
			this.error = maxLength && this.text.length > maxLength
		},
		emit() {
			this.$emit("update:text", this.text)
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>