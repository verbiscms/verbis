<!-- =====================
	Errors (404)
	===================== -->
<template>
	<section :class="{ 'auth' : !isAuth }">
		<div :class="{ 'container' : !isAuth, 'auth-container' : isAuth }">
			<div class="row" :class="{ 'auth-row' : !isAuth }">
				<div class="col-12">
					<figure class="error-warning">
						<img src="@/assets/images/warning.svg">
					</figure>
					<div class="auth-card">
						<div class="auth-card-cont">
							<div class="auth-text">
								<h2>404</h2>
								<p>Page not found, click the button below to go back to the last page.</p>
								<div class="auth-btn-cont">
									<button class="btn" @click.prevent="goBack">Go back</button>
								</div>
							</div>
						</div><!-- /Card Cont -->
					</div><!-- /Card -->
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
	name: "ServerDown",
	title: 'Server Down',
	data: () => ({
		initialLoad: true,
		isAuth: false,
	}),
	beforeMount() {
		this.isAuth = this.$store.state.auth;

	},
	methods: {
		goBack() {
			this.$router.back();
		},
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


.row {
	display: flex;
	align-items: center;
	height: 100%;
	min-height: 85vh;
	justify-content: center;
}

.auth-text {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;

	p {
		margin-bottom: 1rem;
	}
}

.error-warning {
	display: flex;
	justify-content: center;
	width: 100%;
	margin-bottom: 2rem;

	img {
		width: 60px;
	}
}

</style>