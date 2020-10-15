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
							<button class="btn btn-fixed-height btn-orange btn-with-icon" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								<i class="far fa-check"></i>
								Update settings
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<form class="form">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="form-row-group">
					<div class="row">
						<div class="col-12">
							<h2>Site options</h2>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- First name -->
					<div class="row form-row form-row-border form-row-border-top">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Title</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['site_title']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['site_title']">{{ errors['site_title'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Description -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Description</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['site_description']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['site_description']">{{ errors['site_description'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Url -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Url</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['site_url']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['site_url']">{{ errors['site_url'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
					<!-- Logo -->
					<div class="row form-row form-row-border">
						<div class="col-12 col-desk-4 col-hd-2">
							<h4>Logo</h4>
						</div>
						<div class="col-12 col-desk-8 col-hd-6">
							<div class="form-group">
								<input class="form-input form-input-white" type="text" v-model="data['site_logo']">
								<!-- Message -->
								<transition name="trans-fade-height">
									<span class="field-message field-message-warning" v-if="errors['site_logo']">{{ errors['site_logo'] }}</span>
								</transition><!-- /Message -->
							</div>
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div><!-- Form Group -->
				<!-- =====================
					Social media
					===================== -->
				<div class="form-row-group">
					<div class="row">
						<div class="col-12">
							<h2>Social media</h2>
						</div><!-- /Col -->
						<div class="col-12 col-tab-6 col-desk-3">
							<div class="card">
								<div class="card-body">
									<div class="card-social">
										<img src="@/assets/images/facebook.svg">
										<h4>Facebook</h4>
									</div>
									<div class="form-group">
										<input class="form-input form-input-white" type="text" v-model="data['social_facebook_url']" placeholder="https://">
										<!-- Message -->
										<transition name="trans-fade-height">
											<span class="field-message field-message-warning" v-if="errors['social_facebook_url']">{{ errors['social_facebook_url'] }}</span>
										</transition><!-- /Message -->
									</div>
								</div><!-- /Card Body -->
							</div><!-- /Card -->
						</div><!-- /Col -->
						<div class="col-12 col-tab-6 col-desk-3">
							<div class="card card-test">
								<div class="card-body">
									<div class="card-social">
										<img src="@/assets/images/twitter.svg">
										<h3>Twitter</h3>
									</div>
									<div class="form-group">
										<input class="form-input form-input-white" type="text" v-model="data['social_twitter_url']" placeholder="https://">
										<!-- Message -->
										<transition name="trans-fade-height">
											<span class="field-message field-message-warning" v-if="errors['social_twitter_url']">{{ errors['social_twitter_url'] }}</span>
										</transition><!-- /Message -->
									</div>
								</div><!-- /Card Body -->
							</div><!-- /Card -->
						</div><!-- /Col -->
						<div class="col-12 col-tab-6 col-desk-3">
							<div class="card">
								<div class="card-body">
									<div class="card-social">
										<img src="@/assets/images/instagram.svg">
										<h3>Instagram</h3>
									</div>
									<div class="form-group">
										<input class="form-input form-input-white" type="text" v-model="data['social_instagram_url']" placeholder="https://">
										<!-- Message -->
										<transition name="trans-fade-height">
											<span class="field-message field-message-warning" v-if="errors['social_instagram_url']">{{ errors['social_instagram_url'] }}</span>
										</transition><!-- /Message -->
									</div>
								</div><!-- /Card Body -->
							</div><!-- /Card -->
						</div><!-- /Col -->
						<div class="col-12 col-tab-6 col-desk-3">
							<div class="card">
								<div class="card-body">
									<div class="card-social">
										<img src="@/assets/images/linkedin.svg">
										<h3>LinkedIn</h3>
									</div>
									<div class="form-group">
										<input class="form-input form-input-white" type="text" v-model="data['social_linkedin_url']" placeholder="https://">
										<!-- Message -->
										<transition name="trans-fade-height">
											<span class="field-message field-message-warning" v-if="errors['social_linkedin_url']">{{ errors['social_linkedin_url'] }}</span>
										</transition><!-- /Message -->
									</div>
								</div><!-- /Card Body -->
							</div><!-- /Card -->
						</div><!-- /Col -->
					</div><!-- /Row -->
				</div>
			</form>
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
	name: "General",
	title: 'General Settings',
	mixins: [optionsMixin],
	components: {
		Breadcrumbs
	},
	data: () => ({
		errorMsg: "Fix the errors before saving settings.",
		successMsg: "Site options updated successfully."
	})
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

	.card-test {
		//box-shadow: none;
		//border: 1px solid #1DA1F2;
	}

	.card-social {
		display: flex;
		align-items: center;
		margin-bottom: 1rem;

		img {
			margin-right: 4px;
			height: 50px;
		}

		h3 {
			color: $copy-light;
			margin-bottom: 0;
		}
	}

</style>