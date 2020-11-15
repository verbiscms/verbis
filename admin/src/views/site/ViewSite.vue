<!-- =====================
	Home
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<iframe :src="getSiteUrl"></iframe>
		</div>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>


export default {
	name: "ViewSite",
	title: "Site",
	components: {
	},
	data: () => ({}),
	methods: {},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.auth-container {
	padding: 0;
}

iframe {
	display: block;
	width: 100%;
	height: 100vh;
	min-height: 100%;
	border: none;
	overflow-y: scroll;
}

</style>