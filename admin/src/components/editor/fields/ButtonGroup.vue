<!-- =====================
	Field - Button Group
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<!-- Button Container -->
		<div class="button-cont">
			<button class="btn"
				v-for="(choice, choiceIndex) in getOptions['choices']"
				:key="choiceIndex"
				@click="value = choice"
				:class="{ 'btn-blue' : value === choice}">
				{{  choice }}</button>
		</div><!-- /Button Container -->
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
	name: "FiledButtonGroup",
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
	}),
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		value: {
			get() {
				return this.setDefaultValue(this.replacePrependAppend());
			},
			set(value) {
				this.$emit("update:fields", this.setPrependAppend(value));
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

.button {

	// Container
	// =========================================================================

	&-cont {
		padding: 6px 0;

		.btn {
			border-radius: 0;
			border-right: 2px solid $grey-light;
			transition: background-color 200ms ease, box-shadow 200ms ease;;
			will-change: background-color, box-shadow;

			&:first-child {
				border-top-left-radius: $btn-border-radius;
				border-bottom-left-radius: $btn-border-radius;
			}

			&:last-child {
				border-top-right-radius: $btn-border-radius;
				border-bottom-right-radius: $btn-border-radius;
				border-left: 0;
			}
		}
	}
}

</style>