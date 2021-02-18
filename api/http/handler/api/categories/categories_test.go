// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"testing"
)


type CategoriesTestSuite struct {
	suite.Suite
}

func TestCategories(t *testing.T) {
	suite.Run(t, new(CategoriesTestSuite))
}

// Mock
//
// Us a helper to obtain a mock categories handler
// for testing.
func (t *CategoriesTestSuite) Mock(m models.CategoryRepository) *Categories {
	return &Categories{
		Deps: &deps.Deps{
			Store: &models.Store{
				Categories: m,
			},
		},
	}
}