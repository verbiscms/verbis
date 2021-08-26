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
			<div class="row proxies">
				<!-- =====================
					Details
					===================== -->
				<div class="col-12 col-desk-5">
					<!-- Header -->
					<div class="proxies-info-header">
						<h2 class="mb-1">Reverse Proxies:</h2>
						<p class="mb-2">A reverse proxy accepts a request from a client, forwards it to a server that can fulfill it, and returns the server’s response to the client. Verbis allows the use of reverse proxies to proxy traffic to external websites.</p>
						<el-link href="https://verbiscms.com" target="_blank">Visit documentation</el-link>
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
						<el-button size="small" @click="showModal = true">New Proxy</el-button>
					</div>
					<el-alert class="mb-3" title="Order" type="warning" description="The order of which the proxies are defined are important for more info visit the link to the left." show-icon></el-alert>
					<!-- Proxies -->
					<draggable
						class="proxies-draggable"
						:class="{ 'proxies-draggable-dragging' : drag }"
						v-if="proxies && proxies.length"
						v-model="proxies"
						@start="drag = true"
						@end="drag = false">
						<div v-for="(proxy, index) in proxies" :key="index" class="proxies-draggable-item">
							<el-card class="box-card" shadow="never">
								<!-- Header -->
								<template #header>
									<div class="proxies-config-item-header">
										<span>{{ proxy.path }}</span>
										<el-button-group>
											<el-button size="small" icon="el-icon-edit" @click="editProxy(index)"></el-button>
											<el-button size="small" icon="el-icon-rank"></el-button>
											<el-popconfirm confirmButtonText="Yes" cancelButtonText="No" icon="el-icon-danger" iconColor="red" title="Are you sure to delete this proxy?" @confirm="deleteProxy(index)">
												<template #reference>
													<el-button size="small" icon="el-icon-delete"></el-button>
												</template>
											</el-popconfirm>
										</el-button-group>
									</div>
								</template>
								<!-- Path -->
								<div class="proxies-config-item">
									<h6>Path:</h6>
									<p>{{ proxy.path }}</p>
								</div>
								<!-- Host -->
								<div class="proxies-config-item">
									<h6>Host:</h6>
									<p>{{ proxy.host }}</p>
								</div>
								<!-- Rewrites -->
								<div class="proxies-config-item">
									<h6>Rewrites:</h6>
									<el-table v-if="proxy.rewrites && proxy['rewrites'].length" size="mini" :data="proxy.rewrites" border style="width: 100%; margin-top: 10px">
										<el-table-column prop="from" label="From Path"></el-table-column>
										<el-table-column prop="to" label="To Path"></el-table-column>
									</el-table>
									<p v-else>No rewrites set</p>
								</div><!-- /Rewrites -->
								<!-- Regex -->
								<div class="proxies-config-item">
									<h6>Regex Rewrites:</h6>
									<el-table v-if="proxy.rewrites && proxy['rewrites'].length" size="mini" :data="proxy.rewrites" border style="width: 100%; margin-top: 10px">
										<el-table-column prop="from" label="From Path"></el-table-column>
										<el-table-column prop="to" label="To Path"></el-table-column>
									</el-table>
									<p v-else>No rewrites set</p>
								</div><!-- /Regex -->
							</el-card>
						</div>
					</draggable><!-- /Proxies -->
					<el-empty v-else :image-size="100">
						<h4>No proxies available</h4>
						<p>Click the button above to create a new proxy</p>
					</el-empty>

				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Create New Modal
			===================== -->
		<NewProxyModal :visible.sync="showModal" @create="createProxy"></NewProxyModal>
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
		showModal: false,
		activeCollapse: "",
		updatingProxy: {},
		drag: false,
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
	created() {
		setTimeout(() => {
			this.proxies =  [
				{
					"host": "https://35.214.23.223:5000/tools/serp-speed",
					"path": "/tools/serp-speed",
					"rewrites": {
						"map/$1":"map/$3",
					},
					"rewrite_regex": [],
				},
				{
					"host": "https://35.214.23.223:5000/tools/serp-speed",
					"path": "/tools/reredirect",
					"rewrites": {
						"map/$1":"map/$3",
					},
					"rewrite_regex": [],
				}
			]
		}, 500)
	},
	methods: {
		/**
		 * Creates a new reverse proxy and adds to the
		 * proxies array.
		 */
		createProxy(proxy) {
			this.proxies.push(proxy);
			this.showModal = false;
		},
		deleteProxy(index) {
			if (index !== -1) {
				this.proxies.splice(index, 1);
			}
		},
		editProxy(index) {
			this.updatingProxy = this.proxies[index];
			this.showModal = true;
		}
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

	// Props
	// =========================================================================

	::v-deep {

		.el-empty__description {
			display: none !important;
		}

		.el-alert__description {
			margin-top: 0;
		}
	}

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

		&-item {
			margin-bottom: 10px;
			display: flex;
			align-items: center;
			justify-content: space-between;

			p {
				margin-bottom: 0;
			}

			&:last-child {
				margin-bottom: 0;
			}

			&-header {
				display: flex;
				justify-content: space-between;
				align-items: center;

				span {
					// TODO change to Element UI colour var
					color: rgb(121, 187, 255);
					font-weight: 600;
					font-size: 0.9rem;
				}
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
			margin-bottom: 2.4rem;
		}
	}

	// Draggable
	// =========================================================================

	&-draggable {
		padding: 1rem;
		// TODO change to Element UI colour var
		border: 1px dashed rgba($grey, 0.3);
		border-radius: 4px;

		&-item {
			margin-bottom: 1rem;
		}

		&-dragging {
			// TODO change to Element UI colour var
			background-color: #ecf5ff;
			border: 1px dashed #409EFF;
		}
	}
}

</style>
