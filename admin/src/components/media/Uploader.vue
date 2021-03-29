<!-- =====================
	Media
	===================== -->
<template>
	<section class="media">
		<!-- =====================
			Tabs
			===================== -->
		<Tabs v-if="filters" @update="filterTabs">
			<template slot="item">Show all</template>
			<template slot="item">JPG's</template>
			<template slot="item">PNG's</template>
			<template slot="item">Files</template>
		</Tabs>
		<!-- =====================
			Insert
			===================== -->
		<div class="media-insert" v-if="modal">
			<h2>Insert media item</h2>
			<div>
				<button v-if="selectMultiple" class="btn btn-margin-right" @click="bulkMode = true">Select multiple</button>
				<slot name="close"></slot>
				<button class="btn btn-green btn-icon-mob" @click="insertItem">
					<i class="feather feather-check"></i>
					<span>Insert</span>
				</button>
			</div>
		</div>
		<!-- Input -->
		<input class="media-input" id="browse-file" type="file" multiple ref="file" @change="addFile($event, true)">
		<!-- Spinner -->
		<div v-show="doingAxios && loadingImages" class="media-spinner spinner-container">
			<div class="spinner spinner-large spinner-grey"></div>
		</div>
		<div v-show="!doingAxios" class="row media-row trans-fade-in-anim">
			<!-- =====================
				Editor
				===================== -->
			<div v-if="options" class="col-12 col-desk-4 col-hd-3 order-desk-last media-col media-side">
				<div v-if="selectedMedia">
					<div class="media-options card card-small-box-shadow card-expand card-expand-full-width card-margin-none">
						<!-- Options -->
						<Collapse  class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Options</h3>
									<div class="card-controls">
										<i class="feather feather-trash-2" @click.prevent="showDeleteModal = true" :class="{ 'btn-loading' : isDeleting }"></i>
										<i class="feather feather-chevron-down"></i>
									</div>
								</div>
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Title -->
									<div class="form-group">
										<label for="media-title" class="form-label">Title</label>
										<input id="media-title" type="text" class="form-input form-input-white" v-model="selectedMedia.title" @keyup="save">
									</div><!-- /Title -->
									<!-- Alt Text -->
									<div class="form-group" v-if="getMediaType(selectedMedia.type) !== 'file'">
										<label for="media-alt" class="form-label">Alternative text</label>
										<input id="media-alt" type="text" class="form-input form-input-white" v-model="selectedMedia.alt" @keyup="save">
									</div><!-- /Alt Text -->
									<!-- Description -->
									<div class="form-group">
										<label for="media-description" class="form-label">Description</label>
										<input id="media-description" type="text" class="form-input form-input-white" v-model="selectedMedia.description" @keyup="save">
									</div><!-- /Description -->
								</div>
							</template>
						</Collapse><!-- /Options -->
						<!-- Editor -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Information</h3>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div>
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Url -->
									<div class="text-cont">
										<h6>Url:</h6>
										<p><a :href="getSiteUrl + selectedMedia.url" target="_blank">{{ selectedMedia.url }}</a></p>
									</div>
									<!-- Filesize -->
									<div class="text-cont">
										<h6>Filesize:</h6>
										<p>{{ formatBytes(selectedMedia['file_size']) }}</p>
									</div>
									<!-- Uploaded by -->
									<div class="text-cont">
										<h6>Uploaded by:</h6>
										<p>{{ selectedMedia['uploaded_by']['full_name'] }}</p>
									</div>
									<!-- Type -->
									<div class="text-cont">
										<h6>Type:</h6>
										<p>{{ selectedMedia.type }}</p>
									</div>
									<!-- Uploaded at -->
									<div class="text-cont text-cont-no-margin">
										<h6>Uploaded at:</h6>
										<p>{{ selectedMedia['created_at'] | moment("dddd, MMMM Do YYYY") }}</p>
									</div>
								</div>
							</template>
						</Collapse><!-- /Editor -->
						<!-- Sizes -->
						<Collapse  v-if="selectedMedia.sizes"  :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Sizes</h3>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div>
							</template>
							<template v-slot:body>
								<div class="card-body" >
									<div class="media-size" v-for="size in sortSizes(selectedMedia.sizes)" :key="size.uuid">
										<div class="media-size-header">
											<h4>{{ size['size_name'] }}</h4>
											<div class="badge badge-green">{{ formatBytes(size['file_size']) }}</div>
										</div>
										<div class="media-size-body">
											<p><span>Crop:</span> {{ size.crop }}</p>
											<p><span>Url:</span> {{ size.url }}</p>
											<p><span>Width:</span> {{ size.width }}px</p>
											<p><span>Height:</span> {{ size.height }}px</p>
										</div>
									</div>
								</div>
							</template>
						</Collapse><!-- /Sizes -->
					</div>
				</div><!-- /Wrapper -->
				<div v-else-if="!selectedMedia && media.length" class="trans-fade-in-anim">
					<div class="card card-small-box-shadow media-select-card">
						<div class="card-body">
							<i class="feather feather-edit"></i>
							<div>
								<h4>No file selected</h4>
								<p>Select media to edit & view information</p>
							</div>
						</div>
					</div>
				</div>
			</div><!-- /Col -->
			<!-- =====================
				Media Items
				===================== -->
			<div class="col-12 media-col media-col-item" :class="{ 'col-desk-8 col-hd-9' : media.length && options }" @dragover.prevent.stop="handleDrag">
				<transition name="trans-fade-quick">
					<div class="media-dragging" v-if="dragging" :class="{ 'media-dragging-centered' : !media.length || media.length < 15 }" @drop.prevent.stop="addFile($event, false)" @dragexit.stop="dragging = false" @dragleave.stop="dragging = false" @mouseleave="dragging = false">
						<i class="feather feather-upload-cloud"></i>
						<h4>Drop!</h4>
						<p>Drop files to upload them instantly to the media library.</p>
					</div>
				</transition>
				<div class="media-files">
					<!-- Placeholder -->
					<div v-if="!media.length && !doingAxios" class="media-placeholder">
						<i class="feather feather-image"></i>
						<h4>No media items found!</h4>
						<p>Drag and drop files here or click the button above.</p>
					</div>
					<!-- Media -->
					<div v-else class="media-item" v-for="(item, itemIndex) in media" :key="item.uuid" @click.prevent.stop="handleMediaClick(item)"
						:class="{ 'media-item-active' : selectedMedia && selectedMedia['uuid'] === item['uuid'],
						'media-item-plain' : item.loading,
						'media-item-bulk' : checked.includes(item.id),
						'media-item-icon' : getMediaType(item.type) !== 'image' && getMediaType(item.type) !== 'video' || (item['unsupported']),
						'media-item-error' : item.loading && item['unsupported'],
						'media-item-no-options' : !options }">
						<!-- Checkbox -->
						<div class="form-checkbox media-item-checkbox">
							<input type="checkbox" checked :id="'media-item-' + item.uuid"/>
							<label :for="'media-item-' + item.uuid">
								<i class="fal fa-check"></i>
							</label>
						</div>
						<!-- Uploading -->
						<div v-if="item.loading && !item['unsupported']" :key="item.uuid + '-loading'" class="media-item-trans">
							<div class="spinner spinner-grey spinner-large"></div>
							<h4>{{ item.name }}</h4>
						</div>
						<!-- Unsupported -->
						<div v-else-if="item['unsupported']" class="media-item-icon-cont media-item-trans">
							<i class="feather feather-alert-circle"></i>
							<p>{{ item['unsupported'] }}</p>
							<i class="media-close feather feather-x" @click="removeErrorItem(item, itemIndex)"></i>
						</div>
						<!-- Image -->
						<div v-else-if="getMediaType(item.type) === 'image'" class="media-item-image media-item-trans" ref="images">
							<img v-onload="getSiteUrl + item.url" :alt="item.alt" @loaded="loadImages($event)">
						</div>
						<!-- Video -->
						<div v-else-if="getMediaType(item.type) === 'video'" class="media-item-video media-item-trans">
							<video controls preload="none" disablepictureinpicture controlslist="nodownload">
								<source :src="getSiteUrl + item.url">
							</video>
						</div>
						<!-- File -->
						<div v-else-if="getMediaType(item.type) === 'file'" class="media-item-icon-cont media-item-trans">
							<i class="feather feather-file"></i>
							<p>{{ item['file_name'] }}</p>
						</div>
					</div><!-- /Media Item -->
				</div><!-- /Media Files -->
				<!-- =====================
					Pagination
					===================== -->
				<transition name="trans-fade">
					<div class="row" v-if="!doingAxios && paginationObj && media.length">
						<div class="col-12 media-col">
							<Pagination :pagination="paginationObj" @update="setPagination"></Pagination>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</transition>
			</div><!-- /Col -->
		</div><!-- /Row -->
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning">
			<template slot="button">
				<button class="btn" @click="deleteItem">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="!checked.length">Are you sure want to delete this media item?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} media items?</p>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Modal from "../../components/modals/General";
import Tabs from "../../components/misc/Tabs";
import Pagination from "@/components/misc/Pagination";
import {mediaMixin} from "@/util/media";
import Collapse from "@/components/misc/Collapse";

export default {
	name: "Uploader",
	mixins: [mediaMixin],
	props: {
		filters: {
			type: Boolean,
			default: false
		},
		bulkAction: {
			type: Boolean,
			default: false
		},
		deleting: {
			type: Boolean,
			default: false,
		},
		modal: {
			type: Boolean,
			default: false,
		},
		selectMultiple: {
			type: Boolean,
			default: false,
		},
		rows: {
			type: Number,
			default: 6,
		},
		options: {
			type: Boolean,
			default: true,
		}
	},
	components: {
		Collapse,
		Modal,
		Pagination,
		Tabs,
	},
	data: () => ({
		doingAxios: true,
		media: [],
		selectedMedia: false,
		users: [],
		filter: "",
		pagination: "",
		paginationObj: {},
		checked: [],
		activeAction: "",
		uploading: false,
		timeout: null,
		isDeleting: false,
		dragging: false,
		loadingImages: true,
		loadedImages: [],
		initial: true,
		showDeleteModal: false,
	}),
	mounted() {
		this.getMedia();
		this.getUsers();
	},
	created() {
		window.addEventListener("resize", this.resizeHandler);
	},
	destroyed() {
		window.removeEventListener("resize", this.resizeHandler);
	},
	watch: {
		deleting: function() {
			if (!this.checked.length) {
				this.$noty.warning("Select items in order to apply bulk actions");
				return;
			}
			this.showDeleteModal = true;
		},
		bulkAction: function(val) {
			if (!val) {
				this.checked = [];
			} else {
				this.selectedMedia = false;
			}
		}
	},
	methods: {
		/*
		 * getMedia()
		 * Obtain the media or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getMedia() {
			this.doingAxios = true;

			this.axios.get(`media?order_by=created_at&order_direction=desc&filter=${this.filter}&${this.pagination}&limit=${this.rows * 5}`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
				.then(res => {
					this.media = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					const media = res.data.data;
					if (media.length) {
						this.media = media;
					}
					if (!this.initialLoad) {
						setTimeout(() => {
							this.doingAxios = false;
						}, 200);
						return
					}
					this.doingAxios = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.initialLoad = true;
				});
		},
		/*
		 * addFile()
		 * Processes file depending on weather the file was dropped or inserted
		 * by the button from parent component.
		 */
		addFile(e, manual = false) {
			if (this.uploading) {
				this.$noty.warning("Wait for the other files to finish uploading")
				this.dragging = false;
				return;
			}

			this.uploading = true;
			this.dragging = false;

			let droppedFiles = manual ? e.target.files : e.dataTransfer.files;
			if (!droppedFiles) return;

			([...droppedFiles].forEach((file, index) => {
				this.media.unshift({
					name: file.name,
					loading: true,
					index: index,
				});
				this.upload(file, index);
			}));
		},
		/*
		 * upload()
		 * Upload file independently to the server.
		 */
		upload(file, index) {
			let formData = new FormData();
			formData.append('file', file);

			this.axios("/media", {
				method: 'post',
				data: formData,
				headers: {'Content-Type': 'multipart/form-data'}
			})
				.then(res => {
					setTimeout(() => {
						this.$set(this.media, index, res.data.data);
					}, Math.floor(Math.random() * (230 - 100 + 1)) + 100);
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 415) {
						setTimeout(() => {
							this.$set(this.media[index], "unsupported", err.response.data.message);
						}, Math.floor(Math.random() * (430 - 200 + 1)) + 200);
						return;
					}
					if (err.response.status === 413) {
						setTimeout(() => {
							this.$set(this.media[index], "unsupported", "File too large, check the server settings");
						}, Math.floor(Math.random() * (430 - 200 + 1)) + 200);
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.uploading = false;
				})
		},
		/*
		 * save()
		 * Save title, description & alt text on key up, throttle input.
		 */
		save() {
			clearTimeout(this.timeout);

			this.timeout = setTimeout(() => {
				this.axios.put("/media/" + this.selectedMedia.id, {
					title: this.selectedMedia.title,
					description: this.selectedMedia.description,
					alt: this.selectedMedia.alt,
				})
					.then(() => {
						this.$noty.success("Media item updated successfully.");
					})
					.catch(err => {
						this.helpers.handleResponse(err);
					});
			}, 1000);
		},
		/*
		 * deleteItem()
		 * Delete media item and remove from media array.
		 */
		deleteItem() {
			this.isDeleting = true;

			let toDelete = [];
			if (this.selectedMedia) {
				toDelete.push(this.selectedMedia.id);
			} else {
				toDelete = this.checked;
			}

			const promises = [];
			toDelete.forEach(id => {
				promises.push(this.deleteItemAxios(id));
			});

			// Send all requests
			Promise.all(promises)
				.then(() => {
					this.$noty.success("Media item deleted successfully.");
					this.getMedia();
					this.$store.dispatch("getProfilePicture");
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					if (this.paginationObj) this.paginationObj.page = 0;
					this.selectedMedia = false;
					this.isDeleting = false;
					this.bulkMode = false;
					this.showDeleteModal = false;

					// Check if the file deleted was the site logo
					this.axios.get("/site").then(res => {
						this.$store.commit("setSite", res.data.data);
					});
				});
		},
		/*
		 * async deleteItemAxios()
		 */
		async deleteItemAxios(id) {
			return await this.axios.delete("/media/" + id);
		},
		/*
		 * handleDrag()
		 * Handler for when user drags item onto media uploader.
		 * Sets bulk mode to falsely.
		 */
		handleDrag() {
			this.dragging = true;
			this.bulkMode = false;
		},
		/*
		 * insertItem()
		 * Update parent component if modal, emit the selected media.
		 */
		insertItem() {
			if (!this.selectedMedia && (!this.checked.length && this.bulkMode)) {
				this.$noty.warning("Select media item to insert.");
				return;
			}
			if (this.checked.length) {
				const media = this.checked.map(m => this.findMediaById(m));
				this.$emit("insert", media);
				return;
			}
			this.$emit("insert", this.selectedMedia);
		},
		/*
		 * handleMediaClick()
		 * If bulk select is enabled, add or remove from the checked array
		 * depending on whether or not the item is in.
		 * Else, change & show the selected media panel.
		 */
		handleMediaClick(item) {
			if (item['unsupported']) {
				return;
			}
			if (this.bulkMode) {
				if (this.checked.includes(item.id)) {
					const index = this.checked.indexOf(item.id);
					this.checked.splice(index, 1);
					return;
				}
				this.checked.push(item.id)
			} else {
				this.changeSelectedMedia(item);
			}
		},
		/*
		 * loadImages()
		 * Push to loaded images array, if the images length is equal
		 * To the amount of images loaded, set the loading images
		 * to falsely.
		 */
		loadImages(e) {
			this.loadedImages.push(e);

			if (this.$refs.images.length === this.loadedImages.length) {
				if (this.initial) {
					setTimeout(() => {
						this.loadingImages = false;
						this.initial = false;
					}, this.timeoutDelay);
					return;
				}
				this.loadingImages = false;
			}
		},
		/*
		 * changeSelectedMedia()
		 * Change selected media item when an item is clicked.
		 * Change author details.
		 */
		changeSelectedMedia(item) {
			this.selectedMedia = {};
			const user = this.findUserById(item['user_id']);
			if (!user) {
				console.warn("No user attached to media item");
				return;
			}
			user['full_name'] = `${user['first_name']} ${user['last_name']}`;
			item['uploaded_by'] = user;
			this.selectedMedia = item;
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain media items.
		 */
		filterTabs(e) {
			let filter = ""
			switch (e) {
				case 2: {
					filter = '{"type":[{"operator":"=", "value": "image/jpeg" }]}';
					break;
				}
				case 3: {
					filter = '{"type":[{"operator":"=", "value": "image/png" }]}';
					break;
				}
				case 4: {
					filter = '{"type":[{"operator": "NOT LIKE", "value": "image/png"},{"operator": "NOT LIKE", "value": "image/jpeg"},{"operator": "NOT LIKE", "value": "image/gif"},{"operator": "NOT LIKE", "value": "image/webp"},{"operator": "NOT LIKE", "value": "image/bmp"},{"operator": "NOT LIKE", "value": "image/svg+xml"}]}'
				}
			}
			this.selectedMedia = false;
			this.filter = filter;
			this.getMedia();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain posts.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getMedia();
		},
		/*
		 * sortSizes()
		 * Sort sizes by width for the side panel.
		 */
		sortSizes(sizes) {
			return sizes;

			//return sizes.slice().sort((a, b) => parseFloat(a.width) - parseFloat(b.width));
		},
		/*
		 * removeErrorItem()
		 * If there is an error on the media item, allow the user to remove
		 * by deleting from the media array.
		 */
		removeErrorItem(item, index) {
			this.$delete(this.media, index);
		},
		/*
		 * getUsers()
		 * Obtain users from store, if none, dispatch users action.
		 */
		getUsers() {
			this.$store.dispatch("getUsers")
				.then(users => {
					this.users = users;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * findUserById()
		 */
		findUserById(id) {
			return this.users.find(u => u.id === id);
		},
		/*
		 * findMediaById()
		 */
		findMediaById(id) {
			return this.media.find(u => u.id === id);
		},
		/*
		 * resizeHandler()
		 * Clear bulk mode & checked array for mobile.
		 */
		resizeHandler() {
			if (window.innerWidth <= 568) {
				this.bulkMode = false;
				this.checked = [];
			}
		},
	},
	computed: {
		/*
		 * bulkMode()
		 * Get & set the bulk action to emit to parent.
		 */
		bulkMode: {
			get() {
				return this.bulkAction;
			},
			set(value) {
				this.$emit("update:bulk-action", value)
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

.media {
	$self: &;

	// Props
	// =========================================================================

	&-row {
		align-items: flex-start;
		max-width: none !important;
	}

	// Placeholder / Dragging Props
	// =========================================================================

	&-placeholder,
	&-dragging {
		height: calc(100% - 1rem);
		min-height: 100%;

		i {
			font-size: 46px;
			margin-bottom: 1rem;
			transition: color 100ms ease;
			will-change: color;
		}

		h4 {
			font-weight: 600;
			color: $copy;
			margin-bottom: 4px;
		}

		p {
			margin-bottom: 0;
		}
	}

	// Placeholder
	// =========================================================================

	&-placeholder {
		position: absolute;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 100%;
		min-height: 300px;
		border-radius: 10px;
		background-color: $white;
		border: 2px dashed $grey-light;
		transition: background-color 100ms ease;
		will-change: background-color;

		i {
			color: $orange;
		}
	}

	// Dragging
	// =========================================================================

	&-dragging {
		position: absolute;
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
		//width: calc(100% - 30px);
		background-color: rgba($white, 0.99);
		z-index: 9999999;
		border: 2px dashed $primary;
		transition: all 200ms;
		padding-top: 20%;
		border-radius: 10px;
		cursor: copy;

		i {
			color: $primary;
		}

		* {
			cursor: copy;
			pointer-events: none !important;
		}

		&-centered {
			padding-top: 0;
			justify-content: center;
		}
	}

	// Files
	// =========================================================================

	&-files {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		width: 100%;
		margin-bottom: -1rem;
		padding-bottom: 20px;
		border-radius: 10px;

		&:after {
			content: "";
			flex: auto;
		}
	}

	// Input
	// =========================================================================

	&-input {
		display: none;
		position: absolute;
		top: -9999999px;
		left: -9999999px;
	}

	// Item
	// =========================================================================

	&-item {
		position: relative;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-basis: calc(50% - 6px);
		height: 160px;
		background-color: transparent;
		margin-bottom: 1rem;
		border-radius: 6px;
		cursor: pointer;
		border: 2px solid rgba($black, 0);
		transition: all 100ms ease;
		will-change: opacity;

		// Image
		// =========================================================================

		&-image {
			display: flex;
			justify-content: center;
			align-items: center;
			width: 100%;
			height: 100%;
			margin: 0;
			border-radius: 4px;
			overflow: hidden;
			background-color: darken($bg-color, 2%);

			img {
				width: 100%;
				height: 100%;
				object-fit: cover;
			}
		}

		// Video
		// =========================================================================

		&-video {
			width: 100%;
			height: 100%;
			border-radius: 4px;
			overflow: hidden;
			border: none;
			background-color: transparent;

			video {
				display: block;
				height: 100%;
				width: 100%;
				border-radius: 4px;
				object-fit: contain;
				outline: none;
			}
		}

		// Icon
		// =========================================================================

		&-icon {
			border-color: $grey-light;
			padding: 0 10px;
			border-radius: 6px;

			&-cont {
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
			}

			h4 {
				text-align: center;
			}

			p,
			i {
				transition: color 100ms ease;
				will-change: color;
			}

			i {
				font-size: 26px;
				margin-bottom: 6px;
				color: $copy-light;
			}

			p {
				text-align: center;
				margin-bottom: 0;
				color: $copy-light;
			}
		}

		// Plain
		// =========================================================================

		&-plain {
			background-color: $white;
			box-shadow: $form-box-shadow;
			cursor: default;

			h4 {
				color: $copy;
				margin-top: 1rem;
				font-size: 0.86rem;
				user-select: none;
			}
		}

		// Error
		// =========================================================================

		&-error {
			border-color: $orange;

			#{$self}-close {
				position: absolute;
				top: 0;
				right: 0;
				padding: 10px;
				font-size: 18px;
				cursor: pointer;
			}

			i {
				color: $orange;
			}

			p {
				color: $orange;
			}
		}

		// Checkbox
		// =========================================================================

		&-checkbox {
			position: absolute;
			top: 0;
			right: 0;
			z-index: 9999;
			opacity: 0;
			transition: opacity 160ms ease;
			transform: translate(-50%, -50%);
		}

		// Bulk View
		// =========================================================================

		&-bulk {
			border-color: $primary;

			#{$self}-item-checkbox {
				opacity: 1;
			}
		}

		// Active (Clicked)
		// =========================================================================

		&-active {
			border-color: $primary;
			box-shadow: 0 0 10px 0 rgba($black, 0.14);


			p,
			i {
				color: $primary;
			}
		}
	}

	// Select Card
	// =========================================================================

	&-select-card {
	//	box-shadow: 0 3px 6px 0 rgba($black, 0.05);

		.card-body {
			display: flex;
			align-items: center;

			i {
				margin-right: 16px;
				color: $grey;
			}

			h4 {
				margin-bottom: -2px;
			}

			p {
				margin-bottom: 0;
			}
		}
	}

	// Options
	// =========================================================================

	&-side {
		margin-bottom: 0;
	}

	// Information
	// =========================================================================

	&-information {

		a {
			font-weight: normal;

		}
		.text-cont {
			//margin-bottom: 1.4rem;
		}
	}

	// Editor
	// =========================================================================

	&-information {

		h6 {
			color: $secondary;
			font-weight: 600;
		}
	}

	// Size (Image Size Card)
	// =========================================================================

	&-size {
		border-bottom: 1px solid $grey-light;
		background-color: $white;
		padding: 16px 0;

		&:first-of-type {
			border-top: 1px solid $grey-light;
			//padding-top: 10px;
		}

		&-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			width: 100%;
			//margin-bottom: 4px;
		}

		&-body {
			span {
				color: $secondary;
				font-weight: 600;
			}
		}

		p {
			margin-bottom: 4px;
			max-width: 100%;
			word-break: keep-all;
			text-overflow: ellipsis;

			&:last-of-type {
				margin-bottom: 0;
			}
		}

		&:last-of-type {
			border-bottom: 0;
			padding-bottom: 0;
		}
	}

	// Insert
	// =========================================================================

	&-insert {
		position: relative;
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding-bottom: 1rem;
		margin-bottom: 2rem;
		border-bottom: 1px solid $grey-light;

		h2 {
			margin-bottom: 0;
		}
	}

	// Modal
	// =========================================================================

	&-modal {
		display: block;
		height: 70vh;

		#{$self}-side {
			padding-bottom: 2rem;
		}

		.pagination {
			margin-bottom: 1rem;
		}
	}

	// Fade Anim
	// =========================================================================

	&-item-trans {
		animation: fade 300ms ease-in;

		@keyframes fade {

			from {
				opacity: 0;
			}

			to {
				opacity: 1;
			}
		}
	}

	// Tablet Down
	// =========================================================================

	@include media-tab-down {

		&-col {
			padding: 0;
		}

		&-item {
			&:last-child {
				margin-left: 6px;
			}
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {

		&-modal {

			.pagination {
				margin-bottom: 2rem;
			}
		}

		&-files {
			min-height: 300px;
		}

		&-item {
			flex-basis: calc(33.333333333% - 6px);
			height: 200px;
		}
	}

	// Desktop
	// =========================================================================

	@include media-desk {

		&-modal {
			height: 70vh;
			max-height: 825px;
		}

		&:not(&-modal) &-col-item {
			padding-left: 0;
		}

		&-files {
			min-height: 400px;
		}

		&-item {
			flex-basis: calc(33.33333% - 10px);
			margin-right: 10px;
			height: 180px;

			&-no-options {
				flex-basis: calc(25% - 15px);
			}
		}
	}

	// HD
	// =========================================================================

	@include media-hd {

		&-item {
			height: 220px;
			flex-basis: calc(20% - 15px);
			margin-right: 15px;

			&-no-options {
				flex-basis: calc(16.6666667% - 15px);
			}
		}
	}
}
</style>