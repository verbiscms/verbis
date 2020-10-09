<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- Head -->
		<div class="field">
			<h4>Header</h4>
			<div class="form-group">
				<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
				<textarea class="form-input form-input-white" type="text" rows="14" v-model="headerVal" @keyup="emit"></textarea>
			</div>
		</div>
		<!-- Footer -->
		<div class="field">
			<h4>Footer</h4>
			<div class="form-group">
				<p>This will appear at the bottom of the search preview, recommended 240 characters.</p>
				<textarea class="form-input form-input-white" type="text" rows="14" v-model="footerVal" @keyup="emit"></textarea>
			</div>
		</div>
		{{ header }}
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "CodeInjection",
	props: {
		header: {
			type: String,
			default: "",
		},
		footer: {
			type: String,
			default: "",
		},
	},
	data: () => ({
		headerVal: "",
		footerVal: "",
	}),
	mounted() {
		this.headerVal = this.header;
		this.footerVal = this.footer;
	},
	methods: {
		emit() {
			this.$emit("update", {
				"header": this.headerVal,
				"footer": this.footerVal,
			})
		}
	},
}

</script>
