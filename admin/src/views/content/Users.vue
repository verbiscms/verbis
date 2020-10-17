<!-- =====================
	Users
	===================== -->
<template>
	<section>
		<div class="auth-container">
			{{ roles }}
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
								<div class="form-select-cont form-input">
									<select class="form-select">
										<option value="" disabled selected>Bulk actions</option>
										<option value="restore">Restore</option>
										<option value="delete">Delete permanently</option>
									</select>
								</div>
								<button class="btn btn-fixed-height btn-margin btn-white" :class="{ 'btn-loading' : savingBulk }" @click.prevent="doBulkAction">Apply</button>
								<div class="btn btn-icon btn-orange">
									<i class="fal fa-plus"></i>
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
					<!-- =====================
						Users
						===================== -->
					<div v-if="!doingAxios">
						<transition name="trans-fade-quick" mode="out-in">
							<div class="table-wrapper" v-if="users.length">
								<div class="table-scroll table-with-hover">
									<table class="table users-table">
										<thead>
											<tr>
												<th>
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
												<th class="table-order" @click="changeOrderBy('role.name')" :class="{ 'active' : activeOrder === 'role.name' }">
													<span>Role</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['role.name'] !== 'asc' }"></i>
												</th>
												<th class="table-order" @click="changeOrderBy('created_at')" :class="{ 'active' : activeOrder === 'created_at' }">
													<span>Created at</span>
													<i class="fas fa-caret-down" :class="{ 'active' : orderBy['created_at'] !== 'asc' }"></i>
												</th>
												<th></th>
											</tr>
										</thead>
										<tbody>
											<tr v-for="(user, userIndex) in users" :key="user.uuid">
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
												<td class="users-table-info">
													<span class="avatar" v-html="getInitials(user)"></span>
													<div class="users-table-info-name">
														<h4>{{ user['first_name'] }} {{ user['last_name'] }}</h4>
														<p>{{ user['email'] }}</p>
													</div>
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
															<div class="popover-item popover-item-icon popover-item-border popover-item-orange" @click="handleDelete(item.post.id);">
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
<!--									<h3>No {{ resource['friendly_name'].toLowerCase() }} available. </h3>-->
									<p>To create a new one, click the plus sign above.</p>
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
		<Modal :show.sync="showDeleteModal">
			<template slot="button">
				<button class="btn" @click="deletePost(false); savingBulk = true;">Delete</button>
			</template>
			<template slot="text">
				<h2>Are you sure?</h2>
				<p v-if="checked.length === 1">Are you sure want to delete this user }}?</p>
				<p v-else>Are you sure want to delete {{ checked.length }} users?</p>
			</template>
		</Modal>
		<!-- =====================
			Create Modal
			===================== -->
		<Modal :show.sync="showCreateModal" class="users-modal-create">
			<template slot="button">
				<button class="btn">Create</button>
			</template>
			<template slot="text">
				<div class="row">
					<div class="col-12">
						<h2>New User</h2>
					</div><!-- /Col -->
					<div class="col-12 col-desk-6">
						<div class="form-group">
							<label class="form-label" for="user-first-name">First name:</label>
							<input class="form-input form-input-white" id="user-first-name" type="text" v-model="newUser['first_name']">
						</div>
						<div class="form-group">
							<label class="form-label" for="user-email">Email:</label>
							<input class="form-input form-input-white" id="user-email" type="text" v-model="newUser['email']">
						</div>
						<div class="form-group">
							<label class="form-label" for="user-password">Password:</label>
							<input class="form-input form-input-white" id="user-password" type="text">
						</div>
					</div><!-- /Col -->
					<div class="col-12 col-desk-6">
						<div class="form-group">
							<label class="form-label" for="user-last-name">Last name:</label>
							<input class="form-input form-input-white" id="user-last-name" type="text" v-model="newUser['last_name']">
						</div>
						<div class="form-group">
							<label class="form-label" for="user-role">Role:</label>
							<input class="form-input form-input-white" id="user-role" type="text">
						</div>
						<div class="form-group">
							<label class="form-label" for="user-role">Confirm Password:</label>
							<input class="form-input form-input-white" id="user-confirm-password" type="text">
						</div>
					</div><!-- /Col -->
				</div><!-- /Row -->
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
	name: "Users",
	title: "Users",
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
		users: [],
		roles: [],
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
		showDeleteModal: false,
		showCreateModal: true,
		selectedDeleteId: null,
		newUser: {

		}
	}),
	mounted() {
		this.getUsers();
		this.getRoles();
	},
	methods: {
		/*
		 * getUsers()
		 * Obtain the users or resources & apply query strings.
		 * NOTE: paramsSerializer is required here.
		 */
		getUsers() {
			this.axios.get(`users?order=${this.order}&filter=${this.filter}&${this.pagination}`, {
				paramsSerializer: function(params) {
					return params;
				}
			})
				.then(res => {
					this.users = {};
					this.paginationObj = {};
					this.paginationObj = res.data.meta.pagination;
					const users = res.data.data
					if (users.length) {
						this.users = this.removeSelf(users);
					}
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
		 * getRoles()
		 * Obtain all roles from API for use with creating a new user.
		 */
		getRoles() {
			this.axios.get("/roles")
				.then(res => {
					this.roles = res.data.data;
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
		filterTabs(e) {
			this.activeTab = e;
			let filter = "";
			switch (e) {
				case 2: {
					filter = '{"roles.name":[{"operator":"=", "value": "b=Banned"}]}';
					break;
				}
				case 3: {
					filter = '{"roles.name":[{"operator":"=", "value": "Contributor"}]}';
					break;
				}
				case 4: {
					filter = '{"roles.name":[{"operator":"=", "value": "Author"}]}';
					break;
				}
				case 5: {
					filter = '{"roles.name":[{"operator":"=", "value": "Editor"}]}';
					break;
				}
				case 6: {
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
		makeid(length) {
			let result = '';
			const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@:\\/@£$%=^&&*()_+?><';
			const charactersLength = characters.length;
			for (let i = 0; i < length; i++) {
				result += characters.charAt(Math.floor(Math.random() * charactersLength));
			}
			return result;
		}
	},
	computed: {
		/*
		 * getUserInfo()
		 * Get the logged in user info from the store.
		 */
		getUserInfo() {
			return this.$store.state.userInfo;
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

	// Create Modal
	// =========================================================================

	&-modal-create {

		.modal-container {
			max-width: 700px;
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