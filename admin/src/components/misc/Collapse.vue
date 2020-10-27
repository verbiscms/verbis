<!-- =====================
	Collapse
	===================== -->
<template>
	<div class="collapse">
		<div class="expand">
			<div v-if="!reverse">
				<div class="collapse-item">
					<div class="collapse-header" ref="header">
						<slot name="header"></slot>
					</div>
					<div class="collapse-content" ref="content">
						<slot name="body"></slot>
					</div>
				</div>
			</div>
			<div v-else>
				<div class="collapse-item">
					<div class="collapse-content" ref="content">
						<slot name="body"></slot>
					</div>
					<div class="collapse-header" ref="header">
						<slot name="header"></slot>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Collapse",
	props: {
		show: {
			type: Boolean,
			default: true,
		},
		useIcon: {
			type: Boolean,
			default: true,
		},
		reverse: {
			type: Boolean,
			default: false,
		},
	},
	data: () => ({
		isOpen: true,
		content: {},
		header: {},
	}),
	mounted() {
		this.content = this.$refs.content;
		this.header = this.$refs.header;
		this.addListener()
		this.setDefaults();
		// this.$nextTick(() => {
		// 	setTimeout(() => {
		// 		this.setDefaults();
		// 	}, 100);
		// })
	},
	methods: {
		/*
		 * addListener()
		 * Add the event listener to the header or chevron based on props.
		 */
		addListener() {
			if (!this.useIcon) {
				this.header.addEventListener("click", () => {
					this.collapse(this.header.querySelector("i"))
				});
			} else {
				const chevron = this.header.querySelector(".feather-chevron-down");
				chevron.addEventListener("click", () => {
					this.collapse(chevron)
				});
			}
		},
		/*
		 * setDefaults()
		 * If the show prop is true, calculate the height and set max height.
		 * Else set to 0.
		 */
		setDefaults() {
			if (this.show) {
				const height = this.content.getBoundingClientRect().height,
					variable = this.content.querySelectorAll(".field").length * 35;
				this.content.style.maxHeight = (height + variable + 40) + "px";
			} else {
				this.isOpen = false;
				this.content.style.maxHeight = 0 + "px"
			}
		},
		/*
		 * collapse()
		 * Change the maximum height on click, add chevron active class.
		 */
		collapse(chevron) {
			if (this.isOpen) {
				this.content.style.maxHeight = "0px";
				chevron.classList.add("active");
			} else {
				this.content.style.maxHeight = this.calcHeight(this.content) + "px";
				chevron.classList.remove("active");
			}
			this.isOpen = !this.isOpen;
		},
		/*
		 * calcHeight()
		 * Calculate the height of the el's children.
		 */
		calcHeight(el) {
			let children = el.children;
			let height = 0;
			for (let i = 0; i < children.length; i++) {
				height += children[i].clientHeight;
			}
			return height;
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.collapse {
		$self: &;

		&-border-bottom {

			#{$self}-content {
				border-bottom: 1px solid rgba($grey-light, 0.7);
			}
		}

		// Content
		// ==========================================================================

		&-content {
			overflow: hidden;
			padding: 0;
			transition: 0.5s;
		}

		// Chevron
		// ==========================================================================

		.feather-chevron-down {
			transition: transform 180ms ease;

			&.active {
				transform: rotate(180deg);
			}
		}
	}


</style>