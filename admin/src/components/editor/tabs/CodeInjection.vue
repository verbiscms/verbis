<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- Head -->
		<div class="form-group">
			<h4 class="card-title">Header</h4>
			<prism-editor class="my-editor" v-model="headerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
		</div>
		<!-- Footer -->
		<h4>Footer</h4>
		<div class="form-group">
			<prism-editor class="my-editor" v-model="footerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
		</div>
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

//https://github.com/koca/vue-prism-editor

// TODO  ! Move to Vue.Use and have a prism component in the scss

// import Prism Editor
import { PrismEditor } from 'vue-prism-editor';
import 'vue-prism-editor/dist/prismeditor.min.css';
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-markup';
import 'prismjs/themes/prism-twilight.css'; // import syntax highlighting styles

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
	components: {
		PrismEditor,
	},
	data: () => ({
		headerVal: "",
		footerVal: "",
		code: 'console.log("Hello World")'
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


<style>
/* required class */
.my-editor {
	/* we dont use `language-` classes anymore so thats why we need to add background and text color manually */
	background: #2d2d2d;
	color: #ccc;
	height: 500px;

	/* you must provide font-family font-size line-height. Example: */
	font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
	font-size: 14px;
	line-height: 1.5;
	padding: 5px;
	border-radius: 6px;
}

/* optional class for removing the outline */
.prism-editor__textarea:focus {
	outline: none;
}
</style>