// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/media"
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
			env, err := environment.Load()
			if err != nil {
				printError(err.Error())
				return
			}

			logger.Init(env)

			//client, err := storage.New(env, &domain.Options{
			//	StorageProvider: domain.StorageAWS,
			//	StorageBucket:   "reddicotest",
			//})

			client, err := storage.New(env, &domain.Options{
				StorageProvider: domain.StorageLocal,
			})

			if err != nil {
				printError(err.Error())
				return
			}

			media := media.New(&domain.Options{
				MediaCompression:     0,
				MediaConvertWebP:     true,
				MediaServeWebP:       false,
				MediaUploadMaxSize:   0,
				MediaUploadMaxWidth:  0,
				MediaUploadMaxHeight: 0,
				MediaOrganiseDate:    true,
				MediaSizes:           domain.MediaSizes{
					"test": domain.MediaSize{
						Width:    300,
						Height:   300,
						Crop:     false,
					},
				},
			}, client, func(fileName string) bool {
				return false
			})

			file, err := File("/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata/images/gopher.png")
			if err != nil {
				printError(err.Error())
				return
			}

			_, err = media.Test(file)
			if err != nil {
				printError(err.Error())
				return
			}

			//
			//upload, err := client.Upload("test.txt", strings.NewReader("this is a test"))
			//if err != nil {
			//	fmt.Println("her")
			//	printError(err.Error())
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
