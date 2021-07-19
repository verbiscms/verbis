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
					Details
					===================== -->
				<div class="col-12 col-desk-5">
					<h2 class="storage-config-title">Configuration:</h2>
					<p>You are able to change storage providers and bucket's below. Each file is stored with it's own provider and bucket information.</p>
					<div class="storage-config">
						<h4>Active Provider</h4><p>{{  config['active_provider'] }}</p>
						<h4>Active Bucket</h4><p>{{ config['active_bucket'] }}</p>
						<h4>Change files provider</h4>
						<p>To cha</p>
					</div>
					<div class="btn-cont">
						<button class="btn" @click="showProviderModal = true">Change Provider</button>
						<button class="btn" v-if="config['active_provider'] !== 'local'" @click="showBucketModal = true">Change Bucket</button>
					</div>
					<div class="btn-cont">
						<button class="btn" @click="doMigrate()">Migrate to Server THIS NEEDS TO BE A MODAL</button>
						<button class="btn" @click="doMigrateToLocal()">Migrate to Local</button>
					</div>
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
				<!-- Role -->
				<FormGroup label="Provider*" :error="errors['provider_modal']">
					<div class="form-select-cont form-input">
						<select class="form-select" id="user-role" v-model="info['provider']">
							<option v-for="(provider, providerIndex) in config['providers']" :value="providerIndex" :key="providerIndex">{{ provider['name'] }}</option>
						</select>
					</div>
				</FormGroup>
			</template>
		</Modal>
		<!-- =====================
			Bucket Modal
			===================== -->
		<Modal :show.sync="showBucketModal">
			<template slot="button">
				<button class="btn" @click="saveInfo">Change Bucket</button>
			</template>
			<template slot="text">
				<div class="text-cont">
					<h2>Bucket</h2>
					<p class="t-left">Select a provider below to change the default storage offloading for Verbis.</p>
				</div>
				<!-- Role -->
				<FormGroup label="Bucket*" :error="errors['bucket_modal']">
					<div class="form-select-cont form-input">
						<select class="form-select" v-model="info['bucket']">
							<option value="" disabled selected>Select a Bucket</option>
							<option v-for="(bucket, bucketIndex) in buckets" :value="bucket['id']" :key="bucketIndex">{{ bucket['name'] }}</option>
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
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import Modal from "@/components/modals/General";

const GCPLogo = "google-storage.png",
	AWSLogo = "amazon-storage.png",
	AzureLogo = "azure-storage.png";

export default {
	name: "Storage",
	title: "Storage",
	components: {
		Breadcrumbs,
		FormGroup,
		Collapse,
		Modal,
	},
	data: () => ({
		doingAxios: true,
		text: "",
		config: {},
		showProviderModal: false,
		showBucketModal: false,
		providerBtnLoading: false,
		bucketBtnLoading: false,
		info: {},
		buckets: [],
		errors: [],
	}),
	mounted() {
		this.getConfig();
	},
	methods: {
		getConfig() {
			this.axios.get("/storage/config")
				.then(res => {
					this.config = res.data.data;
					this.info = {
						"provider": this.config['active_provider'],
						"bucket": this.config['active_bucket']
					};
					this.doingAxios = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		filteredProviders() {
			const providers =  Object.assign({}, this.config['providers']);
			delete providers["local"];
			return providers;
		},
		changeProvider() {
			const provider = this.info['provider'];
			if (provider === "local") {
				this.info['bucket'] = "";
				this.saveInfo();
				return;
			}

			const cfg = this.config['providers'][provider];
			if (cfg['error']) {
				this.$set(this.errors, 'provider_modal', cfg['error']);
				return;
			}

			this.showProviderModal = false;
			this.listBuckets(this.info['provider']).then(res => this.buckets = res);
			setTimeout(() => {
				this.showBucketModal = true;
			}, 100);
		},
		listBuckets(provider) {
			return this.axios.get("/storage/bucket/" + provider).then(res => res.data.data)
		},
		doMigrate() {
			const migration = {
				to: {
					provider: "aws",
					bucket: "reddicotest",
				},
				from: {
					provider: "local",
				}
			}

			this.axios.post("/storage/migrate", migration)
				.then(res => {
					console.log(res);
				})

		},
		doMigrateToLocal() {
			const migration = {
				from: {
					provider: "aws",
					bucket: "reddicotest",
				},
				to: {
					provider: "local",
				}
			}

			this.axios.post("/storage/migrate", migration)
				.then(res => {
					console.log(res);
				})

		},
		// deleteBucket(bucket, provider) {
		// 	this.axios.delete("/storage/bucket")
		// },
		saveInfo() {
			this.axios.post("/storage", this.info)
				.then(res => {
					console.log(res);
					this.showProviderModal = false;
					this.showBucketModal = false;
					this.$noty.success(res.data.message);
					this.getConfig();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		getLogo(name) {
			switch (name) {
				case "google":
					return GCPLogo;
				case "aws":
					return AWSLogo;
				case "azure":
					return AzureLogo;
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

	.storage {

		&-cont {
			display: flex;
			align-items: center;
		}

		&-config {
			margin-top: 2rem;

			&-title {
				margin-bottom: 4px;
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
	img {
		max-width: 100%;
	}

</style>
