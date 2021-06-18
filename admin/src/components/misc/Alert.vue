<!-- =====================
	Alert
	===================== -->
<template>
	<transition name="trans-fade-quick" mode="out-in">
		<div v-if="show" class="alert alert-background" :class="'alert-' + colour">
			<!-- Icon -->
			<div class="alert-icon">
				<i v-if="type === 'error'" class="feather feather-alert-triangle"></i>
				<i v-if="type === 'warning'" class="feather feather-alert-circle"></i>
				<i v-if="type === 'success'" class="feather feather-check-circle"></i>
			</div><!-- /Icon -->
			<!-- Text -->
			<div class="alert-text">
				<slot></slot>
			</div><!-- /Text -->
			<!-- Close -->
			<button v-if="cross" type="button" class="alert-close" aria-label="Close" @click="show = false">
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
			default: "",
		},
		type: {
			type: String,
			default: "error",
		},
		cross: {
			type: Boolean,
			default: true,
		}
	},
	data: () => ({
		show: true,
	})
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

.alert {
	$self: &;

	position: relative;
	display: flex;
	align-items: center;
	padding: 18px 20px;
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
		margin-right: 18px;

		i {
			font-size: 24px;
		}
	}

	// Text
	// ==========================================================================

	&-text {
		flex-grow: 2;

		p {
			margin-bottom: 0;
		}

		h3 {
			margin-bottom: 4px;
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
		border-left-color: $orange;

		#{$self}-icon,
		#{$self}-text h3 {
			color: $orange;
		}
	}

	&-green {
		border-left-color: $green;

		#{$self}-icon,
		#{$self}-text h3  {
			color: $green;
		}
	}

	// Tablet
	// ==========================================================================

	@include media-tab {
		padding: 26px 24px;
	}

	// Desktop
	// ==========================================================================

	@include media-desk {
		padding: 26px 24px;

		&-icon {
			margin-right: 24px;
		}
	}
}

</style>