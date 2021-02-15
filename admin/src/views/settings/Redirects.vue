<!-- =====================
	Settings - Redirects
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Redirects</h1>
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
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Redirect Modal
			===================== -->
		<Redirect :show.sync="showRedirectModal" :redirect-key="selectedRedirectKey" :redirect-update="selectedRedirect" @update="updateRedirect"></Redirect>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';
import Redirect from "@/components/modals/Redirect";

export default {
	name: "SeoMeta",
	mixins: [optionsMixin],
	components: {
		Redirect,
		Breadcrumbs,
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
		sitemap: "",
		showRedirectModal: false,
		selectedRedirect: false,
		selectedRedirectKey: false,
	}),
	mounted() {
		this.getSitemap();
	},
	watch: {
		/*
		 * show()
		 * Watch if the model has been closed/opened &
		 * delete keys & redirect.
		 */
		showRedirectModal: function(val) {
			if (!val) {
				this.selectedRedirect = false;
				this.selectedRedirectKey = false;
				this.selectedRedirect = false;
			}
		}
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
		updateHeight(ref) {
			this.$nextTick(() => {
				console.log(ref);
				console.log(this.$refs);
				this.helpers.setHeight(this.$refs[ref].closest(".collapse-content"));
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
		/*
		 * updateRedirect()
		 */
		updateRedirect(redirect, key) {
			if (!this.data['seo_redirects']) {
				this.$set(this.data, 'seo_redirects', []);
			}

			if (key === "new") {
				this.data['seo_redirects'].push(redirect);
				this.updateHeight('redirects');
			}

			if (Number.isInteger(key)) {
				console.log(redirect, key)
				this.$set(this.data['seo_redirects'], key, redirect);
			}

			this.showRedirectModal = false;
		},
		/*
		 * updateRedirectHandler()
		 */
		updateRedirectHandler(key) {
			this.selectedRedirectKey = key;
			this.selectedRedirect = this.data['seo_redirects'][key];
			this.showRedirectModal = true;
		},
		/*
		 * deleteRedirect()
		 * Remove a redirect from the array.
		 */
		deleteRedirect(index) {
			if (index > -1) {
				this.data['seo_redirects'].splice(index, 1);
			}
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

		.btn {
			background-color: $white;
		}
	}

	// Redirects
	// =========================================================================

	&-redirects {


		//.table {
		//	border-top: 1px solid $grey-light;
		//}

		.card-controls {

		}

		.feather-trash-2 {
			color: $orange;
		}
		//&-header {
		//	display: flex;
		//	align-items: center;
		//	justify-content: space-between;
		//	margin-bottom: 1rem;
		//
		//	p {
		//		margin-bottom: 0;
		//	}
		//}
	}
}

</style>