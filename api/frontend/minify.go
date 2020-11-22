package frontend

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	min "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"os"
	"regexp"
)

// Minifier represents functions for executing the minify package.
type Minifier interface {
	Minify(f *os.File, mime string) ([]byte, error)
	MinifyBytes(b *bytes.Buffer, mime string) ([]byte, error)
}

// Minify represents the minify type along with the minify package
// and options to determine whether or not to minify the asset.
type Minify struct {
	pkg     *min.M
	options models.OptionsRepository
}

// New - Construct, sets minify functions
func New(o models.OptionsRepository) *Minify {
	m := min.New()

	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	return &Minify{
		pkg:     m,
		options: o,
	}
}

// Minify minifies a file & calls the compare function to render
// the file.
func (m *Minify) Minify(f *os.File, mime string) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	return m.compare(&buf, mime)
}

// MinifyBytes minifies existing bytes.Buffer & calls the compare
// function to render the file.
// Usually used for HTML files.
func (m *Minify) MinifyBytes(b *bytes.Buffer, mime string) ([]byte, error) {
	return m.compare(b, mime)
}

// compare gets the options struct in order to see if the user has
// selected the type of minification.
// It then compares mime's and executes the file to be minified.
func (m *Minify) compare(b *bytes.Buffer, mime string) ([]byte, error) {
	const op = "Minify.Compare"

	options, optsErr := m.options.GetStruct()
	if optsErr != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: fmt.Errorf("could not get the options struct")},
		}).Fatal()
	}

	var (
		render []byte
		err    error
	)

	switch mime {
	case "text/html":
		{
			render, err = m.execute(b, options.MinifyHTML, mime)
			if err != nil {
				return nil, err
			}
			break
		}
	case "text/css":
		{
			render, err = m.execute(b, options.MinifyCSS, mime)
			if err != nil {
				return nil, err
			}
			return render, nil
		}
	case "application/javascript":
		{
			render, err = m.execute(b, options.MinifyJS, mime)
			if err != nil {
				return nil, err
			}
			break
		}
	case "image/svg+xml":
		{
			render, err = m.execute(b, options.MinifySVG, mime)
			if err != nil {
				return nil, err
			}
			break
		}
	case "application/json":
		{
			render, err = m.execute(b, options.MinifyJSON, mime)
			if err != nil {
				return nil, err
			}
			break
		}
	case "text/xml":
		{
			render, err = m.execute(b, options.MinifyXML, mime)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return render, nil
}

// execute the buffer.Bytes depending onn the user selection in the options
// table.
// Returns the original bytes if the minification failed.
// Returns errors.INTERNAL if something went wrong minifying the file.
func (m *Minify) execute(buf *bytes.Buffer, allow bool, mime string) ([]byte, error) {
	const op = "Minifier.execute"

	var (
		render []byte
		err    error
	)

	if allow {
		render, err = m.pkg.Bytes(mime, buf.Bytes())
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not minify the file"), Operation: op, Err: err}
		}
	} else {
		return buf.Bytes(), nil
	}

	return render, nil
}
