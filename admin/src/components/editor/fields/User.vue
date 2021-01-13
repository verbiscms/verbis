<!-- =====================
	Field - User // TODO  - Sort out roles!
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" v-if="!loadingUsers" ref="user">
		<!-- User Tags -->
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
			placeholder="Add User"
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
	name: "FieldPostObject",
	mixins: [fieldMixin],
	components: {
		VueTagsInput,
	},
	data: () => ({
		errors: [],
		selectedTags: [],
		users: [],
		focused: false,
		user: '',
		tag: '',
		tags: [],
		disabled: false,
		loadingUsers: true,
	}),
	mounted() {
		this.getUsers()
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
				this.helpers.setHeight(this.$refs.user.closest(".collapse-content"));
			});
		},
		/*
		 * handleBlur()
		 * Validate minimum and maximum length in options.
		 */
		handleBlur() {
			//this.validateRequired();
			if (this.getMinTags > this.selectedTags.length) {
				this.errors.push(`Enter a minimum of ${this.getMinTags} users.`)
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
		 * getUsers()
		 * Retrieve all users from the API.
		 */
		getUsers() {
			this.$store.dispatch("getUsers")
				.then(users => {
					this.users = users;
					this.mapUsers();
					this.loadingUsers = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		/*
		 * mapUsers()
		 * Set the users for filtering and set tags.
		 */
		mapUsers() {
			this.users = this.users.map(a => ({text: a['first_name'] + " " + a['last_name'], id: a.id, role: a.role.name}));
			this.setTags();
		},
		/*
		 * hasRole()
		 */
		hasRole(user) {
			const roles = this.getOptions['role']
			return !!(roles.length && roles.includes(user.role));
		},
		/*
		 * setTags()
		 * Set the existing tags on mounted if there is any by
		 * filtering through existing users.
		 */
		setTags() {
			if (this.field !== "") {
				this.selectedTags = this.field.reduce((r, x) => {
					const user = this.users.find(u => u.id === parseInt(x));
					if (user) {
						r.push({text: user.text, id: user.id, role: user.role.name})
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
		 * getRoles() TODO
		 * Get the roles allowed.
		 */
		getRoles() {
			const roles = this.getOptions['roles']
			return !roles || !roles.length ? false : roles;
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
			return this.users.filter(i => {
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
				return this.getValue === "" || !this.getValue ? "" : this.getValue.split(",");
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(value))
			}
		}
	}
}

</script>