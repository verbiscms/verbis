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
							<h1 v-if="isSelf">Edit profile</h1>
							<h1 v-else>Edit User {{ data['first_name'] }} {{ data['last_name'] }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								<i class="far fa-check"></i>
								Update Profile
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>First name*</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Last name*</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Password*</h4>
							<p>Enter a new password, a minimum of 8 alphanumeric characters are required.</p>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Email address*</h4>
							<p>Enter a valid email address, this will be used for signing in to Verbis.</p>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Website</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
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
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Facebook</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['facebook']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Twitter -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Twitter</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['twitter']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- LinkedIn -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>LinkedIn</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['linked_in']">
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
				<!-- =====================
					Profile
					===================== -->
				<div class="form-row-group">
					<div class="row">
						<div class="col-12">
							<h2>Profile</h2>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Profile Picture -->
					<div class="row form-row form-row-border form-row-border-top">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Profile picture</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">

							<div class="form-group">
								<button class="btn" @click.prevent="showImageModal = true">Add photo</button>
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
			</form>
		</div><!-- /Container -->
		<!-- =====================
			Insert Photo Modal
			===================== -->
		<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
			<template slot="text">
				<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertProfilePhoto">
					<template slot="close">
						<button class="btn btn-margin-right btn-icon-mob" @click.prevent="showImageModal = false">
							<i class="feather feather-x"></i>
							<span>Close</span>
						</button>
					</template>
				</Uploader>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Modal from "@/components/modals/General";
import Uploader from "@/components/media/Uploader";

export default {
	name: "Home",
	title: 'Profile',
	components: {
		Uploader,
		Modal,
		Breadcrumbs
	},
	data: () => ({
		doingAxios: false,
		saving: false,
		data: {
			website: "",
		},
		errors: [],
		newPassword: "",
		confirmPassword: "",
		userId: -1,
		isSelf: false,
		showImageModal: false,
	}),
	mounted() {
		this.init();
	},
	watch: {
		'$route.params.id': function() {
			this.init();
		},
	},
	methods: {
		/*
		 * init()
		 * Determine if the profile to edit is the user logged in,
		 * or a user that needs to be obtained from the API.
		 */
		init() {
			this.userId = this.$route.params.id;
			if (!this.userId) {
				this.data = this.getUserInfo;
				this.isSelf = true;
			} else {
				this.getUser();
			}
		},
		/*
		 * save()
		 * Save the updated profile, check for field validation.
		 */
		save() {
			this.saving = true;
			if (this.errors.length) {
				this.$noty.error("Fix the errors before saving your profile.")
				return
			}
			this.data['password'] = this.newPassword;
			this.axios.put("/users/" + this.userId, this.data)
				.then(res => {
					this.errors = [];
					this.$noty.success("Profile updated successfully.");

					// IMPORTANT: Don't commit to the store, if the user isn't the one logged in!
					if (this.isSelf) {
						this.$store.commit("setUser", res.data.data);
					}
				})
				.catch(err => {
					console.log(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before saving your profile.");
						return;
					}
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					setTimeout(() => {
						this.saving = false;
					}, 100);
				});
		},
		/*
		 * getUser()
		 * Obtains data from API, if the user being edited is not the one
		 * logged in.
		 */
		getUser() {
			this.axios.get("/users/" + this.userId)
				.then(res => {
					this.data = res.data.data;
				})
				.catch(err => {
					console.log(err);
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
		insertProfilePhoto(e) {
			console.log(e);
		}
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