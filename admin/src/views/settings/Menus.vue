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
				<div class="col-12 col-desk-4" v-if="activeMenuKey !== ''">
					<h3 class="mb-3">Add items</h3>
					<el-card class="box-card" :body-style="{ padding: '0' }" shadow="never">
						<el-collapse class="collapse collapse-bg-header" accordion v-model="activeCollapse">
							<!-- Posts -->
							<el-collapse-item title="Posts" name="posts">
								<MenuPostsFilter @update="addPosts"></MenuPostsFilter>
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
									<el-button size="small" plain @click="addExternal('newExternalItem')">
										Add to Menu
									</el-button>
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
				<div class="col-12 col-desk-8" v-if="activeMenuKey !== '' && !doingAxios">
					<!-- Hierarchy -->
					<h3 class="mb-3">Menu hierarchy</h3>
					<MenuItems class="mb-5" :items.sync="menus[activeMenuKey].items"></MenuItems>
					<!-- Attributes -->
					<h3 class="mb-3">Menu attributes</h3>
					<el-card class="menu-attributes" shadow="never">

					</el-card>
				</div><!-- /Col -->
				<el-empty v-if="activeMenuKey === ''" description="No Menus">
					<el-button type="primary" @click="showNewModal = true">Create a new Menu</el-button>
				</el-empty>
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Create New Modal
			===================== -->
		<MenuNewModal :visible.sync="showNewModal" :menus="getThemeMenus" @create="createMenu"></MenuNewModal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import MenuPostsFilter from "../../components/menu/MenuPostsFilter";
import MenuItems from "../../components/menu/MenuItems";
import MenuNewModal from "../../components/menu/MenuNewModal";

const ID_START = 10000;

export default {
	name: "Menus",
	title: 'Menu Settings',
	mixins: [optionsMixin],
	components: {
		MenuNewModal,
		Breadcrumbs,
		MenuPostsFilter,
		MenuItems,
	},
	data: () => ({
		errorMsg: "Fix the errors before saving performance settings.",
		successMsg: "Performance options updated successfully.",
		activeMenuKey: "main-menu",
		showNewModal: false,
		newItem: {},
		dialogVisible: false,
		isDragging: false,
		activeCollapse: "",
		doingBulk: false
	}),
	methods: {
		/**
		 * Creates a new menu and adds it the menus
		 * object. If form validation is rejected
		 * the function will return.
		 */
		createMenu(menu) {
			if (menu.id in this.menus) {
				this.$set(this.menus, "unassigned-" + this.helpers.randomString("10"), this.menus[this.newMenu.id]);
			}
			this.$set(this.menus, menu.id, {
				name: menu.name,
				items: [],
			});
			this.activeMenuKey = menu.id;
			this.showNewModal = false;
		},
		/**
		 * Deletes the currently active menu.
		 */
		deleteMenu() {
			let key = this.activeMenuKey;
			this.activeMenuKey = "";
			this.$delete(this.menus, key);
		},
		/**
		 * Adds a post to the menu hierarchy.
		 * @param items
		 */
		addPosts(items) {
			items.forEach(item => {
				this.addNewItem(item.post.title, item.post.url, {
					post_id: item.post.id,
				});
			})
		},
		/**
		 * Adds an external item to the menu hierarchy.
		 * @param formName
		 */
		addExternal(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.addNewItem(this.newItem.text, this.newItem.href);
					this.newItem = {};
				}
			});
		},
		/**
		 * Generates a new item with ID, text, href and
		 * default values.
		 * @param text
		 * @param href
		 * @param opts
		 */
		addNewItem(text, href, opts = {}) {
			this.menus[this.activeMenuKey].items.push({
				id: this.getID(),
				text: text,
				href: href,
				rel: [],
				li_classes: [],
				...opts
			});
		},
		/**
		 * Returns a unique identifier for the menu item,
		 * if there are no items in the menu, ID_START
		 * will be returned.
		 * @return {number}
		 */
		getID() {
			const items = this.menus[this.activeMenuKey].items;
			if (!items.length) {
				return ID_START;
			}
			const reduced = this.flattenItems(items);
			return Math.max.apply(Math, reduced.map(function(o) { return o.id; })) + 1;
		},
		/**
		 * Reduces the menu hierarchy and returns a
		 * flattened array.
		 * @param items
		 */
		flattenItems(items) {
			return items.reduce((acc, item) => {
				acc.push(item);
				if (item.children && item.children.length) {
					return acc.concat(this.flattenItems(item.children));
				}
				return acc;
			}, []);
		},
		/**
		 * Returns the current active menu.
		 * @returns {*}
		 */
		activeMenu() {
			return this.menus[this.activeMenuKey];
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

// Menu
// =========================================================================

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
