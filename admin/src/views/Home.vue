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
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<Alert colour="green" type="success">
						<slot><strong>Welcome back {{ getUserInfo['first_name'] }}</strong></slot>
					</Alert>
					<!-- Updater -->
					<Alert v-if="getSite['has_update']" colour="orange" type="warning" :cross="false">
						<slot>
							<div class="updater">
								<p><strong>Verbis update available:</strong> Update to version {{ getSite['remote_version'] }}</p>
								<button class="btn" @click="doUpdate" :class="{ 'btn-loading' : updating }">Update</button>
							</div>
						</slot>
					</Alert><!--/ Updater -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Alert from "@/components/misc/Alert";

export default {
	name: "Home",
	components: {
		Alert,
	},
	data: () => ({
		updating: false,
	}),
	methods: {
		/*
		 * doUpdate()
		 * Posts to /update and updates Verbis to the
		 * latest version.
		 */
		doUpdate() {
			this.updating = true;
			this.axios.post("/system/update")
			.then(() => {
				setTimeout(() => {
					this.$store.dispatch("getSiteConfig")
					this.$noty.success("Successfully updated Verbis, restarting system.")
				}, 2000);
			})
			.catch(err => {
				this.helpers.handleResponse(err);
			})
			.finally(() => {
				this.update = false;
			})
		},
	},
	computed: {
		/*
		 * getSite()
		 */
		getSite() {
			return this.$store.state.site;
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Updater
	// ==========================================================================

	.updater {
		width: 100%;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

</style>
