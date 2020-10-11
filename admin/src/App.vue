<!-- =====================
	Home
	===================== -->
<template>
	<div id="app">
		<Sidebar v-if="isLoggedIn"></Sidebar>
		<main>
			<router-view/>
		</main>
	</div>
</template><!-- /App -->

<!-- =====================
	Scripts
	===================== -->
<script>

import Sidebar from '@/components/layout/Sidebar.vue'

export default {
	name: 'App',
	components: {
		Sidebar,
	},
	metaInfo: {
		meta: [
			{ charset: 'utf-8' },
			{ name: 'description', content: 'Verbis CMS' },
			{ name: 'viewport', content: 'width=device-width, initial-scale=1.0' }
		],
	},
	computed: {
		isLoggedIn() {
			return this.$store.state.auth ? this.$store.state.auth : false;
		},
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

	@import "assets/scss/app.scss";

</style>
