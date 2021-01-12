<!-- =====================
	Field - Radio
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<!-- Radio Container -->
		<div class="form-radio-cont radio-cont">
			<div class="form-radio" v-for="(choice, choiceIndex) in getOptions['choices']" :key="choiceIndex">
				<input type="radio"
					:id="getLayout.uuid + '-' + choiceIndex"
					:name="getLayout.uuid"
					v-model="field"
					:value="choice">
				<label :for="getLayout.uuid + '-' + choiceIndex"></label>
				<div class="form-radio-text">{{ choice }}</div>
			</div>
		</div><!-- /Radio Container -->
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

import { fieldMixin } from "@/util/fields/fields"
import {choiceMixin} from "@/util/fields/choice"

export default {
	name: "FieldRadio",
	mixins: [fieldMixin, choiceMixin],
	data: () => ({
		errors: [],
	}),
	mounted() {
		this.setDefault();
	},
	created() {
		this.fields.key = this.getFormat;
	},
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			if (!this.getOptions["allow_null"]) {
				this.validateRequired();
			}
		},
	},
	computed: {
		/*
		 * field()
		 * Fire's value back up to the parent.
		 */
		field: {
			get() {
				return this.getMultipleFormat();
			},
			set(value) {
				this.setMultipleFormat(value);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.radio {

	// Container
	// =========================================================================

	&-cont {
		padding: 6px 0;
	}
}

</style>