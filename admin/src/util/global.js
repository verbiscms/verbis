/**
 * global.js
 * Global util functions & data.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const globalMixin = {
	data: () => ({
		timeoutDelay: 150,
	}),
	methods: {

	},
	computed: {
		/*
		 * getUserInfo()
		 */
		getUserInfo() {
			return this.$store.state.userInfo;
		},
		/*
		 * getSiteUrl()
		 */
		getSiteUrl() {
			return this.$store.state.site.url;
		},
		/*
		 * getSite()
		 */
		getSite() {
			return this.$store.state.site;
		},
		/*
		 * getTheme()
		 */
		getTheme() {
			return this.$store.state.theme;
		},
	}
};