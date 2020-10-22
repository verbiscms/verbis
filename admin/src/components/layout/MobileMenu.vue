<!-- =====================
	Mobile Menu
	===================== -->
<template>
	<nav class="menu">
		<ul class="menu-list">
			<!-- Home -->
			<li class="menu-item" :class="{ 'menu-item-active' : activePage === 'home' }">
				<router-link class="menu-link" :to="{ name: 'home' }">
					<i class="feather feather-home"></i>
					<span>Home</span>
				</router-link>
			</li>
			<!-- Pages -->
			<li class="menu-item" :class="{ 'menu-item-active' : activePage === 'pages' }">
				<router-link class="menu-link" :to="{ name: 'resources', params: { resource: 'pages' }}">
					<i class="feather feather-file"></i>
					<span>Pages</span>
				</router-link>
			</li>
			<!-- Settings -->
			<li class="menu-item" :class="{ 'menu-item-active' : activePage === 'settings-general' }">
				<router-link class="menu-link" :to="{ name: 'settings-general' }">
					<i class="feather feather-settings"></i>
					<span>General</span>
				</router-link>
			</li><!-- /Settings -->
			<!-- Menu -->
			<li class="menu-item" :class="{ 'menu-item-active' : sidebarOpen }">
				<div class="menu-link" @click="sidebarOpen = !sidebarOpen">
					<i class="feather feather-menu"></i>
					<span>Menu</span>
				</div>
			</li><!-- /Menu -->
		</ul>
	</nav>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "MobileMenu",
	props: {
		sidebar: {
			type: Boolean,
			default: false,
		}
	},
	data: () => ({
		activePage: "",
	}),
	watch: {
		'$route'() {
			this.setActivePage();
		}
	},
	methods: {
		/*
		 * setActivePage()
		 * Define the active class by resource or query
		 */
		setActivePage() {
			const editorResource = this.$route.query.resource;
			const archiveResource = this.$route.params.resource;
			if (editorResource) {
				this.activePage = editorResource;
			} else if (archiveResource) {
				this.activePage = archiveResource;
			} else {
				this.activePage = this.$route.name;
			}
		},
	},
	computed: {
		/*
		 * sidebarOpen()
		 * Get & set the sidebar variable to emit to the parent.
		 */
		sidebarOpen: {
			get() {
				return this.sidebar;
			},
			set(value) {
				this.$emit("update:sidebar", value)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.menu {
	$self: &;

	position: fixed;
	bottom: 0;
	left: 0;
	width: 100vw;
	height: $mobile-menu-height-mob;
	z-index: 9997;
	background-color: $white;
	box-shadow: -3px 0 10px 0 rgba($black, 0.16);

	// List
	// ==========================================================================

	&-list {
		display: flex;
		justify-content: space-between;
		align-items: center;
		height: 100%;
	}

	// Item
	// ==========================================================================

	&-item {
		flex-basis: 25%;
		height: 100%;
		padding: 6px 4px;

		&-active {

			#{$self}-link {
				background-color: $primary;

				* {
					color: $white;
				}
			}
		}
	}

	// Link
	// ==========================================================================

	&-link {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
		color: $grey;
		border-radius: 2px;
		transition: background-color 250ms ease;
		will-change: background-color;

		* {
			transition: color 250ms ease;
			will-change: color;
		}

		i {
			display: block;
			font-size: 16px;
			height: 14px;
			margin-bottom: 4px;
			color: $grey;
		}
	}

	// Tablet
	// ==========================================================================

	@include media-tab {
		height: $mobile-menu-height-tab;
	}

	// Desktop
	// ==========================================================================

	@include media-desk {
		display: none;
	}
}

</style>