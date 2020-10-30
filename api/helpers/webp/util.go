package webp

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image"
	"io/ioutil"
	"strings"
)

// Accepts checks to see if the browser accepts WebP images
func Accepts(g *gin.Context) bool {
	acceptHeader := g.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, "image/webp")
}

// GetData first checks to see if the browser accepts WebP images
// and if the mime type is jpg or a png.
// Returns a data was found, nil if it hasn't.
func GetData(g *gin.Context, path string, mime string) []byte {
	if Accepts(g) && mime == "image/jpeg" || mime == "image/png" {
		data, found := ioutil.ReadFile(path + ".webp")
		if found != nil {
			return nil
		}
		return data
	}
	return nil
}

// Converts an image to webp based on compression and decoded image.
// Compression level is also set.
func Convert(image image.Image, path string, compression int) {
	var buf bytes.Buffer
	if err := webp.Encode(&buf, image, &webp.Options{Lossless: true}); err != nil {
		log.Error(err)
	}

	if err := ioutil.WriteFile(path + ".webp", buf.Bytes(), 0666); err != nil {
		log.Error(err)
	}
}