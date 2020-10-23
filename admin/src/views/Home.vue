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
					<Alert colour="green">
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