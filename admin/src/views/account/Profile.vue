<!-- =====================
	User - Profile
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Edit profile</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : doingAxios }">
								<i class="far fa-check"></i>
								Update Profile
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- =====================
				Info
				===================== -->
			<form class="form">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="form-row-group">
					<div class="row">
						<div class="col-12">
							<h2>Basic options</h2>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- First name -->
					<div class="row form-row form-row-border form-row-border-top">
						<div class="col-12 col-desk-2">
							<h4>First name*</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['first_name']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['first_name']">{{ errors['first_name'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Last name -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>Last name*</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['last_name']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['last_name']">{{ errors['last_name'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Password -->
					<transition name="trans-fade-height">
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>Password*</h4>
							<p>Enter a new password, a minimum of 8 alphanumeric characters are required.</p>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="password" placeholder="New password" v-model="newPassword" @keyup="validatePassword">
								<transition name="trans-fade-height">
									<div v-show="newPassword != ''" class="profile-confirm-password">
										<input class="form-input form-input-white form-input-test" type="password" placeholder="Confirm password" v-model="confirmPassword" @keyup="validatePassword">
									</div>
								</transition>
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['password']">{{ errors['password'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					</transition>
				</div><!-- Form Group -->
				<!-- =====================
					Contact Info
					===================== -->
				<div class="form-row-group">
					<div class="row">
						<div class="col-12">
							<h2>Contact Info</h2>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Email -->
					<div class="row form-row form-row-border form-row-border-top">
						<div class="col-12 col-desk-2">
							<h4>Email address*</h4>
							<p>Enter a valid email address, this will be used for signing in to Verbis.</p>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['email']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['email']">{{ errors['email'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Website -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>Website</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['website']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['website']">{{ errors['website'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Facebook -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>Facebook</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['facebook']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Twitter -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>Twitter</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['twitter']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- LinkedIn -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-2">
							<h4>LinkedIn</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['linked_in']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
				<!-- =====================
					Profile
					===================== -->
				<div class="form-group">
					<div class="row">
						<div class="col-12">
							<h2>Profile</h2>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Profile Picture -->
					<div class="row form-row form-row-border form-row-border-top">
						<div class="col-12 col-desk-2">
							<h4>Profile picture</h4>
						</div>
						<div class="col-12 col-desk-6">
							<div class="form-group">
								<p class="c-primary t-bold mb-0">Coming soon!</p>
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
			</form>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";

export default {
	name: "Home",
	title: 'Profile',
	components: {
		Breadcrumbs
	},
	data: () => ({
		newPassword: "",
		confirmPassword: "",
		data: {
			website: "",
		},
		errors: [],
		doingAxios: false,
	}),
	mounted() {
		this.data = this.getUserInfo;
	},
	methods: {
		/*
		 * save()
		 * Save the updated profile, check for field validation.
		 */
		save() {
			this.doingAxios = true;
			if (this.errors.length) {
				this.$noty.error("Fix the errors before saving your profile.")
				return
			}
			this.data['password'] = this.newPassword;
			this.axios.put("/users/" + this.$store.state.userInfo.id, this.data)
				.then(res => {
					this.errors = [];
					this.$noty.success("Profile updated successfully.");
					this.$store.commit("setUser", res.data.data);
				})
				.catch(err => {
					console.log(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before saving your profile.")
						return
					}
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * validatePassword()
		 * Check if the new nad confirm passwords match
		 */
		validatePassword() {
			if (this.newPassword !== "" && this.newPassword !== this.confirmPassword && this.confirmPassword !== "") {
				this.$set(this.errors, 'password', "The new password & confirm password must match.")
			} else {
				this.$delete(this.errors, 'password')
			}
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
		/*
		 * getUserInfo()
		 * Get the logged in user info from the store.
		 */
		getUserInfo() {
			return this.$store.state.userInfo;
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.profile {

		// Password
		// =========================================================================

		&-confirm-password {
			margin-top: 10px;

		}
	}

</style>