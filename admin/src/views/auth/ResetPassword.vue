<!-- =====================
	Send Reset Password
	===================== -->
<template>
	<section class="auth reset-password">
		<div class="container">
			<div class="row auth-row">
				<div class="col-12">
					<!-- Logo -->
					<figure class="auth-logo" v-if="getLogo">
						<img :src="getSiteUrl + getLogo">
					</figure><!-- /Col -->
					<div class="auth-card">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text">
								<h2>Forgot password?</h2>
								<p>Enter something about password here</p>
							</div>
							<form class="form form-center">
								<!-- Password -->
								<div class="form-group">
									<input type="password" autocomplete="new-password" placeholder="Password" class="form-input" v-model="password">
								</div>
								<router-link :to="{ name: 'login' }" class="auth-link">Back to login</router-link>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow"
											@click.prevent="doReset">Reset
									</button>
								</div>
							</form>
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
	name: "ResetPassword",
	data: () => ({
		doingAxios: false,
		password: "",
		token: "",
	}),
	mounted: function () {
		this.token = this.$route.params.token;
		this.doVerify();
	},
	methods: {
		doVerify() {
			this.axios.get("/email/verify/" + this.token)
		},
		doReset() {
			this.axios.post("/password/reset", {
				password: this.password,
				token: this.token,
			})
				.then(res => {
					console.log(res);
					this.$router.push({name: 'login'});
				})
				.catch(err => {
					console.log(err);
				})
		},
	},
	computed: {
		getLogo() {
			return this.$store.state.logo;
		}
	}
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

//
// =========================================================================

</style>