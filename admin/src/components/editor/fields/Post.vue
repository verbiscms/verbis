<!-- =====================
	Field - Post Object
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" ref="post">
		<!-- Post Tags -->
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
			placeholder="Add Post"
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

export default {
	name: "FieldPost",
	mixins: [fieldMixin],
	components: {
		VueTagsInput,
	},
	data: () => ({
		errors: [],
		selectedTags: [],
		posts: [],
		post: '',
		tag: '',
		tags: [],
		disabled: false,
	}),
	mounted() {
		this.getPosts();
	},
	methods: {
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
		 * updateHeight()
		 * Update the height of the container when searching.
		 */
		updateHeight() {
			this.$nextTick(() => {
				this.helpers.setHeight(this.$refs.post.closest(".collapse-content"));
			});
		},
		/*
		 * handleBlur()
		 * Validate minimum and maximum length in options.
		 */
		handleBlur() {
			//this.validateRequired();
			if (this.getMinTags > this.selectedTags.length) {
				this.errors.push(`Enter a minimum of ${this.getMinTags} posts.`)
			}
			this.updateHeight();
		},
		/*
		 * maxTagsReached()
		 * Handler for maximum tags reached.
		 */
		maxTagsReached() {
			let label = this.getMaxTags === 0 ? 'item' : 'items'
			this.errors.push(`Enter a maximum of ${this.getMaxTags} ${label} can be inserted in to the ${this.layout.label}`);
			setTimeout(() => {
				this.updateHeight();
			}, 10)
		},
		/*
		 * getPosts()
		 * Retrieve all posts from the API.
		 */
		getPosts() {
			this.axios.get("/posts")
				.then(res => {
					this.mapPosts(res.data.data);
					this.setTags();
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				});
		},
		/*
		 * mapPosts()
		 * If the posts are undefined, return an empty array.
		 * If resources are set, they are filters and returned.
		 * Or all posts are returned for filtering.
		 */
		mapPosts(posts) {
			if (posts === undefined || Object.keys(posts).length === 0 && posts.constructor === Object) {
				this.posts = [];
				return
			}

			if (this.getResources) {
				this.posts = posts.filter(p => {
					if (this.getResources.includes(p.post.resource)) {
						return p
					}
				}).map(p => ({text: p.post.title, id: p.post.id}))
				return
			}

			this.posts = posts.map(a => ({text: a.post.title, id: a.post.id}));
		},
		/*
		 * setTags()
		 * Set the existing tags on mounted if there is any by
		 * filtering through existing posts.
		 */
		setTags() {
			if (this.field !== "") {
				this.selectedTags = this.field.reduce((r, x) => {
					const post = this.posts.find(p => p.id === parseInt(x));
					if (post) {
						r.push({text: post.text, id: post.id})
					}
					return r;
				}, [])
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
			this.field = tags.map(t => t.id).join(",")
		},
	},
	computed: {
		/*
		 * getResource()
		 * Get the resources allowed.
		 */
		getResources() {
			const resources = this.getOptions['resource']
			return !resources || !resources.length ? false : resources;
		},
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
		 * filteredItems()
		 * Filter tags.
		 */
		filteredItems() {
			return this.posts.filter(i => {
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
				return this.getValue === "" ? "" : this.getValue.split(",");
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(value))
			}
		}
	}
}

</script>