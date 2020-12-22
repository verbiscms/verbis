package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestSuite - Sets up up a TemplateFunctions with gin read
// for testing.
func newTestSuite(args ...string) *TemplateFunctions {
	gin.SetMode(gin.TestMode)
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get", nil)

	p := &domain.PostData{}
	if len(args) == 1 {
		data := []byte(args[0])
		var m map[string]interface{}
		err := json.Unmarshal(data, &m)
		if err != nil {
			fmt.Println(err)
		}
		p = &domain.PostData{
			Post: domain.Post{
				Fields: m,
			},
		}
	}

	mockOptions := mocks.OptionsRepository{}
	mockOptions.On("GetStruct").Return(domain.Options{}, nil)

	mockSite := mocks.SiteRepository{}
	mockSite.On("GetGlobalConfig").Return(&domain.Site{}, nil)
	mockSite.On("GetThemeConfig").Return(domain.ThemeConfig{})

	return NewFunctions(g, &models.Store{
		Options: &mockOptions,
		Site:    &mockSite,
	}, p, config.Configuration{})
}

// runtv
//
// Run the template test by executing the tpl given
// with data.
func runtv(t *testing.T, tf *TemplateFunctions, tpl string, expect interface{}, data interface{}) {
	tt := template.Must(template.New("test").Funcs(tf.GetFunctions()).Parse(tpl))

	var b bytes.Buffer
	err := tt.Execute(&b, data)

	if err != nil {
		fmt.Println(err)
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	assert.Equal(t, got, b.String())
}

// runt
//
//Run the template test by executing the tpl given
func runt(t *testing.T, tf *TemplateFunctions, tpl string, expect interface{}) {
	tt := template.Must(template.New("test").Funcs(tf.GetFunctions()).Parse(tpl))

	var b bytes.Buffer
	err := tt.Execute(&b, map[string]string{})

	if err != nil {
		fmt.Println(err)
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	assert.Equal(t, got, b.String())
}

// CLEAN UP

func rune(t *testing.T, tf *TemplateFunctions, tpl string, expect interface{}, e error) {
	tt := template.Must(template.New("test").Funcs(tf.GetFunctions()).Parse(tpl))

	var b bytes.Buffer
	err := tt.Execute(&b, map[string]string{})
	assert.Contains(t, err.Error(), e.Error())

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")
	assert.Equal(t, got, b.String())
}
