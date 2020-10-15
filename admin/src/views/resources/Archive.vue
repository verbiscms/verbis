<!-- =====================
	Archive
	===================== -->
<template>
	<section class="">
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>{{ resource['friendly_name'] }}</h1>
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
						<template slot="item">Show all</template>
						<template slot="item">Published</template>
						<template slot="item">Drafts</template>
					</Tabs>
					<!-- =====================
						Posts
						===================== -->
<!--					<transition name="trans-fade-quick" mode="out-in">-->
					<div class="table-wrapper">
						<div class="table-scroll table-with-hover" v-if="posts.length">
							<table class="table archive-table">
								<thead>
									<tr>
										<th>
											<div class="form-checkbox form-checkbox-dark">
												<input type="checkbox" id="archive-check-all" v-model="checkedAll"/>
												<label for="archive-check-all">
													<i class="fal fa-check"></i>
												</label>
											</div>
										</th>
										<th class="archive-table-order" @click="changeOrderBy('title')" :class="{ 'active' : activeOrder === 'title' }">
											<span>Name</span>
											<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
										</th>
										<th class="archive-table-order" @click="changeOrderBy('user_id')" :class="{ 'active' : activeOrder === 'user_id' }">
											<span>Author</span>
											<i class="fas fa-caret-down" :class="{ 'active' : orderBy['user_id'] !== 'asc' }"></i>
										</th>
										<th class="archive-table-order" @click="changeOrderBy('status')" :class="{ 'active' : activeOrder === 'status' }">
											<span>Status</span>
											<i class="fas fa-caret-down" :class="{ 'active' : orderBy['status'] !== 'asc' }"></i>
										</th>
										<th class="archive-table-order" @click="changeOrderBy('published_at')" :class="{ 'active' : activeOrder === 'published_at' }">
											<span>Published at</span>
											<i class="fas fa-caret-down" :class="{ 'active' : orderBy['published_at'] !== 'asc' }"></i>
										</th>
										<th></th>
									</tr>
								</thead>
								<tbody>
									<tr v-for="(item, itemIndex) in posts" :key="item.post.uuid">
										<td class="table-checkbox">
											<div class="form-checkbox form-checkbox-dark">
												<input type="checkbox" :id="item.post.uuid" :value="item.post.id" v-model="checked"/>
												<label :for="item.post.uuid">
													<i class="fal fa-check"></i>
												</label>
											</div>
										</td>
										<!-- Title & Slug -->
										<td class="archive-table-title">
											<router-link :to="{ name: 'editor', params: { id: item.post.id }}">
												<h4>{{ item.post.title }}</h4>
												<p>{{ item.post.slug }}</p>
											</router-link>
										</td>
										<!-- Author -->
										<td class="archive-table-author">
											{{ item.author['first_name'] }} {{ item.author['last_name'] }}
										</td>
										<!-- Status -->
										<td class="archive-table-status">
											<div class="badge capitalize" :class="{ 'badge-warning' : item.post.status  === 'draft' }">{{ item.post.status }}</div>
										</td>
										<!-- Published at -->
										<td class="archive-table-date">
											<span v-if="!item.post['published_at']">Not published</span>
											<span v-else>{{ item.post['published_at'] | moment("dddd, MMMM Do YYYY") }}</span>
										</td>
										<!-- =====================
											Actions
											===================== -->
										<td class="archive-table-actions">
											<Popover :triangle="false" :position="(itemIndex + 1) > (posts.length - 6) ? 'top-left' : 'bottom-left'" @update="updateActions($event, item.post.uuid)" :item-key="item.post.uuid" :active="activeAction">
												<template slot="items">
													<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }}" >
														<i class="feather feather-edit"></i>
														<span>Edit</span>
													</router-link>
													<a  class="popover-item popover-item-icon" :href="getSiteUrl + item.post.slug" target="_blank">
														<i class="feather feather-eye"></i>
														<span>View</span>
													</a>
													<div class="popover-item popover-item-icon" @click="moveToDrafts(item.post.id)">
														<i class="feather feather-archive"></i>
														<span>Move to drafts</span>
													</div>
													<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }, query: { tab: 'meta' }}" >
														<i class="feather feather-external-link"></i>
														<span>Meta</span>
													</router-link>
													<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }, query: { tab: 'seo' }}" >
														<i class="feather feather-search"></i>
														<span>SEO</span>
													</router-link>
													<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }, query: { tab: 'code-injection' }}" >
														<i class="feather feather-code"></i>
														<span>Code Injection</span>
													</router-link>
													<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }, query: { tab: 'insights' }}" >
														<i class="feather feather-bar-chart-2"></i>
														<span>Insights</span>
													</router-link>
													<div class="popover-line"></div>
													<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="deletePost(item.post.id)">
														<i class="feather feather-trash-2"></i>
														<span>Delete</span>
													</div>
												</template>
												<template slot="button">
													<i class="icon icon-square far fa-ellipsis-h" :class="{'icon-square-active' : activeAction === item.post.uuid}"></i>
												</template>
											</Popover>
										</td>
									</tr>
								</tbody>
							</table>
						</div><!-- /Table Scroll -->
						<Alert v-else colour="orange">
							<slot><strong>No {{ resource['friendly_name'] }} available. </strong>To create a new one, click the plus sign above.</slot>
						</Alert>
<!--					</transition>-->
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<p>Sort out the pagination ui, add singular and plural names to config, add doing axios and transition, sort out alert, it looks shit, need to do action buttons</p>
					<div v-if="paginationObj">
						<button v-if="paginationObj['prev']" class="btn" @click="setPagination('prev')">Prev</button>
						<button v-if="paginationObj['next']" class="btn" @click="setPagination('next')">Next</button>
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
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
import Popover from "@/components/misc/Popover";

export default {
	name: "Pages",
	title: "Archive",
	components: {
		Alert,
		Breadcrumbs,
		Tabs,
		Popover,
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
		this.getPosts();
		this.setResource();
	},
	watch: {
		'$route.params.resource': function() {
			this.setResource();
			this.getPosts();
		},
	},
	methods: {
		updateActions(e, uuid) {
			this.activeAction = e ? uuid : "";
		},
		/*
		 * getPosts()
		 * Obtain the posts or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getPosts() {
			const resource = this.$route.params.resource;
			const url = resource === "pages" ? "/posts" : "/resource/" + resource;
			this.axios.get(`${url}?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
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
		 * setResource()
		 * Set the resource from the query parameter, if none defined,
		 * set default page 'resource'.
		 */
		setResource() {
			const resource = this.getResources[this.$route.params.resource]
			this.resource = resource === undefined ? {
				"name": "page",
				"friendly_name": "Page",
				"slug": "",
				"icon": 'fal fa-file'
			} : resource
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
			this.getPosts()
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
					filter = '{"status":[{"operator":"=", "value": "published" }]}';
					break;
				}
				case 3: {
					filter = '{"status":[{"operator":"=", "value": "draft" }]}';
					break;
				}
			}
			this.filter = filter;
			this.getPosts()
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain posts.
		 */
		setPagination(direction) {
			this.pagination = "";
			const page = direction === "next" ? this.paginationObj.page + 1 : this.paginationObj.page - 1
			this.pagination = `page=${page}`
			this.getPosts();
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

	.archive {

		// Table
		// =========================================================================

		&-table {

			// Props
			// =========================================================================

			tbody tr:hover {

				.icon-square {
					background-color: $white;
				}
			}

			// Order
			// =========================================================================

			&-order {
				cursor: pointer;
				user-select: none;

				span {
					transition: color 180ms ease;
					will-change: color;
				}

				i {
					color: rgba($secondary, 0.7);
					font-size: 12px;
					margin-left: 4px;
					transition: 180ms ease, color 180ms ease;
					will-change: transform, color;

					&.active {
						transform: rotate(180deg);
					}
				}

				&.active {

					i,
					span {
						color: $primary;
					}
				}
			}

			// Title
			// =========================================================================

			&-title {
				width: 400px;

				h4 {
					transition: color 200ms ease;
					will-change: color;
				}

				&:hover {

					h4 {
						color: $primary;
					}
				}
			}

			// Author
			// =========================================================================

			&-author {
				width: 215px;
			}

			// Actions
			// =========================================================================

			&-actions {

				.icon-square {
					transition: background-color 160ms ease;
					will-change: background-color;
				}

				i {
					font-size: 16px;
				}
			}
		}
	}

</style>