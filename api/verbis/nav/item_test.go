// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItems_HasItems(t *testing.T) {
	tt := map[string]struct {
		input Items
		want  bool
	}{
		"True": {
			Items{Item{Href: "href"}},
			true,
		},
		"False": {
			Items{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasItems()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestItems_Length(t *testing.T) {
	tt := map[string]struct {
		input Items
		want  int
	}{
		"Zero": {
			Items{},
			0,
		},
		"One": {
			Items{Item{Href: "href"}},
			1,
		},
		"Two": {
			Items{Item{Href: "href"}, Item{Href: "href"}},
			2,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Length()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestItem_HasChildren(t *testing.T) {
	tt := map[string]struct {
		input Item
		want  bool
	}{
		"True": {
			Item{Children: Items{Item{}}},
			true,
		},
		"False": {
			Item{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasChildren()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestItem_LiClasses(t *testing.T) {
	tt := map[string]struct {
		input Item
		want  string
	}{
		"Has Children": {
			Item{Children: Items{Item{}}},
			LiClass + "-has-children",
		},
		"Is Active": {
			Item{IsActive: true},
			LiClass + "-active",
		},
		"Post ID": {
			Item{PostID: &postID},
			fmt.Sprintf("%s-post-id-%d", LiClass, postID),
		},
		"Category ID": {
			Item{CategoryID: &categoryID},
			fmt.Sprintf("%s-category-id-%d", LiClass, categoryID),
		},
		"All": {
			Item{Children: Items{Item{}}, IsActive: true, PostID: &postID, CategoryID: &categoryID},
			fmt.Sprintf("%s-has-children %s-active %s-post-id-1 %s-category-id-1", LiClass, LiClass, LiClass, LiClass),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.LiClasses()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestItem_HasDownload(t *testing.T) {
	tt := map[string]struct {
		input Item
		want  bool
	}{
		"True": {
			Item{Download: "download.jpg"},
			true,
		},
		"False": {
			Item{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasDownload()
			assert.Equal(t, test.want, got)
		})
	}
}
