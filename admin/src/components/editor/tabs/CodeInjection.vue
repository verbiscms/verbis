<!-- =====================
	CodeInjection
	===================== -->
<template>
	<section>
		<!-- =====================
			Head
			===================== -->
		<div class="card">
			<collapse>
				<template v-slot:header>
					<div class="card-header">
						<h3 class="card-title">Header</h3>
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
			</collapse>
		</div><!-- /Card -->
		<!-- =====================
			Footer
			===================== -->
		<div class="card">
			<collapse>
				<template v-slot:header>
					<div class="card-header">
						<h3 class="card-title">Footer</h3>
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
			</collapse>
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