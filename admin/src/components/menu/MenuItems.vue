<!-- =====================
	Menu Items
	===================== -->
<template>
	<!-- Menu Items -->
	<div class="items">
		<el-card class="items-card" :body-style="{ padding: '0px' }" shadow="never">
			<!-- Nestable Items -->
			<div class="items-nestable">
				<vue-nestable
					v-model="elements"
					:max-depth="10"
					:threshold="30"
					children-prop="children"
					@input="isDragging = true"
					@change="handleAfterDrag">
					<vue-nestable-handle slot-scope="{ item }" :item="item">
						<!-- Specific Menu Item -->
						<MenuItem
							@remove="removeItem(item.id)"
							@checked="handleChecked(item)"
							:disabled="isDragging"
							:item="item"
							:bulk="doingBulk">
						</MenuItem>
					</vue-nestable-handle>
				</vue-nestable>
			</div><!-- /Nestable Items -->
			<!-- Footer -->
			<div class="items-footer">
				<!-- Bulk Select -->
				<el-checkbox v-model="doingBulk">Bulk select</el-checkbox>
				<!--		<el-link class="ml-2" type="danger" @click="this.deleteMenu">Remove menu</el-link>-->
				<el-link :disabled="checked.length === 0" class="ml-2" @click="removeItems()" v-if="doingBulk">Remove selected items</el-link>
			</div><!-- /Footer -->
		</el-card>
	</div><!-- /Menu Items -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import MenuItem from "@/components/menu/MenuItem";

export default {
	name: "MenuItems",
	components: {
		MenuItem,
	},
	props: {
		items: {
			type: Array,
			default: () => [],
		},
	},
	data: () => ({
		doingBulk: false,
		isDragging: false,
		checked: [],
	}),
	created() {
		setTimeout(() => {
			this.isDragging = false;
		}, 200);
	},
	methods: {
		/**
		 * removeItem()
		 * Removes a singular item from the hierarchy.
		 * @param id
		 */
		removeItem(id) {
			this.elements = this.filterItems(this.elements, [id]);
		},
		/**
		 * removeItems()
		 * Removes multiple items from the hierarchy.
		 */
		removeItems() {
			this.elements = this.filterItems(this.elements, this.checked);
			this.checked = [];
		},
		/**
		 * filterItems()
		 * Iterates over the data and finds a menu item by
		 * ID and continues to delete it and any
		 * children associated with it.
		 * @param data
		 * @param ids
		 */
		filterItems(data, ids) {
			return data.filter(item => {
				if (item.children) {
					item.children = this.filterItems(item.children, ids);
				}
				return !ids.includes(item.id);
			})
		},
		/**
		 * handleChecked()
		 * Pushes or removes the item ID to the checked array
		 * when a child's checkbox is ticked.
		 * @param item
		 */
		handleChecked(item) {
			const index = this.checked.indexOf(item.id);
			if (index > -1) {
				this.checked.splice(index, 1);
				return;
			}
			this.checked.push(item.id);
		},
		/**
		 * handleAfterDrag()
		 */
		handleAfterDrag() {
			this.isDragging = false;
			// this.$nextTick(() => {
			// 	this.isDragging = false;
			// }, 500);
		},
	},
	computed: {
		/**
		 * Computed prop for emitting the items to
		 * the parent.
		 */
		elements: {
			get() {
				return this.items;
			},
			set(val) {
				this.$emit("update:items", val)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Items
// =========================================================================

.items {


	&-nestable {
		padding: 1.4rem;
	}

	&-footer {
		padding: 1rem 1.6rem;
	}

	// Card
	// =========================================================================

	&-card {
		margin-bottom: 1rem;
	}

	// Footer
	// =========================================================================

	&-footer {
		border-top: 1px solid rgba($grey, 0.3);
	}
}


</style>
