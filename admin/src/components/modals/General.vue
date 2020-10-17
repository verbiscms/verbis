<!-- =====================
	Modal
	===================== -->
<template>
	<div class="modal modal-centered" :class="{ 'modal-open' : showModal }" aria-hidden="true">
		<!-- Container -->
		<div class="modal-container">
			<!-- Close -->
			<div class="modal-close" @click="showModal = false">
				<i class="feather feather-x"></i>
			</div>
			<!-- Icon -->
			<div class="modal-icon">
				<i class="feather feather-alert-triangle"></i>
			</div>
			<!-- Text -->
			<div class="modal-text">
				<slot name="text"></slot>
			</div>
			<slot name="button"></slot>
		</div><!-- /Container -->
		<!-- /Overlay -->
		<div class="modal-overlay"></div>
	</div><!-- /Modal -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Modal",
	props: {
		show: {
			type: Boolean,
			default: false,
		}
	},
	computed: {
		showModal: {
			get() {
				return this.show;
			},
			set(value) {
				this.$emit("update:show", value);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

// Variables
$modal-padding-desk: 25px;
$modal-transition-time: 400ms;

.modal {
	$self: &;

	visibility: visible;
	position: fixed;
	display: flex;
	align-items: flex-start;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	z-index: -1;
	overflow-y: auto;
	padding: 75px 0;
	transition: z-index $modal-transition-time step-end;

	// Icon
	// ==========================================================================

	&-icon {
		display: inline-block;
		margin: 0 auto 1rem auto;

		i {
			font-size: 56px;
			background: linear-gradient(0deg, rgba($orange, 100%) 0, rgba($orange, 0.7) 100%);
			-webkit-background-clip: text;
			-webkit-text-fill-color: transparent;
		}
	}

	// Text
	// ==========================================================================

	&-text {
		margin-bottom: 1.6rem;

		h2 {
			text-align: center;
			margin-bottom: 2px;
		}

		p {
			text-align: center;
			margin-bottom: 0;
		}
	}

	// Close
	// ==========================================================================

	&-close {
		position: absolute;
		top: 0;
		right: 4px;
		padding: 16px;
		cursor: pointer;

		i {
			color: $secondary;
		}
	}

	// Container
	// ==========================================================================

	&-container {
		max-width: 500px;
		min-width: 400px;
		position: relative;
		display: flex;
		flex-direction: column;
		justify-content: center;
		background-color: #fff;
		box-shadow: 0 0 20px 2px rgba($black, 0.16);
		border-radius: 6px;
		border: 0;
		margin: 0 auto;
		opacity: 0;
		transform: translateY(-50px);
		transition: opacity $modal-transition-time ease, transform $modal-transition-time ease;
		padding: $modal-padding-desk;
		will-change: transform, opacity
	}

	// Open
	// ==========================================================================

	&-overlay {
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		position: fixed;
		display: block;
		cursor: default;
		background: rgba($black, 0.1);
		z-index: -1;
		opacity: 0;
		transition: opacity $modal-transition-time ease;
	}

	// Open
	// ==========================================================================

	&-open {
		visibility: visible;
		z-index: 9999999;
		transition: z-index $modal-transition-time step-start;

		#{$self}-container {
			opacity: 1;
			transform: translateY(0);
		}

		#{$self}-overlay {
			opacity: 1;
		}
	}

	// Centered
	// ==========================================================================

	&-centered {

		#{$self}-container {
			margin: auto;
		}
	}
}

</style>