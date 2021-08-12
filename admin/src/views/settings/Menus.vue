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
							<h1>Performance</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Settings</span>
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

				<pre>{{ menus }}</pre>
				<!-- =====================
					Menu Picker
					===================== -->
				<div class="col-12">
					<div v-for="(menu, key) in menus" :key="key">
						<select name="" id="">
							<option :value="key">{{ menu['name'] }}</option>
						</select>
					</div>
				</div>
				<!-- =====================
					Item Picker
					===================== -->
				<div class="col-12 col-desk-3">
					<h2>Hell</h2>

				</div><!-- /Col -->
				<!-- =====================
					Items
					===================== -->
				<div class="col-12 col-desk-9" v-if="!doingAxios">
					<div v-for="(item, itemIndex) in menus[activeMenu].items" :key="itemIndex">
							<div class="card card-small-box-shadow card-expand card-expand-full-width">
								<!--  -->
								<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_frontend'] }">
									<template v-slot:header>
										<div class="card-header">
											<h4 class="card-title">{{ item['text'] }}</h4>
											<div class="card-controls">
												<i class="feather feather-chevron-down"></i>
											</div>
										</div><!-- /Card Header -->
									</template>
									<template v-slot:body>
										<div class="card-body">
											<!-- Label -->
											<FormGroup label="Label*">
												<input class="form-input form-input-white" type="text" v-model="menus[activeMenu]['items'][itemIndex]['text']">
											</FormGroup>
											<!-- Title -->
											<FormGroup label="Title">
												<input class="form-input form-input-white" type="text" v-model="menus[activeMenu]['items'][itemIndex]['title']">
											</FormGroup>
											<!-- Rel -->
											<FormGroup label="Rel">
												<input class="form-input form-input-white" type="text" v-model="menus[activeMenu]['items'][itemIndex]['rel']">
											</FormGroup>
											<!-- Open Tab -->
											<FormGroup label="Open link in new tab">
												<div class="toggle">
													<input type="checkbox" class="toggle-switch" id="cache-frontend" checked v-model="menus[activeMenu]['items'][itemIndex]['new_tab']" :true-value="true" :false-value="false" />
													<label for="cache-frontend"></label>
												</div>
											</FormGroup>
										</div><!-- /Card Body -->
									</template>
								</Collapse><!-- /Cache assets? -->
							</div>

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
import Collapse from "@/components/misc/Collapse";
import {optionsMixin} from "@/util/options";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "Menus",
	title: 'Menu Settings',
	mixins: [optionsMixin],
	components: {
		Breadcrumbs,
		Collapse,
		FormGroup,
	},
	data: () => ({
		errorMsg: "Fix the errors before saving performance settings.",
		successMsg: "Performance options updated successfully.",
		activeMenu: "main-menu",
	}),
	methods: {
		//runAfterGet() {
		//	this.activeMenu =
		//},
		createMenu(menu) {
			this.menus[menu] = [];
		},
		deleteMenu(menu) {
			this.$delete(this.menus, menu)
		},
		// addItem(item) {
		//
		// },

	},
	computed: {
		menus: {
			get() {
				return this.data['nav_menus'];
			},
			set(el) {
				this.data['nav_menus'] = el;
			}
		}
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
