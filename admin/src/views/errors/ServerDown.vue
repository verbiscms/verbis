<!-- =====================
	Errors (Server Down)
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
								<h2>Server Down</h2>
								<p>The Verbis server could not be initialised, please contact the system administrator.</p>
								<div class="auth-btn-cont">
									<button class="btn" @click.prevent="checkServer">Try again</button>
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
			this.checkServer();
		},
		methods: {
			checkServer() {
				this.axios.get("/site").then(res => {
					if (this.$store.state.auth) {
						this.$router.push({ name: "home" });
					} else {
						this.$store.commit("setSite", res.data.data);
						this.$router.push({ name: "login" });
					}
				}).catch(err => {
					console.log(err);
					if (!this.initialLoad) {
						this.$noty.error("Server still down.");
					}
				})
				.finally(() => {
					this.initialLoad = false;
				})
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
		min-height: 90vh;
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