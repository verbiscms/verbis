<!-- =====================
	Install
	===================== -->
<template>
	<section class="auth auth-install">
		<div class="container">
			<div class="row auth-row">
				<div class="col-12">
					<!-- =====================
						Preflight (Step 1)
						===================== -->
					<div class="auth-card" v-if="step === 1">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text auth-text-margin">
								<h2>Install Verbis</h2>
								<p>Welcome to the Verbis installer, please fill out your database credentials below.</p>
							</div>
							<form class="form form-center">
								<span class="form-error" v-html="installMessage" :class="{ 'form-error-show' : installMessage }"></span>
								<!-- Host -->
								<FormGroup :error="errors['db_host']">
									<input type="text" placeholder="Host" class="form-input" v-model="data['db_host']">
								</FormGroup>
								<!-- Port -->
								<FormGroup :error="errors['db_port']">
									<input type="text" placeholder="Port" class="form-input" v-model="data['db_port']">
								</FormGroup>
								<!-- Database -->
								<FormGroup :error="errors['db_database']">
									<input type="text" placeholder="Database" class="form-input" v-model="data['db_database']">
								</FormGroup>
								<!-- Username -->
								<FormGroup :error="errors['db_user']">
									<input type="text" placeholder="User" class="form-input" v-model="data['db_user']">
								</FormGroup>
								<!-- Password -->
								<FormGroup :error="errors['db_password']">
									<input type="password" placeholder="Password" class="form-input" v-model="data['db_password']">
								</FormGroup>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doValidate" :class="{ 'btn-loading' : doingAxios }">Next</button>
								</div>
							</form>
						</div><!-- /Card Cont -->
					</div><!-- /Card -->
					<!-- =====================
						User (Step 2)
						===================== -->
					<div class="auth-card" v-if="step === 2">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text auth-text-margin">
								<h2>Create a user</h2>
								<p>Welcome to the Verbis installer, please fill out your database credentials below.</p>
							</div>
							<form class="form form-center">
								<span class="form-error" v-html="installMessage" :class="{ 'form-error-show' : installMessage }"></span>
								<!-- First Name -->
								<FormGroup :error="errors['user_first_name']">
									<input type="text" placeholder="First name" class="form-input" v-model="data['user_first_name']">
								</FormGroup>
								<!-- Last Name -->
								<FormGroup :error="errors['user_last_name']">
									<input type="text" placeholder="Last name" class="form-input" v-model="data['user_last_name']">
								</FormGroup>
								<!-- Email -->
								<FormGroup :error="errors['user_email']">
									<input type="text" placeholder="Email" class="form-input" v-model="data['user_email']">
								</FormGroup>
								<!-- Password -->
								<FormGroup :error="errors['user_password']">
									<input type="text" placeholder="User" class="form-input" v-model="data['user_password']">
								</FormGroup>
								<!-- Confirm Password -->
								<FormGroup :error="errors['user_confirm_password']">
									<input type="text" placeholder="User" class="form-input" v-model="data['user_confirm_password']">
								</FormGroup>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button v-if="this.step !== 0" type="submit" class="btn btn-arrow btn-arrow-left btn-transparent btn-arrow mr-2" @click.prevent="step--;">Previous</button>
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doValidate" :class="{ 'btn-loading' : doingAxios }">Next</button>
								</div>
							</form>
						</div><!-- /Card Cont -->
					</div><!-- /Card -->
					<!-- =====================
						Site (Step 3)
						===================== -->
					<div class="auth-card" v-if="step === 3">
						<div class="auth-card-cont">
							<!-- Auth Text -->
							<div class="auth-text auth-text-margin">
								<h2>About your site</h2>
								<p>Welcome to the Verbis installer, please fill out your database credentials below.</p>
							</div>
							<form class="form form-center">
								<span class="form-error" v-html="installMessage" :class="{ 'form-error-show' : installMessage }"></span>
								<!-- Host -->
								<FormGroup :error="errors['site_title']">
									<input type="text" placeholder="Site Title" class="form-input" v-model="data['site_title']">
								</FormGroup>
								<!-- Port -->
								<FormGroup :error="errors['site_url']">
									<input type="text" placeholder="Site URL" class="form-input" v-model="data['site_url']">
								</FormGroup>
								<!-- Database -->
								<div class="d-flex align-items-center">
									<div>
										<h4 class="card-title">Allow robots to crawl this site?</h4>
										<p>Enabling private will place a <code>no robots</code> meta tag on the site.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="cache-frontend" checked v-model="data['cache_frontend']" :true-value="true" :false-value="false" />
										<label for="cache-frontend"></label>
									</div>
								</div>
								<!-- Submit -->
								<div class="auth-btn-cont">
									<button type="submit" class="btn btn-arrow btn-transparent btn-arrow" @click.prevent="doInstall" :class="{ 'btn-loading' : doingAxios }">Install</button>
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
	name: "Install",
	title: "Install",
	components: {FormGroup},
	data: () => ({
		doingAxios: false,
		data: {},
		installMessage: "",
		errors: {},
		step: 1,
	}),
	methods: {
		/*
		 * doValidate()
		 * Validates each step of the install.
		 */
		doValidate() {
			this.installMessage = '';
			this.doingAxios = true;

			this.axios.post("/install/validate/" + this.step, this.data)
				.then(res => {
					console.log(res);
					this.errors = [];
					this.$noty.success("Connected to database");
					this.step++;
				})
				.catch(err => {
					if (err.response.status === 400) {
						const errors = err.response.data.data.errors;
						if (!errors) {
							this.installMessage = err.response.data.message;
							return;
						}
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors to install Verbis.",)
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
		 * doInstall()
		 * Installs the application.
		 */
		doInstall() {
			this.doValidate();

			this.installMessage = '';
			this.doingAxios = true;
			this.errors = [];

			this.axios.post("/install", this.data)
				.then(() => {
					this.$noty.success("Successfully installed verbis");
				})
				.catch(err => {
					console.log(err);
					if (err.response.status === 400) {
						const errors = err.response.data.data.errors;
						if (!errors) {
							this.installMessage = err.response.data.message;
							return;
						}
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors to install Verbis.")
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
	}
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>
