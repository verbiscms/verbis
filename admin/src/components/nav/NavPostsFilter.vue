<!-- =====================
	Post Search
	===================== -->
<template>
	<div class="filter">
		<el-tabs type="border-card" v-model="activeTab" @tab-click="handleTabClick">
			<!-- Most Recent -->
			<el-tab-pane :loading="loading" :name="RECENT_TAB" label="Most Recent"></el-tab-pane>
			<!-- View All -->
			<el-tab-pane :name="ALL_TAB" label="View All"></el-tab-pane>
			<!-- Search Posts -->
			<el-tab-pane :name="SEARCH_TAB" label="Search">
				<SearchPosts class="filter-search" :class="{ 'filter-search-margin' : posts.length }" size="small" @input="updatePosts"></SearchPosts>
			</el-tab-pane>
			<!-- Search Results -->
			<el-checkbox-group class="filter-results" v-model="selected" @change="handlePostChange">
				<el-checkbox class="filter-results-check" v-for="(post, index) in posts" :key="index" :label="index">
					{{ post.post.title }}
				</el-checkbox>
			</el-checkbox-group>
			<!-- Pagination -->
			<el-pagination v-if="showPagination" small layout="prev, pager, next" :total="posts.length"></el-pagination>
			<!-- Add to Menu -->
			<div class="filter-btn-cont" v-if="posts.length">
				<el-checkbox :indeterminate="isIndeterminate" @change="handleCheckAllChange" v-model="checkAll">Select All</el-checkbox>
				<el-button class="filter-btn" size="small" @click="update" plain>
					Add to Menu
				</el-button>
			</div>
		</el-tabs>
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
		LIMIT: 20,
		RECENT_TAB: "recent",
		ALL_TAB: "all",
		SEARCH_TAB: "search",
		activeTab: "",
		posts: [],
		selected: [],
		items: [],
		noResults: true,
		isIndeterminate: false,
		checkAll: false,
		loading: true,
		showPagination: false,
	}),
	mounted() {
		this.getPosts();
		this.activeTab = this.RECENT_TAB;
	},
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
		},
		handlePostChange(value) {
			let checkedCount = value.length;
			this.checkAll = checkedCount === this.posts.length;
			this.isIndeterminate = checkedCount > 0 && checkedCount < this.posts.length;
		},
		handleCheckAllChange(val) {
			this.selected = val ? this.posts.map((post, index) => index) : [];
			this.isIndeterminate = false;
		},

		handleTabClick(tab) {
			this.selected = [];
			this.isIndeterminate = false;
			this.checkAll = false;

			switch (tab.name) {
				case this.RECENT_TAB:
					this.getPosts();
					break;
				case this.ALL_TAB:
					this.getPosts();
					break;
				case this.SEARCH_TAB:
					this.posts = [];
					break;
			}
		},
		getPosts() {
			this.axios.get(`/posts?limit=${this.LIMIT}`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
				.then(res => {
					this.posts = res.data.data;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
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

			&-margin {
				margin-bottom: 10px;
			}
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

			&-cont {
				margin-top: 10px;
				display: flex;
				align-items: center;
				justify-content: space-between;
			}
		}
	}
}


</style>
