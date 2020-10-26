<!-- =====================
	Login
	===================== -->
<template>
	<section class="auth auth-login">
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
							<div class="auth-text auth-text-margin">
								<h2>Login</h2>
							</div>
							<form class="form form-center" :class="{ 'form-is-invalid' : authMessage }">
								<span class="form-error" v-html="authMessage" :class="{ 'form-error-show' : authMessage }"></span>
								<!-- Email -->
								<FormGroup :error="errors['email']">
									<input type="text" autocomplete="email" placeholder="Email" class="form-input" v-model="authInfo.email">
								</FormGroup>
								<FormGroup :error="errors['password']">
									<input type="password" autocomplete="password" placeholder="Password" class="form-input" v-model="authInfo.password">
								</FormGroup>
								<router-link :to="{ name: 'password-reset' }" class="login-password">Forgot your password?</router-link>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doLogin" :class="{ 'btn-loading' : doingAxios }">Login</button>
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
	name: "Login",
	title: "Login",
	components: {FormGroup},
	data: () => ({
		doingAxios: false,
		authInfo: {
			email: "",
			password: "",
		},
		authMessage: "",
		errors: {},
	}),
	methods: {
		doLogin() {
			this.doingAxios = true;
			this.authMessage = '';

			this.axios.post('/login', {email: this.authInfo.email, password: this.authInfo.password})
				.then(res => {
					this.$store.commit('login', res.data.data)
					Promise.all([this.$store.dispatch("getTheme"), this.$store.dispatch("getRoles")])
						.then(() => {
							this.$router.push({ name: 'home' })
						})
						.catch(err => {
							console.log(err);
						});
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						console.log(err.response.data.data)
						this.validate(err.response.data.data.errors);
						return;
					} else {
						this.errors = {};
						this.authMessage = err.response.data.message
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.doingAxios = false;
					}, this.helpers.timeoutDelay)
				})
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
<style scoped lang="scss">

	// Login Page
	// =========================================================================

	.login {

		&-password {
			display: block;
			text-align: right;
			margin-top: -5px;
		}
	}

	.auth-btn-cont {
		margin-top: 1.8rem;
	}

</style>