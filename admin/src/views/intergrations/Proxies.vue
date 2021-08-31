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
							<h1>Reverse Proxies</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<el-button type="primary" @click.prevent="saveProxies" :loading="saving">Update Proxies</el-button>
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
							<h2 class="mb-1">Information:</h2>
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
							<el-table size="small" :data="regexExamples" border style="width: 100%; margin-top: 10px">
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
					<div v-loading="doingAxios">
						<!-- Proxies -->
						<el-form v-if="proxies && proxies.length && !doingAxios">
							<el-collapse class="proxies" v-model="activeCollapse">
								<draggable
									class="proxies-draggable"
									:class="{ 'proxies-draggable-dragging' : drag }"
									v-model="proxies"
									handle=".proxies-handle"
									@start="handleDragStart"
									@end="drag = false">
									<el-collapse-item v-for="(proxy, index) in proxies" :key="index" :disabled="true" class="proxies-item" :name="index">
										<template slot="title">
											<!-- Header -->
											<div class="proxies-header" ref="proxyHeader">
												<!-- Info -->
												<div class="proxies-header-info">
													<i class="el-icon-warning-outline" color="#F56C6C"></i>
													<span>{{ proxy.name }}</span>
												</div>
												<!-- Actions -->
												<el-button-group class="proxies-header-btns">
													<!-- Edit -->
													<el-button size="mini" icon="el-icon-edit" @click="handleCollapse(index)"></el-button>
													<!-- Move -->
													<el-button size="mini" icon="el-icon-rank" class="proxies-handle"></el-button>
													<!-- Preview -->
													<el-link :disabled="proxy.path === ''" target="_blank" :href="proxy.path" :underline="false">
														<el-button style="border-radius: 0;" :disabled="proxy.path === ''" size="mini" icon="el-icon-view"></el-button>
													</el-link>
													<!-- Delete -->
													<el-popconfirm
														confirmButtonText="Yes"
														cancelButtonText="No"
														icon="el-icon-danger"
														iconColor="red"
														title="Are you sure to delete this proxy?"
														@confirm="deleteProxy(index)">
														<template #reference>
															<el-button style="border-radius: 0 4px 4px 0; border-left: 0;" size="mini" icon="el-icon-delete"></el-button>
														</template>
													</el-popconfirm>
												</el-button-group>
											</div><!-- /Header -->
										</template>
										<!-- Proxy Item -->
										<ProxyItem :item="proxy" :validating="isValidating" @validate="handleValidate($event, index)"></ProxyItem>
									</el-collapse-item>
								</draggable>
							</el-collapse><!-- /Proxies -->
						</el-form>
						<el-empty v-else-if="!doingAxios" :image-size="100">
							<h4>No proxies available</h4>
							<p>Click the button above to create a new proxy</p>
						</el-empty>
					</div>
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
import ProxyItem from "../../components/proxies/ProxyItem";

const UNASSIGNED_PREFIX = "Unassigned";

export default {
	name: "Proxies",
	title: "Proxies",
	mixins: [optionsMixin],
	components: {
		ProxyItem,
		Breadcrumbs,
		draggable,
	},
	data: () => ({
		activeCollapse: [],
		isValidating: false,
		validForm: true,
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
	methods: {
		/**
		 * Checks if the form is valid, if it
		 * is it will continue to save to
		 * the backend.
		 */
		saveProxies() {
			this.isValidating = true;
			this.validForm = true;
			setTimeout(() => {
				this.isValidating = false;
				if (!this.validForm) {
					this.$message({showClose: true, message: 'Fix the errors before saving.', type: 'error'});
					return;
				}
				this.save();
			},100);
		},
		/**
		 * Creates a new reverse proxy and adds to the
		 * proxies array.
		 * @param proxy
		 */
		createProxy(proxy) {
			this.proxies.push(proxy);
		},
		/**
		 * Deletes a reverse proxy from the array.
		 * @param index
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
				rewrite: {},
				rewrite_regex: {},
			});
			this.activeCollapse = [this.proxies.length - 1];
		},
		/**
		 * Handles the collapsing or proxy items, if the
		 * index is already expanded,
		 * it will be collapsed and visa versa.
		 * @param index
		 */
		handleCollapse(index) {
			if (this.activeCollapse.includes(index)) {
				this.activeCollapse.splice(this.activeCollapse.indexOf(index), 1);
				return;
			}
			this.activeCollapse.push(index);
		},
		/**
		 * Handle the start of a drag item, collapses
		 * all accordion items.
		 */
		handleDragStart() {
			this.activeCollapse = [];
			this.drag = true;
		},
		/**
		 * Adds an error class to the proxies when
		 * the user saves.
		 * @param isValid
		 * @param index
		 */
		handleValidate(isValid, index) {
			if (isValid) {
				this.$refs['proxyHeader'][index].classList.remove("proxies-header-error");
				return;
			}
			this.$refs['proxyHeader'][index].classList.add("proxies-header-error");
			this.validForm = false;
		},
		/**
		 * Retrieves the unassigned name for a proxy
		 * when no name is set. It will increment
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
		}
	},
	computed: {
		/**
		 * Gets and sets the proxies to the options
		 * data.
		 */
		proxies: {
			get() {
				return this.data['proxies'];
			},
			set(el) {
				this.$set(this.data, 'proxies', el);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


// Props
// =========================================================================

::v-deep {

	.el-empty__description {
		display: none !important;
	}
}


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

		&-info {
			display: flex;
			align-items: center;
		}

		i {
			display: none;
			margin-right: 6px;
			font-size: 16px;
			// TODO change to Element UI colour var
			color: #F56C6C;
		}

		&-error {

			i {
				display: block;
			}

			span {
				// TODO change to Element UI colour var
				color: #F56C6C;
			}
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
		border-top: 0;

		&:first-child {
			border-top: 1px solid rgba(#DCDFE6, 0.5);
		}

		&.sortable-chosen {

			#{$self}-header {
				background-color: #409eff1c;
				border: 1px dashed #409EFF;
			}
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
			cursor: pointer;

			#{$self}-card {
				background-color: #ecf5ff;
			}
		}
	}
}

</style>
