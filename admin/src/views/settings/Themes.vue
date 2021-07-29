<!-- =====================
	Settings - General
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Themes</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-if="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Basic Options
					===================== -->
				<div class="col-12 col-tab-6 col-desk-4" v-for="(theme, themeIndex) in themes" :key="themeIndex">
					<div class="card">
						<div class="card-image theme-image">
							<img :src="theme['screenshot']" alt="">
						</div>
						<div class="card-body">
							<h3 class="mb-1">{{ theme['title'] }}</h3>
							<p class="mb-1">{{ theme['description'] }}</p>
							<p>Version: {{ theme['version'] }}</p>
							<button class="btn btn-block" :class='{ "btn-green" : !isActive(theme.name) }' :disabled="isActive(theme.name)" @click="setTheme(theme.name)">
								<span v-if="isActive(theme.name)">Activated</span>
								<span v-else>Activate</span>
							</button>
						</div>
					</div>
				</div>
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import {validationMixin} from "@/util/validation";

export default {
	name: "General",
	title: 'General Settings',
	mixins: [optionsMixin, validationMixin],
	components: {
		Breadcrumbs,
	},
	data: () => ({
		errorMsg: "Fix the errors before saving settings.",
		successMsg: "Site options updated successfully.",
		themes: [],
	}),
	mounted() {
		this.getThemes();
	},
	methods: {
		/*
		 * getThemes()
		 */
		getThemes() {
			this.axios.get("/themes")
				.then(res => {
					this.themes = res.data.data.map(t => t.theme);
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
				});
		},
		/*
		 * setTheme()
		 */
		setTheme(name) {
			this.axios.post("/themes/" + name)
				.then(res => {
					this.$store.commit("setTheme", res.data.data);
					this.$set(this.data, 'active_theme', name)
				}).catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * isActive()
		 * Determines if a theme is active by comparing
		 * the options & name passed.
		 */
		isActive(name) {
			return name === this.data['active_theme'];
		}
	},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Theme
	// =========================================================================

	.theme {

		&-image {
			height: 350px;
			border-bottom: 1px solid $grey-light;

			img {
				width: 100%;
				height: 100%;
				border-top-left-radius: 6px;
				border-top-right-radius : 6px;
				object-fit: cover;
			}
		}
	}

</style>
