// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"mime/multipart"
)

type Reader struct {
	Form   *domain.Form
	Reader dynamicstruct.Reader
}

type FormValues map[string]interface{}

func (f FormValues) JSON() ([]byte, error) {
	const op = "FormValues.JSON"
	v, err := json.Marshal(f)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not process the form fields for storing"), Operation: op, Err: err}
	}
	return v, nil
}

type Sender struct {
	Attachments []Attachment
	Fields      FormValues
}

func NewReader(form *domain.Form) *Reader {
	return &Reader{
		Form:   form,
		Reader: dynamicstruct.NewReader(form.Body),
	}
}

func (r *Reader) Values() (FormValues, Attachments, error) {
	const op = "FormReader.Values"

	m := make(FormValues)
	var attachments Attachments

	var totalSize int64 = 0
	for _, v := range r.Form.Fields {
		field := r.Reader.GetField(v.Label.Name())

		switch v.Type {
		case "file":
			a, err := getAttachment(field.Interface())
			if err != nil {
				return nil, nil, err
			}

			// Add to the total size to ensure its below
			// the UploadLimit
			totalSize += a.Size

			// Append to the attachments
			attachments = append(attachments, a)

			// Set the key of the FormValues to the MD5
			// name of the file.
			m[v.Key] = a.MD5name
		default:

			// html.escapestring here

			m[v.Key] = field.Interface()
		}
	}

	if int(totalSize/1024) > UploadLimit {
		return nil, nil, &errors.Error{Code: errors.INVALID, Message: "File attachments have exceeded the upload limit", Operation: op, Err: fmt.Errorf("attachements exceed the upload limit defined")}
	}

	return m, attachments, nil
}

// Struct returns the dynamic struct used for validation.
func ToStruct(form domain.Form) interface{} {
	instance := dynamicstruct.NewStruct()

	for _, v := range form.Fields {
		tag := fmt.Sprintf("json:\"%s\" form:\"%s\"", v.Key, v.Key)
		if v.Required {
			tag = fmt.Sprintf("%s binding:\"required\"", tag)
		}
		instance.AddField(v.Label.Name(), getType(v.Type), tag)
	}

	return instance.Build().New()
}

// getType
//
//
func getType(typ string) interface{} {
	var i interface{} = nil

	switch typ {
	case "text":
		i = ""
	case "int":
		i = 0
	case "float":
		i = 0.0
	case "boolean":
		i = false
	case "file":
		i = multipart.FileHeader{}
	}

	return i
}
