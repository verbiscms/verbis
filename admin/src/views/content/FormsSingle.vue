<!-- =====================
	Forms - Single
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<!-- Header -->
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions">
						<div class="header-title">
							<h1>{{ form['name'] }} Submissions</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div class="row trans-fade-down-anim" v-if="!doingAxios" >
				<div class="col-12">
					<div class="table-wrapper" v-if="form['submissions'] && form['submissions'].length">
						<div class="table-scroll table-with-hover">
							<table class="table archive-table">
								<thead>
								<tr>
									<th v-for="(field, fieldIndex) in form['fields']" :key="fieldIndex">
										<span>{{ field['label'] }}</span>
									</th>
								</tr>
								</thead>
								<tbody>
								<tr v-for="(submission, submissionIndex) in form['submissions']" :key="submissionIndex">
									<td v-for="(field, fieldIndex) in orderFields(submission['fields'])" :key="fieldIndex">
										{{ field }}
									</td>
								</tr>
								</tbody>
							</table>
						</div><!-- /Table Scroll -->
					</div><!-- /Table Wrapper -->
				</div>
			</div>
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";

export default {
	name: "Forms",
	components: {
		Breadcrumbs,
	},
	data: () => ({
		doingAxios: true,
		form: {},
	}),
	mounted() {
		this.getFormById(this.$route.params.id);
	},
	methods: {
		/*
		 * getFormById()
		 */
		getFormById(id) {
			this.axios.get('/forms/' + id)
				.then(res => {
					const form = res.data.data;
					// Return 404 if there is no ID
					if (!form) {
						this.$router.push({ name : 'not-found' })
					}
					this.form = form;
					this.doingAxios = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		orderFields(fields) {
			let obj = {};
			this.form['fields'].forEach(f => {
				const key = f.key;
				// eslint-disable-next-line no-prototype-builtins
				if (fields.hasOwnProperty(key)) {
					obj[key] = fields[key];
					console.log(fields[key]);
				}
			});
			return obj;
		}
	},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.table {
		background-color: $white;
		border-radius: 6px;

		tbody {

			td {
				max-width: 600px;
				white-space: pre-wrap;
			}
		}
	}

	// URL
	// =========================================================================

	@include media-desk {
		.form-url {
			width: 50%;

			input {
				width: 100%;
			}
		}
	}

</style>