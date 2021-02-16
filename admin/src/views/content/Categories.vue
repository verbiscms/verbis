<!-- =====================
	Categories
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Categories</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<div class="form-select-cont form-input header-hide-mob">
									<select class="form-select" v-model="bulkType">
										<option value="" disabled selected>Bulk actions</option>
										<option value="delete">Delete permanently</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white header-hide-mob" :class="{ 'btn-loading' : isDoingBulk }" @click.prevent="doBulkAction">Apply</button>
								<router-link :to="{ name: 'categories-single', params: { id: 'new' }}"  class="btn btn-orange btn-fixed-height btn-flex">
									New Category
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
						<template slot="item">Pages</template>
						<template slot="item" v-for="resource in getTheme['resources']">{{ resource['friendly_name'] }}</template>
					</Tabs>
					<!-- Spinner -->
					<div v-if="doingAxios" class="media-spinner spinner-container">
						<div class="spinner spinner-large spinner-grey"></div>
					</div>
					<!-- =====================
						Categories
						===================== -->
					<div v-else>
						<transition name="trans-fade" mode="out-in">
							<div class="table-wrapper" v-if="categories.length">
								<div class="table-scroll table-with-hover">
									<table class="table categories-table">
										<thead>
										<tr>
											<th class="table-header-checkbox">
												<div class="form-checkbox form-checkbox-dark">
													<input type="checkbox" id="categories-check-all" v-model="checkedAll"/>
													<label for="categories-check-all">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</th>
											<th class="table-order" @click="changeOrderBy('name')" :class="{ 'active' : activeOrder === 'name' }">
												<span>Name</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['first_name'] !== 'asc' }"></i>
											</th>
											<th class="table-order" @click="changeOrderBy('resource')" :class="{ 'active' : activeOrder === 'resource' }">
												<span>Resource</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['resource'] !== 'asc' }"></i>
											</th>
											<th>
												<span>Parent</span>
											</th>
											<th class="table-order" @click="changeOrderBy('updated_at')" :class="{ 'active' : activeOrder === 'updated_at' }">
												<span>Updated at</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['updated_at'] !== 'asc' }"></i>
											</th>
											<th class="table-order" @click="changeOrderBy('created_at')" :class="{ 'active' : activeOrder === 'created_at' }">
												<span>Created at</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['created_at'] !== 'asc' }"></i>
											</th>
											<th></th>
										</tr>
										</thead>
										<tbody>
										<tr class="trans-fade-in-anim-slow" v-for="(category, categoryIndex) in categories" :key="category.uuid" >
											<!-- Checkbox -->
											<td class="table-checkbox">
												<div class="form-checkbox form-checkbox-dark">
													<input type="checkbox" :id="category.uuid" :value="category.id" v-model="checked"/>
													<label :for="category.uuid">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</td>
											<!-- Name -->
											<td>
												<router-link :to="{ name: 'categories-single', params: { id: category.id }}">
													{{ category['name'] }}
												</router-link>
											</td>
											<!-- Resource -->
											<td>
												<div class="badge badge-green">{{ capitalize(category.resource) }}</div>
											</td>
											<!-- Parent -->
											<td>
												<p v-if="filterCategoryById(category['parent_id'])">{{ filterCategoryById(category['parent_id'])['name'] }}</p>
												<p v-else>No parent</p>
											</td>
											<!-- Updated At -->
											<td>
												<span>{{ category['updated_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<!-- Created At -->
											<td>
												<span>{{ category['created_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<td class="table-actions">
												<Popover :triangle="false"
														@update="updateActions($event, category.uuid)"
														:classes="(categoryIndex + 1) > (categories.length - 4) ? 'popover-table popover-table-top' : 'popover-table popover-table-bottom'"
														:item-key="category.uuid"
														:active="activeAction">
													<template slot="items">
														<router-link class="popover-item popover-item-icon" :to="{ name: 'categories-single', params: { id: category.id }}">
															<i class="feather feather-edit"></i>
															<span>Edit</span>
														</router-link>
														<div class="popover-line"></div>
														<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(category);">
															<i class="feather feather-trash-2"></i>
															<span>Delete</span>
														</div>
													</template>
													<template slot="button">
														<i class="icon icon-square far fa-ellipsis-h" :class="{'icon-square-active' : activeAction === category.uuid}"></i>
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
									<h3>No {{ activeTabName === "all" ? "" : activeTabName }} categories available. </h3>
									<p>To create a new one, click the button above.</p>
								</slot>
							</Alert>
						</transition>
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<transition name="trans-fade-in-anim">
				<div class="row" v-if="!doingAxios && paginationObj">
					<div class="col-12">
						<Pagination :pagination="paginationObj" @update="setPagination"></Pagination>
					</div><!-- /Col -->
				</div><!-- /Row -->
			</transition>
		</div>
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning">
			<template slot="button">
				<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="deleteCategory(false);">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="selectedCategory">Are you sure want to delete this category?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} categories?</p>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Tabs from "@/components/misc/Tabs";
import Alert from "@/components/misc/Alert";
import Pagination from "@/components/misc/Pagination";
import Modal from "@/components/modals/General";
import Popover from "@/components/misc/Popover";

export default {
	name: "Categories",
	components: {
		Modal,
		Pagination,
		Alert,
		Tabs,
		Breadcrumbs,
		Popover,
	},
	data: () => ({
		doingAxios: true,
		categories: [],
		selectedCategory: false,
		errors: [],
		paginationObj: {},
		activeTab: 1,
		activeTabName: "all",
		order: ["", ""],
		orderBy: {
			name: "asc",
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
		showCreateModal: false,
		selectedDeleteId: null,
		isDeleting: false,
		isDoingBulk: false,
	}),
	mounted() {
		this.getCategories();
		this.getMessage();
	},
	methods: {
		/*
		 * getMessage()
		 * Determine if the category has been deleted.
		 */
		getMessage() {
			if (this.$route.query.delete) {
				this.$noty.success("Successfully deleted category.")
				let query = Object.assign({}, this.$route.query);
				delete query.delete;
				this.$router.replace({ query });
			}
		},
		/*
		 * getCategories()
		 * Obtain the categories & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getCategories() {
			this.axios.get(`/categories?order_by=${this.order[0]}&order_direction=${this.order[1]}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.categories = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					this.categories = res.data.data
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * deleteCategory()
		 */
		deleteCategory() {
			this.isDeleting = true;

			const toDelete = this.selectedCategory ? [this.selectedCategory.id] : this.checked;

			const promises = [];
			toDelete.forEach(id => {
				console.log(id);
				promises.push(this.deleteCategoryAxios(id));
			});

			// Send all requests
			Promise.all(promises)
				.then(() => {
					const successMsg = toDelete.length === 1 ? `Category deleted successfully.` : `Categories deleted successfully.`
					this.$noty.success(successMsg);
					this.getCategories();
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
					this.selectedCategory = false;
				});
		},
		/*
		 * async deleteCategoryAxios()
		 */
		async deleteCategoryAxios(id) {
			return await this.axios.delete("/categories/" + id);
		},
		/*
		 * handleDelete()
		 * Changes the selected category to the given input,
		 * & show's the delete category modal.
		 */
		handleDelete(category) {
			this.selectedCategory = category;
			this.showDeleteModal = true;
		},
		/*
		 * changeOrderBy()
		 * Update the order by object when clicked, obtain categories.
		 */
		changeOrderBy(column) {
			this.activeOrder = column;
			if (this.orderBy[column] === "desc" || this.orderBy[column] === "") {
				this.$set(this.orderBy, column, 'asc');
			} else {
				this.$set(this.orderBy, column, 'desc');
			}
			this.order = [column,this.orderBy[column]];
			this.getCategories();
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain categories.
		 */
		filterTabs(tab) {
			this.pagination = "page=1";
			this.activeTab = tab;
			let filter = "";
			switch (tab) {
				case 1: {
					this.activeTabName = "all";
					break;
				}
				case 2: {
					this.activeTabName = 'page';
					filter = `{"resource":[{"operator":"=", "value": "pages"}]}`;
					break;
				}
				default: {
					const resources = this.getTheme['resources'];
					const resource = resources[Object.keys(resources)[tab - 3]];
					if (resource) {
						const key = Object.keys(resources).find(key => resources[key] === resource)
						this.activeTabName = resource['singular_name'];
						filter = `{"resource":[{"operator":"=", "value": "${key}"}]}`;
					}
					break;
				}

			}
			this.filter = filter;
			this.getCategories();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain categories.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getCategories();
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
				}, this.timeoutDelay)
				return
			}

			// Delete
			if (this.bulkType === "delete") {
				setTimeout(() => {
					this.isDoingBulk = false;
				}, this.timeoutDelay)
				this.showDeleteModal = true;
			}
		},
		/*
		 * updateActions()
		 *  Update the action uuid for clearing the popover.
		 */
		updateActions(e, uuid) {
			this.activeAction = e ? uuid : "";
		},
		/*
		 * filterCategoryById()
		 * Filter all of the categories and get by ID.
		 */
		filterCategoryById(id) {
			let category = false;
			this.categories.forEach(c => {
				if (id === c.id) {
					category = c;
				}
			});
			return category;
		},
		/*
 		 * capitalize()
		 * Capitalize the first letter of the resource.
		 */
		capitalize(str) {
			return str.replace(/(?:^|\s|["'([{])+\S/g, match => match.toUpperCase());
		}
	},
	computed: {
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return this.checked.length === this.categories.length;
			},
			set(value) {
				if (value) {
					this.checked = this.categories.map(m => {
						return m.id;
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

	// Dummy
	// =========================================================================

</style>