<!-- =====================
	Field - Flexible Content
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<draggable @start="drag=true" :list="fields" :group="fields" :sort="true" handle=".repeater-handle">
			<div class="field-group" v-for="(group, groupIndex) in getFields" :key="groupIndex">
				<div class="field-controls">
					<i class="feather icon-trash-2" @click="deleteRow(groupIndex)"></i>
					<i class="feather icon-arrow-up" @click="moveUp(groupIndex)"></i>
					<i class="feather icon-arrow-down" @click="moveDown(groupIndex)"></i>
					<i class="repeater-handle fal fa-arrows"></i>
				</div>
				<div class="field-content" v-for="(layout, layoutKey) in getSubFields(group.type)" :key="layoutKey">
					<!-- Text -->
					<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldEmail>
					<!-- Richtext -->
					<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldRichText>
					<!-- Repeater -->
					<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]"></FieldRepeater>
				</div>
			</div>
		</draggable>
		<div class="popover-cont">
			<button class="btn btn-blue"><i class="fal fa-plus-circle"></i>Add Row</button>
			<div class="popover">
				<span class="popover-triangle"></span>
				<div class="popover-item" v-for="(group) in getLayouts" :key="group.uuid" @click="addRow(group.name)">{{ group.name }}</div>
			</div>
		</div>
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import FieldText from "@/components/editor/fields/Text";
import FieldTextarea from "@/components/editor/fields/Textarea";
import FieldNumber from "@/components/editor/fields/Number";
import FieldRange from "@/components/editor/fields/Range";
import FieldEmail from "@/components/editor/fields/Email";
import FieldRichText from "@/components/editor/fields/RichText";
import FieldRepeater from "@/components/editor/fields/Repeater";
import draggable from 'vuedraggable'


export default {
	name: "FieldFlexible",
	props: {
		layout: Object,
		fields: [Array, Object],
	},
	components: {
		FieldText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldRichText,
		FieldRepeater,
		draggable,
	},
	data: () => ({
		errors: [],
		layouts: [],
		showPopover: false,
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
	//background-color: rgba(red, 0.4);
	margin-bottom: 1rem;
	border: 2px solid $secondary;
	padding: 15px;
	border-radius: 4px;
}

	.flexible {

		// Btn
		// =========================================================================

		&-btn {

		}

	}
</style>