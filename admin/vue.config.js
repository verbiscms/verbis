module.exports = {
	//! Need to inject the admin path from the go config.
	publicPath: '/admin',
	css: {
		loaderOptions: {
			sass: {
				additionalData:
					`@import "@/assets/scss/vendor/mesh/src/abstracts/mixins.scss";
                    @import "@/assets/scss/abstracts/mixins.scss";
                    @import "@/assets/scss/abstracts/variables.scss";`
			}
		}
	}
};