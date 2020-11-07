<!-- =====================
	Categories - Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1 v-if="newItem">New category</h1>
							<h1 v-else>Update {{ data['name']  }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : doingAxios }">
									<span v-if="newItem">Save</span>
									<span v-else>Update</span>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<h6 class="margin">General</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Title & description -->
						<Collapse :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['name'] || errors['description']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Name & description</h4>
										<p>Enter a name and description for the category.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Name -->
									<FormGroup label="Name*" :error="errors['name']">
										<input class="form-input form-input-white" type="text" v-model="data['name']">
									</FormGroup><!-- /Name -->
									<!-- Description -->
									<FormGroup label="Description" :error="errors['description']">
										<input class="form-input form-input-white" type="text" v-model="data['description']">
									</FormGroup><!-- /Description -->
								</div>
							</template>
						</Collapse><!-- /Title & description -->
						<!-- Slug-->
						<Collapse :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['slug']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Slug</h4>
										<p>Enter a slug for the category, by default it will use the name, and will be assigned after the resource, for example: /news/tech</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup class="form-url" label="Slug*" :error="errors['slug']">
										<div class="form-url-cont">
											<input class="form-input form-input-white" type="text" id="options-url" v-model="slug" :disabled="!slugBtn">
											<i class="feather feather-edit" @click="slugBtn = !slugBtn"></i>
										</div>
										<h4>{{computedSlug }}</h4>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Slug-->
						<!-- Resource-->
						<Collapse :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['resource']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Resource</h4>
										<p>Choose a resource the category will be assigned too.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup v-if="data['parent_id'] === '' || data['parent_id'] === null"   label="Resource*" :error="errors['resource']">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model="data['resource']">
												<option disabled selected value=""></option>
												<option v-for="resource in getTheme['resources']" :value="resource['friendly_name']" :key="resource.name">{{ resource['friendly_name'] }}</option>
											</select>
										</div>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Resource-->
						<!-- Parent -->
						<Collapse :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['parent']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Parent</h4>
										<p>Choose a parent category.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Parent">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model="data['parent_id']">
												<option selected value="">No parent</option>
												<option v-for="category in categories" :value="category.id" :key="category.uuid">{{ category['name'] }}</option>
											</select>
										</div>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Parent-->
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

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import slugify from "slugify";

export default {
	name: "Categories",
	components: {
		FormGroup,
		Collapse,
		Breadcrumbs
	},
	data: () => ({
		doingAxios: false,
		categories: [],
		errors: {},
		data: {
			name: "",
			description: "",
			slug: "",
			resource: "",
			parent_id: "",
		},
		newItem: true,
		slug: "",
		slugBtn: false,
	}),
	beforeMount() {
		this.setNewUpdate();
	},
	mounted() {
		this.getCategories();
	},
	methods: {
		/*
		 * getSuccessMessage()
		 * Determine if the page has been created.
		 */
		getSuccessMessage() {
			if (this.$route.query.success) {
				this.$noty.success("Successfully created new category.")
			}
		},
		/*
		 * getCategories()
		 * Obtain the categories.
		 */
		getCategories() {
			this.axios.get('/categories')
				.then(res => {
					this.categories = res.data.data;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * getCategoryById()
		 */
		getCategoryById(id) {
			this.axios.get('/categories/' + id)
				.then(res => {
					const category = res.data.data

					// Return 404 if there is no ID
					if (!this.data) {
						this.$router.push({ name : 'not-found' })
					}

					this.data = category;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * save()
		 * Save or update the new category.
		 */
		save() {
			this.doingAxios = true;

			if (this.data['parent_id'] === "") {
				this.$set(this.data, 'parent_id', null)
			}

			if (this.newItem) {
				this.axios.post('/categories', this.data)
					.then(res => {
						console.log(res);
						this.errors = {};
						this.$noty.success("Successfully created category");

						// Push to new page if successful
						this.$router.push({
							name: 'categories-single',
							params: { id : res.data.data.id },
							query: { success : "true" }
						})

						this.data = res.data.data;
						this.newItem = false;
					})
					.catch(err => {

						this.helpers.checkServer(err);
						if (err.response.status === 400) {
							this.validate(err.response.data.data.errors);
							this.$noty.error("Fix the errors before saving the category.");
							this.setAllHeight();
							return;
						}
						console.log(err.response)
						this.$noty.error(err.response.data.message);
					})
					.finally(() => {
						setTimeout(() => {
							this.doingAxios = false;
						}, this.timeoutDelay);
					});
			} else {
				this.axios.put('/categories/' + this.$route.params.id, this.data)
					.then(() => {
						this.errors = {};
						this.$noty.success("Successfully updated category");
					})
					.catch(err => {
						this.helpers.checkServer(err);
						if (err.response.status === 400) {
							this.validate(err.response.data.data.errors);
							this.$noty.error(this.errorMsg);
							this.setAllHeight();
							return;
						}
						console.log(err.response)
						this.$noty.error(err.response.data.message);
					})
					.finally(() => {
						setTimeout(() => {
							this.doingAxios = false;
						}, this.timeoutDelay);
					});
			}
		},
		/*
		 * setNewUpdate()
		 * Determine if the page is new or if it already exists.
		 */
		setNewUpdate() {
			const isNew = this.$route.params.id === "new"
			this.newItem = isNew
			if (!isNew) {
				this.getCategoryById(this.$route.params.id);
			} else {
				//this.loadingResourceData = false;
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
		/*
		 * slugify()
		 * Slugify's given input.
		 */
		slugify(text) {
			return slugify(text, {
				replacement: '-',    // replace spaces with replacement
				remove: null,        // regex to remove characters
				lower: true          // result in lower case
			})
		},
	},
	computed: {
		computedSlug: {
			get() {

				getTheme['resources']
				return  + "/" + this.slugify(this.data['name']);
			},
			set(value) {
				let slug = this.slugify(value);
				this.data.slug = slug;
				return slug;
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// URL
	// =========================================================================

	@include media-desk {
		.form-url {
			width: 50%;

			input {
				width: 100%;
			}
		}
	}

</style>