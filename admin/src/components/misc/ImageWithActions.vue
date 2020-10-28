<!-- =====================
	Image With Actions
	===================== -->
<template>
	<figure class="image">
		<slot></slot>
		<div class="image-buttons">
			<i class="feather feather-edit image-buttons-choose" @click="$emit('choose', true)"></i>
			<i class="feather feather-trash image-buttons-remove" @click="$emit('remove', true)"></i>
		</div>
	</figure>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "ImageWithActions",
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.image {
	$self: &;

	position: relative;
	height: 250px;
	width: 100%;
	border-radius: 4px;
	overflow: hidden;

	// Image
	// =========================================================================

	img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 250ms ease;
		will-change: transform;
	}

	// Overlay
	// =========================================================================

	&:after {
		content: "";
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		z-index: 99;
		background-color: rgba($black, 0.3);
		opacity: 0;
		transition: opacity 250ms ease;
		will-change: opacity;
	}

	// Hover
	// =========================================================================

	&:hover {

		img {
			transform: scale(1.1);
		}

		#{$self}-buttons,
		&:after {
			opacity: 1;
		}
	}

	// Buttons
	// =========================================================================

	&-buttons {
		position: absolute;
		display: flex;
		top: 20px;
		right: 20px;
		z-index: 100;
		opacity: 0;
		transition: opacity 250ms ease;
		will-change: opacity;

		i {
			color: $white;
			background-color: rgba($white, 1);
			padding: 10px;
			border-radius: 4px;
			cursor: pointer;
			box-shadow: none;
			transition: box-shadow 100ms ease;
			will-change: box-shadow;

			&:first-of-type {
				margin-right: 5px;
			}

			&:hover {
				box-shadow: 0 3px 10px 0 rgba($black, 0.14);
			}
		}

		&-remove {
			color: $orange !important;
		}

		&-choose {
			color: $green !important;
		}
	}

	// Full Width
	// =========================================================================

	&-cover {
		width: 100% !important;
		height: 100% !important;
	}

	// Tablet
	// =========================================================================

	@include media-tab {
		width: 250px;
	}
}

</style>