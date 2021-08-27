<!-- =====================
	New Menu Modal
	===================== -->
<template>
	<!-- Body -->
	<el-form class="item" :model="form" size="small" ref="form" label-width="120px" label-position="left">
		<!-- Name -->
		{{ test }}
		<el-form-item label="Name" prop="name" :rules="{ required: true, message: 'Enter a Name.', trigger: 'blur' }">
			<el-input placeholder="Add a name" v-model="test.name"></el-input>
		</el-form-item>
<!--		&lt;!&ndash; Path &ndash;&gt;-->
<!--		<el-form-item label="Path" prop="path" :rules="{ required: true, message: 'Enter a Path.', trigger: 'blur' }">-->
<!--			<el-input placeholder="Add a path" v-model="proxy.path"></el-input>-->
<!--		</el-form-item>-->
<!--		&lt;!&ndash; Host &ndash;&gt;-->
<!--		<el-form-item label="Host" prop="host" :rules="{ required: true, message: 'Enter a Host.', trigger: 'blur' }">-->
<!--			<el-input placeholder="Add a host" v-model="proxy.host"></el-input>-->
<!--		</el-form-item>-->
<!--		&lt;!&ndash; Rewrites &ndash;&gt;-->
<!--		<div class="item-rewrite">-->
<!--			<el-form-item label="Rewrites" prop="rewrites">-->
<!--				<el-button size="mini" icon="el-icon-plus" @click="addRewrite()"></el-button>-->
<!--			</el-form-item>-->
<!--			<el-table v-if="proxy.rewrite && proxy['rewrite'].length" size="mini" :data="proxy.rewrite" border style="width: 100%; margin: 10px 0;">-->
<!--				&lt;!&ndash; From Path &ndash;&gt;-->
<!--				<el-table-column prop="from" label="From Path">-->
<!--					<template slot-scope="scope">-->
<!--						<el-input placeholder="e.g. /api/*" v-model="scope.row.from" size="mini"></el-input>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--				&lt;!&ndash; To Path &ndash;&gt;-->
<!--				<el-table-column prop="to" label="To Path">-->
<!--					<template slot-scope="scope">-->
<!--						<el-input placeholder="e.g. /$1" v-model="scope.row.to" size="mini"></el-input>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--				&lt;!&ndash; Delete &ndash;&gt;-->
<!--				<el-table-column label="Actions" width="74px" align="right">-->
<!--					<template slot-scope="scope">-->
<!--						<el-button icon="el-icon-delete" size="mini" type="danger"  @click="handleDeleteRewrite(scope.$index)"></el-button>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--			</el-table>-->
<!--			<span v-else>-->
<!--			<p>No rewrites</p>-->
<!--		</span>-->
<!--		</div>-->
<!--		&lt;!&ndash; Regex &ndash;&gt;-->
<!--		<div class="item-rewrite">-->
<!--			<el-form-item label="Regex Rewrites" prop="regex">-->
<!--				<el-button size="mini" icon="el-icon-plus" @click="addRewrite(true)"></el-button>-->
<!--			</el-form-item>-->
<!--			<el-table v-if="proxy.rewrite_regex && proxy.rewrite_regex.length" size="mini" :data="proxy.rewrite_regex" border style="width: 100%; margin-top: 10px">-->
<!--				&lt;!&ndash; From Path &ndash;&gt;-->
<!--				<el-table-column prop="from" label="From Path">-->
<!--					<template slot-scope="scope">-->
<!--						<el-input placeholder="e.g. /old/[0.9]+/" v-model="scope.row.from" size="mini"></el-input>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--				&lt;!&ndash; To Path &ndash;&gt;-->
<!--				<el-table-column prop="to" label="To Path">-->
<!--					<template slot-scope="scope">-->
<!--						<el-input placeholder="e.g. /new" v-model="scope.row.to" size="mini"></el-input>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--				&lt;!&ndash; Delete &ndash;&gt;-->
<!--				<el-table-column label="Actions" width="74px" align="right">-->
<!--					<template slot-scope="scope">-->
<!--						<el-button icon="el-icon-delete" size="mini" type="danger"  @click="handleDeleteRewrite(scope.$index, true)"></el-button>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--			</el-table>-->
<!--			<span v-else>-->
<!--			<p>No regex rewrites</p>-->
<!--		</span>-->
<!--		</div>-->
	</el-form><!-- /Body -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>
export default {
	name: "ProxyItem",
	props: {
		// item: {
		// 	type: Object,
		// },
	},
	data: () => ({
		isValid: false,
		form: {},
		item: {
			name: "",
			path: "",
			host: "",
			rewrite: [],
			rewrite_regex: [],
		},
		itemTest: {},
	}),
	methods: {
		/**
		 * Adds a rewrite regex to the proxy.
		 * @param regex
		 */
		addRewrite(regex = false) {
			const key = regex ? "rewrite_regex" : "rewrite";
			this.proxy[key].push({
				from: "",
				to: "",
			});
		},
		/**
		 *
		 * @param index
		 * @param regex = false
		 */
		handleDeleteRewrite(index, regex = false) {
			const key = regex ? "rewrite_regex" : "rewrite";
			this.proxy[key].splice(index, 1)
		},
		/**
		 *
		 */
		validate() {
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
		 * Loops over the proxies and assigns the
		 * rewrites and regex rewrites to an
		 * array ready for processing.
		 * @param proxy
		 */
		expandProxy(proxy) {
			proxy.rewrite = this.expandRewrites(proxy.rewrite);
			proxy.rewrite_regex = this.expandRewrites(proxy.rewrite_regex);
			return proxy;
		},
		/**
		 * Reduces the rewrites and regex rewrites to a
		 * flattened array ob objects for the API to
		 * store.
		 * @param proxy
		 */
		flattenRewrites(proxy) {
			const rewrites = proxy.rewrite;
			if (rewrites.length) {
				proxy.rewrites = rewrites.reduce((obj, item) => (obj[item.from] = item.to, obj) ,{});
			} else {
				proxy.rewrites = {};
			}
			const regex = proxy.rewrite_regex;
			if (regex.length) {
				proxy.rewrite_regex = regex.reduce((obj, item) => (obj[item.from] = item.to, obj) ,{});
			} else {
				proxy.rewrite_regex = {};
			}
			return proxy;
		}
	},
	computed: {
		/**
		 *
		 */
		proxy: {
			get() {
				return this.item;
			},
			set(el) {
				console.log("hereee")
				console.log(el);
			}
		},
		test: {
			get() {
				return this.itemTest;
			},
			set(el) {
				console.log(el, "here");
			}
		}
	}
}
</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Item
// =========================================================================


.item {


}

</style>
