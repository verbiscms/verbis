<!-- =====================
	Media
	===================== -->
<template>
	<section class="">
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Media</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<div class="form-select-cont form-input">
									<select class="form-select" v-model="bulkType">
										<option value="" disabled selected>Bulk actions</option>
										<option value="drafts">Move to drafts</option>
										<option value="delete">Delete</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white" :class="{ 'btn-loading' : savingBulk }" @click.prevent="doBulkAction">Apply</button>
								<router-link class="btn btn-icon btn-orange" :to="{ name: 'editor', params: { id: 'new' }, query: { resource: resource['name'] }}">
									<i class="fal fa-plus"></i>
								</router-link>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<!-- =====================
						Tabs
						===================== -->
					<Tabs @update="filterTabs">
						<template slot="item">JPG's</template>
						<template slot="item">PNG's</template>
						<template slot="item">Files</template>
					</Tabs>
					<!-- =====================
						Posts
						===================== -->
					<div v-if="!doingAxios">
						<transition name="trans-fade-quick" mode="out-in">

							<Alert colour="orange">
								<slot>
									<h3>No media items available. </h3>
									<p>To create a new one, click the plus sign above.</p>
								</slot>
							</Alert>
						</transition>
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<transition name="archive-pagination-trans">
				<div class="row" v-if="!doingAxios && paginationObj">
					<div class="col-12">
						<Pagination :pagination="paginationObj" @update="setPagination"></Pagination>
					</div><!-- /Col -->
				</div><!-- /Row -->
			</transition>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Tabs from "../../components/misc/Tabs";
import Alert from "@/components/misc/Alert";
import Pagination from "@/components/misc/Pagination";

export default {
	name: "Pages",
	title: "Archive",
	components: {
		Pagination,
		Alert,
		Breadcrumbs,
		Tabs,
	},
	data: () => ({
		doingAxios: true,
		activeTab: 0,
		resource: {},
		posts: [],
		paginationObj: {},
		order: "",
		orderBy: {
			title: "asc",
			user_id: "asc",
			status: "asc",
			published_at: "asc",
		},
		activeOrder: "",
		filter: "",
		pagination: "",
		savingBulk: false,
		bulkType: "",
		checked: [],
		activeAction: "",
	}),
	mounted() {
		this.getMedia();
	},
	methods: {
		/*
		 * getMedia()
		 * Obtain the media or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getMedia() {
			this.axios.get(`media?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.posts = {};
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					this.posts = res.data.data;
				})
				.catch(err => {
					console.log(err)
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * changeOrderBy()
		 * Update the order by object when clicked, obtain posts.
		 */
		changeOrderBy(column) {
			this.activeOrder = column;
			if (this.orderBy[column] === "desc" || this.orderBy[column] === "") {
				this.$set(this.orderBy, column, 'asc');
			} else {
				this.$set(this.orderBy, column, 'desc');
			}
			this.order = column + "," + this.orderBy[column]
			this.getMedia();
		},
		/*
		 * getPostsById()
		 */
		getPostsById(id) {
			return this.posts.find(p => p.post.id === id)
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain posts.
		 */
		filterTabs($event) {
			let filter = ""
			switch ($event) {
				case 2: {
					//filter = '{"status":[{"operator":"=", "value": "published" }]}';
					break;
				}
				case 3: {
					//filter = '{"status":[{"operator":"=", "value": "draft" }]}';
					break;
				}
			}
			this.filter = filter;
			this.getMedia();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain posts.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getMedia();
		},
		/*
		 * updateActions()
		 *  Update the action uuid for clearing the popover.
		 */
		updateActions(e, uuid) {
			this.activeAction = e ? uuid : "";
		},
		/*
		 * doBulkAction()
		 * When bulk action is clicked, this function will call drafts or delete.
		 * Validation on bulk type action and checked length performed.
		 */
		doBulkAction() {
			this.savingBulk = true;
			if (this.bulkType === "") {
				this.$noty.warning("Select a bulk action.");
				this.savingBulk = false;
				return
			}
			if (!this.checked.length) {
				this.$noty.warning("Select items in order to apply bulk actions");
				this.savingBulk = false;
				return
			}
			// Move to drafts
			if (this.bulkType === "drafts") {
				this.moveToDrafts(false);
			} else if (this.bulkType === "delete") {
				this.deletePost(false);
			}
		},
		/*
		 * moveToDrafts()
		 */
		moveToDrafts(id = false) {
			let checkedArr = [];
			if (id) {
				checkedArr.push(id);
			} else {
				checkedArr = this.checked;
			}
			checkedArr.forEach(id => {
				const post =  this.getPostsById(id).post
				post.status = "draft"
				this.axios.put("/posts/" + id, post)
					.then(() => {
						this.$noty.success("Posts updated successfully.");
						this.getPosts();
					})
					.catch((err) => {
						console.log(err);
						this.$noty.error("Error occurred, please refresh the page.");
					})
					.finally(() => {
						this.savingBulk = false;
					});
			});
		},
		/*
		 * deletePost()
		 */
		deletePost(id = false) {
			let checkedArr = [];
			if (id) {
				checkedArr.push(id);
			} else {
				checkedArr = this.checked;
			}
			checkedArr.forEach(id => {
				this.axios.delete("/posts/" + id)
					.then(() => {
						this.$noty.success("Posts deleted successfully.");
						this.getPosts();
					})
					.catch(err => {
						console.log(err);
						this.$noty.error("Error occurred, please refresh the page.");
					})
					.finally(() => {
						this.savingBulk = false;
					});
			});
		}
	},
	computed: {
		/*
		 * getResources()
		 * Get the theme resources from store.
		 */
		getResources() {
			return this.$store.state.theme.resources;
		},
		/*
		 * getSiteUrl()
		 * Get the site url from the store for previewing.
		 */
		getSiteUrl() {
			return this.$store.state.site.url;
		},
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return false;
			},
			set(value) {
				if (value) {
					this.checked = this.posts.map(m => {
						return m.post.id
					});
					return;
				}
				this.checked = [];
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.media {

}

</style>