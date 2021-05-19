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
			const label = this.getLayout['options']['button_label'];
			if (!label) {
				return "Add Row";
			}
			return label;
		},
	}
};