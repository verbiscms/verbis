/**
 * layout.js
 * Common util functions for fields.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const layoutMixin = {
	props: {

	},
	data: () => ({
	}),
	computed: {
		/*
		 * getButtonLabel()
		 * Retrieves the button label for the layout, if there
		 * is none set, 'Add Row' will be returned.
		 */
		getButtonLabel() {
			const def = "Add Row",
				layout = this.getLayout;

			// eslint-disable-next-line no-prototype-builtins
			if (!layout.hasOwnProperty("options")) {
				return def;
			}
			// eslint-disable-next-line no-prototype-builtins
			if (!layout.hasOwnProperty("button_label")) {
				return def;
			}
			const label = layout['options']['button_label'];
			if (!label) {
				return def;
			}

			return label;
		},
	}
};