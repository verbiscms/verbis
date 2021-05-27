<!-- =====================
	Field - Flexible Content
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="flexible">
		<draggable @start="drag=true" :list="fields['children']" :group="fields['children']" :sort="true" handle=".flexible-handle">
			<div class="flexible" v-for="(group, groupIndex) in getFields['children']" :key="groupIndex">
				<div class="card-header">
					<h4>{{ getParentLabel(groupIndex) }}</h4>
					<div class="card-controls">
						<i class="feather feather-trash-2" @click="deleteRow(groupIndex)"></i>
						<i class="feather feather-arrow-up" @click="moveUp(groupIndex)"></i>
						<i class="feather feather-arrow-down" @click="moveDown(groupIndex)"></i>
						<i class="flexible-handle fal fa-arrows"></i>
					</div>
				</div><!-- /Card Header -->
				<div class="flexible-body">
					<div class="card-body card-body-border-bottom" v-for="(layout, layoutKey) in getSubFields(groupIndex)" :key="layoutKey" :style="{ width: layout.wrapper['width'] + '%' }">
						<!-- Field Title -->
						<div class="field-title">
							<h4>{{ layout.label }}</h4>
							<p>{{ layout.instructions }}</p>
						</div>
						<!-- =====================
							Basic
							===================== -->
						<!-- Text -->
						<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldText>
						<!-- Textarea -->
						<FieldTextarea v-else-if="layout.type === 'textarea'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldTextarea>
						<!-- Number -->
						<FieldNumber v-if="layout.type === 'number'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldNumber>
						<!-- Range -->
						<FieldRange v-if="layout.type === 'range'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldRange>
						<!-- Email -->
						<FieldEmail v-if="layout.type === 'email'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldEmail>
						<!-- Url -->
						<FieldUrl v-if="layout.type === 'url'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldUrl>
						<!-- Password -->
						<FieldPassword v-if="layout.type === 'password'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldPassword>
						<!-- =====================
							Content
							===================== -->
						<!-- Richtext -->
						<FieldRichText v-else-if="layout.type === 'richtext'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger" :updating="updatingIndex"></FieldRichText>
						<!-- Image -->
						<FieldImage v-else-if="layout.type === 'image'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldImage>
						<!-- =====================
							Choice
							===================== -->
						<!-- Select -->
						<FieldSelect v-else-if="layout.type === 'select'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldSelect>
						<!-- Multi Select -->
						<FieldTags v-else-if="layout.type === 'multi_select'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldTags>
						<!-- Checkbox -->
						<FieldCheckbox v-else-if="layout.type === 'checkbox'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldCheckbox>
						<!-- Radio -->
						<FieldRadio v-else-if="layout.type === 'radio'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldRadio>
						<!-- Button Group -->
						<FieldButtonGroup v-else-if="layout.type === 'button_group'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldButtonGroup>
						<!-- =====================
							Relational
							===================== -->
						<!-- Post Object -->
						<FieldPost v-if="layout.type === 'post'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldPost>
						<!-- User -->
						<FieldUser v-if="layout.type === 'user'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldUser>
						<!-- =====================
							Layout
							===================== -->
						<!-- Repeater -->
						<FieldRepeater v-if="layout.type === 'repeater'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldRepeater>
						<!-- Flexible -->
						<FieldFlexible v-if="layout.type === 'flexible'" :layout="layout" :fields.sync="fields['children'][groupIndex][layout.name]" :field-key="getKey(groupIndex, layout.name)" :error-trigger="errorTrigger"></FieldFlexible>
					</div><!-- /Card Body -->
				</div><!-- /Flexible Body -->
			</div><!-- /Card -->
		</draggable>
		<div class="field-btn">
			<Popover :triangle="true" :classes="'popover-hover'">
				<template slot="items">
					<div class="popover-item" v-for="(group, index) in getLayouts" :key="index" @click="addRow(group.name)">{{ group.label }}</div>
				</template>
				<template slot="button">
					<button class="btn btn-blue"><i class="fal fa-plus-circle"></i>{{ getButtonLabel }}</button>
				</template>
			</Popover>
		</div>
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import {fieldMixin} from "@/util/fields/fields"
import {layoutMixin} from "@/util/fields/layout"

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

import Popover from "@/components/misc/Popover";
import draggable from 'vuedraggable'

export default {
	name: "FieldFlexible",
	mixins: [fieldMixin, layoutMixin],
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
		Popover,
		draggable,
	},
	data: () => ({
		errors: [],
		layouts: [],
		showPopover: false,
		layoutStr: [],
		updatingIndex: false,
	}),
	mounted() {
		this.init();
	},
	watch: {
		/*
		 * fields()
		 * Watch the fields and if they update, set the 'flexible'
		 * object to a comma separated array of layouts, also
		 * update the child index.
		 */
		fields: {
			deep: true,
			handler(val) {
				if (!('flexible' in val)) {
					return;
				}
				val.flexible.value =  this.layoutStr.join(",");
				this.$nextTick(() => {
					this.updateChildIndex();
				}, 20);
			},
		},
	},
	methods: {
		/*
		 * init()
		 * Set the flexible fields to the original fields.
		 * Set the a new flexible parent and children if
		 * there are no fields stored.
		 */
		init() {
			this.layoutFields = this.getFields;

			if (!this.layoutFields['flexible']) {
				const value = this.layoutStr === "" ? "" : this.layoutStr.join(",")
				this.$set(this.layoutFields, 'flexible', {
					uuid: this.getLayout.uuid,
					value: value,
					name: this.getLayout.name,
					type: this.getLayout.type,
					key: this.fieldKey
				});
			}

			if (!this.layoutFields['children']) {
				this.$set(this.layoutFields, 'children', [])
			}

			this.layoutStr = this.fields['flexible'].value === "" ? [] : this.fields['flexible'].value.split(",")
		},
		/*
		 * getKey()
		 * Get the key of repeater item to send to the child
		 * component. For nested flexible content, the layout's
		 * name is added.
		 */
		getKey(index, name) {
			if (this.fieldKey === "") {
				return this.getLayout.name + "|" + index + "|" + name;
			}
			return this.fieldKey + "|" + index + "|" + name
		},
		/*
		 * updateHeight()
		 * Update height when row is added
		 */
		updateHeight() {
			this.$nextTick(() => {
				setTimeout(() => {
					this.helpers.setHeight(this.$refs.flexible.closest(".collapse-content"));
				}, 20)
			});
		},
		/*
		 * deleteRow()
		 */
		deleteRow(index) {
			this.layoutFields['children'].splice(index, 1);
			this.layoutStr.splice(index, 1);
			this.updatingIndex = !this.updatingIndex;
		},
		/*
		 * addRow()
		 * Add a flexible field to the children and update
		 * the height.
		 */
		addRow(name) {
			let max = -1;
			// eslint-disable-next-line no-prototype-builtins
			if (this.getOptions && this.getOptions.hasOwnProperty("max")) {
				max = this.getOptions['max'];
			}

			if (max && this.layoutFields['children'].length === max) {
				this.$noty.error(`Only ${max} items allowed in this block.`)
				return;
			}

			const layouts = this.getLayouts;
			for (const key in layouts) {
				// eslint-disable-next-line no-prototype-builtins
				if (!layouts.hasOwnProperty(key)) {
					return;
				}
				const layout = layouts[key];
				if (name === layout.name) {
					this.layouts.push(layouts[key]);
					this.layoutFields['children'].push({});
					this.layoutStr.push(layout.name);
					this.updateHeight();
				}
			}
		},
		/*
		 * moveUp()
		 * Move flexible item up.
		 */
		moveUp(index) {
			this.moveItem(index, index - 1)
		},
		/*
		 * moveDown()
		 * Move flexible item down.
		 */
		moveDown(index) {
			this.moveItem(index, index + 1)
		},
		/*
		 * moveItem()
		 * Move from, to, moves flexible children
		 * up or down.
		 */
		moveItem(from, to) {
			this.layoutFields['children'].splice(to, 0, this.layoutFields['children'].splice(from, 1)[0]);
			this.layoutStr.splice(to, 0, this.layoutStr.splice(from, 1)[0]);
			this.updatingIndex = !this.updatingIndex;
		},
		/*
		 * getSubFields()
		 */
		getSubFields(index) {
			let name = this.layoutStr[index],
				layouts = this.getLayouts,
				layout = false;

			for (const key in layouts) {
				let l = layouts[key];
				if (l.name === name) {
					layout = l;
				}
			}
			// OLD Amendment:
			//const layout = this.getLayouts[this.layoutStr[index]];
			if (layout) {
				if ('sub_fields' in layout) {
					return layout['sub_fields'];
				}
			}
			return [];
		},
		/*
		 * getParentLabel()
		 */
		getParentLabel(index) {
			const layout = this.getLayouts[this.layoutStr[index]]
			if (layout) {
				return layout.label;
			}
			return "";
		},
		/*
		 * updateChildIndex()
		 * When an item is moved around, the index of the child
		 * should change within the key.
		 */
		updateChildIndex() {
			this.layoutFields['children'].forEach((child, index) => {
				const fields = child
				for (const key in fields) {
					// eslint-disable-next-line no-prototype-builtins
					if (fields.hasOwnProperty(key)) {
						if ("key" in child[key] && 'uuid' in child[key]) {
							fields[key].key = this.getKey(index, fields[key].name)
						}
						if ("repeater" in fields[key]) {
							fields[key]['repeater'].key = this.getKey(index, fields[key]['repeater'].name)
							return
						}
						if ("flexible" in fields[key]) {
							fields[key]['flexible'].key = this.getKey(index, fields[key]['flexible'].name)
							return
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
		 * getLayout()
		 * Get the children layouts.
		 */
		getLayouts() {
			return this.layout['layouts'];
		},
		/*
		 * getFields()
		 */
		getFields() {
			return this.fields
		},
		/*
		 * layoutFields()
		 * Fire's the repeater fields back up to the parent.
		 */
		layoutFields: {
			get() {
				return this.fields === undefined ? [] : this.fields
			},
			set() {
				this.$emit("update:fields", this.layoutFields)
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


	.flexible {
		border: 1px solid $grey-light;
		margin-bottom: 1.6rem;
		border-radius: 6px;

		// Header
		// =========================================================================

		&-header {
			display: flex;
			justify-content: space-between;
			width: 100%;
			margin-bottom: 1rem;
			padding: 0 20px 15px 20px;
			border-bottom: 1px solid $grey-light;

			h4 {
				margin-bottom: 0;
			}
		}

		// Body
		// =========================================================================

		&-body {
			display: flex;
			flex-wrap: wrap;
		}
	}

</style>