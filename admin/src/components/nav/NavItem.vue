<!-- =====================
	Nav Item
	===================== -->
<template>
	<el-card class="item box-card" shadow="never" :class="{ 'item-disabled' : disabled }">
		<el-collapse v-model="activeItem">
			<el-collapse-item :disabled="disabled" :title="item.text" class="item-collapse" name="item">
				<!-- Header -->
				<template class="item-header" #title>
					{{ item.text }}
				</template>
				<el-form ref="form" :model="item" :rules="rules" label-position="top">
					<!-- Link Text -->
					<el-form-item label="Link Text" prop="name">
						<el-input placeholder="Link Text*" label="Link Text*" v-model="item.text" clearable></el-input>
					</el-form-item>
					<!-- Link Text -->
					<el-form-item label="Link Title">
						<el-input placeholder="Title" label="Link Title" v-model="item.title" clearable></el-input>
					</el-form-item>
					<!-- Rel -->
					<el-form-item label="Rel Attribute">
						<el-input placeholder="Rel" label="Rel Attribute" v-model="item.rel" clearable></el-input>
					</el-form-item>
					<!-- Description -->
					<el-form-item label="Description" class="form-group">
						<el-input type="textarea" :rows="4" placeholder="Description" v-model="item.description" resize="none"></el-input>
					</el-form-item>
					<!-- Open Tab -->
					<el-form-item>
						<el-checkbox v-model="item['new_tab']">Open in new tab</el-checkbox>
					</el-form-item>
					<!-- Toolbar -->
					<div class="item-toolbar">
						<el-link class="item-toolbar-link" type="danger" @click="removeItem">Remove</el-link>
						<el-link class="item-toolbar-link" @click="collapse">Cancel</el-link>
					</div>
				</el-form>
			</el-collapse-item>
		</el-collapse>
	</el-card>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "NavItem",
	props: {
		item: {
			type: Object,
		},
		disabled: {
			type: Boolean,
			default: false,
		}
	},
	data: () => ({
		activeItem: "",
		rules: {
			name: [
				{ required: true, message: 'Enter link text for the menu item', trigger: 'blur' },
			],
		},
	}),
	methods: {
		/*
		 * removeItem()
		 * Remove an item from the menu, emits
		 * back up to the parent.
		 */
		removeItem() {
			this.$emit("remove", this.item);
		},
		/*
		 * collapse()
		 * Collapses the current item (when cancel)
		 * is clicked.
		 */
		collapse() {
			this.activeItem = ""
		},
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
		width: 50%;
		background-color: $white;

		// Props
		// =========================================================================

		::v-deep {

			.el-collapse,
			.el-collapse-item__wrap,
			.el-collapse-item__header {
				border: none !important;
			}

			.el-collapse-item__content {
				padding-bottom: 1rem;
				user-select: none;
			}

			.el-card__body {
				padding: 0;
			}
		}

		// Toolbar
		// =========================================================================

		&-toolbar {
			display: flex;


			&-link:first-child {
				margin-right: 10px;
			}
		}

		// Disabled
		// =========================================================================

		&-disabled {
			pointer-events: none;
			user-select: none;
		}
	}

</style>
