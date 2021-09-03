// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	vstrings "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/domain"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
	"sync"
)

// Errors defines a slice of validation errors used
// within the application.
type Errors []Error

// Error defines the structure when returning
// validation errors.
type Error struct {
	Key     string `json:"key"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// defaultValidator defines site wide validation for endpoints
// and using the validate helper.
type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// v is the singleton of the defaultProvider.
var v = &defaultValidator{}

// init assigns the gin binding to the default
// validator. It also initialises the validation
// package.
func init() {
	binding.Validator = v
	v.lazyInit()
}

// Validator returns to package validation used for
// validation structs and maps etc..
func Validator() *validator.Validate {
	return v.validate
}

// Process handles validation errors and passes back to respond.
// If the kind of the error is not of type ValidationErrors
// nil will be returned.
func Process(err error) Errors {
	var returnErrors []Error

	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	for _, e := range errors {
		returnErrors = append(returnErrors, Error{
			Key:     e.Field(),
			Type:    e.Tag(),
			Message: message(e.Tag(), e.Field(), e.Param()),
		})
	}

	return returnErrors
}

// getKey traverses over the field and converts the string to
// title case.
func getKey(field string) string {
	split := strings.Split(field, "_")
	var str []string
	for _, item := range split {
		if item == "id" || item == "Id" || item == "url" || item == "Url" {
			str = append(str, strings.ToUpper(item))
			continue
		}
		str = append(str, strings.Title(item))
	}
	return strings.Join(str, " ")
}

// message checks the kind, field and parameters and binds custom
// error messages.
func message(kind, field, param string) string {
	var errorMsg string

	field = getKey(field)
	param = vstrings.AddSpace(param)

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
		errorMsg = "Enter a valid URL."
	case "eqfield":
		errorMsg = field + " must equal the " + param + "."
	case "password":
		errorMsg = field + " doesn't match our records."
	default:
		errorMsg = "Validation failed on the " + field + " field."
	}
	return errorMsg
}

// comparePassword for the password field on the domain.UserPasswordReset
// (custom validation)
func comparePassword(fl validator.FieldLevel) bool {
	curPass := fl.Field().String()
	reset := fl.Parent().Interface().(domain.UserPasswordReset)
	err := bcrypt.CompareHashAndPassword([]byte(reset.DBPassword), []byte(curPass))
	return err == nil
}

// Engine is the implementation of the StructValidator
// for gin.
func (v *defaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

// ValidateStruct is the implementation of the StructValidator
// for gin.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// kindOfData returns the reflect.Kind of the data passed,
// if the item is a pointer, it will be de referenced.
func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// lazyInit initialises the defaultValidator on initialisation.
// It creates a new validator, sets the tag key and
// registers a tag name function.
func (v *defaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.RegisterTagNameFunc(tagNameFunc)
		v.validate.SetTagName("binding")
		_ = v.validate.RegisterValidation("password", comparePassword)
	})
}

// tagNameFunc sets the validation key. If the validation_key
// exists it will be returned, otherwise the `json` key
// will be returned.
func tagNameFunc(fld reflect.StructField) string {
	key := fld.Tag.Get("validation_key")
	if key != "" {
		return key
	}
	return fld.Tag.Get("json")
}
