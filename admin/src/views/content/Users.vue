<!-- =====================
	Users
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Users</h1>
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
								<div class="btn btn-icon btn-orange btn-text-mob" @click="showCreateModal = true">
									<i class="fal fa-plus"></i>
									<span>New User</span>
								</div>
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
						<template slot="item">Banned</template>
						<template slot="item">Contributor</template>
						<template slot="item">Author</template>
						<template slot="item">Editor</template>
						<template slot="item">Admin</template>
					</Tabs>
					<!-- Spinner -->
					<div v-if="doingAxios" class="media-spinner spinner-container">
						<div class="spinner spinner-large spinner-grey"></div>
					</div>
					<!-- =====================
						Users
						===================== -->
					<div v-else>
						<transition name="trans-fade" mode="out-in">
							<div class="table-wrapper" v-if="users.length">
								<div class="table-scroll table-with-hover">
									<table class="table users-table">
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
											<tr class="trans-fade-in-anim-slow" v-for="(user, userIndex) in users" :key="user.uuid" >
												<!-- Checkbox -->
												<td class="table-checkbox">
													<div class="form-checkbox form-checkbox-dark">
														<input type="checkbox" :id="user.uuid" :value="user.id" v-model="checked"/>
														<label :for="user.uuid">
															<i class="fal fa-check"></i>
														</label>
													</div>
												</td>
												<!-- Name, Email & Avatar -->
												<td>
													<router-link  class="users-table-info" :to="{ name: 'edit-user', params: { id: user.id }}">
														<img v-if="getProfilePicture(user['profile_picture_id'])" class="avatar" :src="getSite.url + getProfilePicture(user['profile_picture_id'])">
														<span v-else class="avatar" v-html="getInitials(user)"></span>
														<div class="users-table-info-name">
															<h4>{{ user['first_name'] }} {{ user['last_name'] }}</h4>
															<p>{{ user['email'] }}</p>
														</div>
													</router-link>
												</td>
												<!-- Role -->
												<td class="users-table-role">
													<div class="badge capitalize" :class="{
														'badge-orange' : user.role.name  === 'Banned',
														'badge-green' : user.role.name  === 'Administrator',
														'badge-blue' : user.role.name  === 'Contributor' || user.role.name  === 'Editor' || user.role.name  === 'Author',
													}">{{ user.role.name }}</div>
												</td>
												<!-- Created at -->
												<td>
													<span>{{ user['created_at'] | moment("dddd, MMMM Do YYYY") }}</span>
												</td>
												<td class="table-actions">
													<Popover :triangle="false"
															@update="updateActions($event, user.uuid)"
															:classes="(userIndex + 1) > (users.length - 4) ? 'popover-table popover-table-top' : 'popover-table popover-table-bottom'"
															:item-key="user.uuid"
															:active="activeAction">
														<template slot="items">
															<router-link class="popover-item popover-item-icon" :to="{ name: 'edit-user', params: { id: user.id }}">
																<i class="feather feather-edit"></i>
																<span>Edit</span>
															</router-link>
															<div class="popover-line"></div>
															<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(user);">
																<i class="feather feather-trash-2"></i>
																<span>Delete</span>
															</div>
														</template>
														<template slot="button">
															<i class="icon icon-square far fa-ellipsis-h" :class="{'icon-square-active' : activeAction === user.uuid}"></i>
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
									<h3>No {{ activeTabName === "all" ? "" : activeTabNamem }} users available. </h3>
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
		</div><!-- /Container -->
		<!-- =====================
			Delete Modal
			===================== -->
		<Modal :show.sync="showDeleteModal" class="modal-with-icon modal-with-warning users-delete">
			<template slot="button">
				<button class="btn" :class="{ 'btn-loading' : isDeleting }" @click="deleteUser">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="selectedUser">Are you sure want to delete {{ selectedUser['first_name'] }}?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} users?</p>
			</template>
		</Modal>
		<!-- =====================
			Create Modal
			===================== -->
		<CreateUser :show.sync="showCreateModal" @update="getUsers"></CreateUser>
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
import CreateUser from "@/components/modals/CreateUser";
import {userMixin} from "@/util/users";

export default {
	name: "Users",
	title: "Users",
	mixins: [userMixin],
	components: {
		CreateUser,
		Alert,
		Breadcrumbs,
		Modal,
		Popover,
		Pagination,
		Tabs,
	},
	data: () => ({
		doingAxios: true,
		users: [],
		selectedUser: false,
		roles: [],
		media: [],
		errors: [],
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
		this.getUsers();
		this.getMedia();
	},
	methods: {
		/*
		 * getUsers()
		 * Obtain the users & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getUsers() {
			this.axios.get(`users?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.users = [];
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					const users = res.data.data
					if (users.length) {
						this.users = this.removeSelf(users);
					}
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * deleteUser()
		 * Marking the interface as deleting (for the button) and detected if the
		 * user has been selected by the actions or bulk, push all of the
		 * promises to an array and send all requests.
		 */
		deleteUser() {
			this.isDeleting = false;

			let toDelete = this.selectedUser ? [this.selectedUser.id] : this.checked;

			const promises = [];
			toDelete.forEach(id => {
				promises.push(this.deleteUserAxios(id));
			});

			// Send all requests
			Promise.all(promises)
				.then(() => {
					const successMsg = toDelete.length === 1 ? "User deleted successfully" : "Users deleted successfully."
					this.$noty.success(successMsg);
					this.getUsers();
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						this.$noty.error(err.response.data.message);
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.activeAction = "";
					this.checked = [];
					this.checkedAll = false;
					this.showDeleteModal = false;
					this.bulkType = "";
					this.isDeleting = false;
					this.selectedUser = false;
				});
		},
		/*
		 * async deleteUserAxios()
		 */
		async deleteUserAxios(id) {
			return await this.axios.delete("/users/" + id);
		},
		/*
		 * handleDelete()
		 * Changes the selected user to the given input,
		 * & show's the delete user modal.
		 */
		handleDelete(user) {
			this.selectedUser = user;
			this.showDeleteModal = true;
		},
		/*
		 * removeSelf()
		 * Remove the current logged in user from the list.
		 */
		removeSelf(users) {
			return users.filter(u => {
				return u.id !== this.getUserInfo.id;
			})
		},
		/*
		 * changeOrderBy()
		 * Update the order by object when clicked, obtain users.
		 */
		changeOrderBy(column) {
			this.activeOrder = column;
			if (this.orderBy[column] === "desc" || this.orderBy[column] === "") {
				this.$set(this.orderBy, column, 'asc');
			} else {
				this.$set(this.orderBy, column, 'desc');
			}
			this.order = column + "," + this.orderBy[column];
			this.getUsers();
		},
		/*
		 * filterTabs()
		 * Update the filter by string when tabs are clicked, obtain users.
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
					this.activeTabName = "banned";
					filter = '{"roles.name":[{"operator":"=", "value": "Banned"}]}';
					break;
				}
				case 3: {
					this.activeTabName = "contributor";
					filter = '{"roles.name":[{"operator":"=", "value": "Contributor"}]}';
					break;
				}
				case 4: {
					this.activeTabName = "author";
					filter = '{"roles.name":[{"operator":"=", "value": "Author"}]}';
					break;
				}
				case 5: {
					this.activeTabName = "editor";
					filter = '{"roles.name":[{"operator":"=", "value": "Editor"}]}';
					break;
				}
				case 6: {
					this.activeTabName = "administrator";
					filter = '{"roles.name":[{"operator":"=", "value": "Administrator"}]}';
					break;
				}
			}
			this.filter = filter;
			this.getUsers();
		},
		/*
		 * setPagination()
		 * Update the pagination string when clicked, obtain users.
		 */
		setPagination(query) {
			this.activeAction = "";
			this.pagination = query;
			this.getUsers();
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
		 * getInitials()
		 */
		getInitials(user) {
			return user['first_name'].charAt(0) + user['last_name'].charAt(0).toUpperCase();
		},
		/*
		 * getMedia()
		 * Return media for filtering profile picture.
		 */
		getMedia() {
			this.axios.get("/media")
				.then(res => {
					this.media = res.data.data;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getProfilePicture()
		 */
		getProfilePicture(id) {
			console.log(this.media);
			if (this.media.length) {
				const picture = this.media.find(m => m.id === id);
				return picture ? picture.url : false;
			}
		}
	},
	computed: {
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return this.checked.length === this.users.length;
			},
			set(value) {
				if (value) {
					this.checked = this.users.map(m => {
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
<style lang="scss">

.users {

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

		// Name, Email & Avatar
		// =========================================================================

		&-info {
			display: flex;

			.avatar {
				margin-right: 16px;
			}

			h4 {
				margin-bottom: -2px;
			}

			p {
				font-size: 14px;
			}
		}
	}
}

</style>