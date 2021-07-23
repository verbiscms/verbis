// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	pkgValidate "github.com/go-playground/validator/v10"
	strings2 "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/domain"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

// Validator defines methods for checking the validation errors
type Validator interface {
	Process(errors pkgValidate.ValidationErrors) []Error
	CmdCheck(key string, data interface{}) error
	message(kind string, field string, param string) string
}

// Validation defines site wide validation for endpoints
// and using the Package validation helper.
type Validation struct {
	Package *pkgValidate.Validate
}

type Errors []Error

// Error defines the structure when returning
// validation errors.
type Error struct {
	Key     string `json:"key"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// New - Construct & set tag name
func New() *Validation {
	const op = "Validation.New"
	v := &Validation{
		Package: pkgValidate.New(),
	}

	v.Package.SetTagName("binding")

	if v, ok := binding.Validator.Engine().(*pkgValidate.Validate); ok {
		err := v.RegisterValidation("password", comparePassword)
		if err != nil {
			fmt.Println(err)
			// Using logger has an import cycle.
			// logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error registering password validation", Operation: op, Err: err})
		}
	}

	return v
}

// Process handles validation errors and passes back to respond.
func (v *Validation) Process(errors pkgValidate.ValidationErrors) []Error {
	var returnErrors []Error
	for _, e := range errors {
		field := e.Field()
		result := strings.Split(e.Namespace(), ".")

		// TODO: Clean up here
		if len(result) > 2 && !strings.Contains(e.Namespace(), "PostCreate") || !strings.Contains(e.Namespace(), "UserCreate") {
			field = ""
			for i := 1; i < len(result); i++ {
				field += result[i]
			}
		}

		field = strings.ReplaceAll(field, "Part", "")
		field = strings.ReplaceAll(field, "User", "")

		reg := regexp.MustCompile(`[A-Z][^A-Z]*`)
		fieldString := ""

		if field == "" {
			field = e.StructField()
		}

		submatchall := reg.FindAllString(field, -1)
		for _, element := range submatchall {
			if element == "User" || element == "post" {
				continue
			}
			fieldString += strings.ToLower(element) + "_"
		}

		returnErrors = append(returnErrors, Error{
			Key:     strings.TrimRight(fieldString, "_"),
			Type:    e.Tag(),
			Message: v.message(e.Tag(), field, e.Param()),
		})
	}

	return returnErrors
}

// CmdCheck is a function for checking validation by struct on the command line.
func (v *Validation) CmdCheck(key string, data interface{}) error {
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
func (v *Validation) message(kind, field, param string) string {
	var errorMsg string

	field = strings2.AddSpace(field)
	param = strings2.AddSpace(param)

	switch kind {
	case "required":
		errorMsg = field + " is required."
	case "email":
		errorMsg = "Enter a valid email address."
	case "min":
		errorMsg = "Enter a minimum of " + param + " characters."
	case "max":
		errorMsg = "Enter a maximum of " + param + " characters."
	case "alpha":
		errorMsg = field + " must be alpha."
	case "alphanum":
		errorMsg = field + " must be alphanumeric."
	case "ip":
		errorMsg = field + " must be valid IP address."
	case "url":
		errorMsg = "Enter a valid url."
	case "eqfield":
		errorMsg = field + " must equal the " + param + "."
	case "password":
		errorMsg = field + " doesn't match our records."
	}

	return errorMsg
}

// comparePassword for the password field on the domain.UserPasswordReset
// (custom validation)
func comparePassword(fl pkgValidate.FieldLevel) bool {
	curPass := fl.Field().String()
	reset := fl.Parent().Interface().(*domain.UserPasswordReset)
	err := bcrypt.CompareHashAndPassword([]byte(reset.DBPassword), []byte(curPass))
	return err == nil
}
