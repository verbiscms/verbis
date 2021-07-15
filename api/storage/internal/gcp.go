// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/google"
	"io/ioutil"
)

type gcp struct {
	json *string
}

func (g *gcp) Dial(env *environment.Env) (stow.Location, error) {
	json, err := g.getGCPJson(env)
	if err != nil {
		return nil, err
	}

	return stow.Dial(google.Kind, stow.ConfigMap{
		google.ConfigJSON:      json,
		google.ConfigProjectId: env.GCPProjectId,
	})
}

func (g *gcp) ConfigValid(env *environment.Env) bool {
	if env.GCPJson == "" || env.GCPProjectId == "" {
		return false
	}
	return true
}

func (g *gcp) getGCPJson(env *environment.Env) (string, error) {
	if g.json != nil {
		return *g.json, nil
	}

	bytes, err := ioutil.ReadFile(env.GCPJson)
	if err != nil {
		return "", fmt.Errorf("error reading google json path: " + env.GCPJson)
	}

	json := string(bytes)

	g.json = &json

	return json, nil
}
