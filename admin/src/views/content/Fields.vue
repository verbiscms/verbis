<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header">
						<div class="header-title">
							<h1>Fields</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<Alert colour="orange">
						<slot>
							<h3>Coming soon!</h3>
							<p>Hold tight, we're working on this..</p>
						</slot>
					</Alert>
				</div>
			</div>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Alert from "../../components/misc/Alert";

export default {
	name: "Fields",
	components: {
		Breadcrumbs,
		Alert
	},
	data: () => ({}),
	methods: {}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

</style>