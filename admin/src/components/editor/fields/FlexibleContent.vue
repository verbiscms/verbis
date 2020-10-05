<!-- =====================
	Field - Flexible Content
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">

		<div class="repeater-item" v-for="(group, groupIndex) in getLayouts" :key="groupIndex">

			<div v-if="getLayouts[groupIndex]">
				<div  class="test" v-for="(layout, layoutIndex) in group['sub_fields']" :key="layoutIndex">

						<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[groupIndex][layoutIndex]"></FieldText>

				</div>
			</div>
		</div>
		<div class="repeater-btn">
			<button class="btn btn-blue" @click="addRow">Add row</button>
		</div>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->
		<pre>{{ fields }}</pre>
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import FieldText from "@/components/editor/fields/Text";

export default {
	name: "FieldFlexible",
	props: {
		layout: Object,
		fields: Object,
	},
	components: {
		FieldText,
	},
	data: () => ({
		errors: [],
		layouts: [],
	}),
	mounted() {
		if (this.layoutFields !== undefined) {
			this.layoutFields = this.getFields
		}
	},
	methods: {
		deleteRow(index) {
			this.fields.splice(index, 1);
			this.layouts.splice(index, 1);
		},
		addRow() {
			//this.layoutFields.push({})
		},
		moveUp(index) {
			this.moveItem(index, index - 1)
		},
		moveDown(index) {
			this.moveItem(index, index + 1)
		},
		moveItem(from, to) {
			this.layoutFields.splice(to, 0, this.layoutFields.splice(from, 1)[0]);
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getLayouts() {
			// Need to do some work here!
			return this.layout['layouts'];
		},
		getFields() {
			return this.fields
		},
		layoutFields: {
			get() {
				if (this.fields === undefined) {
					let temp = {};
					for (const layout in this.getLayouts) {
						if (this.getLayouts[layout] !== undefined) {
							temp[layout] = {};
							const fields = this.getLayouts[layout]['sub_fields'];
							if (fields !== undefined) {
								for (const field in fields) {
									temp[layout][field] = ""
								}
							} else {
								console.log("in fields undef")
							}
						} else {
							console.log("in")
						}
					}
					console.log(temp)
					return temp
				} else {
					return this.fields
				}
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