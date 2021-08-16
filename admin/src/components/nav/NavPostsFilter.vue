<!-- =====================
	Post Search
	===================== -->
<template>
	<div class="filter">
		<!-- Search Posts -->
		<SearchPosts class="filter-search" size="small" @input="updatePosts"></SearchPosts>
		<!-- Search Results -->
		<el-checkbox-group class="filter-results" v-model="selected">
			<el-checkbox class="filter-results-check" v-for="(post, index) in posts" :key="index" :label="index">
				{{ post.post.title }}
			</el-checkbox>
		</el-checkbox-group>
		<!-- Add to Menu -->
		<el-button class="filter-btn" size="medium" @click="update" plain>
			Add to Menu
		</el-button>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import SearchPosts from "@/components/search/PostSearch";

export default {
	name: "NavPostsFilter",
	components: {
		SearchPosts,
	},
	props: {},
	data: () => ({
		posts: [],
		selected: [],
		items: [],
		noResults: false,
	}),
	methods: {
		/*
		 * handleChange()
		 * Updates the selectedItems array
		 */
		handleChange(items) {
			this.items = items;
		},
		/*
		 * update()
		 * Updates the parent with the selected post
		 * items.
		 */
		update() {
			this.$emit("update", this.selected.map(id => this.posts[id]));
		},
		updatePosts(posts) {
			this.posts = [];
			this.noResults = false;
			if (!posts.length) {
				this.noResults = true;
				return;
			}
			this.posts = posts;
		}
	},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

::v-deep {


	// Filter
	// =========================================================================

	.filter {

		// Search
		// =========================================================================

		&-search {
			margin-bottom: 10px;
		}

		// Table
		// =========================================================================

		&-results {
			display: flex;
			flex-direction: column;

			&-check {
				margin-bottom: 4px;
			}
		}

		// Button
		// =========================================================================

		&-btn {
			display: block;
			margin-top: 1rem;
			margin-left: auto;
			margin-right: 0;
		}
	}
}


</style>
