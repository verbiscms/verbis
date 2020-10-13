package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
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

// MediaController defines the handler for Posts
type MediaController struct {
	controller	Controller
	mediaModel 	models.MediaRepository
	userModel 	models.UserRepository
}

// newMedia - Construct
func newMedia(m models.MediaRepository, um models.UserRepository) *MediaController {
	return &MediaController{
		mediaModel: m,
		userModel: um,
	}
}

// Get all media items
func (c *MediaController) Get(g *gin.Context) {
	const op = "MediaHandler.Get"

	params := http.GetParams(g)
	media, err := c.mediaModel.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	totalAmount, err := c.mediaModel.Total()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	pagination := http.GetPagination(params, totalAmount)

	Respond(g, 200, "Successfully obtained media", media, pagination)
}

// Get By ID
// Returns errors.INVALID if the Id is not a string or passed.
func (c *MediaController) GetById(g *gin.Context) {
	const op = "MediaHandler.GetById"

	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Respond(g, 400,  "Pass a valid number to obtain the media item by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	media, err := c.mediaModel.GetById(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully obtained media with the ID: " + paramId, media)
}

// Upload
// Returns errors.INVALID if there were no files attached to the body,
// more than 1 attached to the body or the validation failed.
func (c *MediaController) Upload(g *gin.Context) {
	const op = "MediaHandler.Upload"

	form, err := g.MultipartForm()
	if err != nil {
		Respond(g, 400, "No files attached to the upload", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	files := form.File["file"]

	if len(files) > 1 {
		Respond(g, 400, "Files are only permitted to be uploaded one at a time", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("too many files uploaded at once"), Operation: op})
		return
	}

	if len(files) == 0 {
		Respond(g, 400, "Attach a file to the request to be uploaded", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no files attached to upload"), Operation: op})
		return
	}

	if err := c.mediaModel.Validate(files[0]); err != nil {
		Respond(g, 415, errors.Message(err), err)
		return
	}

	token := g.Request.Header.Get("token")
	user, err := c.userModel.CheckToken(token)
	if err != nil {
		Respond(g, 400, errors.Message(err), err)
		return
	}

	media, err := c.mediaModel.Upload(files[0], user.Id)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully uploaded media", media)
}

// Update
func (c *MediaController) Update(g *gin.Context) {
	const op = "MediaHandler.Update"

	var media domain.Media
	if err := g.ShouldBindJSON(&media); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500,"A valid ID is required to update the Media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	media.Id = id

	err = c.mediaModel.Update(&media)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated media item with the ID: " + strconv.Itoa(id), media)
}

// Delete
func (c *MediaController) Delete(g *gin.Context) {
	const op = "MediaHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"A valid ID is required to delete a media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	err = c.mediaModel.Delete(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted media item with the ID: " + strconv.Itoa(id), nil)
}