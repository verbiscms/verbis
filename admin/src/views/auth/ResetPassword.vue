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
								<p>Enter your new password below.</p>
							</div>
							<form class="form form-center">
								<!-- Password -->
								<FormGroup :error="errors['new_password']">
									<input type="password" autocomplete="new-password" placeholder="New password*" class="form-input" v-model="password">
								</FormGroup>
								<!-- Confirm Password -->
								<FormGroup :error="errors['confirm_password']">
									<input type="password" autocomplete="new-password" placeholder="Confirm password*" class="form-input" v-model="confirmPassword">
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
	name: "ResetPassword",
	title: "Reset Password",
	components: {FormGroup},
	data: () => ({
		doingAxios: false,
		password: "",
		confirmPassword: "",
		token: "",
		errors: [],
	}),
	mounted() {
		this.token = this.$route.params.token;
		this.doVerify();
	},
	methods: {
		/*
		 * doVerify()
		 * Verify if the token is valid, if it isn't, return to the login page.
		 */
		doVerify() {
			this.axios.get("/password/verify/" + this.token)
				.catch(() => {
					this.$router.push({ name: 'login' })
				})
		},
		/*
		 * doReset()
		 * Send the new password off to the backend & validate.
		 * If successfully return to login with success message.
		 */
		doReset() {
			this.doingAxios = true;
			this.axios.post("/password/reset", {
				"new_password": this.password,
				"confirm_password": this.confirmPassword,
				token: this.token,
			})
				.then(() => {
					this.$noty.success("Successfully reset password")
					this.$router.push({name: 'login', query: { reset: "true" }});
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before resetting your password.",)
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
		margin-bottom: 1.8rem;
	}
}

</style>
