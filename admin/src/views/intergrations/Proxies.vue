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
			<el-row class="row">
				<!-- =====================
					Details
					===================== -->
				<div class="col-12 col-desk-5">
					<!-- Info -->
					<div class="info">
						<!-- Header -->
						<div class="info-header">
							<h2 class="mb-1">Reverse Proxies:</h2>
							<p class="mb-2">A reverse proxy accepts a request from a client, forwards it to a server that
								can fulfill it, and returns the server’s response to the client. Verbis allows the use of
								reverse proxies to proxy traffic to external websites.</p>
							<el-link href="https://verbiscms.com" target="_blank">Visit documentation</el-link>
						</div>
						<!-- Rewrites -->
						<div class="info-rewrite">
							<h3 class="mb-1">Rewrites</h3>
							<p>Rewrite defines URL path rewrite rules. The values captured in asterisk can be retrieved by
								index e.g. $1, $2 and so on.</p>
							<h6>Examples:</h6>
							<el-table size="small" :data="rewriteExamples" border style="width: 100%; margin-top: 10px">
								<el-table-column prop="from" label="From Path"></el-table-column>
								<el-table-column prop="to" label="To Path"></el-table-column>
							</el-table>
						</div>
						<!-- Regex Rewrites -->
						<div class="info-rewrite">
							<h3 class="mb-1">Regex Rewrites</h3>
							<p>Regex Rewrites defines rewrite rules using regexp exppresions with captures. Every capture
								group in the values can be retrieved by index e.g. $1, $2 and so on.</p>
							<h6>Examples:</h6>
							<el-table size="small" :data="rewriteExamples" border style="width: 100%; margin-top: 10px">
								<el-table-column prop="from" label="From Path"></el-table-column>
								<el-table-column prop="to" label="To Path"></el-table-column>
							</el-table>
						</div>
					</div><!-- /Info -->
				</div><!-- /Col -->
				<!-- =====================
					Proxies
					===================== -->
				<div class="col-12 col-desk-6 offset-desk-1">
					<!-- Config -->
					<div class="config">
						<!-- Header -->
						<div class="config-header">
							<h2>Configuration</h2>
							<el-button size="small" @click="newProxy">New Proxy</el-button>
						</div>
						<!-- Order Alert -->
						<el-alert class="config-alert" title="Order" type="warning" description="The order of which the proxies are defined are important for more info visit the link to the left." show-icon></el-alert>
					</div><!-- /Config -->
					<!-- Proxies -->
					<el-form :model="form" size="small" ref="proxiesForm" label-width="80px" label-position="left" v-if="proxies && proxies.length">
						<el-collapse class="proxies" v-model="activeCollapse" accordion>
							<draggable
								class="proxies-draggable"
								:class="{ 'proxies-draggable-dragging' : drag }"
								v-model="proxies"
								handle=".proxies-handle"
								v-loading="doingAxios"
								@start="handleDragStart"
								@end="drag = false">
								<el-collapse-item v-for="(proxy, index) in proxies" :key="index" :disabled="true" class="proxies-item" :name="proxy.name">
									<!-- Header -->
									<template slot="title">
										<div class="proxies-header">
											<span>{{ proxy.name }}</span>
											<el-button-group class="proxies-header-btns">
												<el-button size="mini" icon="el-icon-edit" @click="handleCollapse(proxy)"></el-button>
												<el-button size="mini" icon="el-icon-rank" class="proxies-handle"></el-button>
												<el-popconfirm confirmButtonText="Yes" cancelButtonText="No" icon="el-icon-danger" iconColor="red" title="Are you sure to delete this proxy?" @confirm="deleteProxy(index)">
													<template #reference>
														<el-button style="border-radius: 0 4px 4px 0;" size="mini" icon="el-icon-delete"></el-button>
													</template>
												</el-popconfirm>
											</el-button-group>
										</div>
									</template><!-- /Header -->
									<!-- Body -->
									<div class="proxies-body">
										<!-- Name -->
										<el-form-item label="Name" prop="name" :rules="{ required: true, message: 'Enter a Name.', trigger: 'blur' }">
											<el-input placeholder="Add a name" v-model="proxies[index].path"></el-input>
										</el-form-item>
										<!-- Path -->
										<el-form-item label="Path" prop="path" :rules="{ required: true, message: 'Enter a Path.', trigger: 'blur' }">
											<el-input placeholder="Add a path" v-model="proxies[index].path"></el-input>
										</el-form-item>
										<!-- Host -->
										<el-form-item label="Host" prop="host" :rules="{ required: true, message: 'Enter a Host.', trigger: 'blur' }">
											<el-input placeholder="Add a host" v-model="proxies[index].path"></el-input>
										</el-form-item>
										<!-- Rewrites -->
										<el-form-item label="Rewrites" prop="rewrites">
											<el-table v-if="proxy.rewrites && proxy['rewrites'].length" size="mini" :data="proxy.rewrites" border style="width: 100%; margin-top: 10px">
												<el-table-column prop="from" label="From Path"></el-table-column>
												<el-table-column prop="to" label="To Path"></el-table-column>
											</el-table>
											<p v-else>No rewrites set</p>
										</el-form-item>
									</div><!-- /Body -->
								</el-collapse-item>
							</draggable>
						</el-collapse><!-- /Proxies -->
					</el-form>
					<el-empty v-else :image-size="100">
						<h4>No proxies available</h4>
						<p>Click the button above to create a new proxy</p>
					</el-empty>
				</div><!-- /Col -->
			</el-row><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import draggable from 'vuedraggable'

const UNASSIGNED_PREFIX = "Unassigned";

export default {
	name: "Proxies",
	title: "Proxies",
	mixins: [optionsMixin],
	components: {
		Breadcrumbs,
		draggable,
	},
	data: () => ({
		form: {},
		rules: {
			name: [
				{required: true, message: 'Enter link text for the menu item', trigger: 'blur'},
			],
		},
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
			this.$set(this.data, 'proxies', [
				{
					"name": "Serp Speed",
					"host": "https://35.214.23.223:5000/tools/serp-speed",
					"path": "/tools/serp-speed",
					"rewrites": [
						{from: "mappp", to: "ap/$"},
					],
					"rewrite_regex": [],
				},
				{
					"name": "ReReDirect",
					"host": "https://35.214.23.223:5000/tools/serp-speed",
					"path": "/tools/reredirect",
					"rewrites": {
						"map/$1": "map/$3",
					},
					"rewrite_regex": [],
				}
			]);
		}, 500)
	},
	methods: {
		/**
		 * Creates a new reverse proxy and adds to the
		 * proxies array.
		 */
		createProxy(proxy) {
			this.proxies.push(proxy);
		},
		/**
		 * Deletes a reverse proxy from the array.
		 */
		deleteProxy(index) {
			if (index !== -1) {
				this.proxies.splice(index, 1);
			}
		},
		/**
		 * Creates a new proxy and adds it to
		 * the array.
		 */
		newProxy() {
			this.proxies.push({
				name: this.getUnassignedName(),
				path: "",
				host: "",
				rewrites: [],
				regex_rewrites: [],
			})
		},
		/**
		 * Handles the accordion.
		 */
		handleCollapse(proxy) {
			this.activeCollapse = proxy.name;
		},
		/**
		 * Handle the start of a drag item, collapses
		 * all accordion items.
		 */
		handleDragStart() {
			this.activeCollapse = "";
			this.drag = true;
		},
		/**
		 * Retrieves the unassigned name for a proxy
		 * when no name is set. It will incrmement
		 * by one if none is found.
		 */
		getUnassignedName() {
			if (!this.proxies.find(el => el.name === UNASSIGNED_PREFIX)) {
				return UNASSIGNED_PREFIX;
			}
			let counter = 1;
			// eslint-disable-next-line no-constant-condition
			while (true) {
				let found = this.proxies.find(el => el.name === UNASSIGNED_PREFIX + "-" + counter);
				if (!found) {
					return UNASSIGNED_PREFIX + "-" + counter;
				}
				counter++;
			}
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


// Info
// =========================================================================

.info {

	&-header {
		margin-bottom: 2rem;
	}

	&-rewrite {
		margin-bottom: 2.4rem;
	}
}

// Config
// =========================================================================


.config {

	&-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 1rem;

		h2 {
			margin-bottom: 0;
		}
	}

	&-alert {
		margin-bottom: 1rem;
	}

	::v-deep {

		.el-alert__description {
			margin-top: 0;
		}
	}
}

// Proxies
// =========================================================================

.proxies {
	$self: &;

	// Props
	// =========================================================================

	::v-deep {

		.el-empty__description {
			display: none !important;
		}

		.el-collapse-item__header {
			border-bottom: 0;
			height: auto;
			line-height: 0;

			.el-collapse-item__arrow {
				display: none;
			}
		}

		.el-collapse-item__wrap {
			border-bottom: 0;
		}

		.el-form-item__label {
			color: $secondary;
		}

		.el-collapse-item.is-disabled .el-collapse-item__header {
			cursor: default;
		}

	}

	// Header
	// =========================================================================

	&-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding: 1rem;
		line-height: 0;

		span {
			// TODO change to Element UI colour var
			color: rgb(121, 187, 255);
			font-weight: 600;
			font-size: 0.9rem;
		}

		&-btns {
			display: flex;
			align-items: center;
		}
	}

	// Body
	// =========================================================================

	&-body {
		padding: 0 1rem;
	}

	// Handle
	// =========================================================================

	&-handle {
		cursor: all-scroll;
	}

	// Item
	// =========================================================================

	&-item {
		border: 1px solid rgba(#DCDFE6, 0.5);
		border-bottom: none;

		&:last-child {
			border-bottom: 1px solid rgba(#DCDFE6, 0.5);
		}
	}

	// Draggable
	// =========================================================================

	&-draggable {
		padding: 1rem;
		// TODO change to Element UI colour var
		border: 1px dashed rgba($grey, 0.3);
		border-radius: 4px;

		&-dragging {
			// TODO change to Element UI colour var
			border: 1px dashed #409EFF;
			cursor: pointer;

			#{$self}-card {
				background-color: #ecf5ff;
			}
		}
	}
}

</style>
