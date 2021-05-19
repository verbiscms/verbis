<!-- =====================
	Settings - General
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>General Settings</h1>
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
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-if="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="col-12">
					<h6 class="margin">Site options</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Title & description -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['site_title'] ||  errors['site_description'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Title & description</h4>
										<p>Details of the website that will be used publicly around the web.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Title -->
									<FormGroup label="Site title*" :error="errors['site_title']">
										<input class="form-input form-input-white" type="text" v-model="data['site_title']">
									</FormGroup>
									<!-- Description -->
									<FormGroup label="Site description*" :error="errors['site_description']">
										<input class="form-input form-input-white" type="text" v-model="data['site_description']">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Title & description -->
						<!-- Url -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['site_url']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Url</h4>
										<p>Set the global site url, be careful when changing the site url, it may have undesired effects.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Url -->
									<FormGroup label="Site title*" :error="errors['site_url']">
										<input class="form-input form-input-white" type="text" v-model="data['site_url']">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!--/ Url -->
						<!-- Logo -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['site_logo']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Logo</h4>
										<p>Set the logo of the site, this will be used for the Verbis backend and the publicly accessible routes.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Logo -->
									<div class="general-logo">
										<div v-show="!hasLogo">
											<button class="btn" @click.prevent="showImageModal = true">Insert logo</button>
										</div>
										<div v-show="hasLogo">
											<ImageWithActions @choose="showImageModal = true" @remove="hasLogo = false">
												<img :src="getSiteUrl + data['site_logo']" />
											</ImageWithActions>
										</div>
									</div>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Logo -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					General
					===================== -->
				<div class="col-12">
					<h6 class="margin">General</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Locale -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['homepage']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Homepage</h4>
										<p>Set the site's homepage.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body" ref="homepage">
									<FormGroup label="Homepage">
										<!-- Post Tags -->
										<vue-tags-input
											v-model="tag"
											:tags="selectedTags"
											:autocomplete-items="filteredItems"
											@tags-changed="updateTags"
											add-only-from-autocomplete
											:max-tags="1"
											:autocomplete-min-length="0"
											@max-tags-reached="maxTagsReached"
											placeholder="Add Post"
											@focus="updateHeight"
											:add-on-key="[13, ':', ';']"
										/>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Locale -->
						<!-- Locale -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['general_locale']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Locale</h4>
										<p>Set the site's location.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Location">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model="computedLocale">
												<option v-for="location in locale" :value="location" :key="location">{{ location }}</option>
											</select>
										</div>
									</FormGroup>
								</div>
							</template>
						</Collapse><!-- /Locale -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Social media
					===================== -->
				<div class="col-12">
					<h6 class="margin">Social</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Social Media -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' :
						errors['social_facebook_url'] || errors['social_twitter_url'] || errors['social_youtube_url'] || errors['social_linked_in'] || errors['social_instagram_url'] || errors['social_pinterest_url']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Social media URL's</h4>
										<p>Enter a valid url for the website's social media accounts.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Facebook" :error="errors['social_facebook']">
										<input class="form-input form-input-white" type="text" v-model="data['social_facebook']">
									</FormGroup>
									<FormGroup label="Twitter" :error="errors['social_twitter']">
										<input class="form-input form-input-white" type="text" v-model="data['social_twitter']">
									</FormGroup>
									<FormGroup label="Youtube" :error="errors['social_youtube']">
										<input class="form-input form-input-white" type="text" v-model="data['social_youtube']">
									</FormGroup>
									<FormGroup label="LinkedIn" :error="errors['social_linkedin']">
										<input class="form-input form-input-white" type="text" v-model="data['social_linkedin']">
									</FormGroup>
									<FormGroup label="Instagram" :error="errors['social_instagram']">
										<input class="form-input form-input-white" type="text" v-model="data['social_instagram']">
									</FormGroup>
									<FormGroup label="Pinterest" :error="errors['social_pinterest']">
										<input class="form-input form-input-white" type="text" v-model="data['social_pinterest']">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Social Media -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Contact Details
					===================== -->
				<div class="col-12">
					<h6 class="margin">Contact details</h6>
					<div class="card card-small-box-shadow card-expand card-margin-none">
						<!-- Email -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['contact_email']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Email</h4>
										<p>Enter an address for {{ getSiteTitle }}.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Email" :error="errors['contact_email']">
										<input class="form-input form-input-white" type="text" v-model="data['contact_email']" @keyup="validateEmail(true, data['contact_email'], 'contact_email')">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Email -->
						<!-- Telephone -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['contact_telephone']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Telephone</h4>
										<p>Enter a telephone number for {{ getSiteTitle }}.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Telephone">
										<input class="form-input form-input-white" type="text" v-model="data['contact_telephone']">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Telephone -->
						<!-- Address -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['contact_address']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Address</h4>
										<p>Enter an address for {{ getSiteTitle }}.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Address" :error="errors['contact_address']">
										<textarea rows="6" class="form-textarea form-input form-input-white" type="text" v-model="data['contact_address']"></textarea>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Address -->
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Insert Photo Modal
			===================== -->
		<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
			<template slot="text">
				<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertLogo" :options="false">
					<template slot="close">
						<button class="btn btn-margin-right btn-icon-mob" @click.prevent="showImageModal = false">
							<i class="feather feather-x"></i>
							<span>Close</span>
						</button>
					</template>
				</Uploader>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import FormGroup from "@/components/forms/FormGroup";
import Modal from "@/components/modals/General";
import Uploader from "@/components/media/Uploader";
import ImageWithActions from "@/components/misc/ImageWithActions";
import Collapse from "@/components/misc/Collapse";
import {validationMixin} from "@/util/validation";
import localeObj from "@/util/locale.js";
import VueTagsInput from '@jack_reddico/vue-tags-input';

export default {
	name: "General",
	title: 'General Settings',
	mixins: [optionsMixin, validationMixin],
	components: {
		Collapse,
		ImageWithActions,
		Uploader,
		Modal,
		FormGroup,
		Breadcrumbs,
		VueTagsInput,
	},
	data: () => ({
		errorMsg: "Fix the errors before saving settings.",
		successMsg: "Site options updated successfully.",
		hasLogo: true,
		showImageModal: false,
		locale: {},
		posts: [],
		tag: '',
		selectedTags: [],
	}),
	mounted() {
		this.locale = localeObj;
		this.getPosts();
	},
	methods: {
		/*
		 * insertLogo()
		 */
		insertLogo(e) {
			this.profilePicture = e;
			this.showImageModal = false;
			this.data['site_logo'] = e.url;
			this.hasLogo = true;
			// Update store with new logo.
			const site = this.$store.state.site;
			site['logo'] = e.url;
			this.$store.commit("setSite", site)
		},
		/*
		 * getPosts()
		 * Obtain the posts or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getPosts() {
			this.axios.get(`/posts?resource=pages&limit=all"`)
				.then(res => {
					this.mapPosts(res.data.data);
					this.setTags();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * mapPosts()
		 * If the posts are undefined, return an empty array.
		 * If resources are set, they are filters and returned.
		 * Or all posts are returned for filtering.
		 */
		mapPosts(posts) {
			if (posts === undefined || Object.keys(posts).length === 0 && posts.constructor === Object) {
				this.posts = [];
				return
			}

			if (this.getResources) {
				this.posts = posts.filter(p => {
					if (this.getResources.includes(p.post.resource)) {
						return p
					}
				}).map(p => ({text: p.post.title, id: p.post.id}))
				return
			}

			this.posts = posts.map(a => ({text: a.post.title, id: a.post.id}));
		},
		/*
		 * setTags()
		 * Set the existing tags on mounted if there is any by
		 * filtering through existing posts.
		 */
		setTags() {
			if (this.data['homepage'] !== "") {
				const post = this.posts.find(p => p.id === this.data['homepage']);
				if (!post) {
					return;
				}
				this.selectedTags = [{text: post.text, id: post.id}];
			}
		},
		/*
		 * updateTags()
		 * Update tags when fired.
		 */
		updateTags(tags) {
			delete this.errors['homepage'];
			this.selectedTags = tags;
			if (tags.length) {
				this.$set(this.data, "homepage", tags[0].id)
			}
		},
		/*
		 * updateHeight()
		 * Update the height of the container when searching.
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.homepage.closest(".collapse-content"));
			});
		},
		/*
		 * maxTagsReached()
		 * Handler for maximum tags reached.
		 */
		maxTagsReached() {
			this.$set(this.errors, 'homepage', true);
			this.$noty.error(`Only one homepage can be assigned.`);
		}
	},
	computed: {
		/*
		 * getSiteTitle()
		 */
		getSiteTitle() {
			return this.getSite.title === "Verbis" ? "the business" : this.getSite.title;
		},
		/*
		 * filteredItems()
		 * Filter tags.
		 */
		filteredItems() {
			return this.posts.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
		computedLocale: {
			/*
			 * getDefaultLocale()
			 * Sets the default locale if there is none or returns
			 * the set locale.
			 */
			get() {
				if (this.data['general_locale'] === undefined || this.data['general_locale'] === "") {
					return "English (United Kingdom)";
				}
				return this.locale[this.data['general_locale']];
			},
			/*
			 * setLocale()
			 * Sets the locale by key.
			 */
			set(value) {
				this.data['general_locale'] = Object.keys(this.locale).find(key => this.locale[key] === value);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

</style>
