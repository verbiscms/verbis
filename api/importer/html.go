package importer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/net/html"
	"mime/multipart"
	"strings"
)

type uploader func(file *multipart.FileHeader, url string, err error) (string, error)

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
					url, err := upload(file, img.Val, err)
					if err != nil {
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
