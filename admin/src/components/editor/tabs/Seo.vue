<!-- =====================
	Seo Options
	===================== -->
<template>
	<section>
		<div class="card card-small-box-shadow card-expand">
			<!-- Public -->
			<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['private'] }">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Private</h4>
							<p>Enabling private will place a <code>no robots</code> meta tag on the page, so the page is not visible to search engines.</p>
						</div>
						<div class="toggle">
							<input type="checkbox" class="toggle-switch" id="public" checked v-model="data['private']" @change="emit" :true-value="true" :false-value="false" />
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
							<input type="checkbox" class="toggle-switch" id="exclude_sitemap" checked  v-model="data['exclude_sitemap']" @change="emit" :true-value="true" :false-value="false" />
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
							<p>Verbis automatically generates a canonical for the page, if you wish to override it, enter a URL below. Ensure to add the entire url, for example <code>https://verbiscms.com/posts</code>.</p>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<!-- Url -->
						<FormGroup label="Url" :error="errors['site_url']">
							<input class="form-input form-input-white" type="text" v-model="data['canonical']" @keyup="emit">
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
				return {
					"private": false,
					"exclude_sitemap": false,
					"canonical": null,
				}
			}
		}
	},
	data: () => ({
		data: {
			"public": false,
			"exclude_sitemap": false,
			"canonical": null,
		},
		errors: [],
	}),
	mounted() {
		if (this.seo) {
			this.data = this.seo;
		}
		// Update the parent with default options when mounted.
		this.emit();
	},
	methods: {
		/*
		 * emit()
		 * Update the parent SEO object.
		 */
		emit() {
			if (this.data['canonical'] === "") {
				this.$set(this.data, 'canonical', null)
			}
			this.$emit("update:seo", this.data);
		}
	},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>