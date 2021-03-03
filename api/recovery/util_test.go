// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
)

//func (t *RecoverTestSuite) Test_GetError() {
//
//	tt := map[string]struct {
//		input interface{}
//		want  *errors.Error
//	}{
//		"Non Pointer": {
//			errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
//			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
//		},
//		"Pointer": {
//			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
//			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
//		},
//		"Standard Error": {
//			fmt.Errorf("error"),
//			&errors.Error{Code: errors.TEMPLATE, Message: "error", Operation: "", Err: fmt.Errorf("error")},
//		},
//		"Nil Input": {
//			nil,
//			&errors.Error{Code: errors.TEMPLATE, Message: "Internal Verbis error, please report", Operation: "Internal", Err: nil},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			got := errors.ToError(test.input)
//			t.Equal(test.want.Operation, got.Operation)
//			t.Equal(test.want.Message, got.Message)
//			t.Equal(test.want.Code, got.Code)
//			if test.want.Err != nil {
//				t.Equal(test.want.Err.Error(), got.Err.Error())
//			}
//		})
//	}
//}

func (t *RecoverTestSuite) Test_TplLineNumber() {
	tt := map[string]struct {
		input *errors.Error
		want  int
	}{
		"Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:4: function "wrong" not defined`)},
			4,
		},
		"Found Second": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:10: function "wrong" not defined`)},
			10,
		},
		"Not Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home10: function "wrong" not defined`)},
			-1,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := tplLineNumber(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *RecoverTestSuite) Test_TplFileContents() {
	tt := map[string]struct {
		input string
		want  string
	}{
		"Found": {
			t.apiPath + "/test/testdata/html/partial.cms",
			"<h1>This is a partial file.</h1>",
		},
		"Not Found": {
			"wrong path",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := tplFileContents(test.input)
			t.Equal(test.want, got)
		})
	}
}
