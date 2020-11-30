package mime

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
)

// Get the mime type by opening the file
// Returns errors.INTERNAL if the file could not be opened.
// Returns errors.NOTFOUND if the mime type was not found.
func TypeByFile(file *multipart.FileHeader) (string, error) {
	const op = "mime.TypeByFile"

	reader, err := file.Open()
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to open file with the name: %s", file.Filename), Operation: op, Err: err}
	}
	defer reader.Close()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Mime type not found: %s", mime), Operation: op, Err: err}
	}

	return mime.String(), nil
}

// IsValidMime checks a whitelist of MIME types
// Returns true if the file is in the whitelist.
func IsValidMime(allowed []string, mime string) bool {
	if mimetype.EqualsAny(mime, allowed...) {
		return true
	}
	return false
}

