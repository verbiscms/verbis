package mime

import (
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
)

// Get the mime type by opening the file
func TypeByFile(file *multipart.FileHeader) (string, error) {
	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return "", err
	}

	return mime.String(), nil
}

// Check a whitelist of MIME types
func IsValidMime(allowed []string, mime string) bool {
	if mimetype.EqualsAny(mime, allowed...) {
		return true
	} else {
		return false
	}
}

// Check if the mime type is an image
func IsImage(mime string) bool {

	var imageMimeTypes = []string{
		"image/vnd.adobe.photoshop",
		"image/png",
		"image/jpeg",
		"image/jp2",
		"image/jpx",
		"image/jpm",
		"image/gif",
		"image/webp",
		"image/tiff",
		"image/bmp",
		"image/x-icon",
		"image/vnd.djvu",
		"image/bpg",
		"image/vnd.dwg",
		"image/x-icns",
		"image/heic",
		"image/heic-sequence",
		"image/heif",
		"image/heif-sequence",
		"image/svg+xml",
	}

	for _, v := range imageMimeTypes {
		if mime == v {
			return true
		}
	}

	return false
}