<!-- =====================
	Meta Options
	===================== -->
<template>
	<section>
		<!-- =====================
			Serp
			===================== -->
		<div class="field">
			<h4>SERP Preview</h4>
			<div class="field-content">
				<div class="serp">
					<span class="serp-title">{{ value.title }}</span>
					<span class="serp-url">{{ getUrl }}</span>
					<span class="serp-description">{{ value.description }}</span>
				</div>
			</div>
		</div>
		<!-- =====================
			General Meta
			===================== -->
		<div class="field">
			<h4>Meta</h4>
			<div class="field-content">
				<div class="form-group">
					<h5>Title</h5>
					<p>This will appear at the top of the search preview</p>
					<input class="form-input form-input-white" type="text" v-model="value.title">
				</div>
				<div class="form-group">
					<h5>Description</h5>
					<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
					<textarea class="form-input form-input-white" type="text" rows="4" v-model="value.description"></textarea>
				</div>
			</div>
		</div>
		<!-- =====================
			Facebook
			===================== -->
		<div class="field">
			<h4>Facebook (Open Graph)</h4>
			<div class="field-content">
				<div class="form-group">
					<h5>Title</h5>
					<p>This will appear at the top of the search preview</p>
<!--					<input class="form-input form-input-white" type="text" v-model="value.facebook.title">-->
				</div>
				<div class="form-group">
					<h5>Description</h5>
					<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
<!--					<textarea class="form-input form-input-white" type="text" rows="4" v-model="value.facebook.description"></textarea>-->
				</div>
			</div>
		</div><!-- /Facebook -->
		<!-- =====================
			Twitter
			===================== -->
		<div class="field">
			<div class="row">
				<div class="col-12">
					<div class="meta-title">
						<h4>Twitter</h4>
						<p>Twitter cards allow you to attach rich photo's, videos and media to Tweets. Add a title & description below.</p>
					</div>
				</div><!-- /Col -->
				<div class="col-12 col-desk-6">
					<div class="field-content">
						<div class="form-group">
							<h5>Title</h5>
							<p>This will appear at the top of the search preview</p>
							<input class="form-input" type="text" v-model="value.twitter.title" :disabled="useTwitterGlobal">
						</div>
						<div class="form-group">
							<h5>Description</h5>
							<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
							<textarea class="form-input" type="text" rows="4" v-model="value.twitter.description" :disabled="useTwitterGlobal"></textarea>
						</div>
						<div class="form-group">
							<h5>Image</h5>
							<p>Add an image for the twitter card.</p>
							<button class="btn">Add image</button>
						</div>
					</div>
				</div><!-- /Col -->
				<div class="col-12 col-desk-6">
					<h5>Global:</h5>
					<p>The twitter preview will appear here:</p>
					<div class="form-checkbox checkbox-cont">
						<input type="checkbox" id="metatwittercheck" @change="updateGlobal('twitter')" v-model="useTwitterGlobal" :true-value="true" :false-value="false">
						<label for="metatwittercheck">
							<i class="fal fa-check"></i>
						</label>
						<div  class="form-checkbox-text">Use global meta?</div>
					</div>
					<h5>Preview:</h5>
					<p>The twitter preview will appear here:</p>
					<div class="twitter">
						<div class="twitter-image">
							<i class="fal fa-file-alt"></i>
						</div>
						<div class="twitter-text">
							<span class="twitter-title">{{ value['twitter']['title'] }}</span>
							<span class="twitter-description">{{ getDescription(140, value['twitter']['description']) }}</span>
							<span class="twitter-url">{{ getSiteUrl }}</span>
						</div>
					</div><!-- /Twitter Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div>

	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "MetaOptions",
	props: {
		url: {
			type: String,
			default: "",
		},
		meta: {
			type: Object,
			default: null,
		}
	},
	data: () => ({
		useFacebookGlobal: false,
		useTwitterGlobal: false,
		facebook: {
			"title": "",
			"description": "",
		},
		twitter: {
			"title": "",
			"description": "",
		},
		defaults: {
			"title": "",
			"description": "",
			"facebook": {
				"title": "",
				"description": "",
			},
			"twitter": {
				"title": "",
				"description": "",
			}
		},
	}),
	watch: {
		value: {
			handler(value){
				this.$emit("update:meta", value);
			},
			deep: true
		}
	},
	created() {
		this.isGlobal();
	},
	methods: {
		getDescription(length, text) {
			if (text && length < text.length) {
				return text.substring(0, length - 3) + "...";
			}
			return text;
		},
		isGlobal() {
			if (this.value.title === this.value.facebook.title && this.value.description === this.value.facebook.description) {
				this.useFacebookGlobal = true;
			}
			if (this.value.title === this.value.twitter.title && this.value.description === this.value.twitter.description) {
				this.useTwitterGlobal = true;
			}
		},
		updateGlobal(type) {
			if (type === "facebook") {
				if (!this.useFacebookGlobal) {
					this.$set(this.value, "facebook", this.facebook)
					this.facebook = {
						"title" : "",
						"description" : ""
					}
				} else {
					this.$set(this.facebook, "title", this.value["facebook"]["title"])
					this.$set(this.facebook, "description", this.value["facebook"]["description"])
					this.$set(this.value['facebook'], "title", this.value.title)
					this.$set(this.value['facebook'], "description", this.value.description)
				}
			} else if (type === "twitter") {
				if (!this.useTwitterGlobal) {
					this.$set(this.value, "twitter", this.twitter)
					this.twitter = {
						"title" : "",
						"description" : ""
					}
				} else {
					this.$set(this.twitter, "title", this.value["twitter"]["title"])
					this.$set(this.twitter, "description", this.value["twitter"]["description"])
					this.$set(this.value['twitter'], "title", this.value.title)
					this.$set(this.value['twitter'], "description", this.value.description)
				}
			}
		},
	},
	computed: {
		getUrl() {
			return this.url;
		},
		getSiteUrl() {
			return this.$store.state.site['url'];
		},
		value: {
			get() {
				return this.meta === null || this.meta === undefined ? this.defaults : this.meta;
			},
			set(value) {
				this.meta = value;
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.field{
	background-color: $white;
}

// Global
// =========================================================================

.meta {

	&-title {
		margin-bottom: 1.6rem;

		h4 {
			margin-bottom: 0;
		}
	}
}

// Twitter
// =========================================================================

.twitter {
	position: relative;
	display: flex;
	width: 100%;
	border: 1px solid #E1E8ED;
	border-radius: 10px;
	overflow: hidden;
	margin-top: 1rem;

	// Image
	// =========================================================================

	&-image {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 125px;
		min-height: 125px;
		min-width: 125px;
		background-color: $grey-light;

		i {
			color: $grey;
			font-size: 2rem;
		}
	}

	// Text
	// =========================================================================

	&-text {
		//flex-grow: 2;
		padding: 14px;
		background-color: $white;
		width: calc(100% - 125px);
	}


	&-title,
	&-description,
	&-url {
		color: $black;
		display: block;
	}

	&-title {
		font-size: 0.9rem;
		font-weight: 600;
		margin-bottom: 0;
		max-height: 23px;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis
	}

	&-description {
		color: $black;
		font-size: 0.8rem;
		height: 52px;
		overflow: hidden;
		line-height: 1.3;
		margin-bottom: 4px;
	}

	&-url {
		color: $grey;
		font-size: 0.8rem;
	}
}


</style>