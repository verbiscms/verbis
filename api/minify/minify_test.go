package minify

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMinify_MinifyBytes(t *testing.T) {

	m := New(defaultConfig)

	tt := map[string]struct {
		input string
		mime  string
		want  interface{}
	}{
		"HTML": {
			` <div> <i> test </i> <b> test </b> </div> `,
			HTML,
			`<div><i>test</i> <b>test</b></div>`,
		},
		"CSS": {
			`/*comment*/`,
			CSS,
			``,
		},
		"JS": {
			`/*comment*/a`,
			Javascript,
			`a`,
		},
		"SVG": {
			`<!-- comment -->`,
			SVG,
			``,
		},
		"JSON": {
			`{ "a": [1, 2] }`,
			JSON,
			`{"a":[1,2]}`,
		},
		"XML": {
			`<!-- comment -->`,
			XML,
			``,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte(test.input))
			got, _ := m.MinifyBytes(b, test.mime)
			assert.Equal(t, test.want, string(got))
		})
	}
}

func TestMinify_MinifyBytesError(t *testing.T) {
	m := New(defaultConfig)

	t.Run("Error", func(t *testing.T) {
		orignal := htmlMime
		defer func() {
			htmlMime = orignal
		}()

		htmlMime = "wrongval"
		b := bytes.NewBuffer([]byte(""))
		_, err := m.MinifyBytes(b, HTML)
		assert.Contains(t, err.Error(), "Minifier.execute: minifier does not exist for mimetype")
	})
}

func TestMinify_MinifyBytesNotParsed(t *testing.T) {

	m := New(Config{
		MinifyHTML: false,
		MinifyCSS:  false,
		MinifyJS:   false,
		MinifySVG:  false,
		MinifyJSON: false,
		MinifyXML:  false,
	})

	tt := map[string]struct {
		input string
		mime  string
		want  interface{}
	}{
		"HTML": {
			` <div> <i> test </i> <b> test </b> </div> `,
			HTML,
			` <div> <i> test </i> <b> test </b> </div> `,
		},
		"CSS": {
			`/*comment*/`,
			CSS,
			`/*comment*/`,
		},
		"JS": {
			`/*comment*/a`,
			Javascript,
			`/*comment*/a`,
		},
		"SVG": {
			`<!-- comment -->`,
			SVG,
			`<!-- comment -->`,
		},
		"JSON": {
			`{ "a": [1, 2] }`,
			JSON,
			`{ "a": [1, 2] }`,
		},
		"XML": {
			`<!-- comment -->`,
			XML,
			`<!-- comment -->`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte(test.input))
			got, _ := m.MinifyBytes(b, test.mime)
			assert.Equal(t, test.want, string(got))
		})
	}
}

func TestMinify_Minify(t *testing.T) {
	m := New(defaultConfig)

	file, err := os.Create(os.TempDir() + "verbis-test-minify")
	assert.NoError(t, err)

	defer os.Remove(file.Name())

	_, err = file.Write([]byte(` <div> <i> test </i> <b> test </b> </div> `))
	assert.NoError(t, err)

	got, err := m.Minify(file.Name(), HTML)
	assert.NoError(t, err)

	assert.Equal(t, []byte(`<div><i>test</i> <b>test</b></div>`), got)

	err = file.Close()
	assert.NoError(t, err)
}

func TestMinify_MinifyError(t *testing.T) {
	m := New(defaultConfig)
	_, err := m.Minify("wrong", HTML)
	assert.Contains(t, err.Error(), "Minifier.Minify: open wrong: no such file or directory")
}
