<!-- =====================
	Console
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Proxies</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<!-- =====================
					Details
					===================== -->
				<div class="col-12 col-desk-5">
					<!-- Header -->
					<div class="proxies-info-header">
						<h2 class="mb-0">Reverse Proxies:</h2>
						<p>A reverse proxy server is a type of proxy server that typically sits behind the firewall in a private network and directs client requests to the appropriate backend server. A reverse proxy provides an additional level of abstraction and control to ensure the smooth flow of network traffic between clients and servers.</p>
						<p></p>
					</div>
					<!-- Rewrites -->
					<div class="proxies-info-rewrite">
						<h3 class="mb-1">Rewrites</h3>
						<p>Rewrite defines URL path rewrite rules. The values captured in asterisk can be retrieved by index e.g. $1, $2 and so on.</p>
						<h6>Examples:</h6>
						<el-table size="small" :data="rewriteExamples" border style="width: 100%; margin-top: 10px">
							<el-table-column prop="from" label="From Path"></el-table-column>
							<el-table-column prop="to" label="To Path"></el-table-column>
						</el-table>
					</div>
					<!-- Regex Rewrites -->
					<div class="proxies-info-rewrite">
						<h3 class="mb-1">Regex Rewrites</h3>
						<p>Regex Rewrites defines rewrite rules using regexp exppresions with captures. Every capture group in the values can be retrieved by index e.g. $1, $2 and so on.</p>
						<h6>Examples:</h6>
						<el-table size="small" :data="rewriteExamples" border style="width: 100%; margin-top: 10px">
							<el-table-column prop="from" label="From Path"></el-table-column>
							<el-table-column prop="to" label="To Path"></el-table-column>
						</el-table>
					</div>
				</div><!-- /Col -->
				<!-- =====================
					Proxies
					===================== -->
				<div v-loading="doingAxios" class="col-12 col-desk-6 offset-desk-1">
					<!-- Header -->
					<div class="proxies-config-header">
						<h2>Configuration</h2>
						<el-button size="small" @click="showNewModal = true">New Proxy</el-button>
					</div>
					<!-- Proxies -->
					<draggable v-model="proxies" draggable=".item">
						<div v-for="(proxy, index) in proxies" :key="index" class="item">
							<el-card class="box-card" shadow="never">
							<h2>{{ proxy.host }}</h2>
							<h2>{{ proxy.path }}</h2>
							<h2>{{ proxy.rewrite }}</h2>
							<h2>{{ proxy.rewrite_regex }}</h2>
							</el-card>
						</div>
					</draggable><!-- /Proxies -->

				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Create New Modal
			===================== -->
		<NewProxyModal :visible.sync="showNewModal" @create="createProxy"></NewProxyModal>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import NewProxyModal from "../../components/proxies/NewProxyModal";
import {optionsMixin} from "@/util/options";
import draggable from 'vuedraggable'

export default {
	name: "Proxies",
	title: "Proxies",
	mixins: [optionsMixin],
	components: {
		Breadcrumbs,
		draggable,
		NewProxyModal,
	},
	data: () => ({
		form: {},
		rules: {
			name: [
				{ required: true, message: 'Enter link text for the menu item', trigger: 'blur' },
			],
		},
		showNewModal: false,
		rewriteExamples: [
			{from: '/old', to: '/new'},
			{from: '/api/*', to: '/$1'},
			{from: '/js/*', to: '/public/javascripts/$1'},
			{from: '/users/*/orders/*', to: '/user/$1/order/$2'},
		],
		regexExamples: [
			{from: '^/old/[0.9]+/', to: '/new'},
			{from: '^/api/.+?/(.*)', to: '/v2/$1'},
		]
	}),
	methods: {
		/**
		 * Creates a new reverse proxy and adds to the
		 * proxies array.
		 */
		createProxy(proxy) {
			this.proxies.push(proxy);
			this.showNewModal = false;
		},
	},
	computed: {
		/**
		 *
		 */
		proxies: {
			get() {
				return this.data['proxies'];
			},
			set(el) {
				this.data['proxies'] = el;
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Proxies
// =========================================================================

.proxies {

	// Config
	// =========================================================================

	&-config {

		&-header {
			display: flex;
			align-items: center;
			justify-content: space-between;
			margin-bottom: 1rem;

			h2 {
				margin-bottom: 0;
			}
		}
	}

	// Info
	// =========================================================================

	&-info {

		&-header {
			margin-bottom: 2rem;
		}

		&-rewrite {
			margin-bottom: 2rem;
		}
	}
}

</style>
