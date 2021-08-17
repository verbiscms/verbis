<!-- =====================
	Post Search
	===================== -->
<template>
	<div class="search">
		<el-input
			:size="size"
			class="search-input"
			prefix-icon="el-icon-search"
			placeholder="Search posts"
			v-model="input"
			@input="search">
		</el-input>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "SearchPosts",
	props: {
		size: {
			type: String,
			default: "medium",
		},
		limit: {
			type: Number,
			default: 15,
		}
	},
	data: () => ({
		input: "",
		awaitingSearch: false,
	}),
	methods: {
		/**
		 * search()
		 * Awaits for a timeout and then calls
		 * the makeRequest() function.
		 */
		search() {
			if (!this.awaitingSearch) {
				setTimeout(() => {
					this.makeRequest();
					this.awaitingSearch = false;
				}, 400); // Add delay
			}
			this.awaitingSearch = true;
		},
		/**
		 * makeRequest()
		 * Searches the API for the posts using the
		 * query string param. Updates the parent
		 * when search results come back.
		 */
		makeRequest() {
			this.axios.get(`/posts?limit=${this.limit}&filter={"title":[{"operator":"LIKE", "value":"${this.input}"}]}`, {
				paramsSerializer: function (params) {
					return params;
				}
			})
			.then(res => {
				let data = res.data.data;
				if (!data.length) {
					data = [];
				}
				this.$emit("input", data);
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

	.search {

		&-input {
			width: 100%;
		}
	}

</style>
