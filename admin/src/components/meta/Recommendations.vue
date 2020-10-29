<!-- =====================
	Meta Recommendations
	===================== -->
<template>
	<div class="recs">
		<p>
			Recommended:
			<span class="recs-text">{{ getRecommendations() }}</span> characters,
			you've used <span class="recs-text" :class="{ 'recs-valid' : isValid, 'recs-error' : !isValid }">{{ getCharacterCount }}</span>
		</p>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

const metaTitleLength = 60,
	metaDescriptionLength = 160,
	faceBookTitleLength = 55,
	facebookDescriptionLength = 200,
	twitterTitleLength = 50,
	twitterDescriptionLength = 200;

export default {
	name: "Recommendations",
	props: {
		text: {
			type: String,
			default: "",
		},
		type: {
			type: String,
			default: "title",
		},
		usage: {
			type: String,
			default: "meta",
		}
	},
	data: () => ({
		isValid: true,
	}),
	mounted() {
		this.checkValidity();
	},
	watch: {
		text: function() {
			this.checkValidity();
		}
	},
	methods: {
		getRecommendations() {
			switch (this.usage) {
				case "meta": {
					return this.type === "title" ? metaTitleLength : metaDescriptionLength;
				}
				case "facebook": {
					return this.type === "title" ? faceBookTitleLength : facebookDescriptionLength;
				}
				case "twitter": {
					return this.type === "title" ? twitterTitleLength : twitterDescriptionLength;
				}
			}
		},
		checkValidity() {
			this.isValid = this.getCharacterCount < this.getRecommendations();
		}
	},
	computed: {
		getCharacterCount() {
			return this.text.length;
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.recs {

		// Valid & Error
		// =========================================================================

		&-valid {
			color: $green;
		}

		&-error {
			color: $orange;
		}

		// Span
		// =========================================================================

		&-text {
			font-weight: 600;
		}
	}

</style>