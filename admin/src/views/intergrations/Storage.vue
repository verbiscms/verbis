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
			<div class="row">
				<!-- =====================
					Details
					===================== -->
				<div class="col-12 col-desk-6">
					<!-- =====================
						Info
						===================== -->
					<div class="storage-config">
						<h2 class="storage-title">Configuration:</h2>
						<p>You are able to change storage providers and bucket's below. Each file is stored with it's own provider and bucket information.</p>
<!--						<el-result icon="info" title="Info Tip" subTitle="Please follow the instructions">-->
<!--							<template slot="extra">-->

<!--							</template>-->
<!--						</el-result>-->
						<!-- Active Provider -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-link @click="showProviderModal = true">Change Provider</el-link>
							</el-col>
							<el-col :span="16">
								<h4>Active Provider</h4>
								<el-tag type="success" effect="plain" size="small">{{ info['provider'] }}</el-tag>
							</el-col>
						</el-row><!-- /Active Provider -->
						<!-- Active Bucket -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-link @click="showBucketModal = true">Change Bucket</el-link>
							</el-col>
							<el-col :span="16">
								<h4>Active Bucket</h4>
								<el-tag type="success" effect="plain" size="small">{{ info['bucket'] }}</el-tag>
							</el-col>
						</el-row><!-- /Active Bucket -->
						<!-- Active Bucket -->
						<el-row class="storage-config-item" align="center">
							<el-col :span="8">
								<el-button @click="disconnect">Disconnect</el-button>
							</el-col>
							<el-col :span="16">
								<h4>Disconnect</h4>
								<p>Removes the remote storage provider and defaults to the local file system.</p>
							</el-col>
						</el-row><!-- /Active Bucket -->
					</div><!-- /Info -->
					<!-- =====================
						Migration
						===================== -->
					<el-divider></el-divider>
					<div class="storage-config">
						<!-- Heading -->
						<div class="storage-config-item">
							<h2 class="storage-title">Migrate:</h2>
						</div><!-- /Heading -->
						<!-- Migrate To Server -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-button @click="handleMigrateModal(true)" plain>Migrate to Server</el-button>
							</el-col>
							<el-col :span="16">
								<h4>Migrate to Server</h4>
								<p>Migrate all of the Verbis library to the remote provider, this will upload all files stored in the library. The provider must be connected before migrating.</p>
							</el-col>
						</el-row><!-- Migrate To Server -->
						<!-- Migrate To Local -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-button @click="handleMigrateModal(false)" plain>Migrate to Local</el-button>
							</el-col>
							<el-col :span="16">
								<h4>Migrate to Local</h4>
								<p>Migrate all of the Verbis library to the local file system, this will download all files stored in the library. The provider must be connected before migrating.</p>
							</el-col>
						</el-row><!-- Migrate To Server -->
						<!-- Progress -->
						<div v-if="config['is_migrating'] && config['migration']" class="storage-migration">
							<h4>Migration in progress</h4>
							<el-progress :text-inside="true" :stroke-width="20" :percentage="config.migration.progress"></el-progress>
							<small>Processing {{ config['migration']['files_processed'] }}/{{ config['migration']['total'] }} files</small>
						</div>
					</div><!-- /Migration -->
					<!-- =====================
						Options
						===================== -->
					<el-divider></el-divider>
					<div class="storage-config">
						<!-- Heading -->
						<div class="storage-config-item">
							<h2 class="storage-title">Options:</h2>
						</div><!-- /Heading -->

						<!-- Upload to Remote -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-switch v-model="info.upload_remote" @change="saveInfo"></el-switch>
							</el-col>
							<el-col :span="16">
								<h4>Upload to Remote</h4>
								<p>Files will automatically be uploaded to and served from the remote provider. If it is disabled, local file storage will be used for storing and serving.</p>
							</el-col>
						</el-row><!-- Upload to Remote -->
						<!-- Keep Local Backup -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-switch v-model="info.local_backup" @change="saveInfo"></el-switch>
							</el-col>
							<el-col :span="16">
								<h4>Local backup</h4>
								<p>Keeps a local backup of a file when it's uploaded if a remote provider is connected.</p>
							</el-col>
						</el-row><!-- /Keep Local Backup -->
						<!-- Keep Server Backup -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-switch v-model="info.remote_backup" @change="saveInfo"></el-switch>
							</el-col>
							<el-col :span="16">
								<h4>Remote backup</h4>
								<p>Keeps a remote backup of a file when it's uploaded if a remote provider is connected.</p>
							</el-col>
						</el-row><!-- /Keep Local Backup -->
						<!-- Download -->
						<el-row class="storage-config-item">
							<el-col :span="8">
								<el-button @click="download" :loading="downloading">Download</el-button>
							</el-col>
							<el-col :span="16">
								<h4>Download Library</h4>
								<p>By clicking download you are able to download the entire media library to your local machine. This will not include any items that are not stored in Verbis.</p>
								<el-alert v-if="downloading" title="Downloading, do not refresh the page" type="warning" show-icon></el-alert>
							</el-col>
						</el-row><!-- /Download -->
					</div><!-- /Options -->
				</div>
				<!-- =====================
					Providers
					===================== -->
				<div class="col-12 col-desk-5 offset-desk-1">
					<h2>Providers:</h2>
					<el-collapse v-model="activeProviders">
						<el-collapse-item class="storage-provider" v-for="(provider, key) in config.providers" :key="key" :name="provider.name">
							<!-- Header -->
							<template #title>
								<el-image class="storage-provider-image" :src="require('@/assets/images/' + getLogo(key))" fit="contain"></el-image>
								<h4>{{ provider['name'] }}</h4>
							</template><!-- /Header -->
							<!-- Body -->
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
							<div v-if="provider.connected">
								<el-tag	type="success">Connected</el-tag>
							</div>
						</el-collapse-item><!-- /Body -->
					</el-collapse>
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
		<!-- =====================
			Provider Modal
			===================== -->
		<el-dialog :visible.sync="showProviderModal" width="30%">
			<!-- Header -->
			<template #title>
				<h2 class="mb-0">Provider</h2>
				<p>Select a provider below to change the default storage offloading for Verbis.</p>
			</template>
			<!-- Form -->
			<el-form ref="provider-form" label-width="70px" label-position="left">
				<!-- Provider -->
				<el-form-item label="Provider">
					<el-select v-model="info.provider" placeholder="Select a provider" style="width: 100%;">
						<el-option v-for="(provider, index) in config['providers']" :key="index" :label="provider.name" :value="index" :disabled="!provider.connected"></el-option>
					</el-select>
				</el-form-item>
			</el-form>
			<!-- Footer -->
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="showProviderModal = false">Cancel</el-button>
					<el-button type="primary" @click="changeProvider">Change Provider</el-button>
				</span>
			</template>
		</el-dialog>
		<!-- =====================
			Bucket Modal
			===================== -->
		<el-dialog :visible.sync="showBucketModal" width="30%">
			<!-- Header -->
			<template #title>
				<h2 class="mb-0">Bucket</h2>
				<p>Select a bucket below to change the default storage offloading for Verbis.</p>
			</template>
			<!-- Form -->
			<el-form ref="form" label-width="120px" label-position="left">
				<!-- Bucket -->
				<el-form-item label="Bucket">
					<el-select v-model="info.bucket" placeholder="Select a bucket" no-match-text="No buckets found" style="width: 100%;">
						<el-option v-for="(bucket, index) in buckets" :key="index" :label="bucket.name" :value="bucket.id"></el-option>
					</el-select>
				</el-form-item>
				<!-- Create Bucket -->
<!--				<el-form-item label="Create Bucket">-->
<!--					&lt;!&ndash; TODO - Update to spacer Vue 3 &ndash;&gt;-->
<!--					<div class="d-flex">-->
<!--						<el-input v-model="bucket.name"></el-input>-->
<!--						<el-button class="ml-1" @click="createBucket()" :loading="bucket.loading">Create</el-button>-->
<!--					</div>-->
<!--				</el-form-item>-->
			</el-form>
			<!-- Footer -->
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="showBucketModal = false">Cancel</el-button>
					<el-button type="primary" @click="saveInfo()">Change Bucket</el-button>
				</span>
			</template>
		</el-dialog>
		<!-- =====================
			Migrate Modal
			===================== -->
		<el-dialog :visible.sync="showMigrateModal" width="30%">
			<!-- Header -->
			<template #title>
				<h2 v-if="migrate['isRemote']" class="mb-0">Migrate to {{ config.providers[info.provider].name }}</h2>
				<h2 v-else class="mb-0">Migrate to Local Storage</h2>
				<p class="t-left">Select a remote provider and bucket below to migrate too.</p>
			</template>
			<!-- Form -->
			<el-form ref="migrate-form" :v-model="migrate" label-width="140px" label-position="left">
				<!-- Delete Files -->
				<el-form-item label="Delete originals?">
					<el-switch v-model="migrate['delete']"></el-switch>
				</el-form-item>
			</el-form>
			<!-- Footer -->
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="showMigrateModal = false">Cancel</el-button>
					<el-button type="primary" @click="doMigrate()">Migrate</el-button>
				</span>
			</template>
		</el-dialog>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";

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
	},
	data: () => ({
		doingAxios: true,
		text: "",
		config: {},
		showProviderModal: false,
		showMigrateModal: false,
		showBucketModal: false,
		providerBtnLoading: false,
		info: {},
		buckets: [],
		activeProviders: [],
		migrate: {
			delete: false,
			isRemote: false,
		},
		bucket: {
			name: "",
			loading: false,
		},
		saving: false,
		backup: false,
		downloading: false,
	}),
	mounted() {
		this.getConfig();
	},
	watch: {
		showMigrateModal: function() {
			this.migrate.provider = "";
			this.migrate.bucket = "";
			this.migrate.validProvider = false;
		},
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
					this.doingAxios = false;
					this.info = this.config.info;
					this.listBuckets(this.info['provider']).then(res => this.buckets = res);


					// this.config.is_migrating = true;
					// this.config.migration = {
					// 	total: 400,
					// 	progress: 10,
					// 	succeeded: 10,
					// 	failed: 10,
					// 	files_processed: 10,
					// }
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
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
			return this.axios.get("/storage/bucket/" + provider).then(res => this.buckets = res.data.data);
		},
		/*
		 * changeMigrateModal()
		 * Handles processing when the select box is
		 * changed in the migration modal.
		 */
		async changeMigrateModal() {
			this.buckets = [];
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
		 * doMigrate()
		 * Posts to the backend and runs the migration.
		 */
		doMigrate() {
			let migration = {
				to_server: this.migrate['isRemote'],
				delete: this.migrate['delete'],
			}
			this.axios.post("/storage/migrate", migration)
				.then(res => {
					this.$message({message: res.data.message, type: "success"});
					this.showMigrateModal = false;
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.showMigrateModal = false;
						this.$message.error(err.response.data.message);
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
					this.$message({message: "Successfully created bucket: " + postData.bucket, type: "success"})
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
			this.saving = true;
			this.axios.post("/storage", this.info)
				.then(res => {
					this.showProviderModal = false;
					this.showBucketModal = false;
					this.$message({message: res.data.message, type: "success"});
					this.getConfig();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.saving = false;
				})
		},
		// THis needs to be in API.
		disconnect() {
			this.info.provider = ""
			this.info.bucket = ""
			this.saveInfo();
		},
		/**
		 * Download's the entire media library and creates
		 * a new blob to download.
		 */
		download() {
			this.downloading = true;
			this.axios.get("/storage/download",
				{responseType: 'blob'}
			).then(res => {
				console.log(res);
				let fileName = res.headers["x-filename"],
					fileUrl = window.URL.createObjectURL(new Blob([res.data], {type: "application/zip"})),
					fileLink = document.createElement("a");
				fileLink.href = fileUrl;

				if (!fileName) {
					fileName = "verbis-library.zip"
				}

				fileLink.setAttribute("download", fileName);
				document.body.appendChild(fileLink);
				fileLink.click();
			})
			.catch(err => {
				this.helpers.handleResponse(err);
			})
			.finally(() => {
				this.downloading = false;
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

		// Props
		// =========================================================================

		&-title {
			margin-bottom: 4px;
		}

		// Config
		// =========================================================================

		&-config {
			margin-bottom: 2rem;

			&-item {
				margin-bottom: 1.6rem;

				p:last-child {
					margin-bottom: 0;
				}
			}
		}

		// Provider
		// =========================================================================

		&-provider {

			&-image {
				width: 50px;
				height: 50px;
				margin-right: 10px;
			}

			::v-deep {

				.el-collapse-item__header {
					padding: 1rem;
					height: 70px;
				}

				.el-collapse-item__content {
					padding: 0 1rem 1rem 1rem;
				}
			}
		}

		// Migration
		// =========================================================================

		&-migration {
			margin-top: 1.7rem;

			h4 {
				margin-bottom: 8px;
			}

			small {
				margin-top: 8px;
				display: block;
				text-align: right;
			}
		}
	}


</style>
