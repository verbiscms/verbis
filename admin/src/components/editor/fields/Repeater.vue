<!-- =====================
	Field - Repeater
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="repeater">
		<draggable @start="drag=true" :list="repeaterFields['children']" :group="repeaterFields['children']" :sort="true" handle=".repeater-handle">
			<div class="repeater" v-for="(repeater, repeaterIndex) in repeaterFields['children']" :key="repeaterIndex">
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
						<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldText>
						<!-- Textarea -->
						<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldTextarea>
						<!-- Number -->
						<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldNumber>
						<!-- Range -->
						<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldRange>
						<!-- Email -->
						<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldEmail>
						<!-- Url -->
						<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldUrl>
						<!-- Password -->
						<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldPassword>
						<!-- =====================
							Content
							===================== -->
						<!-- Richtext -->
						<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldRichText>
						<!-- Image -->
						<FieldImage v-else-if="layout.type === 'image'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldImage>
						<!-- =====================
							Choice
							===================== -->
						<!-- Select -->
						<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldSelect>
						<!-- Multi Select -->
						<FieldTags v-else-if="layout.type === 'multi_select'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldTags>
						<!-- Checkbox -->
						<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldCheckbox>
						<!-- Radio -->
						<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldRadio>
						<!-- Button Group -->
						<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldButtonGroup>
						<!-- =====================
							Relational
							===================== -->
						<!-- Post Object -->
						<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldPost>
						<!-- User -->
						<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldUser>
						<!-- =====================
							Layout
							===================== -->
						<!-- Repeater -->
						<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldRepeater>
						<!-- Flexible -->
						<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields['children'][repeaterIndex][layout.name]" :field-key="getKey(repeaterIndex, layout.name)" :error-trigger="errorTrigger"></FieldFlexible>
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
import FieldImage from "@/components/editor/fields/Image";

// Choice
import FieldSelect from "@/components/editor/fields/Select";
import FieldTags from "@/components/editor/fields/Tags";
import FieldCheckbox from "@/components/editor/fields/Checkbox";
import FieldRadio from "@/components/editor/fields/Radio";
import FieldButtonGroup from "@/components/editor/fields/ButtonGroup";

// Relational
import FieldPost from "@/components/editor/fields/Post";
import FieldUser from "@/components/editor/fields/User";

import draggable from 'vuedraggable'

export default {
	name: "FieldRepeater",
	props: {
		layout: Object,
		fields: {
			type: Object,
			default: () => {
				return {};
			}
		},
		fieldKey: {
			type: String,
			default: "",
		},
		errorTrigger: {
			type: Boolean,
			default: false,
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
		FieldImage,
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
		FieldRepeater: () => import('@/components/editor/fields/Repeater'),
		FieldFlexible: () => import('@/components/editor/fields/FlexibleContent'),
		draggable,
	},
	data: () => ({
		errors: [],
	}),
	mounted() {
		this.init();
	},
	watch: {
		/*
		 * fields()
		 * Watch the fields and if they update, set the 'repeater'
		 * object to the length of the children, also update
		 * the child index.
		 */
		fields: {
			deep: true,
			handler(val) {
				this.$nextTick(() => {
					this.fields['repeater'].value = val.children.length.toString();
					this.$nextTick(() => {
						this.updateChildIndex();
					}, 20)
				});
			},
		},
	},
	methods: {
		/*
		 * init()
		 * Set the repeater fields to the original fields.
		 * Set the a new repeater parent and children if
		 * there are no fields stored.
		 */
		init() {
			this.repeaterFields = this.getFields;

			if (!this.repeaterFields['repeater']) {
				this.$set(this.repeaterFields, 'repeater', {
					uuid: this.getLayout.uuid,
					value: "0",
					name: this.getLayout.name,
					type: this.getLayout.type,
					key: this.fieldKey
				});
			}

			if (!this.repeaterFields['children']) {
				this.$set(this.repeaterFields, 'children', [])
			}
		},
		/*
		 * getKey()
		 * Get the key of repeater item to send to the child
		 * component. For nested repeaters, the layout's
		 * name is added.
		 */
		getKey(index, name) {
			if (this.fieldKey === "") {
				return this.getLayout.name + "|" + index + "|" + name;
			}
			return this.fieldKey + "|" + index + "|" + name
		},
		/*
		 * addRow()
		 * Add a repeater row to the children and update
		 * the height.
		 */
		addRow() {
			this.repeaterFields['children'].push({})
			this.updateHeight();
		},
		/*
		 * updateHeight()
		 * Update height when row is added
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.repeater.closest(".collapse-content"));
			});
		},
		/*
		 * deleteRow()
		 */
		deleteRow(index) {
			this.fields['children'].splice(index, 1);
		},
		/*
		 * moveUp()
		 * Move repeater item up.
		 */
		moveUp(index) {
			this.moveItem(index, index - 1)
		},
		/*
		 * moveDown()
		 * Move repeater item down.
		 */
		moveDown(index) {
			this.moveItem(index, index + 1)
		},
		/*
		 * moveItem()
		 * Move from, to, moves repeater children
		 * up or down.
		 */
		moveItem(from, to) {
			this.repeaterFields['children'].splice(to, 0, this.repeaterFields['children'].splice(from, 1)[0]);
		},
		/*
		 * updateChildIndex()
		 * When an item is moved around, the index of the child
		 * should change within the key.
		 */
		updateChildIndex() {
			this.repeaterFields['children'].forEach((child, index) => {
				for (const key in child) {
					// eslint-disable-next-line no-prototype-builtins
					if (child.hasOwnProperty(key)) {
						if ("key" in child[key] && 'uuid' in child[key]) {
							child[key]['key'] = this.getKey(index, child[key].name)
						}
						if ("repeater" in child[key]) {
							child[key]['repeater']['key'] = this.getKey(index, child[key]['repeater'].name)
						}
						if ("flexible" in child[key]) {
							child[key]['flexible']['key'] = this.getKey(index, child[key]['flexible'].name)
						}
					}
				}
			});
		}
	},
	computed: {
		/*
		 * getOptions()
		 * Get the field options.
		 */
		getOptions() {
			return this.layout.options
		},
		/*
		 * getLayout()
		 * Get the field layout.
		 */
		getLayout() {
			return this.layout;
		},
		/*
		 * getSubFields()
		 * Get the sub field layouts for looping over.
		 */
		getSubFields() {
			return this.layout['sub_fields'];
		},
		/*
		 * repeaterFields()
		 * Fire's the repeater fields back up to the parent.
		 */
		repeaterFields: {
			get() {
				return this.fields;
			},
			set() {
				this.$emit("update:fields", this.repeaterFields)
			}
		}
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