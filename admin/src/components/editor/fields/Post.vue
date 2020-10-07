<!-- =====================
	Field - Post Object
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }">
		<!-- Post Tags -->
		<vue-tags-input
			v-model="tag"
			:tags="selectedTags"
			:autocomplete-items="filteredItems"
			@tags-changed="updateTags"
			add-only-from-autocomplete
			:disabled="disabled"
			:max-tags="getMaxTags"
			@max-tags-reached="validate(`Only one post can be inserted in to the ${layout.label}`)"
			placeholder="Add Post"
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

export default {
	name: "FieldPost",
	props: {
		layout: Object,
		fields: Array,
	},
	components: {
		VueTagsInput,
	},
	data: () => ({
		errors: [],
		selectedTags: [],
		posts: [],
		focused: false,
		post: '',
		tag: '',
		tags: [],
		disabled: false,
	}),
	mounted() {
		this.getPosts();
	},
	methods: {
		validate(msg) {
			this.errors = [];
			this.errors.push(msg)
		},
		validateRequired() {
			if (!this.selectedTags.length && !this.getOptions["allow_null"]) {
				this.errors.push(`The ${this.layout.label} field is required.`)
			}
		},
		getPosts() {
			this.axios.get("/posts")
				.then(res => {
					this.posts = res.data.data.map(a => {
						return {
							text: a.post.title,
							id: a.post.id
						};
					});
					this.setTags()
				})
				.catch(err => {
					// TODO: Add toast
					console.log(err)
				});
		},
		setTags() {
			this.value.forEach(val => {
				this.posts.forEach(post => {
					const id = val.replace("verbis_post_", "");
					if (parseInt(id) === post.id) {
						this.selectedTags.push({
							text: post.text,
							id: post.id
						})
					}
				});
			});
		},
		updateTags(tags) {
			this.errors = [];
			this.selectedTags = tags;
			this.validateRequired()
			let tagsArr = []
			tags.forEach(tag => {
				tagsArr.push("verbis_post_" + tag.id)
			})
			this.value = tagsArr
		},
	},
	computed: {
		getOptions() {
			return this.layout.options
		},
		getMaxTags() {
			return this.layout.options['multiple'] ? 999999999999999999 : 1;
		},
		filteredItems() {
			return this.posts.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
		value: {
			get() {
				return this.fields === undefined ? [] : this.fields
			},
			set(value) {
				console.log(!this.errors.length)
				//this.$emit("update:field-error", !this.errors.length ? this.layout.label : "")
				this.$emit("update:fields", value)
			}
		}
	}
}

</script>