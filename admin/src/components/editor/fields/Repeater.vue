<!-- =====================
	Field - Repeater
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="repeater">
		<pre>{{ repeaterFields }}</pre>
		<draggable @start="drag=true" :list="fields" :group="fields" :sort="true" handle=".repeater-handle">
			<div class="repeater" v-for="(repeater, repeaterIndex) in getFields" :key="repeaterIndex">
				<div class="card-header">
					<h4>{{ layout.label }} item {{ repeaterIndex + 1 }}</h4>
					<div class="card-controls">
						<i class="feather feather-trash-2" @click="deleteRow(repeaterIndex)"></i>
						<i class="feather feather-arrow-up" @click="moveUp(repeaterIndex)"></i>
						<i class="feather feather-arrow-down" @click="moveDown(repeaterIndex)"></i>
						<i class="repeater-handle fal fa-arrows"></i>
					</div>
				</div>
				<div class="card-body" v-for="(layout, layoutIndex) in getSubFields" :key="layoutIndex">
					<!-- Field Title -->
					<div class="field-title">
						<h4>{{ layout.label }}</h4>
						<p>{{ layout.instructions }}</p>
					</div>
					<!-- =====================
						Basic
						===================== -->
					<!-- Text -->
					<FieldText v-if="layout.type === 'text'" :layout="layout" :fields="getValue(layout)" @update:fields="pushValue($event, layout)" :error-trigger="errorTrigger"></FieldText>
<!--					&lt;!&ndash; Textarea &ndash;&gt;-->
<!--					<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldTextarea>-->
<!--					&lt;!&ndash; Number &ndash;&gt;-->
<!--					<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldNumber>-->
<!--					&lt;!&ndash; Range &ndash;&gt;-->
<!--					<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldRange>-->
<!--					&lt;!&ndash; Email &ndash;&gt;-->
<!--					<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldEmail>-->
<!--					&lt;!&ndash; Url &ndash;&gt;-->
<!--					<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldUrl>-->
<!--					&lt;!&ndash; Password &ndash;&gt;-->
<!--					<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldPassword>-->
<!--					&lt;!&ndash; =====================-->
<!--						Content-->
<!--						===================== &ndash;&gt;-->
<!--					&lt;!&ndash; Richtext &ndash;&gt;-->
<!--					<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldRichText>-->
<!--					&lt;!&ndash; Image &ndash;&gt;-->
<!--					<FieldImage v-else-if="layout.type === 'image'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldImage>-->
<!--					&lt;!&ndash; =====================-->
<!--						Choice-->
<!--						===================== &ndash;&gt;-->
<!--					&lt;!&ndash; Select &ndash;&gt;-->
<!--					<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldSelect>-->
<!--					&lt;!&ndash; Multi Select &ndash;&gt;-->
<!--					<FieldTags v-else-if="layout.type === 'multi_select'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldTags>-->
<!--					&lt;!&ndash; Checkbox &ndash;&gt;-->
<!--					<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldCheckbox>-->
<!--					&lt;!&ndash; Radio &ndash;&gt;-->
<!--					<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldRadio>-->
<!--					&lt;!&ndash; Button Group &ndash;&gt;-->
<!--					<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldButtonGroup>-->
<!--					&lt;!&ndash; =====================-->
<!--						Relational-->
<!--						===================== &ndash;&gt;-->
<!--					&lt;!&ndash; Post Object &ndash;&gt;-->
<!--					<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldPost>-->
<!--					&lt;!&ndash; User &ndash;&gt;-->
<!--					<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldUser>-->
					<!-- =====================
						Layout
						===================== -->
					<!-- Repeater -->
<!--					<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldRepeater>-->
<!--					&lt;!&ndash; Flexible &ndash;&gt;-->
<!--					<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[repeaterIndex][layout.name]" :error-trigger="errorTrigger"></FieldFlexible>-->
				</div><!-- /Card Body -->
			</div><!-- /Card -->
		</draggable>
		<div class="field-btn">
			<button class="btn btn-blue" @click="addRow"><i class="fal fa-plus-circle"></i>Add Row</button>
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
// import FieldTextarea from "@/components/editor/fields/Textarea";
// import FieldNumber from "@/components/editor/fields/Number";
// import FieldRange from "@/components/editor/fields/Range";
// import FieldEmail from "@/components/editor/fields/Email";
// import FieldImage from "@/components/editor/fields/Image";
// import FieldRichText from "@/components/editor/fields/RichText";
// import FieldRepeater from "@/components/editor/fields/Repeater";
import draggable from 'vuedraggable'

export default {
	name: "FieldRepeater",
	props: {
		layout: Object,
		fields: Array,
		errorTrigger: {
			type: Boolean,
			default: false,
		},
	},
	components: {
		FieldText,
		// FieldTextarea,
		// FieldNumber,
		// FieldRange,
		// FieldEmail,
		// FieldRichText,
		// FieldRepeater,
		// FieldImage,
		draggable,
	},
	data: () => ({
		errors: [],
		repeaters: [],
		repeaterFields: [],
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
			console.log("cliocked");
			this.repeaterFields.push({})
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.repeater.closest(".collapse-content"));
			});
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
		pushValue(value, layout) {
			const fieldData = this.fields.filter(function(field) {
				return field.uuid === layout.uuid;
			});

			if (fieldData.length) {
				fieldData[0].value = value;
				return
			}

			this.fields.push({
				uuid: layout.uuid,
				value: value,
				name: layout.name,
				type: layout.type,
			});
		},
		getValue(layout) {
			const fieldData = this.fields.filter(function(field) {
				return field.uuid === layout.uuid;
			});
			if (fieldData.length) {
				return fieldData[0].value;
			}
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
		// repeaterFields: {
		// 	get() {
		// 		return this.fields === undefined ? [{}] : this.fields
		// 	},
		// 	set() {
		// 		this.$emit("update:fields", this.repeaterFields)
		// 	}
		// }
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	.repeater {
		border: 1px solid $grey-light;
		margin-bottom: 1.6rem;
		border-radius: 6px;

		// Item
		// =========================================================================

		&-item {
			border: 1px solid $grey-light;
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