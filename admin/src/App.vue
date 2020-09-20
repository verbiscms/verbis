<!-- =====================
	Home
	===================== -->
<template>
	<div id="app">
		<Sidebar v-if="isLoggedIn"></Sidebar>
		<main>
			<router-view/>
		</main>
	</div><!-- /App -->
</template>

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
	mounted: function(){
		// this.$moment({
		// 	timezone: 'Europe/London'
		// })
		// this.$noty.info("Welcome to the CMS");
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
