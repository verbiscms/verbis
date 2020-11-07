<!-- =====================
	User - Create
	===================== -->
<template>
	<section class="">
		<Modal :show.sync="showCreateModal" class="modal-large create">
			<!-- =====================
				Buttons
				===================== -->
			<template slot="button">
				<div class="create-buttons">
					<button class="btn" @click.prevent="generatePassword"><span>Generate password</span></button>
					<button class="btn btn-blue btn-margin-left" :class="{ 'btn-loading' : doingAxios }" @click="create">Create</button>
				</div>
			</template>
			<!-- =====================
				Content
				===================== -->
			<template slot="text">
				<div class="row">
					<div class="col-12">
						<h2>Create user</h2>
					</div><!-- /Col -->
					<div class="col-12">
						<div class="row">
							<div class="col-12 col-desk-6">
								<!-- First name -->
								<FormGroup label="First name*" :error="errors['first_name']">
									<input class="form-input form-input-white" type="text" v-model="newUser['first_name']" tabindex="1">
								</FormGroup>
							</div><!-- /Col -->
							<div class="col-12 col-desk-6">
								<!-- Last name -->
								<FormGroup label="Last name*" :error="errors['last_name']">
									<input class="form-input form-input-white" type="text" v-model="newUser['last_name']" tabindex="2">
								</FormGroup>
							</div><!-- /Col -->
						</div><!-- /Row -->
						<div class="row">
							<div class="col-12 col-desk-6">
								<!-- Email -->
								<FormGroup label="Email*" :error="errors['email']">
									<input class="form-input form-input-white" type="text" v-model="newUser['email']" tabindex="3">
								</FormGroup>
							</div><!-- /Col -->
							<div class="col-12 col-desk-6">
								<!-- Role -->
								<FormGroup label="Role*" :error="errors['role_id']">
									<div class="form-select-cont form-input">
										<select class="form-select" id="user-role" v-model="newUser['role']">
											<option v-for="(role, roleIndex) in getRoles" :selected="roleIndex === 0" :value="role" :key="role.id" tabindex="4">{{ role.name }}</option>
										</select>
									</div>
								</FormGroup>
							</div><!-- /Col -->
						</div><!-- /Row -->
						<div class="row create-last-row">
							<div class="col-12 col-desk-6">
								<!-- New password -->
								<FormGroup label="Password*" :error="errors['password']">
									<input class="form-input form-input-white" :type="isGeneratedPassword ? 'text' : 'password'" v-model="newUser['password']" tabindex="5">
								</FormGroup>
							</div><!-- /Col -->
							<div class="col-12 col-desk-6">
								<!-- Confirm password -->
								<FormGroup label="Confirm password*" class="form-group-no-margin" :error="errors['confirm_password']">
									<input class="form-input form-input-white" :type="isGeneratedPassword ? 'text' : 'password'" v-model="newUser['confirm_password']" tabindex="6">
								</FormGroup>
							</div><!-- /Col -->
						</div><!-- /Row -->
					</div><!-- /Col -->
				</div><!-- /Row -->
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Modal from "@/components/modals/General";
import FormGroup from "@/components/forms/FormGroup";
import {userMixin} from "@/util/users";

export default {
	name: "CreateUser",
	mixins: [userMixin],
	props: {
		show: {
			type: Boolean,
			default: false,
		}
	},
	components: {
		FormGroup,
		Modal
	},
	data: () => ({
		doingAxios: false,
		newUser: {},
		roles: [],
		errors: [],
		isGeneratedPassword: false,
	}),
	mounted() {
		// Set the new users role to Contributor
		this.$set(this.newUser, "role", this.getRoles.find(r => r.name === "Contributor"))
	},
	methods: {
		/*
		 * create()
		 * Save the new user, check for validation.
		 */
		create() {
			this.doingAxios = true;

			this.axios.post("/users", this.newUser)
				.then(() => {
					this.$noty.success("User created successfully");
					this.showCreateModal = false;
					this.newUser = {};
					this.$emit("update", true)
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before saving the user.");
						return;
					}
					if (err.response.status === 409) {
						this.errors = [];
						this.$noty.error(err.response.data.message);
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.doingAxios = false;
					}, this.timeoutDelay);
				})
		},
		/*
		 * validate()
		 * Add errors if the post failed.
		 */
		validate(errors) {
			this.errors = {};
			errors.forEach(err => {
				this.$set(this.errors, err.key, err.message);
			})
		},
		/*
		 * generatePassword()
		 * Generate random hash.
		 */
		generatePassword() {
			const password = this.createPassword();
			this.$set(this.newUser, 'password', password);
			this.$set(this.newUser, 'confirm_password', password);
			this.isGeneratedPassword = true;
		},
	},
	computed: {
		/*
		 * showCreateModal()
		 */
		showCreateModal: {
			get() {
				return this.show;
			},
			set(value) {
				this.$emit("update:show", value)
			}

		},
		/*
		 * getRoles()
		 * Obtain all roles from API for use with creating a new user.
		 * Don't allow the user to create the owner by filter.
		 */
		getRoles() {
			return this.$store.state.roles.filter(r => r.name !== "Owner");
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.create {

	// Props
	// =========================================================================

	h2 {
		margin-bottom: 1rem;
	}

	.modal-container {
		max-width: 700px;
	}

	// Buttons
	// =========================================================================

	&-buttons {
		display: flex;
		justify-content: flex-end;
		width: 100%;
	}

	// Desktop
	// =========================================================================

	@include media-desk {

		&-last-row {

			.form-group {
				margin-bottom: 0;
			}
		}
	}
}

</style>