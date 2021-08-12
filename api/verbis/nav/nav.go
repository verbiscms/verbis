// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"encoding/json"
	"fmt"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Getter defines the method used for obtaining
// nav menus from Verbis.
type Getter interface {
	Get(args Args) (Items, error)
}

// Service defines the helper for obtaining navigational
// menus within Verbis.
type Service struct {
	deps *deps.Deps
	post *domain.PostDatum
	nav  Nav
}

// Nav defines the type for obtaining navigational
// menus. It contains a key which
type Nav map[string]Items

// Items defines the slice of navigational Item's.
type Items []Item

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
	Rel []string `json:"rel"`
	// Specifies specifies an optional link to download a file.
	Download string `json:"download"`
}

var (
	// ErrMenuNotFound is returned by Get() when no
	// menu was found with the given name.
	ErrMenuNotFound = errors.New("no menu found")
)

// New creates a new navigation service.
// Returns errors.INTERNAL if there was an error converting
// the NavMenus field in the options from a map[string]interface
// to a nav item.
func New(d *deps.Deps, post *domain.PostDatum) (*Service, error) {
	const op = "Nav.New"

	m, err := json.Marshal(d.Options.NavMenus)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling navigation menus", Operation: op, Err: err}
	}

	nav := Nav{}
	err = json.Unmarshal(m, &nav)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling navigation menus", Operation: op, Err: err}
	}

	return &Service{
		deps: d,
		post: post,
		nav:  nav,
	}, nil
}

// Get parses the arguments and retrieves the navigational Items
// from the database.
func (s *Service) Get(args Args) (Items, error) {
	const op = "Nav.Get"

	opts, err := args.toOptions()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error converting arguments to navigation options", Operation: op, Err: err}
	}

	err = opts.Validate()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error validating navigation options", Operation: op, Err: err}
	}

	items, err := s.getNavItems(opts)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error obtaining navigation items", Operation: op, Err: err}
	}

	return items, nil
}

// getNavItems obtains the navigational items by comparing
// the menu name with the options.
// Returns ErrMenuNotFound if lookup failed.
func (s *Service) getNavItems(opts Options) (Items, error) {
	for menu, nav := range s.nav {
		if menu == opts.Menu {
			return s.processItems(nav), nil
		}
	}
	return nil, fmt.Errorf("%s: %s", ErrMenuNotFound, opts.Menu)
}

// processItems iterates over the navigational items and processes
// them. The function is recursively called if the item has
// any children.
func (s *Service) processItems(items Items) Items {
	for idx, item := range items {
		// Obtain the child navigation menu if there is
		// one present.
		if item.Children != nil {
			items[idx].HasChildren = true
			items[idx].Children = s.processItems(item.Children)
		}

		// If the item is external, we can presume there is
		// no post or category attached to it, continue
		// the loop.
		if item.External {
			continue
		}

		if item.PostID != nil {
			// Obtain the post if there is one attached and
			// assign the permalink to the href value.
			post, err := s.deps.Store.Posts.Find(*item.PostID, false)
			if err != nil {
				continue // We can assume it's been deleted or removed.
			}
			items[idx].Href = post.Permalink

			// Check if the link text is empty, if it is
			// assign the post title.
			if item.Title == "" {
				items[idx].Title = post.Title
			}

			// Check if the current navigational item is the same
			// as the one being viewed and mark it as active.
			if s.post.ID == *item.PostID {
				items[idx].IsActive = true
			}
		}
	}

	return items
}
