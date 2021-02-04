// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// CategoryHandler defines methods for categories to interact with the server
type CategoryHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Categories defines the handler for Categories
type Categories struct {
	*deps.Deps
}

// newCategories - Construct
func NewCategories(d *deps.Deps) *Categories {
	return &Categories{d}
}

// Get all categories
//
// Returns 200 if there are no categories or success.
// Returns 500 if there was an error getting the categories.
// Returns 400 if there was conflict or the request was invalid.
func (c *Categories) Get(g *gin.Context) {
	const op = "CategoryHandler.Get"

	p := params.ApiParams(g, DefaultParams).Get()

	categories, total, err := c.Store.Categories.Get(p)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(p, total)

	Respond(g, 200, "Successfully obtained categories", categories, pagination)
}

// Get By ID
//
// Returns 200 if the category was obtained.
// Returns 500 if there as an error obtaining the category.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *Categories) GetById(g *gin.Context) {
	const op = "CategoryHandler.GetById"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "Pass a valid number to obtain the category by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	category, err := c.Store.Categories.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained category with ID: "+strconv.Itoa(id), category)
}

// Create
//
// Returns 200 if the category was created.
// Returns 500 if there was an error creating the category.
// Returns 400 if the the validation failed or there was a conflict.
func (c *Categories) Create(g *gin.Context) {
	const op = "CategoryHandler.Create"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newCategory, err := c.Store.Categories.Create(&category)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created category with ID: "+strconv.Itoa(category.Id), newCategory)
}

// Update
//
// Returns 200 if the category was updated.
// Returns 500 if there was an error updating the category.
// Returns 400 if the the validation failed or the category wasn't found.
func (c *Categories) Update(g *gin.Context) {
	const op = "CategoryHandler.Update"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to update the category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	category.Id = id

	updatedCategory, err := c.Store.Categories.Update(&category)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	defer c.clearCache(updatedCategory.Id)

	Respond(g, 200, "Successfully updated category with ID: "+strconv.Itoa(category.Id), updatedCategory)
}

// Delete
//
// Returns 200 if the category was deleted.
// Returns 500 if there was an error deleting the category.
// Returns 400 if the the category wasn't found or no ID was passed.
func (c *Categories) Delete(g *gin.Context) {
	const op = "CategoryHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to delete a category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.Store.Categories.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted category with ID: "+strconv.Itoa(id), nil)
}

// clearCache
// Clear the post cache that have the given category ID
// attached to it.
func (c *Categories) clearCache(id int) {
	go func() {
		posts, _, err := c.Store.Posts.Get(params.Params{LimitAll: true}, false, "", "")
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error()
		}
		cache.ClearCategoryCache(id, posts)
	}()
}
