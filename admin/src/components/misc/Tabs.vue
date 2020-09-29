<!-- =====================
	Tabs
	===================== -->
<template>
	<ul class="tabs" ref="tabs" :class="{ 'tabs-loading' : loading }">
		<li class="tabs-item" @click="changeTab($event, 0)" :class="{ 'tabs-item-active' : activeTab === 0 }">
			<h5 class="tabs-title">Show All</h5>
		</li>
		<li class="tabs-item" @click="changeTab($event, 1)" :class="{ 'tabs-item-active' : activeTab === 1 }">
			<h5 class="tabs-title">Published</h5>
		</li>
		<li class="tabs-item" @click="changeTab($event, 2)" :class="{ 'tabs-item-active' : activeTab === 2 }">
			<h5 class="tabs-title">Drafts</h5>
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
	data: () => ({
		loading: true,
		activeTab: 0,
	}),
	mounted() {
		const firstTab = this.$refs.tabs.childNodes[0];
		this.updatePosition(firstTab, 0);
		setTimeout(() => {
			this.loading = false;
		}, 200);
	},
	methods: {
		changeTab(e, index) {
			this.updatePosition(e.target, index);
		},
		updatePosition(el, index) {
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
	flex-wrap: wrap;
	margin: 36px 0;
	width: 100%;
	border-bottom: 1.4px solid $grey-light;
	overflow-x: auto;
	overflow-y: hidden;
	white-space: nowrap;

	// Item
	// =========================================================================

	&-item {
		position: relative;
		display: block;
		margin: 0 24px;
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
		width: 100px;
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
}


</style>