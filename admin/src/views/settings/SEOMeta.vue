<!-- =====================
	Settings - SEO & Meta
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Seo & Meta Settings</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Settings</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios && loadingMeta" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios && !loadingMeta" class="row trans-fade-in-anim">
				<!-- =====================
					Visibility
					===================== -->
				<div class="col-12">
					<h6 class="margin">Visibility</h6>
						<div class="card card-small-box-shadow card-expand">
						<!-- Public -->
						<div class="collapse-border-bottom">
							<div class="card-header card-header-block">
								<div>
									<h4 class="card-title">Public</h4>
									<p>By disabling public, no social media meta data will be outputted and a <code v-text="'<meta name=\'robots\' content=\'noindex\'>'"></code> will be placed globally.</p>
								</div>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="seo-public" v-model="data['seo_public']" :true-value="true" :false-value="false" />
									<label for="seo-public"></label>
								</div>
							</div><!-- /Card Header -->
						</div><!-- /Public -->
					</div><!-- /Card -->
				</div>
				<!-- =====================
					Meta Information
					===================== -->
				<div class="col-12">
					<h6 class="margin">Global meta</h6>
					<div v-if="!loadingMeta" class="card card-small-box-shadow card-expand">
						<MetaForm :meta="meta" @update="updateMeta"></MetaForm>
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Sitemap
					===================== -->
				<div class="col-12">
					<h6 class="margin">Sitemap</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Serve Sitemap -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header card-header-block">
									<div>
										<h4 class="card-title">Serve sitemap?</h4>
										<p>By disabling this selection the <code>/sitemap.xml</code> file will not be automatically served from the Verbis server.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="seo-sitemap-serve" v-model="data['seo_sitemap_serve']" checked :true-value="true" :false-value="false" />
										<label for="seo-sitemap-serve"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Serve Sitemap -->
						<!-- View -->
						<Collapse v-if="data['seo_sitemap_serve']" :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">View sitemap</h4>
										<p>View the XML sitemap that Verbis generates.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<div class="seo-sitemap-btn">
										<a :href="getSiteUrl + '/sitemap.xml'" class="btn" target="_blank">Open in new tab</a>
									</div>
									<prism-editor class="prism prism-large" v-model="sitemap" :readonly="true" :highlight="highlighter"></prism-editor>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!--/ View -->
						<!-- Exclude Resources -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Exclude resources</h4>
										<p>Select resources to exclude from the XML sitemap.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body" ref="resources">
									<vue-tags-input
										v-model="tag"
										:tags="selectedTags"
										:autocomplete-items="filteredResources"
										@tags-changed="updateTags"
										add-only-from-autocomplete
										:autocomplete-min-length="0"
										@focus="updateHeight"
										@blur="updateHeight"
										placeholder="Add excluded resource"
									/>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!--/ View -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Robots
					===================== -->
				<div class="col-12">
					<h6 class="margin">Robots</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Serve Robots -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header card-header-block">
									<div>
										<h4 class="card-title">Serve robots?</h4>
										<p>By disabling this selection the <code>/robots.text</code> file will not be automatically served from the Verbis server.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="seo-robots-serve" v-model="data['seo_robots_serve']" checked :true-value="true" :false-value="false" />
										<label for="seo-robots-serve"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Serve Robots -->
						<!-- Content -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Edit robots.txt</h4>
										<p>Edit the <code>/robots.txt</code> file which is automatically served by the Verbis server.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Robots File" :error="errors['robots']">
										<textarea rows="6" class="form-textarea form-input form-input-white" type="text" v-model="data['seo_robots']"></textarea>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!--/Content -->
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
import {optionsMixin} from "@/util/options";
import MetaForm from "@/components/meta/Meta";
import Collapse from "../../components/misc/Collapse";
import FormGroup from "../../components/forms/FormGroup";
import VueTagsInput from '@jack_reddico/vue-tags-input';
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';

export default {
	name: "SeoMeta",
	mixins: [optionsMixin],
	components: {
		FormGroup,
		Collapse,
		MetaForm,
		Breadcrumbs,
		VueTagsInput,
	},
	data: () => ({
		errorMsg: "Fix the errors before saving SEO & Meta settings.",
		successMsg: "Seo & Meta options updated successfully.",
		showImageModal: false,
		selectedImageType: "",
		facebookImage: false,
		twitterImage: false,
		meta: {},
		loadingMeta: true,
		selectedTags: [],
		tags: [],
		tag: "",
		sitemap: "hey",
	}),
	mounted() {
		this.getSitemap();
	},
	methods: {
		/*
		 * runAfterGet()
		 * Insert facebook or twitter image & update the height
		 * after the options have been obtained.
		 */
		runAfterGet() {
			for (const key in this.data) {
				if (key.includes('meta_')) {
					this.$set(this.meta, key, this.data[key]);
				}
			}
			this.loadingMeta = false;
			this.setTags();
		},
		/*
		 * updateMeta()
		 * Sets the data using the key when the meta updates.
		 */
		updateMeta(val, key) {
			this.$set(this.data, key, val);
		},
		/*
		 * setTags()
		 * Sets the tags after the data has been obtained from
		 * the API.
		 */
		setTags() {
			// Push to the autocomplete items.
			this.tags = this.getResources.map(m => {
				return {
					text: m.friendly_name,
					name: m.name,
				}
			});
			// Push to the selected tags.
			const excluded = this.data['seo_sitemap_excluded']
			if (excluded && excluded.length) {
				excluded.forEach(r => {
					const tag = this.tags.find(t => t.name === r);
					if (tag) {
						this.selectedTags.push({
							text: tag.text,
							name: tag.name,
						})
					}
				});
			}
		},
		/*
		 * updateTags()
		 * Update the data once the tag/resource has been inputted.
		 */
		updateTags(resources) {
			this.data['seo_sitemap_excluded'] = {};
			let arr = [];
			resources.forEach(r => {
				arr.push(r.name);
			})
			this.$set(this.data, 'seo_sitemap_excluded', arr)
		},
		/*
		 * updateHeight()
		 * Update excluded resources on focus & blur.
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.resources.closest(".collapse-content"));
			});
		},
		/*
		 * getSitemap()
		 * Obtain the sitemap.xml file from the API.
		 */
		getSitemap() {
			this.axios.get("/sitemap.xml", {
				baseURL: ""
			})
				.then(res => {
					this.sitemap = res.data;
				});
		},
		/*
		 * highlighter()
		 * Return xml for prism editor.
		 */
		highlighter(code) {
			return highlight(code, languages.xml, "xml");
		},
	},
	computed: {
		/*
		 * getResources()
		 * Obtain the sites resources for use in the exclude resources section
		 * of the sitemap.xml.
		 */
		getResources() {
			let resources = this.$store.state.theme.resources;
			let arr = [];
			for (const resource in resources) {
				arr.push(resources[resource]);
			}
			arr.push({
				"name": "pages",
				"friendly_name": "Pages",
				"singular_name": "Page",
				"slug": "",
				"icon": 'fal fa-file'
			});
			return arr;
		},
		/*
		 * filteredResources()
		 * Retrieve the resources for the select tags.
		 */
		filteredResources() {
			return this.tags.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.seo {

	// Sitemap
	// =========================================================================

	&-sitemap-btn {
		top: 0;
		right: 1.2rem;
		position: absolute;
		display: flex;
		justify-content: flex-end;
		z-index: 99;
	}
}

</style>