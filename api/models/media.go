package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
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

type MediaRepository interface {
	GetAll(meta http.Params) ([]domain.Media, error)
	GetById(id int) (domain.Media, error)
	GetByName(name string) (domain.Media, error)
	GetByUrl(url string) (domain.Media, error)
	Serve(uploadPath string, acceptWeb bool) ([]byte, string, error)
	Upload(file *multipart.FileHeader, userId int) (domain.Media, error)
	Validate(file *multipart.FileHeader) error
	Update(m *domain.Media) error
	Delete(id int) error
	Exists(name string) bool
	Total() (int, error)
}

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

//Construct
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

// init the model
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

// GetAll Gets all media files from the database
func (s *MediaStore) GetAll(meta http.Params) ([]domain.Media, error) {
	var m []domain.Media
	q := fmt.Sprintf("SELECT * FROM media ORDER BY media.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, meta.Page * meta.Limit)
	if err := s.db.Select(&m, q); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not get media - %w", err)
	}

	if len(m) == 0 {
		return []domain.Media{}, nil
	}

	for k, _ := range m {
		publicSizes, err := s.getSizesForPublic(m[k])
		if err != nil {
			log.Error(err)
		}
		m[k].Sizes = publicSizes
	}

	return m, nil
}

// GetBydId Gets all media files from the database
func (s *MediaStore) GetById(id int) (domain.Media, error) {
	var m domain.Media
	if err := s.db.Get(&m, "SELECT * FROM media WHERE id = ?", id); err != nil {
		log.Info(err)
		return domain.Media{}, fmt.Errorf("Could not get media with the ID: %v", id)
	}
	return m, nil
}

// Gets a media file by the file name
func (s *MediaStore) GetByName(name string) (domain.Media, error) {
	var m domain.Media
	if err := s.db.Get(&m, "SELECT * FROM media WHERE name = ?", name); err != nil {
		log.Info(err)
		return domain.Media{}, fmt.Errorf("Could not get media with the name: %v", name)
	}
	return m, nil
}

// GetByUrl Obtains a media file by the URL from the database
func (s *MediaStore) GetByUrl(url string) (domain.Media, error) {
	var m domain.Media
	if err := s.db.Get(&m, "SELECT * FROM media WHERE url = ?", url); err != nil {
		log.Info(err)
		return domain.Media{}, fmt.Errorf("Could not get media with the url: %v", url)
	}
	return m, nil
}

// Serve is responsible for serving the correct data to the front end
func (s *MediaStore) Serve(uploadPath string, acceptWebP bool) ([]byte, string, error) {
	m, err := s.GetByUrl(uploadPath)
	if err != nil {
		return nil, "", err
	}

	extension := files.GetFileExtension(uploadPath)

	var mimeType = m.Type
	var data []byte
	var found error
	if acceptWebP && s.serveWebP {
		data, found = ioutil.ReadFile(m.FilePath + "/" + m.UUID.String() + extension + ".bollox")
		if found != nil {
			data, found = ioutil.ReadFile(m.FilePath + "/" + m.UUID.String() + extension)
		} else {
			mimeType = "image/webp"
		}
	} else {
		data, found = ioutil.ReadFile(m.FilePath + "/" + m.UUID.String() + extension)
	}

	if found != nil {
		return nil, "", fmt.Errorf("File does not exist with the path: %v", uploadPath)
	}

	return data, mimeType, nil
}

// Upload the media files, return bool is for server or user error
func (s *MediaStore) Upload(file *multipart.FileHeader, userId int) (domain.Media, error) {

	// E.G  /Users/admin/cms/storage/uploads
	path := s.createDirectory()

	// E.G: Image20@.png
	name := file.Filename

	// E.G: .png
	extension := files.GetFileExtension(name)

	// E.G: image.png
	cleanName := s.processFileName(name, extension)

	// E.G: 180ea4324ab2556032141e956ca1f141
	key, err := encryption.GenerateRandomHash()
	if err != nil {
		return domain.Media{}, err
	}

	// E.G image/png
	mimeType, _ := mime.TypeByFile(file)

	// Save the uploaded file
	if err := files.Save(file, path + "/" + key + extension); err != nil {
		log.Error(err)
		return domain.Media{}, fmt.Errorf("Could not save the media file, please try again")
	}

	// Convert to WebP
	if s.convertWebP {
		decodedImage := s.decodeImage(file, mimeType)
		go convertWebP(*decodedImage, path + "/" + key + extension, s.compression)
	}

	// Resize
	sizes := s.saveResizedImages(file, cleanName, path, mimeType, extension)

	// Insert into the database
	dm, err := s.insert(key, cleanName + extension, path, int(file.Size), mimeType, sizes, userId)
	if err != nil {
		log.Info(err)
		return domain.Media{}, nil
	}

	publicSizes, err := s.getSizesForPublic(dm)
	if err != nil {
		log.Info(err)
		return domain.Media{}, nil
	}

	dm.Sizes = publicSizes

	return dm, nil
}

// Validate the file before uploading
func (s *MediaStore) Validate(file *multipart.FileHeader) error {

	// Get the mime type of the file
	mimeType, err := mime.TypeByFile(file)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Unable to get the mime type of the file")
	}

	// Check if the mime is valid in the allowed file types configuration array
	valid := mime.IsValidMime(config.Media.AllowedFileTypes, mimeType)
	if !valid {
		err := fmt.Errorf("The %s, is not in the whitelist for uploading, please upload a correct file format", mimeType)
		log.Info(err)
		return err
	}

	// Check for max file upload limit
	fileSize := int(file.Size / 1024)
	if fileSize > s.maxFileSize {
		return fmt.Errorf("The file exceeds the upload restriction of: %v", s.maxFileSize)
	}

	return nil
}

// Inserts a media item into the database
func (s *MediaStore) insert(key string, name string, filePath string, fileSize int, mime string, sizes []domain.MediaSizeDB, userId int) (domain.Media, error) {

	marshal, _ := json.Marshal(sizes)
	marshalledSizes := json.RawMessage(marshal)

	m := domain.Media{
		UUID: 			uuid.New(),
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
		log.Error(err)
		return domain.Media{}, fmt.Errorf("Could not create the media item with the name: %v", m.FileName)
	}

	id, err := c.LastInsertId()
	if err != nil {
		log.Error(err)
		return domain.Media{}, fmt.Errorf("Could not get the newly created media with the name: %v", m.FileName)
	}
	m.Id = int(id)

	return m, nil
}

// Update the media item (title, alt & description)
func (s *MediaStore) Update(m *domain.Media) error {
	_, err := s.GetById(m.Id)
	if err != nil {
		log.Info(err)
		return err
	}

	q := "UPDATE media SET title = ?, alt = ?, description = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, m.Title, m.Alt, m.Description, m.Id)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update the media item with the ID: %v", m.Id)
	}

	return nil
}

// Delete the record from the database and all files
func (s *MediaStore) Delete(id int) error {
	m, err := s.GetById(id)
	if err != nil {
		return err
	}

	extension := files.GetFileExtension(m.Url)

	// Delete the main file
	if err := files.CheckAndDelete(m.FilePath + "/" + m.UUID.String() + extension); err != nil {
		fmt.Println(m.FilePath + "/" + m.FileName)
		log.Error(err)
		return fmt.Errorf("Could not delete the original media file with the ID: %v", id)
	}

	// Delete the sizes and webp versions if stored
	sizes, err := s.unmarshalSizes(m)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete the media file sizes with the ID: %v", id)
	}
	for _, v := range sizes {
		filePath := v.FilePath + "/" + v.UUID.String() + extension
		if err := files.CheckAndDelete(filePath); err != nil {
			log.Error(err)
			return fmt.Errorf("Could not delete the media size file with the name: %v", v.Name)
		}
		if err := files.CheckAndDelete(filePath + ".webp"); err != nil {
			log.Error(err)
			return fmt.Errorf("Could not delete the media size webp file with the name: %v", v.Name)
		}
	}

	// Delete entry from database
	if _, err := s.db.Exec("DELETE FROM media WHERE id = ?", id); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not delete media with the ID : %v", id)
	}

	return nil
}

// Check if the media item exists in the database
func (s *MediaStore) Exists(name string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM media WHERE file_name = ?)", name).Scan(&exists)
	return exists
}

// Get the total number of posts
func (s *MediaStore) Total() (int, error) {
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM media").Scan(&total); err != nil {
		log.Error(err)
		return -1, fmt.Errorf("Could not get the total number of media items")
	}
	return total, nil
}

// Saves all of the resized images and returns an array if successful
func (s *MediaStore) saveResizedImages(file *multipart.FileHeader, name string, path string, mime string, extension string) []domain.MediaSizeDB {
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

// Process image sizes, convert WebPs and saves various image sizes based on configuration
func (s *MediaStore) processImageSize(file *multipart.FileHeader, filePath string, mime string, size domain.MediaSize) error {

	// PNG Type
	if mime == "image/png" {
		filePath = filePath + ".png"

		decodedImage := s.decodeImage(file, mime)
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.PNGCompressionLevel(png.CompressionLevel(s.compression)));
			err != nil {
			log.Error(err)
			return err
		}
		log.Info("Saved image resized image with path: " + filePath)

		if s.convertWebP {
			go convertWebP(resized, filePath, s.compression)
		}
	}

	// Jpg Type
	if mime == "image/jpeg" || mime == "image/jp2" {
		filePath = filePath + ".jpg"

		decodedImage := s.decodeImage(file, mime)
		resized := resizeImage(*decodedImage, size.Width, size.Height, size.Crop)

		if err := imaging.Save(resized, filePath, imaging.JPEGQuality(s.compression)); err != nil {
			log.Error(err)
			return err
		}
		log.Info("Saved image resized image with path: " + filePath)

		if s.convertWebP {
			go convertWebP(resized, filePath, s.compression)
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

// Creates the media directory year path if the organise year variable in the media
// store is set to true. Date and year folders are created recursively.
func (s *MediaStore) createDirectory() string {
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

// Decode the image from a file dependant on the mim type
func (s *MediaStore) decodeImage(file *multipart.FileHeader, mime string) *image.Image {
	reader, err := file.Open()
	if err != nil {
		log.Error(err)
	}

	if mime == "image/png" {
		pngFile, err := png.Decode(reader)
		if err != nil {
			log.Error(err)
		}
		return &pngFile

	} else if mime == "image/jpeg" || mime == "image/jp2" {
		jpgFile, err := jpeg.Decode(reader)
		if err != nil {
			log.Error(err)
		}
		return &jpgFile
	}

	return nil
}

// Unmarshal the media sizes for processing
func (s *MediaStore) unmarshalSizes(m domain.Media) ([]domain.MediaSizeDB, error) {
	var sizes []domain.MediaSizeDB
	if err := json.Unmarshal(*m.Sizes, &sizes); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not unmarshal the media sizes with the ID: %v", m.Id)
	}
	return sizes, nil
}

// Unmarshal the media sizes for public use
// TODO: Sort by name or size
func (s *MediaStore) getSizesForPublic(m domain.Media) (*json.RawMessage, error) {
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
		return nil, fmt.Errorf("Could not marshal the media sizes with the ID: %v", m.Id)
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


