<!-- =====================
	Archive
	===================== -->
<template>
	<section>
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
								<div class="form-select-cont form-input header-hide-mob">
									<select class="form-select" v-model="bulkType" v-if="activeTab === 4">
										<option value="" disabled selected>Bulk actions</option>
										<option value="restore">Restore</option>
										<option value="delete">Delete permanently</option>
									</select>
									<select class="form-select" v-model="bulkType" v-else-if="activeTab === 3">
										<option value="" disabled selected>Bulk actions</option>
										<option value="publish">Publish</option>
										<option value="bin">Move to bin</option>
									</select>
									<select class="form-select" v-model="bulkType" v-else>
										<option value="" disabled selected>Bulk actions</option>
										<option value="publish">Publish</option>
										<option value="draft">Move to drafts</option>
										<option value="bin">Move to bin</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white header-hide-mob" :class="{ 'btn-loading' : isDoingBulk }" @click.prevent="doBulkAction">Apply</button>
								<router-link class="btn btn-icon btn-orange btn-text-mob" :to="{ name: 'editor', params: { id: 'new' }, query: { resource: resource['name'] }}">
									<i class="fal fa-plus"></i>
									<span>New {{ resource['singular_name'] }}</span>
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
						<template slot="item">Bin</template>
					</Tabs>
					<!-- =====================
						Posts
						===================== -->
					<div v-if="!doingAxios">
						<transition name="trans-fade-quick" mode="out-in">
							<div class="table-wrapper" v-if="posts.length">
								<div class="table-scroll table-with-hover">
									<table class="table archive-table">
										<thead>
											<tr>
												<th class="table-header-checkbox">
													<div class="form-checkbox form-checkbox-dark">
														<input type="checkbox" id="archive-check-all" v-model="checkedAll"/>
														<label for="archive-check-all">
															<i class="fal fa-check"></i>
														</label>
													</div>
												</th>
												<th class="table-order" @click="changeOrderBy('title')" :class="{ 'active' : activeOrder === 'title' }">
													<span>Name</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
												</th>
												<th class="table-order" @click="changeOrderBy('user_id')" :class="{ 'active' : activeOrder === 'user_id' }">
													<span>Author</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['user_id'] !== 'asc' }"></i>
												</th>
												<th class="table-order" @click="changeOrderBy('status')" :class="{ 'active' : activeOrder === 'status' }">
													<span>Status</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['status'] !== 'asc' }"></i>
												</th>
												<th>
													<span>Category</span>
												</th>
												<th class="table-order" @click="changeOrderBy('published_at')" :class="{ 'active' : activeOrder === 'published_at' }">
													<span>Published at</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['published_at'] !== 'asc' }"></i>
												</th>
												<th></th>
											</tr>
										</thead>
										<tbody>
											<tr v-for="(item, itemIndex) in posts" :key="item.post.uuid">
												<!-- Checkbox -->
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
													<router-link :to="{ name: 'editor', params: { id: item.post.id }, query: {resource: item.post.resource ? item.post.resource : '' }}">
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
													<div class="badge capitalize" :class="{
														'badge-yellow' : item.post.status  === 'draft',
														'badge-green' : item.post.status  === 'published',
														'badge-orange' : item.post.status  === 'bin',
													}">{{ item.post.status }}</div>
												</td>
												<!-- Category -->
												<td>
													<span v-if="item.categories.length">{{ item.categories[0].name }}</span>
													<span v-else>No category</span>
												</td>
												<!-- Published at -->
												<td class="archive-table-date">
													<span v-if="!item.post['published_at']">Not published</span>
													<span v-else>{{ item.post['published_at'] | moment("dddd, MMMM Do YYYY") }}</span>
												</td>
												<!-- Actions -->
												<td class="table-actions">
													<Popover :triangle="false"
															:classes="(itemIndex + 1) > (posts.length - 4) ? 'popover-table popover-table-top' : 'popover-table popover-table-bottom'"
															@update="updateActions($event, item.post.uuid)"
															:item-key="item.post.uuid"
															:active="activeAction">
														<template slot="items">
															<router-link class="popover-item popover-item-icon" :to="{ name: 'editor', params: { id: item.post.id }, query: {resource: item.post.resource ? item.post.resource : '' }}">
																<i class="feather feather-edit"></i>
																<span>Edit</span>
															</router-link>
															<a class="popover-item popover-item-icon" :href="getSiteUrl + item.post.slug" target="_blank">
																<i class="feather feather-eye"></i>
																<span>View</span>
															</a>
															<div v-if="item.post.status === 'published'" class="popover-item popover-item-icon" @click="updateStatus(item.post.id, 'draft')">
																<i class="feather feather-archive"></i>
																<span>Move to drafts</span>
															</div>
															<div v-else-if="item.post.status === 'draft'" class="popover-item popover-item-icon" @click="updateStatus(item.post.id, 'published')">
																<i class="feather feather-archive"></i>
																<span>Publish</span>
															</div>
															<div v-else-if="item.post.status === 'bin'" class="popover-item popover-item-icon" @click="updateStatus(item.post.id, 'draft')">
																<i class="feather feather-archive"></i>
																<span>Restore</span>
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
															<div v-if="item.post.status === 'bin'" class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(item.post);">
																<i class="feather feather-trash-2"></i>
																<span>Delete</span>
															</div>
															<div v-else class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="updateStatus(item.post.id, 'bin')">
																<i class="feather feather-trash-2"></i>
																<span>Move to bin</span>
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
							</div><!-- /Table Wrapper -->
							<Alert v-else colour="orange">
								<slot>
									<h3>No {{ activeTabName === "all" ? "" : (activeTabName === "bin" ? "binned" : activeTabName ) }} {{ resource['friendly_name'].toLowerCase() }} available. </h3>
									<p>To create a new one, click the button above.</p>
								</slot>
							</Alert>
						</transition>
					</div><!-- /Doing Axios -->
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
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning">
			<template slot="button">
				<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="deletePost(false);">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="selectedPost">Are you sure want to delete this {{ resource['singular_name'] }}?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} {{ resource['friendly_name'] }}?</p>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Alert from "@/components/misc/Alert";
import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Modal from "@/components/modals/General";
import Popover from "@/components/misc/Popover";
import Pagination from "@/components/misc/Pagination";
import Tabs from "../../components/misc/Tabs";

export default {
	name: "Pages",
	title: "Archive",
	components: {
		Alert,
		Breadcrumbs,
		Modal,
		Popover,
		Pagination,
		Tabs,
	},
	data: () => ({
		doingAxios: true,
		resource: {},
		posts: [],
		selectedPost: false,
		paginationObj: {},
		activeTab: 1,
		activeTabName: "all",
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
		bulkType: "",
		checked: [],
		activeAction: "",
		showDeleteModal: false,
		selectedDeleteId: null,
		isDoingBulk: false,
		isDeleting: false,
	}),
	mounted() {
		this.filterTabs(1);
		this.setResource();
	},
	watch: {
		'$route.params.resource': function() {
			this.setResource();
			this.getPosts();
			this.activeTab = 1;
		},
	},
	methods: {
		/*
		 * getPosts()
		 * Obtain the posts or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getPosts() {
			const resource = this.$route.params.resource;
			const url = /resources/ + resource;

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
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * deletePost()
		 */
		deletePost() {
			this.isDeleting = true;

			const toDelete = this.selectedPost ? [this.selectedPost.id] : this.checked;

			const promises = [];
			toDelete.forEach(id => {
				promises.push(this.deleteUserAxios(id));
			});

			// Send all requests
			Promise.all(promises)
				.then(() => {
					const successMsg = toDelete.length === 1 ? `${this.resource['singular_name']} deleted successfully.` : `${this.resource['name']} deleted successfully.`
					this.$noty.success(successMsg);
					this.getPosts();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.activeAction = "";
					this.checked = [];
					this.checkedAll = false;
					this.showDeleteModal = false;
					this.bulkType = "";
					this.isDeleting = false;
					this.selectedPost = false;
				});
		},
		/*
		 * async deletePostAxios()
		 */
		async deleteUserAxios(id) {
			return await this.axios.delete("/posts/" + id);
		},
		/*
		 * handleDelete()
		 * Pushes ID to array on single click (not bulk action) &
		 * show the delete post modal.
		 */
		handleDelete(post) {
			this.selectedPost = post;
			this.showDeleteModal = true;
		},
		/*
		 * updateStatus()
		 */
		updateStatus(id = false, status = 'draft') {
			let checkedArr = [];
			if (id) {
				checkedArr.push(id);
			} else {
				checkedArr = this.checked;
			}

			const promises = [];
			checkedArr.forEach(id => {
				const post =  this.getPostsById(id).post
				post.status = status;
				promises.push(this.updateStatusAxios(id, post));
			});

			Promise.all(promises)
				.then(() => {
					const successMsg = checkedArr.length === 1 ? `${this.resource['singular_name']} updated successfully.` : `${this.resource['name']} updated successfully.`
					this.$noty.success(successMsg);
					this.getPosts();
				})
				.catch((err) => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.activeAction = "";
					this.checked = [];
					this.checkedAll = false;
					this.showDeleteModal = false;
					this.bulkType = "";
					setTimeout(() => {
						this.isDoingBulk = false;
					}, this.timeoutDelay);
				});
		},
		/*
		 * async deletePostAxios()
		 */
		async updateStatusAxios(id, post) {
			return await this.axios.put("/posts/" + id, post);
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
				"friendly_name": "Pages",
				"singular_name": "Page",
				"slug": "",
				"icon": 'fal fa-file'
			} : resource;
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
			this.order = column + "," + this.orderBy[column];
			this.getPosts();
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain posts.
		 */
		filterTabs(tab) {
			this.pagination = "page=1";
			this.activeTab = tab;
			let filter = "";
			switch (tab) {
				case 1: {
					this.activeTabName = "all";
					filter = '{"status":[{"operator":"NOT LIKE", "value": "bin" }]}';
					break;
				}
				case 2: {
					this.activeTabName = "published";
					filter = '{"status":[{"operator":"=", "value": "published" }]}';
					break;
				}
				case 3: {
					this.activeTabName = "draft";
					filter = '{"status":[{"operator":"=", "value": "draft" }]}';
					break;
				}
				case 4: {
					this.activeTabName = "bin";
					filter = '{"status":[{"operator":"=", "value": "bin" }]}';
					break;
				}
			}
			this.checkedAll = false;
			this.filter = filter;
			this.getPosts();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain posts.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getPosts();
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
			this.isDoingBulk = true;

			// Check if there no items
			if (!this.checked.length) {
				this.$noty.warning("Select items in order to apply bulk actions");
				setTimeout(() => {
					this.isDoingBulk = false;
				}, this.timeoutDelay);
				return;
			}

			// Move to drafts / restore
			if (this.bulkType === "draft" || this.bulkType === "restore") {
				this.updateStatus(false, 'draft');
			// Publish
			} else if (this.bulkType === "publish") {
				this.updateStatus(false, 'published');
			// Move to bin
			} else if (this.bulkType === "bin") {
				this.updateStatus(false, 'bin');
			// Delete
			} else if (this.bulkType === "delete") {
				this.showDeleteModal = true;
			} else {
				this.$noty.warning("Select a bulk action.");
				this.isDoingBulk = false;
			}

			setTimeout(() => {
				this.isDoingBulk = false;
			}, this.timeoutDelay);
		},
		/*
		 * getPostsById()
		 */
		getPostsById(id) {
			return this.posts.find(p => p.post.id === id)
		},
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
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return this.checked.length === this.posts.length;
			},
			set(value) {
				if (value) {
					this.checked = this.posts.map(m => {
						return m.post.id;
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

	// Pagination
	// =========================================================================

	&-pagination {

		&-trans {
			&-enter-active, &-leave-active {
				transition: opacity 200ms;
				transition-delay: 200ms;
			}

			&-enter, &-leave-to /* .fade-leave-active below version 2.1.8 */ {
				opacity: 0;
				transition-delay: 200ms;
			}
		}
	}

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
	}
}

</style>