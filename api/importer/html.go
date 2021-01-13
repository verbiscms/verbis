package importer

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

func ParseHTML(content string, downloader func(url string) (string, error)) (string, error) {

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return "", err
	}

	var f func(*html.Node) bool
	f = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "img" {
			for index, img := range n.Attr {
				if img.Key == "src" {
					newUrl, err := downloader(img.Val)
					if err != nil {
						break
					}
					n.Attr[index].Val = newUrl
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
		return "", err
	}

	replacer := strings.NewReplacer("<html><head></head><body>", "", "</body></html>", "")

	return replacer.Replace(b.String()), nil
}
