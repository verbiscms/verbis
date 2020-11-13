<!-- =====================
	Seo Options
	===================== -->
<template>
	<section>
		<div class="card card-small-box-shadow card-expand">
			{{ value }}
			<!-- Public -->
			<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['public'] }">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Public</h4>
							<p>Disabling public will place a <code>no robots</code> meta tag on the page, so the page is not visible to search engines.</p>
						</div>
						<div class="toggle">
							<input type="checkbox" class="toggle-switch" id="public" checked v-model="value['public']" :true-value="true" :false-value="false" />
							<label for="public"></label>
						</div>
					</div><!-- /Card Header -->
				</template>
			</Collapse><!-- /Public -->
			<!-- Exclude Sitemap -->
			<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['exclude_sitemap'] }">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Exclude from sitemap?</h4>
							<p>Check this toggle to exclude this page from the dynamically generated sitemap.</p>
						</div>
						<div class="toggle">
							<input type="checkbox" class="toggle-switch" id="exclude_sitemap" checked v-model="value['exclude_sitemap']" :true-value="true" :false-value="false" />
							<label for="exclude_sitemap"></label>
						</div>
					</div><!-- /Card Header -->
				</template>
			</Collapse><!-- /Exclude Sitemap -->
			<!-- Url -->
			<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['canonical']}">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Override canonical?</h4>
							<p>Verbis automatically generates a canonical for the page, if you wish to override it, enter a URL below.</p>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<!-- Url -->
						<FormGroup label="Site title*" :error="errors['site_url']">
							<input class="form-input form-input-white" type="text" v-model="value['canonical']">
						</FormGroup>
					</div><!-- /Card Body -->
				</template>
			</Collapse><!--/ Url -->
		</div><!-- /Card -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";
import FormGroup from "../../forms/FormGroup";

export default {
	name: "SeoOptions",
	components: {
		FormGroup,
		Collapse,
	},
	props: {
		seo: {
			type: Object,
			default: function () {
				return {}
			}
		}
	},
	data: () => ({
		data: {
			"public": true,
			"exclude_sitemap": false,
			"canonical": null,
		},
		errors: [],
	}),
	mounted() {
		//this.value = this.seo;
	},
	methods: {
		emit() {
		//	this.value = this.data;
		}
	},
	computed: {
		value: {
			get() {
				if (this.seo === null) {
					return {
						"public": true,
						"exclude_sitemap": false,
						"canonical": null,
					}
				}
				return this.seo;
			},
			set(value) {
				console.log("----");
				console.log(value);
				this.value = value;
				this.$emit("update:seo", value);
			},
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>