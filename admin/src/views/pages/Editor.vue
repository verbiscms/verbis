<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container editor-auth-container" v-if="!loadingResourceData">
			<!-- =====================
				Header
				===================== -->
			<div class="row">
				<div class="col-12">
					<!-- Header -->
					<header class="header header-with-actions">
						<div class="header-title">
							<h1 v-if="newItem">Add a new {{ resource.friendly_name }}</h1>
							<h1 v-else>Edit {{ resource.friendly_name }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<button class="btn btn-fixed-height btn-margin btn-white">Preview</button>
								<button class="btn btn-fixed-height btn-orange btn-popover" @click.prevent="save">
									<span v-if="newItem">Publish</span>
									<span v-else>Update</span>
									<span class="btn-popover-icon">
										<i class="fal fa-chevron-down"></i>
									</span>
									<div class="popover"></div>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<!-- =====================
					Title & Tabs
					===================== -->
				<div class="col-12 col-desk-9 editor-main-col">
					<Tabs @update="activeTab = $event - 1">
						<template slot="item">Content</template>
						<template slot="item">Meta</template>
						<template slot="item">SEO</template>
						<template slot="item">Code Injection</template>
					</Tabs>
					<!-- Content & Fields -->
					<transition name="trans-fade" mode="out-in">
						<div class="editor-slide" :class="{ 'editor-slide-active' : activeTab === 0}" v-if="fieldLayout.length && activeTab === 0" :key="1">
							<!-- Title -->
							<div class="card">
								<collapse :show="true">
									<template v-slot:header>
										<div class="card-header">
											<h3 class="card-title">Title</h3>
											<div class="card-controls">
												<i class="feather feather-chevron-down"></i>
											</div>
										</div><!-- /Card Header -->
									</template>
									<template v-slot:body>
										<div class="card-body">
											<div class="card-input">
												<input class="form-input form-input-white" type="text" placeholder="Add title" v-model="data.title">
												<!-- Message -->
												<transition name="trans-fade-height">
													<span class="field-message field-message-warning" v-if="errors.title">{{ errors.title }}</span>
												</transition><!-- /Message -->
											</div>
										</div><!-- /Card Body -->
									</template>
								</collapse>
							</div><!-- /Card -->
							<Fields :layout="fieldLayout" :fields.sync="data.fields" :error-trigger="errorTrigger"></Fields>
						</div>
						<!-- Meta Options -->
						<div class="editor-slide" :class="{ 'editor-slide-active' : activeTab === 1}" v-if="activeTab === 1" :key="2">
							<MetaOptions :meta.sync="data.options.meta" url="CHANGE ME"> </MetaOptions>
						</div>
						<!-- Seo Options -->
						<div class="editor-slide" :class="{ 'editor-slide-active' : activeTab === 2}" v-if="activeTab === 2" :key="3">
							<SeoOptions></SeoOptions>
						</div>
						<!-- Code Injection -->
						<div class="editor-slide" :class="{ 'editor-slide-active' : activeTab === 3 }" v-if="activeTab === 3" :key="4">
							<CodeInjection :header="data.codeinjection_head" :footer="data.codeinjection_foot" @update="updateCodeInjection"></CodeInjection>
						</div>
					</transition>
				</div><!-- /Col -->
				<!-- =====================
					Options & Properties
					===================== -->
				<div class="col-12 col-desk-3">
					<div class="editor-sidebar">
						<!-- Options -->
						<h2>Options</h2>
						<div class="editor-sidebar-options">
							<!-- URL -->
							<div class="form-group">
								<label class="form-label" for="options-url">URL</label>
								<input class="form-input form-input-white" type="text" id="options-url" v-model="slug">
								<h2>Tes  {{ computedSlug }}</h2>
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors.slug">{{ errors.slug }}</span>
								</transition><!-- /Message -->
							</div>
							<!-- Status -->
							<div class="form-group">
								<label class="form-label" for="options-status">Status</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="options-status" v-model="data.status">
										<option value="" disabled selected>Select status</option>
										<option value="draft">Draft</option>
										<option value="published">Published</option>
									</select>
								</div>
							</div>
							<!-- Author -->
							<div class="form-group">
								<label class="form-label" for="options-author">Author</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="options-author" v-model="data['author']" @change="getFieldLayout">
										<option value="0" disabled selected>Select author</option>
										<option v-for="user in users" :value="user.id" :key="user.uuid">{{ user.first_name }} {{ user.last_name }}</option>
									</select>
								</div>
							</div>
							<!-- Date -->
							<div class="form-group">
								<label class="form-label">Published Date</label>
								<DatePicker class="date" color="blue" :value="data['published_at']" v-model="data['published_at']"></DatePicker>
							</div>
							{{data['published_at']}}
						</div><!-- /Options -->
						<!-- Properties -->
						<div class="editor-sidebar-properties">
							<h2>Properties</h2>
							<!-- Template -->
							<div class="form-group">
								<label class="form-label" for="properties-template">Template</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="properties-template" v-model="data['page_template']" @change="getFieldLayout">
										<option value="" disabled selected>Select template</option>
										<option v-for="template in templates" :value="template.key" :key="template.key">{{ template.name }}</option>
									</select>
								</div>
							</div>
							<!-- Layout -->
							<div class="form-group">
								<label class="form-label" for="properties-layout">Layout</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="properties-layout" v-model="data['layout']" @change="getFieldLayout">
										<option value="" disabled selected>Select layout</option>
										<option v-for="layout in layouts" :value="layout.key" :key="layout.key">{{ layout.name }}</option>
									</select>
								</div>
							</div>
						</div><!-- /Properties -->
					</div><!-- /Sidebar -->
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
import MetaOptions from "@/components/editor/tabs/Meta";
import SeoOptions from "@/components/editor/tabs/Seo";
import CodeInjection from "@/components/editor/tabs/CodeInjection";
import DatePicker from 'v-calendar/lib/components/date-picker.umd'
import Fields from "@/components/editor/tabs/Fields";
import slugify from "slugify";
import Tabs from "@/components/misc/Tabs";
import Collapse from "@/components/misc/Collapse";

export default {
	name: "Single",
	title: 'Editor',
	components: {
		Tabs,
		Fields,
		Breadcrumbs,
		DatePicker,
		MetaOptions,
		SeoOptions,
		CodeInjection,
		Collapse,
	},
	data: () => ({
		activeTab: 0,
		fieldHeights: [],
		users: [],
		slug: "",
		fieldLayout: [],
		templates: [],
		layouts: [],
		resource: {},
		newItem: false,
		errorTrigger: false,
		errors: {},
		data: {
			"title": "",
			"slug": "/",
			"fields": {},
			"author": 0,
			"status": "",
			"page_template": "",
			"layout": "",
			"options": {},
			"categories": [],
			"codeinjection_head": "",
			"codeinjection_foot": "",
			"published_at": new Date(),
		},
		doingAxios: true,
		loadingResourceData: true,
	}),
	beforeMount() {
		this.setResource()
		this.setNewUpdate();
	},
	mounted() {
		this.getFieldLayout();
		this.getUsers();
		this.getTemplates();
		this.getLayouts();
	},
	methods: {
		/*
		 * getSuccessMessage()
		 * Determine if the page has been created.
		 */
		getSuccessMessage() {
			if (this.$route.query.success) {
				console.log("in")
				this.$noty.success("Successfully created new page.")
			}
		},
		/*
		 * getResourceData()
		 * Get the page data, if none exists return 404.
		 */
		getResourceData() {
			const id = this.$route.params.id;
			this.axios.get(`/posts/${id}`)
				.then(res => {
					const post = res.data.data.post;
					this.data = post;

					// Return 404 if there is no ID
					if (!this.data) {
						this.$router.push({ name : 'not-found' })
					}

					// Compare slugs & set
					if (this.slugify(this.slug) !== this.slugify(this.data.slug)) {
						this.slug = this.data.slug;
					}

					// Set author
					this.data.author = res.data.data.author.id;

					this.setDates()

					this.loadingResourceData = false;
				})
				.catch(err => {
					console.log(err);
					this.$noty.error("Error occurred, please refresh the page.")
				})
		},
		/*
		 * getFieldLayout()
		 * Obtain the field layout, on change.
		 */
		getFieldLayout() {
			this.axios.get("/fields", {
				params: {
					"layout": this.data['layout'],
					"page_template": this.data['page_template'],
					"user_id": this.data['author'],
				}
			})
			.then(res => {
				this.fieldLayout = res.data.data
			})
		},
		/*
		 * getTemplates()
		 * Obtain page templates from API.
		 */
		getTemplates() {
			this.axios.get("/templates")
			.then(res => {
				this.templates = res.data.data.templates
			})
		},
		/*
		 * getLayouts()
		 * Obtain page layouts from API.
		 */
		getLayouts() {
			this.axios.get("/layouts")
			.then(res => {
				this.layouts = res.data.data.layouts
			})
		},
		/*
		 * getUsers()
		 * Obtain users from store, if none, dispatch users action.
		 */
		getUsers() {
			this.$store.dispatch("getUsers")
				.then(users => {
					this.users = users;
				})
				.catch(() => {
					this.$noty.error("Error occured when loading authors, please refresh.")
				})
		},
		/*
		 * setResource()
		 * Set the resource from the query parameter, if none defined,
		 * set default page 'resource'.
		 */
		setResource() {
			const resource = this.getResources[this.$route.query.resource]
			this.resource = resource === undefined ? {
				"name": "page",
				"friendly_name": "Page",
				"slug": "",
				"icon": 'fal fa-file'
			} : resource
		},
		/*
		 * setNewUpdate()
		 * Determine if the page is new or if it already exists.
		 */
		setNewUpdate() {
			const isNew = this.$route.params.id === "new"
			this.newItem = isNew
			if (!isNew) {
				this.getResourceData()
			} else {
				this.loadingResourceData = false;
			}
		},
		setDates() {
			this.data["created_at"] = new Date(this.data['created_at']);
			this.data["updated_at"] = new Date(this.data['updated_at']);
			this.data["published_at"] = new Date(this.data['published_at']);
		},
		/*
		 * save()
		 * Save the new page, check for field validation.
		 */
		save() {
			this.errorTrigger = true;
			this.$nextTick().then(() => {
				if (document.querySelectorAll(".field-cont-error").length === 0) {
					this.data.slug = this.computedSlug;
					if (this.newItem) {
						this.axios.post("/posts", this.data)
							.then(res => {
								// Push to new page if successful
								this.$router.push({
									name: 'editor',
									params: { id : res.data.data.post.id },
									query: { success : "true" }
								})

								// Set defaults
								this.data = res.data.data.post;
								this.data.author = res.data.data.author.id;
								this.newItem = false;
								this.setDates();
								this.getSuccessMessage()
							})
							.catch(err => {
								console.log(err);
								if (err.response.status === 400) {
									this.validate(err.response.data.data.errors)
								}
								this.$noty.error("Error occurred, please refresh the page.")
							})
					} else {
						this.axios.put("/posts/" + this.$route.params.id, this.data)
							.then(() => {
								this.$noty.success("Page updated successfully.")
							})
							.catch(err => {
								console.log(err)
								this.$noty.error("Error occurred, please refresh the page.")
							})
					}
				} else {
					this.$noty.error("Fix the errors before saving the post.")
				}
			})
		},
		validate(errors) {
			this.errors = {}
			errors.forEach(err => {
				this.$set(this.errors, err.key, err.message)
			})
		},
		/*
		 * updateCodeInjection()
		 * Update code injection from component.
		 */
		updateCodeInjection(e) {
			this.data['codeinjection_head'] = e.header;
			this.data['codeinjection_foot'] = e.footer;
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
		/*
		 * getResources()
		 * Get the theme resources from store.
		 */
		getResources() {
			return this.$store.state.theme.resources;
		},
		/*
		 * getBaseSlug()
		 * Get the base slug (resource).
		 */
		getBaseSlug() {
			return  this.resource.name === "page" ? "/" : "/" + this.resource.name + "/";
		},
		/*
		 * computedSlug()
		 * Obtain the computed slug from the input & title.
		 */
		computedSlug: {
			get() {
				return this.getBaseSlug + this.slugify(this.slug ? this.slug : this.data.title);
			},
			set(value) {
				let slug = this.slugify(value)
				this.data.slug = slug;
				return slug;
			}
		}
	}
};
</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">


	.editor {


		// Sidebar
		// =========================================================================

		&-sidebar {
			position: sticky;
			top: 100px;
			background-color: $bg-color;

			&-options {
				margin-bottom: 1.6rem;
			}
		}
	}

</style>