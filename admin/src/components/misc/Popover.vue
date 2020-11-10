<!-- =====================
	Popover
	===================== -->
<template>
	<div class="popover-cont" :class="{ 'popover-hover' : hover }">
		<div class="popover-btn" @click="update">
			<slot name="button"></slot>
		</div>
		<div class="popover" :class="[{ 'popover-triangle' : triangle }, { 'popover-active' : show && itemKey === active }, classes]" ref="popover">
			<div class="popover-items-cont">
				<slot name="items"></slot>
			</div>
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
		},
		triangle: {
			type: Boolean,
			default: false,
		},
		classes: {
			type: String,
			default: "",
		},
		itemKey: {
			type: String,
			default: "",
		},
		active: {
			type: String,
			default: "",
		},
		hover: {
			type: Boolean,
			default: false
		}
	},
	data: () => ({
		show: false,
	}),
	methods: {
		/*
		 * update()
		 * Show & hide the popover, and emit data.
		 */
		update() {
			if (!this.hover) {
				this.show = !this.show;
				this.$emit("update", this.show);
			}
		},
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
<style lang="scss">


// Variables
$popover-triangle-size: 20px;

.popover {
	$self: &;

	position: absolute;
	top: 0;
	left: 50%;
	transform: translate(-50%, -100%);
	border: 1px solid $grey-light;
	background-color: $white;
	border-radius: 8px;
	box-shadow: 0 3px 20px 0 rgba($black, 0.08);
	z-index: -1;
	opacity: 0;
	transition: opacity 200ms ease, z-index 200ms step-end;

	// Container
	// =========================================================================

	&-cont {
		display: inline-block;
		position: relative;
		width: auto;

		&#{$self}-hover {

			.popover {
				transition: opacity 200ms ease, z-index 200ms step-end !important;
			}


			.icon,
			i,
			button {
				z-index: 99;
			}

			.popover:hover,
			.icon:hover,
			i:hover + #{$self},
			#{$self}-btn:hover + #{$self} {
				opacity: 1 !important;
				z-index: 999999;
				transition: opacity 200ms ease, z-index 200ms step-start !important;
			}

			#{$self} {
				bottom: 0;
				left: 0;
				top: auto;
				transform: translate(calc(-50% + 2px), 100%);
			}
		}
	}

	// Triangle
	// =========================================================================

	&-triangle {

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
	}

	// Line
	// =========================================================================

	&-line {
		&:before {
			content: "";
			display: block;
			position: relative;
			width: 100%;
			height: 1px;
			background-color: $grey-light;
		}
	}

	// Item
	// =========================================================================

	&-item {
		position: relative;
		width: 180px;
		font-size: 13px;
		font-weight: 500;
		color: rgba($secondary, 0.7);
		padding: 6px 0;
		text-align: center;
		cursor: pointer;
		margin: 4px;
		border-radius: 4px;

		&-icon {
			display: flex;
			justify-content: flex-start;
			align-items: center;
			padding-left: 16px;
			padding-right: 16px;
			text-align: left;

			i {
				margin-right: 12px;
			}
		}

		&:not(div):not(&-orange) {

			i,
			span {
				color: rgba($secondary, 0.7) !important;
			}
		}

		&:hover {
			background-color: rgba($primary, 0.07);
		}

		&-orange {

			i,
			span {
				color: $orange;
			}

			&:hover {
				background-color: rgba($orange, 0.07);

				i,
				span {
					color: $orange;
				}
			}
		}
	}

	// Button
	// =========================================================================

	&-btn {
		position: relative;
		display: block;
		//padding: 10px;
		//margin-right: -10px;
		z-index: 9;
	}

	// Active
	// =========================================================================

	&-active {
		opacity: 1;
		z-index: 99999;
		transition: opacity 200ms ease, z-index 200ms step-start;
	}

	// Mods
	// =========================================================================

	&-no-arrow {

		&::before {
			display: none !important;
		}
	}

	// Table
	// =========================================================================

	&-table {
		position: fixed;

		&-top {
			top: auto;
			bottom: auto;
			transform: translate(-110px, calc(-100% - 42px));
			left: auto;
			right: 0;
		}

		&-bottom {
			top: auto;
			bottom: auto;
			transform: translate(-110px, 10px);
			left: auto;
			right: 0;
		}
	}

	// Positions
	// =========================================================================

	&-bottom {

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