

/**
 * validation.js
 * Common util functions for validation.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

const timeoutDelay = 1000;

export const validationMixin = {
	data: () => ({
		timeout: null,

	}),
	mounted() {
		this.getOptions();
	},
	methods: {
		validateEmail(timeout = true, email, key) {
			const errMsg =  "Enter a valid email address."

			if (timeout) {
				clearTimeout(this.timeout);
				this.timeout = setTimeout(() => {
					if (email !== "" && !this.helpers.validateEmail(email)) {
						this.$set(this.errors, key, errMsg);
					} else {
						this.$delete(this.errors, key);
					}
				}, timeoutDelay);
			} else {
				if (email !== "" && !this.helpers.validateEmail(email)) {
					this.$set(this.errors, key, errMsg);
				} else {
					this.$delete(this.errors, key);
				}
			}
		},
		validateUrl(timeout = true, url, key) {
			const errMsg =  "Enter a valid url."

			if (timeout) {
				clearTimeout(this.timeout);
				this.timeout = setTimeout(() => {
					if (url !== "" && !this.helpers.validateUrl(url)) {
						this.$set(this.errors, key, errMsg);
					} else {
						this.$delete(this.errors, key);
					}
				}, timeoutDelay);
			} else {
				console.log(url === "")
				if (url !== "" && !this.helpers.validateUrl(url)) {
					this.$set(this.errors, key, errMsg);
				} else {
					this.$delete(this.errors, key);
				}
			}
		},
		/*
		 * validate()
		 * Add errors if the post/put failed.
		 */
		validate(errors) {
			this.errors = {};
			errors.forEach(err => {
				this.$set(this.errors, err.key, err.message);
			})
		},
	}
}
