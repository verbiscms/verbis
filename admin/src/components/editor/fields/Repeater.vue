<!-- =====================
	Field - Repeater
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<draggable @start="drag=true" :list="fields" :group="fields" :sort="true" handle=".repeater-handle">
			<div class="repeater-item" v-for="(repeater, repeaterIndex) in getFields" :key="repeaterIndex">
				<div class="field-controls">
					<i class="feather icon-trash-2" @click="deleteRow(repeaterIndex)"></i>
					<i class="feather icon-arrow-up" @click="moveUp(repeaterIndex)"></i>
					<i class="feather icon-arrow-down" @click="moveDown(repeaterIndex)"></i>
					<i class="repeater-handle fal fa-arrows"></i>
				</div>
				<div v-for="(layout, layoutIndex) in getSubFields" :key="layoutIndex">
					<!-- Text -->
					<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldEmail>
					<!-- Richtext -->
					<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]"></FieldRichText>
					<!-- Repeater -->
					<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRepeater>
				</div>
			</div>
		</draggable>
		<div class="repeater-btn">
			<button class="btn btn-blue" @click="addRow">Add row</button>
		</div>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->
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
	name: "FieldRepeater",
	props: {
		layout: Object,
		fields: Array,
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
		repeaters: [],
	}),
	mounted() {
		if (this.repeaterFields !== undefined) {
			this.repeaterFields = this.getFields
		}
	},
	methods: {
		deleteRow(index) {
			this.fields.splice(index, 1);
			this.repeaters.splice(index, 1);
		},
		addRow() {
			this.repeaterFields.push({})
		},
		moveUp(index) {
			this.moveItem(index, index - 1)
		},
		moveDown(index) {
			this.moveItem(index, index + 1)
		},
		moveItem(from, to) {
			this.repeaterFields.splice(to, 0, this.repeaterFields.splice(from, 1)[0]);
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getSubFields() {
			return this.layout['sub_fields'];
		},
		getFields() {
			return this.fields
		},
		repeaterFields: {
			get() {
				return this.fields === undefined ? [{}] : this.fields
			},
			set() {
				this.$emit("update:fields", this.repeaterFields)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


	.repeater {

		// Item
		// =========================================================================

		&-item {
			border: 2px solid $grey-light;
			padding: 10px;
			margin-bottom: 1rem;
			border-radius: 4px;
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