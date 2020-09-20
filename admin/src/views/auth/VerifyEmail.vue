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
						<img :src="globalBasePath + getLogo">
					</figure><!-- /Col -->
					<div class="auth-card">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text">
								<h2>Forgot password?</h2>
								<p>Enter your email below to get a verification link sent to reset your password.</p>
							</div>
							<form class="form form-center">
								<!-- Password -->
								<div class="form-group">
									<input type="text" autocomplete="email" placeholder="Email" class="form-input" v-model="email">
								</div>
								<router-link :to="{ name: 'login' }" class="auth-link">Back to login</router-link>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doReset">Reset</button>
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
	name: "SendResetPassword",
	data: () => ({
		doingAxios: false,
		email: "",
	}),
	methods: {
		doReset() {
			console.log(this.email)
			this.axios.post("/password/email", {
				email: this.email
			})
			.then(res => {
				console.log(res);
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