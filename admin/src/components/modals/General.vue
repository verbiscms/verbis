<!-- =====================
	Modal
	===================== -->
<template>
	<transition name="trans-fade">
		<div v-if="showModal" class="modal modal-centered" :class="{ 'modal-open' : showModal }" aria-hidden="true">
			<!-- Container -->
			<div class="modal-container trans-fade-down-anim">
				<!-- Close -->
				<div class="modal-close" @click="showModal = false">
					<i class="feather feather-x"></i>
				</div>
				<!-- Warning -->
				<div class="modal-warning">
					<div class="modal-icon">
						<i class="feather feather-alert-triangle"></i>
					</div>
				</div>
				<!-- Text -->
				<div class="modal-text">
					<slot name="text"></slot>
				</div>
				<slot name="button"></slot>
			</div><!-- /Container -->
			<!-- /Overlay -->
			<div class="modal-overlay" @click="showModal = false"></div>
		</div><!-- /Modal -->
	</transition>
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
	overflow-y: hidden;
	padding: 75px 0;
	z-index: 9999;

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

	// Full Width
	// ==========================================================================

	&-full-width {

		#{$self}-container {
			width: 90vw;
			max-width: 1800px;
			min-width: 0 !important;
		}
	}

	// Text
	// ==========================================================================

	&-text {
		margin-bottom: 1.6rem;

		h2 {
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
		z-index: 999;

		i {
			color: $secondary;
			transition: color 100ms ease;
			will-change: color;
		}

		&:hover i {
			color: $orange;
		}
	}

	// Sizes
	// ==========================================================================

	@include media-desk {

		&-large {

			#{$self}-container {
				max-width: 700px !important;
				min-width: 600px !important;
			}
		}
	}


	// Container
	// ==========================================================================

	&-container {
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
		overflow-y: auto;
		transition: z-index $modal-transition-time step-end;
		padding: $modal-padding-desk;
		will-change: transform, opacity;

		@include media-tab {
			max-width: 500px;
			min-width: 400px;
		}
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
		cursor: pointer;
		background: rgba($black, 0.2);
		z-index: -1;
		opacity: 0;
		transition: opacity $modal-transition-time ease;
	}

	// Open
	// ==========================================================================

	&-open {
		visibility: visible;

		#{$self}-container {
			z-index: 9999999;
			transition: z-index $modal-transition-time step-start;
		}

		#{$self}-overlay {
			opacity: 1;
		}
	}

	// Hide Close Button
	// ==========================================================================

	&-hide-close {

		#{$self}-close {
			display: none !important;
		}
	}

	// Centered
	// ==========================================================================

	&-centered {

		#{$self}-container {
			margin: auto;
		}
	}

	// With Icon
	// ==========================================================================

	&-with-icon {

		#{$self}-container {
			align-items: center;

			* {
				text-align: center;
			}
		}
	}

	// Warning
	// ==========================================================================

	&-warning {
		display: none;
	}

	&-with-warning {

		#{$self}-warning {
			display: block;
		}
	}
}

</style>