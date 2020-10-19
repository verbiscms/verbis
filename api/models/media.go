package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kamoljan/webp"
	log "github.com/sirupsen/logrus"
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
	Get(meta http.Params) ([]domain.Media, int, error)
	GetById(id int) (domain.Media, error)
	GetByName(name string) (domain.Media, error)
	GetByUrl(url string) (string, string, error)
	Serve(uploadPath string, acceptWeb bool) ([]byte, string, error)
	Upload(file *multipart.FileHeader, userId int) (domain.Media, error)
	Validate(file *multipart.FileHeader) error
	Update(m *domain.Media) error
	Delete(id int) error
	Exists(name string) bool
	Total() (int, error)
}

// MediaStore defines the data layer for Media
type MediaStore struct {
	db          	*sqlx.DB
	imageSizes 		domain.MediaSizes
	convertWebP     bool
	serveWebP     	bool
	compression 	int
	maxWidth    	int
	maxHeight   	int
	maxFileSize 	int
	organiseYear	bool
}

// newMedia - Construct
func newMedia(db *sqlx.DB) *MediaStore {
	ms := &MediaStore{
		db: db,
	}
	ms.init()
	return ms
}

// processImageSizes Processes image sizes from options
func (s *MediaStore) processImageSizes(sizes map[string]interface{}) domain.MediaSizes {
	sizesArr := make(domain.MediaSizes)
	for name, mediaSize := range sizes {
		size := mediaSize.(map[string]interface{})
		sizesArr[name] = domain.MediaSize{
			Name:   size["name"].(string),
			Width:  int(size["width"].(float64)),
			Height: int(size["height"].(float64)),
			Crop:   size["crop"].(bool),
		}
	}
	return sizesArr
}

// init the model with options
func (s *MediaStore) init() {
	om := newOptions(s.db)

	opts, err := om.GetStruct()
	if err != nil {
		log.Fatal(err)
	}

	s.imageSizes = s.processImageSizes(opts.MediaSizes)
	s.convertWebP = opts.MediaConvertWebP
	s.serveWebP = opts.MediaServeWebP
	s.compression = opts.MediaCompression
	s.maxWidth = opts.MediaUploadMaxWidth
	s.maxHeight = opts.MediaUploadMaxHeight
	s.maxFileSize = opts.MediaUploadMaxSize
	s.organiseYear = opts.MediaOrganiseDate
}

// Get all media
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no media available.
func (s *MediaStore) Get(meta http.Params) ([]domain.Media, int, error) {
	const op = "MediaRepository.Get"

	var m []domain.Media
	q := fmt.Sprintf("SELECT * FROM media")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM media")

	// Apply filters to total and original query
	filter, err := filterRows(s.db, meta.Filters, "media")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply pagination
	q += fmt.Sprintf(" ORDER BY media.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, (meta.Page - 1) * meta.Limit)

	// Select media
	if err := s.db.Select(&m, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get media", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	if len(m) == 0 {
		return []domain.Media{}, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No media available", Operation: op}
	}

	// Count the total number of media
	var total int
	if err := s.db.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of media items", Operation: op, Err: err}
	}

	// Loop over and process the sizes
	for k, _ := range m {
		publicSizes, err := s.getSizesForPublic(m[k])
		if err != nil {
			return []domain.Media{}, -1, err
		}
		m[k].Sizes = publicSizes
	}

	return m, total, nil
}

// GetById returns a media item by Id
// Returns errors.NOTFOUND if the media item was not found by the given Id.
func (s *MediaStore) GetById(id int) (domain.Media, error) {
	const op = "MediaRepository.GetById"
	var m domain.Media
	if err := s.db.Get(&m, "SELECT * FROM media WHERE id = ?", id); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the ID: %d", id), Operation: op}
	}
	return m, nil
}

// Gets a media file by the file name
// Returns errors.NOTFOUND if the media item was not found by the given name.
func (s *MediaStore) GetByName(name string) (domain.Media, error) {
	const op = "MediaRepository.GetByName"
	var m domain.Media
	if err := s.db.Get(&m, "SELECT * FROM media WHERE name = ?", name); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the name: %s", name), Operation: op}
	}
	return m, nil
}

// GetByUrl Obtains a media file by the URL from the database
// Returns errors.NOTFOUND if the media item was not found by the given url.
func (s *MediaStore) GetByUrl(url string) (string, string, error) {
	const op = "MediaRepository.GetByUrl"
	var m domain.Media

	// Test normal size
	if err := s.db.Get(&m, "SELECT * FROM media WHERE url = ?", url); err == nil {
		return m.FilePath + "/" + m.UUID.String(), m.Type, nil
	}

	// Test Sizes
	if err := s.db.Get(&m, "SELECT * FROM media WHERE sizes LIKE '%" + url + "%' LIMIT 1"); err == nil {
		sizes, err := s.unmarshalSizes(m)
		if err == nil {
			for _, v := range sizes {
				if v.Url == url {
					return v.FilePath + "/" + v.UUID.String(), m.Type, nil
				}
			}
		}
	}

	return "", "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the media item with the url: %s", url), Operation: op}
}

// Serve is responsible for serving the correct data to the front end
// Returns errors.NOTFOUND if the media item was not found.
func (s *MediaStore) Serve(uploadPath string, acceptWebP bool) ([]byte, string, error) {
	const op = "MediaRepository.Serve"

	path, mimeType, err := s.GetByUrl(uploadPath)
	if err != nil {
		return nil, "", err
	}

	extension := files.GetFileExtension(uploadPath)

	var data []byte
	var found error
	if acceptWebP && s.serveWebP {
		data, found = ioutil.ReadFile(path + extension + ".webp")
		if found != nil {
			data, found = ioutil.ReadFile(path + extension)
		} else {
			mimeType = "image/webp"
		}
	} else {
		data, found = ioutil.ReadFile(path + extension)
	}

	if found != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("File does not exist with the path: %v", uploadPath), Operation: op}
	}

	return data, mimeType, nil
}

// Upload the media files, return bool is for server or user error
// Returns errors.INTERNAL if the uploaded file failed to save.
func (s *MediaStore) Upload(file *multipart.FileHeader, userId int) (domain.Media, error) {
	const op = "MediaRepository.Upload"

	// E.G  /Users/admin/cms/storage/uploads
	path := s.createDirectory()

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
	if err := files.Save(file, path + "/" + key.String() + extension); err != nil {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "Could not save the media file, please try again", Operation: op}
	}

	// Convert to WebP
	if s.convertWebP {
		//decodedImage, err := s.decodeImage(file, mimeType)
		//if err != nil {
			//log.WithFields(log.Fields{"error": err}).Error()
		//}
		//go convertWebP(*decodedImage, path + "/" + key + extension, s.compression)
	}

	// Resize
	sizes := s.saveResizedImages(file, cleanName, path, mimeType, extension)

	// Insert into the database
	dm, err := s.insert(key, cleanName + extension, path, int(file.Size), mimeType, sizes, userId)
	if err != nil {
		return domain.Media{}, err
	}

	publicSizes, err := s.getSizesForPublic(dm)
	if err != nil {
		return domain.Media{}, err
	}

	dm.Sizes = publicSizes

	return dm, nil
}

// Validate the file before uploading
// Returns errors.INVALID if the file was not in the whitelist or
// the file was too big.
func (s *MediaStore) Validate(file *multipart.FileHeader) error {
	const op = "MediaRepository.Validate"

	mimeType, err := mime.TypeByFile(file)
	if err != nil {
		return err
	}

	valid := mime.IsValidMime(config.Media.AllowedFileTypes, mimeType)
	if !valid {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The %s, is not in the whitelist for uploading, please upload a correct file format", mimeType), Operation: op, Err: err}
	}

	fileSize := int(file.Size / 1024)
	if fileSize > s.maxFileSize {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The file exceeds the upload restriction of: %v", s.maxFileSize), Operation: op, Err: err}
	}

	return nil
}

// Inserts a media item into the database
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *MediaStore) insert(uuid uuid.UUID, name string, filePath string, fileSize int, mime string, sizes []domain.MediaSizeDB, userId int) (domain.Media, error) {
	const op = "MediaRepository.insert"

	marshal, err := json.Marshal(sizes)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not marshal the media sizes", Operation: op, Err: err}
	}
	marshalledSizes := json.RawMessage(marshal)

	m := domain.Media{
		UUID: 			uuid,
		Url:			s.getUrl() + "/" + name,
		Title: 			nil,
		Description:	nil,
		Alt: 			nil,
		FilePath:    	filePath,
		FileSize:    	fileSize,
		FileName:    	name,
		Sizes:       	&marshalledSizes,
		Type:        	mime,
		UserID:      	userId,
	}

	q := "INSERT INTO media (uuid, url, title, alt, description, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.db.Exec(q, m.UUID, m.Url, m.Title, m.Alt, m.Description, m.FilePath, m.FileSize, m.FileName, m.Sizes, m.Type, m.UserID)
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

	_, err := s.GetById(m.Id)
	if err != nil {
		return err
	}

	q := "UPDATE media SET title = ?, alt = ?, description = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, m.Title, m.Alt, m.Description, m.Id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the media item with the ID: %v", m.Id), Operation: op, Err: err}
	}

	return nil
}

// Delete the record from the database and all files
// Returns errors.NOTFOUND if the media item was not found.
// Returns errors.INTERNAL if any file (original, webp or any sizes) were not deleted.
// Or if the SQL query was invalid
func (s *MediaStore) Delete(id int) error {
	const op = "MediaRepository.Delete"

	m, err := s.GetById(id)
	if err != nil {
		return err
	}

	extension := files.GetFileExtension(m.Url)

	// Delete the main file
	if err := files.CheckAndDelete(m.FilePath + "/" + m.UUID.String() + extension); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the original media file with the ID: %v", id), Operation: op, Err: err}
	}

	// Delete the sizes and webp versions if stored
	sizes, err := s.unmarshalSizes(m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the media file sizes with the ID: %v", id), Operation: op, Err: err}
	}

	for _, v := range sizes {
		filePath := v.FilePath + "/" + v.UUID.String() + extension
		if err := files.CheckAndDelete(filePath); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the media size file with the name: %v", v.Name), Operation: op, Err: err}
		}
		if err := files.CheckAndDelete(filePath + ".webp"); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the media size webp file with the name: %v", v.Name), Operation: op, Err: err}
		}
	}

	// Delete entry from database
	if _, err := s.db.Exec("DELETE FROM media WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete media item with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Exists Checks if a media items exists by the given name
func (s *MediaStore) Exists(name string) bool {
	const op = "MediaRepository.Exists"
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM media WHERE file_name = ?)", name).Scan(&exists)
	return exists
}

// Total gets the total number of media items
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *MediaStore) Total() (int, error) {
	const op = "MediaRepository.Total"
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM media").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of media items", Operation: op, Err: err}
	}
	return total, nil
}

// saveResizedImages saves all of the resized images and returns
// an array of media DB sizes if successful.
func (s *MediaStore) saveResizedImages(file *multipart.FileHeader, name string, path string, mime string, extension string) []domain.MediaSizeDB {
	const op = "MediaRepository.saveResizedImages"

	var savedSizes []domain.MediaSizeDB
	if mime == "image/png" || mime == "image/jpeg" {
		for _, size := range s.imageSizes {
			mediaUUID := uuid.New()
			fileName := name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + extension

			if err := s.processImageSize(file, path + "/" + mediaUUID.String(), mime, size); err == nil {
				savedSizes = append(savedSizes, domain.MediaSizeDB{
					FilePath: path,
					MediaSize: domain.MediaSize{
						UUID: mediaUUID,
						Url: s.getUrl() + "/" + fileName,
						Name: fileName,
						SizeName: size.Name,
						FileSize: files.GetFileSize(path + "/" + mediaUUID.String() + extension),
						Width: size.Width,
						Height: size.Height,
						Crop: size.Crop,
					},
				})
			}
		}
	}

	return savedSizes
}

// processImageSize processes image sizes, convert WebPs and saves various image sizes based on configuration
// Returns errors.INTERNAL if the image was unable to be saved or decoded.
func (s *MediaStore) processImageSize(file *multipart.FileHeader, filePath string, mime string, size domain.MediaSize) error {
	const op = "MediaRepository.processImageSize"

	// PNG Type
	if mime == "image/png" {
		filePath = filePath + ".png"

		decodedImage, err  := s.decodeImage(file, mime)
		if err != nil {
			return err
		}
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.PNGCompressionLevel(png.CompressionLevel(s.compression))); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not save the resized image"), Operation: op, Err: err}
		}

		if s.convertWebP {
			//go convertWebP(resized, filePath, s.compression)
		}
	}

	// Jpg Type
	if mime == "image/jpeg" || mime == "image/jp2" {
		filePath = filePath + ".jpg"

		decodedImage, err := s.decodeImage(file, mime)
		if err != nil {
			return err
		}
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.JPEGQuality(s.compression)); err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not save the resized image"), Operation: op, Err: err}
		}

		if s.convertWebP {
			//go convertWebP(resized, filePath, s.compression)
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

// Converts an image to webp based on compression and decoded image.
// Compression level is also set.
func convertWebP(image image.Image, path string, compression int) {
	const op = "MediaRepository.insert"

	var buf bytes.Buffer
	var opts = webp.Options{
		Lossless: true,
		Quality:  float32(compression),
	}

	if err := webp.Encode(&buf, image, &opts); err != nil {
		log.Error(err)
	}

	if err := ioutil.WriteFile(path + ".webp", buf.Bytes(), 0666); err != nil {
		log.Error(err)
	}

	log.Info("WebP conversion ok with path: " + path + ".webp")
}

// createDirectory creates the media directory year path if the organise year variable in the media
// store is set to true. Date and year folders are created recursively.
func (s *MediaStore) createDirectory() string {
	const op = "MediaRepository.createDirectory"

	uploadsPath := paths.Uploads()

	if !s.organiseYear {
		return uploadsPath
	} else {
		t := time.Now()
		path := uploadsPath + "/" + t.Format("2006") + "/" + t.Format("01")

		if _, err := os.Stat(path); os.IsNotExist(err) {
			_ = os.MkdirAll(path, os.ModePerm)
		}
		return path
	}
}

// Get the public url of the file according to date and month if the organise
// year variable in the media store is set to true. If not the function will
// return the public uploads folder by default.
func (s *MediaStore) getUrl() string {
	if !s.organiseYear {
		return paths.PublicUploads()
	} else {
		t := time.Now()
		return paths.PublicUploads() + "/" + t.Format("2006") + "/" + t.Format("01")
	}
}

// decodeImage decodes the image from a file dependant on the mime type
// Returns errors.INTERNAL if the file was unable to be decoded or the file
// was unable to be opened.
func (s *MediaStore) decodeImage(file *multipart.FileHeader, mime string) (*image.Image, error) {
	const op = "MediaRepository.decodeImage"

	reader, err := file.Open()
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to open the file with the filename: %s", file.Filename), Operation: op, Err: err}
	}

	if mime == "image/png" {
		pngFile, err := png.Decode(reader)
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not decode the image with the filename: %s", file.Filename), Operation: op, Err: err}
		}
		return &pngFile, nil

	} else if mime == "image/jpeg" || mime == "image/jp2" {
		jpgFile, err := jpeg.Decode(reader)
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not decode the image with the filename: %s", file.Filename), Operation: op, Err: err}
		}
		return &jpgFile, nil
	}

	return nil, &errors.Error{Code: errors.INTERNAL, Message: "Something went wrong decoding the image", Operation: op, Err: err}
}

// Unmarshal the media sizes for processing
// Returns errors.INTERNAL if the size was unable to be unmarshalled
func (s *MediaStore) unmarshalSizes(m domain.Media) ([]domain.MediaSizeDB, error) {
	const op = "MediaRepository.unmarshalSizes"
	var sizes []domain.MediaSizeDB
	if err := json.Unmarshal(*m.Sizes, &sizes); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could unmarshal the image sizes", Operation: op, Err: err}
	}
	return sizes, nil
}

// Unmarshal the media sizes for public use
// Returns errors.INTERNAL if the size was unable to be marshalled
// TODO: Sort by name or size
func (s *MediaStore) getSizesForPublic(m domain.Media) (*json.RawMessage, error) {
	const op = "MediaRepository.getSizesForPublic"

	ms, err := s.unmarshalSizes(m)
	if err != nil {
		return nil, err
	}

	var returnData []domain.MediaSize
	for _, v := range ms {
		returnData = append(returnData, domain.MediaSize{
			UUID: 	  v.UUID,
			Url:      v.Url,
			Name:     v.Name,
			FileSize: v.FileSize,
			Width:    v.Width,
			Height:   v.Height,
			Crop:     v.Crop,
		})
	}

	marshalled, err := json.Marshal(returnData)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could marshal the image sizes", Operation: op, Err: err}
	}
	jsonMessage := json.RawMessage(marshalled)

	return &jsonMessage, nil
}

// Process file name
func (s *MediaStore) processFileName(file string, extension string) string {

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


