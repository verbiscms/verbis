<!-- =====================
	Categories - Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>New category</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<h6 class="margin">General</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Title & description -->
						<Collapse :show="true" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Name</h4>
										<p>Set the site's location.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Name -->
									<FormGroup label="Name*">
										<input class="form-input form-input-white" type="text">
									</FormGroup><!-- /Name -->
									<!-- Slug -->
									<FormGroup label="Slug*">
										<input class="form-input form-input-white" type="text">
									</FormGroup><!-- /Slug -->
									<!-- Description -->
									<FormGroup label="Description">
										<input class="form-input form-input-white" type="text">
									</FormGroup><!-- /Description -->
									<!-- Resource -->
									<FormGroup label="Resource">
										<div class="form-select-cont form-input">
											<select class="form-select">
												<option v-for="resource in getTheme['resources']" :value="resource['friendly_name']" :key="resource.name">{{ resource['friendly_name'] }}</option>
											</select>
										</div>
									</FormGroup><!-- /Resource -->
								</div>
							</template>
						</Collapse><!-- /Title & description -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				Resource, Slug, Name, Descripition, ParentId
			</div><!-- /Row -->
		</div><!-- /Container -->

	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Categories",
	components: {
		FormGroup,
		Collapse,
		Breadcrumbs
	},
	data: () => ({
		errors: [],
	}),
	mounted() {

	},
	methods: {

	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Dummy
// =========================================================================

</style>