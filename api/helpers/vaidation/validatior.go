package validation

import (
	"fmt"
	pkgValidate "github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

// Validator defines methods for checking the validation errors
type Validator interface {
	Process(errors pkgValidate.ValidationErrors) []ValidationError
	CmdCheck(key string, data interface{}) error
	message(kind string, field string, param string) string
}

// Validation defines site wide validation for endpoints
// and using the Package validation helper.
type Validation struct {
	Package *pkgValidate.Validate
}

// ValidationError defines the structure when returning
// validation errors.
type ValidationError struct {
	Key     string	`json:"key"`
	Type	string	`json:"type"`
	Message string	`json:"message"`
}

// New - Construct & set tag name
func New() *Validation {
	v := &Validation{
		Package: pkgValidate.New(),
	}

	v.Package.SetTagName("binding")

	return v
}

// Process handles validation errors and passes back to respond.
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

// CmdCheck is a function for checking validation by struct on the command line.
func (v* Validation) CmdCheck(key string, data interface{}) error {

	err := v.Package.Struct(data)

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

// message checks the kind, field and parameters and binds custom
// error messages.
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
		errorMsg = field + " must be alpha."
		break
	case "alphanum":
		errorMsg = field + " must be alphanumeric."
		break
	case "ip":
		errorMsg = field + " must be valid IP address."
		break
	}

	return errorMsg
}

