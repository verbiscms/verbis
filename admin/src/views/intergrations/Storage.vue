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
							<h1>Storage</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-if="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-else class="row">
				<!-- =====================
					Progress
					===================== -->
				<div class="col-12" v-if="config['is_migrating']">
					<Progress :percent="config['migration']['progress']">
						<slot>
							<div class="storage-progress">
								<h4>Migration in progress</h4>
								<p>Processing {{ config['migration']['files_processed'] }}/{{ config['migration']['total'] }}</p>
							</div>
						</slot>
					</Progress>
				</div>
				<!-- =====================
					Details
					===================== -->
				<div class="col-12 col-desk-5">
					<!-- Configuration -->
					<div class="storage-config">
						<h2 class="storage-title">Configuration:</h2>
						<p>You are able to change storage providers and bucket's below. Each file is stored with it's own provider and bucket information.</p>
						<!-- Provider Info -->
						<div class="storage-config-cont">
							<div>
								<h4>Active Provider</h4>
								<p>{{ config['active_provider'] }}</p>
							</div>
							<span class="link" @click="showProviderModal = true">Change Provider</span>
						</div><!-- /Provider Info -->
						<!-- Bucket Info -->
						<div class="storage-config-cont"  v-if="config['active_provider'] !== 'local'">
							<div>
								<h4>Active Bucket</h4>
								<p>{{ config['active_bucket'] }}</p>
							</div>
							<span class="link" @click="showBucketModal = true">Change Bucket</span>
						</div><!-- /Bucket Info -->
					</div><!-- /Configuration -->
					<!-- Migration -->
					<div class="storage-config">
						<h2 class="storage-title">Migrate:</h2>
						<p>You can migrate all of the Verbis file library to a remote cloud provider or the local file system with the buttons below. The provider must be connected before migrating.</p>
						<div class="storage-config-btn-cont">
							<button class="btn" @click="handleMigrateModal(true)">Migrate to Server</button>
							<button class="btn" @click="handleMigrateModal(false)">Migrate to Local</button>
						</div>
					</div><!-- /Migration -->
				</div>
				<!-- =====================
					Providers
					===================== -->
				<div class="col-12 col-desk-6 offset-desk-1">
					<h2>Providers:</h2>
					<div class="card card-small-box-shadow card-expand">
						<!-- Local -->
						<Collapse v-for="(provider, key) in filteredProviders()" :key="key" :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div class="storage-cont">
										<figure class="storage-image">
											<img :src="require('@/assets/images/' + getLogo(key))">
										</figure>
										<h4>{{ provider['name'] }}</h4>
									</div>
									<div v-if="provider['connected']">
										<div class="badge badge-green">Connected</div>
									</div>
									<div v-else class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body v-if="!provider['connected']">
								<div class="card-body">
									<div v-if="!provider['environment_set']">
										<p>In order to use {{ provider.name }} as a storage provider please add the following keys to the <code>.env</code> file or use <code>export VARIABLE=value</code> and proceed to restart Verbis.</p>
										<div v-for="(envKey, index) in provider['environment_keys']" :key="index">
											<code>{{ envKey }}</code>
										</div>
									</div>
									<div v-if="provider['error'] && provider['environment_set']">
										<p>There is an error connecting to {{ provider['name'] }}</p>
										<p class="storage-error">{{ provider['error'] }}</p>
									</div>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Title & description -->
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Migrate Modal
			===================== -->
		<Modal :show.sync="showMigrateModal">
			<template slot="button">
				<button class="btn" @click="doMigrate()">Migrate</button>
			</template>
			<template slot="text">
				<div class="text-cont">
					<h2 v-if="this.migrate['isRemote']">Migrate to Remote Provider</h2>
					<h2 v-else>Migrate to Local Storage</h2>
					<p class="t-left">Select a remote provider and bucket below to migrate too.</p>
				</div>
				<!-- Provider -->
				<FormGroup label="Provider*" :error="errors['migrate_modal_provider']">
					<div class="form-select-cont form-input">
						<select class="form-select" id="user-role" v-model="migrate['provider']" @change="changeMigrateModal()">
							<option value="" disabled selected>Select a Provider</option>
							<option v-for="(provider, providerIndex) in filteredProviders()" :value="providerIndex" :key="providerIndex">{{ provider['name'] }}</option>
						</select>
					</div>
				</FormGroup><!-- /Provider -->
				<!-- Bucket -->
				<FormGroup v-if="migrate['validProvider']" label="Bucket*" :error="errors['migrate_modal_bucket']">
					<div class="form-select-cont form-input">
						<select class="form-select" v-model="migrate['bucket']">
							<option value="" disabled selected>Select a Bucket</option>
							<option v-for="(bucket, bucketIndex) in buckets" :value="bucket['id']" :key="bucketIndex">{{ bucket['name'] }}</option>
						</select>
					</div>
				</FormGroup><!-- /Bucket -->
				<!-- Delete -->
				<div class="migrate-delete">
					<h6 class="margin">Delete original files?</h6>
					<div class="toggle">
						<input type="checkbox" class="toggle-switch" id="migration-delete" v-model="migrate['delete']" :true-value="true" :false-value="false" />
						<label for="migration-delete"></label>
					</div>
				</div><!-- /Delete -->
			</template>
		</Modal>
		<!-- =====================
			Provider Modal
			===================== -->
		<Modal :show.sync="showProviderModal">
			<template slot="button">
				<button class="btn" @click="changeProvider()">Change Provider</button>
			</template>
			<template slot="text">
				<div class="text-cont">
					<h2>Provider</h2>
					<p class="t-left">Select a provider below to change the default storage offloading for Verbis.</p>
				</div>
				<!-- Provider -->
				<FormGroup label="Provider*" :error="errors['provider_modal']">
					<div class="form-select-cont form-input">
						<select class="form-select" v-model="info['provider']" @change="errors = []">
							<option value="" disabled selected>Select a Provider</option>
							<option v-for="(provider, providerIndex) in config['providers']" :value="providerIndex" :key="providerIndex">{{ provider['name'] }}</option>
						</select>
					</div>
				</FormGroup><!-- /Provider -->
			</template>
		</Modal>
		<!-- =====================
			Bucket Modal
			===================== -->
		<Modal :show.sync="showBucketModal">
			<template slot="button">
				<button class="btn" @click="saveInfo()">Change Bucket</button>
			</template>
			<template slot="text">
				<div class="text-cont">
					<h2>Bucket</h2>
					<p class="t-left">Select a provider below to change the default storage offloading for Verbis.</p>
				</div>
				<!-- Select Bucket -->
				<FormGroup label="Bucket*" :error="errors['bucket_modal']">
					<div class="form-select-cont form-input">
						<select class="form-select" v-model="info['bucket']">
							<option value="" disabled selected>Select a Bucket</option>
							<option v-for="(bucket, bucketIndex) in buckets" :value="bucket['id']" :key="bucketIndex">{{ bucket['name'] }}</option>
						</select>
					</div>
				</FormGroup><!-- /Select Bucket -->
				<FormGroup class="bucket" label="Create Bucket" :error="errors['create_bucket']">
					<div class="bucket-cont">
						<input class="form-input form-input-white" type="text" v-model="bucket['name']">
						<button class="btn" @click="createBucket()" :class="{ 'btn-loading' : bucket['loading'] }">Create</button>
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
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import Modal from "@/components/modals/General";
import Progress from "../../components/misc/Progress";

const logos = {
	google: "google-storage.png",
	aws: "amazon-storage.png",
	azure:  "azure-storage.png"
}

export default {
	name: "Storage",
	title: "Storage",
	components: {
		Breadcrumbs,
		FormGroup,
		Collapse,
		Modal,
		Progress,
	},
	data: () => ({
		doingAxios: true,
		text: "",
		config: {},
		showProviderModal: false,
		showMigrateModal: false,
		showBucketModal: false,
		providerBtnLoading: false,
		info: {
			provider: "",
			bucket: "",
		},
		buckets: [],
		errors: [],
		migrate: {
			provider: "",
			bucket: "",
			delete: false,
			validProvider: false,
			isRemote: false,
		},
		bucket: {
			name: "",
			loading: false,
		}
	}),
	mounted() {
		this.getConfig();
		//setInterval(this.getConfig, 3000);
	},
	watch: {
		showProviderModal: function () {
			this.errors = [];
		},
		showMigrateModal: function() {
			this.migrate.provider = "";
			this.migrate.bucket = "";
			this.migrate.validProvider = false;
		}
	},
	methods: {
		/*
		 * getConfig()
		 * Returns the configuration for the storage
		 * service.
		 */
		getConfig() {
			this.axios.get("/storage/config")
				.then(res => {
					this.config = res.data.data;
					// this.info = {
					// 	"provider": this.config['active_provider'],
					// 	"bucket": this.config['active_bucket']
					// };
					this.doingAxios = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * filteredProviders()
		 * Returns providers without local.
		 */
		filteredProviders() {
			const providers =  Object.assign({}, this.config['providers']);
			delete providers["local"];
			return providers;
		},
		/*
		 * changeProvider()
		 * Updates the storage provider. Checks for errors in the
		 * config if the provider is not connected. If the
		 * provider is local, info will be saved
		 * straight away.
		 */
		async changeProvider() {
			const provider = this.info['provider'];
			if (provider === "local") {
				this.info['bucket'] = "";
				this.saveInfo();
				return;
			}

			const err = this.getProviderError(provider)
			if (err) {
				this.$set(this.errors, 'provider_modal', err);
				return;
			}

			this.showProviderModal = false;

			await this.listBuckets(this.info['provider']).then(res => this.buckets = res);

			setTimeout(() => {
				this.showBucketModal = true;
			}, 100);
		},
		/*
		 * listBuckets()
		 * Sets the buckets with the given provider.
		 */
		listBuckets(provider) {
			return this.axios.get("/storage/bucket/" + provider).then(res => this.buckets = res.data.data)
		},
		/*
		 * changeMigrateModal()
		 * Handles processing when the select box is
		 * changed in the migration modal.
		 */
		async changeMigrateModal() {
			this.migrate['validProvider'] = false;
			this.$delete(this.errors, 'migrate_modal_provider');

			const err = this.getProviderError(this.migrate['provider'])
			if (err) {
				this.$set(this.errors, 'migrate_modal_provider', err);
				return;
			}

			this.migrate['validProvider'] = true;
			await this.listBuckets(this.migrate['provider']);
		},
		/*
		 * handleMigrateModal()
		 * Updates the migration data and shows the
		 * migration modal.
		 */
		handleMigrateModal(toRemote) {
			this.$set(this.migrate, 'isRemote', toRemote);
			this.showMigrateModal = true;
		},
		/*
		 * getProviderError()
		 * Determines if the provider is connected and
		 * has an error.
		 */
		getProviderError(provider) {
			const cfg = this.config['providers'][provider];
			if (cfg['error']) {
				return cfg['error'];
			}
			return false;
		},
		/*
		 * doMigrate()
		 * Posts to the backend and runs the migration.
		 */
		doMigrate() {
			let migration =  {
				from: {
					provider: this.migrate['provider'],
					bucket: this.migrate['bucket'],
				},
				to: {
					provider: "local"
				},
			};
			if (this.migrate['isRemote']) {
				migration = {
					from: {
						provider: "local"
					},
					to: {
						provider: this.migrate['provider'],
						bucket: this.migrate['bucket'],
					},
				}
			}
			migration['delete'] = this.migrate['delete'];
			this.axios.post("/storage/migrate", migration)
				.then(res => {
					this.$noty.success(res.data.message);
					this.showMigrateModal = false;
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.$noty.error(err.response.data.message);
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.getConfig();
				})
		},
		/*
		 * createBucket()
		 * Creates a new bucket and pushes to the bucket
		 * array on successful post.
		 */
		createBucket() {
			this.$set(this.bucket, "loading", true)

			// Validation check
			if (this.bucket['name'] === "") {
				this.$set(this.errors, 'create_bucket', "Enter a bucket name");
				setTimeout(() => {
					this.$set(this.bucket, "loading", false)
				}, this.timeoutDelay);
				return;
			}

			// Remove errors
			this.$delete(this.errors, 'create_bucket');

			// Post to backend
			let postData = {
				provider: this.info['provider'],
				bucket: this.bucket.name,
			};

			this.axios.post("/storage/bucket", postData)
				.then(res => {
					this.buckets.push(res.data.data);
					this.$noty.success("Successfully created bucket: " + postData.bucket)
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.$set(this.errors, 'create_bucket', err.response.data.message);
						return
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.$set(this.bucket, "loading", false)
					}, this.timeoutDelay);
				})
		},
		/*
		 * saveInfo()
		 * Changes the storage provider (info).
		 */
		saveInfo() {
			this.axios.post("/storage", this.info)
				.then(res => {
					this.showProviderModal = false;
					this.showBucketModal = false;
					this.$noty.success(res.data.message);
					this.getConfig();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * getLogo()
		 * Returns the relevant provider logo.
		 */
		getLogo(name) {
			return logos[name];
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Storage
	// =========================================================================

	.storage {

		&-config-cont {
			display: flex;
			align-items: center;
			justify-content: space-between;
			margin-bottom: 10px;

			.link {
				text-decoration: underline;
				font-size: 14px;
				cursor: pointer;
				color: $secondary;

				&:hover {
					color: $primary;
				}
			}
		}

		&-config {
			margin-bottom: 2rem;

			&-btn-cont {
				.btn:first-child {
					margin-right: 10px;
				}
			}
		}

		&-title {
			margin-bottom: 4px;
		}

		&-progress {
			display: flex;
			align-items: center;
			justify-content: space-between;
			margin-bottom: 2rem;

			p {
				margin-bottom: 0;
			}
		}

		&-error {
			color: $orange;
			font-weight: 500;
			margin-bottom: 0;
		}

		&-image {
			height: 50px;
			width: 60px;
			margin-right: 1rem;

			img {
				width: 100%;
				height: 100%;
				object-fit: contain;
			}
		}
	}

	// Bucket
	// =========================================================================

	.bucket {

		&-cont {
			display: flex;
			align-items: center;
		}

		.btn {
			margin-left: 10px;
			height: 50px;
		}
	}

	// Migrate
	// =========================================================================

	.migrate {

		&-delete {
			display: flex;
			justify-content: space-between;
		}
	}


</style>
