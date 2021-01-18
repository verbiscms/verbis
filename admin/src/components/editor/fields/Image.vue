<!-- =====================
	Field - Image
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="image">
		<div class="image">
			<!-- =====================
				No Image
				===================== -->
			<div v-if="!field">
				<button class="btn" @click="showImageModal = true;">Add image</button>
			</div>
			<!-- =====================
				With Image
				===================== -->
			<div v-else class="image-wrapper">
				<div class="image-cont">
					<ImageWithActions class="image-cover" @choose="showImageModal = true" @remove="remove">
						<img :src="getSiteUrl + media['url']" @error="handleError">
					</ImageWithActions>
				</div>
				<div class="image-content">
					<div class="image-text">
						<div class="image-header">
							<h3>Information</h3>
						</div>
						<!-- Url -->
						<div class="text-cont">
							<h6>Url:</h6>
							<p><a :href="getSiteUrl + field.url" target="_blank">{{ field.url }}</a></p>
						</div>
						<!-- Title -->
						<div class="text-cont" v-if="media['title']">
							<h6>Title</h6>
							<p>{{ media['title'] }}</p>
						</div>
						<!-- Alt Text -->
						<div class="text-cont" v-if="media['alt']">
							<h6>Alternative text</h6>
							<p>{{ media['alt'] }}</p>
						</div>
						<!-- Description -->
						<div class="text-cont" v-if="media['description']">
							<h6>Description</h6>
							<p>{{ media['description'] }}</p>
						</div>
						<!-- File Size -->
						<div class="text-cont">
							<h6>Filesize:</h6>
							<p>{{ formatBytes(media['file_size']) }}</p>
						</div>
						<!-- File Size -->
						<div class="text-cont">
							<h6>Uploaded at:</h6>
							<p>{{ media['created_at'] | moment("dddd, MMMM Do YYYY") }}</p>
						</div>
					</div>
				</div>
			</div>
			<!-- =====================
				Modal
				===================== -->
			<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
				<template slot="text">
					<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertMedia">
						<template slot="close">
							<button class="btn btn-margin-right btn-icon-mob" @click="showImageModal = false">
								<i class="feather feather-x"></i>
								<span>Close</span>
							</button>
						</template>
					</Uploader>
				</template>
			</Modal>
		</div>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Uploader from "@/components/media/Uploader";
import Modal from "@/components/modals/General";
import { mediaMixin } from "@/util/media"
import {fieldMixin} from "@/util/fields/fields"
import ImageWithActions from "@/components/misc/ImageWithActions";

export default {
	name: "FieldImage",
	components: {
		ImageWithActions,
		Modal,
		Uploader
	},
	mixins: [mediaMixin,fieldMixin],
	data: () => ({
		media: {},
		showImageModal: false,
		selectedImage: {},
	}),
	mounted() {
		if (this.getValue !== "") {
			this.getMediaById();
		}
	},
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			this.validateRequired();
		},
		/*
		 * validateRequired()
		 * Return a error message if the options are required & the value is nil.
		 */
		validateRequired() {
			if (this.field === false && this.getLayout["required"] === true) {
				this.errors.push(`The ${this.getLayout.label.toLowerCase()} field is required.`);
			}
		},
		/*
		 * handleError()
		 * If there was an error getting the media file,
		 * defaults will be set and no broken image will be displayed.
		 */
		handleError() {
			this.media = {};
			this.field = "";
		},
		/*
		 * insertMedia()
		 * Insert a new media item when clicked.
		 */
		insertMedia(e) {
			this.field = e.id;
			this.$nextTick(() => {
				this.showImageModal = false;
				this.getMediaById();
				setTimeout(() => {
					this.helpers.setHeight(this.$refs.image.closest(".collapse-content"));
				}, 200);
			});
		},
		/*
		 * remove()
		 * Remove's a media item when clicked.
		 */
		getMediaById() {
			this.axios.get('/media/' + this.field)
				.then(res => {
					this.media = res.data.data;
				})
				.catch(() => {
					this.field = "";
				});
		},
		/*
		 * remove()
		 * Remove's a media item when clicked.
		 */
		remove() {
			this.field = false;
			this.media = {};
		}
	},
	computed: {
		/*
		 * field()
		 * Cast the ID to a string on emitting
		 * and parses the ID back to an integer when retrieving.
		 * Fire's back up to the parent. Check when setting the value
		 * for an empty string IS REQUIRED.
		 */
		field: {
			get() {
				return this.getValue !== "" ? parseInt(this.getValue) : false;
			},
			set(value) {
				if (value !== "") {
					this.$emit("update:fields", this.getFieldObject(value.toString()));
				}
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.field-cont {
	margin-top: 1.6rem;
}

.image {
	$self: &;

	// Wrapper
	// =========================================================================

	&-wrapper {
		display: flex;
		flex-direction: column;
	}

	// Container
	// =========================================================================

	&-cont {
		height: 250px;
		width: 100%;
	}

	// Header
	// =========================================================================

	&-header {
		padding-bottom: 1rem;
		margin-bottom: 1.6rem;
		border-bottom: 1px solid $grey-light;
	}

	// Tablet
	// =========================================================================

	&-content {

		.text-cont:last-of-type {
			margin-bottom: 0 !important;
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {

		&-wrapper {
			flex-direction: row;
		}

		&-cont {
			height: 400px;
			width: 65%;
			margin-bottom: 0;
		}

		&-content {
			flex-grow: 2;
			padding: 0 20px;
		}
	}

	// Desktop
	// =========================================================================

	@include media-desk {

		&-cont {
			height: 450px;
			width: 65%;
		}

		&-content {
			display: flex;
			flex-direction: column;
			justify-content: space-between;
		}
	}
}

</style>