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
				<figure>
					<img :src="getSiteUrl + media['url']">
					<div class="image-buttons">
						<i class="feather feather-edit image-buttons-choose" @click="showImageModal = true"></i>
						<i class="feather feather-trash image-buttons-remove" @click="remove"></i>
					</div>
				</figure>
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

export default {
	name: "FieldImage",
	components: {
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
				url: e.url,
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
					console.log(res.data.data)
					this.$set(this.media, {})
					this.media = res.data.data;
				})
				.catch(() => {
					this.value = false;
				})
				.finally(() => {
					//console.log("don")
				})

		},
		remove() {
			this.value = false;
			this.media = {};
		}
	},
	computed: {
		/*
		 * getSiteUrl()
		 * Get the site url from the store for previewing.
		 */
		getSiteUrl() {
			return this.$store.state.site.url;
		},
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

.image {
	$self: &;

	margin-top: 1.6rem;

	// Wrapper
	// =========================================================================

	&-wrapper {
		display: flex;
		flex-direction: column;
	}

	// Figure
	// =========================================================================

	figure {
		position: relative;
		height: 250px;
		width: 100%;
		border-radius: 4px;
		overflow: hidden;
		margin-bottom: 2rem;

		img {
			width: 100%;
			height: 100%;
			object-fit: cover;
			transition: transform 250ms ease;
			will-change: transform;
		}

		&:after {
			content: "";
			position: absolute;
			width: 100%;
			height: 100%;
			top: 0;
			left: 0;
			z-index: 99;
			background-color: rgba($black, 0.3);
			opacity: 0;
			transition: opacity 250ms ease;
			will-change: opacity;
		}

		&:hover {

			img {
				transform: scale(1.1);
			}

			#{$self}-buttons,
			&:after {
				opacity: 1;
			}
		}
	}

	// Buttons
	// =========================================================================

	&-buttons {
		position: absolute;
		display: flex;
		top: 20px;
		right: 20px;
		z-index: 100;
		opacity: 0;
		transition: opacity 250ms ease;
		will-change: opacity;

		i {
			color: $white;
			background-color: rgba($white, 1);
			padding: 10px;
			border-radius: 4px;
			cursor: pointer;
			box-shadow: none;
			transition: box-shadow 100ms ease;
			will-change: box-shadow;

			&:first-of-type {
				margin-right: 5px;
			}

			&:hover {
				box-shadow: 0 3px 10px 0 rgba($black, 0.14);
			}
		}

		&-remove {
			color: $orange !important;
		}

		&-choose {
			color: $green !important;
		}
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

		figure {
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

		figure {
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