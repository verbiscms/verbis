<!-- =====================
	Fields
	===================== -->
<template>
	<section>

		<!-- Text -->
		<div class="field-group" v-for="(group, groupIndex) in layout" :key="group.uuid">
			<h4>{{ group.title }}</h4>
			<div v-for="(layout) in group.fields" :key="layout.uuid" :style="{ width: layout.wrapper['width'] + '%' }">
				<transition name="trans-fade">
					<div class="field" v-if="parseLogic(layout, groupIndex)">
						<div class="field-wrapper">
							<div class="field-title-cont">
								<span class="field-collapse" v-on:click="isActive = !isActive"></span>
								<div class="field-title">
									<h5>{{ layout.label }}</h5>
									<p>{{ layout.instructions }}</p>
								</div>
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
							<!-- Post Object -->
							<FieldPostObject v-if="layout.type === 'post_object'" :layout="layout" :fields.sync="fields[layout.name]" @update:field-error="fieldError = $event"></FieldPostObject>
							<!-- Repeater -->
							<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRepeater>
							<!-- Flexible -->
							<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[layout.name]"></FieldFlexible>
						</div>
					</div>
				</transition>
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
import FieldPostObject from "@/components/editor/fields/PostObject";
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
		fieldError: {
			type: String,
			default: "",
		},
	},
	components: {
		FieldText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldRichText,
		FieldPostObject,
		FieldRepeater,
		FieldFlexible,
	},
	data: () => ({
		computedHeight: 'auto',
		isActive: true
	}),
	mounted() {
		this.initHeight()
	},
	methods: {
		initHeight() {
			//this.computedHeight= getComputedStyle(this.$refs['myText']).height;
		},
		getLayoutByName(groupIndex, name) {
			return this.getLayout[groupIndex]['fields'].find(f => f.name === name)
		},
		parseLogic(layout, groupIndex) {
			const logic = layout['conditional_logic']
			let passed = true

			if (logic) {
				logic.forEach(block => {
					block.forEach(location => {
						const field = this.getLayoutByName(groupIndex, location.field),
							operator = location.operator,
							fieldEval = location.value;

						let value = this.fields[location.field],
							prepend = field.options['prepend'],
							append = field.options['append']

						value = value === undefined ? "" : value
						value = value.replace(prepend, "").replace(append, "")

						switch (operator) {
							case '>':
								passed = fieldEval > value;
								break;
							case '<':
								passed = fieldEval < value;
								break;
							case '>=':
								passed = fieldEval >= value;
								break;
							case '<=':
								passed = fieldEval <= value;
								break;
							case '==':
								passed = fieldEval == value;
								break;
							case '!=':
								passed = fieldEval != value;
								break;
							case '===':
								passed = fieldEval === value;
								break;
							case '!==':
								passed = fieldEval !== value;
								break;
						}
					})
				});
			}

			return passed;
		},
	},
	computed: {
		getLayout() {
			return this.layout
		},
		getFields() {
			return this.fields
		}
	},
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

</style>