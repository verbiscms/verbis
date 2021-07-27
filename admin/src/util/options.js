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
		errors: {},
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
					if (typeof this.runAfterGet == 'function') {
						this.runAfterGet();
					}
					setTimeout(() => {
						this.doingAxios = false;
					}, this.axiosTimeout ? 200 : 0)
				});
		},
		/*
		 * save()
		 * Save the updated options, check for field validation.
		 */
		save() {
			this.saving = true;

			if (typeof this.runBeforeSave == 'function') {
				this.runBeforeSave();
			}

			this.axios.post("/options", this.data)
				.then(() => {
					this.errors = [];
					this.$noty.success("Site options updated successfully.");
				})
				.catch(err => {
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error(this.errorMsg);
						this.setAllHeight();
						return;
					}
					this.helpers.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.saving = false;
					}, this.timeoutDelay);

					setTimeout(() => {
						this.axios.get("/site")
							.then(res => {
								this.$store.commit('setSite', res.data.data);
							});
					}, 500)
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
