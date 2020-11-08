*<!-- =====================
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
							<h1 v-else>Edit {{ resource['singular_name'] ? resource['singular_name'] : resource['friendly_name']  }}</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<button class="btn btn-icon btn-white btn-margin-right" @click.prevent="sidebarOpen = !sidebarOpen">
									<i class="feather feather-settings"></i>
								</button>
								<a :href="getSiteUrl + computedSlug" target="_blank" class="btn btn-fixed-height btn-margin-right btn-white btn-flex">Preview</a>
								<button class="btn btn-fixed-height btn-orange" @click.prevent="save">
									<span v-if="newItem">Publish</span>
									<span v-else-if="newItem">Publish</span>
									<span v-else>Update</span>
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
				<div class="col-12 col-desk-12 editor-main-col">
					<Tabs @update="activeTab = $event - 1" :default-tab="activeTab">
						<template slot="item">Content</template>
						<template slot="item">Meta</template>
						<template slot="item">SEO</template>
						<template slot="item">Code Injection</template>
						<template slot="item">Insights</template>
					</Tabs>
					<!-- Spinner -->
					<div v-if="loadingResourceData" class="media-spinner spinner-container">
						<div class="spinner spinner-large spinner-grey"></div>
					</div>
					<!-- Content & Fields -->
					<transition v-else name="trans-fade" mode="out-in">
						<div v-if="activeTab === 0" :key="1">
							<!-- Title -->
							<div class="editor-title">
								<FormGroup class="form-group-no-margin" :error="errors['title']">
									<input type="text" placeholder="Add title" v-model="data.title">
								</FormGroup>
								<div @click="handleSlugClick" class="editor-title-slug">
									<i class="feather feather-edit-2"></i>
									<p>{{ computedSlug }}</p>
								</div>
							</div>
							<Fields :layout="fieldLayout" :fields.sync="data.fields" :error-trigger="errorTrigger"></Fields>
						</div>
						<!-- Meta Options -->
						<MetaOptions v-if="activeTab === 1" :key="2" :meta.sync="data.options.meta" :url="computedSlug"></MetaOptions>
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
					<h6 class="margin">Properties</h6>
					<!-- URL -->
					<FormGroup class="form-url" label="Url" :error="errors['slug']">
						<div class="form-url-cont">
							<input class="form-input form-input-white" type="text" id="options-url" v-model="slug" :disabled="!slugBtn">
							<i class="feather feather-edit" @click="slugBtn = !slugBtn"></i>
						</div>
						<h4>{{ computedSlug }}</h4>
					</FormGroup><!-- /Url -->
					<!-- Status -->
					<FormGroup label="Status">
						<div class="form-select-cont form-input">
							<select class="form-select" id="options-status" v-model="data.status">
								<option value="" disabled selected>Select status</option>
								<option value="draft">Draft</option>
								<option value="published">Published</option>
							</select>
						</div>
					</FormGroup><!-- /Status -->
					<!-- Author -->
				</div>
				<div class="editor-sidebar-cont">
					<FormGroup label="Author">
						<div class="form-select-cont form-input">
							<select class="form-select" id="options-author" v-model="data['author']" @change="getFieldLayout">
								<option value="0" disabled selected>Select author</option>
								<option v-for="user in users" :value="user.id" :key="user.uuid">{{ user.first_name }} {{ user.last_name }}</option>
							</select>
						</div>
					</FormGroup><!-- /Author -->
					<FormGroup label="Category">
						<!-- User Tags -->
						<vue-tags-input
							v-model="tag"
							:tags="selectedTags"
							:autocomplete-items="filteredCategories"
							@tags-changed="updateCategoriesTags"
							add-only-from-autocomplete
							@max-tags-reached="$noty.warning('Only one category per post is permitted')"
							placeholder="Add category"
							:max-tags="1"
						/>
					</FormGroup>
				</div>
				<div class="editor-sidebar-cont">
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
								<option value="" disabled selected>Select layout</option>
								<option v-for="layout in layouts" :value="layout.key" :key="layout.key">{{ layout.name }}</option>
							</select>
						</div>
					</FormGroup><!-- /Layout -->
				</div>
				<div class="editor-sidebar-cont">
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
import slugify from "slugify";
import Tabs from "@/components/misc/Tabs";
import FormGroup from "@/components/forms/FormGroup";
import VueTagsInput from '@jack_reddico/vue-tags-input';

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
		Insights,
		VueTagsInput,
	},
	data: () => ({
		activeTab: 0,
		users: [],
		slug: "",
		slugBtn: false,
		fieldLayout: [],
		templates: [],
		layouts: [],
		resource: {},
		categories: [],
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
		sidebarOpen: false,
		tag: "",
		selectedTags: [],
	}),
	beforeMount() {
		this.setResource()
		this.setNewUpdate();
		this.setTab();
	},
	mounted() {
		if (this.newItem) {
			this.getFieldLayout();
		}
		this.getUsers();
		this.getTemplates();
		this.getLayouts();
		this.getCategories();
	},
	methods: {
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

					// Set field layouts
					this.fieldLayout = res.data.data.layout;

					// Set categories
					res.data.data.categories.forEach(category => {
						this.selectedTags.push({
							text: category.name,
							id: category.id,
						})
					});

					// Set date format
					this.setDates()

				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.loadingResourceData = false;
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
		 * getCategories()
		 * Obtain the categories.
		 */
		getCategories() {
			this.axios.get(`/categories?filter={"resource":[{"operator":"=", "value": "${this.resource['name']}"}]}`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
				.then(res => {
					this.mapCategories(res.data.data);
				})
				.catch(err => {
					console.log(err);
					this.helpers.handleResponse(err);
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
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * setResource()
		 * Set the resource from the query parameter, if none defined,
		 * set default page 'resource'.
		 */
		setResource() {
			const resource = this.getTheme['resources'][this.$route.query.resource]
			this.resource = resource === undefined ? {
				"name": "page",
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
			} else {
				this.loadingResourceData = false;
			}
		},
		/*
		 * setDates()
		 * Set the dates for moment when getting back from API.
		 */
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
									query: { success : "true" }
								})

								// Set defaults
								this.data = res.data.data.post;

								this.data.author = res.data.data.author.id;
								this.newItem = false;
								this.setDates();
								this.getSuccessMessage();
							})
							.catch(err => {
								this.helpers.checkServer(err);
								if (err.response.status === 400) {
									const msg = err.response.data.message,
										errors = err.response.data.data.errors;
									if (msg && !errors) this.$noty.error(msg);
									this.validate(errors);
									return;
								}
								this.helpers.handleResponse(err);
							})
					} else {
						this.axios.put("/posts/" + this.$route.params.id, this.data)
							.then(() => {
								this.$noty.success("Page updated successfully.")
							})
							.catch(err => {
								this.helpers.handleResponse(err);

							});
					}
				} else {
					this.$noty.error("Fix the errors before saving the post.")
				}
			})
		},
		/*
		 * mapCategories()
		 * Create a new categories array from input & map.
		 */
		mapCategories(categories) {
			if (categories && !this.helpers.isEmptyObject(categories)) {
				this.categories = categories.map(a => {
					return {
						text: a.name,
						id: a.id
					}
				});
			}
		},
		/*
		 * updateTags()
		 * Updates the categories when the tags changes.
		 */
		updateCategoriesTags(categories) {
			this.$set(this.data, 'categories', []);
			let catArr = [];
			categories.forEach(category => {
				catArr.push(category.id);
			});
			this.$set(this.data, 'categories', catArr);
		},
		/*
		 * handleSlugClick()
		 * Open the sidebar and set to custom slug if the user
		 * clicks the slug edit button under the title
		 */
		handleSlugClick() {
			this.sidebarOpen = true;
			this.slugBtn = true;
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
		 * getBaseSlug()
		 * Get the base slug (resource).
		 */
		getBaseSlug() {
			return  this.resource.name === "page" ? "/" : "/" + this.resource.name + "/";
		},
		/*
		 * filteredCategories()
		 */
		filteredCategories() {
			return this.categories.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
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

	$editor-side-input-height: 50px;

	.editor {

		// Title
		// =========================================================================

		&-title {
			margin: 3rem 0;

			input {
				background-color: transparent;
				outline: none;
				border: none;
				width: 100%;
				font-size: 2rem;
				color: $black;
				font-weight: 600;
			}

			&-slug {
				display: inline-flex;
				align-items: center;
				margin-top: 6px;
				cursor: pointer;

				p {
					color: $grey;
					margin: 0;
				}

				i {
					color: $grey;
					margin-right: 6px;
					font-size: 14px;
				}

				p,
				i {
					transition: 200ms ease color;
				}

				&:hover {

					p,
					i {
						color: $primary;
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
			width: 340px;
			background-color: $bg-color;
			z-index: 999;
			transform: translateX(100%);
			transition: transform .4s cubic-bezier(.1,.7,.1,1);
			box-shadow: 0 0 50px 3px rgba(0, 0, 0, 0.11);
			border-left: 1px solid $grey-light;

			&-active {
				transform: translateX(0);
			}


			.form-label {
				font-size: 0.7rem;
				//color: $secondary;
			}

			&-header {
				width: 100%;
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 16px 24px;

				i {
					cursor: pointer;
					padding: 10px;
					margin-right: -5px;
					color: $secondary;
				}

				h3 {
					font-weight: 600;
					font-size: 1.1rem;
					margin-bottom: 0;
				}
			}

			&-body {
				//padding: 0 24px;

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
			}
		}
	}

</style>