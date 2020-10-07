<!-- =====================
	Field - Checkbox
	===================== -->
<template>
	<div class="field-cont">
		<div class="form-checkbox checkbox-cont">
			<input type="checkbox"
				:id="layout.uuid"
				v-model="value"
				:true-value="true"
				:false-value="false">
			<label :for="layout.uuid">
				<i class="fal fa-check"></i>
			</label>
			<div v-if="getMessage" class="form-checkbox-text">{{ getMessage }}</div>
		</div>
	</div><!-- /Container -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import { fieldMixin } from "@/util/fields"

export default {
	name: "FieldText",
	mixins: [fieldMixin],
	props: {
		layout: Object,
		fields: {
			type: Boolean,
		},
	},
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		getMessage() {
			const msg = this.getOptions['message'];
			return msg === undefined || msg === "" ? false : msg;
		},
		value: {
			get() {
				if (this.fields === undefined) {
					return this.getOptions['default_value']
				}
				return this.fields;
			},
			set(value) {
				this.$emit("update:fields", value);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.checkbox {

	// Container
	// =========================================================================

	&-cont {
		padding: 6px 0;
	}
}

</style>