<!-- =====================
	Popover
	===================== -->
<template>
	<div class="popover-cont">
		<button class="btn btn-blue"><i class="fal fa-plus-circle"></i>Add Row</button>
		<div class="popover">
			<span class="popover-triangle"></span>
			<div class="popover-items-cont">
				<slot name="items"></slot>
			</div>
<!--			<div class="popover-item" v-for="item in getItems" :key="item.key">{{ item.name }}</div>-->
		</div>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Alert",
	props: {
		items: {
			type: Array,
		}
	},
	data: () => ({
		test: [],
	}),
	mounted() {
	//	this.setUpItems();
	},
	methods: {
	},
	computed: {
		/*
		 * getItems()
		 * Get the computed items props.
		 */
		getItems() {
			return this.items;
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


// Variables
$popover-triangle-size: 20px;

.popover {
	$self: &;

	position: absolute;
	top: 0;
	left: 50%;
	border: 1px solid $grey-light;
	background-color: $white;
	padding: 4px;
	border-radius: 8px;
	box-shadow: 0 3px 20px 0 rgba($black, 0.08);
	transform: translate(-50%, -100%);
	z-index: 99;
	//opacity: 0;
	transition: opacity 200ms ease;

	// Container
	// =========================================================================

	&-cont {
		display: inline-block;
		position: relative;
		width: auto;

		.icon,
		i,
		button {
			z-index: 99;
		}

		.popover:hover,
		.icon:hover,
		i:hover + #{$self},
		button:hover + #{$self} {
			opacity: 1;
		}
	}

	&::before {
		content: "";
		position: absolute;
		width: 22px;
		height: 14px;
		bottom: 0;
		left: 50%;
		background:  url("~@/assets/images/popover-triangle.svg") no-repeat;
		background-size: cover;
		z-index: 100;
		transform: translate(-50%, calc(100% - 1px)) rotate(180deg);
	}

	// Item
	// =========================================================================

	&-items-cont {

		div {
			position: relative;
			width: 120px;
			font-size: 13px;
			font-weight: 600;
			color: $secondary;
			padding: 6px 0;
			text-align: center;
			cursor: pointer;

			&-icon {
				display: flex;
				justify-content: flex-start;
				align-items: center;
				padding-left: 12px;
				padding-right: 12px;

				i {
					margin-right: 10px;
				}
			}

			&:hover {
				background-color: rgba($primary, 0.07);
			}
		}
	}

	// Mods
	// =========================================================================

	&-no-arrow {

		&::before {
			display: none !important;
		}
	}

	// Positions
	// =========================================================================

	&-top {

		&-left {
			left: 0;
			right: auto;
			transform: translate(-85%, -85%);
		}

		&-right {
			left: auto;
			right: 0;
			transform: translate(90%, -90%);
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {
	}

	// Desktop
	// =========================================================================

	@include media-desk {
	}
}


</style>