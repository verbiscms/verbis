<!-- =====================
	Field - Tags
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<!-- Multi Select Tags -->
		<vue-tags-input
			v-model="tag"
			:tags="selectedTags"
			:autocomplete-items="filteredItems"
			@tags-changed="updateTags"
			add-only-from-autocomplete
			:disabled="disabled"
			:max-tags="getMaxTags"
			@max-tags-reached="validate(`Only one tag can be inserted in to the ${layout.label}`)"
			:placeholder="getButtonLabel"
			@blur="validateRequired"
		/>
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

import VueTagsInput from '@jack_reddico/vue-tags-input';
import { fieldMixin } from "@/util/fields"

export default {
	name: "FieldTags",
	mixins: [fieldMixin],
	props: {
		layout: Object,
		fields: {
			type: Array,
			default: () => {
				return [];
			},
		},
	},
	components: {
		VueTagsInput,
	},
	data: () => ({
		errors: [],
		selectedTags: [],
		items: [],
		item: '',
		tag: '',
		tags: [],
		disabled: false,
	}),
	mounted() {
		this.setDefault();
		this.getItems();
	},
	methods: {
		validate() {
			this.errors = [];
			if (!this.getOptions["allow_null"]) {
				this.validateRequired()
			}
		},
		getItems() {
			const choices = this.getOptions['choices'];
			for (const choice in choices) {
				this.items.push({
					text: choices[choice],
					key: choice,
				});
			}
		},
		updateTags(tags) {
			this.errors = [];
			this.selectedTags = tags;
			this.validateRequired()
			let tagsArr = []
			tags.forEach(tag => {
				tagsArr.push({
					text: tag.text,
					key: tag.key,
				})
			})
			this.value = tagsArr
		},
		setDefault() {
			if (!this.fields.length && this.getOptions['default_value'] && this.getOptions['default_value'].length) {
				const opts = this.getOptions['default_value'];
				let defaultVal = [];
				opts.forEach(opt => {
					defaultVal.push({
						text: this.getOptions['choices'][opt],
						key: opt,
					})
				});
				if (defaultVal.length) {
					this.selectedTags = defaultVal;
					this.value = defaultVal
				}
			} else {
				this.value.forEach(tag => {
					this.selectedTags.push(tag)
				});
			}
		}
	},
	computed: {
		getOptions() {
			return this.layout.options;
		},
		getLayout() {
			return this.layout;
		},
		getMaxTags() {
			return !this.getOptions['maximum'] ? 999999999999999999 :  this.getOptions['maximum'];
		},
		getButtonLabel() {
			return this.getOptions['button_label'] ? this.getOptions['button_label'] : "Add item"
		},
		filteredItems() {
			return this.items.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
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