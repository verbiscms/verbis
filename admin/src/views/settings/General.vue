<!-- =====================
	Settings - General
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>General Settings</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : saving }">Update settings</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->

			<div class="row" v-if="!doingAxios">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="col-12">
					<div class="card card-small-box-shadow">
						<collapse>
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Site Options</h3>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Title -->
									<FormGroup label="Site title*" :error="errors['site_title']">
										<input class="form-input form-input-white" type="text" v-model="data['site_title']">
									</FormGroup>
									<!-- Description -->
									<FormGroup label="Site description*" :error="errors['site_description']">
										<input class="form-input form-input-white" type="text" v-model="data['site_description']">
									</FormGroup>
									<!-- Url -->
									<FormGroup label="Site title*" :error="errors['site_url']">
										<input class="form-input form-input-white" type="text" v-model="data['site_url']">
									</FormGroup>
									<!-- Logo -->
									<div class="general-logo">
										<label class="form-label">Logo</label>
										<div v-show="!hasLogo">
											<button class="btn" @click.prevent="showImageModal = true">Insert Logo</button>
										</div>
										<div v-show="hasLogo">
											<ImageWithActions @choose="showImageModal = true" @remove="hasLogo = false">
												<img :src="getSiteUrl + data['logo']" @error="hasLogo = false"/>
											</ImageWithActions>
										</div>
									</div>
								</div><!-- /Card Body -->
							</template>
						</collapse>
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Social media
					===================== -->
				<div class="col-12">
					<div class="card card-small-box-shadow">
						<collapse>
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Social media</h3>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<div class="row no-gutter">
										<div class="col-12 col-desk-6">
											<FormGroup label="Facebook" :error="errors['social_facebook_url']">
												<input class="form-input form-input-white" type="text" v-model="data['social_facebook_url']">
											</FormGroup>
											<FormGroup label="Twitter" :error="errors['social_twitter_url']">
												<input class="form-input form-input-white" type="text" v-model="data['social_twitter_url']">
											</FormGroup>
											<FormGroup label="Youtube" :error="errors['social_youtube_url']">
												<input class="form-input form-input-white" type="text" v-model="data['social_youtube_url']">
											</FormGroup>
										</div><!-- /Card -->
										<div class="col-12 col-desk-6">
											<FormGroup label="LinkedIn" :error="errors['social_linked_in']">
												<input class="form-input form-input-white" type="text" v-model="data['social_linked_in']">
											</FormGroup>
											<FormGroup label="Instagram" :error="errors['social_instagram_url']">
												<input class="form-input form-input-white" type="text" v-model="data['social_instagram_url']">
											</FormGroup>
											<FormGroup label="Pinterest" :error="errors['social_pinterest_url']">
												<input class="form-input form-input-white" type="text" v-model="data['social_pinterest_url']">
											</FormGroup>
										</div>
									</div><!-- /Col -->
								</div><!-- /Card Body -->
							</template>
						</collapse>
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Contact Details
					===================== -->
				<div class="col-12">
					<div class="card card-small-box-shadow">
						<collapse>
							<template v-slot:header>
								<div class="card-header">
									<h3 class="card-title">Contact details</h3>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<div class="row no-gutter">
										<div class="col-12 col-desk-6">
											<FormGroup label="Address" :error="errors['contact_telephone']">
												<input class="form-input form-input-white" type="text" v-model="data['social_facebook_url']">
											</FormGroup>
										</div><!-- /Card -->
										<div class="col-12 col-desk-6">
											<FormGroup label="LinkedIn" :error="errors['contact_address']">
												<input class="form-input form-input-white" type="text" v-model="data['social_linked_in']">
											</FormGroup>
										</div>
									</div><!-- /Col -->
								</div><!-- /Card Body -->
							</template>
						</collapse>
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Insert Photo Modal
			===================== -->
		<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close">
			<template slot="text">
				<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertLogo" :options="false">
					<template slot="close">
						<button class="btn btn-margin-right btn-icon-mob" @click.prevent="showImageModal = false">
							<i class="feather feather-x"></i>
							<span>Close</span>
						</button>
					</template>
				</Uploader>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import FormGroup from "@/components/forms/FormGroup";
import Modal from "@/components/modals/General";
import Uploader from "@/components/media/Uploader";
import ImageWithActions from "@/components/misc/ImageWithActions";
import Collapse from "@/components/misc/Collapse";

export default {
	name: "General",
	title: 'General Settings',
	mixins: [optionsMixin],
	components: {
		Collapse,
		ImageWithActions,
		Uploader,
		Modal,
		FormGroup,
		Breadcrumbs
	},
	data: () => ({
		errorMsg: "Fix the errors before saving settings.",
		successMsg: "Site options updated successfully.",
		errors: [],
		hasLogo: true,
		showImageModal: false,
	}),
	methods: {
		/*
		 * insertLogo()
		 */
		insertLogo(e) {
			this.profilePicture = e;
			this.showImageModal = false;
			this.data['logo'] = e.url;
			this.hasLogo = true;
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Text
	// =========================================================================

	h2 {
		margin-bottom: 4px;
	}

	p {
		font-size: 0.9rem;
	}

</style>