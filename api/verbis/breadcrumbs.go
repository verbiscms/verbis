// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbis

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"regexp"
	"strings"
)

// Breadcrumbs represent the navigation aid for displaying
// hierarchical content within verbis. Various options
// are included within the type as well as the
// trail of crumbs (a slice of Items).
type Breadcrumbs struct {
	Enabled   bool   `json:"enabled"`
	Title     string `json:"title"`
	Separator string `json:"separator"`
	Items     Items  `json:"crumbs"`
}

// Item represents a singular crumb. It includes the link,
// link text, (the post title or a clean version of the
// URI segment if no post was found), the position of
// the breadcrumb and weather or not the item was
// successfully retrieved from the database.
type Item struct {
	Link     string `json:"link"`
	Text     string `json:"text"`
	Position int    `json:"position"`
	Found    bool   `json:"found"`
	Active   bool   `json:"active"`
}

// Items defines the slice of crumbs within the breadcrumbs
// type.
type Items []Item

// Length determines how many crumbs there within items.
func (i Items) Length() int {
	return len(i)
}

// Reverse reverse's the items slice to use the crumbs
// from back to front order such as Technology ->
// News -> Home. It returns a copy of the
// slice without modifying the original.
func (i Items) Reverse() Items {
	tmp := make(Items, len(i))
	copy(tmp, i)
	for j, k := 0, len(tmp)-1; j < k; j, k = j+1, k-1 {
		tmp[j], tmp[k] = tmp[k], tmp[j]
	}
	return tmp
}

// GetBreadcrumbs retrieves the breadcrumbs for a specific
// post. If the breadcrumbs are not enabled from the
// options, an empty Breadcrumbs type will be
// returned. The homepage is automatically
// prepended and subsequent URI's are
// split and matches are looked up
// from the database.
func GetBreadcrumbs(post *domain.PostDatum, d *deps.Deps) Breadcrumbs {
	bc := Breadcrumbs{
		Enabled:   d.Options.BreadcrumbsEnable,
		Title:     d.Options.BreadcrumbsTitle,
		Separator: d.Options.BreadcrumbsSeparator,
	}

	if !d.Options.BreadcrumbsEnable {
		bc.Enabled = false
		return bc
	}

	isHome := post.IsHomepage(d.Options.Homepage)
	if d.Options.BreadcrumbsHideHomePage && isHome {
		return bc
	}

	bc.Items = Items{
		{
			Link:     d.Options.SiteUrl,
			Text:     d.Options.BreadcrumbsHomepageText,
			Position: 1,
			Found:    true,
		},
	}

	if isHome {
		bc.Items[0].Active = true
		return bc
	}

	bc = traverseChild(post, d, bc)

	return bc
}

// traverseChild ranges over the split permalink and
// builds out the breadcrumb trail.
func traverseChild(post *domain.PostDatum, d *deps.Deps, bc Breadcrumbs) Breadcrumbs {
	b := urlSlice(post.Permalink)

	link := ""
	for i := 0; i < len(b); i++ {
		pos := i + 2
		link += "/" + b[i]
		slash := ""
		active := false

		if len(b)-1 == i {
			active = true
		}

		if d.Options.SeoEnforceSlash {
			slash = "/"
		}

		pt, err := d.Store.Posts.FindBySlug(b[i])
		if err != nil {
			bc.Items = append(bc.Items, Item{
				Link:     d.Options.SiteUrl + link + slash,
				Text:     cleanTitle(b[i]),
				Position: pos,
				Found:    false,
				Active:   active,
			})
			continue
		}

		bc.Items = append(bc.Items, Item{
			Link:     d.Options.SiteUrl + pt.Permalink,
			Text:     pt.Title,
			Position: pos,
			Found:    true,
			Active:   active,
		})
	}

	return bc
}

// urlSlice splits the permalink by a forward slash and
// cleans any values that are empty.
func urlSlice(permalink string) []string {
	var p = strings.Split(permalink, "/")
	var r []string
	for _, str := range p {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// titleRegex is the regex used to clean the post title
// upon failure of the database lookup.
var titleRegex = "[^a-zA-Z -]+"

// cleanTitle the title of the post by removing any
// unwanted characters, and converting the string
// to Title case.
func cleanTitle(title string) string {
	const op = "Breadcrumbs.cleanTitle"

	reg, err := regexp.Compile(titleRegex)
	t := title
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error compiling regex", Operation: op, Err: err})
	} else {
		t = reg.ReplaceAllString(title, "")
	}

	t = strings.ToLower(strings.ReplaceAll(t, "-", " "))
	t = strings.Title(strings.TrimSpace(t))

	return t
}
