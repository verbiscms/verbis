<!-- =====================
	New Menu Modal
	===================== -->
<template>
	<el-dialog :visible.sync="show" width="30%">
		<!-- Title -->
		<div slot="title">
			<h2 class="mb-0">Create a proxy</h2>
			<p class="mb-0">Enter the fields below to create a new reverse proxy.</p>
		</div>
		<!-- Form -->
		<el-form :model="proxy" ref="newMenu" label-position="top" label-width="auto" :rules="rules">
			<!-- Path -->
			<el-form-item label="Path" prop="path" :rules="{ required: true, message: 'Enter a Path.', trigger: 'blur' }">
				<el-input placeholder="Path" v-model="proxy.path"></el-input>
			</el-form-item>
			<!-- Host -->
			<el-form-item label="Host" prop="host" :rules="{ required: true, message: 'Enter a Host.', trigger: 'blur' }">
				<el-input placeholder="Host" v-model="proxy.host"></el-input>
			</el-form-item>
			<!-- Rewrite -->
			<div class="proxy-rewrite mb-4">
				<!-- Header -->
				<div class="proxy-rewrite-header">
					<h5>Rewrites</h5>
					<el-button @click="addRewrite" size="small" icon="el-icon-plus"></el-button>
				</div>
				<!-- Item -->
				<div class="proxy-rewrite-item" v-for="(rewrites, index) in proxy['rewrites']" :key="index" prop="rewrites">
					<el-input size="small" placeholder="From" v-model="proxy['rewrites'][index].from"></el-input>
					<el-input class="mx-1" size="small" placeholder="To" v-model="proxy['rewrites'][index].to"></el-input>
					<el-button @click="removeRewrite(index)" size="small" type="danger" icon="el-icon-delete"></el-button>
				</div>
			</div><!-- /Rewrite -->
			<!-- Regex Rewrite -->
			<div class="proxy-rewrite">
				<!-- Header -->
				<div class="proxy-rewrite-header">
					<h5>Regex Rewrites</h5>
					<el-button @click="addRewriteRegex" size="small" icon="el-icon-plus"></el-button>
				</div>
				<!-- Item -->
				<div class="proxy-rewrite-item" v-for="(rewrites, index) in proxy['rewrite_regex']" :key="index" prop="regex_rewrites">
					<el-input size="small" placeholder="From" v-model="proxy['rewrite_regex'][index].from"></el-input>
					<el-input class="mx-1" size="small" placeholder="To" v-model="proxy['rewrite_regex'][index].to" ></el-input>
					<el-button @click="removeRewriteRegex(index)" size="small" type="danger" icon="el-icon-delete"></el-button>
				</div>
			</div><!-- Regex Rewrite -->
		</el-form>
		<!-- Footer -->
		<span slot="footer" class="dialog-footer">
			<el-button @click="show = false">Cancel</el-button>
			<el-button type="primary" @click="create">Create Proxy</el-button>
		</span>
	</el-dialog>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>
export default {
	name: "NewProxyModal",
	props: {
		visible: {
			type: Boolean,
			default: false,
		},
	},
	data: () => ({
		proxy: {},
		rules: {
			rewrite: [
				{ required: true, message: 'Enter a from path.', trigger: 'blur' },
				{ required: true, message: 'Enter a to path.', trigger: 'blur' },
			]

		}
	}),
	mounted() {
		//this.show = true;
		this.proxy = {
			"host": "https://35.214.23.223:5000/tools/serp-speed",
			"path": "/tools/serp-speed",
			"rewrites": [{
				from: "map/$1",
				to: "map/$3",
			}],
			"rewrite_regex": [],
		}
	},
	methods: {
		/**
		 * Checks to see if the form data is valid then
		 * updates the parent with the new menu
		 * information. Resets the proxy data
		 * upon success.
		 */
		create() {
			this.$refs['newMenu'].validate((valid) => {
				if (valid) {

					// Logic here to flatten wrirtes

					this.$emit("create", this.proxy);
					this.proxy = {};
				}
			});
		},
		/**
		 * Adds a rewrite to the proxy.
		 */
		addRewrite() {
			this.proxy['rewrites'].push({
				from: "",
				to: "",
			});
		},
		/**
		 * Removes a rewrite from the proxy.
		 */
		removeRewrite(index) {
			if (index !== -1) {
				this.proxy.rewrites.splice(index, 1);
			}
		},
		/**
		 * Adds a rewrite regex to the proxy.
		 */
		addRewriteRegex() {
			this.proxy['rewrite_regex'].push({
				from: "",
				to: "",
			});
		},
		/**
		 * Removes a rewrite regex from the proxy.
		 */
		removeRewriteRegex(index) {
			if (index !== -1) {
				this.proxy['rewrite_regex'].splice(index, 1);
			}
		},
	},
	computed: {
		/**
		 * Computed prop for showing or hiding the
		 * modal.
		 */
		show: {
			get() {
				return this.visible;
			},
			set(val) {
				this.$emit("update:visible", val)
			}
		}
	}
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.proxy {

	&-rewrite {

		&-header {
			display: flex;
			align-items: center;
			justify-content: space-between;
			margin-bottom: 10px;
		}

		&-item {
			display: flex;
			margin-bottom: 8px;
		}
	}
}

</style>
