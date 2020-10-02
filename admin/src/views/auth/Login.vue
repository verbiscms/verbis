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
						<img :src="globalBasePath + getSite.logo">
					</figure><!-- /Col -->
					<div class="auth-card">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text">
								<h2>Login</h2>
							</div>
							<form class="form form-center" :class="{ 'form-is-invalid' : authMessage }">
								<span class="form-error" v-html="authMessage" :class="{ 'form-error-show' : authMessage }"></span>
								<!-- Email -->
								<div class="form-group" :class="{ 'form-group-error' : hasError('email') }">
									<input type="text" autocomplete="email" placeholder="Email" class="form-input" v-model="authInfo.email">
									<span class="form-message">{{ getError('email') }}</span>
								</div>
								<!-- Password -->
								<div class="form-group" :class="{ 'form-group-error' : hasError('password') }">
									<input type="password" autocomplete="password" placeholder="Password" class="form-input" v-model="authInfo.password">
									<span class="form-message">{{ getError('password') }}</span>
								</div>
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

export default {
	name: "Login",
	data: () => ({
		doingAxios: false,
		authInfo: {
			email: "",
			password: "",
		},
		authMessage: "",
		errors: [],
	}),
	methods: {
		doLogin() {
			this.doingAxios = true;
			this.authMessage = '';

			this.axios.post('/login', {email: this.authInfo.email, password: this.authInfo.password})
				.then(res => {
					this.$store.commit('login', res.data.data.user, res.data.data.session)
					this.$store.commit('setSession', res.data.data.session);

					if (!this.$store.state.theme.title) {
						this.axios.get("/theme")
							.then(res => {
								this.$store.commit('setTheme', res.data.data);
								this.$router.push({ name: 'home' })
								this.doingAxios = false;
							})
							.catch(err => {
								this.helpers.handleResponse(err);
								this.doingAxios = false;
							})
					}
				})
				.catch(e => {

					const response = e.response.data,
						errors = response.data.errors;

					if (errors) {
						this.errors = errors
					} else {
						this.errors = []
						this.authMessage = response.message
					}

					setTimeout(() => {
						this.doingAxios = false;
					}, 200)
				})
		},
		getError(key) {
			const err = this.errors.find(e => e.key === key)
			return err !== undefined ? err.message : "";
		},
		hasError(key) {
			return this.errors.find(e => e.key === key)
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

</style>