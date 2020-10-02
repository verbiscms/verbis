<!-- =====================
	Errors (Server Down)
	===================== -->
<template>
	<section class="auth">
		<div class="container">
			<div class="row auth-row">
				<div class="col-12">
					<figure class="error-warning">
						<img src="@/assets/images/warning.svg">
					</figure>
					<div class="auth-card">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text">
								<h2>Server Down</h2>
								<p>The Verbis server could not be initialised, please contact the system administrator.</p>
								<button class="btn" @click.prevent="toLogin">Try again</button>
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
		beforeMount() {
			if (this.getSite) {
				this.$router.push({name: 'login'})
			}
		},
		methods: {
			toLogin() {
				this.$store.dispatch("getSiteConfig")
					.then(() => {
						this.$router.push({name: 'login'})
					})
					.catch(() => {
						console.log("Server still down")
					})
			}
		},
		computed: {
			getSite() {
				return this.$store.state.site;
			}
		}
    }
</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

	h2,
	.auth-text {
		margin-bottom: 0;
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