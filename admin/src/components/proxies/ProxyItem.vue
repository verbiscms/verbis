<!-- =====================
	New Menu Modal
	===================== -->
<template>
	<!-- Body -->
	<el-form class="item" :model="item" size="small" ref="form" label-width="120px" label-position="left">
		<!-- Name -->
		<el-form-item label="Name" prop="name" :rules="{ required: true, message: 'Enter a Name.', trigger: 'blur' }">
			<el-input placeholder="Add a name" v-model="item.name"></el-input>
		</el-form-item>
		<!-- Path -->
		<el-form-item label="Path" prop="path" :rules="{ required: true, message: 'Enter a Path.', trigger: 'blur' }">
			<el-input placeholder="Add a path" v-model="item.path"></el-input>
		</el-form-item>
		<!-- Host -->
		<el-form-item label="Host" prop="host" :rules="{ required: true, message: 'Enter a Host.', trigger: 'blur' }">
			<el-input placeholder="Add a host" v-model="item.host"></el-input>
		</el-form-item>
		<!-- Rewrites -->
		<div class="item-rewrite">
			<el-form-item label="Rewrites" prop="rewrites">
				<el-button size="mini" icon="el-icon-plus" @click="addRewrite()"></el-button>
			</el-form-item>
			<el-table v-if="rewrites && rewrites.length" size="mini" v-model="rewrites" :data="rewrites" border style="width: 100%; margin: 10px 0;">
				<!-- From Path -->
				<el-table-column prop="from" label="From Path">
					<template slot-scope="scope">
						<el-input placeholder="e.g. /api/*" :prop="'rewrite.' + scope.$index+ '.from'" v-model="scope.row.from" size="mini" :rules="{ required: true, message: 'Enter a From Path.', trigger: 'blur' }"></el-input>
					</template>
				</el-table-column>
				<!-- To Path -->
				<el-table-column prop="to" label="To Path">
					<template slot-scope="scope">
						<el-input placeholder="e.g. /$1" v-model="scope.row.to" size="mini"></el-input>
					</template>
				</el-table-column>
				<!-- Delete -->
				<el-table-column label="Actions" width="74px" align="right">
					<template slot-scope="scope">
						<el-button icon="el-icon-delete" size="mini" type="danger" @click="handleDeleteRewrite(scope.$index)"></el-button>
					</template>
				</el-table-column>
			</el-table>
			<p v-else>No rewrites</p>
		</div>
		<!-- Regex -->
		<div class="item-rewrite">
			<el-form-item label="Regex Rewrites" prop="regex">
				<el-button size="mini" icon="el-icon-plus" @click="addRewrite(true)"></el-button>
			</el-form-item>
			<el-table v-if="regex && regex.length" size="mini" :data="regex" border style="width: 100%; margin-top: 10px">
				<!-- From Path -->
				<el-table-column prop="from" label="From Path">
					<template slot-scope="scope">
						<el-input placeholder="e.g. /old/[0.9]+/" v-model="scope.row.from" size="mini"></el-input>
					</template>
				</el-table-column>
				<!-- To Path -->
				<el-table-column prop="to" label="To Path">
					<template slot-scope="scope">
						<el-input placeholder="e.g. /new" v-model="scope.row.to" size="mini"></el-input>
					</template>
				</el-table-column>
				<!-- Delete -->
				<el-table-column label="Actions" width="74px" align="right">
					<template slot-scope="scope">
						<el-button icon="el-icon-delete" size="mini" type="danger"  @click="handleDeleteRewrite(scope.$index, true)"></el-button>
					</template>
				</el-table-column>
			</el-table>
			<p v-else>No regex rewrites</p>
		</div>
	</el-form><!-- /Body -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>
export default {
	name: "ProxyItem",
	props: {
		item: {
			type: Object,
			required: true,
			default: () => {},
		},
		validating: {
			type: Boolean,
			required: false,
			default: false,
		}
	},
	data: () => ({
		isValid: false,
		form: {},
		rewrites: [],
		regex: [],
	}),
	watch: {
		rewrites: {
			handler: function(val) {
				this.item.rewrite = this.flattenRewrites(val);
			},
			deep: true
		},
		regex: {
			handler: function(val) {
				this.item.rewrite_regex = this.flattenRewrites(val);
			},
			deep: true
		},
		validating: function () {
			this.validateForm();
		},
	},
	mounted() {
		this.init();
	},
	methods: {
		/**
		 * Transforms the rewrites and regex rewrites
		 * and expands objects to arrays.
		 */
		init() {
			this.rewrites = this.expandRewrites(this.item.rewrite);
			this.regex = this.expandRewrites(this.item.rewrite_regex);
		},
		/**
		 * Adds a rewrite to the proxy.
		 * @param regex
		 */
		addRewrite(regex = false) {
			if (regex) {
				this.regex.push({from: "", to: ""});
				return;
			}
			this.rewrites.push({from: "", to: ""});
		},
		/**
		 * Deletes a rewrite or regex rewrite.
		 * @param index
		 * @param regex = false
		 */
		handleDeleteRewrite(index, regex = false) {
			if (regex) {
				this.regex.splice(index, 1);
				return;
			}
			this.rewrites.splice(index, 1);
		},
		/**
		 * Checks the form for errors and emits
		 * back up to the parent.
		 */
		validateForm() {
			this.$refs['form'].validate((valid) => {
				this.$emit("validate", valid)
			});
		},
		/**
		 * Flattens a rewrite or regex rewrite
		 * to an object
		 * @param obj
		 */
		expandRewrites(obj) {
			let arr = [];
			for (const key in obj) {
				arr.push({'from': key, 'to': obj[key]});
			}
			return arr;
		},
		/**
		 * Reduces the rewrites and regex rewrites to a
		 * flattened array ob objects for the API to
		 * store.
		 * @param rewrites
		 */
		flattenRewrites(rewrites) {
			if (rewrites.length) {
				return rewrites.reduce((obj, item) => (obj[item.from] = item.to, obj) ,{});
			}
			return {};
		}
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Item
// =========================================================================

.item {
	padding: 0 1rem;

	&-rewrite {

		p {
			margin-bottom: 0;
		}
	}
}

</style>
