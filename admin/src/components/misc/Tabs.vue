<!-- =====================
	Tabs
	===================== -->
<template slot="buttons" scope="props">
	<ul class="tabs" ref="tabs" :class="{ 'tabs-loading' : loading }" >
		<li class="tabs-item" v-for="(tab, tabIndex) in tabs" :key="tabIndex" ref="tab" @click="changeTab($event, tabIndex)" :class="{ 'tabs-item-active' : activeTab === tabIndex }">
			<span class="tabs-title">{{ tab }}</span>
		</li>
		<li class="tabs-indicator" ref="indicator"></li>
	</ul><!-- /Tabs -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Tabs",
	props: {
		defaultTab: {
			type: Number,
			default: 0,
		}
	},
	data: () => ({
		loading: true,
		activeTab: 0,
		tabs: []
	}),
	mounted() {
		this.setUpTabs();
		this.activeTab = this.defaultTab;

		this.$nextTick(() => {
			const firstTab = this.$refs['tab'][this.defaultTab]
			this.updatePosition(firstTab, this.defaultTab);
			setTimeout(() => {
				this.loading = false;
			}, 100);
		});
	},
	methods: {
		/*
		 * setUpTabs()
		 * Init the tabs and push to the tabs array.
		 */
		setUpTabs() {
			this.$slots.item.forEach(item => {
				this.tabs.push(item.text)
			});
		},
		/*
		 * changeTab()
		 * Update the position and emit the index.
		 */
		changeTab(e, index) {
			this.updatePosition(e.target, index);
			this.$emit("update", index + 1)
		},
		/*
 		 * updatePosition()
		 * Change the left and right width of the tab by boundingClientRect().
		 */
		updatePosition(el = false, index) {
			this.activeTab = index;

			const tabs = this.$refs.tabs,
				indicator = this.$refs.indicator,
				bounding = el.getBoundingClientRect(),
				tabsBounding = tabs.getBoundingClientRect();

			let props = {
				left: (bounding.left - tabsBounding.left),
				width: bounding.width,
			}

			indicator.style.left = props.left - 8 + "px";
			indicator.style.width = props.width + 16 + "px";
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Variables
$tabs-padding: 16px 6px;
$tabs-underline-height: 3px;

.tabs {
	$self: &;

	position: relative;
	display: flex;
	flex-wrap: nowrap;
	margin-bottom: $header-margin-bottom-mob;
	width: 100%;
	border-bottom: 1.4px solid $grey-light;
	overflow-x: scroll;
	white-space: nowrap;

	// Item
	// =========================================================================

	&-item {
		position: relative;
		display: block;
		margin: 0 16px;
		cursor: pointer;
		padding: $tabs-padding;

		&:last-child {
			margin-right: 0;
		}

		&:first-child {
			margin-left: 0;
		}

		&:hover {

			#{$self}-title {
				color: $primary;
			}
		}
	}

	// Title
	// =========================================================================

	&-title {
		font-size: 1rem;
		color: $grey;
		font-weight: 600;
		transition: color 200ms ease;
		will-change: color;
	}

	// Underline
	// =========================================================================

	&-indicator {
		position: absolute;
		bottom: 0;
		left: 0;
		display: block;
		height: $tabs-underline-height;
		width: auto;
		background-color: $primary;
		opacity: 1;
		transition: left 200ms ease, width 200ms ease, opacity 1000ms ease;
		will-change: left, width, opacity;
	}

	// Active
	// =========================================================================

	&-item-active {

		#{$self}-title {
			color: $copy;
		}
	}

	// Loading
	// =========================================================================

	&-loading {

		#{$self}-indicator {
			transition: none;
		}
	}

	// Tablet
	// =========================================================================

	@include media-tab {
		margin-bottom: $header-margin-bottom-tab;

		&-item {
			margin: 0 20px;
		}
	}

	// Desktop
	// =========================================================================

	@include media-desk {
		margin-bottom: $header-margin-bottom-desk;
		overflow: hidden;

		&-item {
			margin: 0 24px;
		}
	}
}


</style>