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
					console.log(err);
					this.$noty.error("Error occurred, please refresh the page.");
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
					console.log(err);
					if (err.response.status === 400) {
						this.validate(err.response.data.data.errors);
						this.$noty.error(this.successMsg)
						return
					}
					this.$noty.error("Error occurred, please refresh the page.");
				})
				.finally(() => {
					setTimeout(() => {
						this.saving = false;
					}, 100);
				});
		},
	}
}