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
	created() {
		const breadcrumbs = this.$options
		console.log(breadcrumbs)
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

	margin-top: 10px;

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

		&:after {
			content: url("~@/assets/images/breadcrumb-arrow.svg");
			position: relative;
			display: block;
			margin: 0 10px;
		}

		&-active {
			color: $primary;
		}

		&:hover {
			color: $secondary;
		}
	}
}

</style>