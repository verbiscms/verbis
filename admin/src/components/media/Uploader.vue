<!-- =====================
	Media
	===================== -->
<template>
	<section class="media">
		<!-- =====================
			TODO:
				- Do the styling for bulk action and the checkbox
				- Tidy up the information sidebar (actions)
				- Fix the transition when going to a different tab.
				- Adding files from the plus button isnt working (the loading)
				- Look into the styling for the spinner when a media item is being uploaded
			===================== -->

		<!-- =====================
			Tabs
			===================== -->
		<div class="row" v-if="filters">
			<Tabs @update="filterTabs">
				<template slot="item">Show all</template>
				<template slot="item">JPG's</template>
				<template slot="item">PNG's</template>
				<template slot="item">Files</template>
			</Tabs>
		</div>
		<!-- Input -->
		<input class="media-input" id="browse-file" type="file" multiple ref="file" @change="addFile($event, true)">
		<!-- Spinner -->
		<div v-if="doingAxios" class="spinner-container">
			<div class="spinner spinner-large spinner-grey"></div>
		</div>
		<div v-else>
			<div class="row">
				<!-- =====================
					Editor
					===================== -->
				<div class="col-12 col-desk-3 order-desk-last">
					<div v-if="selectedMedia">
						<!-- Options -->
						<form class="form media-options">
							<h2>Options</h2>
							<!-- Title -->
							<div class="form-group">
								<label for="media-title" class="form-label">Title</label>
								<input id="media-title" type="text" class="form-input form-input-white" v-model="selectedMedia.title" @keyup="save">
							</div><!-- /Title -->
							<!-- Alt Text -->
							<div class="form-group">
								<label for="media-alt" class="form-label">Alternative text</label>
								<input id="media-alt" type="text" class="form-input form-input-white" v-model="selectedMedia.alt" @keyup="save">
							</div><!-- /Alt Text -->
							<!-- Description -->
							<div class="form-group">
								<label for="media-description" class="form-label">Description</label>
								<input id="media-description" type="text" class="form-input form-input-white" v-model="selectedMedia.description" @keyup="save">
							</div><!-- /Description -->
						</form><!-- /Options -->
						<!-- Editor -->
						<div class="media-information">
							<h2>Information</h2>
							<div class="text-cont">
								<h6>Url:</h6>
								<p>{{ selectedMedia.url }}</p>
							</div>
							<div class="text-cont">
								<h6>Filesize:</h6>
								<p>{{ formatBytes(selectedMedia['file_size']) }}</p>
							</div>
							<div class="text-cont">
								<h6>Uploaded By:</h6>
								<p>{{ selectedMedia['uploaded_by']['full_name'] }}</p>
							</div>
							<div class="text-cont">
								<h6>Type:</h6>
								<p>{{ selectedMedia.type }}</p>
							</div>
							<div class="text-cont">
								<h6>Uploaded at:</h6>
								<p>{{ selectedMedia['created_at'] | moment("dddd, MMMM Do YYYY") }}</p>
							</div>
							<div class="text-cont" v-if="selectedMedia.sizes.length">
								<h6>Sizes:</h6>
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
						</div><!-- /Editor -->
						<!-- Actions -->
						<div class="media-actions">
							<h2>Actions</h2>
							<button class="btn btn-orange" @click="deleteItem" :class="{ 'btn-loading' : isDeleting }">Delete</button>
						</div>
					</div><!-- /Wrapper -->
					<div v-else-if="!selectedMedia && media.length">
						<div class="card media-select-card">
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
				<div class="col-12" :class="{ 'col-desk-9' : media.length }">
					<div class="media-files" :class="{ 'media-files-selected' : selectedMedia, 'media-dragging' : dragging, 'media-dragging-empty' : !media.length && dragging }" @click.stop="selectedMedia = false" @drop.prevent="addFile($event, false)" @dragover.prevent @dragover="dragging = true" @dragleave="dragging = false" @dragend="dragging = false" @drop="dragging = false">
						<!-- Placeholder -->
						<div v-if="!media.length" class="media-placeholder">
							<i class="feather feather-image"></i>
							<h4>No media items found!</h4>
							<p>Drag and drop files here or click the button above.</p>
						</div>
						<!-- Media -->
						<div v-else class="media-item" v-for="item in media" :key="item.uuid" @click.prevent.stop="handleMediaClick(item)"
							:class="{ 'media-item-active' : selectedMedia && selectedMedia['uuid'] === item['uuid'], 'media-item-loading' : item.loading, 'media-item-bulk' : bulkDelete }">
							<!-- Checkbox -->
							<div class="form-checkbox media-item-checkbox">
								<input type="checkbox" checked :id="'media-item-' + item.uuid"/>
								<label :for="'media-item-' + item.uuid">
									<i class="fal fa-check"></i>
								</label>
							</div>
							<!-- Image -->
							<figure v-if="getMediaType(item.type) === 'image'" class="media-item-image">
								<img :src="getSiteUrl + item.url" :alt="item.alt">
							</figure>
							<!-- Video -->
							<div v-else-if="getMediaType(item.type) === 'video'">
								<i class="feather feather-video"></i>
							</div>
							<!-- File -->
							<div v-else-if="getMediaType(item.type) === 'file'">
								<i class="feather feather-file"></i>
							</div>
							<!-- Uploading -->
							<div v-else-if="item.loading" class="media-item-loading-cont">
								<transition name="trans-fade">
									<div v-if="item['unsupported']" class="media-item-loading-cont-error">
										<i class="feather feather-alert-circle"></i>
										<p>Media type unsupported</p>
									</div>
									<div v-else>
										<div class="spinner spinner-grey"></div>
										<h4>{{ item.name }}</h4>
									</div>
								</transition>
							</div>
						</div><!-- /Media Item -->
					</div><!-- /Media Files -->
					<!-- =====================
						Pagination
						===================== -->
					<div class="row">
						<div class="col-12">
							<transition name="trans-fade">
								<div class="row" v-if="!doingAxios && paginationObj">
									<div class="col-12">
										<Pagination :pagination="paginationObj" @update="setPagination"></Pagination>
									</div><!-- /Col -->
								</div><!-- /Row -->
							</transition>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Doing Axios -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Tabs from "../../components/misc/Tabs";
import Pagination from "@/components/misc/Pagination";

export default {
	name: "Uploader",
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
		}
	},
	components: {
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
		dragging: false,
		timeout: null,
		isDeleting: false,
	}),
	mounted() {
		this.getMedia();
		this.getUsers();
	},
	watch: {
		deleting: function() {
			this.checked.forEach(id => {
				this.deleteItem(id);
			});
			this.$noty.success("Media items deleted successfully.");
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

			this.axios.get(`media?order=${this.order}&filter=${this.filter}&${this.pagination}&limit=35`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
				.then(res => {
					console.log(res);
					this.media = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;

					const media = res.data.data;
					if (media.length) {
						this.media = media;
					}
				})
				.catch(err => {
					console.log(err)
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					setTimeout(() => {
						this.doingAxios = false;
					}, 100)
				});
		},
		/*
		 * addFile()
		 * Processes file depending on weather the file was dropped or inserted
		 * by the button from parent component.
		 */
		addFile(e, manual = false) {
			this.uploading = true;

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
					console.log(res.data.data);
					//this.media = res.data.data;
				})
				.catch(err => {
					console.log(err);
					console.log(index);
					if (err.response.status === 415) {
						this.$set(this.media[index], "unsupported", true)
					}
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
						console.log(err);
						this.$noty.error("Error occurred, please refresh the page.");
					});
			}, 1000);
		},
		/*
		 * deleteItem()
		 * Delete media item and remove from media array.
		 */
		deleteItem(id = false) {
			this.isDeleting = true;
			const isMultiple = id;
			if (!id) id = this.selectedMedia.id;

			this.axios.delete("/media/" + id)
				.then(() => {
					if (!isMultiple) {
						this.$noty.success("Media item deleted successfully.");
					}
					// Clear selected item
					this.selectedMedia = false;
					// Remove from media array
					const index = this.media.findIndex(m => m.id === id);
					this.media.splice(index, 1);
					return true;
				})
				.catch(err => {
					console.log(err);
					if (!isMultiple) {
						this.$noty.error("Error occurred, please refresh the page.");
					}
					return false;
				})
				.finally(() => {
					this.isDeleting = false;
					this.bulkDelete = false;
				})
		},
		/*
		 * handleMediaClick()
		 * If bulk select is enabled, add or remove from the checked array
		 * depending on whether or not the item is in.
		 * Else, change & show the selected media panel.
		 */
		handleMediaClick(item) {
			if (this.bulkDelete) {
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
		 * changeSelectedMedia()
		 * Change selected media item when an item is clicked.
		 * Change author details.
		 */
		changeSelectedMedia(item) {
			this.selectedMedia = {};
			const user = this.findUserById(item['user_id']);
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
					filter = '{"type":[{"operator": "!=", "value": "image/png"},{"operator": "!=", "value": "image/jpeg"},{"operator": "!=", "value": "image/gif"},{"operator": "!=", "value": "image/webp"},{"operator": "!=", "value": "image/bmp"},{"operator": "!=", "value": "image/svg+xml"}]}'
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
		 * getMediaType()
		 * Determines if the file is a image, video or file.
		 */
		getMediaType(type) {
			const images = ['image/png', 'image/jpeg', 'image/gif', 'image/webp', 'image/bmp', 'image/svg+xml'],
				video = ['video/mpeg', 'video/mp4', 'video/webm'];
			if (images.includes(type)) {
				return "image";
			} else if (video.includes(type)) {
				return "video";
			} else {
				return "file";
			}
		},
		/*
		 * formatBytes()
		 * Return formatted byte information for file size.
		 */
		formatBytes(bytes, decimals = 2) {
			if (bytes === 0) return '0 Bytes';

			const k = 1024,
				dm = decimals < 0 ? 0 : decimals,
				sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
				i = Math.floor(Math.log(bytes) / Math.log(k));

			return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
		},
		/*
		 * sortSizes()
		 * Sort sizes by width for the side panel.
		 */
		sortSizes(sizes) {
			return sizes.slice().sort((a, b) => parseFloat(a.width) - parseFloat(b.width));
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
				.catch(() => {
					this.$noty.error("Error occurred when loading authors, please refresh.")
				})
		},
		/*
		 * findUserById()
		 */
		findUserById(id) {
			return this.users.find(u => u.id === id);
		},
	},
	computed: {
		/*
		 * getSiteUrl()
		 * Get the site url from the store for previewing.
		 */
		getSiteUrl() {
			return this.$store.state.site.url;
		},
		bulkDelete: {
			get() {
				return this.bulkAction;
			},
			set(value) {
				this.$emit("update:bulk-action", value)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

.media {
	$self: &;


	// Placeholder
	// =========================================================================

	&-placeholder {
		position: relative;
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
			font-size: 46px;
			color: $orange;
			margin-bottom: 1rem;
			transition: color 100ms ease;
			will-change: color;
		}

		h4 {
			font-weight: 600;
			color: $copy;
		}
	}

	// Dragging
	// =========================================================================

	&-dragging {
		border: 2px dashed $grey-light;
		background-color: rgba($primary, 0.14);
		cursor: copy;
		padding: 1rem 1rem 0 1rem;

		&-empty {
			padding: 0;
			border: none;
			background-color: $white;

			#{$self}-placeholder {
				background-color: rgba($primary, 0.08);

				i {
					color: $primary;
				}
			}
		}

		* {
			cursor: copy;
		}
	}

	// Files
	// =========================================================================

	&-files {
		display: flex;
		flex-wrap: wrap;
		justify-content: flex-start;
		width: 100%;
		margin: 0 -15px -1rem -15px;
		padding-bottom: 20px;
		border-radius: 10px;
		transition: all 100ms ease;

		&-selected {

			#{$self}-item:not(#{$self}-item-active) {
				filter: grayscale(100%);
				opacity: 0.5;
			}
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
		flex-basis: calc(20% - 15px);
		margin-right: 15px;
		height: 200px;
		background-color: $grey;
		margin-bottom: 1rem;
		border-radius: 6px;
		cursor: pointer;
		transition: all 200ms ease;

		&-image {
			width: 100%;
			height: 100%;
			margin: 0;

			img {
				border-radius: 6px;
				width: 100%;
				height: 100%;
				object-fit: cover;
			}
		}


		&:nth-child(5n) {
			margin-right: 0;
		}

		&-loading {
			background-color: $white;
			box-shadow: $form-box-shadow;
			cursor: default;

			&-cont {
				position: relative;
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
				width: 100%;
				height: 100%;
				z-index: 99;
				box-shadow: none;
				user-select: none;

				h4 {
					color: $copy;
					margin-top: 1rem;
					font-size: 0.86rem;
					user-select: none;
				}

				&-error {
					display: flex;
					flex-direction: column;
					align-items: center;

					i {
						display: inline-block;
						font-size: 26px;
						color: $orange;
						margin-bottom: 4px;
					}

					p {
						color: $orange;
					}
				}
			}
		}



		// Checkbox
		// =========================================================================

		&-checkbox {
			position: absolute;
			bottom: 10px;
			right: 10px;
			z-index: 9999;
		}

		// Active (Clicked)
		// =========================================================================

		&-active {
			opacity: 1;
			box-shadow: 0 0 10px 0 rgba($black, 0.14);
		}
	}

	// Select Card
	// =========================================================================

	&-select-card {
		box-shadow: 0 3px 6px 0 rgba($black, 0.05);

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

	&-options {
		margin-bottom: 1.6em;
	}

	// Information
	// =========================================================================

	&-information {

		.text-cont {
			margin-bottom: 1.4rem;
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

	// Size
	// =========================================================================

	&-size {
		padding: 1rem;
		margin-bottom: 10px;
		border-radius: 10px;
		background-color: $white;

		&:first-of-type {
			margin-top: 10px;
		}

		&-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			width: 100%;
			margin-bottom: 4px;
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
		}
	}
}

</style>