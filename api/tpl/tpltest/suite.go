package tpltest_test

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TplTestSuite struct {
	test *testing.T
	*tpl.TemplateManager
}

func New(t *testing.T) *TplTestSuite {
	_ = logger.Init(config.Configuration{})
	gin.SetMode(gin.TestMode)
	log.SetOutput(ioutil.Discard)

	manager := getManager()
	return &TplTestSuite{
		t,
		manager,
	}
}

// GetManager
//
// Sets up up a TemplateManager with gin read
// for testing.
func getManager() *tpl.TemplateManager {
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get", nil)

	mockOptions := mocks.OptionsRepository{}
	mockOptions.On("GetStruct").Return(domain.Options{}, nil)

	mockSite := mocks.SiteRepository{}
	mockSite.On("GetGlobalConfig").Return(domain.Site{}, nil)
	mockSite.On("GetThemeConfig").Return(domain.ThemeConfig{})

	return tpl.NewManager(g, &models.Store{
		Options: &mockOptions,
		Site:    &mockSite,
	}, &domain.PostData{}, config.Configuration{})
}

// Execute
//
// Executes the templates with functions and returns the resulting
// html or an error if there was a problem executing the template.
func (t *TplTestSuite) Execute(tpl string, data interface{}) (string, error) {
	funcMap := internal.GetFuncMap(&deps.Deps{})

	tt := template.Must(template.New("test").Funcs(funcMap).Parse(tpl))

	var b bytes.Buffer
	err := tt.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// RunWithData
//
// Run the template test by executing the tpl given with data
func (t *TplTestSuite) RunTWithData(tpl string, expect interface{}, data interface{}) {
	b, err := t.Execute(tpl, data)
	if err != nil {
		assert.Contains(t.test, err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	assert.Equal(t.test, got, b)
}

// Run
//
// Run the template test by executing the tpl given
func (t *TplTestSuite) RunT(tpl string, expect interface{}) {
	b, err := t.Execute(tpl, map[string]string{})
	if err != nil {
		assert.Contains(t.test, err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	assert.Equal(t.test, got, b)
}
