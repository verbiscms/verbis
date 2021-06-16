module.exports = {
	//! Need to inject the admin path from the go config.
	devServer: {
		proxy: 'http://127.0.0.1:8080'
	},
	publicPath: '/admin',
	css: {
		loaderOptions: {
			sass: {
				prependData:
					`@import "~@/assets/scss/vendor/mesh/src/abstracts/mixins.scss";
                    @import "~@/assets/scss/abstracts/mixins.scss";
                    @import "~@/assets/scss/abstracts/variables.scss";`
			}
		}
	},
};
