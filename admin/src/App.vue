<!-- =====================
	Home
	===================== -->
<template>
	<div id="app">
		<!-- Sidebar -->
		<Sidebar v-if="isLoggedIn" :open="sidebarOpen" @close="sidebarOpen = false"></Sidebar>
		<!-- Mobile Menu -->
		<MobileMenu v-if="isLoggedIn" :sidebar.sync="sidebarOpen"></MobileMenu>
		<main>
			<!-- Main Router View -->
			<div class="router">
				<span class="router-overlay" :class="{ 'router-overlay-active' : sidebarOpen }" @click="sidebarOpen = false"></span>
				<TransitionPage>
					<router-view class="router" :key="$route.fullPath" />
				</TransitionPage>
			</div>
		</main>
	</div><!-- /App -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Sidebar from '@/components/layout/Sidebar.vue'
import MobileMenu from "@/components/layout/MobileMenu";
import TransitionPage from "@/components/misc/TransitionPage";

export default {
	name: 'App',
	components: {
		MobileMenu,
		Sidebar,
		TransitionPage
	},
	metaInfo: {
		meta: [
			{ charset: 'utf-8' },
			{ name: 'description', content: 'Verbis CMS' },
			{ name: 'viewport', content: 'width=device-width, initial-scale=1.0' }
		],
	},
	data: () => ({
		sidebarOpen: false,
	}),
	methods: {
		sendSocketMessage: function(message) {
			this.ws.send(message);
		},
	},
	created: function() {
		// console.log(this.$store.state.site)
		//
		// this.ws = new WebSocket('ws://' + window.location.host + '/api/v1/ws');
		//
		// this.ws.onopen = function(ctx){
		// 	console.log(ctx);
		// 	console.log("Web socket open")
		// }
		//
		// this.ws.onerror = function(ctx){
		// 	console.log(ctx);
		// 	console.log("** WS ERROR **")
		// }
		//
		// this.ws.onmessage = (msg) => {
		// 	console.log(msg);
		// 	let response = JSON.parse(msg.data);
		// 	console.log(response);
		// 	//this.$store.commit("setTheme", response)
		// };
	},
	computed: {
		/*
		 * isLoggedIn()
		 * Determines if the user is logged in from the store.
		 */
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

	// Import
	// =========================================================================

	@import "assets/scss/app.scss";

	// App
	// =========================================================================

	#app {
		overflow: hidden;
	}

	// Router Container
	// =========================================================================

	.router {
		$self: &;

		position: relative;
		// Transition Page fix
		height: auto !important;

		&-overlay {
			content: "";
			position: fixed;
			top: 0;
			left: 0;
			display: block;
			width: 100vw;
			height: 100vh;
			background-color: black;
			z-index: 0;
			opacity: 0;
			transition: opacity 300ms ease-in-out, z-index 300ms step-end;
			will-change: opacity;

			&-active {
				transition: opacity 300ms ease-in-out, z-index 300ms step-start;
				opacity: 0.4;
				z-index: 99;
			}
		}
	}

	// Desktop
	// =========================================================================

	@include media-desk {

		.router-overlay {
			display: none;
		}
	}

</style>
