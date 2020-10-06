<!-- =====================
	Field - Flexible Content
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="field-group" v-for="(group, groupIndex) in getFields" :key="groupIndex">
			<div class="field-content" v-for="(layout, layoutKey) in getSubFields(group.type)" :key="layoutKey">
				<!-- Text -->
				<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldText>
			</div>
		</div>
		<div class="repeater-btn" v-for="(group) in getLayouts" :key="group.uuid">
			<button class="btn btn-blue" @click="addRow(group.name)">Add {{ group.name }}</button>
		</div>
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import FieldText from "@/components/editor/fields/Text";

export default {
	name: "FieldFlexible",
	props: {
		layout: Object,
		fields: [Array, Object],
	},
	components: {
		FieldText,
	},
	data: () => ({
		errors: [],
		layouts: [],
	}),
	mounted() {
		if (this.layoutFields !== undefined) {
			this.layoutFields = this.getFields
		}
	},
	methods: {
		deleteRow(index) {
			this.fields.splice(index, 1);
			this.layouts.splice(index, 1);
		},
		addRow(key) {
			this.layouts.push(this.getLayouts[key])
			const subFields = this.getLayouts[key]['sub_fields']

			let temp = {
				type: key,
				fields: {},
			}
			for (const fieldKey in subFields) {
				temp['fields'][fieldKey] = "";
			}

			this.layoutFields.push(temp)
		},
		moveUp(index) {
			this.moveItem(index, index - 1)
		},
		moveDown(index) {
			this.moveItem(index, index + 1)
		},
		moveItem(from, to) {
			this.layoutFields.splice(to, 0, this.layoutFields.splice(from, 1)[0]);
		},
		getSubFields(key) {
			const layout = this.getLayouts[key],
				subFields = layout['sub_fields'];
			return subFields;
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getLayouts() {
			return this.layout['layouts'];
		},
		getFields() {
			return this.fields
		},
		layoutFields: {
			get() {
				return this.fields === undefined ? [] : this.fields
			},
			set() {
				this.$emit("update:fields", this.layoutFields)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.field-group {
	background-color: rgba(red, 0.4);
	margin-bottom: 1rem;
}

	.repeater {


		// Item
		// =========================================================================

		&-item {
			border: 2px solid $grey-light;
			padding: 10px;
		}

		// Button
		// =========================================================================

		&-btn {
			display: flex;
			justify-content: flex-end;
			margin-top: 1rem;
		}
	}
</style>