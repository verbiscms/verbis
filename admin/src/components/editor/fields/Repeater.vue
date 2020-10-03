<!-- =====================
	Field - Repeater
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<div class="repeater-item" v-for="(repeater, repeaterIndex) in repeaters" :key="repeaterIndex">
			<button class="btn btn-orange" @click="deleteRow(repeaterIndex)">Delete</button>
			<div v-for="(field, fieldIndex) in getSubFields" :key="field.uuid">
				rIndex{{ repeaterIndex }}
				fIndex {{ fieldIndex }}
				<!-- Text -->
				<FieldText v-if="field.type === 'text'" :layout="field" @update:text="updateField($event, repeaterIndex, fieldIndex, field.name)"></FieldText>
				<!-- Textarea -->
				<FieldTextarea v-else-if="field.type === 'textarea'" :layout="field" @update="updateField($event, repeaterIndex, fieldIndex, field.name)"></FieldTextarea>
				<!-- Number -->
				<FieldNumber v-if="field.type === 'number'" :layout="field" @update:number="updateField($event, repeaterIndex, fieldIndex, field.name)"></FieldNumber>
				<!-- Range -->
				<FieldRange v-if="field.type === 'range'" :layout="field" @update:range="updateField($event, repeaterIndex, fieldIndex, field.name)"></FieldRange>
				<!-- Email -->
				<FieldEmail v-if="field.type === 'email'" :layout="field" @update:email="updateField($event, repeaterIndex, fieldIndex, field.name)"></FieldEmail>
				<!-- Richtext -->
				<FieldRichText v-else-if="field.type === 'richtext'"></FieldRichText>
			</div>
		</div>

		<div class="repeater-btn">
			<button class="btn btn-blue" @click="addRow">Add row</button>
		</div>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->

		{{ fields }}
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

export default {
	name: "FieldRepeater",
	props: {
		layout: Object,
	},
	components: {
		FieldText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldRichText,
	},
	data: () => ({
		fields: {},
		errors: [],
		repeaters: [],
		focused: false,
	}),
	mounted() {
		this.addRow()
	},
	methods: {
		updateField(e, repeaterIndex, fieldIndex, field) {
			console.log(fieldIndex)
			console.log(field)
			this.$set(this.fields, field, e)
			//this.$emit("update:repeater", this.fields)
		},
		deleteRow(index) {
			this.repeaters.splice(index, 1);
		},
		addRow() {
			this.repeaters.push(this.getSubFields)
		}
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getSubFields() {
			return this.layout['sub_fields'];
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