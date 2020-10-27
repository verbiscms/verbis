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
		/*
		 * setAllHeight()
		 * Set all height of collapse content.
		 */
		setAllHeight() {
			this.$nextTick(() => {
				document.querySelectorAll(".auth-container .collapse-content").forEach(el => {
					if (el.style.maxHeight !== "0px") {
						const messages = el.querySelectorAll(".field-message-warning");
						if (messages.length) {
							const heightToAdd = messages.length * 26 + (messages.length + 7);
							el.style.maxHeight = (heightToAdd + parseInt(el.style.maxHeight.replace("px", ""))) + "px";
						}
					}
				});
			});
		}
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