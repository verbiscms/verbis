<!-- =====================
	Home
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header">
						<div class="header-title">
							<h1>Home</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<Alert colour="green" type="success">
						<slot><strong>Welcome back {{ getUserInfo['first_name'] }}</strong></slot>
					</Alert>
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
import Alert from "@/components/misc/Alert";

export default {
	name: "Home",
	components: {
		Alert,
		Breadcrumbs
	},
	data: () => ({}),
	methods: {},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>