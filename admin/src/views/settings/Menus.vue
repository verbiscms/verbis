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
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Menus</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Menu Picker
					===================== -->
				<div class="col-12">
					<h3 class="mb-3">Select a menu to edit</h3>
					<div class="menu-picker card card-padding-large">
						<div class="menu-picker-select">
							<!-- Menu -->
							<FormGroup class="form-group-no-margin form-select-cont form-input">
								<select class="form-select" v-model="activeMenuKey">
									<option v-for="(menu, key) in menus" :key="key" :value="key">{{ menu['name'] }}</option>
								</select>
							</FormGroup>
						</div>
						<a @click="showNewModal = true;">Or create a new menu</a>
					</div>
				</div>
				<!-- =====================
					Item Picker
					===================== -->
				<div class="col-12 col-desk-3">
					<h3 class="mb-3">Add items</h3>
					<div class="card card-padding-large card-small-box-shadow">

					</div>
				</div><!-- /Col -->
				<!-- =====================
					Items
					===================== -->
				<div class="col-12 col-desk-9" v-if="!doingAxios">
					<h3 class="mb-3">Menu hierarchy</h3>
					<vue-nestable v-model="activeMenu().items" :max-depth="10" :threshold="30" children-prop="children">
						<vue-nestable-handle slot-scope="{ item }" :item="item">
							<NavItem :item="item"></NavItem>
						</vue-nestable-handle>
					</vue-nestable>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Provider Modal
			===================== -->
		<Modal :show.sync="showNewModal">
			<template slot="button">
				<button class="btn" @click="createMenu()">Create Menu</button>
			</template>
			<template slot="text">
				<!-- Create Menu -->
				<div class="text-cont">
					<h2>Create a menu</h2>
					<p class="t-left">Select a menu location and assign the new menu a name.</p>
				</div>
				<!-- Name -->
				<FormGroup label="Name*" :error="errors['new_menu_name']">
					<input placeholder="Enter a menu name" class="form-input form-input-white" type="text" v-model="newMenu['name']">
				</FormGroup>
				<!-- Location -->
				<FormGroup label="Location*" :error="errors['new_menu_id']">
					<div class="form-select-cont form-input">
						<select class="form-select" v-model="newMenu['id']">
							<option value="" disabled selected>Select a location</option>
							<option v-for="menu in getThemeMenus" :key="menu['id']" :value="menu['id']">{{ menu.name }}</option>
						</select>
					</div>
				</FormGroup>
			</template>
		</Modal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import Modal from "@/components/modals/General";
import NavItem from "@/components/nav/NavItem";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Menus",
	title: 'Menu Settings',
	mixins: [optionsMixin],
	components: {
		Breadcrumbs,
		FormGroup,
		Modal,
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
	}),
	methods: {
		//runAfterGet() {
		//	this.activeMenu =
		//},
		createMenu() {
			this.errors = {};

			if (this.newMenu['name'] === "") {
				this.$set(this.errors, 'new_menu_name', "Enter a name for the menu");
			}
			if (this.newMenu['id'] === "") {
				this.$set(this.errors, 'new_menu_id', "Select a menu location");
			}

			if (!this.helpers.isEmptyObject(this.errors)) {
				return
			}

			if (this.newMenu.id in this.menus) {
				this.$set(this.menus, "unassigned-" + this.helpers.randomString("10"), this.menus[this.newMenu.id]);
			}

			this.$set(this.menus, this.newMenu.id, {
				name: this.newMenu.name
			});

			this.activeMenuKey = this.newMenu.id;
			this.newMenu = {};
			this.showNewModal = false;
		},

		deleteMenu(menu) {
			this.$delete(this.menus, menu)
		},
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
