<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<pre>{{ data }}7</pre>
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
					<!-- Title -->
					<div class="title">
						<div class="form-group">
							<input class="form-input form-input-white" type="text" placeholder="Title" v-model="data.title">
						</div>
					</div>
					<!-- Tabs -->
					<div class="tabs">
						<div class="tabs-header">
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 1 }" @click="activeTab = 1">Content</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 2 }" @click="activeTab = 2">Meta</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 3 }" @click="activeTab = 3">SEO</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 4 }" @click="activeTab = 4">Code Injection</div>
						</div>
						<!-- Fields -->
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 1 }">
							<Fields :layout="fieldLayout" :fields.sync="data.fields"></Fields>
						</div>
						<!-- Meta Options -->
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 2 }">
							<MetaOptions @update="updateMeta"></MetaOptions>
						</div>
						<!-- Seo Options -->
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 3 }">
							<SeoOptions></SeoOptions>
						</div>
						<!-- Code Injection -->
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 4 }">
							<CodeInjection @update="updateCodeInjection"></CodeInjection>
						</div>
					</div>
				</div><!-- /Col -->
				<!-- =====================
					Options & Properties
					===================== -->
				<div class="col-12 col-desk-3">
					<!-- Options -->
					<div class="options">
						<h2>Options</h2>
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
								<select class="form-select" id="options-author" v-model="selectedAuthor" @change="getFieldLayout">
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
					</div><!-- /Properties -->
					<!-- Properties -->
					<div class="options">
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
					</div>
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

export default {
	name: "Single",
	components: {
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
		publishedDate: new Date(),
		fieldLayout: [],
		templates: [],
		layouts: [],
		resource: {},
		newItem: false,
		data: {
			"title": "",
			"slug": "",
			"fields": {},
			"meta": {},
			"options": {
				"codeinjection_header": "",
				"codeinjection_footer": "",
			},
			"updated_at": new Date(),
			"created_at": new Date(),
		},
		selectedAuthor: "",
		selectedTemplate: "",
		selectedLayout: "",
	}),
	beforeMount() {
		this.getFieldLayout()
		this.setResource()
		this.setNewUpdate()
	},
	mounted() {
		//this.getResourceDataTest()
		this.getUsers()
		this.getTemplates()
		this.getLayouts()
	},
	methods: {
		getResourceData() {
			const id = this.$route.params.id;
			this.axios.get(`/posts/${id}`)
				.then(res => {
					this.data = res.data.data.post
				})
				.catch(err => {
					console.log(err)
				})
		},
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
		getTemplates() {
			this.axios.get("/templates")
				.then(res => {
					this.templates = res.data.data.templates
				})
		},
		getLayouts() {
			this.axios.get("/layouts")
				.then(res => {
					this.layouts = res.data.data.layouts
				})
		},
		getUsers() {
			if (!this.$store.state.users.length) {
				this.axios.get(`/users`)
					.then(res => {
						const users = res.data.data
						this.$store.commit("setUsers", users)
						console.log(users)
						this.users = users
					})
					.catch(err => {
						this.helpers.handleResponse(err)

					})
			} else {
				this.users = this.$store.state.users
			}
		},
		setResource() {
			const resource = this.getResources[this.$route.query.resource]
			this.resource = resource === undefined ? {
				"name": "page",
				"friendly_name": "Page",
				"slug": "",
				"icon": 'fal fa-file'
			} : resource
		},
		setNewUpdate() {
			const isNew = this.$route.params.id === "new"
			this.newItem = isNew
			if (!isNew) {
				this.getResourceData()
			}
		},
		save() {
			if (this.newItem) {
				this.axios.post("/posts", this.data)
					.then(res => {
						console.log(res);
					})
					.catch(err => {
						console.log(err);
					})
			} else {
				this.axios.put("/posts/" + this.$route.params.id, this.data)
					.then(res => {
						console.log(res);
					})
					.catch(err => {
						console.log(err);
					})
			}

		},
		updateFields(e) {
			this.data.fields = e
		},
		updateMeta(e) {
			this.data.meta = e
		},
		updateCodeInjection(e) {
			this.data.codeinjection_header = e.header;
			this.data.codeinjection_footer = e.footer;
		},
	},
	computed: {
		getResources() {
			return this.$store.state.theme.resources;
		},
		computedSlug() {
			let slugResult = '';

			const test = slugify("kjfhsdf£$fdsfldshf", {
				replacement: '-',    // replace spaces with replacement
				remove: null,        // regex to remove characters
				lower: true          // result in lower case
			})

			console.log(test)

			return slugResult;
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
		margin-bottom: 1.4rem;
	}

</style>