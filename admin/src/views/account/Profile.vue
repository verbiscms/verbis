<!-- =====================
	User - Profile
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1 v-if="isSelf">Edit profile</h1>
							<h1 v-else>Edit User {{ data['first_name'] }} {{ data['last_name'] }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-orange" @click.prevent="save" :class="{ 'btn-loading' : isSavingUser }">
								Update&nbsp;<span class="btn-hide-text-mob">Profile</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-if="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<form class="form" v-else>
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
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>First name*</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['first_name']">
								<input class="form-input form-input-white" type="text" v-model="data['first_name']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Last name -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Last name*</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['last_name']">
								<input class="form-input form-input-white" type="text" v-model="data['last_name']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
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
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Email address*</h4>
							<p>Enter a valid email address, this will be used for signing in to Verbis.</p>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['email']">
								<input class="form-input form-input-white" type="text" v-model="data['email']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Website -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Website</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['website']">
								<input class="form-input form-input-white" type="text" v-model="data['website']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Facebook -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Facebook</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['facebook']">
								<input class="form-input form-input-white" type="text" v-model="data['facebook']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Twitter -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Twitter</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['twitter']">
								<input class="form-input form-input-white" type="text" v-model="data['twitter']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- LinkedIn -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>LinkedIn</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['linked_in']">
								<input class="form-input form-input-white" type="text" v-model="data['linked_in']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Instagram -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Instagram</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['instagram']">
								<input class="form-input form-input-white" type="text" v-model="data['instagram']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Instagram -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Bio</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['biography']">
								<textarea rows="8" class="form-input form-input-white" type="text" v-model="data['biography']"></textarea>
							</FormGroup>
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
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Profile picture</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div v-if="!profilePicture">
								<button class="btn" @click.prevent="showImageModal = true">Add photo</button>
							</div>
							<div v-else>
								<ImageWithActions @choose="showImageModal = true" @remove="removeProfilePicture">
									<img :src="getSiteUrl + profilePicture['url']" />
								</ImageWithActions>
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
				<!-- =====================
					Reset Password
					===================== -->
				<div class="form-row-group profile-last-row" v-if="isSelf">
					<div class="row">
						<div class="col-12">
							<div class="profile-reset-password">
								<h2>Reset password</h2>
								<div>
<!--									<button class="btn btn-orange btn-margin-right">Forgot Password?</button>-->
									<button class="btn btn-orange" @click.prevent="resetPassword" :class="{ 'btn-loading' : isSavingPassword }">Reset</button>
								</div>
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Current Password -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Current password</h4>
							<p>Type in your current password.</p>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['current_password']">
								<input class="form-input form-input-white" type="password" placeholder="Current password" v-model="password['current_password']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Password -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Password</h4>
							<p>Enter a new password, a minimum of 8 alphanumeric characters are required.</p>
							<button class="btn profile-generate-pass" @click.prevent="generatePassword">Generate password</button>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['new_password']">
								<input class="form-input form-input-white" placeholder="New password" :type="isGeneratedPassword ? 'text' : 'password'" v-model="password['new_password']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Confirm Password -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-3">
							<h4>Confirm Password</h4>
							<p>Enter the same password in again.</p>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<FormGroup class="form-group-no-margin" :error="errors['confirm_password']">
								<input class="form-input form-input-white form-input-test" :type="isGeneratedPassword ? 'text' : 'password'" placeholder="Confirm password" v-model="password['confirm_password']">
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
				{{ data }}
			</form>
		</div><!-- /Container -->
		<!-- =====================
			Insert Photo Modal
			===================== -->
		<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
			<template slot="text">
				<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertProfilePicture" :options="false">
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
import ImageWithActions from "@/components/misc/ImageWithActions";
import {userMixin} from "@/util/users";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Home",
	title: 'Profile',
	mixins: [userMixin],
	components: {
		FormGroup,
		ImageWithActions,
		Uploader,
		Modal,
		Breadcrumbs
	},
	data: () => ({
		doingAxios: true,
		data: {
			website: "",
		},
		password: {
			id: false,
			"current_password": "",
			"new_password": "",
			"confirm_password": "",
		},
		media: [],
		errors: [],
		isSelf: false,
		showImageModal: false,
		profilePicture: false,
		timeout: null,
		isSavingUser: false,
		isSavingPassword: false,
		isGeneratedPassword: false,
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
		 * Obtains profile picture if the user is already set.
		 */
		init() {
			this.userId = this.$route.params.id;

			// Return 404 if the user is self
			if (parseFloat(this.userId) === this.getUserInfo.id) {
				this.$router.push({ name : 'not-found' })
			}

			if (!this.userId) {
				this.data = this.getUserInfo;
				this.userId = this.data.id;
				this.isSelf = true;
				this.doingAxios = false;
				Promise.all([this.getMedia(), this.getUser()]).then(() => this.getProfilePicture())
			} else {
				Promise.all([this.getMedia(), this.getUser()]).then(() => this.getProfilePicture())
			}
		},
		/*
		 * save()
		 * Save the updated profile, check for field validation.
		 */
		save() {
			this.isSavingUser = true;

			if (this.errors.length) {
				this.$noty.error("Fix the errors before saving your profile.")
				return
			}
			this.data['password'] = this.newPassword;

			if (this.profilePicture) {
				this.data['profile_picture_id'] = this.profilePicture.id;
			}

			this.$delete(this.data, 'confirm_password');
			this.axios.put("/users/" + this.userId, this.data)
				.then(res => {
					this.errors = [];

					const successMsg = this.isSelf ? "Profile updated successfully." : "User updated successfully."
					this.$noty.success(successMsg);

					// IMPORTANT: Don't commit to the store, if the user isn't the one logged in!
					if (this.isSelf) {
						this.$store.commit("setUser", res.data.data);
					}
				})
				.catch(err => {
					console.log(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						const errorMsg =  this.isSelf ? "Fix the errors before saving your profile." : "User updated successfully."
						this.$noty.error(errorMsg);
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.isSavingUser = false;
					}, this.timeoutDelay);
				});
		},
		/*
		 * resetPassword()
		 */
		resetPassword() {
			this.isSavingPassword = true;
			this.password.id = this.userId;

			this.axios.post("/users/" + this.userId + "/reset-password", this.password)
				.then(() => {
					this.errors = [];
					this.password = {};
					this.isGeneratedPassword = false;
					this.$noty.success("Password updated successfully.");
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before resetting password.");
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.isSavingPassword = true;
					}, this.timeoutDelay);
				})
		},
		// TODO: Implement
		sendResetPassword() {
			this.axios.post("/password/email", {
				email: this.data.email
			})
				.then(res => {
					console.log(res);
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getUser()
		 * Obtains data from API, if the user being edited is not the one
		 * logged in.
		 */
		async getUser() {
			await this.axios.get("/users/" + this.userId)
				.then(res => {
					const user = res.data.data;

					// Return 404 if not found
					if (this.helpers.isEmptyObject(user)) {
						this.$router.push({ name : 'not-found' })
					}

					console.log(user);

					this.data = user;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
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
		/*
		 * insertProfilePicture()
		 * Set the profile picture and show modal, update the store.
		 * Commit to the store if self.
		 */
		insertProfilePicture(e) {
			this.profilePicture = e;
			this.showImageModal = false;
			if (this.isSelf) {
				this.$store.commit("setProfilePicture", e);
			}
		},
		/*
		 * removeProfilePicture()
		 * Set profile picture to false, update the store.
		 */
		removeProfilePicture() {
			this.profilePicture = false;
			this.$store.commit("setProfilePicture", false);
		},
		/*
		 * getMedia()
		 * Return media for filtering profile picture.
		 */
		async getMedia() {
			await this.axios.get("/media?limit=all")
				.then(res => {
					this.media = res.data.data;
				})
		},
		/*
		 * getProfilePicture()
		 */
		getProfilePicture() {
			this.profilePicture = this.media.find(m => m.id === this.data['profile_picture_id']);
		},
		/*
 		* generatePassword()
		 * Generate random hash.
		 */
		generatePassword() {
			const password = this.createPassword()
			this.password['new_password'] = password;
			this.password['confirm_password'] = password;
			this.isGeneratedPassword = true;
		},
	},
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

	&-reset-password {
		display: flex;
		justify-content: space-between;
		align-items: center;

		h2 {
			margin-bottom: 0;
		}
	}

	&-generate-pass {
		margin-bottom: 1rem;
	}

	// Picture
	// =========================================================================

	&-picture {
		width: 300px;
		height: 260px;

		img {
			width: 100%;
			height: 100%;
			object-fit: cover;
			border-radius: 6px;
			//border-radius: 100%;
			box-shadow: 0 0 12px 2px rgba($black, 0.12);
		}
	}

	// Last Row
	// =========================================================================

	&-last-row {
		margin-bottom: 0;
	}

	// Desktop
	// =========================================================================

	@include media-desk {

		&-generate-pass {
			margin-top: 1rem;
			margin-bottom: 0;
		}
	}
}

</style>