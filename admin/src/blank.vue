<!-- =====================
	Section
	===================== -->
<template>
	<section class="">
		<div class="container">
			<div class="row">
				<div class="col-12">

				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>
export default {
	name: "Login",
	data: () => ({
		doingAxios: false,
		authInfo: {
			email: '',
			pass: '',
		},
		authMessage: ''
	}),
	methods: {
		dummyMethod() {
		},
	}
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

</style>