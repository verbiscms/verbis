<!-- =====================
	Field - Checkbox
	===================== -->
<template>
	<div class="field-cont">
		<div class="form-checkbox checkbox-cont">
			<input type="checkbox"
				:id="layout.uuid"
				v-model="field"
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

import { fieldMixin } from "@/util/fields/fields"

export default {
	name: "FieldText",
	mixins: [fieldMixin],
	computed: {
		/*
		 * getMessage()
		 * Get's the check box message from the options
		 * if there is one.
		 */
		getMessage() {
			const msg = this.getOptions['message'];
			return msg === undefined || msg === "" ? false : msg;
		},
		/*
		 * field()
		 * Returns the default value if there is none.
		 * Fire's value back up to the parent.
		 */
		field: {
			get() {
				return this.getValue === "" ? this.getOptions['default_value'] : (this.getValue === "true");
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(value.toString()));
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