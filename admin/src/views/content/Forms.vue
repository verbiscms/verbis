<!-- =====================
	Forms
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Form Submissions</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<!-- Spinner -->
					<div v-if="doingAxios" class="media-spinner spinner-container">
						<div class="spinner spinner-large spinner-grey"></div>
					</div>
					<!-- =====================
						Submissions
						===================== -->
					<div v-else>
						<transition name="trans-fade" mode="out-in">
							<div class="table-wrapper" v-if="forms.length">
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
												<span>Name</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['from_path'] !== 'asc' }"></i>
											</th>
											<th class="table-order" @click="changeOrderBy('to_path')" :class="{ 'active' : activeOrder === 'to_path' }">
												<span>Submissions</span>
												<i class="fas fa-caret-down" :class="{ 'active' : orderBy['to_path'] !== 'asc' }"></i>
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
										<tr class="trans-fade-in-anim-slow" v-for="(form, formIndex) in forms" :key="formIndex">
											<!-- Checkbox -->
											<td class="table-checkbox">
												<div class="form-checkbox form-checkbox-dark">
													<input type="checkbox" :id="form.uuid" :value="form.id" v-model="checked"/>
													<label :for="form.uuid">
														<i class="fal fa-check"></i>
													</label>
												</div>
											</td>
											<!-- Name -->
											<td>
												<router-link :to="{ name: 'forms-single', params: { id: form.id }}">
													<span>{{ form['name'] }}</span>
												</router-link>
											</td>
											<!-- Name -->
											<td>
												<span>{{ getSubmissionCount(form) }}</span>
											</td>
											<!-- Updated At -->
											<td>
												<span>{{ form['updated_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<!-- Created At -->
											<td>
												<span>{{ form['created_at'] | moment("dddd, MMMM Do YYYY") }}</span>
											</td>
											<td class="table-actions">
												<Popover :triangle="false" @update="updateActions($event, form.id.toString())" :classes="(formIndex + 1) > (form.length - 4) ? 'popover-table popover-table-top' : 'popover-table popover-table-bottom'" :item-key="form.id.toString()" :active="activeAction">
													<template slot="items">
														<div class="popover-item popover-item-icon" @click="editHandler(form)">
															<i class="feather feather-edit"></i>
															<span>Edit</span>
														</div>
														<div class="popover-line"></div>
														<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(form);">
															<i class="feather feather-trash-2"></i>
															<span>Delete</span>
														</div>
													</template>
													<template slot="button">
														<i class="icon icon-square far fa-ellipsis-h" :class="{'icon-square-active' : activeAction === form.id}"></i>
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
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import Alert from "@/components/misc/Alert";
import Popover from "@/components/misc/Popover";

export default {
	name: "Forms",
	components: {
		Alert,
		Breadcrumbs,
		Popover,
	},
	data: () => ({
		doingAxios: true,
		forms: [],
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
			this.axios.get(`/forms?order_by=${this.order[0]}&order_direction=${this.order[1]}&filter=${this.filter}&limit=all`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.forms = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					this.forms = res.data.data
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
			this.order = [column,this.orderBy[column]];
			this.getRedirects();
		},
		/*
		 * getSubmissionCount()
		 * Gets the total amount of submissions per form.
		 */
		getSubmissionCount(form) {
			if (!form['submissions'] || !form['submissions'].length) {
				return 0;
			}
			return form['submissions'].length
		},
	},
	computed: {
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return this.checked.length === this.forms.length;
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