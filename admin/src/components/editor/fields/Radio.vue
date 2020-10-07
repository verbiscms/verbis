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
					v-model="value"
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

import { fieldMixin } from "@/util/fields"

export default {
	name: "FieldRadio",
	mixins: [fieldMixin],
	props: {
		layout: Object,
		fields: {
			type: String,
			default: ''
		},
	},
	data: () => ({
		errors: [],
		test: "",
	}),
	mounted() {
		this.setDefaultValueChoices()
	},
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		value: {
			get() {
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

.radio {

	// Container
	// =========================================================================

	&-cont {
		padding: 6px 0;
	}
}

</style>