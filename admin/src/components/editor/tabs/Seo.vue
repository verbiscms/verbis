<!-- =====================
	Seo Options
	===================== -->
<template>
	<section>
		<div class="card">
			<collapse :show="true">
				<template v-slot:header>
					<div class="card-header">
						<h3 class="card-title">SEO</h3>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<h4>Coming soon!</h4>
					</div><!-- /Card Body -->
				</template>
			</collapse>
		</div><!-- /Card -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";

export default {
	name: "SeoOptions",
	components: {
		Collapse,
	},
	data: () => ({

	}),
	methods: {

	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>