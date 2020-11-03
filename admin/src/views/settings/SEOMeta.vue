<!-- =====================
	Settings - SEO & Meta
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Seo & Meta Settings</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Settings</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios && loadingMeta" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios && !loadingMeta" class="row trans-fade-in-anim">
				<!-- =====================
					Visibility
					===================== -->
				<div class="col-12">
					<h6 class="margin">Visibility</h6>
						<div class="card card-small-box-shadow card-expand">
						<!-- Public -->
						<div class="collapse-border-bottom">
							<div class="card-header card-header-block">
								<div>
									<h4 class="card-title">Public</h4>
									<p>By disabling public, no social media meta data will be outputted and a <code v-text="'<meta name=\'robots\' content=\'noindex\'>'"></code> will be placed globally.</p>
								</div>
								<div class="toggle">
									<input type="checkbox" class="toggle-switch" id="site-public" v-model="data['seo_public']" :true-value="true" :false-value="false" />
									<label for="site-public"></label>
								</div>
							</div><!-- /Card Header -->
						</div><!-- /Public -->
					</div><!-- /Card -->
				</div>
				<!-- =====================
					Meta Information
					===================== -->
				<div class="col-12">
					<h6 class="margin">Global meta</h6>
					<div v-if="!loadingMeta" class="card card-small-box-shadow card-expand">
						<MetaForm :meta="meta" @update="updateMeta"></MetaForm>
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import MetaForm from "@/components/meta/Meta";

export default {
	name: "SeoMeta",
	mixins: [optionsMixin],
	components: {
		MetaForm,
		Breadcrumbs
	},
	data: () => ({
		errorMsg: "Fix the errors before saving SEO & Meta settings.",
		successMsg: "Seo & Meta options updated successfully.",
		showImageModal: false,
		selectedImageType: "",
		facebookImage: false,
		twitterImage: false,
		meta: {},
		loadingMeta: true,
	}),
	methods: {
		/*
		 * runAfterGet()
		 * Insert facebook or twitter image & update the height
		 * after the options have been obtained.
		 */
		runAfterGet() {
			for (const key in this.data) {
				if (key.includes('meta_')) {
					this.$set(this.meta, key, this.data[key]);
				}
			}
			this.loadingMeta = false;
		},
		/*
		 * updateMeta()
		 * Sets the data using the key when the meta updates.
		 */
		updateMeta(val, key) {
			this.$set(this.data, key, val);
		}
	}
}

</script>