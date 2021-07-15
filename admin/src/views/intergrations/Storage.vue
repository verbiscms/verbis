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
				<div class="col-12 col-desk-5">
					<h2 class="storage-config-title">Configuration:</h2>
					<p>You are able to change storage providers and bucket's below. Each file is stored with it's own provider and bucket information.</p>
					<div class="storage-config">
						<h4>Active Provider</h4><p>Local</p>
						<h4>Active Bucket</h4><p>Local</p>
						<h4>Change files provider</h4>
						<p>To cha</p>
					</div>

					<div class="btn-cont">
						<button class="btn">Change Provider</button>
						<button class="btn">Change Bucket</button>
					</div>
				</div>
				<!-- =====================
					Providers
					===================== -->
				<div class="col-12 col-desk-6 offset-desk-1">
					<div class="card card-small-box-shadow card-expand">
						<!-- Local -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<h4>Local Storage</h4>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<p>Lorem ipsum</p>
									<code v-html="'{{ verbisHead }}'"></code>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Title & description -->
						<!-- GCP -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div class="storage-cont">
										<figure class="storage-image">
											<img src="@/assets/images/google-storage.png">
										</figure>
										<h4>Google Cloud</h4>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<p>Lorem ipsum</p>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /GCP -->
						<!-- Amazon -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div class="storage-cont">
										<figure class="storage-image">
											<img src="@/assets/images/amazon-storage.png">
										</figure>
										<h4>Amazon S3</h4>

									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<code v-html="'{{ verbisHead }}'"></code>
									<p>Lorem ipsum</p>
								</div><!-- / Body -->
							</template>
						</Collapse><!-- /Amazon -->
						<!-- Azure -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div class="storage-cont">
										<figure class="storage-image">
											<img src="@/assets/images/azure-storage.png">
										</figure>
										<h4>Microsoft Azure</h4>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<p>Lorem ipsum</p>
								</div><!-- / Body -->
							</template>
						</Collapse><!-- /Azure -->
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
import Collapse from "@/components/misc/Collapse";

export default {
	name: "Storage",
	title: "Storage",
	components: {
		Breadcrumbs,
		Collapse,
	},
	data: () => ({
		text: "",
	}),
	mounted() {
		this.getBuckets();
	},
	methods: {
		getBuckets() {
			this.axios.get("/storage/config").then(res => {
				console.log(res);
			})
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

		&-image {
			height: 50px;
			width: 80px;
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
