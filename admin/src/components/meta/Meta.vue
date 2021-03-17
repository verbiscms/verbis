<!-- =====================
	Meta
	===================== -->
<template>
	<div>
		<!-- =====================
			General
			===================== -->
		<Collapse :show="false" class="collapse-border-bottom">
			<template v-slot:header>
				<div class="card-header">
					<div>
						<h4 class="card-title">Meta data</h4>
						<p>Meta information & content for search engines.</p>
					</div>
					<div class="card-controls">
						<i class="feather feather-chevron-down"></i>
					</div>
				</div><!-- /Card Header -->
			</template>
			<template v-slot:body>
				<div class="card-body">
					<!-- Title -->
					<FormGroup label="Title" :error="errors['meta_title']">
						<input class="form-input form-input-white" type="text" v-model="meta['meta_title']" @keyup="$emit('update', meta['meta_title'], 'meta_title')">
						<Recommendations type="title" usage="meta" :text="meta['meta_title']"></Recommendations>
					</FormGroup>
					<!-- Description -->
					<FormGroup label="Description" :error="errors['meta_description']">
						<input class="form-input form-input-white" type="text" v-model="meta['meta_description']" @keyup="$emit('update', meta['meta_description'], 'meta_description')">
						<Recommendations type="description" usage="meta" :text="meta['meta_description']"></Recommendations>
					</FormGroup>
				</div><!-- /Card Body -->
			</template>
		</Collapse><!-- /Meta -->
		<!-- =====================
			Facebook
			===================== -->
		<Collapse :show="false" class="collapse-border-bottom meta-card meta-card-facebook">
			<template v-slot:header>
				<div class="card-header">
					<div>
						<h4 class="card-title">Facebook</h4>
						<p>Facebook card information & content for social sharing.</p>
					</div>
					<div class="card-controls">
						<i class="feather feather-chevron-down"></i>
					</div>
				</div><!-- /Card Header -->
			</template>
			<template v-slot:body>
				<div class="card-body" ref="facebook">
					<div class="row">
						<div class="col-12 col-desk-6 col-hd-5 offset-hd-1 meta-card-col-right order-desk-last">
							<!-- Toggle -->
							<div class="meta-toggle">
								<h6 class="margin">Use global meta?</h6>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="use-facebook-global" :true-value="true" :false-value="false" v-model="useFacebookGlobal" @change="updateGlobal('facebook')" />
									<label for="use-facebook-global"></label>
								</div>
							</div><!-- /Toggle -->
							<h6 class="margin">Preview:</h6>
							<SocialCard :title="meta['meta_facebook_title']" :description="meta['meta_facebook_description']" :image="facebookImage"></SocialCard>
						</div><!-- /Col -->
						<div class="col-12 col-desk-6 meta-card-col-left">
							<!-- Title -->
							<FormGroup label="Title" :error="errors['meta_facebook_title']">
								<input class="form-input form-input-white" type="text" v-model="meta['meta_facebook_title']" @keyup="$emit('update', meta['meta_facebook_title'], 'meta_facebook_title')">
								<Recommendations type="title" usage="facebook" :text="meta['meta_facebook_title']"></Recommendations>
							</FormGroup>
							<!-- Description -->
							<FormGroup label="Description" :error="errors['meta_facebook_description']">
								<input class="form-input form-input-white" type="text" v-model="meta['meta_facebook_description']" @keyup="$emit('update', meta['meta_facebook_description'], 'meta_facebook_description')">
								<Recommendations type="description" usage="facebook" :text="meta['meta_facebook_description']"></Recommendations>
							</FormGroup>
							<!-- Image -->
							<FormGroup label="Image">
								<div v-show="!facebookImage">
									<button class="btn" @click.prevent="showModal('facebook')">Insert image</button>
								</div>
								<div v-show="facebookImage">
									<ImageWithActions @choose="showImageModal = true" @remove="removeImage('facebook')">
										<img :src="getSiteUrl + facebookImage['url']"/>
									</ImageWithActions>
								</div>
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- /Card Body -->
			</template>
		</Collapse><!-- /Facebook -->
		<!-- =====================
			Twitter
			===================== -->
		<Collapse :show="false" class="collapse-border-bottom meta-card meta-card-twitter">
			<template v-slot:header>
				<div class="card-header">
					<div>
						<h4 class="card-title">Twitter</h4>
						<p>Twitter card information & content for social sharing.</p>
					</div>
					<div class="card-controls">
						<i class="feather feather-chevron-down"></i>
					</div>
				</div><!-- /Card Header -->
			</template>
			<template v-slot:body>
				<div class="card-body" ref="twitter">
					<div class="row">
						<div class="col-12 col-desk-6 col-hd-5 offset-hd-1 meta-card-col-right order-desk-last">
							<!-- Toggle -->
							<div class="meta-toggle">
								<h6 class="margin">Use global meta?</h6>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="use-twitter-global" :true-value="true" :false-value="false" v-model="useTwitterGlobal" @change="updateGlobal('twitter')" />
									<label for="use-twitter-global"></label>
								</div>
							</div><!-- /Toggle -->
							<h6 class="margin">Preview:</h6>
							<SocialCard :title="meta['meta_twitter_title']" :description="meta['meta_twitter_description']" :image="twitterImage"></SocialCard>
						</div><!-- /Col -->
						<div class="col-12 col-desk-6 meta-card-col-left">
							<!-- Title -->
							<FormGroup label="Title" :error="errors['meta_twitter_title']">
								<input class="form-input form-input-white" type="text" v-model="meta['meta_twitter_title']" @keyup="$emit('update', meta['meta_twitter_title'], 'meta_twitter_title')">
								<Recommendations type="title" usage="twitter" :text="meta['meta_twitter_title']"></Recommendations>
							</FormGroup>
							<!-- Description -->
							<FormGroup label="Description" :error="errors['meta_twitter_description']">
								<input class="form-input form-input-white" type="text" v-model="meta['meta_twitter_description']" @keyup="$emit('update', meta['meta_twitter_description'], 'meta_twitter_description')">
								<Recommendations type="description" usage="twitter" :text="meta['meta_twitter_description']"></Recommendations>
							</FormGroup>
							<!-- Image -->
							<FormGroup label="Image">
								<div v-show="!twitterImage">
									<button class="btn" @click.prevent="showModal('twitter')">Insert image</button>
								</div>
								<div v-show="twitterImage">
									<ImageWithActions @choose="showImageModal = true" @remove="removeImage('twitter')">
										<img :src="getSiteUrl + twitterImage['url']"/>
									</ImageWithActions>
								</div>
							</FormGroup>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- /Card Body -->
			</template>
		</Collapse><!-- /Twitter -->
		<!-- =====================
			Insert Image Modal
			===================== -->
		<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
			<template slot="text">
				<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertImage" :options="false">
					<template slot="close">
						<button class="btn btn-margin-right btn-icon-mob" @click.prevent="showImageModal = false">
							<i class="feather feather-x"></i>
							<span>Close</span>
						</button>
					</template>
				</Uploader>
			</template>
		</Modal>
	</div>
</template>

<!-- =====================
	Scripts
	=====
	================ -->
<script>

import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import Recommendations from "@/components/meta/Recommendations";
import ImageWithActions from "@/components/misc/ImageWithActions";
import Modal from "@/components/modals/General";
import Uploader from "@/components/media/Uploader";
import SocialCard from "@/components/meta/SocialCard";

export default {
	name: "MetaForm",
	components: {
		SocialCard,
		Uploader,
		Modal,
		ImageWithActions,
		Collapse,
		FormGroup,
		Recommendations,
	},
	props: {
		meta: {
			type: [Object],
			default: function () {
				return {
					"meta_title": "",
					"meta_description": "",
					"meta_facebook_title": "",
					"meta_facebook_description": "",
					"meta_twitter_title": "",
					"meta_twitter_description": "",
				};
			},
		},
	},
	data: () => ({
		facebookImage: false,
		twitterImage: false,
		errors: {},
		showImageModal: false,
		useFacebookGlobal: false,
		useTwitterGlobal: false,
		facebook: {
			"title": "",
			"description": "",
			"image_id": null,
		},
		twitter: {
			"title": "",
			"description": "",
			"image_id": null,
		},
	}),
	mounted() {
		this.getMediaById('facebook', this.meta['meta_facebook_image_id']);
		this.getMediaById('twitter', this.meta['meta_twitter_image_id']);
	},
	created() {
		this.isGlobal();
	},
	methods: {
		/*
		 * insertImage()
		 * Insert facebook or twitter image & update the height.
		 * Close the image modal.
		 */
		insertImage(e) {
			if (this.selectedImageType === "facebook") {
				this.facebookImage = e;
				this.meta['meta_facebook_image_id'] = e.id;
				this.$emit('update', e.id, 'meta_facebook_image_id');
				this.$nextTick(() => {
					this.helpers.setHeight(this.$refs.facebook.closest(".collapse-content"));
				})
			} else {
				this.twitterImage = e;
				this.meta['meta_twitter_image_id'] = e.id;
				this.$emit('update', e.id, 'meta_twitter_image_id');
				this.$nextTick(() => {
					this.helpers.setHeight(this.$refs.twitter.closest(".collapse-content"));
				})
			}
			this.showImageModal = false;
		},
		/*
		 * showModal()
		 */
		showModal(type) {
			this.selectedImageType = type;
			this.showImageModal = true;
		},
		/*
		 * removeImage()
		 * Remove the image from the global meta dependant on type.
		 */
		removeImage(type) {
			if (type === "facebook") {
				this.facebookImage = false;
				this.meta['meta_facebook_image_id'] = null;
			} else {
				this.twitterImage = false;
				this.meta['meta_twitter_image_id'] = null;
			}
		},
		/*
		 * getMediaById()
		 * Checks if the id exists, and proceeds to obtain the
		 * face/book image & sets vars.
		 */
		getMediaById(type, id) {
			if (!id) return;
			this.axios.get('/media/' + id)
				.then(res => {
					if (type === "facebook") {
						this.facebookImage = res.data.data;
					} else {
						this.twitterImage = res.data.data;
					}
				})
		},
		/*
		 * isGlobal()
		 * Checks to see if the facebook/twitter meta information
		 * is the same as the global.
		 */
		isGlobal() {
			if (this.meta['meta_title'] === this.meta['meta_facebook_title'] && this.meta['meta_description'] === this.meta['meta_facebook_description'] && this.meta['meta_facebook_title'] !== "" && this.meta['meta_facebook_description'] !== "") {
				this.useFacebookGlobal = true;
			}
			if (this.meta['meta_title'] === this.meta['meta_twitter_title'] && this.meta['meta_description'] === this.meta['meta_twitter_description'] && this.meta['meta_twitter_title'] !== "" && this.meta['meta_twitter_description'] !== "") {
				this.useTwitterGlobal = true;
			}
		},
		/*
		 * updateGlobal()
		 * Checks to see if the use global is truthy and
		 * updates twitter/facebook objects.
		 * Emits data to parent.
		*/
		updateGlobal(type) {
			if (type === "facebook") {
				if (!this.useFacebookGlobal) {
					this.$set(this.meta, "meta_facebook_title", this.facebook['title'])
					this.$set(this.meta, "meta_facebook_description", this.facebook['description'])
					this.facebook = {
						"title" : "",
						"description" : ""
					}
				} else {
					this.$set(this.facebook, "title", this.meta['meta_facebook_title'])
					this.$set(this.facebook, "description", this.meta['meta_facebook_description'])
					this.$set(this.meta, "meta_facebook_title", this.meta['meta_title'])
					this.$set(this.meta, "meta_facebook_description", this.meta['meta_description'])
				}

				this.$emit('update', this.meta['meta_facebook_title'], 'meta_facebook_title');
				this.$emit('update', this.meta['meta_facebook_description'], 'meta_facebook_description');
			} else if (type === "twitter") {
				if (!this.useTwitterGlobal) {
					this.$set(this.meta, "meta_twitter_title", this.twitter['title'])
					this.$set(this.meta, "meta_twitter_description", this.twitter['description'])
					this.twitter = {
						"title" : "",
						"description" : ""
					}
				} else {
					this.$set(this.twitter, "title", this.meta['meta_twitter_title'])
					this.$set(this.twitter, "description", this.meta['meta_twitter_description'])
					this.$set(this.meta, "meta_twitter_title", this.meta['meta_title'])
					this.$set(this.meta, "meta_twitter_description", this.meta['meta_description'])
				}

				this.$emit('update', this.meta['meta_twitter_title'], 'meta_twitter_title');
				this.$emit('update', this.meta['meta_twitter_description'], 'meta_twitter_description');
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">


.meta {

	// Toggle
	// =========================================================================

	&-toggle {
		display: flex;
		justify-content: space-between;
		margin-bottom: 2rem;
		margin-top: 10px;
	}

	// Card
	// =========================================================================

	&-card {

		.btn {
			margin-top: 10px;
		}

		input {
			width: 100%;
		}

		.image {
			margin-top: 10px;
		}

		&-col-left {
			padding: 0;
		}

		&-col-right {
			padding: 0;
		}


		&-facebook,
		&-twitter {

			input {
				width: 100% !important;
			}
		}

		@include media-desk {

			&-col-left {
				padding-right: 15px;
			}

			&-col-right {
				padding-left: 15px;
			}
		}
	}
}


</style>