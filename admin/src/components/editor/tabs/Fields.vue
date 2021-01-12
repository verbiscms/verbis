<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Field Group -->
		<div v-for="(group, groupIndex) in layout" :key="group.uuid" class="field-group">
			<Collapse :show="true" :use-icon="true">
				<template v-slot:header>
					<div class="field-header">
						<h4 class="card-title">{{  group.title }}</h4>
						<div class="card-controls">
							<i class="feather feather-arrow-up" @click="moveGroupUp(groupIndex)"></i>
							<i class="feather feather-arrow-down" @click="moveGroupDown(groupIndex)"></i>
							<i class="feather feather-chevron-down" @click="collapseGroup(group.uuid)"></i>
						</div>
					</div>
				</template>
				<!-- Field Layout -->
				<template v-slot:body>
					<div class="field-body">
						<div class="field" v-for="(layout) in group.fields" :key="layout.uuid" :style="{ width: layout.wrapper['width'] + '%' }">
							<transition name="trans-fade">
								<div v-if="parseLogic(layout, groupIndex)">
									<!-- Field Title -->
									<div class="field-title" :class="{ 'field-title-margin-bottom' : layout.type === 'flexible' || layout.type === 'repeater' }">
										<h4>{{ layout.label }}</h4>
										<p>{{ layout.instructions }}</p>
									</div>
									<!-- Field Content -->
									<div class="field-content">
										<!-- =====================
											Basic
											===================== -->
										<!-- Text -->
										<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldText>
										<!-- Textarea -->
										<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldTextarea>
										<!-- Number -->
										<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldNumber>
										<!-- Range -->
										<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldRange>
										<!-- Email -->
										<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldEmail>
										<!-- Url -->
										<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldUrl>
										<!-- Password -->
										<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldPassword>
										<!-- =====================
											Content
											===================== -->
										<!-- Richtext -->
										<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldRichText>
										<!-- Image -->
										<FieldImage v-else-if="layout.type === 'image'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldImage>
										<!-- =====================
											Choice
											===================== -->
										<!-- Select -->
										<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldSelect>
										<!-- Tags -->
										<FieldTags v-else-if="layout.type === 'tags'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldTags>
										<!-- Checkbox -->
										<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldCheckbox>
										<!-- Radio -->
										<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldRadio>
										<!-- Button Group -->
										<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldButtonGroup>
										<!-- =====================
											Relational
											===================== -->
										<!-- Post Object -->
										<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldPost>
										<!-- User -->
										<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldUser>

										<!-- Category -->
										<!-- =====================
											Layout
											===================== -->
										<!-- Repeater -->
										<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldRepeater>
<!--										&lt;!&ndash; Flexible &ndash;&gt;-->
<!--										<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[layout.name]" :error-trigger="errorTrigger"></FieldFlexible>-->
									</div><!-- /Field Content -->
								</div>
							</transition>
						</div><!-- /Field Layout -->
					</div><!-- /Field Group Layout -->
				</template>
			</Collapse>
		</div><!-- /Field Group -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

// Basic
import FieldText from "@/components/editor/fields/Text";
import FieldTextarea from "@/components/editor/fields/Textarea";
import FieldNumber from "@/components/editor/fields/Number";
import FieldRange from "@/components/editor/fields/Range";
import FieldEmail from "@/components/editor/fields/Email";
import FieldUrl from "@/components/editor/fields/Url";
import FieldPassword from "@/components/editor/fields/Password";

// Content
import FieldRichText from "@/components/editor/fields/RichText";

// Choice
import FieldSelect from "@/components/editor/fields/Select";
import FieldTags from "@/components/editor/fields/Tags";
import FieldCheckbox from "@/components/editor/fields/Checkbox";
import FieldRadio from "@/components/editor/fields/Radio";
import FieldButtonGroup from "@/components/editor/fields/ButtonGroup";

// Relational
import FieldPost from "@/components/editor/fields/Post";
import FieldUser from "@/components/editor/fields/User";

// Layout
import FieldRepeater from "@/components/editor/fields/Repeater";
// import FieldFlexible from "@/components/editor/fields/FlexibleContent";
//
import Collapse from "@/components/misc/Collapse";
import FieldImage from "@/components/editor/fields/Image";

export default {
	name: "Fields",
	props: {
		layout: Array,
		fields: {
			required: true,
			type: Object,
		},
		errorTrigger: {
			type: Boolean,
			default: false,
		},
	},
	components: {
		FieldImage,
		Collapse,
		// Basic
		FieldText,
		FieldTextarea,
		FieldNumber,
		FieldRange,
		FieldEmail,
		FieldUrl,
		FieldPassword,
		// Content
		FieldRichText,
		// Choice
		FieldSelect,
		FieldTags,
		FieldCheckbox,
		FieldRadio,
		FieldButtonGroup,
		// Relational
		FieldPost,
		FieldUser,
		// Layout
		FieldRepeater,
		// FieldFlexible,
	},
	data: () => ({
		heights: {},
		computedHeights: {},
		isActive: true,
		errors: {},
	}),
	methods: {
		getLayoutByName(groupIndex, name) {
			return this.getLayout[groupIndex]['fields'].find(f => f.name === name)
		},
		moveGroupUp(index) {
			this.moveItem(index, index - 1)
		},
		moveGroupDown(index) {
			this.moveItem(index, index + 1)
		},
		moveItem(from, to) {
			this.layout.splice(to, 0, this.layout.splice(from, 1)[0]);
		},
		collapseGroup(uuid) {
			if (this.computedHeights[uuid] === "0px") {
				this.$set(this.computedHeights, uuid, this.heights[uuid])
			} else {
				this.$set(this.computedHeights, uuid, "0px")
			}
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