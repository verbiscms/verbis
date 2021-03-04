<!-- =====================
	Sidebar
	===================== -->
<template>
	<aside class="aside" :class="{ 'aside-active' : open }" v-if="getTheme">
		<div class="aside-container">
			<div class="aside-top">
				<!-- Logo -->
				<router-link class="aside-logo" :to="{ name: 'home' }" v-if="getSite.logo">
					<img :src="getSiteUrl + getSite.logo">
					<h2>Verbis</h2>
				</router-link>
				<!-- =====================
					Site
					===================== -->
				<div class="aside-block">
					<collapse :use-icon="false" :open="true">
						<template v-slot:header>
							<div class="aside-block-nav">
								<h6>Site</h6>
								<i class="feather feather-chevron-down"></i>
							</div>
						</template>
						<template v-slot:body>
							<nav class="aside-nav">
								<ul>
									<!-- Pages -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'site' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'site' }">
											<i class="feather feather-eye"></i>
											<span>View Site</span>
										</router-link>
									</li><!-- /Pages -->
								</ul>
							</nav>
						</template>
					</collapse>
				</div><!-- /Resources -->
				<!-- =====================
					Content
					===================== -->
				<div class="aside-block">
					<collapse :use-icon="false" :open="true">
						<template v-slot:header>
							<div class="aside-block-nav">
								<h6>Content</h6>
								<i class="feather feather-chevron-down"></i>
							</div>
						</template>
						<template v-slot:body>
							<nav class="aside-nav">
								<ul>
									<!-- Pages -->
									<li class="aside-nav-item aside-nav-item-icon" :class="{ 'aside-nav-item-active' : activePage === 'pages' || activePage === 'page' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'resources', params: { resource: 'pages' }}">
											<div>
												<i class="feather feather-file"></i>
												<span>Pages</span>
											</div>
											<router-link class="aside-icon" :to="{ name: 'editor', params: { id: 'new' }, query: { resource: 'pages' }}">
												<i class="feather feather-plus"></i>
											</router-link>
										</router-link>
									</li><!-- /Pages -->
									<!-- Resources -->
									<li class="aside-nav-item aside-nav-item-icon" v-for="(resource) in getTheme.resources"
										v-bind:key="resource.name"
										:class="{ 'aside-nav-item-active' : activePage === resource.name.toLowerCase() }"
										@click="$emit('close', true)">
										<router-link v-if="resource.name !== ''" class="aside-nav-link" :to="{ name: 'resources', params: { resource: resource.name }}">
											<div>
												<i v-if="resource.icon" :class="resource.icon"></i>
												<i v-else class="fal fa-file"></i>
												<span>{{ resource['friendly_name'] }}</span>
											</div>
											<router-link class="aside-icon" :to="{ name: 'editor', params: { id: 'new' }, query: { resource: resource.name }}">
												<i class="feather feather-plus"></i>
											</router-link>
										</router-link>
									</li><!-- /Resources -->
								</ul>
							</nav>
						</template>
					</collapse>
				</div><!-- /Resources -->
				<!-- =====================
					Assets
					===================== -->
				<div class="aside-block">
					<collapse :use-icon="false">
						<template v-slot:header>
							<div class="aside-block-nav">
								<h6>Assets</h6>
								<i class="feather feather-chevron-down"></i>
							</div>
						</template>
						<template v-slot:body>
							<nav class="aside-nav">
								<ul>
									<!-- Categories -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'categories' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'categories' }">
											<i class="feather feather-tag"></i>
											<span>Categories</span>
										</router-link>
									</li><!-- /Categories -->
									<!-- Media -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'media' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'media' }">
											<i class="feather feather-image"></i>
											<span>Media</span>
										</router-link>
									</li><!-- /Media -->
									<!-- Fields -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'fields' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'fields' }">
											<i class="feather feather-layout"></i>
											<span>Fields</span>
										</router-link>
									</li><!-- /Users -->
									<!-- Users -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'users' || activePage === 'edit-user' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'users' }">
											<i class="feather feather-users"></i>
											<span>Users</span>
										</router-link>
									</li><!-- /Users -->
								</ul>
							</nav>
						</template>
					</collapse>
				</div><!-- /Content -->
				<!-- =====================
					Integrations
					===================== -->
				<div class="aside-block">
					<collapse :use-icon="false">
						<template v-slot:header>
							<div class="aside-block-nav">
								<h6>Integrations</h6>
								<i class="feather feather-chevron-down"></i>
							</div>
						</template>
						<template v-slot:body>
							<nav class="aside-nav">
								<ul>
									<!-- Categories -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'console' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'console' }">
											<i class="feather feather-terminal"></i>
											<span>Console</span>
										</router-link>
									</li><!-- /Categories -->
								</ul>
							</nav>
						</template>
					</collapse>
				</div><!-- /Content -->
				<!-- =====================
					Settings
					===================== -->
				<div class="aside-block">
					<collapse :use-icon="false">
						<template v-slot:header>
							<div class="aside-block-nav">
								<h6>Settings</h6>
								<i class="feather feather-chevron-down"></i>
							</div>
						</template>
						<template v-slot:body>
							<nav class="aside-nav">
								<ul>
									<!-- General -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-general' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-general' }">
											<i class="feather feather-settings"></i>
											<span>General</span>
										</router-link>
									</li><!-- /General -->
									<!-- Code Injection -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-code-injection' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-code-injection' }">
											<i class="feather feather-code"></i>
											<span>Code Injection</span>
										</router-link>
									</li><!-- /Code Injection -->
									<!-- Performance -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-performance' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-performance' }">
											<i class="feather feather-clock"></i>
											<span>Performance</span>
										</router-link>
									</li><!-- /Performance -->
									<!-- SEO & Meta -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-seo-meta' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-seo-meta' }">
											<i class="feather feather-search"></i>
											<span>SEO & Meta</span>
										</router-link>
									</li><!-- /Performance -->
									<!-- Media -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-media' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-media' }">
											<i class="feather feather-film"></i>
											<span>Media</span>
										</router-link>
									</li><!-- /Media -->
									<!-- Themes -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-themes' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-themes' }">
											<i class="feather feather-monitor"></i>
											<span>Themes</span>
										</router-link>
									</li><!-- /Media -->
									<!-- Redirects -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'settings-redirects' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'settings-redirects' }">
											<i class="feather feather-repeat"></i>
											<span>Redirects</span>
										</router-link>
									</li><!-- /Media -->
								</ul>
							</nav>
						</template>
					</collapse>
				</div><!-- /Content -->
			</div>
			<!-- =====================
				Bottom (User)
				===================== -->
			<div class="aside-bottom">
				<collapse :use-icon="false" :show="false" :reverse="true">
					<template v-slot:header>
						<div class="aside-bottom-content" @click="accountActive = !accountActive">
							<img v-if="getProfilePicture['url']" class="avatar" :src="getSite.url + getProfilePicture.url">
							<span v-else class="avatar" v-html="getInitials"></span>
							<!--Aside Bottom User -->
							<div class="aside-bottom-user">
								<!-- User Text -->
								<div class="aside-bottom-user-text">
									<h4>{{ getUserInfo['first_name'] }} {{ getUserInfo['last_name'] }}</h4>
									<p>{{ getUserInfo['email'] }}</p>
								</div><!-- /User Text -->
								<div class="icon icon-naked aside-bottom-chevron">
									<i class="feather feather-chevrons-up" :class="{ 'active' : accountActive }"></i>
								</div>
							</div><!--/Aside Bottom User -->
						</div>
					</template>
					<template v-slot:body>
						<!-- =====================
							Account
							===================== -->
						<div class="aside-block aside-block-account">
							<div class="aside-block-nav">
								<h6>Account</h6>
							</div>
							<nav class="aside-nav">
								<ul>
									<!-- Profile -->
									<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'profile' }" @click="$emit('close', true)">
										<router-link class="aside-nav-link" :to="{ name: 'profile' }">
											<i class="feather feather-user"></i>
											<span>Profile</span>
										</router-link>
									</li><!-- /Profile -->
									<!-- Logout -->
									<li class="aside-nav-item" @click="doLogout">
										<div class="aside-nav-link">
											<i class="feather feather-log-out"></i>
											<span>Logout</span>
										</div>
									</li><!-- /Logout -->
								</ul>
							</nav>
						</div><!-- /Account -->
					</template>
				</collapse>
			</div>
		</div><!-- Aside Container -->
	</aside>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";

export default {
	name: "Sidebar",
	props: {
		open: {
			type: Boolean,
			default: false,
		}
	},
	components: {
		Collapse,
	},
	data: () => ({
		doingAxios: false,
		themeConfig: {},
		resources: {},
		activePage: "",
		collapsed: false,
		accountActive: false,
	}),
	watch: {
		'$route'() {
			this.setActivePage();
		}
	},
	mounted() {
		//console.log(this.getTheme.resources)
	},
	methods: {
		/*
		 * doLogout()
		 */
		doLogout() {
			this.axios.post("/logout", {})
				.then(() => {
					this.$store.commit("logout");
					this.$store.dispatch("getSiteConfig");
					this.$router.push("/login")
				});
		},
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
		/*
		 * collapse()
		 */
		collapse() {
			document.querySelector(".auth-container").classList.toggle("auth-container-collapsed");
			this.collapsed = !this.collapsed;
		},

	},
	computed: {
		/*
		 * getInitials()
		 */
		getInitials() {
			const info = this.getUserInfo
			return info['first_name'].charAt(0) + info['last_name'].charAt(0).toUpperCase();
		},
		/*
		 * getProfilePicture()
		 */
		getProfilePicture() {
			return this.$store.state.profilePicture;
		},
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Variables
$aside-padding-mob: 10px 15px;
$aside-padding-tab: 10px 18px;
$aside-padding-desk: 12px 20px;

.aside {
	$self: &;

	position: fixed;
	left: 0;
	top: 0;
	width: $sidebar-width-mob;
	height: 100%;
	overflow: hidden;
	background-color: $bg-color;
	transform: translateX(-100%);
	transition: transform 300ms ease;
	will-change: transform;
	z-index: 9998;
	box-shadow: 0 0 50px 0 rgba(0, 0, 0, 0.11);

	// Container
	// =========================================================================

	&-container {
		width: 100%;
		height: 100%;
		overflow-y: scroll;
		padding-right: 17px; /* Increase/decrease this value for cross-browser compatibility */
		box-sizing: content-box; /* So the width will be 100% + 17px */
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	// Block
	// =========================================================================

	&-block {
		padding: $aside-padding-mob;
		border-bottom: 1px solid $grey-light;

		&:first-of-type {
			padding-top: 10px;
		}

		&-nav {
			position: relative;
			display: flex;
			justify-content: space-between;
			align-items: center;
			width: 100%;
			padding: 10px 0;
			cursor: pointer;

			i {
				color: rgba($secondary, 0.8);
			}
		}

		&-block-account {
			max-height: 0;
		}
	}

	// Logo
	// =========================================================================

	&-logo {
		display: flex;
		align-items: center;
		padding: $aside-padding-mob $aside-padding-mob 0;

		img {
			width: 26px;
			margin-right: 10px;
		}

		h2 {
			margin-bottom: 0;
			transform: translateY(5px);
		}
	}

	// Bottom
	// =========================================================================

	&-bottom {

		&-content {
			display: flex;
			align-items: center;
			padding: $aside-padding-mob;
			border-top: 1px solid $grey-light;
			cursor: pointer;
		}

		&-chevron {

			i {
				transition: transform 200ms ease;
				will-change: transform;

				&.active {
					transform: rotate(180deg);
				}
			}
		}

		.avatar {
			margin-right: 10px;
		}

		&-user {
			display: flex;
			align-items: center;
			justify-content: space-between;
			width: 100%;

			.icon-naked {
				padding: 6px;
				margin-right: -6px;
				font-size: 18px;
				color: $grey;
				cursor: pointer;
			}

			p {
				margin-bottom: 0;
			}

			h4 {
				line-height: 1;
			}
		}
	}

	// Active
	// =========================================================================

	&-active {
		transform: translateX(0);
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
			margin-bottom: 2px;

			#{$self}-nav-link {
				display: flex;
				align-items: center;
				justify-content: flex-start;
				padding: 8px 14px;
				border-radius: 4px;
				background-color: transparent;
				cursor: pointer;
				transition: background-color 100ms linear;
				will-change: box-shadow;
			}

			span {
				font-size: 14px;
				font-weight: 500;
			}

			span,
			i {
				color: rgba($secondary, 0.6);
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

			#{$self}-nav-link {
				background-color: rgba(#EBEDF9, 1);
			}

			span,
			i {
				color: $primary
			}
		}

		// Hover
		// =========================================================================

		&-item:hover {

			#{$self}-nav-link {
				background-color: rgba($primary, 1);
			}

			span,
			i {
				color: $white;
			}
		}

		// Icon
		// =========================================================================

		&-item-icon {

			#{$self}-nav-link {
				width: 100%;
				display: flex;
				justify-content: space-between;
				align-items: center;
			}

			#{$self}-icon {
				display: flex;
				justify-content: center;
				align-items: center;
				opacity: 0;
				margin-right: 0;
				background-color: rgba($white, 0.2);
				padding: 4px;
				border-radius: 2px;
				z-index: 100;
				transition: opacity 200ms ease;

				i {
					margin: 0;
					width: auto;
				}
			}

			&:hover {

				#{$self}-icon {
					opacity: 1;
				}
			}
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {
		width: $sidebar-width-tab;


		&-block {
			&:first-of-type {
				padding-top: 0;
			}
		}


		&-logo,
		&-block,
		&-bottom &-content {
			padding: $aside-padding-tab;
		}
	}

	// Desktop
	// =========================================================================

	@include media-desk {
		width: $sidebar-width-desk;
		height: 100vh;
		transform: none;
		transition: none;
		z-index: 8;
		box-shadow: none;

		&-logo,
		&-block,
		&-bottom &-content {
			padding: $aside-padding-desk;
		}

		&-block-account {
			border-top: 1px solid $grey-light;
		}

		&:after {
			content: "";
			position: absolute;
			right: -1px;
			top: 0;
			height: 100%;
			width: 2px;
			background-color: $grey-light;
			box-shadow: 0 0 50px 10px rgba($black, 0.11) !important;
		}
	}

	// HD
	// =========================================================================

	@include media-hd {
		width: $sidebar-width-hd;
	}
}

</style>