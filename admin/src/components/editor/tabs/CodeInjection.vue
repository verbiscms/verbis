<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- Head -->
		<div class="form-group">
			<h4 class="card-title">Header</h4>
			<prism-editor class="prism" v-model="headerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
		</div>
		<!-- Footer -->
		<h4>Footer</h4>
		<div class="form-group">
			<prism-editor class="prism" v-model="footerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
		</div>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';

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
		highlighter(code) {
			return highlight(code, languages.html);
		},
		emit() {
			this.$emit("update", {
				"header": this.headerVal,
				"footer": this.footerVal,
			})
		}
	},
}

</script>