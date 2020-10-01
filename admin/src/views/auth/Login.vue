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
							<form class="form form-center" >
								<span class="form-error" v-html="authMessage" :class="{ 'form-error-show' : authMessage }"></span>
								<!-- Email -->
								<div class="form-group">
									<input type="text" autocomplete="email" placeholder="Email" class="form-input" v-model="authInfo.email">
								</div>
								<!-- Password -->
								<div class="form-group">
									<input type="password" autocomplete="password" placeholder="Password" class="form-input" v-model="authInfo.password">
								</div>
								<router-link :to="{ name: 'password-reset' }" class="login-password">Forgot your password?</router-link>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doLogin">Login</button>
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
	}),
	beforeMount() {
		//console.log(this.$store.state.logo)
	},
	methods: {
		doLogin() {
			this.doingAxios = true;
			this.authMessage = '';//{withCredentials: true }
			this.axios.post('/login', {email: this.authInfo.email, password: this.authInfo.password})
				.then(res => {
					console.log(res)
					//this.$store.dispatch("login", res.data.data)
					this.$store.commit('login', res.data.data);
					this.$router.push({ name: 'home' })
				})
				.catch(e => {


					const response = e.response.data;

					console.log(response);
					if (response.data.message === "Validation failed") {
						console.log("in")
					} else {
						this.authMessage = response.data.message
					}

					console.log(e.response.data.data.message);
					this.$store.commit('logout');
					this.doingAxios = false;
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

</style>