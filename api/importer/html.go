// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/net/html"
	"mime/multipart"
	"strings"
)

// uploader defines the return return function thats called when a image
// is found crawling the content.
type uploader func(file *multipart.FileHeader, url string, err error) string

// ParseHTML
//
func ParseHTML(content string, upload uploader) (string, error) {
	const op = "Importer.ParseHTML"

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return "", err
	}

	var f func(*html.Node) bool
	f = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "img" {
			for index, img := range n.Attr {
				if img.Key == "src" {
					file, err := DownloadFile(img.Val)
					url := upload(file, img.Val, err)

					if url == "" {
						break
					}

					n.Attr[index].Val = url
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		return false
	}

	f(doc)

	var b bytes.Buffer
	err = html.Render(&b, doc)
	if err != nil {
		return "", &errors.Error{Code: errors.INVALID, Message: "Could not parse HTML", Operation: op, Err: fmt.Errorf("unable to parse HTML")}
	}

	replacer := strings.NewReplacer("<html><head></head><body>", "", "</body></html>", "")

	return replacer.Replace(b.String()), nil
}
