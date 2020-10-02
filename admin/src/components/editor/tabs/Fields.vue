<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Text -->
		<div class="field-group"  v-for="group in fieldGroups" :key="group.uuid">
			<div class="field" v-for="field in group.fields" :key="field.uuid" >
				<div class="field-wrapper">
					<div class="field-title-cont">
						<span class="field-collapse" v-on:click="this.isActive = !this.isActive"></span>
						<div class="field-title">
							<h5>{{ field.label }}</h5>
							<p>{{  field.instructions }}</p>
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
				<div class="field-content" ref="myText" :style="[isActive ? { height : computedHeight } : {}]">
					<!-- Text -->
					<div v-if="field.type === 'text'">
						<FieldText></FieldText>
					</div>
					<!-- Textarea -->
					<div v-else-if="field.type === 'textarea'">
						<FieldTextarea></FieldTextarea>
					</div>
					<!-- Richtext -->
					<div v-else-if="field.type === 'richtext'">
						<FieldRichText></FieldRichText>
					</div>
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
import FieldRichText from "@/components/editor/fields/RichText";

export default {
	name: "Fields",
	props: {
		fieldGroups: Object,
	},
	components: {
		FieldRichText,
		FieldTextarea,
		FieldText,
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
			this.computedHeight= getComputedStyle(this.$refs['myText']).height;
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