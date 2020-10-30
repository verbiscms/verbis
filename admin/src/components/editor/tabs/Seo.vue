<!-- =====================
	Seo Options
	===================== -->
<template>
	<section>
		<div class="card card-small-box-shadow card-expand">
			<Collapse :show="false" class="collapse-border-bottom">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">SEO</h4>
							<p>Coming soon!</p>
						</div>
					</div><!-- /Card Header -->
				</template>
			</Collapse>
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