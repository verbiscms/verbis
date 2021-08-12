// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

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
	// The link href value of the item, this could be
	// a post URL, category URL or an external link.
	Href string `json:"href"`
	// The text value of the link. If the text is empty
	// the post title will be used (if it's not
	// external).
	Text string `json:"text"`
	// Is true when the current page being viewed is equal
	// to the item.
	IsActive bool `json:"is_active"`
	// Is true when the item has another navigation menu
	// below it.
	HasChildren bool `json:"has_children"`
	// The title of the link, can be empty.
	Title string `json:"title"`
	// Children contains a slice of items recursively. It can
	// be used to nest additional menus.
	Children Items `json:"children"`
	// Optional Post ID that can be attached to the item.
	PostID *int `json:"post_id"`
	// Optional Category ID that can be attached to the item.
	CategoryID *int `json:"category_id"`
	// Is true if the link is marked as external (i.e. not a
	// Post or Category).
	External bool `json:"external"`
	// Is true if the link should open in a new tab or
	// window.
	NewTab bool `json:"new_tab"`
	// Specifies the relationship between the current page
	// and the item. See https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/rel
	// for more details.
	Rel Relative `json:"rel"`
	// Specifies specifies an optional link to download a file.
	Download string `json:"download"`
}

// HasDownload determines if the item has a download
// link attached to it.
func (i *Item) HasDownload() bool {
	return i.Download != ""
}
