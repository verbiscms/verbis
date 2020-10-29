<!-- =====================
	Settings - Performance
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header">
						<div class="header-title">
							<h1>Performance</h1>
							<Breadcrumbs></Breadcrumbs>
							<p>Caching</p>
							<p>gzip compression</p>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="col-12">
					<h6 class="margin">Cache control</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Title & description -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['site_title'] ||  errors['site_description'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Cache assets?</h4>
										<p>By ticking the box, assets will be globally.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="media-size-year" v-model="data['media_organise_year_month']" :true-value="true" :false-value="false" />
										<label for="media-size-year"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Title & description -->
						<!-- Title & description -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['site_title'] ||  errors['site_description'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Cacheability</h4>
										<p>By ticking the box, assets will be globally.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="media-size-year" v-model="data['media_organise_year_month']" :true-value="true" :false-value="false" />
										<label for="media-size-year"></label>
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
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Title & description -->
					</div><!-- /Card -->
				</div><!-- /Col -->

				Global options

				Each mime type has a maximum age (experiation)
				Expiration: max-age

				cacheability,
				disable cache no transform

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
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Performance",
	title: 'Performance Settings',
	mixins: [optionsMixin],
	components: {
		FormGroup,
		Collapse,
		Breadcrumbs
	},
	data: () => ({

	}),
	methods: {}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

</style>