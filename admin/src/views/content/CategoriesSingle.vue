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
								<button v-if="!newItem" class="btn btn-fixed-height btn-orange btn-margin-right" @click.prevent="handleDelete" :class="{ 'btn-loading' : isDoingBulk }">
									<span>Delete</span>
								</button>
								<button class="btn btn-fixed-height btn-orange" @click.prevent="saveHandler" :class="{ 'btn-loading' : doingAxios }">
									<span v-if="newItem">Save</span>
									<span v-else>Update</span>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-if="!doingAxios" class="row trans-fade-in-anim">
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
										<textarea rows="6" type="text" class="form-textarea form-input form-input-white" v-model="data['description']"></textarea>
									</FormGroup><!-- /Description -->
								</div>
							</template>
						</Collapse><!-- /Title & description -->
						<!-- Resource-->
						<Collapse v-if="data['parent_id'] === '' || data['parent_id'] === null" :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['resource']}">
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
									<FormGroup label="Resource*" :error="errors['resource']">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model="data['resource']" @change="getPosts">
												<option disabled selected value="">Select resource</option>
												<option value="pages">Pages</option>
												<option v-for="(resource, resourceKey) in getTheme['resources']" :value="resourceKey" :key="resource.name">{{ resource['friendly_name'] }}</option>
											</select>
										</div>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Resource -->
						<!-- Slug-->
						<Collapse :show="newItem" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['slug']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Slug</h4>
										<p v-if="!getHideCategorySlug">Enter a slug for the category, by default it will use the name, and will be assigned after the resource, for example: /news/tech</p>
										<p v-else>Category slugs for this resource are hidden.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-if="!getHideCategorySlug" v-slot:body>
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
									{{ data['parent_id'] }}
									<FormGroup label="Parent">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model="data['parent_id']">
												<option selected :value="null">No parent</option>
												<option v-for="category in categories" :value="category.id" :key="category.uuid">{{ category['name'] }}</option>
											</select>
										</div>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Parent-->
						<!-- Archive Page -->
						<Collapse :show="newItem" v-if="(data['resource'] !== '' && !newItem) || (data['resource'] !== '' && posts.length)" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['parent']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Archive page</h4>
										<p>Choose an archive page for this category.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body" ref="archive">
									<FormGroup label="Parent">
										<vue-tags-input
											v-model="tag"
											:tags="selectedTags"
											:autocomplete-items="filteredPosts"
											@tags-changed="updateTags"
											add-only-from-autocomplete
											:autocomplete-min-length="0"
											@focus="updateHeight"
											@blur="updateHeight"
											:max-tags="1"
											@max-tags-reached="$noty.warning('Only one archive page can be assigned')"
											placeholder="Add post"
										/>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Archive Page -->
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning">
			<template slot="button">
				<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="deleteCategory">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p>Are you sure want to delete this category?</p>
			</template>
		</Modal>
		<!-- =====================
			Change Archive ID Modal
			===================== -->
		<Modal :show.sync="showWarningModal" class="modal-with-icon modal-with-warning modal-large">
			<template slot="button">
				<div class="category-modal-btns">
					<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="save(false)">Update</button>
					<button v-if="warnings['archive']" class="btn" :class="{ 'btn-loading' : isDeleting }" @click="save(true)">Update & delete old archive</button>
				</div>
			</template>
			<template slot="text">
				<h2>Warnings</h2>
				<ul class="list">
					<li v-if="warnings['slug']">
						Changing the category slug will automatically rename all of the {{ data['resource'] }} slugs from<br /><code>{{ getResourceSlug }}{{ currentSavedSlug }}</code> to <code>{{ computedSlug }}</code>
					</li>
					<li v-if="warnings['archive']">
						Changing the archive page will rename the <span class="t-bold">{{ getPostById(currentSavedArchive)  }}</span> post slug from <code>{{ getPostById(currentSavedArchive)['post']['slug'] }}</code> to <code>/untitled</code>.
						You may want to change the slug or delete the old archive using the `Update & delete old archive` button below.
					</li>
				</ul>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import Modal from "@/components/modals/General";
import VueTagsInput from '@jack_reddico/vue-tags-input';

export default {
	name: "Categories",
	components: {
		Modal,
		FormGroup,
		Collapse,
		Breadcrumbs,
		VueTagsInput,
	},
	data: () => ({
		doingAxios: true,
		categories: [],
		posts: [],
		errors: {},
		data: {
			name: "",
			description: "",
			slug: "",
			resource: "",
			parent_id: null,
			archive_id: null,
		},
		newItem: true,
		slug: "",
		slugBtn: false,
		isDeleting: false,
		isDoingBulk: false,
		showDeleteModal: false,
		showWarningModal: false,
		selectedTags: [],
		tags: [],
		tag: "",
		currentSavedSlug: false,
		currentSavedArchive: false,
		warnings: {},
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
			this.axios.get(`/categories`)
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

					//category['parent_id'] = category['parent_id'] === null ? "" : category['parent_id'];
					category['resource'] = category['resource'] === null ? "" : category['resource'];
					this.currentSavedSlug = category['slug'];
					this.currentSavedArchive = category['archive_id'] === null ? false :  category['archive_id']

					this.data = category;

					this.getPosts();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getPosts()
		 * Obtain posts for the archive page selection.
		 */
		getPosts() {
			this.axios.get(`/posts?limit=all`, {
				params: {
					resource: "pages",
				}
			})
				.then(res => {
					const posts = res.data.data;
					this.posts = posts;
					if (posts.length) {
						this.tags = posts.map(a => {
							return {
								text: a.post.title,
								id: a.post.id
							};
						});
					}
					this.setTags();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				})
		},
		/*
		 * save()
		 * Save or update the new category.
		 */
		save(deleteArchive = false) {
			this.doingAxios = true;

			// Set parent to null if empty string
			// or set the resource if there is a parent association
			if (this.data['parent_id'] !== null) {
				const parent = this.findParent(this.data['parent_id']);
				this.$set(this.data, 'resource', parent['resource']);
			}

			console.log(deleteArchive);

			// Set the computed slug
			this.$set(this.data, 'slug', this.saveSlug);

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
						this.$noty.error(err.response.data.message);
					})
					.finally(() => {
						setTimeout(() => {
							this.doingAxios = false;
							this.showWarningModal = false;
						}, this.timeoutDelay);
						this.getPosts();
					});
			} else {
				this.axios.put('/categories/' + this.$route.params.id, this.data)
					.then(() => {
						this.errors = {};

						if (!deleteArchive) {
							this.$noty.success("Successfully updated category");
						} else {
							this.deleteArchive()
								.then(() => {
									this.$noty.success("Successfully updated category & deleted archive");
								})
								.catch(err => {
									throw err
								});
						}
					})
					.catch(err => {
						console.log(err);
						this.helpers.checkServer(err);
						if (err.response.status === 400) {
							this.validate(err.response.data.data.errors);
							this.$noty.error(this.errorMsg);
							this.setAllHeight();
							return;
						}
						this.$noty.error(err.response.data.message);
					})
					.finally(() => {
						this.showWarningModal = false;
						this.getPosts();
						if (!deleteArchive) {
							setTimeout(() => {
								this.doingAxios = false;
							}, this.timeoutDelay)
						}
					});
			}
		},
		/*
		 * saveHandler()
		 */
		saveHandler() {
			this.warnings = {};

			if (this.data['slug'] !== '' && !this.newItem && (this.saveSlug !== this.currentSavedSlug)) {
				this.$set(this.warnings, "slug", true)
			}

			if (this.data['archive_id'] !== null && this.currentSavedArchive && (this.currentSavedArchive !== this.data['archive_id'])) {
				this.$set(this.warnings, "archive", true)
			}

			if (!this.helpers.isEmptyObject(this.warnings)) {
				this.showWarningModal = true;
				return
			}

			this.save(false);
		},
		/*
		 * findParent()
		 * Find a parent category by given ID.
		 */
		findParent(id) {
			if (this.categories.length) return this.categories.find(c => c.id === id);
		},
		/*
		 * handleDelete()
		 * Show delete modal and spinner.
		 */
		handleDelete() {
			this.isDoingBulk = true;
			this.showDeleteModal = true;
			setTimeout(() => {
				this.isDoingBulk = false;
			}, this.timeoutDelay)
		},
		/*
		 * deleteCategory()
		 */
		deleteCategory() {
			this.axios.delete("/categories/" + this.data['id'])
				.then(() => {
					this.$router.push({name: 'categories', query: { delete : "true" }})
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * deleteArchive()
		 */
		async deleteArchive() {
			return this.axios.delete("/posts/" + this.currentSavedArchive)
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
				this.doingAxios = false;
			}
		},
		/*
		 * setTags()
		 * Push to the selected tags if the archive ID is set.
		 */
		setTags() {
			const archiveId = this.data['archive_id'];
			if (archiveId && this.posts.length) {
				const post = this.getPostById(archiveId);
				if (post) {
					this.selectedTags.push({
						text: post.post.title,
						id: this.data['archive_id'],
					});
				}
			}
		},
		/*
 		 * updateTags()
 		 * Assign the new category ID to the data if it exists.
		 */
		updateTags(category) {
			if (category.length) {
				this.$set(this.data, 'archive_id', parseInt(category[0].id));
			} else {
				this.$set(this.data, 'archive_id', null);
			}
		},
		/*
		 * updateHeight()
		 * Update archive height on focus.
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.archive.closest(".collapse-content"));
			});
		},
		/*
		 * getPostById()
		 */
		getPostById(id) {
			return this.posts.find(p => {
				console.log(id);
				console.log(p.post.id)
				return p.post.id === id
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
	},
	computed: {
		/*
		 * getResourceSlug()
		 * Get resource slug from the data to show in the warnings.
		 */
		getResourceSlug() {
			if (this.data['resource'] === "pages") return "/";
			const resource = this.getTheme['resources'][this.data.resource]
			if ('slug' in resource) {
				return resource.slug + "/";
			}
			return "/";
		},
		/*
		 * getHideCategorySlug()
		 * Obtains the resource from the store and returns tru
		 * if the 'hide_category_slug' is truthy.
		 */
		getHideCategorySlug() {
			const resources = this.getTheme['resources'],
				key = this.data.resource;
			if (key in resources) {
				console.log(resources[key]);
				return resources[key]['hide_category_slug'];
			}
			return false
		},
		/*
		 * filteredPosts()
		 * Retrieve the posts for the select tags.
		 */
		filteredPosts() {
			return this.tags.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
		/*
		 * computedSlug()
		 * Pretty slug for user.
		 */
		computedSlug() {
			if (this.data['parent_id'] !== "") {
				const parent = this.findParent(this.data['parent_id']);
				if (parent) {
					return parent['slug'] + "/" + this.slugify(this.slug ? this.slug : this.data['name']);
				}
			}
			if (this.data['resource'] === "pages") {
				return "/" + this.slugify(this.slug ? this.slug : this.data['name']);
			}
			const resourceSlug = this.data['resource'] === "" ? "/" : "/" + this.data["resource"] + "/";
			return resourceSlug + this.slugify(this.slug ? this.slug : this.data['name']);
		},
		/*
		 * saveSlug()
		 * Slug to be saved in the backend.
		 */
		saveSlug() {
			return this.slugify(this.slug ? this.slug : this.data['name']);
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