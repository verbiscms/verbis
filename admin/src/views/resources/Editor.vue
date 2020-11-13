*<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container editor-auth-container">
			{{ categoryArchive }}
			<!-- =====================
				Header
				===================== -->
			<div class="row">
				<div class="col-12">
					<!-- Header -->
					<header class="header header-with-actions">
						<div class="header-title">
							<h1 v-if="newItem">Add a new {{ resource.friendly_name }}</h1>
							<h1 v-else>Edit {{ resource['singular_name'] ? resource['singular_name'] : resource['friendly_name'] }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions editor-actions">
							<div class="form form-actions">
								<button class="btn btn-icon btn-white btn-margin-right" @click.prevent="sidebarOpen = !sidebarOpen">
									<i class="feather feather-settings"></i>
								</button>
								<button class="btn btn-fixed-height btn-orange btn-popover" :class="{ 'btn-loading' : isSaving }">
									<span class="btn-popover-text" @click.prevent="saveWithStatus('published')">Publish</span>
									<Popover :hover="true" :arrow="true">
										<template slot="button">
											<i class="btn-popover-click feather feather-chevron-down"></i>
										</template>
										<template slot="items">
											<a v-if="!newItem" :href="getSiteUrl + data.slug" class="popover-item popover-item-icon" target="_blank">
												<i class="feather feather-eye"></i>
												<span>Preview</span>
											</a>
											<div class="popover-item popover-item-icon" @click.prevent="saveWithStatus('draft')">
												<i class="feather feather-edit"></i>
												<span>Safe draft</span>
											</div>
											<div class="popover-item popover-item-icon" @click.prevent="saveWithStatus('private')">
												<i class="feather feather-lock"></i>
												<span>Make private</span>
											</div>
											<div class="popover-item popover-item-icon" @click.prevent="saveWithStatus('published')">
												<i class="feather feather-send"></i>
												<span>Publish</span>
											</div>
										</template>
									</Popover>
								</button>
							</div>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<!-- =====================
					Title & Tabs
					===================== -->
				<div class="col-12 col-desk-12 editor-main-col">
					<Tabs @update="activeTab = $event - 1" :default-tab="activeTab">
						<template slot="item">Content</template>
						<template slot="item">Meta</template>
						<template slot="item">SEO</template>
						<template slot="item">Code Injection</template>
						<template slot="item">Insights</template>
					</Tabs>
					<!-- Spinner -->
					<div v-if="doingAxios || loadingLayouts" class="media-spinner spinner-container">
						<div class="spinner spinner-large spinner-grey"></div>
					</div>
					<!-- Content & Fields -->
					<transition v-else name="trans-fade-in-anim" mode="out-in">
						<div v-if="activeTab === 0 && !loadingLayouts" :key="1">
							<!-- Title -->
							<div class="editor-title">
								<FormGroup class="form-group-no-margin" :error="errors['title']">
									<input class="editor-title-text" type="text" placeholder="Add title" v-model="data.title">
								</FormGroup>
								<div class="editor-slug" :class="{ 'editor-slug-disabled' : categoryArchive }">
									<div class="editor-slug-text" @click="slugBtn = !categoryArchive ">
										<i class="feather feather-edit-2" v-if="!categoryArchive"></i>
										<i class="feather feather-slash" v-else></i>
										<p>{{ computedSlug }}</p>
									</div>
									<div v-if="slugBtn" class="editor-slug-form" :class="{ 'editor-slug-form-active' : slugBtn }">
										<input type="text" class="form-input-white" v-model="editSlug">
										<i class="editor-slug-save feather feather-save" @click.prevent="saveSlug"></i>
										<i class="editor-slug-close feather feather-x-circle" @click="closeSlug"></i>
									</div>
									<div v-if="categoryArchive" class="badge badge-orange">Category archive</div>
								</div>
							</div>
							<Fields :layout="fieldLayout" :fields.sync="data.fields" :error-trigger="errorTrigger"></Fields>
						</div>
						<!-- Meta Options -->
						<MetaOptions v-if="activeTab === 1" :key="2" :meta.sync="data.options.meta" :url="data.slug"></MetaOptions>
						<!-- Seo Options -->
						<SeoOptions v-if="activeTab === 2" :key="3"></SeoOptions>
						<!-- Code Injection -->
						<CodeInjection v-if="activeTab === 3" :key="4" :header="data.codeinjection_head" :footer="data.codeinjection_foot" @update="updateCodeInjection"></CodeInjection>
						<!-- Seo Options -->
						<Insights v-if="activeTab === 4" :key="4"></Insights>
					</transition>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Sidebar
			===================== -->
		<aside class="editor-sidebar" :class="{ 'editor-sidebar-active' : sidebarOpen }">
			<div class="editor-sidebar-header">
				<h3>Page options</h3>
				<i class="feather feather-x" @click.prevent="sidebarOpen = false"></i>
			</div>
			<div class="editor-sidebar-body">
				<div class="editor-sidebar-cont">
					<h6 class="margin">Publishing</h6>
					<!-- Status -->
					<FormGroup label="Status">
						<div class="form-select-cont form-input">
							<select class="form-select" id="options-status" v-model="data.status">
								<option value="draft">Draft</option>
								<option value="published">Published</option>
								<option value="private">Private</option>
							</select>
						</div>
					</FormGroup><!-- /Status -->
					<!-- Author -->
				</div>
				<div class="editor-sidebar-cont">
					<h6 class="margin">Content</h6>
					<!-- Author -->
					<FormGroup label="Author">
						<div class="form-select-cont form-input">
							<select class="form-select" id="options-author" v-model="data['author']" @change="getFieldLayout">
								<option value="0" disabled selected>Select author</option>
								<option v-for="user in users" :value="user.id" :key="user.uuid">{{ user.first_name }} {{ user.last_name }}</option>
							</select>
						</div>
					</FormGroup><!-- /Author -->
					<!-- Categories -->
					<FormGroup v-if="!categoryArchive" label="Category">
						<div v-if="categories.length" class="form-select-cont form-input">
							<select class="form-select" id="options-categories" v-model="data['category']" @change="getFieldLayout">
								<option :value="null" selected>No category</option>
								<option v-for="category in categories" :value="category.id" :key="category.uuid">{{ category.name }}</option>
							</select>
						</div>
						<div v-else class="editor-sidebar-category">
							<p>No categories available, click below to create one</p>
							<router-link to="/categories/new" class="btn btn-white btn-small btn-block">Create category</router-link>
						</div>
					</FormGroup><!-- /Categories -->
				</div>
				<div class="editor-sidebar-cont">
					<h6 class="margin">Properties</h6>
					<!-- Page Template -->
					<FormGroup label="Page template">
						<div class="form-select-cont form-input">
							<select class="form-select" id="properties-template" v-model="data['page_template']" @change="getFieldLayout">
								<option value="" disabled selected>Select template</option>
								<option v-for="template in templates" :value="template.key" :key="template.key">{{ template.name }}</option>
							</select>
						</div>
					</FormGroup><!-- /Page Template -->
					<!-- Layout -->
					<FormGroup label="Layout">
						<div class="form-select-cont form-input">
							<select class="form-select" id="properties-layout" v-model="data['layout']" @change="getFieldLayout">
								<option v-for="(layout, layoutKey) in layouts" :value="layout.key" :key="layout.key" :selected="layoutKey === 1">{{ layout.name }}</option>
							</select>
						</div>
					</FormGroup><!-- /Layout -->
					<!-- Published Date -->
					<FormGroup label="Published date">
						<DatePicker class="date" color="blue" :value="data['published_at']" v-model="data['published_at']"></DatePicker>
					</FormGroup><!-- /Published Date -->
				</div>
			</div>
		</aside>
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
import Insights from "@/components/editor/tabs/Insights";
import DatePicker from 'v-calendar/lib/components/date-picker.umd'
import Fields from "@/components/editor/tabs/Fields";
// import slugify from "slugify";
import Tabs from "@/components/misc/Tabs";
import Popover from "@/components/misc/Popover";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Single",
	title: 'Editor',
	components: {
		FormGroup,
		Tabs,
		Fields,
		Breadcrumbs,
		DatePicker,
		MetaOptions,
		SeoOptions,
		CodeInjection,
		Popover,
		Insights,
	},
	data: () => ({
		doingAxios: true,
		loadingLayouts: true,
		activeTab: 0,
		users: [],
		isCustomSlug: false,
		editSlug: "",
		slugBtn: false,
		fieldLayout: [],
		templates: [],
		layouts: [],
		resource: {},
		categories: [],
		newItem: false,
		categoryArchive: false,
		errorTrigger: false,
		errors: {},
		data: {
			"title": "",
			"slug": "/",
			"fields": {},
			"archive_id": "",
			"author": 0,
			"status": "draft",
			"page_template": "",
			"layout": "",
			"options": {},
			"category": null,
			"codeinjection_head": "",
			"codeinjection_foot": "",
			"published_at": new Date(),
		},
		defaultLayout: `{"uuid":"6a4d7442-1020-490f-a3e2-436f9135bc24","title":"Default Options","fields":[{"uuid":"39ca0ea0-c911-4eaa-b6e0-67dfd99e1225","label":"RichText","name":"content","type":"richtext","instructions":"Add content to the page.","required":true,"conditional_logic":null,"wrapper":{"width":100},"options":{"default_value":"","tabs":"all","toolbar":"full","media_upload":1}}]}`,
		isSaving: false,
		sidebarOpen: false,
	}),
	beforeMount() {
		this.setNewUpdate();
	},
	mounted() {
		this.init();
	},
	watch: {
		getCategory: function () {
			if (!this.doingAxios) {
				this.computedSlug = this.getBaseSlug + this.slugify(this.editSlug);
			}
		}
	},
	methods: {
		/*
		 * init()
		 */
		init() {
			this.getSuccessMessage();
			this.setResource();
			this.setTab();
			if (this.newItem) {
				Promise.all([this.getUsers(), this.getCategories(), this.getLayouts(), this.getTemplates()])
					.then(() => {
						this.doingAxios = false;
						this.loadingLayouts = false;
						if (this.layouts.length >= 2) {
							this.$set(this.data, 'layout', this.layouts[1].key);
						}
						this.getFieldLayout();
					})
			} else {
				Promise.all([this.getUsers(), this.getLayouts(), this.getTemplates()])
					.then(() => {
						this.doingAxios = false;
						this.loadingLayouts = false;
					})
			}
		},
		/*
		 * getSuccessMessage()
		 * Determine if the page has been created.
		 */
		getSuccessMessage() {
			if (this.$route.query.success) {
				this.$noty.success("Successfully created new page.")
			}
		},
		/*
		 * setDefaultLayout()
		 * If there are no layouts, add a richtext to the default field layout.
		 */
		setDefaultLayout() {
			if (!this.fieldLayout.length) {
				this.fieldLayout.push(JSON.parse(this.defaultLayout));
			}
		},
		/*
		 * setTab()
		 * Determine if there is a tab query paramater and set the active tab.
		 */
		setTab() {
			const tab = this.$route.query.tab
			if (tab) {
				switch (tab) {
					case "meta" : {
						this.activeTab = 1;
						break;
					}
					case "seo" : {
						this.activeTab = 2;
						break;
					}
					case "code-injection" : {
						this.activeTab = 3;
						break;
					}
					case "insights" : {
						this.activeTab = 4;
						break;
					}
				}
			}
		},
		/*
		 * getResourceData()
		 * Get the page data, if none exists return 404.
		 */
		async getResourceData() {
			const id = this.$route.params.id;
			await this.axios.get(`/posts/${id}`)
				.then(res => {
					const post = res.data.data.post;
					this.data = post;

					// Return 404 if there is no ID
					if (!this.data) {
						this.$router.push({ name : 'not-found' })
					}

					// Compare slugs & set
					if (this.data.slug !== this.getBaseSlug + this.slugify(this.data['title'])) this.isCustomSlug = true;

					// Set author
					this.data.author = res.data.data.author.id;

					// Set field layouts
					this.fieldLayout = res.data.data.layout;
					this.setDefaultLayout();

					// Set category
					const category = res.data.data.category;
					if (category) {
						this.$set(this.data, 'category', category.id)
					} else {
						this.$set(this.data, 'category', null)
					}

					// Set date format
					this.setDates();

					this.getCategories();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getFieldLayout()
		 * Obtain the field layout, on change.
		 */
		async getFieldLayout() {
			this.loadingLayouts = true;
			await this.axios.get("/fields", {
				params: {
					"layout": this.data['layout'],
					"resource": this.resource.name,
					"page_template": this.data['page_template'],
					"user_id": this.data['author'],
				}
			})
			.then(res => {
				this.fieldLayout = res.data.data;
				this.setDefaultLayout();
			})
			.finally(() => {
				setTimeout(() => {
					this.loadingLayouts = false;
				}, this.timeoutDelay);
			});
		},
		/*
		 * getCategories()
		 * Obtain the categories.
		 */
		async getCategories() {
			await this.axios.get(`/categories?filter={"resource":[{"operator":"=", "value": "${this.resource['name']}"}]}`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
				.then(res => {
					const categories = res.data.data;
					this.categories = categories;
					if (!this.newItem) {
						categories.forEach(c => {
							if (c.archive_id === this.data.id) this.categoryArchive = true;
						});
					}
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getTemplates()
		 * Obtain page templates from API.
		 */
		async getTemplates() {
			await this.axios.get("/templates")
			.then(res => {
				this.templates = res.data.data.templates
			})
			.catch(err => {
				this.helpers.handleResponse(err);
			})
		},
		/*
		 * getLayouts()
		 * Obtain page layouts from API.
		 */
		async getLayouts() {
			await this.axios.get("/layouts")
			.then(res => {
				this.layouts = res.data.data.layouts
			})
			.catch(err => {
				this.helpers.handleResponse(err);
			})
		},
		/*
		 * getUsers()
		 * Obtain users from store, if none, dispatch users action.
		 */
		async getUsers() {
			await this.$store.dispatch("getUsers")
			.then(users => {
				this.users = users;
			})
			.catch(err => {
				this.helpers.handleResponse(err);
			})
		},
		/*
		 * setResource()
		 * Set the resource from the query parameter, if none defined,
		 * set default page 'resource'.
		 */
		async setResource() {
			const resource = this.getTheme['resources'][this.$route.query.resource];
			this.resource = resource === undefined ? {
				"name": "pages",
				"friendly_name": "Page",
				"singular_name": "Page",
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
				this.getResourceData();
			}
		},
		/*
		 * setDates()
		 * Set the dates for moment when getting back from API.
		 */
		setDates() {
			this.data["created_at"] = new Date(this.data['created_at']);
			this.data["updated_at"] = new Date(this.data['updated_at']);
			if (this.data.published_at !== null) {
				this.data["published_at"] = new Date(this.data['published_at']);
			} else {
				this.data["published_at"] = new Date();
			}
		},
		/*
		 * save()
		 * Save the new page, check for field validation.
		 */
		save() {
			this.isSaving = true;
			this.errorTrigger = true;
			this.$nextTick().then(() => {
				if (document.querySelectorAll(".field-cont-error").length === 0) {
					this.$set(this.data, 'slug', this.computedSlug);
					if (this.resource.name !== "page") {
						this.data.resource = this.resource.name;
					}

					if (this.newItem) {
						this.axios.post("/posts", this.data)
							.then(res => {
								// Push to new page if successful
								this.$router.push({
									name: 'editor',
									params: { id : res.data.data.post.id },
									query: { success : "true", resource : res.data.data.post.resource }
								})
							})
							.catch(err => {
								this.helpers.checkServer(err);
								if (err.response.status === 400) {
									const msg = err.response.data.message,
										errors = err.response.data.data.errors;
									if (msg && !errors) this.$noty.error(msg);
									this.validate(errors);
									this.$noty.error("Fix the errors before saving the " + this.resource['singular_name'] + ".");
									return;
								}
								this.helpers.handleResponse(err);
							})
							.finally(() => {
								setTimeout(() => {
									this.isSaving = false;
								}, this.timeoutDelay);
							})
					} else {
						this.axios.put("/posts/" + this.$route.params.id, this.data)
							.then(() => {
								this.$noty.success("Page updated successfully.")
							})
							.catch(err => {
								if (err.response.status === 400) {
									const msg = err.response.data.message,
										errors = err.response.data.data.errors;
									if (msg) {
										this.$noty.error(msg);
										return;
									}
									this.validate(errors);
									this.$noty.error("Fix the errors before saving the " + this.resource['singular_name'] + ".");
									return;
								}
								this.helpers.handleResponse(err);
							})
							.finally(() => {
								setTimeout(() => {
									this.isSaving = false;
								}, this.timeoutDelay);
							})
					}
				} else {
					this.$noty.error("Fix the errors before saving the post.")
					setTimeout(() => {
						this.isSaving = false;
					}, this.timeoutDelay);
				}
			})
		},
		/*
 		 * save()
		 * Save the new page, with status (for popover).
		 */
		saveWithStatus(status) {
			this.$set(this.data, 'status', status);
			this.save();
		},
		/*
		 * validate()
		 * Add errors if the post/put failed.
		 */
		validate(errors) {
			if (errors && errors.length) {
				this.errors = {};
				errors.forEach(err => {
					this.$set(this.errors, err.key, err.message);
				});
			}
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
		 * resolveCategorySlug()
		 * Find the category by ID and work the way up the tree of categories
		 * until the parent ID is undefined, reverse the slug array and
		 * return the nice slug.
		 */
		resolveCategorySlug() {
			let categorySlugs = [];

			if (this.data['category']) {
				let category = this.categories.find(c => c.id === this.data['category']);
				categorySlugs.push(category['slug']);

				while (category['parent_id'] !== null) {
					category = this.categories.find(c => c.id === category['parent_id']);
					categorySlugs.push(category['slug']);
				}
			}
			categorySlugs = categorySlugs.reverse();

			let slug = '';
			categorySlugs.forEach(c => {
				slug += c + "/";
			});

			return slug;
		},
		/*
		 * saveSlug()
		 * Slugify the new edited slug and set the custom slug
		 * to true, close the slug editing area.
		 */
		saveSlug() {
			if (this.editSlug === "") {
				this.closeSlug();
				return;
			}
			else {
				const newSlug = this.getBaseSlug + this.slugify(this.editSlug);
				this.computedSlug = newSlug;
				this.slugBtn = false;
				this.isCustomSlug = true;
				this.editSlug = "";
			}
		},
		/*
		 * closeSlug()
		 * Handler for closing the slug edit button,
		 * restore default values.
		 */
		closeSlug() {
			this.isCustomSlug = false;
			this.editSlug = "";
			this.slugBtn = false;
		}
	},
	computed: {
		/*
		 * getBaseSlug()
		 * Get the base slug (resource).
		 */
		getBaseSlug() {
			return this.resource.name === "page" ? "/" + this.resolveCategorySlug() : "/" + this.resource.name + "/" + this.resolveCategorySlug();
		},
		/*
		 * getCategory()
		 * Get the category from the data.
		 */
		getCategory() {
			return this.data.category;
		},
		/*
		 * computedSlug()
		 * If the slug is custom, return the slug that is stored in the data.
		 * Otherwise get the base slug and slugify the title or the slug
		 * that has been edited.
		 */
		computedSlug: {
			get() {
				if (this.isCustomSlug) return this.data.slug;
				if (this.editSlug !== "") return this.getBaseSlug + this.slugify(this.data['title']);
				return this.getBaseSlug + this.slugify(this.editSlug ? this.editSlug : this.data['title']);
			},
			set(value) {
				this.$set(this.data, 'slug', value)
			}
		}
	}
};
</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

	$editor-side-input-height: 50px;

	.editor {

		// Actions
		// =========================================================================

		&-actions {

			@include media-mob-down {

				i {
					font-size: 1.2rem;
				}
			}
		}

		// Buttons
		// =========================================================================

		&-save-btn {
			display: none;
		}

		// Title
		// =========================================================================

		&-title {
			margin: 3rem 0;

			&-text {
				background-color: transparent;
				outline: none;
				border: none;
				width: 100%;
				font-size: 2rem;
				color: $black;
				font-weight: 600;
				overflow-y: scroll;
			}
		}

		// Slug
		// =========================================================================

		&-slug {
			display: inline-flex;
			align-items: center;
			flex-wrap: wrap;
			width: 100%;
			margin-top: 10px;
			cursor: pointer;
			min-height: 25px;

			i {
				color: $grey;
				margin-right: 6px;
				font-size: 16px;
			}

			p,
			i {
				transition: 200ms ease color;
			}

			.badge {
				margin-left: 10px;
			}

			&-text {
				display: inline-flex;
				align-items: center;
				min-width: 100%;

				p {
					color: $grey;
					margin: 0;
					line-height: 1.3;
				}
			}

			&:not(&-disabled) &-text:hover  {

				p,
				i {
					color: $primary;
				}
			}

			&-form {
				display: flex;
				align-items: center;
				opacity: 0;
				transition: opacity 200ms ease;
				margin-top: 10px;

				input {
					border: 1px solid $grey-light;
					color: $secondary;
					outline: none;
					padding: 4px 6px;
					margin-right: 6px;
					width: auto;
					font-size: 0.8rem;
					border-radius: 4px;
				}

				&-active {
					opacity: 1;
				}
			}

			&-disabled {

				i {
					color: $orange;
				}
			}

			&-save:hover {
				color: $green;
			}

			&-close:hover {
				color: $orange;
			}

			@include media-tab {
				flex-wrap: nowrap;

				&-text {
					min-width: 0;
				}

				&-form {
					margin-top: 0;

					input {
						margin: 0 6px;
					}
				}
			}
		}

		// Sidebar
		// =========================================================================

		&-sidebar {
			position: fixed;
			top: 0;
			right: 0;
			height: 100%;
			width: 250px;
			background-color: $bg-color;
			z-index: 999999;
			transform: translateX(100%);
			transition: transform .4s cubic-bezier(.1,.7,.1,1);
			box-shadow: 0 0 50px 3px rgba(0, 0, 0, 0.11);
			border-left: 1px solid $grey-light;
			overflow-y: scroll;

			p {
				font-size: 0.8rem;
				margin-bottom: 6px;
			}

			&-active {
				transform: translateX(0);
			}

			&-header {
				width: 100%;
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 16px 24px 16px 24px;

				i {
					cursor: pointer;
					padding: 10px;
					margin-right: -5px;
					color: $secondary;
				}

				h3 {
					font-weight: 700;
					font-size: 1.2rem;
					margin-bottom: 0;
				}
			}

			&-body {

				input {
					height: $editor-side-input-height;
					min-height: $editor-side-input-height;
				}


				.form-select-cont,
				.form-select {
					height: $editor-side-input-height
				}
			}

			&-cont {
				padding: 24px;
				border-bottom: 1px solid $grey-light;

				.form-group:last-child {
					margin-bottom: 0;
				}

				&:first-child {
					padding-top: 10px;
				}
			}

		}

		// Mobile Down
		// =========================================================================

		@include media-mob-down {

			&-preview-btn {
				display: none;
			}
		}


			// Tablet
		// =========================================================================

		@include media-tab {

			&-sidebar {
				width: 340px;

				&-header {
					margin-top: 1rem;
				}
			}

			&-save-btn {
				display: inline;
			}
		}
	}

</style>