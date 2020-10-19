<!-- =====================
	Modal
	===================== -->

<template>
	<div class="media" :class="{ 'media-loading' : uploading, 'media-dragging' : dragging }" @drop.prevent="addFile($event, false)" @dragover.prevent @dragover="dragging = true" @dragleave="dragging = false">
<!--		<div class="media-placeholder" >-->
<!--			<i  class="feather feather-image"></i>-->
<!--			<h4>Drop your images here or click the button below to browse.</h4>-->
<!--		</div>-->

		<div class="media-files">
			<div class="row">
				<div class="col-12 col-tab-6 col-desk-3" v-for="item in media" :key="item.uuid">
					<div class="media-item">
						<img :src="getSiteUrl + item.url" :alt="item.alt">
						{{ item }}
<!--						<div class="spinner spinner-grey"></div>-->
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div>
	</div>
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
	},
	methods: {
		/*
		 * getMedia()
		 * Obtain the media or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getMedia() {

			console.log("dfsgklhsdfg")

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

	padding: 20px;
	min-height: 300px;
	border-radius: 10px;

	// Dragging
	// =========================================================================

	&-dragging {
		border: 2px dashed $grey-light;
	}


	// Placeholder
	// =========================================================================

	&-placeholder {
		position: absolute;
		top: 50%;
		left: 50%;
		display: flex;
		flex-direction: column;
		align-items: center;
		transform: translate(-50%, -50%);
		transition: opacity 100ms ease;
		will-change: opacity;

		i {
			font-size: 60px;
			color: $grey-light;
			margin-bottom: 6px;
		}

		.btn {
			margin-top: 1rem;
		}

		input {
			position: absolute;
			display: none;
			top: -9999999px;
			left: -9999999px;
		}
	}

	// Files
	// =========================================================================

	&-files {
		width: 100%;
		//opacity: 0;
		transition: opacity 100ms ease;
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
	width: 100%;
	min-height: 200px;
	border: 1px solid $grey;
	background-color: $grey;
}


</style>