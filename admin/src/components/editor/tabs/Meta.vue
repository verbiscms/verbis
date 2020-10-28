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
<!--						<span class="card-serp-title" v-text="meta.title === '' ? 'Title will appear here' : meta.title"></span>-->
<!--						<span class="card-serp-url">{{ getUrl }}</span>-->
<!--						<span class="card-serp-description" v-text="meta.description === '' ? 'Description will appear here' : meta.description"></span>-->
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
		<div v-if="!loadingMeta" class="card card-small-box-shadow card-expand">
			<MetaForm :meta="test" @update="updateMeta"></MetaForm>
		</div>
		{{ test }}
		{{ emitVal }}
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
			},
			"twitter": {
				"title": "",
				"description": "",
			}
		}
	}),
	created() {
		this.loadingMeta = false;
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
					if (this.emitVal['facebook'] === undefined) {
						this.$set(this.emitVal, "facebook", {});
					}
					this.$set(this.emitVal['facebook'], "title", val);
					break;
				}
				case "meta_facebook_description": {
					if (this.emitVal['facebook'] === undefined) {
						this.$set(this.emitVal, "facebook", {});
					}
					this.$set(this.emitVal['facebook'], "description", val);
					break;
				}
				case "meta_twitter_title": {
					if (this.emitVal['twitter'] === undefined) {
						this.$set(this.emitVal, "twitter", {});
					}
					this.$set(this.emitVal['twitter'], "title", val);
					break;
				}
				case "meta_twitter_description": {
					if (this.emitVal['twitter'] === undefined) {
						this.$set(this.emitVal, "twitter", {});
					}
					this.$set(this.emitVal['twitter'], "description", val);
					break;
				}
			}
			this.test = this.emitVal;
		},
	},
	computed: {
		getUrl() {
			return this.url;
		},
		getSiteUrl() {
			return this.$store.state.site['url'];
		},
		test: {
			get() {
				const meta = this.meta;
				if (this.meta === null) {
					console.log("in")
					return {
						meta_title: "",
						meta_description: "",
						meta_facebook_title: "",
						meta_facebook_description: "",
						meta_twitter_title: "",
						meta_twitter_description: "",
					}
				} else {
					return {
						meta_title: meta['title'],
						meta_description: meta['description'],
						meta_facebook_title: meta['facebook']['title'],
						meta_facebook_description: meta['facebook']['description'],
						meta_twitter_title: meta['twitter']['title'],
						meta_twitter_description: meta['twitter']['description'],
					}
				}
			},
			set() {
				this.$emit("update:meta", this.emitVal);
			}
		}
	}
}

</script>

<!--<span class="social-title">{{ value['facebook']['title'] }}</span>-->
<!--<span class="social-description">{{ getDescription(140, value['facebook']['description']) }}</span>-->
<!--<span class="social-url">{{ getSiteUrl }}</span>-->