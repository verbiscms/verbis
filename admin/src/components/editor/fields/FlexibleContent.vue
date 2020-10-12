<!-- =====================
	Field - Flexible Content
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="flexible">
		<draggable @start="drag=true" :list="fields" :group="fields" :sort="true" handle=".flexible-handle">
			<div class="card card-margin-small" v-for="(group, groupIndex) in getFields" :key="groupIndex">
				<div class="card-header">
					<h4>{{ group.type }}</h4>
					<div class="card-controls">
						<i class="feather feather-trash-2" @click="deleteRow(groupIndex)"></i>
						<i class="feather feather-arrow-up" @click="moveUp(groupIndex)"></i>
						<i class="feather feather-arrow-down" @click="moveDown(groupIndex)"></i>
						<i class="flexible-handle fal fa-arrows"></i>
					</div>
				</div><!-- /Card Header -->
				<div class="card-body card-body-border-bottom" v-for="(layout, layoutKey) in getSubFields(group.type)" :key="layoutKey">
					<!-- Field Title -->
					<div class="field-title">
						<h4>{{ layout.label }}</h4>
						<p>{{ layout.instructions }}</p>
					</div>
					<!-- =====================
						Basic
						===================== -->
					<!-- Text -->
					<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldEmail>
					<!-- Url -->
					<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldUrl>
					<!-- Password -->
					<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldPassword>
					<!-- =====================
						Content
						===================== -->
					<!-- Richtext -->
					<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldRichText>
					<!-- =====================
						Choice
						===================== -->
					<!-- Select -->
					<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldSelect>
					<!-- Multi Select -->
					<FieldTags v-else-if="layout.type === 'multi_select'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldTags>
					<!-- Checkbox -->
					<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldCheckbox>
					<!-- Radio -->
					<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldRadio>
					<!-- Button Group -->
					<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldButtonGroup>
					<!-- =====================
						Relational
						===================== -->
					<!-- Post Object -->
					<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldPost>
					<!-- User -->
					<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldUser>
					<!-- =====================
						Layout
						===================== -->
					<!-- Repeater -->
					<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldRepeater>
					<!-- Flexible -->
					<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[groupIndex]['fields'][layout.name]" :error-trigger="errorTrigger"></FieldFlexible>
				</div><!-- /Card Body -->
			</div><!-- /Card -->
		</draggable>
		<div class="field-btn">
			<div class="popover-cont">
				<button class="btn btn-blue"><i class="fal fa-plus-circle"></i>Add Row</button>
				<div class="popover">
					<span class="popover-triangle"></span>
					<div class="popover-item" v-for="(group) in getLayouts" :key="group.uuid" @click="addRow(group.name)">{{ group.name }}</div>
				</div>
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
		errorTrigger: {
			type: Boolean,
			default: false,
		},
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
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.flexible.closest(".collapse-content"));
			});
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

// Field
// =========================================================================

.field {

	&-group {
		margin-bottom: 1rem;
		border: 1px solid $grey-light;
		padding: 15px 0;
		border-radius: 6px;
		background-color: $white;
	}

	&-content {
		margin-bottom: 1rem;
		padding: 0 20px;

		&:last-of-type {
			margin-bottom: 0;
		}
	}

	&-title {
		color: $secondary;
	}
}

	.flexible {


		// Header
		// =========================================================================

		&-header {
			display: flex;
			justify-content: space-between;
			width: 100%;
			margin-bottom: 1rem;
			padding: 0 20px 15px 20px;
			border-bottom: 1px solid $grey-light;

			h4 {
				margin-bottom: 0;
			}
		}



	}
</style>