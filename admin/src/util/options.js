/**
 * options.js
 * Common util functions for options.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const optionsMixin = {
	data: () => ({
		doingAxios: true,
		saving: false,
		errors: [],
		data: {},
	}),
	mounted() {
		this.getOptions();
	},
	methods: {
		/*
		 * save()
		 * Save the updated options, check for field validation.
		 */
		getOptions() {
			this.axios.get("/options")
				.then(res => {
					this.data = res.data.data;
				})
				.catch(err => {
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					this.doingAxios = false;
					if (typeof this.runAfter == 'function') {
						this.runAfter();
					}
				});
		},
		/*
		 * save()
		 * Save the updated options, check for field validation.
		 */
		save() {
			this.saving = true;
			if (this.errors.length) {
				this.$noty.error(this.errorMsg)
				return
			}
			this.axios.post("/options", this.data)
				.then(() => {
					this.errors = [];
					this.$noty.success("Site options updated successfully.");
				})
				.catch(err => {
					this.helpers.checkServer(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error(this.successMsg)
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.saving = false;
					}, 100);
				});
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