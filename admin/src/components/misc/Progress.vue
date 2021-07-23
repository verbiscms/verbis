<!-- =====================
	Progress
	===================== -->
<template>
	<div class="progress">
		<!-- Bar -->
		<div class="progress-bar">
			<div class="progress-bar-line" :style="{ width: percent + '%' }"></div>
		</div><!-- /Bar -->
		<!-- Text -->
		<div class="progress-text">
			<slot></slot>
		</div><!-- /Text -->
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Progress",
	props: {
		percent: {
			type: Number,
			default: 0,
		},
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

.progress {
	$self: &;
	margin-bottom: 1rem;

	// Bar
	// ==========================================================================

	&-bar {
		width: 100%;
		height: 26px;
		background-color: $white;
		box-shadow: $form-box-shadow;
		border-radius: 4px;
		overflow: hidden;
		margin-bottom: 10px;

		&-line {
			position: relative;
			background-color: $primary;
			width: 30%;
			height: 100%;
			transition: width 30ms ease;
		}
	}
}

</style>
