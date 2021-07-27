<!-- =====================
	Redirect - Create
	===================== -->
<template>
	<section class="">
		<Modal :show.sync="showModal" class="modal-large create">
			<!-- =====================
				Buttons
				===================== -->
			<template slot="button">
				<div class="create-buttons">
					<button class="btn btn-blue btn-margin-left" :class="{ 'btn-loading' : doingAxios }" @click="save" v-text="redirectUpdate ? 'Update' : 'Create'"></button>
				</div>
			</template>
			<!-- =====================
				Content
				===================== -->
			<template slot="text">
				<div class="row">
					<div class="col-12">
						<h2>{{ redirectUpdate ? 'Update' : 'Create' }} redirect</h2>
					</div><!-- /Col -->
					<div class="col-12">
						<!-- From -->
						<FormGroup label="From*" :error="errors['from_path']">
							<input class="form-input form-input-white" type="text" v-model="redirect['from_path']" tabindex="1" placeholder="Enter a relative URL such as /blog">
						</FormGroup>
						<!-- To -->
						<FormGroup label="To*" :error="errors['to_path']">
							<input class="form-input form-input-white" type="text" v-model="redirect['to_path']" tabindex="2" placeholder="Enter a URL">
						</FormGroup>
						<!-- To -->
						<FormGroup class="form-group-no-margin" label="Code*" :error="errors['code']">
							<div class="form-select-cont form-input">
								<select class="form-select" v-model.number="redirect['code']">
									<option disabled selected value="">Select a code</option>
									<option v-for="code in codes" :key="code.value" :value="code.value">{{ code.text }}</option>
								</select>
							</div>
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
	name: "Redirect",
	props: {
		show: {
			type: Boolean,
			default: false,
		},
		redirectUpdate: {
			type: [Boolean, Object],
		}
	},
	components: {
		FormGroup,
		Modal
	},
	data: () => ({
		doingAxios: false,
		redirect: {},
		codes: [
			{ text: "300 - Multiple Choices", value: 300, },
			{ text: "301 - Moved Permanently", value: 301, },
			{ text: "302 - Found", value: 302, },
			{ text: "303 - See Other", value: 303, },
			{ text: "304 - Not Modified", value: 304, },
			{ text: "307 - Temporary Redirect", value: 307, },
			{ text: "308 - Permanent Redirect", value: 308, },
		],
		errors: {},
	}),
	mounted() {
		this.init(false);
	},
	watch: {
		/*
		 * show()
		 * Watch if the model has been closed/opened &
		 * run init().
		 */
		show: function(val) {
			this.init(val);
		}
	},
	methods: {
		/*
		 * init()
		 * Check if a redirect already exists (for updating).
		 */
		init(open) {
			if (open && this.redirectUpdate) {
				this.redirect = this.redirectUpdate;
				return;
			}
			this.redirect = {
				to_path: "",
				from_path: "",
				code: "",
			};
		},
		/*
		 * save()
		 * Check for errors and update parent.
		 */
		save() {
			this.doingAxios = true;
			this.validate();
			if (this.helpers.isEmptyObject(this.errors)) {
				if (this.redirectUpdate) {
					this.update();
				} else {
					this.create();
				}
			}
		},
		/*
		 * create()
		 * Create a new redirect
		 */
		create() {
			this.axios.post('/redirects', this.redirect)
				.then(() => {
					this.errors = {};
					this.$noty.success("Successfully created redirect");
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before saving the category.");
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.emit();
					setTimeout(() => {
						this.doingAxios = false;
					}, this.timeoutDelay);
				});
		},
		/*
		 * update()
		 * Update existing redirect
		 */
		update() {
			console.log(this.redirect);
			this.axios.put('/redirects/' + this.redirect.id, this.redirect)
				.then(res => {
					console.log(res);
					this.errors = {};
					this.$noty.success("Successfully updated redirect");
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error("Fix the errors before saving the category.");
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.emit();
					setTimeout(() => {
						this.doingAxios = false;
					}, this.timeoutDelay);
				});
		},
		/*
		 * validate()
		 * Add errors if the post failed.
		 */
		validate() {
			this.errors = {};


			// // Validate from
			if (this.redirect['from_path'] === "") {
				this.$set(this.errors, "from_path", "Enter a from URL");
			} //else if (!this.helpers.validateUrl(this.redirect['from'])) {
				//this.$set(this.errors, "from", "Enter a valid URL");
			//}

			// // Validate to
			if (this.redirect['to_path'] === "") {
				this.$set(this.errors, "to", "Enter a to URL");
			} //else if (!this.helpers.validateUrl(this.redirect['to'])) {
				//this.$set(this.errors, "to", "Enter a valid URL");
			//}

			// Validate equals
			if (this.redirect['from_path'] === this.redirect['to_path']) {
				this.$set(this.errors, "to", "Enter a different 'to' URL");
			}

			// Validate code
			const found = this.codes.find(c => c.value === this.redirect['code'])
			if (!found) {
				this.$set(this.errors, "code", "Enter a redirect code");
			}

			setTimeout(() => {
				this.doingAxios = false;
			}, this.timeoutDelay);
		},
		/*
		 * emit()
		 * Fires back upto the parent once redirect has been
		 * created or updated.
		 */
		emit() {
			this.$emit("update", this.redirect);
			this.redirect = {};
		},
	},
	computed: {
		/*
		 * showModal()
		 */
		showModal: {
			get() {
				return this.show;
			},
			set(value) {
				this.$emit("update:show", value)
			}
		},
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
