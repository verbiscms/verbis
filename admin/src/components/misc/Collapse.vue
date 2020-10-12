<!-- =====================
	Collapse
	===================== -->
<template>
	<div class="expand">
		<div class="collapse">
			<div class="collapse-item">
				<div class="collapse-header" ref="header">
					<slot name="header"></slot>
				</div>
				<div class="collapse-content" ref="content">
					<slot name="body"></slot>
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
		chevron: {
			default: null
		}
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
		this.$nextTick(() => {
			setTimeout(() => {
				this.setDefaults();
			}, 100)
		})
	},
	methods: {
		addListener() {
			const chevron = this.header.querySelector(".feather-chevron-down");
			chevron.addEventListener("click", () => {
				this.collapse(chevron)
			});
		},
		setDefaults() {
			const height = this.content.getBoundingClientRect().height,
				variable = this.content.querySelectorAll(".field").length * 35;
			this.content.style.maxHeight = (height + variable + 40) + "px";
		},
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