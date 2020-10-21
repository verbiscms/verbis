package res

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
	"io/ioutil"
)

// Converts an image to webp based on compression and decoded image.
// Compression level is also set.
//func convertWebP(image image.Image, path string, compression int) {
//	const op = "MediaRepository.convertWebP"
//
//	var buf bytes.Buffer
//	var opts = webp.Options{
//		Lossless: true,
//		Quality:  float32(compression),
//	}
//
//	if err := webp.Encode(&buf, image, &opts); err != nil {
//		log.Error(err)
//	}
//
//	if err := ioutil.WriteFile(path + ".webp", buf.Bytes(), 0666); err != nil {
//		log.Error(err)
//	}
//
//	log.Info("WebP conversion ok with path: " + path + ".webp")
//}
//
//// Converts an image to webp based on compression and decoded image.
//// Compression level is also set.
//func convertWebP(image image.Image, path string, compression int) {
//	const op = "MediaRepository.convertWebP"
//
//	webpConfig, _ := webp.ConfigPreset(webp.PresetDefault, float32(compression))
//
//	// Create file and buffered writer
//	io, err := os.Create(path + ".webp")
//	if err != nil {
//		return
//	}
//	w := bufio.NewWriter(io)
//	defer func() {
//		w.Flush()
//		io.Close()
//	}()
//
//	if err := webp.EncodeRGBA(w, image, webpConfig); err != nil {
//		return
//	}
//
//	log.Info("WebP conversion ok with path: " + path + ".webp")
//}


// Converts an image to webp based on compression and decoded image.
// Compression level is also set.
//func convertWebP(image image.Image, path string, compression int) {
//	const op = "MediaRepository.convertWebP"
//
//	var buf bytes.Buffer
//
//	if err := webp.Encode(&buf, image, &webp.Options{Lossless: true}); err != nil {
//		log.Println(err)
//	}
//
//	if err := ioutil.WriteFile(path + ".webp", buf.Bytes(), 0666); err != nil {
//		log.Println(err)
//	}
//}