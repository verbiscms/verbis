// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// TODO:
// Check if the partial works within the tpl file.
// Better naming conventions for interface and structs
// https://vuejsdevelopers.com/2017/10/23/vue-js-tree-menu-recursive-components/

// Getter defines the method used for obtaining
// nav menus from Verbis.
type Getter interface {
	Get(args Args) (Menu, error)
	GetFromOptions(opts Options) (Menu, error)
}

// Service defines the helper for obtaining navigational
// menus within Verbis.
type Service struct {
	deps *deps.Deps
	post *domain.PostDatum
	nav  Menus
}

// Menus defines the type for obtaining navigational
// menus. It contains a key which is mapped
// to the ID defined in the theme config.
type Menus map[string]Nav

// TODO, Should this be called nav?
type Nav struct {
	Name  string `json:"name"`
	Items Items  `json:"items"`
}

// Menu defines the data retrieved by calling Get.
// It returns a slice if Item's as well as the
// options passed to it.
type Menu struct {
	Name    string
	Items   Items
	Options Options
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
	const op = "Menus.New"

	m, err := json.Marshal(d.Options.NavMenus)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling navigation menus", Operation: op, Err: err}
	}

	nav := Menus{}
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

// Get parses the arguments and retrieves the navigational
// Items from the database.
func (s *Service) Get(args Args) (Menu, error) {
	opts, err := args.ToOptions()
	if err != nil {
		return Menu{}, err
	}
	return s.GetFromOptions(opts)
}

// GetFromOptions is an alias for retrieving nav items
// by options as opposed to passing a
// map[string]interface{}.
func (s *Service) GetFromOptions(opts Options) (Menu, error) {
	const op = "Menus.Get"

	err := opts.Validate()
	if err != nil {
		return Menu{}, &errors.Error{Code: errors.INVALID, Message: "Error validating navigation options", Operation: op, Err: err}
	}

	// Try and obtain the cached navigational items
	// if it exists.
	c, err := s.deps.Cache.Get(context.Background(), "nav-menu-"+opts.Menu)
	if err == nil {
		m, ok := c.(Menu)
		// If the cast was successful and the options
		// remain the same, return the cached
		// version.
		if ok && m.Options == opts {
			return m, nil
		}
	}

	items, name, err := s.getNavItems(opts)
	if err != nil {
		return Menu{}, &errors.Error{Code: errors.INVALID, Message: "Error obtaining navigation items", Operation: op, Err: err}
	}

	m := Menu{
		Name:    name,
		Options: opts,
		Items:   items,
	}

	// Cache the results forever.
	go s.deps.Cache.Set(context.Background(), "nav-menu-"+opts.Menu, m, cache.Options{
		Expiration: cache.RememberForever,
		Tags:       []string{"menus", "options"},
	})

	return m, nil
}

// getNavItems obtains the navigational items by comparing
// the menu name with the options.
// Returns ErrMenuNotFound if lookup failed.
func (s *Service) getNavItems(opts Options) (Items, string, error) {
	for menu, nav := range s.nav {
		if menu == opts.Menu {
			return s.processItems(nav.Items, opts.Depth, 1), nav.Name, nil
		}
	}
	return nil, "", fmt.Errorf("%s: %s", ErrMenuNotFound, opts.Menu)
}

// processItems iterates over the navigational items and processes
// them. The function is recursively called if the item has
// any children.
func (s *Service) processItems(items Items, depth, currDepth int) Items {
	for idx, item := range items {
		// Assign the item the current depth.
		items[idx].Depth = currDepth

		// Obtain the child navigation menu if there is
		// one present.
		if item.Children != nil {
			// If the depth in the options is greater
			// than the depth passed as argument, or the
			// depth is 0 (all items).
			// we need to exit, and not run
			// this conditional.
			if depth == 0 || currDepth < depth {
				items[idx].Children = s.processItems(item.Children, depth, currDepth+1)
			} else {
				// Reset the children as there shouldn't be
				// any here.
				items[idx].Children = nil
			}
		}

		// If the item is external, we can presume there is
		// no post or category attached to it, continue
		// the loop.
		if item.External {
			continue
		}

		// Assign post attributes.
		if item.PostID != nil {
			// Obtain the post if there is one attached and
			// assign the permalink to the href value.
			post, err := s.deps.Store.Posts.Find(*item.PostID, false)
			if err != nil {
				items[idx].Invalid = true
				continue // We can assume it's been deleted or removed, but mark it as invalid.
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

		// TODO: Assign category attributes.
	}

	return items
}
