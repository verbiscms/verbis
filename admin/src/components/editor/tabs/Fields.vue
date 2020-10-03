<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Text -->
		<div class="field-group" v-for="group in layout" :key="group.uuid">
			<div class="field" v-for="field in group.fields" :key="field.uuid" :style="{ width: field.wrapper['width'] + '%' }">
				<div class="field-wrapper">
					<div class="field-title-cont">
						<span class="field-collapse" v-on:click="isActive = !isActive"></span>
						<div class="field-title">
							<h5>{{ field.label }}</h5>
							<p>{{ field.instructions }}</p>
						</div>
					</div>
					<!-- TODO: If flexible -->
					<div class="field-controls">
						<i class="feather icon-trash-2"></i>
						<i class="feather icon-arrow-up"></i>
						<i class="feather icon-arrow-down"></i>
						<i class="fal fa-arrows"></i>
					</div>
				</div>
				<!--Content -->
				<div class="field-content">
					<!-- Text -->
					<FieldText v-if="field.type === 'text'" :layout="field" @update:text="updateField($event, field.name)"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="field.type === 'textarea'" :layout="field" @update="updateField($event, field.name)"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="field.type === 'number'" :layout="field" @update:number="updateField($event, field.name)"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="field.type === 'range'" :layout="field" @update:range="updateField($event, field.name)"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="field.type === 'email'" :layout="field" @update:email="updateField($event, field.name)"></FieldEmail>
					<!-- Richtext -->
					<FieldRichText v-else-if="field.type === 'richtext'"></FieldRichText>
				</div>
			</div>
		</div>
	</section>
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
	name: "Fields",
	props: {
		layout: Array,
	},
	components: {
		FieldRichText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldText,
	},
	data: () => ({
		computedHeight: 'auto',
		isActive: true,
		fields: {},
	}),
	mounted() {
		this.initHeight()
	},
	methods: {
		initHeight() {
			//this.computedHeight= getComputedStyle(this.$refs['myText']).height;
		},
		updateField(e, field) {
			this.$set(this.fields, field, e)
			this.$emit("update", this.fields)
		},
		parseLogic() {
		// 	//if(field.hasOwnProperty('logic')) {
		// 		if (field.logic) {
		// 			const logic = field.logic;
		// 			const fieldEval = this.fieldValue[logic.field];
		// 			let value = logic.value;
		// 			const operator = logic.operator.toString();
		//
		// 			if (value === 'true') {
		// 				value = true;
		// 			} else if (value === 'false') {
		// 				value = false;
		// 			}
		//
		// 			switch (operator) {
		// 				case '>':
		// 					return fieldEval > value;
		// 				case '<':
		// 					return fieldEval < value;
		// 				case '>=':
		// 					return fieldEval >= value;
		// 				case '<=':
		// 					return fieldEval <= value;
		// 				case '==':
		// 					return fieldEval == value;
		// 				case '!=':
		// 					return fieldEval != value;
		// 				case '===':
		// 					return fieldEval === value;
		// 				case '!==':
		// 					return fieldEval !== value;
		// 			}
		// 		} else {
		// 			return false;
		// 		}
		// 	} else {
		// 		return false;
		// 	}
		//
		// },
			}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>