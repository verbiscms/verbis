package api

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// MediaHandler defines methods for Media Items to interact with the server
type MediaHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Upload(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Media defines the handler for Posts
type Media struct {
	store  *models.Store
	config config.Configuration
}

// newMedia - Construct
func NewMedia(m *models.Store, config config.Configuration) *Media {
	return &Media{
		store:  m,
		config: config,
	}
}

// Get all media items
//
// Returns 200 if there are no media items or success.
// Returns 500 if there was an error getting the media items.
// Returns 400 if there was conflict or the request was invalid.
func (c *Media) Get(g *gin.Context) {
	const op = "MediaHandler.Get"

	params := http.NewParams(g).Get()
	media, total, err := c.store.Media.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(params, total)

	handler.Respond(g, 200, "Successfully obtained media", media, pagination)
}

// Get By ID
//
// Returns 200 if the media items were obtained.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there as an error obtaining the media items.
func (c *Media) GetById(g *gin.Context) {
	const op = "MediaHandler.GetById"

	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		handler.Respond(g, 400, "Pass a valid number to obtain the media item by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	media, err := c.store.Media.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully obtained media item with ID: "+paramId, media)
}

// Upload - if there were no files attached to the body,
// more than 1 attached to the body or the validation failed.
//
// Returns 401 if the user wasn't authenticated.
// Returns 415 if the media item failed to validate.
// Returns 200 if the media item was successfully uploaded.
// Returns 500 if there as an error uploading the media item.
// Returns 400 if the file length was incorrect or there were no files.
func (c *Media) Upload(g *gin.Context) {
	const op = "MediaHandler.Upload"

	form, err := g.MultipartForm()
	if err != nil {
		handler.Respond(g, 400, "No files attached to the upload", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	files := form.File["file"]

	if len(files) > 1 {
		handler.Respond(g, 400, "Files are only permitted to be uploaded one at a time", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("too many files uploaded at once"), Operation: op})
		return
	}

	if len(files) == 0 {
		handler.Respond(g, 400, "Attach a file to the request to be uploaded", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no files attached to upload"), Operation: op})
		return
	}

	if err := c.store.Media.Validate(files[0]); err != nil {
		handler.Respond(g, 415, errors.Message(err), err)
		return
	}

	media, err := c.store.Media.Upload(files[0], g.Request.Header.Get("token"))
	if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully uploaded media item", media)
}

// Update
//
// Returns 200 if the media item was updated successfully.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there was an error updating the media item.
func (c *Media) Update(g *gin.Context) {
	const op = "MediaHandler.Update"

	var m domain.Media
	if err := g.ShouldBindJSON(&m); err != nil {
		handler.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		handler.Respond(g, 400, "A valid ID is required to update the media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	m.Id = id

	err = c.store.Media.Update(&m)
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully updated media item with ID: "+strconv.Itoa(id), m)
}

// Delete
//
// Returns 200 if the media item was deleted.
// Returns 500 if there was an error updating the media item.
// Returns 400 if the the media item wasn't found or no ID was passed.
func (c *Media) Delete(g *gin.Context) {
	const op = "MediaHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		handler.Respond(g, 400, "A valid ID is required to delete a media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.store.Media.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully deleted media item with ID: "+strconv.Itoa(id), nil)
}
