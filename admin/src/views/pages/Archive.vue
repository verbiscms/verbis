<!-- =====================
	Archive
	===================== -->
<template>
	<section class="">
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Posts</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<div class="form-select-cont form-input">
									<select class="form-select">
										<option value="" disabled selected>Bulk actions</option>
										<option value="">Move to drafts</option>
										<option value="">Delete</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white">Apply</button>
								<button class="btn btn-icon btn-orange">
									<i class="fal fa-plus"></i>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Tabs -->
			<div class="row">
				<div class="col-12">
					<Tabs></Tabs>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">

				</div>
			</div>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Tabs from "../../components/misc/Tabs";

export default {
	name: "Pages",
	components: {
		Breadcrumbs,
		Tabs
	},
	data: () => ({
		doingAxios: false,
		authInfo: {
			email: '',
			pass: '',
		},
		authMessage: ''
	}),
	mounted() {
		this.getResourceByName()
		this.test();
	},
	watch: {
		'$route.params.resource': function () {
			this.getResourceByName();
		}
	},
	methods: {
		getResourceByName() {

		},
		test() {
			this.axios.get("/posts/pages")
				.then(res => {
					console.log(res);
				});
		}
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