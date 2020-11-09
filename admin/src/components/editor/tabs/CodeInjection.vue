<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- =====================
			Head
			===================== -->
		<h6 class="margin">Header</h6>
		<div class="card card-small-box-shadow card-expand">
			<Collapse class="collapse-border-bottom">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Head</h4>
							<p>Any code inputted here will be injected the <br class="d-mob-none"><code v-html="'{{ verbisHead }}'"></code> of the site.</p>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<prism-editor class="prism" v-model="headerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
					</div><!-- /Card Body -->
				</template>
			</Collapse>
		</div><!-- /Card -->
		<!-- =====================
			Footer
			===================== -->
		<h6 class="margin">Footer</h6>
		<div class="card card-small-box-shadow card-expand">
			<Collapse class="collapse-border-bottom">
				<template v-slot:header>
					<div class="card-header">
						<div>
							<h4 class="card-title">Foot</h4>
							<p>Any code inputted here will be injected to the <code v-html="'{{ verbisFooter }}'"></code> before the closing body tag.</p>
						</div>
						<div class="card-controls">
							<i class="feather feather-chevron-down"></i>
						</div>
					</div><!-- /Card Header -->
				</template>
				<template v-slot:body>
					<div class="card-body">
						<prism-editor class="prism" v-model="footerVal" :highlight="highlighter" line-numbers @keyup="emit"></prism-editor>
					</div><!-- /Card Body -->
				</template>
			</Collapse>
		</div><!-- /Card -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';
import Collapse from "@/components/misc/Collapse";

export default {
	name: "CodeInjection",
	components: {
		Collapse,
	},
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