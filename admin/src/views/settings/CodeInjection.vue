<!-- =====================
	Settings - Code Injection
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Code Injection</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Settings</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
				<div class="col-12">
					<div class="text-cont">
						<h2>How to use</h2>
						<p>Code injection allows you to inject a small snippet of HTML into your site. It can be a css override, analytics of a block javascript. To insert code into a specific page, visit the page in the resources section.</p>
					</div>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<h6 class="margin">Header</h6>
					<!-- =====================
						Head
						===================== -->
					<div class="card card-small-box-shadow card-expand">
						<Collapse :show="true" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['codeinjection_head']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Head</h4>
										<p>Any code inputted here will be injected the <code v-html="'{{ verbisHead }}'"></code> of the site.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<prism-editor class="prism" v-model="data['codeinjection_head']" :highlight="highlighter" line-numbers></prism-editor>
								</div><!-- /Card Body -->
							</template>
						</Collapse>
					</div><!-- /Card -->
					<!-- =====================
						Footer
						===================== -->
					<h6 class="margin">Footer</h6>
					<div class="card card-small-box-shadow card-expand">
						<Collapse :show="true" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['codeinjection_foot']}">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Foot</h4>
										<p>Any code inputted here will be injected to the <code v-html="'{{ verbisFooter }}'"></code> before the closing body tag.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<prism-editor class="prism" v-model="data['codeinjection_foot']" :highlight="highlighter" line-numbers></prism-editor>
								</div><!-- /Card Body -->
							</template>
						</Collapse>
					</div><!-- /Card -->
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
import {optionsMixin} from "@/util/options";
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';
import Collapse from "@/components/misc/Collapse";

export default {
	name: "CodeInjection",
	title: "Code Injection",
	mixins: [optionsMixin],
	components: {
		Breadcrumbs,
		Collapse
	},
	data: () => ({
		errorMsg: "Fix the errors before saving code injection.",
		successMsg: "Code injection updated successfully.",
	}),
	methods: {
		/*
		 * highlighter()
		 * Return html for prism editor.
		 */
		highlighter(code) {
			return highlight(code, languages.html);
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Text
	// =========================================================================

	h2 {
		margin-bottom: 4px;
	}

	p {
		font-size: 0.9rem;
	}

	.card-body {
		width: 100%;
		padding-left: 10px
	}

</style>