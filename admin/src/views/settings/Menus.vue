<!-- =====================
	Settings - Menus
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Menus</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<el-button type="primary" @click.prevent="save" :loading="saving">Update Menus</el-button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div v-if="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Menu Picker
					===================== -->
				<div class="col-12">
					<h3 class="mb-3">Select a menu to edit</h3>
					<el-card class="box-card mb-4" shadow="never">
						<el-select v-model="activeMenuKey" placeholder="Select">
							<el-option
								v-for="(menu, key) in menus"
								:key="key"
								:label="menu.name"
								:value="key">
							</el-option>
						</el-select>
						<el-link class="ml-2" @click="showNewModal = true">Or create a new one</el-link>
					</el-card>
				</div>
				<!-- =====================
					Item Picker
					===================== -->
				<div class="col-12 col-desk-3">
					<h3 class="mb-3">Add items</h3>
					<el-card class="box-card" :body-style="{ padding: '0' }" shadow="never">
						<el-collapse class="test" v-model="activeCollapse">
							<!-- Posts -->
							<el-collapse-item title="Posts" name="posts">
								<NavPostsFilter @update="addPosts"></NavPostsFilter>
							</el-collapse-item>
							<!-- External -->
							<el-collapse-item title="External" name="external">
								<el-form :model="newItem" ref="newExternalItem" label-position="top">
									<!-- Link Text -->
									<el-form-item label="Link Text" prop="text" :rules="{ required: true, message: 'Enter a link text.', trigger: 'blur' }">
										<el-input placeholder="Link Text*" label="Link Text*" v-model="newItem.text" clearable></el-input>
									</el-form-item>
									<!-- URL -->
									<el-form-item label="Link Text" prop="url" :rules="{ required: true, message: 'Enter a link URL.', trigger: 'blur' }">
										<el-input placeholder="URL*" label="Link URL*" v-model="newItem.url" clearable></el-input>
									</el-form-item>
									<!-- Add to Menu -->
									<el-button size="medium" plain @click="addExternal('newExternalItem')">Add to Menu</el-button>
								</el-form>
							</el-collapse-item>
							<!-- Categories -->
							<el-collapse-item disabled title="Categories" name="categories">
								<!-- TODO -->
							</el-collapse-item>
						</el-collapse>
					</el-card>
				</div><!-- /Col -->
				<!-- =====================
					Items
					===================== -->
				<div class="col-12 col-desk-9" v-if="!doingAxios">
					<h3 class="mb-3">Menu hierarchy</h3>
					<vue-nestable
						v-model="activeMenu().items"
						:max-depth="10"
						:threshold="30"
						children-prop="children"
						@input="isDragging = true"
						@change="handleAfterDrag">
						<vue-nestable-handle slot-scope="{ item }" :item="item">
							<NavItem
								@remove="removeItem"
								:disabled="isDragging"
								:item="item"></NavItem>
						</vue-nestable-handle>
					</vue-nestable>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Create New Modal
			===================== -->
		<el-dialog :visible.sync="showNewModal" width="30%">
			<!-- Title -->
			<div slot="title">
				<h2 class="mb-0">Create a menu</h2>
				<p>Select a menu location and assign the new menu a name.</p>
			</div>
			<!-- Form -->
			<el-form :model="newItem" ref="newMenu" label-position="left" label-width="auto ">
				<!-- Name -->
				<el-form-item
					label="Name"
					prop="name"
					:rules="{ required: true, message: 'Enter a Menu Name.', trigger: 'blur' }">
					<el-input placeholder="Name" v-model="newMenu.name"></el-input>
				</el-form-item>
				<!-- Location -->
				<el-form-item
					label="Location"
					prop="location"
					:rules="{ required: true, message: 'Enter a Menu Location.', trigger: 'change' }">
					<el-select v-model="newMenu.id" placeholder="Select">
						<el-option v-for="menu in getThemeMenus" :key="menu.id" :label="menu.name" :value="menu.id"></el-option>
					</el-select>
				</el-form-item>
			</el-form>
			<!-- Footer -->
			<span slot="footer" class="dialog-footer">
				<el-button @click="showNewModal = false">Cancel</el-button>
				<el-button type="primary" @click="createMenu('newMenu')">Create Menu</el-button>
			</span>
		</el-dialog>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import NavItem from "@/components/nav/NavItem";
import NavPostsFilter from "../../components/nav/NavPostsFilter";


export default {
	name: "Menus",
	title: 'Menu Settings',
	mixins: [optionsMixin],
	components: {
		NavPostsFilter,
		Breadcrumbs,
		NavItem,

	},
	data: () => ({
		errorMsg: "Fix the errors before saving performance settings.",
		successMsg: "Performance options updated successfully.",
		activeMenuKey: "main-menu",
		showNewModal: false,
		newMenu: {
			name: "",
			id: "",
		},
		newItem: {},
		dialogVisible: false,
		isDragging: false,
		activeCollapse: "",
	}),
	created() {
		setTimeout(() => {
			this.isDragging = false;
		}, 200);
	},
	methods: {
		createMenu(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					if (this.newMenu.id in this.menus) {
						this.$set(this.menus, "unassigned-" + this.helpers.randomString("10"), this.menus[this.newMenu.id]);
					}

					this.$set(this.menus, this.newMenu.id, {
						name: this.newMenu.name
					});

					this.activeMenuKey = this.newMenu.id;
					this.newMenu = {};
					this.showNewModal = false;
				}
			});
		},
		deleteMenu(menu) {
			this.$delete(this.menus, menu)
		},
		/*
		 * removeItem()
		 * Removes an item from the hierarchy.
		 */
		removeItem(item) {
			const items = this.menus[this.activeMenuKey].items;

			this.testRemove(items, item.id);
		},

		testRemove(items, id) {

			items.forEach(item => {
				if (item.id === id ) {
					this.menus.splice()
				}
				console.log(item.id, id);
			});
		},

		addPosts(items) {
			items.forEach(item => {
				this.menus[this.activeMenuKey].items.push({
					post_id: item.post.id,
					text: item.post.title,
				});
			})
		},
		addExternal(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					console.log(this.newItem);
				}
			});
		},
		activeMenu() {
			return this.menus[this.activeMenuKey];
		},
		handleAfterDrag() {
			this.$nextTick(() => {
				this.isDragging = false;
			}, 400);
		},
	},
	computed: {
		/*
		 * menus()
		 * Returns the menus already saved within
		 * the store. And sets the options
		 * whe modelling.
		 */
		menus: {
			get() {
				return this.data['nav_menus'];
			},
			set(el) {
				this.data['nav_menus'] = el;
			}
		},
		/*
		 * getThemeMenus()
		 * Returns the theme menus stored in the
		 * theme config file.
		 */
		getThemeMenus() {
			return this.$store.state.theme['menus'];
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">


/*
* Style for nestable
*/
.nestable {
	position: relative;
}
.nestable-rtl {
	direction: rtl;
}
.nestable .nestable-list {
	margin: 0;
	padding: 0 0 0 40px;
	list-style-type: none;
}
.nestable-rtl .nestable-list {
	padding: 0 40px 0 0;
}
.nestable > .nestable-list {
	padding: 0;
}
.nestable-item,
.nestable-item-copy {
	margin: 10px 0 0;
}
.nestable-item:first-child,
.nestable-item-copy:first-child {
	margin-top: 0;
}
.nestable-item .nestable-list,
.nestable-item-copy .nestable-list {
	margin-top: 10px;
}
.nestable-item {
	position: relative;
}
.nestable-item.is-dragging .nestable-list {
	pointer-events: none;
}
.nestable-item.is-dragging * {
	opacity: 0;
	filter: alpha(opacity=0);
}

.nestable-item.is-dragging {

	.item {
		height: 55px !important;
	}
}

.nestable-item.is-dragging:before {
	content: ' ';
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-color: rgba($primary, 0.274);
	border: 1px dashed rgb(73, 100, 241);
	-webkit-border-radius: 5px;
	border-radius: 5px;
}
.nestable-drag-layer {
	position: fixed;
	top: 0;
	left: 0;
	z-index: 100;
	pointer-events: none;
}
.nestable-rtl .nestable-drag-layer {
	left: auto;
	right: 0;
}
.nestable-drag-layer > .nestable-list {
	//position: absolute;
	//top: 0;
	//left: 0;
	//padding: 0;
	//background-color: rgba(106, 127, 233, 0.274);
}
.nestable-rtl .nestable-drag-layer > .nestable-list {
	padding: 0;
}
.nestable [draggable="true"]  {
	cursor: move;

	.collapse-content {
		cursor: default;
	}
}


	.menu {

		// Picker
		// =========================================================================

		&-picker {
			display: flex;
			align-items: center;
			padding: 20px;

			&-select {
				margin-bottom: 0;
				min-width: 30%;
				margin-right: 1rem;
			}

			a {
				cursor: pointer;
			}
		}
	}

</style>
