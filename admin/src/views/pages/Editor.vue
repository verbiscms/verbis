<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<div class="header-icon-cont">
								<i :class="getResourceData.icon"></i>
								<h1>New Page</h1>
							</div>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<div class="form-select-cont form-input">
									<select class="form-select">
										<option value="" disabled selected>Bulk actions</option>
										<option value="">Move to drafts</option>
										<option value="">Delete</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white">Apply</button>
								<button class="btn btn-icon btn-orange">
									<i class="fal fa-plus"></i>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12 col-desk-9">
					<div class="tabs">
						<div class="tabs-header">
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 1 }" @click="activeTab = 1">Content</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 2 }" @click="activeTab = 2">Meta</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 3 }" @click="activeTab = 3">SEO</div>
							<div class="tabs-label" :class="{ 'tabs-label-active' : activeTab === 4 }" @click="activeTab = 4">Code Injection</div>
						</div>
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 1 }">
							<!-- Text -->
							<div class="field">
								<div class="field-wrapper">
									<div class="field-title-cont">
										<span class="field-collapse" v-on:click="this.isActive = !this.isActive"></span>
										<div class="field-title">
											<h5>Text</h5>
											<p>Add a title for the page</p>
										</div>
									</div>
									<div class="field-controls">
										<i class="feather icon-trash-2"></i>
										<i class="feather icon-arrow-up"></i>
										<i class="feather icon-arrow-down"></i>
										<i class="fal fa-arrows"></i>
									</div>
								</div>
								<div class="field-content" ref="myText" :style="[isActive ? { height : computedHeight } : {}]">
									<input class="form-input form-input-white" type="text" value="How the development team go their jobs at Reddico">
								</div>
							</div>
							<!-- Rich Text -->
							<div class="field">
								<div class="field-wrapper">
									<div class="field-title-cont">
										<span class="field-collapse"></span>
										<div class="field-title">
											<h5>Rich text</h5>
											<p>Add rich text content for the news page.</p>
										</div>
									</div>
									<div class="field-controls">
										<i class="feather icon-trash-2"></i>
										<i class="feather icon-arrow-up"></i>
										<i class="feather icon-arrow-down"></i>
										<i class="fal fa-arrows"></i>
									</div>
								</div>
								<div class="field-content">
									<RichText></RichText>
								</div>
							</div>
						</div>
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 2 }">
							<!-- Serp -->
							<div class="field">
								<h4>SERP Preview</h4>
								<div class="field-content">
									<div class="serp">
										<span class="serp-title">How the development team got their jobs at Reddico</span>
										<span class="serp-url">https://reddico.co.uk/how-we-got-our-development-jobs/</span>
										<span class="serp-description">Have you ever wondered how to become a developer? Our team has taken different routes to get to their jobs here – with some surprising diversions along the way. We asked them to explain how they’ve progressed over the years.</span>
									</div>
								</div>
							</div>
							<!-- Meta -->
							<div class="field">
								<h4>Meta</h4>
								<div class="field-content">
									<div class="form-group">
										<h5>Title</h5>
										<p>This will appear at the top of the search preview</p>
										<input class="form-input form-input-white" type="text">
									</div>
									<div class="form-group">
										<h5>Description</h5>
										<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
										<textarea class="form-input form-input-white" type="text" rows="4"></textarea>
									</div>
								</div>
							</div>
							<!-- Facebook -->
							<div class="field">
								<h4>Facebook (Open Graph)</h4>
								<div class="field-content">
									<div class="form-group">
										<h5>Title</h5>
										<p>This will appear at the top of the search preview</p>
										<input class="form-input form-input-white" type="text">
									</div>
									<div class="form-group">
										<h5>Description</h5>
										<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
										<textarea class="form-input form-input-white" type="text" rows="4"></textarea>
									</div>
								</div>
							</div>
						</div>
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 3 }">
							<h3>Third Tab Content</h3>
							<p>Vexillologist messenger bag master...</p>
						</div>
						<!-- Code Injection -->
						<div class="tabs-panel" :class="{ 'tabs-panel-active' : activeTab === 4 }">
							<!-- Head -->
							<div class="field">
								<h4>Header</h4>
								<div class="form-group">
									<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
									<textarea class="form-input form-input-white" type="text" rows="14"></textarea>
								</div>
							</div>
							<!-- Footer -->
							<div class="field">
								<h4>Footer</h4>
								<div class="form-group">
									<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
									<textarea class="form-input form-input-white" type="text" rows="14"></textarea>
								</div>
							</div>
						</div>
					</div>
				</div><!-- /Col -->
				<!-- Page Options -->
				<div class="col-12 col-desk-3">
					<div class="options">
						<h2>Options</h2>
						<!-- URL -->
						<div class="form-group">
							<label class="form-label" for="options-url">URL</label>
							<input class="form-input form-input-white" type="text" id="options-url">
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
							<input class="form-input form-input-white" type="text" id="options-author">
						</div>
						<!-- Date -->
						<div class="form-group">
							<label class="form-label">Published Date</label>
							<DatePicker class="date" color="blue" :value="null"></DatePicker>
						</div>
					</div><!-- /Options -->
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
import RichText from "@/components/editor/RichText";
import DatePicker from 'v-calendar/lib/components/date-picker.umd'

export default {
	name: "Single",
	components: {
		Breadcrumbs,
		RichText,
		DatePicker,
	},
	data: () => ({
		activeTab: 1,
		computedHeight: 'auto',
		fieldHeights: [],
		isActive: true,
		users: [],
	}),
	mounted() {
		//this.getResourceDataTest()
		this.initHeight()
		this.getUsers()
	},
	methods: {
		findResourceByName(resources, name) {
			const resource = resources.find(r => r.name.toLowerCase() === name);
			if (resource === undefined) {
				//return this.$router.push('/404')
			}
			return resource
		},
		getUsers() {
			this.axios.get(`/users`)
				.then(res => {
					console.log(res)
				})
				.catch(err => {
					console.log(err)
				})
		},
		getResourceDataTest() {
			const id = this.$route.params.id;
			this.axios.get(`/posts/${id}`)
				.then(res => {
					console.log(res)
				})
				.catch(err => {
					console.log(err)
				})
		},
		initHeight() {
			this.computedHeight= getComputedStyle(this.$refs['myText']).height;
		}
	},
	computed: {
		getResourceData() {
			// const resources = this.$store.state.resources,
			// 	currentSlug = this.$route.params.resource,
			// 	resource = {};
			//
			// if (currentSlug === "pages") {
			// 	resource.icon = "fal fa-file"
			// } else {
			// 	const r = this.findResourceByName(resources, currentSlug)
			// 	resource.icon = r.options.icon
			// }
			//
			// resource.name = currentSlug
			//
			// return resource;

			return 0;
		},
	}
};
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Dummy
// =========================================================================

</style>