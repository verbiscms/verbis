<!-- =====================
	New Menu Modal
	===================== -->
<template>
	<el-dialog :visible.sync="show" width="30%">
		<!-- Title -->
		<div slot="title">
			<h2 class="mb-0">Create a menu</h2>
			<p>Select a menu location and assign the new menu a name.</p>
		</div>
		<!-- Form -->
		<el-form :model="menu" ref="newMenu" label-position="left" label-width="auto ">
			<!-- Name -->
			<el-form-item label="Name" prop="name" :rules="{ required: true, message: 'Enter a Menu Name.', trigger: 'blur' }">
				<el-input placeholder="Name" v-model="menu.name"></el-input>
			</el-form-item>
			<!-- Location -->
			<el-form-item label="Location" prop="id" :rules="{ required: true, message: 'Enter a Menu Location.', trigger: 'change' }">
				<el-select v-model="menu.id" placeholder="Select">
					<el-option v-for="menu in menus" :key="menu.id" :label="menu.name" :value="menu.id"></el-option>
				</el-select>
			</el-form-item>
		</el-form>
		<!-- Footer -->
		<span slot="footer" class="dialog-footer">
			<el-button @click="show = false">Cancel</el-button>
			<el-button type="primary" @click="create">Create Menu</el-button>
		</span>
	</el-dialog>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "MenuNewModal",
	props: {
		visible: {
			type: Boolean,
			default: false,
		},
		menus: {
			type: Array,
		}
	},
	data: () => ({
		menu: {},
	}),
	methods: {
		/**
		 * Checks to see if the form data is valid then
		 * updates the parent with the new menu
		 * information. Resets the menu data
		 * upon success.
		 */
		create() {
			this.$refs['newMenu'].validate((valid) => {
				if (valid) {
					this.$emit("create", this.menu);
					this.menu = {};
				}
			});
		}
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
<style lang="scss"></style>
