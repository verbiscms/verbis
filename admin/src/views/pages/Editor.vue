<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container" v-if="!loadingResourceData">
			<!-- =====================
				Header
				===================== -->
			<div class="row">
				<div class="col-12">
					<!-- Header -->
					<header class="header header-with-actions">
						<div class="header-title">
							<div class="header-icon-cont">
								<i :class="resource.icon"></i>
								<h1 v-if="newItem">Add a new {{ resource.friendly_name }}</h1>
								<h1 v-else>Edit {{ resource.friendly_name }}</h1>
							</div>
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
				<div class="col-12 col-desk-9">
					<Tabs @update="activeTab = $event">
						<template slot="item">Content</template>
						<template slot="item">Meta</template>
						<template slot="item">SEO</template>
						<template slot="item">Code Injection</template>
					</Tabs>
					<div v-if="!loadingResourceData">
						<!-- Content & Fields -->
						<div v-if=" fieldLayout.length" class="tabs-panel tabs-panel-naked" :class="{ 'tabs-panel-active' : activeTab === 1 }">
							<!-- Title -->
							<div class="title">
								<div class="form-group">
									<input class="form-input form-input-white" type="text" placeholder="Add title" v-model="data.title">
								</div>
							</div>
							<Fields :layout="fieldLayout" :fields.sync="data.fields" :error-trigger="errorTrigger"></Fields>
						</div>
						<!-- Meta Options -->
						<div class="tabs-panel tabs-panel-naked" :class="{ 'tabs-panel-active' : activeTab === 2 }">
							<MetaOptions :meta.sync="data.options.meta" :url="computedBaseSlug"> </MetaOptions>
						</div>
						<!-- Seo Options -->
						<div class="tabs-panel tabs-panel-naked" :class="{ 'tabs-panel-active' : activeTab === 3 }">
							<SeoOptions></SeoOptions>
						</div>
						<!-- Code Injection -->
						<div class="tabs-panel tabs-panel-naked" :class="{ 'tabs-panel-active' : activeTab === 4 }">
							<CodeInjection :header="data.codeinjection_head" :footer="data.codeinjection_foot" @update="updateCodeInjection"></CodeInjection>
						</div>
					</div>
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
								<input class="form-input form-input-white" type="text" id="options-url" v-model="data.slug">
							</div>
							<!-- Status -->
							<div class="form-group">
								<label class="form-label" for="options-status">Status</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="options-status">
										<option value="" disabled selected>Select status</option>
										<option value="drafts">Draft</option>
										<option value="published">Published</option>
									</select>
								</div>
							</div>
							<!-- Author -->
							<div class="form-group">
								<label class="form-label" for="options-author">Author</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="options-author" v-model="data.author" @change="getFieldLayout">
										<option value="" disabled selected>Select author</option>
										<option v-for="user in users" :value="user.id" :key="user.uuid">{{ user.first_name }} {{ user.last_name }}</option>
									</select>
								</div>
							</div>
							<!-- Date -->
							<div class="form-group">
								<label class="form-label">Published Date</label>
								<DatePicker class="date" color="blue" :value="null" v-model="publishedDate"></DatePicker>
							</div>
						</div><!-- /Options -->
						<!-- Properties -->
						<div class="editor-sidebar-properties">
							<h2>Properties</h2>
							<!-- Template -->
							<div class="form-group">
								<label class="form-label" for="properties-template">Template</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="properties-template" v-model="selectedTemplate" @change="getFieldLayout">
										<option value="" disabled selected>Select template</option>
										<option v-for="template in templates" :value="template.key" :key="template.key">{{ template.name }}</option>
									</select>
								</div>
							</div>
							<!-- Layout -->
							<div class="form-group">
								<label class="form-label" for="properties-layout">Layout</label>
								<div class="form-select-cont form-input">
									<select class="form-select" id="properties-layout" v-model="selectedLayout" @change="getFieldLayout">
										<option value="" disabled selected>Select template</option>
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

export default {
	name: "Single",
	components: {
		Tabs,
		Fields,
		Breadcrumbs,
		DatePicker,
		MetaOptions,
		SeoOptions,
		CodeInjection,
	},
	data: () => ({
		activeTab: 1,
		fieldHeights: [],
		users: [],
		slug: "",
		rootSlug: "",
		isCustomSlug: false,
		publishedDate: new Date(),
		fieldLayout: [],
		templates: [],
		layouts: [],
		resource: {},
		newItem: false,
		selectedAuthor: "",
		selectedTemplate: "",
		selectedLayout: "",
		errorTrigger: false,
		data: {
			"title": "",
			"slug": "",
			"fields": {},
			"author": 0,
			"options": {},
			"categories": [],
			"codeinjection_head": "",
			"codeinjection_foot": "",
			"updated_at": new Date(),
			"created_at": new Date(),
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
		this.getSuccessMessage();
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
		 * getResourceData()
		 * Get the page data, if none exists return 404.
		 */
		getResourceData() {
			const id = this.$route.params.id;
			this.axios.get(`/posts/${id}`)
				.then(res => {
					this.data = res.data.data.post;
					if (!this.data) {
						this.$router.push({ name : 'not-found' })
					}
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
					"layout": this.selectedLayout,
					"page_template": this.selectedTemplate,
					"user_id": this.selectedAuthor,
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
		/*
		 * save()
		 * Save the new page, check for field validation.
		 */
		save() {
			this.errorTrigger = true;
			this.$nextTick().then(() => {
				if (document.querySelectorAll(".field-cont-error").length === 0) {
					if (this.newItem) {
						this.axios.post("/posts", this.data)
							.then(res => {
								// Push to new page if successfull
								this.$router.push({
									name: 'editor',
									params: { id : res.data.data.post.id },
									query: { success : "true" }
								})
							})
							.catch(err => {
								console.log(err);
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
		/*
		 * updateCodeInjection()
		 * Update code injection from component.
		 */
		updateCodeInjection(e) {
			this.data['codeinjection_head'] = e.header;
			this.data['codeinjection_foot'] = e.footer;
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
		computedSlug() {
			//let slugResult = '';

			const rootSlug = this.resource.name === "page" ? "" : this.resource.name;
			//this.rootSlug = ""

			const test = slugify(rootSlug + this.title, {
				replacement: '-',    // replace spaces with replacement
				remove: null,        // regex to remove characters
				lower: true          // result in lower case
			})

			return test;
		},
		computedBaseSlug: {
			get: function(){
				return this.customSlug;
			},
			set: function(value){
				if(value.length < 1){
					this.customSlug = '/';
					this.isCustomSlug = false;
				} else {
					this.customSlug = value;
					this.isCustomSlug = true;
				}
			}
		},
	}
};
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Title
	// =========================================================================

	.title {
		margin-bottom: 1.6rem;
	}

	.editor {

		// Sidebar
		// =========================================================================

		&-sidebar {
			position: sticky;
			top: 100px;

			&-options {
				margin-bottom: 1.6rem;
			}
		}

	}


</style>