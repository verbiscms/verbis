// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package minify

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/errors"
	min "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"io/ioutil"
	"regexp"
)

// Minifier represents functions for executing the minify package.
type Minifier interface {
	Minify(name string, mime string) ([]byte, error)
	MinifyBytes(b *bytes.Buffer, mime string) ([]byte, error)
}

// Minify represents the minify type along with the minify
// package and options to determine whether or not to
// minify the asset passed.
type minify struct {
	pkg    *min.M
	config Config
}

// New
//
// Creates a new Minify instance, if no config options are
// passed the defaultConfig is used.
func New(cfg ...Config) Minifier {
	m := min.New()

	m.AddFunc(htmlMime, html.Minify)
	m.AddFunc(cssMime, css.Minify)
	m.AddFuncRegexp(regexp.MustCompile(jsMime), js.Minify)
	m.AddFunc(svgMime, svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile(jsonMime), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile(xmlMime), xml.Minify)

	config := defaultConfig
	if len(cfg) == 1 {
		config = cfg[0]
	}

	return &minify{
		pkg:    m,
		config: config,
	}
}

// Minify
//
// Minifies a file & calls the compare function to render
// the file.
func (m *minify) Minify(name, mime string) ([]byte, error) {
	const op = "Minifier.Minify"

	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to read file contents", Operation: op, Err: err}
	}

	return m.compare(bytes.NewBuffer(b), mime)
}

// MinifyBytes
//
// Minifies existing bytes.Buffer & calls the compare function
// to render the file. Usually used for HTML files.
func (m *minify) MinifyBytes(b *bytes.Buffer, mime string) ([]byte, error) {
	return m.compare(b, mime)
}

// compare
//
// Gets the options struct in order to see if the user has
// selected the type of minification. It then compares
// mime's and executes the file to be minified.
func (m *minify) compare(b *bytes.Buffer, mime string) ([]byte, error) {
	var (
		render []byte
		err    error
	)

	switch mime {
	case HTML:
		render, err = m.execute(b, m.config.MinifyHTML, mime)
	case CSS:
		render, err = m.execute(b, m.config.MinifyCSS, mime)
	case Javascript:
		render, err = m.execute(b, m.config.MinifyJS, mime)
	case SVG:
		render, err = m.execute(b, m.config.MinifySVG, mime)
	case JSON:
		render, err = m.execute(b, m.config.MinifyJSON, mime)
	case XML:
		render, err = m.execute(b, m.config.MinifyXML, mime)
	default:
		return b.Bytes(), nil
	}

	if err != nil {
		return b.Bytes(), err
	}

	return render, nil
}

// execute
//
// Execute the buffer.Bytes depending onn the user selection
// in the Config table.
//
// Returns the original bytes if the minification failed.
// Returns errors.INTERNAL if something went wrong minifying the file.
func (m *minify) execute(buf *bytes.Buffer, allow bool, mime string) ([]byte, error) {
	const op = "Minifier.execute"

	if allow {
		render, err := m.pkg.Bytes(mime, buf.Bytes())
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not minify the file", Operation: op, Err: err}
		}
		return render, nil
	}

	return buf.Bytes(), nil
}
