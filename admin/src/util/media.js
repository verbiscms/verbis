/**
 * media.js
 * Common util functions for media.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

export const mediaMixin = {
	methods: {
		/*
		 * formatBytes()
		 * Return formatted byte information for file size.
		 */
		formatBytes(bytes, decimals = 2) {
			if (bytes === 0) return '0 Bytes';

			const k = 1024,
				dm = decimals < 0 ? 0 : decimals,
				sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
				i = Math.floor(Math.log(bytes) / Math.log(k));

			return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
		},
		/*
		 * getMediaType()
		 * Determines if the file is a image, video or file.
		 */
		getMediaType(type) {
			const images = ['image/png', 'image/jpeg', 'image/gif', 'image/webp', 'image/bmp', 'image/svg+xml'],
				video = ['video/mpeg', 'video/mp4', 'video/webm'];
			if (images.includes(type)) {
				return "image";
			} else if (video.includes(type)) {
				return "video";
			} else {
				return "file";
			}
		}
	}
}