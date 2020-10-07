<!-- =====================
	Fields
	===================== -->
<template>
	<section>
		<!-- Field Group -->
		<div class="field-group" v-for="(group, groupIndex) in layout" :key="group.uuid">
			<h4>{{ group.title }}</h4>
			<!-- Field Layout -->
			<div v-for="(layout) in group.fields" :key="layout.uuid" :style="{ width: layout.wrapper['width'] + '%' }">
				<transition name="trans-fade">
					<div class="field" v-if="parseLogic(layout, groupIndex)">
						<!-- Field Wrapper -->
						<div class="field-wrapper">
							<div class="field-title-cont">
								<span class="field-collapse" v-on:click="isActive = !isActive"></span>
								<div class="field-title">
									<h5>{{ layout.label }}</h5>
									<p>{{ layout.instructions }}</p>
								</div>
							</div>
						</div><!-- /Field Wrapper -->
						<!-- /Field Content -->
						<div class="field-content">
							<!-- =====================
								Basic
								===================== -->
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
							<!-- Url -->
							<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields[layout.name]"></FieldUrl>
							<!-- Password -->
							<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields[layout.name]"></FieldPassword>
							<!-- =====================
								Content
								===================== -->
							<!-- Richtext -->
							<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRichText>
							<!-- =====================
								Choice
								===================== -->
							<!-- Select -->
							<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields[layout.name]"></FieldSelect>
							<!-- Checkbox -->
							<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields[layout.name]"></FieldCheckbox>
							<!-- Radio -->
							<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRadio>
							<!-- Button Group -->
							<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields[layout.name]"></FieldButtonGroup>
							<!-- =====================
								Relational
								===================== -->
							<!-- Post Object -->
							<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields[layout.name]"></FieldPost>
							<!-- User -->
							<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields[layout.name]"></FieldUser>
							<!-- =====================
								Layout
								===================== -->
							<!-- Repeater -->
							<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields[layout.name]"></FieldRepeater>
							<!-- Flexible -->
							<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields[layout.name]"></FieldFlexible>
						</div><!-- /Field Content -->
					</div>
				</transition>
			</div><!-- /Field Layout -->
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
import FieldCheckbox from "@/components/editor/fields/Checkbox";
import FieldRadio from "@/components/editor/fields/Radio";
import FieldButtonGroup from "@/components/editor/fields/ButtonGroup";

// Relational
import FieldPost from "@/components/editor/fields/Post";
import FieldUser from "@/components/editor/fields/User";

// Layout
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
		FieldCheckbox,
		FieldRadio,
		FieldButtonGroup,
		// Relational
		FieldPost,
		FieldUser,
		// Layout
		FieldRepeater,
		FieldFlexible,
	},
	data: () => ({
		computedHeight: 'auto',
		isActive: true,
		errors: {},
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