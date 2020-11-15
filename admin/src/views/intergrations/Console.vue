<!-- =====================
	Console
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Console</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<prism-editor class="prism prism-console" v-model="text" :highlight="highlighter"></prism-editor>
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
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-bash';

export default {
	name: "Console",
	title: "Console",
	components: {
		Breadcrumbs,
	},
	data: () => ({
		text: "",
	}),
	methods: {
		/*
		 * highlighter()
		 * Return html for prism editor.
		 */
		highlighter(code) {
			return highlight(code, languages.bash, "bash");
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