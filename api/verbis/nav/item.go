// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"fmt"
	"strings"
)

// Items defines the slice of navigational Item's.
type Items []Item

// HasItems determines if there are any nav menu's
// to output.
func (i Items) HasItems() bool {
	return len(i) != 0
}

// Length retrieves the amount of items.
func (i Items) Length() int {
	return len(i)
}

// Item defines a singular navigation item within the
// nav menu tree. An item can embed multiple
// children.
type Item struct {
	// ID is the unique identifier of a menu item.
	ID int `json:"id"`
	// Href link value of the item, this could be a post
	// URL, category URL or an external link.
	Href string `json:"href"`
	// Text is the text value of the link. If the text is
	// empty the post title will be used (if it's not external).
	Text string `json:"text"`
	// IsActive defines if the current page being viewed is
	// equal to the item being displayed.
	IsActive bool `json:"is_active"`
	// Title of the link, can be empty.
	Title string `json:"title"`
	// Depth is the integer defining how deep the item is
	// in the nav hierarchy.
	Depth int `json:"depth"`
	// Children contains a slice of items recursively. It can
	// be used to nest additional menus.
	Children Items `json:"children"`
	// PostID - Optional Post ID that can be attached to
	// the item.
	PostID *int `json:"post_id"`
	// CategoryID - Optional Category ID that can be attached
	// to the item.
	CategoryID *int `json:"category_id"`
	// External defines if the link is marked as external (i.e.
	// not a Post or Category).
	External bool `json:"external"`
	// Classes are the CSS classes to be outputted on the
	// item. Specifically on the <li> item.
	Classes TagSlice `json:"li_classes"`
	// NewTab defines if the link should open in a new tab or
	// window.
	NewTab bool `json:"new_tab"`
	// Rel Specifies the relationship between the current page
	// and the item. See https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/rel
	// for more details.
	Rel TagSlice `json:"rel"`
	// Download specifies an optional link to download
	// a file.
	Download string `json:"download"`
	// Is true when a post has been deleted or modified.
	Invalid bool `json:"invalid"`
	// TODO:
	// Should we have type here? For example, download, post, category or external?
}

const (
	// LiClass defines the CSS class for appending
	// class names to the nav item.
	LiClass = "nav-item"
)

// HasChildren determines if an item has any children
// beneath it by comparing length.
func (i *Item) HasChildren() bool {
	return len(i.Children) > 0
}

// LiClasses returns a string of classes to be added to
// the <li> element when executing a template. These
// classes can act as helpers to target with CSS.
func (i *Item) LiClasses() string {
	var classes []string
	if i.HasChildren() {
		classes = append(classes, "has-children")
	}
	if i.IsActive {
		classes = append(classes, "active")
	}
	if i.PostID != nil {
		classes = append(classes, fmt.Sprintf("post-id-%d", *i.PostID))
	}
	if i.CategoryID != nil {
		classes = append(classes, fmt.Sprintf("category-id-%d", *i.CategoryID))
	}
	for idx, class := range classes {
		classes[idx] = LiClass + "-" + class
	}
	if i.Classes.HasTags() {
		classes = append(classes, i.Classes...)
	}
	return strings.Join(classes, " ")
}

// HasDownload determines if the item has a download
// link attached to it.
func (i *Item) HasDownload() bool {
	return i.Download != ""
}
