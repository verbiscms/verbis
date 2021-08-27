<!-- =====================
	New Menu Modal
	===================== -->
<template>
	<!-- Body -->
	<el-form class="item" :model="proxy" size="small" ref="proxiesForm" label-width="120px" label-position="left">
		<!-- Name -->
		<el-form-item label="Name" prop="name" :rules="{ required: true, message: 'Enter a Name.', trigger: 'blur' }">
			<el-input placeholder="Add a name" v-model="proxy.name"></el-input>
		</el-form-item>
		<!-- Path -->
		<el-form-item label="Path" prop="path" :rules="{ required: true, message: 'Enter a Path.', trigger: 'blur' }">
			<el-input placeholder="Add a path" v-model="proxy.path"></el-input>
		</el-form-item>
		<!-- Host -->
		<el-form-item label="Host" prop="host" :rules="{ required: true, message: 'Enter a Host.', trigger: 'blur' }">
			<el-input placeholder="Add a host" v-model="proxy.host"></el-input>
		</el-form-item>
		<!-- Rewrites -->
		<div class="item-rewrite">
			<el-form-item label="Rewrites" prop="rewrites">
				<el-button size="mini" icon="el-icon-plus" @click="addRewrite(index)"></el-button>
			</el-form-item>
			<el-table v-if="proxy.rewrite && proxy['rewrite'].length" size="mini" :data="proxy.rewrite" border style="width: 100%; margin: 10px 0;">
				<!-- From Path -->
				<el-table-column prop="from" label="From Path">
					<template slot-scope="scope">
						<el-input placeholder="e.g. /api/*" v-model="scope.row.from" size="mini"></el-input>
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
						<el-button icon="el-icon-delete" size="mini" type="danger"  @click="handleDeleteRewrite(index, scope.$index)"></el-button>
					</template>
				</el-table-column>
			</el-table>
			<span v-else>
				<p>No rewrites</p>
			</span>
		</div>
		<!-- Regex -->
		<div class="item-rewrite">
			<el-form-item label="Regex Rewrites" prop="regex">
				<el-button size="mini" icon="el-icon-plus" @click="addRewrite(index, true)"></el-button>
			</el-form-item>
			<el-table v-if="proxy.rewrite_regex && proxy.rewrite_regex.length" size="mini" :data="proxy.rewrite_regex" border style="width: 100%; margin-top: 10px">
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
						<el-button icon="el-icon-delete" size="mini" type="danger"  @click="handleDeleteRewrite(index, scope.$index, true)"></el-button>
					</template>
				</el-table-column>
			</el-table>
			<span v-else>
				<p>No regex rewrites</p>
			</span>
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
		},
	},
	data: () => ({

	}),
	methods: {
		/**
		 * Adds a rewrite regex to the proxy.
		 * @param index
		 * @param regex
		 */
		addRewrite(index, regex = false) {
			const key = regex ? "rewrite_regex" : "rewrite";
			this.proxy[index][key].push({
				from: "",
				to: "",
			});
		},
		handleDeleteRewrite(proxyIndex, rowIndex, regex = false) {
			const key = regex ? "rewrite_regex" : "rewrite";
			this.proxies[proxyIndex][key].splice(rowIndex, 1)
		},

	},
	computed: {
		proxy: {
			get() {
				return this.item;
			},
			set(el) {
				console.log(el);
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
