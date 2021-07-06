// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/spf13/cobra"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			config, _, err := doctor(false)
			if err != nil {
				printError(err.Error())
				return
			}

			//client, err := storage.New(config.Env, &domain.Options{
			//	StorageProvider: domain.StorageAWS,
			//	StorageBucket:   "reddicotest",
			//}, config.Store.Files)

			client, err := storage.New(config.Env, &domain.Options{
				StorageProvider: domain.StorageLocal,
			}, config.Store.Files)

			if err != nil {
				printError(err.Error())
				return
			}
			//
			//text, item, err := client.Find("/test2.txt")
			//if err != nil {
			//	printError(err.Error())
			//	return
			//}
			//color.Green.Printf(string(text))
			//color.Red.Printf("%+v\n", item)

			err = client.Delete(18)
			if err != nil {
				printError(err.Error())
				return
			}

			//contents := "This is a new file stored in the cloud"
			//r := strings.NewReader(contents)
			//
			//upload, err := client.Upload(domain.Upload{
			//	Path:       "test2.txt",
			//	Size:       int64(len(contents)),
			//	Contents:   r,
			//	Private:    false,
			//	SourceType: "media",
			//})
			//if err != nil {
			//	printError(err.Error())
			//}

			//color.Red.Printf("%+v\n", upload)

			//media := media.New(&domain.Options{
			//	MediaCompression:     0,
			//	MediaConvertWebP:     true,
			//	MediaServeWebP:       false,
			//	MediaUploadMaxSize:   0,
			//	MediaUploadMaxWidth:  0,
			//	MediaUploadMaxHeight: 0,
			//	MediaOrganiseDate:    true,
			//	MediaSizes: domain.MediaSizes{
			//		"test": domain.MediaSize{
			//			Width:  300,
			//			Height: 300,
			//			Crop:   false,
			//		},
			//	},
			//}, client, func(fileName string) bool {
			//	return false
			//})
			//
			//file, err := File("/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata/images/gopher.png")
			//if err != nil {
			//	printError(err.Error())
			//	return
			//}
			//
			//_, err = media.Upload(file)
			//if err != nil {
			//	printError(err.Error())
			//	return
			//}

			//

			//fmt.Println(upload)

		},
	}
)

// File converts a file path into a *multipart.FileHeader.
func File(path string) (*multipart.FileHeader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	mr := multipart.NewReader(body, writer.Boundary())
	mt, err := mr.ReadForm(99999)
	if err != nil {
		return nil, err
	}

	ft := mt.File["file"][0]

	return ft, nil
}

func res() {
	//_, _, err = client.Find("test.txt")
	//if err != nil {
	//	printError(err.Error())
	//	return
	//}
	//
	//err = client.Delete("test.txt")
	//if err != nil {
	//	printError(err.Error())
	//	return
	//}
}
