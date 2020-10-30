<!-- =====================
	Meta Options
	===================== -->
<template>
	<section>
		<!-- =====================
			Serp
			===================== -->
		<div class="card card-serp card-small-box-shadow card-expand card-expand-margin-small">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Serp Preview</h4>
							<p>Hello</p>
						</div>

						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<span class="card-serp-title" v-text="metaProcessed['meta_title'] === '' ? 'Title will appear here' : metaProcessed['meta_title']"></span>
						<span class="card-serp-url">{{ getSiteUrl + getUrl }}</span>
						<span class="card-serp-description" v-text="metaProcessed['meta_description'] === '' ? 'Description will appear here' : metaProcessed['meta_description']"></span>
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
		<div v-if="!loadingMeta" class="card card-small-box-shadow card-expand">
			<MetaForm :meta="metaProcessed" @update="updateMeta"></MetaForm>
		</div>
	<pre>	{{ metaProcessed }}</pre>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";
import MetaForm from "@/components/meta/Meta";

export default {
	name: "MetaOptions",
	props: {
		url: {
			type: String,
			default: "",
		},
		meta: {
			type: Object,
		}
	},
	components: {
		MetaForm,
		Collapse,
	},
	data: () => ({
		processedVal: {},
		loadingMeta: true,
		emitVal:  {
			"title": "",
			"description": "",
			"facebook": {
				"title": "",
				"description": "",
				"image_id": "",
			},
			"twitter": {
				"title": "",
				"description": "",
				"image_id": "",
			}
		}
	}),
	created() {
		this.loadingMeta = false;
		if (this.meta !== null && this.meta !== undefined) {
			this.emitVal = this.meta;
		}
	},
	methods: {
		/*
		 * updateMeta()
		 * Sets the data using the key when the meta updates.
		 */
		updateMeta(val, key) {
			switch (key) {
				case "meta_title": {
					this.$set(this.emitVal, "title", val);
					break;
				}
				case "meta_description": {
					this.$set(this.emitVal, "description", val);
					break;
				}
				case "meta_facebook_title": {
					this.$set(this.emitVal['facebook'], "title", val);
					break;
				}
				case "meta_facebook_description": {
					this.$set(this.emitVal['facebook'], "description", val);
					break;
				}
				case "meta_facebook_image_id": {
					this.$set(this.emitVal['facebook'], "image_id", val);
					break;
				}
				case "meta_twitter_title": {
					this.$set(this.emitVal['twitter'], "title", val);
					break;
				}
				case "meta_twitter_description": {
					this.$set(this.emitVal['twitter'], "description", val);
					break;
				}
				case "meta_twitter_image_id": {
					this.$set(this.emitVal['twitter'], "image_id", val);
					break;
				}
			}
			this.metaProcessed = this.emitVal;
		},
	},
	computed: {
		/*
		 * getUrl()
		 */
		getUrl() {
			return this.url;
		},
		metaProcessed: {
			get() {
				const meta = this.meta;
				if (this.meta === null || this.meta === undefined) {
					return {
						meta_title: "",
						meta_description: "",
						meta_facebook_title: "",
						meta_facebook_image_id: "",
						meta_facebook_description: "",
						meta_twitter_title: "",
						meta_twitter_description: "",
						meta_twitter_image_id: "",
					}
				} else {
					return {
						meta_title: meta['title'],
						meta_description: meta['description'],
						meta_facebook_title: meta['facebook']['title'],
						meta_facebook_description: meta['facebook']['description'],
						meta_facebook_image_id: meta['facebook']['image_id'],
						meta_twitter_title: meta['twitter']['title'],
						meta_twitter_description: meta['twitter']['description'],
						meta_twitter_image_id: meta['twitter']['image_id'],
					}
				}
			},
			set(value) {
				this.$emit("update:meta", value);
			}
		}
	}
}

</script>
