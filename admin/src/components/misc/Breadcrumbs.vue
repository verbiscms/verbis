<!-- =====================
	Section
	===================== -->
<template>
	<nav class="breadcrumbs">
		<ul class="breadcrumbs-list">
			<li class="breadcrumbs-item" v-for="breadcrumb in breadcrumbs" v-bind:key="breadcrumb.url">
				<router-link class="breadcrumbs-link" :to="breadcrumb.url" :class="{ 'breadcrumbs-link-active' : breadcrumb.active }">{{ breadcrumb.name }}
				</router-link>
			</li>
		</ul>
	</nav>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Breadcrumbs",
	data: () => ({
		breadcrumbs: []
	}),
	beforeMount() {
		this.updateList()
	},
	watch: {
		'$route'() {
			this.updateList()
		}
	},
	methods: {
		/*
		 * updateList()
		 *
		 */
		updateList() {
			this.breadcrumbs = [];

			const fullPath = this.$route.fullPath,
				pathArr = fullPath.split("/");

			if (this.$route.name === "home") {

				this.breadcrumbs.push({
					name: "Home",
					url: "/",
					active: true,
				})

			} else {
				let bPath = "",
					bPathLoop = true;

				pathArr.forEach((path) => {

					bPath += path + "/";

					let temp;
					if (bPathLoop) {
						temp = bPath;
						bPathLoop = false;
					} else {
						temp = bPath.replace(/\/$/, "")
					}

					path = path.split("?")[0];

					this.breadcrumbs.push({
						name: path === "" ? "Home" : this.capitalize(path),
						url: temp,
						active: this.$route.fullPath === temp,
					})
				});
			}
		},
		/*
		 * capitalize()
		 * Capitalize the first letter of the breadcrumb..
		 */
		capitalize(str) {
			return str.replace(/(?:^|\s|["'([{])+\S/g, match => match.toUpperCase());
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.breadcrumbs {
	$self: &;

	margin-top: 6px;

	// List
	// =========================================================================

	&-list {
		display: flex;
		align-items: center;

		li:last-child #{$self}-link:after {
			display: none;
		}
	}

	// Link
	// =========================================================================

	&-link {
		display: flex;
		align-items: center;
		color: $grey;
		font-size: 13px;

		&:after {
			content: url("~@/assets/images/breadcrumb-arrow.svg");
			position: relative;
			display: block;
			margin: 0 8px;
		}

		&-active {
			color: $primary;
		}

		&:hover {
			color: $secondary;
		}
	}

	// Small Phones
	// =========================================================================

	@media screen and (max-width: 350px) {

		&-link {
			font-size: 11px;

			&:after {
				margin: 0 6px;
			}
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {
		margin-top: 8px;

		&-link {
			font-size: 0.875rem;

			&:after {
				margin: 0 10px;
			}
		}
	}
}

</style>