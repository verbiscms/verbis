<!-- =====================
	Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<div class="header-icon-cont">
								<i :class="getResourceData.icon"></i>
								<h1>How the Development Team Got Their Jobs at Reddico</h1>
							</div>
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
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

	import Breadcrumbs from "../../components/misc/Breadcrumbs";

	export default {
		name: "Single",
		components: {
			Breadcrumbs,
		},
		data: () => ({

		}),
		mounted() {
			this.getResourceDataTest()
		},
		methods: {
			findResourceByName(resources, name) {
				const resource =  resources.find(r => r.name.toLowerCase() === name);
				if (resource === undefined) {
					//return this.$router.push('/404')
				}
				return resource
			},
			getResourceDataTest() {
				const id = this.$route.params.id;
				this.axios.get(`/posts/${id}`)
				.then(res => {
					console.log(res)
				})
				.catch(err => {
					console.log(err)
				})
			}
		},
		computed: {
			getResourceData() {
				// const resources = this.$store.state.resources,
				// 	currentSlug = this.$route.params.resource,
				// 	resource = {};
				//
				// if (currentSlug === "pages") {
				// 	resource.icon = "fal fa-file"
				// } else {
				// 	const r = this.findResourceByName(resources, currentSlug)
				// 	resource.icon = r.options.icon
				// }
				//
				// resource.name = currentSlug
				//
				// return resource;

				return 0;
			},
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