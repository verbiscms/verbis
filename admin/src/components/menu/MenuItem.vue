<!-- =====================
	Nav Item
	===================== -->
<template>
	<el-card class="item box-card" shadow="never" :class="{ 'item-disabled' : disabled }">
		<el-collapse :disabled="disabled" v-model="activeItem">
			<el-collapse-item :title="item.text" class="item-collapse" name="item">
				<!-- Header -->
				<template class="item-header" #title>
					<el-checkbox class="item-header-checkbox" v-if="bulk" v-model="checked" @change="handleCheckChange"></el-checkbox>
					{{ item.text }}
				</template>
				<el-form ref="form" :model="item" :rules="rules" label-position="top">
					<!-- Link Text -->
					<el-form-item label="Link Text" prop="name">
						<el-input placeholder="Link Text*" label="Link Text*" v-model="item.text" clearable></el-input>
					</el-form-item>
					<!-- Link Text -->
					<el-form-item class="mb-0" label="Link Title">
						<el-input placeholder="Title" label="Link Title" v-model="item.title" clearable></el-input>
					</el-form-item>
					<!-- Open Tab -->
					<el-form-item>
						<el-checkbox v-model="item['new_tab']">Open in new tab</el-checkbox>
					</el-form-item>
					<!-- CSS Clases -->
					<el-form-item label="Li Classes">
						<el-tag class="item-li-class-tag" :key="tag" v-for="tag in item['li_classes']" closable :disable-transitions="false" @close="handleClose(tag)">
							{{ tag }}
						</el-tag>
						<el-input v-if="inputVisible" v-model="inputValue" ref="saveLiClasses" size="mini" @keyup.native.enter="handleInputConfirm" @blur="handleInputConfirm">
						</el-input>
						<el-button v-else class="item-li-class-btn" size="small" @click="showInput">+ New class</el-button>
					</el-form-item>
					<!-- Rel -->
					<el-form-item label="Rel Attribute">
						<el-select  v-model="item.rel" multiple collapse-tags placeholder="Rel">
							<el-option v-for="item in REL" :key="item" :label="item" :value="item">
							</el-option>
						</el-select>
					</el-form-item>
					<!-- Description -->
					<el-form-item label="Description" class="form-group">
						<el-input type="textarea" :rows="4" placeholder="Description" v-model="item.description" resize="none"></el-input>
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
	name: "MenuItem",
	props: {
		item: {
			type: Object,
		},
		disabled: {
			type: Boolean,
			default: false,
		},
		bulk: {
			type: Boolean,
			default: false,
		},
	},
	data: () => ({
		REL: [
			'alternate', 'author', 'bookmark', 'external', 'help', 'license', 'next',
			'nofollow', 'noopener', 'noreferrer', 'prev', 'search', 'tag'
		],
		activeItem: "",
		checked: false,
		inputVisible: false,
		inputValue: '',
		rules: {
			name: [
				{ required: true, message: 'Enter link text for the menu item', trigger: 'blur' },
			],
		},
	}),
	watch: {
		/*
		 * Collapse the item if it's in bulk mode.
		 */
		bulk: function() {
			this.collapse();
		}
	},
	methods: {
		/*
		 * Remove an item from the menu, emits
		 * back up to the parent.
		 */
		removeItem() {
			this.$emit("remove", this.item);
		},
		/*
		 * Collapses the current item (when cancel)
		 * is clicked.
		 */
		collapse() {
			this.activeItem = ""
		},
		/*
		 * Updates the parent when the checkbox is
		 * changed.
		 */
		handleCheckChange() {
			console.log("somethings changed");
			this.$emit("checked", this.checked);
		},
		/**
		 * Removes a class from the li_classes array,
		 * when the close button is clicked.
		 * @param tag
		 */
		handleClose(tag) {
			this.item['li_classes'].splice(this.item['li_classes'].indexOf(tag), 1);
		},
		/**
		 * Shows the input for entering <li> classes.
		 */
		showInput() {
			this.inputVisible = true;
			this.$nextTick(() => {
				this.$refs.saveLiClasses.$refs.input.focus();
			});
		},
		/**
		 * handleInputConfirm
		 */
		handleInputConfirm() {
			let inputValue = this.inputValue;
			// Check if the array exists, if not create
			// a new one, required!
			if (!this.item['li_classes']) {
				this.item['li_classes'] = [];
			}
			if (inputValue) {
				this.item['li_classes'].push(inputValue);
			}
			this.inputVisible = false;
			this.inputValue = '';
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
		width: 60%;
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

		// Header
		// =========================================================================

		&-header {

			&-checkbox {
				margin-right: 10px;
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


			* {
				pointer-events: none !important;
				user-select: none !important;
			}
			::v-deep .el-collapse-item__header {
				pointer-events: none;
				user-select: none;
			}
		}

		// Li Class
		// =========================================================================

		&-li-class {

			&-tag {
				margin-right: 5px;
			}
		}
	}

</style>
