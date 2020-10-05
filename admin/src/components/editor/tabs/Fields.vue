<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Text -->
		<div class="field-group" v-for="group in layout" :key="group.uuid">
			<h4>{{ group.title }}</h4>
			<div class="field" v-for="layout in group.fields" :key="layout.uuid" :style="{ width: layout.wrapper['width'] + '%' }">
				<div class="field-wrapper">
					<div class="field-title-cont">
						<span class="field-collapse" v-on:click="isActive = !isActive"></span>
						<div class="field-title">
							<h5>{{ layout.label }}</h5>
							<p>{{ layout.instructions }}</p>
						</div>
					</div>
					<!-- TODO: If flexible -->
					<div class="field-controls" v-if="layout.type === 'flexible'">
						<i class="feather icon-trash-2"></i>
						<i class="feather icon-arrow-up"></i>
						<i class="feather icon-arrow-down"></i>
						<i class="fal fa-arrows"></i>
					</div>
				</div>
				<!--Content -->
				<div class="field-content">
					<!-- Text -->
					<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[layout.name]"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[layout.name]"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[layout.name]"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[layout.name]"></FieldEmail>
					<!-- Richtext -->
					<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRichText>
					<!-- Repeater -->
					<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRepeater>
					<!-- Flexible -->
					<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[layout.name]"></FieldFlexible>
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
import FieldRepeater from "@/components/editor/fields/Repeater";
import FieldFlexible from "@/components/editor/fields/FlexibleContent";

export default {
	name: "Fields",
	props: {
		layout: Array,
		fields: {
			required: true,
			type: Object
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
		FieldFlexible,
	},
	data: () => ({
		computedHeight: 'auto',
		isActive: true,
	}),
	mounted() {
		this.initHeight()
	},
	methods: {
		initHeight() {
			//this.computedHeight= getComputedStyle(this.$refs['myText']).height;
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