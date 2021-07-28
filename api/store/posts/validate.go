// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// validate checks to see if the slug already exists, the
// post has a valid page template attached to it, and a
// valid page layout.
// Returns nil if all checks are passed.
func (s *Store) validate(p *domain.PostCreate, checkSlug bool) error {
	if checkSlug {
		err := s.validateSlug(p)
		if err != nil {
			return err
		}
	}

	cfg := config.Get()

	resource, ok := cfg.Resources[p.Resource]
	if ok {
		if resource.Hidden {
			return nil
		}
	}

	//err := s.validatePageTemplate(cfg, p)
	//if err != nil {
	//	return err
	//}
	//
	//err = s.validatePageLayout(cfg, p)
	//if err != nil {
	//	return err
	//}

	return nil
}

// validateSlug compares the slug passed to see if the
// slug has already been taken, by looking up the
// post category and resource.
// Returns ErrPostsExists if there is an existing slug.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) validateSlug(p *domain.PostCreate) error {
	const op = "PostStore.ValidateSlug"

	q := s.Builder().
		Select(s.Schema()+TableName+".id").
		From(s.Schema()+TableName).
		Where(s.Schema()+TableName+".slug", "=", p.Slug)

	if p.Category != nil {
		q.LeftJoin(s.Schema()+"post_categories", "pc", s.Schema()+"posts.id = "+s.Schema()+"pc.post_id").
			LeftJoin(s.Schema()+"categories", "c", "pc.category_id = c.id").
			Where(s.Schema()+"c.id", "=", p.Category)
	}

	if p.Resource != "" {
		q.Where(s.Schema()+TableName+".resource", "=", p.Resource)
	} else {
		q.Where(s.Schema()+TableName+".resource", "=", "")
	}

	var exists bool
	err := s.DB().QueryRow(q.Exists()).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	if exists {
		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, the slug already exists", Operation: op, Err: ErrPostsExists}
	}

	return nil
}

//
//// validatePageTemplate checks to see if a page template
//// is found in the current theme.
//// Returns nil if it was found.
//// Returns ErrNoPageTemplate if the page template was not found.
//func (s *Store) validatePageTemplate(cfg *domain.ThemeConfig, p *domain.PostCreate) error {
//	const op = "PostStore.ValidatePageTemplate"
//
//	tpl, err := s.ThemeService.Templates(cfg.Theme.Name)
//	if err != nil {
//		return err
//	}
//
//	found := false
//	for _, v := range tpl {
//		if v.Key == p.PageTemplate {
//			found = true
//		}
//	}
//
//	if !found {
//		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, no page template exists: " + p.PageTemplate, Operation: op, Err: ErrNoPageTemplate}
//	}
//
//	return nil
//}
//
//// validatePageLayout Checks to see if a page layout is
//// found in the current theme.
//// Returns nil if it was found.
//// Returns ErrNoPageLayout if the page layout was not found.
//func (s *Store) validatePageLayout(cfg *domain.ThemeConfig, p *domain.PostCreate) error {
//	const op = "PostStore.ValidatePageLayout"
//
//	tpl, err := s.ThemeService.Layouts(cfg.Theme.Name)
//	if err != nil {
//		return err
//	}
//
//	found := false
//	for _, v := range tpl {
//		if v.Key == p.PageLayout {
//			found = true
//		}
//	}
//
//	if !found {
//		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, no page layout exists: " + p.PageLayout, Operation: op, Err: ErrNoPageLayout}
//	}
//
//	return nil
//}
