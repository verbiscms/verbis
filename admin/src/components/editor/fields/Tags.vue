<!-- =====================
	Field - Tags
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="tags">
		<!-- Multi Select Tags -->
		<vue-tags-input
			v-model="tag"
			:tags="selectedTags"
			:autocomplete-items="filteredItems"
			@tags-changed="updateTags"
			add-only-from-autocomplete
			:disabled="disabled"
			:max-tags="getMaxTags"
			:autocomplete-min-length="0"
			@max-tags-reached="maxTagsReached"
			:placeholder="getButtonLabel"
			@focus="updateHeight"
			:add-on-key="[13, ':', ';']"
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
import { fieldMixin } from "@/util/fields/fields"
import { choiceMixin } from "@/util/fields/choice"

export default {
	name: "FieldTags",
	mixins: [fieldMixin,choiceMixin],
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
		/*
		 * setDefault()
		 * Sets default values on mounted, if no value exists
		 * and the default value exists.
		 */
		setDefault() {
			if (!this.field && this.getOptions['default_value'] && this.getOptions['default_value'].length) {
				this.selectedTags = this.getOptions['default_value'].map(t => {
					const key = Object.keys(this.getChoices).find(key => this.getChoices[key] === t);
					return {key: key,text: t}
				});
			} else {
				this.setTags();
			}
		},
		/*
		 * updateHeight()
		 * Update the height of the container when searching.
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.tags.closest(".collapse-content"));
			});
		},
		/*
		 * handleBlur()
		 * Validate minimum and maximum length in options.
		 */
		handleBlur() {
			this.validateRequired();
			if (this.getMinTags > this.selectedTags.length) {
				this.errors.push(`Enter a minimum of ${this.getMinTags} ${this.getLayout.label.toLowerCase()}.`)
			}
		},
		/*
		 * maxTagsReached()
		 * Handler for maximum tags reached.
		 */
		maxTagsReached() {
			this.errors.push(`Only a maximum of ${this.getMaxTags} ${this.getLayout.label.toLowerCase()}.`);
			setTimeout(() => {
				this.updateHeight();
			}, 10)
		},
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			this.validateRequired();
			this.updateHeight();
		},
		/*
		 * getItems()
		 * Pushes to the items array on mounted for vue tags
		 * input.
		 */
		getItems() {
			for (const choice in this.getChoices) {
				this.items.push({
					text: this.getChoices[choice],
					key: choice,
				});
			}
		},
		/*
		 * setTags()
		 * Set the existing tags on mounted if there is any.
		 */
		setTags() {
			if (this.field) {
				this.selectedTags = this.field.map(t => {
					return {key: t,text: this.getOptions['choices'][t]}
				});
			}
		},
		/*
		 * updateTags()
		 * Update tags when fired.
		 */
		updateTags(tags) {
			this.errors = [];
			this.selectedTags = tags;
			this.updateHeight();
			this.handleBlur();
			this.field = tags.map(t => t.key).join(",")
		},
	},
	computed: {
		/*
		 * getMaxTags()
		 * Get minimum amount of tags required.
		 */
		getMinTags() {
			return !this.getOptions['min'] ? -1 : this.getOptions['min'];
		},
		/*
		 * getMaxTags()
		 * Get maximum amount of tags required.
		 */
		getMaxTags() {
			return !this.getOptions['max'] ? 999999999999999999 : this.getOptions['max'];
		},
		/*
		 * getButtonLabel()
		 * Obtain the button label for the tags component.
		 * Defaults to "Add Item"
		 */
		getButtonLabel() {
			return this.getOptions['button_label'] ? this.getOptions['button_label'] : "Add item"
		},
		/*
		 * filteredItems()
		 * Filter tags.
		 */
		filteredItems() {
			return this.items.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
		/*
		 * field()
		 * Splits comma separated list and for use and
		 * fire's value back up to the parent.
		 */
		field: {
			get() {
				return this.getValue === "" ? false : this.getValue.split(",");
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(value));
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