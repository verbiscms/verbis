<!-- =====================
	Settings - Media
	===================== -->
<template>
	<section>
		<div class="auth-container" v-if="!doingAxios">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Media Settings</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								<i class="far fa-check"></i>
								Update settings
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- =====================
				General
				===================== -->
			<div class="form-row-group">
				<div class="row">
					<div class="col-6">

					</div>
				</div>
				<div class="row">
					<div class="col-12">
						<h2>General media settings:</h2>
					</div><!-- /Col -->
				</div><!-- /Row -->
				<!-- First name -->
				<div class="row form-row form-row-border form-row-border-top">
					<div class="col-12 col-desk-6">
						<h4>Upload maximum size</h4>
						<div class="form-group">
							<input class="form-input form-input-white" type="text">
							<!-- Message -->
							<transition name="trans-fade-height">
<!--								<span class="field-message field-message-warning" v-if="errors['site_title']">{{ errors['site_title'] }}</span>-->
							</transition><!-- /Message -->
						</div>
						<h4>Upload maximum width</h4>
						<div class="form-group">
							<input class="form-input form-input-white" type="text">
							<!-- Message -->
							<transition name="trans-fade-height">
								<!--								<span class="field-message field-message-warning" v-if="errors['site_title']">{{ errors['site_title'] }}</span>-->
							</transition><!-- /Message -->
						</div>
					</div><!-- /Col -->
					<div class="col-12 col-desk-4">

					</div><!-- /Col -->
				</div><!-- /Row -->
			</div>
			<div class="row">
				<div class="col-12">
					<div class="text-cont">
						<h2>Image sizes:</h2>
					</div>
				</div><!-- /Col -->
				<!-- Thumbnail -->
				<div class="col-12 col-tab-6 col-desk-3" v-for="(size, sizeKey) in data['media_images_sizes']" :key="size.name">
					<div class="card card-small-box-shadow">
						<div class="card-header card-header-icon card-header-naked">
							<h3>{{ size.name}}</h3>
							<div class="badge">{{ helpers.capitalize(sizeKey) }}</div>
						</div>
						<div class="card-body">
							<div class="form-group">
								<label class="form-label" for="media-size-small-width">Width</label>
								<input class="form-input" id="media-size-small-width" type="number" placeholder="Enter a width" v-model="data['media_images_sizes'][sizeKey]['width']">
							</div>
							<div class="form-group">
								<label class="form-label" for="media-size-small-height">Height</label>
								<input class="form-input" id="media-size-small-height" type="number" placeholder="Enter a height" v-model="data['media_images_sizes'][sizeKey]['height']">
							</div>
							{{ sizeKey }}
							{{ data['media_images_sizes'][sizeKey]['crop'] }}
							<div class="form-checkbox">
								<input type="checkbox" id="media-size-small-crop" v-model="data['media_images_sizes'][sizeKey].crop" :true-value="true" :false-value="false" />
								<label for="media-size-small-crop">
									<i class="fal fa-check"></i>
								</label>
								<div class="form-checkbox-text">Crop image?</div>
							</div>
						</div><!-- /Card Body -->
					</div><!-- /Card  -->
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

export default {
	name: "Media",
	title: 'Media Settings',
	mixins: [optionsMixin],
	components: {
		Breadcrumbs
	},
	data: () => ({
		sizes: [],
	}),
	mounted() {

	},
	methods: {
		/*
		 * runAfter()
		 * The options hae
		 */
		runAfter() {
			this.sortImageSizes()
		},
		/*
		 * sortSizes()
		 * Sort sizes by width for the size cards.
		 */
		sortImageSizes() {
			this.data['media_images_sizes'] = Object.fromEntries(
				Object.entries(this.data['media_images_sizes']).sort(([,a],[,b]) => {
					return parseFloat(a.width) - parseFloat(b.width)
				})
			);
		}
	},
	computed: {
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

	h3 {
		//color: $green;
		//color: $green;
	}

</style>