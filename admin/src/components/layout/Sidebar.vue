<!-- =====================
	Sidebar
	===================== -->
<template>
	<aside class="aside">
		<div class="aside-cont">
			<div class="aside-info">
				<!-- Logo -->
				<router-link class="aside-logo" :to="{ name: 'home' }" v-if="getLogo">
					<img :src="globalBasePath + getLogo">
				</router-link><!-- /Col -->
				<h2>{{ themeConfig.name }}</h2>
				<p>Admin CMS</p>
			</div>
			<!-- Navigation -->
			<nav class="aside-nav">
				<ul>
					<!-- Pages -->
					<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'pages' }">
						<router-link class="aside-nav-link" to="/content/pages">
							<i class="fal fa-file"></i>
							<span>Pages</span>
						</router-link>
					</li><!-- /Pages -->
					<!-- Resources -->
					<li class="aside-nav-item" v-for="(resource) in resources" v-bind:key="resource.name" :class="{ 'aside-nav-item-active' : activePage === resource.name.toLowerCase() }">
						<router-link class="aside-nav-link" :to="'/content/' + resource.options.slug">
							<i v-if="resource.options.icon" class="fa" :class="resource.options.icon"></i>
							<i v-else class="fal fa-file"></i>
							<span>{{ resource.name }}</span>
						</router-link>
					</li><!-- /Resources -->
					<!-- Media -->
					<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'analytics' }">
						<router-link class="aside-nav-link" to="/analytics">
							<i class="fal fa-chart-line"></i>
							<span>Analytics</span>
						</router-link>
					</li><!-- /Media -->
					<!-- Media -->
					<li class="aside-nav-item" :class="{ 'aside-nav-item-active' : activePage === 'media' }">
						<router-link class="aside-nav-link" to="media">
							<i class="fal fa-images"></i>
							<span>Media</span>
						</router-link>
					</li><!-- /Media -->
					<!-- Users -->
					<li class="aside-nav-item" v-if="getUserInfo.accessLevel === 2" :class="{ 'aside-nav-item-active' : activePage === 'users' }">
						<router-link class="aside-nav-link" to="/pages">
							<i class="fal fa-users"></i>
							<span>Users</span>
						</router-link>
					</li><!-- /Users -->
					<!-- Settings -->
					<li class="aside-nav-item">
						<router-link class="aside-nav-link" to="/pages" :class="{ 'aside-nav-item-active' : activePage === 'settings' }">
							<i class="fal fa-cogs"></i>
							<span>Settings</span>
						</router-link>
					</li><!-- /Settings -->
				</ul>
			</nav>
			<!-- Controls -->
			<nav class="aside-controls">
				<h2 @click="doLogout" style="background-color: red; display: block; cursor: pointer;">Logout temp</h2>
				<ul>
					<li class="aside-controls-item">
<!--						<router-link class="aside-nav-link" to="/pages">-->
<!--							<i class="fal fa-cogs"></i>-->
<!--							<span>Settings</span>-->
<!--						</router-link>-->
					</li>
				</ul>
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
		}),
		beforeMount() {
			this.getThemeConfig();
			this.getResourceData();
		},
		watch: {
			'$route'() {
				this.setActivePage()
			}
		},
		methods: {
			getThemeConfig() {
				this.axios.get("/theme/config")
					.then(res => {
						this.themeConfig = res.data.data
					})
			},
			getResourceData() {
				this.axios.get("/resources")
				.then(res => {
					this.resources = res.data.data
					this.$store.commit("setResources", res.data.data)
				})
			},
			doLogout() {
				this.axios.post("/logout", {})
				.then(res => {
					console.log(res);
					this.$store.commit("logout")
					location.reload()
				});
			},
			setActivePage() {
				const resource = this.$route.params.resource
				this.activePage = resource === undefined ? this.$route.name : resource
			}
		},
		computed: {
			getUserInfo() {
				return this.$store.state.userInfo
			},
			getLogo() {
				return this.$store.state.logo;
			},
		},
	}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Variables
	$aside-padding: 30px;

	.aside {
		position: fixed;
		display: flex;
		flex-direction: column;
		align-items: center;
		left: 0;
		top: 0;
		width: $sidebar-width;
		height: 100vh;
		background-color: $bg-color;
		z-index: 8;

		// Container
		// =========================================================================

		&-cont {
			width: 100%;
			padding: $auth-container-padding-y $aside-padding $aside-padding;
		}

		// Info
		// =========================================================================

		&-info {
			display: flex;
			flex-direction: column;
			align-items: center;
			margin-bottom: 2rem;

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
			width: 130px;

			img {
				width: 100%;
			}
		}

		// Nav
		// =========================================================================

		&-nav {

			// Item
			// =========================================================================

			&-item {
				margin-bottom: 14px;

				a {
					display: flex;
					align-items: center;
					justify-content: flex-start;
					padding: 14px 20px;
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