<!-- =====================
	Settings - Redirects
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Redirects</h1>
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
								<button class="btn btn-orange btn-fixed-height btn-flex" @click.prevent="showRedirectModal = true">
									New Redirect
								</button>
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
						<template slot="item">300's</template>
						<template slot="item">301's</template>
						<template slot="item">302's</template>
						<template slot="item">303's</template>
						<template slot="item">304's</template>
						<template slot="item">305's</template>
						<template slot="item">306's</template>
						<template slot="item">307's</template>
						<template slot="item">308's</template>
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
							<div class="table-wrapper" v-if="redirects.length">
								<div class="table-scroll table-with-hover">
									<table class="table redirects-table">
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
											<th class="table-order" @click="changeOrderBy('from_path')" :class="{ 'active' : activeOrder === 'from_path' }">
												<span>From</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['from_path'] !== 'asc' }"></i>
											</th>
											<th class="table-order" @click="changeOrderBy('to_path')" :class="{ 'active' : activeOrder === 'to_path' }">
												<span>To</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['to_path'] !== 'asc' }"></i>
											</th>
											<th class="table-order" @click="changeOrderBy('code')" :class="{ 'active' : activeOrder === 'code' }">
												<span>Code</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['code'] !== 'asc' }"></i>
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
										<tr class="trans-fade-in-anim-slow" v-for="(redirect, redirectsIndex) in redirects" :key="redirectsIndex" >
											<!-- Checkbox -->
											<td class="table-checkbox">
												<div class="form-checkbox form-checkbox-dark">
													<input type="checkbox" :id="redirect.uuid" :value="redirect.id" v-model="checked"/>
													<label :for="redirect.uuid">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</td>
											<!-- From -->
											<td>
												<span>{{ redirect['from_path'] }}</span>
											</td>
											<!-- To -->
											<td>
												<span>{{ redirect['to_path'] }}</span>
											</td>
											<!-- Code -->
											<td>
												<span>{{ redirect['code'] }}</span>
											</td>
											<!-- Updated At -->
											<td>
												<span>{{ redirect['updated_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<!-- Created At -->
											<td>
												<span>{{ redirect['created_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<td class="table-actions">
												<Popover :triangle="false" @update="updateActions($event, redirect.id.toString())" :classes="(redirectsIndex + 1) > (redirects.length - 4) ? 'popover-table popover-table-top' : 'popover-table popover-table-bottom'" :item-key="redirect.id.toString()" :active="activeAction">
													<template slot="items">
														<div class="popover-item popover-item-icon" @click="editHandler(redirect)">
															<i class="feather feather-edit"></i>
															<span>Edit</span>
														</div>
														<div class="popover-line"></div>
														<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(redirect);">
															<i class="feather feather-trash-2"></i>
															<span>Delete</span>
														</div>
													</template>
													<template slot="button">
														<i class="icon icon-square far fa-ellipsis-h" :class="{'icon-square-active' : activeAction === redirect.id}"></i>
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
									<h3>No {{ activeTabName === "all" ? "" : activeTabName }} redirects available. </h3>
									<p>To create a new one, click the button above.</p>
								</slot>
							</Alert>
						</transition>
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Redirect Modal
			===================== -->
		<Redirect :show.sync="showRedirectModal" :redirect-update="selectedRedirect" @update="updateCreateHandler"></Redirect>
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning">
			<template slot="button">
				<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="deleteRedirect(false)" @update="getRedirects">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="selectedRedirect">Are you sure want to delete this redirect?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} redirects?</p>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Alert from "@/components/misc/Alert";
import Redirect from "@/components/modals/Redirect";
import Tabs from "@/components/misc/Tabs";
import Modal from "@/components/modals/General";
import Popover from "@/components/misc/Popover";

export default {
	name: "Redirects",
	components: {
		Alert,
		Redirect,
		Breadcrumbs,
		Tabs,
		Modal,
		Popover,
	},
	data: () => ({
		doingAxios: true,
		redirects: [],
		selectedRedirect: false,
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
		showRedirectModal: false,
		selectedRedirectKey: false,
	}),
	mounted() {
		this.getRedirects();
	},
	methods: {
		/*
		 * getRedirects()
		 * Obtain the redirects & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getRedirects() {
			this.axios.get(`/redirects?order_by=${this.order[0]}&order_direction=${this.order[1]}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.redirects = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					this.redirects = res.data.data
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * deleteRedirect()
		 */
		deleteRedirect() {
			this.isDeleting = true;

			const toDelete = this.selectedRedirect ? [this.selectedRedirect.id] : this.checked;

			const promises = [];
			toDelete.forEach(id => {
				promises.push(this.deleteRedirectAxios(id));
			});

			// Send all requests
			Promise.all(promises)
				.then(() => {
					const successMsg = toDelete.length === 1 ? `Redirect deleted successfully.` : `Redirects deleted successfully.`
					this.$noty.success(successMsg);
					this.getRedirects();
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
		 * async deleteRedirectAxios()
		 */
		async deleteRedirectAxios(id) {
			return await this.axios.delete("/redirects/" + id);
		},
		/*
		 * handleDelete()
		 * Changes the selected redirect to the given input,
		 * & show's the delete redirect modal.
		 */
		handleDelete(redirect) {
			this.selectedRedirect = redirect;
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
			this.getRedirects();
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain redirects.
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
					this.activeTabName = '300';
					filter = `{"code":[{"operator":"=", "value": "300"}]}`;
					break;
				}
				case 3: {
					this.activeTabName = '301';
					filter = `{"code":[{"operator":"=", "value": "301"}]}`;
					break;
				}
				case 4: {
					this.activeTabName = '302';
					filter = `{"code":[{"operator":"=", "value": "302"}]}`;
					break;
				}
				case 5: {
					this.activeTabName = '303';
					filter = `{"code":[{"operator":"=", "value": "303"}]}`;
					break;
				}
				case 7: {
					this.activeTabName = '304';
					filter = `{"code":[{"operator":"=", "value": "304"}]}`;
					break;
				}
				case 8: {
					this.activeTabName = '305';
					filter = `{"code":[{"operator":"=", "value": "305"}]}`;
					break;
				}
				case 9: {
					this.activeTabName = '305';
					filter = `{"code":[{"operator":"=", "value": "305"}]}`;
					break;
				}
				case 10: {
					this.activeTabName = '306';
					filter = `{"code":[{"operator":"=", "value": "306"}]}`;
					break;
				}
				case 11: {
					this.activeTabName = '307';
					filter = `{"code":[{"operator":"=", "value": "307"}]}`;
					break;
				}
				case 12: {
					this.activeTabName = '308';
					filter = `{"code":[{"operator":"=", "value": "308"}]}`;
					break;
				}
			}
			this.filter = filter;
			this.getRedirects();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain categories.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getRedirects();
		},
		/*
		 * doBulkAction()
		 * When bulk action is clicked, this function call delete.
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
		 * editHandler()
		 * Fires when the redirect edit button is clicked.
		 */
		editHandler(redirect) {
			this.selectedRedirect = redirect;
			this.showRedirectModal = true;
			this.activeAction = "";
		},
		/*
		 * updateCreateHandler()
		 * Fires when a new redirect is created or updated.
		 */
		updateCreateHandler() {
			this.getRedirects();
			this.showRedirectModal = false;
			this.selectedRedirect = false;
		},
		/*
		 * updateActions()
		 *  Update the action uuid for clearing the popover.
		 */
		updateActions(e, id) {
			this.activeAction = e ? id.toString() : "";
		},
	},
	computed: {
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return this.checked.length === this.redirects.length;
			},
			set(value) {
				if (value) {
					this.checked = this.redirects.map(m => {
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

.seo {

	// Sitemap
	// =========================================================================

	&-sitemap-btn {
		top: 0;
		right: 1.2rem;
		position: absolute;
		display: flex;
		justify-content: flex-end;
		z-index: 99;

		.btn {
			background-color: $white;
		}
	}

	// Redirects
	// =========================================================================

	&-redirects {


		//.table {
		//	border-top: 1px solid $grey-light;
		//}

		.card-controls {

		}

		.feather-trash-2 {
			color: $orange;
		}
		//&-header {
		//	display: flex;
		//	align-items: center;
		//	justify-content: space-between;
		//	margin-bottom: 1rem;
		//
		//	p {
		//		margin-bottom: 0;
		//	}
		//}
	}
}

</style>