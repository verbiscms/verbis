// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"regexp"
	"strings"

	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// MediaRepository defines methods for Media to interact with the database
type MediaRepository interface {
	Get(meta params.Params) (domain.MediaItems, int, error)
	GetByID(id int) (domain.Media, error)
	GetByName(name string) (domain.Media, error)
	GetByURL(url string) (string, domain.Mime, error)
	Serve(uploadPath string, acceptWeb bool) ([]byte, domain.Mime, error)
	Upload(file *multipart.FileHeader, token string) (domain.Media, error)
	Validate(file *multipart.FileHeader) error
	Update(m *domain.Media) error
	Delete(id int) error
	Exists(name string) bool
	Total() (int, error)
}

// MediaStore defines the data layer for Media
type MediaStore struct {
	*StoreCfgOld
	optionsModel OptionsRepository
	userModel    UserRepository
	options      domain.Options
}

// newMedia - Construct
func newMedia(cfg *StoreCfgOld) *MediaStore {
	ms := &MediaStore{
		StoreCfgOld:  cfg,
		optionsModel: cfg.Options,
		userModel:    newUser(cfg),
	}
	ms.getOptionsStruct()
	return ms
}

// getOptionsStruct - Init the model with options
func (s *MediaStore) getOptionsStruct() {
	const op = "MediaRepository.getOptionsStruct"
	s.options = s.optionsModel.GetStruct()
}

// Get all media
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no media available.
func (s *MediaStore) Get(meta params.Params) (domain.MediaItems, int, error) {
	const op = "MediaRepository.Get"

	var m domain.MediaItems
	q := `SELECT id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at,
    CASE WHEN title IS NULL THEN '' ELSE title END AS 'title',
    CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt',
    CASE WHEN description IS NULL THEN '' ELSE description END AS 'description'
	FROM media`

	countQ := "SELECT COUNT(*) FROM media"

	// Apply filters to total and original query
	filter, err := filterRows(s.DB, meta.Filters, "media")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply order
	q += fmt.Sprintf(" ORDER BY media.%s %s", meta.OrderBy, meta.OrderDirection)

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	// Select media
	if err := s.DB.Select(&m, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get media", Operation: op, Err: err}
	}

	// Return not found error if no forms are available
	if len(m) == 0 {
		return domain.MediaItems{}, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No media available", Operation: op}
	}

	// Count the total number of media
	var total int
	if err := s.DB.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of media items", Operation: op, Err: err}
	}

	return m, total, nil
}

// GetByID returns a media item by ID
// Returns errors.NOTFOUND if the media item was not found by the given ID.
func (s *MediaStore) GetByID(id int) (domain.Media, error) {
	const op = "MediaRepository.GetByID"
	var m domain.Media
	q := `SELECT id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at,
    CASE WHEN title IS NULL THEN '' ELSE title END AS 'title',
    CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt',
    CASE WHEN description IS NULL THEN '' ELSE description END AS 'description'
	FROM media 
	WHERE id = ? LIMIT 1`
	if err := s.DB.Get(&m, q, id); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the ID: %d", id), Operation: op, Err: err}
	}
	return m, nil
}

// Gets a media file by the file name
// Returns errors.NOTFOUND if the media item was not found by the given name.
func (s *MediaStore) GetByName(name string) (domain.Media, error) {
	const op = "MediaRepository.GetByName"
	var m domain.Media
	q := `SELECT id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at,
    CASE WHEN title IS NULL THEN '' ELSE title END AS 'title',
    CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt',
    CASE WHEN description IS NULL THEN '' ELSE description END AS 'description'
	FROM media 
	WHERE name = ? LIMIT 1`
	if err := s.DB.Get(&m, q, name); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the name: %s", name), Operation: op, Err: err}
	}
	return m, nil
}

// GetByURL Obtains a media file by the url from the database
// Returns errors.NOTFOUND if the media item was not found by the given url.
func (s *MediaStore) GetByURL(url string) (string, domain.Mime, error) {
	const op = "MediaRepository.GetByURL"
	var m domain.Media

	q := `SELECT id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at,
    CASE WHEN title IS NULL THEN '' ELSE title END AS 'title',
    CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt',
    CASE WHEN description IS NULL THEN '' ELSE description END AS 'description'
	FROM media `

	// Test normal size
	if err := s.DB.Get(&m, q+"WHERE url = ? LIMIT 1", url); err == nil {
		return m.FilePath + "/" + m.UUID.String(), m.Mime, nil
	}

	// Test Sizes
	err := s.DB.Get(&m, q+"WHERE sizes LIKE '%"+url+"%' LIMIT 1")
	if err != nil {
		return "", "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the url: %s", url), Operation: op, Err: fmt.Errorf("no media item exists with the url: %s", url)}
	}

	for _, v := range m.Sizes {
		if v.Url == url {
			return m.FilePath + "/" + v.UUID.String(), m.Mime, nil
		}
	}

	return "", "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the url: %s", url), Operation: op, Err: fmt.Errorf("no media item exists with the url: %s", url)}
}

// Serve is responsible for serving the correct data to the front end
// Returns errors.NOTFOUND if the media item was not found.
func (s *MediaStore) Serve(uploadPath string, acceptWebP bool) ([]byte, domain.Mime, error) {
	const op = "MediaRepository.Serve"

	// NOTE: Not concurrent safe
	// s.getOptionsStruct()

	path, mimeType, err := s.GetByURL(uploadPath)
	if err != nil {
		return nil, "", err
	}

	extension := files.GetFileExtension(uploadPath)
	absPath := s.Paths.Uploads + "/" + path

	var data []byte
	var found error

	if acceptWebP && s.options.MediaServeWebP {
		data, found = ioutil.ReadFile(absPath + extension + ".webp")
		if found != nil {
			data, found = ioutil.ReadFile(absPath + extension)
		} else {
			mimeType = "image/webp"
		}
	} else {
		data, found = ioutil.ReadFile(absPath + extension)
	}

	if found != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("File does not exist with the path: %v", uploadPath), Operation: op, Err: err}
	}

	return data, mimeType, nil
}

// Upload the media files, return bool is for server or user error
// Returns errors.INTERNAL if the uploaded file failed to save.
func (s *MediaStore) Upload(file *multipart.FileHeader, token string) (domain.Media, error) {
	const op = "MediaRepository.Upload"

	s.getOptionsStruct()

	// Get the current logged in user
	user, err := s.userModel.CheckToken(token)
	if err != nil {
		return domain.Media{}, nil
	}

	// E.G  /Users/admin/cms/storage/uploads/2021/1, /2021/1
	absPath, path := s.createDirectory()

	// E.G: Image20@.png
	name := file.Filename

	// E.G: .png
	extension := files.GetFileExtension(name)

	// E.G: image.png
	cleanName := s.processFileName(name, extension)

	// E.G: 53252e77-308a-4587-a078-637bf1b0e186
	key := uuid.New()

	// E.G image/png
	mimeType, _ := mime.TypeByFile(file)

	// Save the uploaded file
	if err := files.Save(file, absPath+"/"+key.String()+extension); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "Could not save the media file, please try again", Operation: op, Err: err}
	}

	// Convert to WebP
	if s.options.MediaConvertWebP && mimeType == "image/jpeg" || mimeType == "image/png" {
		go webp.Convert(absPath+"/"+key.String()+extension, s.options.MediaCompression)
	}

	// Resize
	sizes := s.saveResizedImages(file, cleanName, path, mimeType, extension)

	// Insert into the database
	dm, err := s.insert(key, cleanName+extension, path, file.Size, mimeType, sizes, user.Id)
	if err != nil {
		return domain.Media{}, err
	}

	return dm, nil
}

// Validate the file before uploading
// Returns errors.INVALID if the file was not in the whitelist or
// the file was too big.
func (s *MediaStore) Validate(file *multipart.FileHeader) error {
	const op = "MediaRepository.Validate"

	s.getOptionsStruct()

	mimeType, err := mime.TypeByFile(file)
	if err != nil {
		return err
	}

	valid := mime.IsValidMime(s.Config.Media.AllowedFileTypes, mimeType)
	if !valid {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The %s mime type, is not in the whitelist for uploading.", mimeType), Operation: op, Err: err}
	}

	fileSize := file.Size / 1024
	if fileSize > s.options.MediaUploadMaxSize && s.options.MediaUploadMaxSize != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The file exceeds the maximum size restriction of %vkb.", s.options.MediaUploadMaxSize), Operation: op, Err: err}
	}

	io, err := file.Open()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error opening file", Operation: op, Err: err}
	}

	img, _, err := image.Decode(io)
	if err != nil {
		return nil // Is not an image
	}

	defer io.Close()

	if int64(img.Bounds().Max.X) > s.options.MediaUploadMaxWidth && s.options.MediaUploadMaxWidth != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The image exceeds the upload max width of %vpx.", s.options.MediaUploadMaxWidth), Operation: op, Err: err}
	}

	if int64(img.Bounds().Max.Y) > s.options.MediaUploadMaxHeight && s.options.MediaUploadMaxHeight != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The image exceeds the upload max height of %vpx.", s.options.MediaUploadMaxHeight), Operation: op, Err: err}
	}

	return nil
}

// Inserts a media item into the database
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *MediaStore) insert(uniq uuid.UUID, name, filePath string, fileSize int64, mimeType string, sizes domain.MediaSizes, userID int) (domain.Media, error) {
	const op = "MediaRepository.insert"

	m := domain.Media{
		UUID:        uniq,
		Url:         s.getURL() + "/" + name,
		Title:       "",
		Description: "",
		Alt:         "",
		FilePath:    filePath,
		FileSize:    fileSize,
		FileName:    name,
		Sizes:       sizes,
		Mime:        domain.Mime(mimeType),
		UserId:      userID,
	}

	q := "INSERT INTO media (uuid, url, title, alt, description, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.DB.Exec(q, m.UUID, m.Url, m.Title, m.Alt, m.Description, m.FilePath, m.FileSize, m.FileName, m.Sizes, m.Mime, m.UserId)

	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the new media item with the name: %v", name), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created media with the name: %v", m.FileName), Operation: op, Err: err}
	}
	m.Id = int(id)

	return m, nil
}

// Update the media item (title, alt & description)
// Returns errors.NOTFOUND if the media item was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *MediaStore) Update(m *domain.Media) error {
	const op = "MediaRepository.Update"

	_, err := s.GetByID(m.Id)
	if err != nil {
		return err
	}

	q := "UPDATE media SET title = ?, alt = ?, description = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.DB.Exec(q, m.Title, m.Alt, m.Description, m.Id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the media item with the ID: %v", m.Id), Operation: op, Err: err}
	}

	// Clear the cache
	cache.Store.Delete(m.Url)

	return nil
}

// Delete the record from the database and all files
// Returns errors.NOTFOUND if the media item was not found.
// Returns errors.INTERNAL if any file (original, webp or any sizes) were not deleted.
// Or if the SQL query was invalid
func (s *MediaStore) Delete(id int) error {
	const op = "MediaRepository.Delete"

	s.getOptionsStruct()

	m, err := s.GetByID(id)
	if err != nil {
		return err
	}

	extension := files.GetFileExtension(m.Url)

	// Delete entry from database
	if _, err := s.DB.Exec("DELETE FROM media WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete media item with the ID: %v", id), Operation: op, Err: err}
	}

	// Delete the main file
	go files.CheckAndDelete(s.Paths.Uploads + m.FilePath + "/" + m.UUID.String() + extension)
	go files.CheckAndDelete(s.Paths.Uploads + m.FilePath + "/" + m.UUID.String() + extension + ".webp")

	// Delete the sizes and webp versions if stored
	for _, v := range m.Sizes {
		filePath := m.FilePath + "/" + v.UUID.String() + extension
		go files.CheckAndDelete(s.Paths.Uploads + filePath)
		go files.CheckAndDelete(s.Paths.Uploads + filePath + ".webp")
	}

	// Check if the file deleted was the one stored in the site logo
	if m.Url == s.options.SiteLogo {
		logo, _ := json.Marshal(api.App.Logo)
		if err := s.optionsModel.Update("site_logo", logo); err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Could not update the site logo", Operation: op, Err: err})
		}
	}

	return nil
}

// Exists Checks if a media items exists by the given name
func (s *MediaStore) Exists(name string) bool {
	const op = "MediaRepository.Exists"
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM media WHERE file_name = ?)", name).Scan(&exists)
	return exists
}

// Total gets the total number of media items
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *MediaStore) Total() (int, error) {
	const op = "MediaRepository.Total"
	var total int
	if err := s.DB.QueryRow("SELECT COUNT(*) FROM media").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of media items", Operation: op, Err: err}
	}
	return total, nil
}

// saveResizedImages saves all of the resized images and returns
// an array of media DB sizes if successful.
func (s *MediaStore) saveResizedImages(file *multipart.FileHeader, name, path, mimeType, extension string) domain.MediaSizes {
	const op = "MediaRepository.saveResizedImages"

	s.getOptionsStruct()

	savedSizes := make(domain.MediaSizes)
	if mimeType == "image/png" || mimeType == "image/jpeg" {
		for key, size := range s.options.MediaSizes {
			mediaUUID := uuid.New()
			fileName := name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + extension

			if err := s.processImageSize(file, s.Paths.Uploads+"/"+path+"/"+mediaUUID.String(), mimeType, size); err == nil {
				savedSizes[key] = domain.MediaSize{
					UUID:     mediaUUID,
					Url:      s.getURL() + "/" + fileName,
					Name:     fileName,
					SizeName: size.Name,
					FileSize: int64(files.FileSize(path + "/" + mediaUUID.String() + extension)),
					Width:    size.Width,
					Height:   size.Height,
					Crop:     size.Crop,
				}
			}
		}
	}

	return savedSizes
}

// processImageSize processes image sizes, convert WebPs and saves various image sizes based on configuration
// Returns errors.INTERNAL if the image was unable to be saved or decoded.
func (s *MediaStore) processImageSize(file *multipart.FileHeader, filePath, mimeType string, size domain.MediaSize) error {
	const op = "MediaRepository.processImageSize"

	s.getOptionsStruct()

	// PNG Mime
	if mimeType == "image/png" {
		filePath = filePath + ".png"

		decodedImage, err := s.decodeImage(file, mimeType)
		if err != nil {
			return err
		}
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.PNGCompressionLevel(png.CompressionLevel(s.options.MediaCompression))); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: "Could not save the resized image", Operation: op, Err: err}
		}

		if s.options.MediaConvertWebP {
			// 100 - compression level
			// go webp.Convert(filePath, 100 - s.options.MediaCompression)
			go webp.Convert(filePath, s.options.MediaCompression)
		}
	}

	// Jpg Mime
	if mimeType == "image/jpeg" || mimeType == "image/jp2" {
		filePath = filePath + ".jpg"

		decodedImage, err := s.decodeImage(file, mimeType)
		if err != nil {
			return err
		}
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.JPEGQuality(s.options.MediaCompression)); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: "Could not save the resized image", Operation: op, Err: err}
		}

		if s.options.MediaConvertWebP {
			go webp.Convert(filePath, s.options.MediaCompression)
		}
	}

	return nil
}

// resizeImage Resize the image based width, height & crop
func resizeImage(srcImage image.Image, width int, height int, crop bool) image.Image {
	if crop {
		return imaging.Fill(srcImage, width, height, imaging.Center, imaging.Lanczos)
	} else {
		return imaging.Resize(srcImage, width, height, imaging.Lanczos)
	}
}

// createDirectory creates the media directory year path if the organise year variable in the media
// store is set to true. Date and year folders are created recursively.
func (s *MediaStore) createDirectory() (string, string) {
	const op = "MediaRepository.createDirectory"

	s.getOptionsStruct()

	if !s.options.MediaOrganiseDate {
		return s.Paths.Uploads, ""
	} else {
		t := time.Now()
		datePath := t.Format("2006") + "/" + t.Format("01")
		path := s.Paths.Uploads + "/" + datePath

		if _, err := os.Stat(path); os.IsNotExist(err) {
			_ = os.MkdirAll(path, os.ModePerm)
		}
		return path, datePath
	}
}

// Get the public url of the file according to date and month if the organise
// year variable in the media store is set to true. If not the function will
// return the public uploads folder by default.
func (s *MediaStore) getURL() string {
	if !s.options.MediaOrganiseDate {
		return s.Config.Media.UploadPath
	} else {
		t := time.Now()
		return s.Config.Media.UploadPath + "/" + t.Format("2006") + "/" + t.Format("01")
	}
}

// decodeImage decodes the image from a file dependant on the mime type
// Returns errors.INTERNAL if the file was unable to be decoded or the file
// was unable to be opened.
// TODO: ptrToRefParam: consider to make non-pointer type for `*image.Image`
func (s *MediaStore) decodeImage(file *multipart.FileHeader, mimeType string) (*image.Image, error) { //nolint
	const op = "MediaRepository.decodeImage"

	reader, err := file.Open()
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to open the file with the filename: %s", file.Filename), Operation: op, Err: err}
	}

	if mimeType == "image/png" {
		pngFile, err := png.Decode(reader)
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not decode the image with the filename: %s", file.Filename), Operation: op, Err: err}
		}
		return &pngFile, nil
	} else if mimeType == "image/jpeg" || mimeType == "image/jp2" {
		jpgFile, err := jpeg.Decode(reader)
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not decode the image with the filename: %s", file.Filename), Operation: op, Err: err}
		}
		return &jpgFile, nil
	}

	return nil, &errors.Error{Code: errors.INTERNAL, Message: "Something went wrong decoding the image", Operation: op, Err: err}
}

// Process file name
func (s *MediaStore) processFileName(file, extension string) string {
	// Remove the file extension
	name := files.RemoveFileExtension(file)

	// Clean the file
	var cleanedFile string
	cleanedFile = strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = reg.ReplaceAllString(cleanedFile, "")
	cleanedFile = strings.ToLower(cleanedFile)

	// Check if the file exists and add a version number, continue if not
	version := 0
	for {
		if version == 0 {
			if exists := s.Exists(cleanedFile + extension); !exists {
				break
			}
		} else {
			if exists := s.Exists(cleanedFile + "-" + strconv.Itoa(version) + extension); !exists {
				cleanedFile = cleanedFile + "-" + strconv.Itoa(version)
				break
			}
		}
		version++
	}

	return cleanedFile
}
