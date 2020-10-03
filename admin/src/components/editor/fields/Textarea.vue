<!-- =====================
	Field - Textarea
	===================== -->
<template>
	<textarea class="form-input form-input-white" @keyup="emit" v-model="text" rows="8"></textarea>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "FieldTextarea",
	data: () => ({
		text: "",
	}),
	methods: {
		emit() {
			this.$emit("update", this.text)
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>