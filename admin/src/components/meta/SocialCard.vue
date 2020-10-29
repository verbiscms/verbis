<!-- =====================
	Social Card
	===================== -->
<template>
	<div class="social">
		<div class="social-image">
			<i v-if="!image" class="fal fa-file-alt"></i>
			<img v-else :src="getSiteUrl + image['url']">
		</div>
		<div class="social-text">
			<div class="social-text-cont">
				<span class="social-title" v-text="clipText(140, title)"></span>
				<span class="social-description" v-text="clipText(230, description)"></span>
			</div>
			<span class="social-url">{{ getSiteUrl }}</span>
		</div>
	</div><!-- /Twitter Card -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "SocialCard",
	props: {
		description: {
			type: String,
			default: "",
		},
		title: {
			type: String,
			default: "",
		},
		image: {
			type: [Object, Boolean],
			default: false,
		},
	},
	methods: {
		/*
 		 * clipText()
		 * Clips the description at a given length.
		 */
		clipText(length, text) {
			if (text && length < text.length) {
				return text.substring(0, length - 3) + "...";
			}
			return text;
		},
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">


.social {
	position: relative;
	display: flex;
	flex-direction: column;
	width: 100%;
	border: 1px solid $grey-light;
	border-radius: 6px;
	overflow: hidden;
	margin-bottom: 1rem;

	// Image
	// =========================================================================

	&-image {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 160px;
		background-color: $grey-light;

		i {
			color: $grey;
			font-size: 2rem;
		}

		img {
			width: 100%;
			height: 100%;
			object-fit: cover;
			border-bottom: 1px solid $grey-light;
		}
	}

	// Text
	// =========================================================================

	&-text {
		padding: 14px;
		background-color: $white;
	}

	&-title,
	&-description,
	&-url {
		color: $black;
		display: block;
	}

	&-title {
		font-size: 1rem;
		font-weight: 600;
		max-height: 38px;
		//white-space: nowrap;
		margin-bottom: 6px;
		line-height: 1.2;
		overflow: hidden;
		text-overflow: ellipsis
	}

	&-description {
		color: rgba($secondary, 0.80);
		font-size: 0.8rem;
		height: 72px;
		overflow: hidden;
		line-height: 1.3;
		margin-bottom: 4px;
		word-break: break-all;
	}

	&-url {
		color: $grey;
		font-size: 0.8rem;
	}

	// Tablet
	// =========================================================================

	@include media-tab {
		flex-direction: row;
		margin-bottom: 0;

		&-image {
			width: 170px;
			min-width: 170px;
			height: auto;

			img {
				border-right: 1px solid $grey-light;
				border-bottom: none;
			}
		}

		&-text {
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			width: calc(100% - 170px);
		}
	}
}

</style>