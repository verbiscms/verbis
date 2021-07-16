// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/google"
	"io/ioutil"
)

type gcp struct {
	json *string
}

const GCPName = "Google Cloud Storage"

var (
	GCPEnvKeys = []string{
		"STORAGE_GCP_JSON_FILE",
		"STORAGE_GCP_PROJECT_ID",
	}
)

func (g *gcp) Dial(env *environment.Env) (stow.Location, error) {
	json, err := g.getGCPJson(env)
	if err != nil {
		return nil, err
	}
	return dialler(google.Kind, stow.ConfigMap{
		google.ConfigJSON:      json,
		google.ConfigProjectId: env.GCPProjectId,
	})
}

func (g *gcp) Info(env *environment.Env) domain.StorageProviderInfo {
	sp := domain.StorageProviderInfo{
		Name:            GCPName,
		Order:           2,
		Connected:       false,
		Error:           false,
		EnvironmentSet:  false,
		EnvironmentKeys: GCPEnvKeys,
	}

	if env.GCPJson == "" && env.GCPProjectId == "" {
		sp.Error = ErrMessageConfigNotSet + domain.StorageGCP.TitleCase().String()
		return sp
	}

	sp.EnvironmentSet = true

	_, err := g.Dial(env)
	if err != nil {
		sp.Error = ErrMessageDial + err.Error()
		return sp
	}

	sp.Connected = true
	sp.Error = false

	return sp
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
