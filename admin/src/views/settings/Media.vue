<!-- =====================
	Settings - Media
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Media Settings</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : saving }">Update settings</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					General
					===================== -->
				<div class="col-12">
					<h6 class="margin">General options</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Maximum size -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['media_upload_max_size']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Maximum size</h4>
										<p>Set the maximum size (in bytes) of a media item allowed to be uploaded to the library.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Max Size -->
									<FormGroup label="Maximum size" :error="errors['media_upload_max_size']">
										<input class="form-input form-input-white" type="number" v-model.number="maxSize">
										<p>If the maximum size is set to 0, no upload file size restrictions will apply.</p>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Maximum size -->
						<!-- Maximum dimensions -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['media_upload_max_width'] || errors['media_upload_max_height'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Maximum dimensions</h4>
										<p>Set the maximum width & height (in pixels) of a media item allowed to be uploaded to the library.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Max Width -->
									<FormGroup label="Maximum width*" :error="errors['media_upload_max_width']">
										<input class="form-input form-input-white" type="number" v-model.number="maxWidth">
										<p>If the maximum width is set to 0, no upload restrictions will apply.</p>
									</FormGroup>
									<!-- Max Height -->
									<FormGroup label="Maximum height*" :error="errors['media_upload_max_height']">
										<input class="form-input form-input-white" type="number" v-model.number="maxHeight">
										<p>If the maximum height is set to 0, no upload restrictions will apply.</p>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Maximum dimensions -->
						<!-- Compression -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['media_compression']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Compression</h4>
										<p>Set the image compression from 0 to 100.</p>
										<p>100 being the most amount of compression & 0 being none.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Compression -->
									<FormGroup label="Media compression*" :error="errors['media_compression']">
										<input class="form-input form-input-white" type="text" v-model.number="compression">
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Compression -->
						<!-- Organise by date -->
						<div class="collapse-border-bottom">
							<div class="card-header">
								<div>
									<h4 class="card-title">Organise by date</h4>
									<p>By ticking the box, the Verbis server will organise media items by year & month, e.g. /uploads/2020/01</p>
								</div>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="media-size-year" v-model="data['media_organise_year_month']" :true-value="true" :false-value="false" />
									<label for="media-size-year"></label>
								</div>
							</div><!-- /Card Header -->
						</div><!-- /Organise by date -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					WebP
					===================== -->
				<div class="col-12">
					<h6 class="margin">WebP Options</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Convert Webp -->
						<div class="collapse-border-bottom">
							<div class="card-header">
								<div>
									<h4 class="card-title">Convert to Webp's</h4>
									<p>By ticking the box, the Verbis media library will automatically convert Jpg's & Png's to Webp's on upload.</p>
								</div>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="media-convert-webp" v-model="data['media_convert_webp']" :true-value="true" :false-value="false" />
									<label for="media-convert-webp"></label>
								</div>
							</div><!-- /Card Header -->
						</div><!-- /Convert Webp -->
						<!-- Serve Webp -->
						<div class="collapse-border-bottom">
							<div class="card-header card-header-block">
								<div>
									<h4 class="card-title">Serve Webp images</h4>
									<p>By ticking the box, the Verbis server will automagically serve Webp images if the browser supports it.</p>
								</div>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="media-serve-webp" v-model="data['media_serve_webp']" :true-value="true" :false-value="false" />
									<label for="media-serve-webp"></label>
								</div>
							</div><!-- /Card Header -->
						</div><!-- /Serve Webp -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Sizes
					===================== -->
				<div class="col-12">
					<h6>Image Sizes</h6>
					<p>The image sizes determines the maximum dimensions in pixels when an image is uploaded to the media library.</p>
					<transition name="trans-fade" mode="out-in">
						<div v-if="sizes.length" class="card card-small-box-shadow card-expand card-margin-none">
							<!-- Sizes -->
							<Collapse v-for="(size, sizeKey) in sizes" :key="size.id" :show="false" class="collapse-border-bottom">
								<template v-slot:header>
									<div class="card-header">
										<h4 class="card-title">{{ size.name }}</h4>
										<div class="card-controls">
											<i class="feather feather-trash-2" @click="deleteSize(sizeKey)"></i>
											<i class="feather feather-chevron-down"></i>
										</div>
									</div><!-- /Card Header -->
								</template>
								<template v-slot:body>
									<div class="card-body media-sizes" ref="sizes">
										<FormGroup label="Key">
											<input class="form-input" :id="'media-size-key-' + sizeKey" type="text" placeholder="Enter a key" v-model="sizes[sizeKey].key">
											<p>Set the key used to access the media size in the templates.</p>
										</FormGroup>
										<FormGroup label="Name">
											<input class="form-input" :id="'media-size-name-' + sizeKey" type="text" placeholder="Enter a key" v-model="sizes[sizeKey].name">
											<p>Set the friendly name for the media size.</p>
										</FormGroup>
										<FormGroup label="Width">
											<input class="form-input" :id="'media-size-width-' + sizeKey" type="number" placeholder="Enter a width" v-model.number="sizes[sizeKey].width">
											<p>Set the the width (in pixels) of the image size.</p>
										</FormGroup>
										<FormGroup label="Height">
											<input class="form-input" :id="'media-size-height-' + sizeKey" type="number" placeholder="Enter a height" v-model.number="sizes[sizeKey].height">
											<p>Set the the height (in pixels) of the image size.</p>
										</FormGroup>
										<FormGroup>
											<div class="form-checkbox">
												<input type="checkbox" :id="'media-size-crop-' + sizeKey" v-model="sizes[sizeKey].crop" :true-value="true" :false-value="false" />
												<label :for="'media-size-crop-' + sizeKey">
													<i class="fal fa-check"></i>
												</label>
												<div class="form-checkbox-text">Crop image?</div>
											</div>
										</FormGroup>
									</div><!-- /Card Body -->
								</template>
							</Collapse><!-- /Sizes -->
						</div><!-- /Card -->
						<div v-else>
							<Alert type="warning" colour="orange">
								No media sizes found, consider making some to use with the <code v-text="'<picture>'"></code> element for increase speed.
							</Alert>
						</div>
					</transition>
					<div class="media-btn-cont">
						<button class="btn" @click="addImageSize">Add image size</button>
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
import {optionsMixin} from "@/util/options";
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import {userMixin} from "@/util/users";
import Alert from "@/components/misc/Alert";

export default {
	name: "Media",
	title: 'Media Settings',
	mixins: [optionsMixin, userMixin],
	components: {
		Alert,
		FormGroup,
		Breadcrumbs,
		Collapse
	},
	data: () => ({
		sizes: [],
		errorMsg: "Fix the errors before saving media settings.",
		successMsg: "Media options updated successfully.",
		axiosTimeout: true,
	}),
	methods: {
		/*
		 * runAfter()
		 * Process the image sizes once axios has finished loading.
		 */
		runAfterGet() {
			this.processImageSizes();
		},
		/*
		 * runBeforeSave()
		 * Process the image sizes once before posting.
		 */
		runBeforeSave() {
			let tempSizes = {};

			this.sizes.forEach((s, index) => {
				let key =  s.key === undefined || s.key === "" ?  `media-size-${index + 1}` : s.key;
				const name = s.name === "" || s.name == "Enter size" ?  `Media size ${index + 1}` : s.name;

				tempSizes[key] = tempSizes[key] || {};
				const obj = {
					name: name,
					width: s.width,
					height: s.height,
					crop: s.crop,
				};
				tempSizes[key] = obj;

				obj.id = s.id;
				obj.key = key;
				this.$set(this.sizes, index, obj);
			});

			// Set null if there are no sizes
			if (this.helpers.isEmptyObject(tempSizes)) {
				tempSizes = null
			}

			// Save for API
			for (const key in tempSizes) {
				delete tempSizes[key].id;
				delete tempSizes[key].key;
			}
			this.$delete(this.data, 'media_images_sizes');
			this.$set(this.data, 'media_images_sizes', tempSizes);
		},
		/*
		 * processImageSizes()
		 */
		processImageSizes() {
			for (const sizeName in this.data['media_images_sizes']) {
				const size = this.data['media_images_sizes'][sizeName];
				this.sizes.push({
					key: sizeName,
					name: size.name,
					width: size.width,
					height: size.height,
					crop: size.crop,
					id: this.createPassword(),
				});
			}
			this.sortImageSizes();
		},
		/*
		 * sortSizes()
		 * Sort sizes by width for the size cards.
		 */
		sortImageSizes() {
			this.sizes.sort((a, b) => parseFloat(a.width) - parseFloat(b.width));
		},
		/*
		 * addImageSize()
		 * Add an image to the image sizes array and expand the el.
		 */
		addImageSize() {
			this.sizes.push({
				crop: false,
				width: 0,
				height: 0,
				name: "Enter size",
			});
			this.$nextTick(function() {
				const newSize = this.$refs.sizes[this.sizes.length - 1];
				this.helpers.setHeight(newSize.closest(".collapse-content"));
			});
		},
		deleteSize(index) {
			this.sizes.splice(index, 1);
		},
	},
	computed: {
		/*
		 * maxSize()
		 * Change the value from an empty string to null if not set.
		 */
		maxSize: {
			get() {
				return this.data['media_upload_max_size'];
			},
			set(value) {
				console.log(value);
				if (value === "") {
					this.$set(this.data, 'media_upload_max_size', null);
					return
				}
				this.$set(this.data, 'media_upload_max_size', value);
			}
		},
		/*
		 * maxWidth()
		 * Change the value from an empty string to null if not set.
		 */
		maxWidth: {
			get() {
				return this.data['media_upload_max_width'];
			},
			set(value) {
				if (value === "") {
					this.$set(this.data, 'media_upload_max_width', null);
				} else {
					this.$set(this.data, 'media_upload_max_width', value);
				}
			}
		},
		/*
	 	 * maxHeight()
		 * Change the value from an empty string to null if not set.
		 */
		maxHeight: {
			get() {
				return this.data['media_upload_max_height'];
			},
			set(value) {
				if (value === "") {
					this.$set(this.data, 'media_upload_max_height', null);
					return;
				}
				this.$set(this.data, 'media_upload_max_height', value);
			}
		},
		/*
		 * compression()
		 * Change the value from an empty string to null if not set.
		 */
		compression: {
			get() {
				return this.data['media_compression'];
			},
			set(value) {
				console.log(value);
				if (value === "") {
					this.$set(this.data, 'media_compression', null);
					return
				}
				this.$set(this.data, 'media_compression', value);
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.media {

	$self: &;

	// Button Container
	// =========================================================================

	&-btn-cont {
		margin-top: 2rem;
		display: flex;
		justify-content: flex-end;
	}
}

</style>