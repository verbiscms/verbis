package validation

import (
	"fmt"
	pkgValidate "github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

type Validation struct {
	pkg *pkgValidate.Validate
}

type ValidationError struct {
	Key     string	`json:"key"`
	Type	string	`json:"type"`
	Message string	`json:"message"`
}

type Validator interface {
	Process(errors pkgValidate.ValidationErrors) []ValidationError
	CmdCheck(key string, data interface{}) error
	message(kind string, field string, param string) string
}

// Construct
func New() *Validation {
	v := &Validation{
		pkg: pkgValidate.New(),
	}

	v.pkg.SetTagName("binding")

	return v
}

// Handle validation errors and pass back, respond.
func (v* Validation) Process(errors pkgValidate.ValidationErrors) []ValidationError {

	var returnErrors []ValidationError
	for _, e := range errors {

		field := e.Field()
		result := strings.Split(e.Namespace(), ".")

		// TODO: Clean up here
		if len(result) > 2 && !strings.Contains(e.Namespace(), "PostCreate") {
			field = "";
			for i := 1; i < len(result); i++ {
				field += result[i]
			}
		}

		reg := regexp.MustCompile(`[A-Z][^A-Z]*`)
		fieldString := ""

		submatchall := reg.FindAllString(field, -1)
		for _, element := range submatchall {
			fieldString += strings.ToLower(element) + "_"
		}

		returnErrors = append(returnErrors, ValidationError{
			Key: strings.TrimRight(fieldString, "_"),
			Type: e.Tag(),
			Message: v.message(e.Tag(), field, e.Param()),
		})
	}

	return returnErrors
}

// Function for checking validation by struct on the command line.
func (v* Validation) CmdCheck(key string, data interface{}) error {

	err := v.pkg.Struct(data)

	if err != nil {
		validationErrors, _ := err.(pkgValidate.ValidationErrors)
		formatted := v.Process(validationErrors)

		for _, e := range formatted {
			if e.Key == key {
				return fmt.Errorf(e.Message)
			}
		}
	}

	return nil
}

// Custom Validator Messages
func (v* Validation) message(kind string, field string, param string) string {
	var errorMsg string

	switch kind {
	case "required":
		errorMsg = field + " is required."
		break
	case "email":
		errorMsg = "Enter a valid email address."
		break
	case "min":
		errorMsg = "Enter a minimum of " + param + " characters."
		break
	case "max":
		errorMsg = "Enter a maximum of " + param + " characters."
		break
	case "alpha":
		errorMsg = field + " must be alphanumeric."
		break
	}

	return errorMsg
}

