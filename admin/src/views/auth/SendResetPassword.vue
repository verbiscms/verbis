<!-- =====================
	Send Reset Password
	===================== -->
<template>
	<section class="auth reset-password">
		<div class="container">
			<div class="row auth-row">
				<div class="col-12">
					<!-- Logo -->
					<figure class="auth-logo" v-if="getSite.logo">
						<img :src="getSiteUrl + getSite.logo">
					</figure><!-- /Logo -->
					<div class="auth-card">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text">
								<h2>Forgot password?</h2>
								<p>Enter your email below to get a verification link sent to reset your password.</p>
							</div>
							<form class="form form-center">
								<!-- Password -->
								<FormGroup :error="errors['email']">
									<input type="text" autocomplete="email" placeholder="Email" class="form-input" v-model="email">
								</FormGroup>
								<router-link :to="{ name: 'login' }" class="auth-link">Back to login</router-link>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" :class="{ 'btn-loading' : doingAxios }" @click.prevent="doReset">Reset</button>
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
import FormGroup from "@/components/forms/FormGroup";
export default {
	name: "SendResetPassword",
	title: "Reset Password",
	components: {FormGroup},
	data: () => ({
		doingAxios: false,
		email: "",
		errors: [],
	}),
	methods: {
		doReset() {
			this.doingAxios = true;
			this.axios.post("/password/email", {
				email: this.email
			})
				.then(res => {
					console.log(res.data.message);
					this.$noty.success("A fresh link has been sent to your email.");
					this.$router.push({name: 'login'});
				})
				.catch(err => {
					this.helpers.checkServer(err);

					if (err.response.status === 400) {
						const errors = err.response.data.data.errors;
						if (errors) {
							this.validate(errors);
							this.$noty.error("Fix the errors before resetting your password.",)
							return
						}
						this.$noty.error("Invalid email address.");
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.doingAxios = false;
					}, this.timeoutDelay)
				});
		},
		/*
		 * validate()
		 * Add errors if the post/put failed.
		 */
		validate(errors) {
			this.errors = {};
			errors.forEach(err => {
				this.$set(this.errors, err.key, err.message);
			})
		},
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.auth {

	// Text
	// =========================================================================

	&-text {
		margin-bottom: 1.2rem;
	}
}

</style>