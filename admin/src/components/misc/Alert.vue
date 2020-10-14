<!-- =====================
	Alert
	===================== -->
<template>
	<transition name="trans-fade-quick">
		<div v-if="show" class="alert alert-background" :class="'alert-' + colour">
			<!-- Icon -->
			<div class="alert-icon">
				<i class="feather feather-alert-circle"></i>
			</div><!-- /Icon -->
			<!-- Text -->
			<span>
				<slot></slot>
			</span><!-- /Text -->
			<!-- Close -->
			<button type="button" class="alert-close" aria-label="Close" @click="show = false">
				<i class="feather feather-x"></i>
			</button><!-- /Close -->
		</div>
	</transition>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Alert",
	props: {
		colour: {
			type: String,
			default: ""
		}
	},
	data: () => ({
		show: true,
	}),
	beforeMount() {
		this.updateList()
	},
	watch: {
		'$route'() {
			this.updateList()
		}
	},
	methods: {
		/*
		 * updateList()
		 *
		 */
		close() {

		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.alert {
	$self: &;

	border: 0;
	position: relative;
	display: flex;
	align-items: center;
	padding: 26px 15px;
	margin-bottom: 20px;
	font-size: 14px;
	font-weight: 400;
	border-radius: 4px;
	background: $white;
	border-left: 3px solid $grey-light;
	color: $secondary;
	will-change: opacity;

	// Icon
	// ==========================================================================

	&-container {
		position: relative;
	}

	&-close {
		padding-right: 2.8rem;
	}


	// Icon
	// ==========================================================================

	&-icon {
		display: flex;
		align-items: center;
		margin-right: 10px;

		i {
			font-size: 20px;
		}
	}

	// Close
	// ==========================================================================

	&-close {
		position: absolute;
		top: 0;
		right: 0;
		height: 100%;
		font-size: 20px;
		color: inherit;
		border: 0;
		background-color: transparent;
		-webkit-appearance: none; // sass-lint:disable-line no-vendor-prefixes
		padding: 0.75rem 1.25rem;
		outline: none;
		cursor: pointer;

		span {
			line-height: 1;
		}
	}

	// Colours
	// ==========================================================================

	&-orange {
		background-color: $orange;
		color: $white;
		box-shadow: 0 12px 20px -10px rgba($orange, 0.28), 0 4px 20px 0 rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba($orange, 0.2);
		border-color: darken($orange, 10%);

		* {
			color: $white;
		}
	}

	&-green {
		background-color: $green;
		color: $white;
		box-shadow: 0 12px 20px -10px rgba($green, 0.28), 0 4px 20px 0 rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba($green, 0.2);
		border-color: darken($green, 10%);

		* {
			color: $white;
		}
	}
}

</style>