<!-- =====================
	Field - Image
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="image">
		<div class="image">
			<!-- =====================
				No Image
				===================== -->
			<div v-if="!value">
				<button class="btn" @click="showImageModal = true;">Add image</button>
			</div>
			<!-- =====================
				With Image
				===================== -->
			<div v-else class="image-wrapper">
				<div class="image-cont">
					<ImageWithActions class="image-cover" @choose="showImageModal = true" @remove="remove">
						<img :src="getSiteUrl + media['url']">
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
							<p><a :href="getSiteUrl + value.url" target="_blank">{{ value.url }}</a></p>
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
import ImageWithActions from "@/components/misc/ImageWithActions";

export default {
	name: "FieldImage",
	components: {
		ImageWithActions,
		Modal,
		Uploader
	},
	mixins: [mediaMixin],
	props: {
		layout: Object,
		fields: {
			type: [Object, Boolean],
			default: false,
		},
	},
	data: () => ({
		errors: [],
		media: {},
		showImageModal: false,
		selectedImage: {},
	}),
	mounted() {
		if (this.value) {
			this.getMediaById();
		}
	},
	methods: {
		validate() {

		},
		insertMedia(e) {
			this.value = {
				id: e.id,
				type: "image"
			};
			this.$nextTick(() => {
				this.showImageModal = false;
				this.getMediaById();
				setTimeout(() => {
					this.helpers.setHeight(this.$refs.image.closest(".collapse-content"));
				}, 200);
			});
		},
		getMediaById() {
			this.axios.get('/media/' + this.value.id)
				.then(res => {
					this.$set(this.media, {})
					this.media = res.data.data;
				})
				.catch(() => {
					this.value = false;
				})
		},
		remove() {
			this.value = false;
			this.media = {};
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getLayout() {
			return this.layout;
		},
		value: {
			get() {
				return this.fields;
			},
			set(value) {
				this.$emit("update:fields", value)
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