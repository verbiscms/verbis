// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var json = "json"

func TestGCP(t *testing.T) {
	UtilTestProvider(&environment.Env{
		GCPJson:      "json",
		GCPProjectID: "secret",
	}, &gcp{json: &json}, t)
}

func TestGCPDial_Error(t *testing.T) {
	g := gcp{}
	_, err := g.Dial(&environment.Env{})
	assert.NotNil(t, err)
}

func TestGCP_JSON(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	testFile := filepath.Join(wd, "testdata", "gcp.json")

	tt := map[string]struct {
		input gcp
		env   *environment.Env
		want  interface{}
	}{
		"Relative": {
			gcp{json: nil},
			&environment.Env{
				GCPJson: testFile,
			},
			`{"test": "value"}`,
		},
		"Defined": {
			gcp{json: &json},
			nil,
			json,
		},
		"Bad Path": {
			gcp{json: nil},
			&environment.Env{},
			"error reading google json path",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := test.input.getGCPJson(test.env)
			if err != nil {
				assert.Contains(t, err.Error(), got)
				return
			}
			assert.Contains(t, got, test.want)
		})
	}
}
