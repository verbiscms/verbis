<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- Head -->
		<div class="field">
			<h4>Header</h4>
			<div class="form-group">
				<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
				<textarea class="form-input form-input-white" type="text" rows="14" v-model="header" @keyup="emit"></textarea>
			</div>
		</div>
		<!-- Footer -->
		<div class="field">
			<h4>Footer</h4>
			<div class="form-group">
				<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
				<textarea class="form-input form-input-white" type="text" rows="14" v-model="footer" @keyup="emit"></textarea>
			</div>
		</div>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "CodeInjection",
	data: () => ({
		header: "",
		footer: "",
	}),
	methods: {
		emit() {
			this.$emit("update", {
				header: this.header,
				footer: this.footer,
			})
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>