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
							<h1>Posts</h1>
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
								<button class="btn btn-icon btn-orange">
									<i class="fal fa-plus"></i>
								</button>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Tabs -->
			<div class="row">
				<div class="col-12">
					<Tabs @update="activeTab = $event - 1">
						<template slot="item">Show all</template>
						<template slot="item">Published</template>
						<template slot="item">Drafts</template>
<!--						<template slot="item">Scheduled</template>-->
					</Tabs>
<!--					<transition name="trans-fade" mode="out-in">-->
					<!-- =====================
						Posts
						===================== -->
					<div class="table-scroll">
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
									<th class="archive-table-order" @click="changeOrderBy('title')">
										Name<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
									</th>
									<th class="archive-table-order" @click="changeOrderBy('user_id')">
										Author<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
									</th>
									<th class="archive-table-order" @click="changeOrderBy('status')">
										Status<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
									</th>
									<th class="archive-table-order" @click="changeOrderBy('published_at')">
										Published at<i class="fas fa-caret-down" :class="{ 'active' : orderBy['title'] !== 'asc' }"></i>
									</th>
									<th>Actions</th>
								</tr>
							</thead>
							<tbody>
								<tr v-for="item in posts" :key="item.post.uuid">
									<td class="table-checkbox">
										<div class="form-checkbox form-checkbox-dark">
											<input type="checkbox" :id="item.post.uuid" :value="item.post.id" v-model="checked"/>
											<label :for="item.post.uuid">
												<i class="fal fa-check"></i>
											</label>
										</div>
									</td>
									<!-- Title & Slug -->
									<td>
										<router-link class="archive-table-title" :to="{ name: 'editor', params: { id: item.post.id }}">
											<h4>{{ item.post.title }}</h4>
											<p>{{ item.post.slug }}</p>
										</router-link>
									</td>
									<!-- Author -->
									<td>
										{{ item.author['first_name'] }} {{ item.author['last_name'] }}
									</td>
									<!-- Status -->
									<td class="capitalize">
										<div class="tag" :class="{ 'tag-warning' : item.post.status  === 'draft' }">{{ item.post.status }}</div>
									</td>
									<!-- Published at -->
									<td>
										<span v-if="!item.post['published_at']">Not published</span>
										<span v-else>{{ item.post['published_at'] | moment("dddd, MMMM Do YYYY") }}</span>
									</td>
									<!-- Actions -->
									<td class="archive-table-actions">
										<i class="far fa-ellipsis-h"></i>
									</td>
								</tr>
							</tbody>
						</table>

					</div><!-- /Table Scroll -->
<!--						<div class="editor-slide" :class="{ 'editor-slide-active' : activeTab === 0}" v-if="fieldLayout.length && activeTab === 0" :key="1">-->
<!--					</transition>-->
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
<!--					<pre>{{ orderBy }}</pre>-->
					<pre>	{{ checked }}</pre>
					{{ checked }}
					{{ bulkType }}
				</div>
			</div>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Tabs from "../../components/misc/Tabs";

export default {
	name: "Pages",
	title: "Archive",
	components: {
		Breadcrumbs,
		Tabs
	},
	data: () => ({
		doingAxios: false,
		activeTab: 1,
		posts: [],
		orderBy: {
			title: "asc",
			user_id: "asc",
			status: "asc",
		},
		savingBulk: false,
		bulkType: "",
		checked: [],
	}),
	mounted() {
		this.getPosts();
	},
	watch: {
		'$route.params.resource': function() {
			this.getResourceByName();
		},
	},
	methods: {
		getPosts(order = "") {
			// TODO: Add filter to go
			console.log(order)
			this.axios.get("/posts", {
				// params: {
				// 	//"order": order,
				// }
			})
			.then(res => {
				this.posts = {};
				this.posts = res.data.data
			})
			.catch(err => {
				console.log(err)
				this.$noty.error("Error occurred, please refresh the page.")
			});
		},
		changeOrderBy(column) {
			if (this.orderBy[column] === "desc" || this.orderBy[column] == "") {
				this.$set(this.orderBy, column, 'asc');
			} else {
				this.$set(this.orderBy, column, 'desc');
			}
			this.getPosts(this.getOrderBySting(column));
		},
		getOrderBySting(column) {
			return column + "," + this.orderBy[column]
		},
		getPostsById(id) {
			return this.posts.find(p => p.post.id === id)
		},
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
				this.moveToDrafts()
					.then(() => {
						this.$noty.success("Posts updated successfully.")
						this.getPosts()
					})
					.catch(err => {
						console.log(err);
						this.$noty.error("Error occurred, please refresh the page.")
					})
					.finally(() => {
						this.savingBulk = false;
					});
			} else if (this.bulkType === "delete") {
				this.deletePost()
					.then(() => {
						this.$noty.success("Posts deleted successfully.")
						this.getPosts()
					})
					.catch(err => {
						console.log(err);
						this.$noty.error("Error occurred, please refresh the page.")
					})
					.finally(() => {
						this.savingBulk = false;
					});
			}
		},
		moveToDrafts() {
			return new Promise((resolve, reject) => {
				this.checked.forEach(id => {
					const post =  this.getPostsById(id).post
					post.status = "draft"
					this.axios.put("/posts/" + id, post)
						.catch((err) => {
							reject(err);
						})
				});
				resolve();
			});
		},
		deletePost() {
			return new Promise((resolve, reject) => {
				this.checked.forEach(id => {
					this.axios.delete("/posts/" + id)
						.catch((err) => {
							reject(err);
						})
				});
				resolve();
			});
		}
	},
	computed: {
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


			// Order
			// =========================================================================

			&-order {
				cursor: pointer;
				user-select: none;

				i {
					color: rgba($secondary, 0.7);
					font-size: 12px;
					margin-left: 4px;
					transition: 180ms ease;
					will-change: transform;

					&.active {
						transform: rotate(180deg);
					}
				}
			}

			// Title
			// =========================================================================

			&-title {

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

			// Actions
			// =========================================================================

			&-actions {

				i {
					font-size: 16px;
				}
			}
		}
	}

</style>