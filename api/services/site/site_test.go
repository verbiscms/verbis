// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/domain"
	mockOptions "github.com/verbiscms/verbis/api/mocks/store/options"
	mockSys "github.com/verbiscms/verbis/api/mocks/sys"
	"github.com/verbiscms/verbis/api/version"
	"testing"
)

func TestSite_Config(t *testing.T) {
	ms := &mockSys.System{}
	ms.On("LatestVersion").Return("v0.0.1")
	ms.On("HasUpdate").Return(false)

	opts := domain.Options{
		SiteTitle:       "title",
		SiteDescription: "description",
		SiteLogo:        "logo",
		SiteUrl:         "url",
	}

	mo := &mockOptions.Repository{}
	mo.On("Struct").Return(opts)

	want := domain.Site{
		Title:         opts.SiteTitle,
		Description:   opts.SiteDescription,
		Logo:          opts.SiteLogo,
		Url:           opts.SiteUrl,
		Version:       version.Version,
		RemoteVersion: "v0.0.1",
		HasUpdate:     false,
	}

	s := New(mo, ms)
	assert.Equal(t, want, s.Global())
}
