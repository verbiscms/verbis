<!-- =====================
	Media
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>Media</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<!-- Actions -->
						<div class="header-actions">
							<form class="form form-actions">
								<button v-if="bulkAction" class="btn btn-fixed-height btn-white" @click.prevent="bulkAction = false">Cancel</button>
								<button v-if="!bulkAction" class="btn btn-fixed-height btn-margin btn-white" @click.prevent="bulkAction = true">Bulk select</button>
								<button v-else class="btn btn-fixed-height btn-margin btn-orange" @click.prevent="deleting = !deleting">Delete permanently</button>
								<label for="browse-file" class="btn btn-icon btn-orange btn-icon-mob media-add-new">
									<i class="fal fa-plus"></i>
									<p>Add new media</p>
								</label>
							</form>
						</div><!-- /Actions -->
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<div class="row">
				<div class="col-12">
					<Uploader :filters="true" :modal="false" :bulk-action.sync="bulkAction" :deleting="deleting"></Uploader>
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
import Uploader from "@/components/media/Uploader";

export default {
	name: "Pages",
	title: "Archive",
	components: {
		Breadcrumbs,
		Uploader
	},
	data: () => ({
		bulkAction: false,
		deleting: false,
	}),
	computed: {
		/*
		 * checkedAll()
		 * Update the checked array to everything/nothing when checked all is clicked.
		 */
		checkedAll: {
			get() {
				return false;
			},
			set(value) {
				if (value) {
					this.checked = this.posts.map(m => {
						return m.post.id
					});
					return;
				}
				this.checked = [];
			}
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->

<style scoped lang="scss">

.media {

	// Tablet Down
	// =========================================================================

	@include media-mob-down {

		&-add-new {
			width: auto;

			i {
				display: none;
			}

			p {
				color: $white;
				margin: 0;
				font-weight: 600;
				font-size: 0.9rem;
				padding: $btn-padding-mob;
			}
		}
	}

	// Desktop
	// =========================================================================

	@include media-tab {

		&-add-new {

			i {
				display: block;
			}

			p {
				display: none;
			}
		}
	}
}

</style>