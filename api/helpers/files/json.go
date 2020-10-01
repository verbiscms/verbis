package files

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"io/ioutil"
	"os"
)

// ReadJson reads & return json file on a given path
// Return errors.INVALID if the file could not be opened.
func ReadJson(path string) ([]byte, error) {
	const op = "files.ReadJson"
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf( "Unable to open file with the path: %s", path), Operation: op, Err: err}
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}

