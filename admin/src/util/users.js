/**
 * users.js
 * Common util functions for users.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const userMixin = {
	methods: {
		/*
		 * createPassword()
		 * Generate a random password with the length of 16.
		 */
		createPassword() {
			let result = "";
			const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@:\\/@£$%=^&&*()_+?><',
				charactersLength = characters.length,
				lengthOfPassword = 24;
			for (let i = 0; i < lengthOfPassword; i++) {
				result += characters.charAt(Math.floor(Math.random() * charactersLength));
			}
			return result;
		}
	}
}



