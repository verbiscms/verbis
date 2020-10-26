<!-- =====================
	Form Group
	===================== -->
<template>
	<div class="form-group" :class="{ 'form-group-error' : getError }">
		<label class="form-label">{{ label }}</label>
		<slot></slot>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="getError">{{ getError }}</span>
		</transition><!-- /Message -->
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "FormGroup",
	props: {
		label: {
			type: String,
			default: "",
		},
		error: {
			type: String,
			default: ""
		}
	},
	data: () => ({
		show: true,
	}),
	computed: {
		/*
		 * getError()
		 * If the error is not set return false.
		 */
		getError() {
			return this.error === "" ? false : this.error;
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">



</style>