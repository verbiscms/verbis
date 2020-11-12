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
		},
		/*
		 * slugify()
		 * Slugify's given input.
		 */
		slugify(text) {
			text = text.replace(/^\s+|\s+$/g, ''); // trim
			text = text.toLowerCase();

			// Remove accents, swap ñ for n, etc
			let from = "ãàáäâẽèéëêìíïîõòóöôùúüûñç·_,:;";
			var to   = "aaaaaeeeeeiiiiooooouuuunc------";
			for (var i = 0, l = from.length ; i < l ; i++) {
				text = text.replace(new RegExp(from.charAt(i), 'g'), to.charAt(i));
			}

			text = text.replace(/[^a-z0-9/ -]/g, '') // remove invalid chars
				.replace(/\s+/g, '-') // collapse whitespace and replace by -
				.replace(/-+/g, '-'); // collapse dashes

			return text;
		},
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