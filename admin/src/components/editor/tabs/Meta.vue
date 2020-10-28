<!-- =====================
	Meta Options
	===================== -->
<template>
	<section>
		<!-- =====================
			Serp
			===================== -->
		<div class="card card-serp">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<h3 class="card-title">Serp Preview</h3>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<span class="card-serp-title" v-text="value.title === '' ? 'Title will appear here' : value.title"></span>
						<span class="card-serp-url">{{ getUrl }}</span>
						<span class="card-serp-description" v-text="value.description === '' ? 'Description will appear here' : value.description"></span>
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
		<!-- =====================
			General Meta
			===================== -->
		<div class="card">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<h3 class="card-title">Meta</h3>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<div class="card-input">
							<h4>Title</h4>
							<p>This will appear at the top of the search preview</p>
							<input class="form-input" type="text" v-model="value.title">
						</div>
						<div class="card-input">
							<h4>Description</h4>
							<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
							<textarea class="form-input" type="text" rows="4" v-model="value.description"></textarea>
						</div>
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
		<!-- =====================
			Facebook
			===================== -->
		<div class="card">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<div class="card-header-icon">
							<div class="icon icon-facebook">
								<i class="fab fa-facebook-f"></i>
							</div>
							<h3 class="card-title">Facebook (Open Graph)</h3>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<div class="row">
							<div class="col-12 col-desk-6">
								<div class="form-group">
									<h4>Title</h4>
									<p>This will appear at the top of the search preview</p>
									<input class="form-input" type="text" v-model="value.facebook.title" :disabled="useFacebookGlobal">
								</div>
								<div class="form-group">
									<h4>Description</h4>
									<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
									<textarea class="form-input" type="text" rows="4" v-model="value.facebook.description" :disabled="useFacebookGlobal"></textarea>
								</div>
								<div class="form-group">
									<h4>Image</h4>
									<p>Add an image for the twitter card.</p>
									<button class="btn btn-green">Add image</button>
								</div>
							</div><!-- /Col -->
							<div class="col-12 col-desk-6">
								<div class="form-group">
									<h4>Global:</h4>
									<p>Tick this box to use the global meta information from above</p>
									<div class="form-checkbox checkbox-cont">
										<input type="checkbox" id="metafacebookcheck" @change="updateGlobal('facebook')" v-model="useFacebookGlobal" :true-value="true" :false-value="false">
										<label for="metafacebookcheck">
											<i class="fal fa-check"></i>
										</label>
										<div  class="form-checkbox-text">Use global meta?</div>
									</div>
								</div>
								<h4>Preview:</h4>
								<p>The Facebook preview will appear here:</p>
								<div class="social">
									<div class="social-image">
										<i class="fal fa-file-alt"></i>
									</div>
									<div class="social-text">
										<span class="social-title">{{ value['facebook']['title'] }}</span>
										<span class="social-description">{{ getDescription(140, value['facebook']['description']) }}</span>
										<span class="social-url">{{ getSiteUrl }}</span>
									</div>
								</div><!-- /Twitter Card -->
							</div><!-- /Col -->
						</div><!-- /Row -->
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
		<!-- =====================
			Twitter
			===================== -->
		<div class="card">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<div class="card-header-icon">
							<div class="icon icon-twitter">
								<i class="fab fa-twitter"></i>
							</div>
							<h3 class="card-title">Twitter</h3>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<div class="row">
							<div class="col-12 col-desk-6">
								<div class="form-group">
									<h4>Title</h4>
									<p>This will appear at the top of the search preview</p>
									<input class="form-input" type="text" v-model="value.twitter.title" :disabled="useTwitterGlobal">
								</div>
								<div class="form-group">
									<h4>Description</h4>
									<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
									<textarea class="form-input" type="text" rows="4" v-model="value.twitter.description" :disabled="useTwitterGlobal"></textarea>
								</div>
								<div class="form-group">
									<h4>Image</h4>
									<p>Add an image for the twitter card.</p>
									<button class="btn">Add image</button>
								</div>
							</div><!-- /Col -->
							<div class="col-12 col-desk-6">
								<div class="form-group">
									<h4>Global:</h4>
									<p>Tick this box to use the global meta information from above</p>
									<div class="form-checkbox checkbox-cont">
										<input type="checkbox" id="metatwittercheck" @change="updateGlobal('twitter')" v-model="useTwitterGlobal" :true-value="true" :false-value="false">
										<label for="metatwittercheck">
											<i class="fal fa-check"></i>
										</label>
										<div  class="form-checkbox-text">Use global meta?</div>
									</div>
								</div>
								<h4>Preview:</h4>
								<p>The Twitter preview will appear here:</p>
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
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";

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
	components: {
		Collapse,
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
			if (this.value.title === this.value.facebook.title && this.value.description === this.value.facebook.description && this.value.facebook.title !== "" && this.value.facebook.description !== "") {
				this.useFacebookGlobal = true;
			}
			if (this.value.title === this.value.twitter.title && this.value.description === this.value.twitter.description && this.value.twitter.title !== "" && this.value.twitter.description !== "") {
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