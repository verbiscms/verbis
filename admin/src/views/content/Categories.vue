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
								<router-link :to="{ name: 'categories-single', params: { id: 'new' }}"  class="btn btn-icon btn-orange btn-text-mob">
									<i class="fal fa-plus"></i>
									<span>New user</span>
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
													<input type="checkbox" id="users-check-all" v-model="checkedAll"/>
													<label for="users-check-all">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</th>
											<th class="table-order" @click="changeOrderBy('first_name')" :class="{ 'active' : activeOrder === 'first_name' }">
												<span>Name</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['first_name'] !== 'asc' }"></i>
											</th>
											<!-- @click="changeOrderBy('role.name')" :class="{ 'active' : activeOrder === 'role.name' }" -->
											<th class="table-order">
												<span>Role</span>
												<!--													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['role.name'] !== 'asc' }"></i>-->
											</th>
											<th class="table-order" @click="changeOrderBy('created_at')" :class="{ 'active' : activeOrder === 'created_at' }">
												<span>Created at</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['created_at'] !== 'asc' }"></i>
											</th>
											<th></th>
										</tr>
										</thead>
										<tbody>
										<tr class="trans-fade-in-anim-slow" v-for="(category) in categories" :key="category.uuid" >
											<!-- Checkbox -->
											<td class="table-checkbox">
												<div class="form-checkbox form-checkbox-dark">
													<input type="checkbox" :id="category.uuid" :value="category.id" v-model="checked"/>
													<label :for="category.uuid">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</td>
											<!-- Name, Email & Avatar -->
											<td>

											</td>
											<!-- Role -->
											<td>
											</td>
											<!-- Created at -->
											<td>

											</td>
											<td class="table-actions">

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
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Tabs from "@/components/misc/Tabs";
import Alert from "@/components/misc/Alert";

export default {
	name: "Categories",
	components: {
		Alert,
		Tabs,
		Breadcrumbs
	},
	data: () => ({
		doingAxios: true,
		categories: [],
		errors: [],
		selectecCategory: false,
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
		showCreateModal: false,
		selectedDeleteId: null,
		isDeleting: false,
		isDoingBulk: false,
	}),
	mounted() {
		this.getCategories();
	},
	methods: {
		/*
		 * getCategories()
		 * Obtain the categories & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getCategories() {
			this.axios.get(`categories?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.categories = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					console.log(res);
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
			this.order = column + "," + this.orderBy[column];
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