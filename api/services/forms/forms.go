// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"fmt"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/services/storage"
	"golang.org/x/net/html"
	"mime/multipart"
)

type Reader struct {
	Form    *domain.Form
	Reader  dynamicstruct.Reader
	Storage storage.Provider
}

type Sender struct {
	Attachments []Attachment
	Fields      domain.FormValues
}

func NewReader(store storage.Provider, form *domain.Form) *Reader {
	return &Reader{
		Form:    form,
		Reader:  dynamicstruct.NewReader(form.Body),
		Storage: store,
	}
}

func (r *Reader) Values() (domain.FormValues, Attachments, error) {
	const op = "FormReader.Values"

	m := make(domain.FormValues)
	var attachments Attachments

	var totalSize int64 = 0
	for _, v := range r.Form.Fields {
		field := r.Reader.GetField(v.Label.Name())

		switch v.Type {
		case "file":

			a, err := getAttachment(field.Interface(), r.Storage)
			if err != nil {
				return nil, nil, err
			}

			if a == nil {
				continue
			}

			// Add to the total size to ensure its below
			// the UploadLimit
			totalSize += a.Size

			// Append to the attachments
			attachments = append(attachments, *a)

			// Set the key of the FormValues to the MD5
			// name of the file.
			m[v.Key] = a.MD5name
		case "checkbox":
			m[v.Key] = field.Bool()
		default:
			str := field.String()
			m[v.Key] = html.EscapeString(str)
		}
	}

	if float64(totalSize/1024)/1024 > UploadLimit { //nolint
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
	case "checkbox":
		i = false
	case "file":
		m := &multipart.FileHeader{}
		i = m
	}

	return i
}
