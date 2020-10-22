<!-- =====================
	User - Create
	===================== -->
<template>
	<section class="create">
		<Modal :show.sync="showCreateModal" class="modal-large">
			<!-- =====================
				Buttons
				===================== -->
			<template slot="button">
				<div class="create-buttons">
					<button class="btn">Generate password</button>
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
					<div class="col-12 col-desk-6">
						<!-- First name -->
						<FormGroup label="First name*" :error="errors['first_name']">
							<input class="form-input form-input-white" type="text" v-model="newUser['first_name']">
						</FormGroup>
						<!-- Email -->
						<FormGroup label="Email*" :error="errors['email']">
							<input class="form-input form-input-white" type="text" v-model="newUser['email']">
						</FormGroup>
						<!-- New password -->
						<FormGroup label="Password*" :error="errors['password']">
							<input class="form-input form-input-white" type="text" v-model="newUser['password']">
						</FormGroup>
					</div><!-- /Col -->
					<div class="col-12 col-desk-6">
						<!-- Last name -->
						<FormGroup label="Last name*" :error="errors['last_name']">
							<input class="form-input form-input-white" type="text" v-model="newUser['last_name']">
						</FormGroup>
						<!-- Role -->
						<FormGroup label="Role*" :error="errors['role_id']">
							<div class="form-select-cont form-input">
								<select class="form-select" id="user-role" v-model="newUser['role']">
									<option v-for="(role, roleIndex) in roles" :selected="roleIndex === 0" :value="role" :key="role.id">{{ role.name }}</option>
								</select>
							</div>
						</FormGroup>
						<!-- Confirm password -->
						<FormGroup label="Confirm password*" :error="errors['confirm_password']">
							<input class="form-input form-input-white" type="text" v-model="newUser['confirm_password']">
						</FormGroup>
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

export default {
	name: "CreateUser",
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
	}),
	mounted() {
		this.getRoles();
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
					this.getUsers();
				})
				.catch(err => {
					console.log(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						const errorMsg =  this.isSelf ? "Fix the errors before saving your profile." : "User updated successfully."
						this.$noty.error(errorMsg);
						return;
					}
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					setTimeout(() => {
						this.doingAxios = false;
					}, 150)
				})
		},
		/*
		 * getRoles()
		 * Obtain all roles from API for use with creating a new user.
		 */
		getRoles() {
			this.axios.get("/roles")
				.then(res => {
					this.roles = res.data.data;
					this.$set(this.newUser, "role", this.roles.find(r => r.name === "Contributor"))
				})
				.catch(err => {
					console.log(err)
					this.$noty.error("Error occurred, please refresh the page.");
				});
		},
		/*
		  * validate()
		 * Add errors if the post failed.
		 */
		validate(errors) {
			console.log(errors)
			this.errors = {};
			errors.forEach(err => {
				this.$set(this.errors, err.key, err.message);
			})
		},
	},
	computed: {
		showCreateModal: {
			get() {
				return this.show;
			},
			set(value) {
				this.$emit("update:show", value)
			}
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
}

</style>