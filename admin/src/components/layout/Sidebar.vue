<!-- =====================
	Sidebar
	===================== -->
<template>
	<aside class="aside">
		<div class="aside-left">
			<!-- Top -->
			<div class="aside-left-top">
				<!-- Logo -->
				<router-link class="aside-logo" :to="{ name: 'home' }" v-if="getSite.logo">
					<img :src="globalBasePath + getSite.logo">
				</router-link>
				<!-- Pages -->
				<div class="aside-left-icon">
					<i class="fal fa-file"></i>
				</div>
				<!-- Users -->
				<div class="aside-left-icon">
					<i class="fal fa-users"></i>
				</div><!-- /Logo -->
				<!-- Media -->
				<div class="aside-left-icon">
					<i class="fal fa-images"></i>
				</div><!-- /Logo -->
				<!-- Settings -->
				<div class="aside-left-icon">
					<i class="fal fa-cog"></i>
				</div><!-- /Logo -->
			</div><!-- /Top -->
			<!-- Bottom -->
			<div class="aside-left-bottom">
				<div class="aside-left-icon aside-collapse" :class="{ 'aside-collapse-active' : collapsed }" @click="collapse">
					<i class="fal fa-arrow-alt-to-left"></i>
				</div>
				<div class="popover-cont">
					<button class="aside-initials" v-html="getInitials"></button>
					<div class="popover popover-top-right popover-no-arrow">
						<!-- Profile -->
						<router-link class="popover-item popover-item-icon" :to="{ name: 'profile' }">
							<i class="fal fa-id-card"></i>
							Profile
						</router-link>
						<!-- Logout -->
						<div class="popover-item popover-item-icon" @click="doLogout">
							<i class="fal fa-sign-out-alt"></i>
							Logout
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="aside-right">
			<div class="aside-info">
				<h2>{{ getSite.title }}</h2>
				<p>Admin CMS</p>
			</div>
			<!-- Navigation -->
			<nav class="aside-nav">
				<h6>Resources</h6>
				<ul>
					<!-- Pages -->
					<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'pages' }">
						<router-link class="aside-nav-link" to="/content/pages">
							<i class="fal fa-file"></i>
							<span>Pages</span>
						</router-link>
					</li><!-- /Pages -->
					<!-- Resources -->
					<li class="aside-nav-item" v-for="(resource) in resources" v-bind:key="resource.name"
						:class="{ 'aside-nav-item-active' : activePage === resource.name.toLowerCase() }">
						<router-link class="aside-nav-link" :to="'/content/' + resource.options.slug">
							<i v-if="resource.options.icon" class="fa" :class="resource.options.icon"></i>
							<i v-else class="fal fa-file"></i>
							<span>{{ resource.name }}</span>
						</router-link>
					</li><!-- /Resources -->
					<!-- Media -->
					<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'media' }">
						<router-link class="aside-nav-link" to="media">
							<i class="fal fa-images"></i>
							<span>Media</span>
						</router-link>
					</li><!-- /Media -->
					<!-- Settings -->
					<li class="aside-nav-item">
						<router-link class="aside-nav-link" to="/pages" :class="{ 'aside-nav-item-active' : activePage === 'settings' }">
							<i class="fal fa-cogs"></i>
							<span>Settings</span>
						</router-link>
					</li><!-- /Settings -->
				</ul>
				<h6>Pages</h6>
			</nav>
		</div><!-- /Container -->
	</aside><!-- /Aside Cont -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Sidebar",
	data: () => ({
		doingAxios: false,
		themeConfig: {},
		resources: {},
		activePage: "",
		collapsed: false
	}),
	beforeMount() {
		//this.getThemeConfig();
		//this.getResources();
	},
	watch: {
		'$route'() {
			this.setActivePage();
		}
	},
	methods: {
		doLogout() {
			this.axios.post("/logout", {})
				.then(res => {
					console.log(res);
					this.$store.commit("logout");
					location.reload();
				});
		},
		setActivePage() {
			const resource = this.$route.params.resource;
			this.activePage = resource === undefined ? this.$route.name : resource;
		},
		collapse() {
			document.querySelector(".auth-container").classList.toggle("auth-container-collapsed");
			this.collapsed = !this.collapsed;
		}
	},
	computed: {
		getUserInfo() {
			return this.$store.state.userInfo;
		},
		getSite() {
			return this.$store.state.site;
		},
		getInitials() {
			const info = this.getUserInfo
			return info['first_name'].charAt(0) + info['last_name'].charAt(0).toUpperCase();
		}
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Variables
$aside-padding: 10px;
$aside-initials-size: 44px;
$aside-left-icon-margin: 40px;
$aside-btn-padding-x: 18px;

.aside {
	$self: &;

	position: fixed;
	display: flex;
	flex-direction: row;
	align-items: flex-start;
	left: 0;
	top: 0;
	width: $sidebar-width;
	height: 100vh;
	background-color: $bg-color;
	z-index: 8;

	// Left
	// =========================================================================

	&-left {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		align-items: center;
		width: $sidebar-left-width;
		min-width: $sidebar-left-width;
		height: 100%;
		padding: 44px 0;
		overflow: visible;

		#{$self}-left-icon {
			display: flex;
			justify-content: center;
			margin-top: 40px;
			cursor: pointer;
			width: 100%;
			padding: 10px 0;

			i {
				font-size: 1.4rem;
				color: $grey;
				transition: 180ms ease;
			}

			&:hover i {
				color: $primary;
			}

			&-margin-bottom {
				margin-top: 0;
				margin-bottom: $aside-left-icon-margin;
			}
		}

		#{$self}-collapse {
			margin-top: 0;
			margin-bottom: $aside-left-icon-margin / 2;
			transition: transform 300ms ease;

			&-active {
				transform: rotate(180deg);
			}
		}
	}

	// Right
	// =========================================================================

	&-right {
		width: 100%;
		padding: $auth-container-padding-y $aside-padding $aside-padding;

		h6 {
			margin-left: $aside-btn-padding-x;
		}
	}

	// Info
	// =========================================================================

	&-info {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		margin-bottom: 2rem;
		margin-left: $aside-btn-padding-x;

		p,
		h2 {
			margin-bottom: 0;
			text-align: center;
		}
	}

	// Logo
	// =========================================================================

	&-logo {
		display: block;
		margin: 0 0 10px 0;
		width: 30px;

		img {
			width: 100%;
		}
	}

	// Initials
	// =========================================================================

	&-initials {
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: $secondary;
		width: $aside-initials-size;
		height: $aside-initials-size;
		border-radius: 100%;
		color: $white;
		font-size: 1rem;
		font-weight: 500;
		cursor: pointer;
		border: none;
	}

	// Nav
	// =========================================================================

	&-nav {

		ul:not(:last-child) {
			margin-bottom: 1rem;
		}

		// Item
		// =========================================================================

		&-item {
			margin-bottom: 10px;

			a {
				display: flex;
				align-items: center;
				justify-content: flex-start;
				padding: 12px $aside-btn-padding-x;
				border-radius: 8px;
				background-color: transparent;
				transition: background-color 400ms ease, box-shadow 400ms ease;
				will-change: background-color, box-shadow;
			}

			span,
			i {
				color: rgba($secondary, 0.4);
				transition: color 400ms ease;
				will-change: color;
			}

			i {
				font-size: 1.2rem;
				margin-right: 14px;
				width: 24px;
			}

			&:last-child {
				margin-bottom: 0;
			}
		}

		// Active
		// =========================================================================

		&-item-active {

			a {
				background-color: $white;
				box-shadow: 0 3px 20px rgba($black, 0.14);
			}

			span,
			i {
				color: $primary;
			}
		}

		// Hover
		// =========================================================================

		&-item:not(&-item-active):hover {

			a {
				background-color: rgba($white, 0.4);
				box-shadow: 0 3px 20px rgba($black, 0.14);
			}

			span,
			i {
				color: $secondary;
			}
		}
	}
}

</style>