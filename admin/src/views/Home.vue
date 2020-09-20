<!-- =====================
	Home
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<h1>Home</h1>
					<Breadcrumbs></Breadcrumbs>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

	import Breadcrumbs from "../components/misc/Breadcrumbs";

	export default {
		name: "Home",
		components: {
			Breadcrumbs
		},
		data: () => ({
		}),
		methods: {
		}
	}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>