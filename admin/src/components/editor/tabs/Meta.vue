<!-- =====================
	Meta Options
	===================== -->
<template>
	<section>
		<!-- Serp -->
		<div class="field">
			<h4>SERP Preview</h4>
			<div class="field-content">
				<div class="serp">
					<span class="serp-title" v-html="data.title">How the development team got their jobs at Reddico</span>
					<span class="serp-url">https://reddico.co.uk/how-we-got-our-development-jobs/</span>
					<span class="serp-description" v-html="data.description">Have you ever wondered how to become a developer? Our team has taken different routes to get to their jobs here – with some surprising diversions along the way. We asked them to explain how they’ve progressed over the years.</span>
				</div>
			</div>
		</div>
		<!-- Meta -->
		<div class="field">
			<h4>Meta</h4>
			<div class="field-content">
				<div class="form-group">
					<h5>Title</h5>
					<p>This will appear at the top of the search preview</p>
					<input class="form-input form-input-white" type="text" v-model="data.title" @keyup="emit">
				</div>
				<div class="form-group">
					<h5>Description</h5>
					<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
					<textarea class="form-input form-input-white" type="text" rows="4" v-model="data.description" @keyup="emit"></textarea>
				</div>
			</div>
		</div>
		<!-- Facebook -->
		<div class="field">
			<h4>Facebook (Open Graph)</h4>
			<div class="field-content">
				<div class="form-group">
					<h5>Title</h5>
					<p>This will appear at the top of the search preview</p>
					<input class="form-input form-input-white" type="text" v-model="data.facebook.title" @keyup="emit">
				</div>
				<div class="form-group">
					<h5>Description</h5>
					<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
					<textarea class="form-input form-input-white" type="text" rows="4" v-model="data.facebook.description" @keyup="emit"></textarea>
				</div>
			</div>
		</div>
		<!-- Twitter -->
		<div class="field">
			<h4>Twitter</h4>
			<div class="field-content">
				<div class="form-group">
					<h5>Title</h5>
					<p>This will appear at the top of the search preview</p>
					<input class="form-input form-input-white" type="text" v-model="data.twitter.title" @keyup="emit">
				</div>
				<div class="form-group">
					<h5>Description</h5>
					<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
					<textarea class="form-input form-input-white" type="text" rows="4" v-model="data.twitter.description" @keyup="emit"></textarea>
				</div>
			</div>
		</div>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "MetaOptions",
	data: () => ({
		data: {
			"title": "",
			"description": "",
			"facebook": {
				"title": "",
				"description": "",
			},
			"twitter": {
				"title": "",
				"description": "",
			}
		}
	}),
	methods: {
		emit() {
			this.$emit("update", this.data)
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>