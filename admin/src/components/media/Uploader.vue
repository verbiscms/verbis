<!-- =====================
	Media
	===================== -->
<template>
	<section class="media" :class="{ 'media-loading' : uploading, 'media-dragging' : dragging }" @drop.prevent="addFile($event, false)" @dragover.prevent @dragover="dragging = true" @dragleave="dragging = false">
		<div class="row">
			<!-- =====================
				Editor
				===================== -->
			<div class="col-12 col-desk-3 order-desk-last">
				<div v-if="selectedMedia">
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
						<p>{{ selectedMedia.type }}</p>
					</div>
					<form class="form">
						<div class="form-group">
							<input type="text" class="form-input form-input-white">
						</div>
					</form>
				</div>

				<pre>{{ selectedMedia }}</pre>

			</div><!-- /Col -->
			{{ users }}
			<!-- =====================
				Modal
				===================== -->
			<div class="col-12 col-desk-9">
				<div class="media-files">
					<div class="media-item" v-for="item in media" :key="item.uuid" @click.prevent="changeSelectedMedia(item)">
						<img :src="getSiteUrl + item.url" :alt="item.alt">
<!--						<div class="spinner spinner-grey"></div>-->
					</div><!-- /Media Item -->
				</div><!-- /Media Files -->
			</div><!-- /Col -->
		</div><!-- /Row -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Uploader",
	// props: {
	// 	insertedFiles: {
	// 		type: [Object, Event],
	// 	},
	// },
	data: () => ({
		files: [
			{
				index: 232324,
				name: "ainsphoto.jpg",
				uploading: true,
			},
			{
				index: 45345,
				name: "ainspgoto22.jpg",
				uploading: true,
			}
		],
		media: [],
		selectedMedia: false,
		users: [],
		uploading: false,
		dragging: false,
	}),
	// watch: {
	// 	insertedFiles(e) {
	// 		this.addFile(e, true);
	// 	}
	// },
	mounted() {
		this.getMedia();
		this.getUsers();
	},
	methods: {
		/*
		 * getMedia()
		 * Obtain the media or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getMedia() {
			this.axios.get(`media?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					console.log(res);
					this.media = {};
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					this.media = res.data.data;
				})
				.catch(err => {
					console.log(err)
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		addFile(e, manual = false) {
			this.uploading = true;

			console.log(e);

			let droppedFiles = manual ? e.target.files : e.dataTransfer.files;
			if(!droppedFiles) return;

			([...droppedFiles].forEach((file, index) => {
				this.files.push({
					name: file.name,
					index: index,
					uploading: true,
				});
				this.upload(file);
			}));

			this.$emit("uploaded", true);
		},
		upload(file) {
			let formData = new FormData();
			formData.append('file', file);

			this.axios("/media", {
				method: 'post',
				data: formData,
				headers: {'Content-Type': 'multipart/form-data' }
			})
				.then(res => {
					console.log(res);
				})
				.catch(err => {
					console.log(err);
				})
		},
		changeSelectedMedia(item) {
			this.selectedMedia = {};
			const user = this.findUserById(item['user_id']);
			user['full_name'] = `${user['first_name']} ${user['last_name']}`;
			item['uploaded_by'] = user;
			this.selectedMedia = item;
		},
		formatBytes(bytes, decimals = 2) {
			if (bytes === 0) return '0 Bytes';

			const k = 1024,
				dm = decimals < 0 ? 0 : decimals,
				sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
				i = Math.floor(Math.log(bytes) / Math.log(k));

			return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
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
		findUserById(id) {
			return this.users.find(u => u.id === id);
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
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

.media {
	$self: &;

	min-height: 300px;
	border-radius: 10px;

	// Dragging
	// =========================================================================

	&-dragging {
		border: 2px dashed $grey-light;
	}


	// Files
	// =========================================================================

	&-files {
		display: flex;
		flex-wrap: wrap;
		width: 100%;
		margin: 0 -15px;
		transition: opacity 100ms ease;
	}

	// Loading
	// =========================================================================

	&-item {
		flex-basis: calc(20% - 15px);
		margin-right: 15px;
		height: 200px;
		background-color: $grey;
		margin-bottom: 1rem;
		cursor: pointer;

		img {
			border-radius: 6px;
			width: 100%;
			height: 100%;
			object-fit: cover;
		}
	}

	// Loading
	// =========================================================================

	&-loading {

		#{$self}-placeholder {
			//opacity: 0;
		}

		#{$self}-files {
		}
	}
}


.media-item {

}


</style>