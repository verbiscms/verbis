<!-- =====================
	Field - User // TODO  - Sort out roles!
	===================== -->
<template>
	<div class="field-cont" :class="{ 'field-cont-error' : errors.length }" v-if="!loadingUsers">
		<!-- User Tags -->
		<vue-tags-input
			v-model="tag"
			:tags="selectedTags"
			:autocomplete-items="filteredItems"
			@tags-changed="updateTags"
			add-only-from-autocomplete
			:disabled="disabled"
			:max-tags="getMaxTags"
			@max-tags-reached="validate(`Only one user can be inserted in to the ${layout.label}`)"
			placeholder="Add User"
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
	name: "FieldPostObject",
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
		validate(msg) {
			this.errors = [];
			this.errors.push(msg)
		},
		validateRequired() {
			if (!this.selectedTags.length && !this.getOptions["allow_null"]) {
				this.errors.push(`The ${this.layout.label.toLowerCase()} field is required.`)
			}
		},
		getUsers() {
			this.$store.dispatch("getUsers")
				.then(users => {
					this.users = users;
					this.mapUsers()
					this.loadingUsers = false;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
		},
		mapUsers() {
			this.users = this.users.map(a => {
				return {
					text: a['first_name'] + " " + a['last_name'],
					id: a.id,
					role: a.role.name,
				};
			});
			this.setTags();
		},
		hasRole(user) {
			const roles = this.getOptions['role']
			return !!(roles.length && roles.includes(user.role));
		},
		setTags() {
			this.value.forEach(val => {
				this.users.forEach(user => {
					if (val.id === user.id) {
						this.selectedTags.push({
							text: user.text,
							id: user.id,
							role: user.role.name
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
				tagsArr.push({
					id: tag.id,
					type: "user",
				})
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
			return this.users.filter(i => {
				return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
			});
		},
		value: {
			get() {
				return this.fields === undefined ? [] : this.fields
			},
			set(value) {
				this.$emit("update:fields", value)
			}
		}
	}
}

</script>