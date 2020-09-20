package helpers

import (
	"io/ioutil"
	"os"
)

// Read & return json file, return error.
func ReadJson(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}
