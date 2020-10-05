<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Text -->
		<div class="field-group" v-for="group in layout" :key="group.uuid">
			<div class="field" v-for="(field) in group.fields" :key="field.uuid" :style="{ width: field.wrapper['width'] + '%' }">
				<div class="field-wrapper">
					<div class="field-title-cont">
						<span class="field-collapse" v-on:click="isActive = !isActive"></span>
						<div class="field-title">
							<h5>{{ field.label }}</h5>
							<p>{{ field.instructions }}</p>
						</div>
					</div>
					<!-- TODO: If flexible -->
					<div class="field-controls" v-if="field.type === 'flexible'">
						<i class="feather icon-trash-2"></i>
						<i class="feather icon-arrow-up"></i>
						<i class="feather icon-arrow-down"></i>
						<i class="fal fa-arrows"></i>
					</div>
				</div>
				<!--Content -->
				<div class="field-content">
					<!-- Text -->
					<FieldText v-if="field.type === 'text'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldText>
					<!-- Textarea -->
					<FieldTextarea v-else-if="field.type === 'textarea'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldTextarea>
					<!-- Number -->
					<FieldNumber v-if="field.type === 'number'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldNumber>
					<!-- Range -->
					<FieldRange v-if="field.type === 'range'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldRange>
					<!-- Email -->
					<FieldEmail v-if="field.type === 'email'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldEmail>
					<!-- Richtext -->
					<FieldRichText v-else-if="field.type === 'richtext'" v-model="fields[field.name]" @input="emit"></FieldRichText>
					<!-- Repeater -->
					<FieldRepeater v-if="field.type === 'repeater'" :layout="field" v-model="fields[field.name]" @input="emit"></FieldRepeater>
				</div>
			</div>
			{{ fields }}
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

export default {
	name: "Fields",
	props: {
		layout: Array,
	},
	components: {
		FieldText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldRichText,
		FieldRepeater
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
		emit() {
			this.$emit("input", this.fields)
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