package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MediaController struct {
	controller	Controller
	mediaModel 	models.MediaRepository
	userModel 	models.UserRepository
}

type MediaHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Upload(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Construct
func newMedia(m models.MediaRepository, um models.UserRepository) *MediaController {
	return &MediaController{
		mediaModel: m,
		userModel: um,
	}
}

// Get All
func (c *MediaController) Get(g *gin.Context) {
	params := http.GetParams(g)
	media, err := c.mediaModel.GetAll(params)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	successMsg := "Successfully obtained media"
	if len(media) == 0 {
		successMsg = "No media available"
	}

	totalAmount, err := c.mediaModel.Total()
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}
	pagination := http.GetPagination(params, totalAmount)

	Respond(g, 200, successMsg, media, *pagination)
}

// Get By ID
func (c *MediaController) GetById(g *gin.Context) {
	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Respond(g, 500,  err.Error(), nil)
		return
	}

	// Get the post
	media, err := c.mediaModel.GetById(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully obtained media with the ID: " + paramId, media)
}

// Upload/Create
func (c *MediaController) Upload(g *gin.Context) {
	form, err := g.MultipartForm()
	if err != nil {
		Respond(g, 400, "No files attached to the upload", nil)
		return
	}
	files := form.File["file"]

	if len(files) > 1 {
		Respond(g, 400, "Files are only permitted to be uploaded one at a time", nil)
		return
	}

	if err := c.mediaModel.Validate(files[0]); err != nil {
		Respond(g, 415, err.Error(), nil)
	}

	token := g.Request.Header.Get("token")
	user, err := c.userModel.CheckToken(token)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
	}

	media, err := c.mediaModel.Upload(files[0], user.Id)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully uploaded media", media)
}

// Update
func (c *MediaController) Update(g *gin.Context) {
	var media domain.Media
	if err := g.ShouldBindJSON(&media); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500,"An ID is required to update the media item", nil)
		return
	}
	media.Id = id

	err = c.mediaModel.Update(&media)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully updated media item with the ID: " + strconv.Itoa(id), media)
}

// Delete
func (c *MediaController) Delete(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500, err.Error(), nil)
	}

	err = c.mediaModel.Delete(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully deleted media item with the ID: " + strconv.Itoa(id), nil)
}