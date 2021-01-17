package tpl

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TplTestSuite struct {
	suite.Suite
	*TemplateManager
}

func TestTpl(t *testing.T) {
	suite.Run(t, new(TplTestSuite))
}

func (t *TplTestSuite) BeforeTest(suiteName, testName string) {
	err := logger.Init(config.Configuration{})

	gin.SetMode(gin.TestMode)
	log.SetOutput(ioutil.Discard)

	t.TemplateManager = t.GetManager()

	t.NoError(err)
}

func (t *TplTestSuite) AfterTest() {
	t.TemplateManager = nil
}

// GetManager
//
// Sets up up a TemplateManager with gin read
// for testing.
func (t *TplTestSuite) GetManager() *TemplateManager {
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get", nil)

	mockOptions := mocks.OptionsRepository{}
	mockOptions.On("GetStruct").Return(domain.Options{}, nil)

	mockSite := mocks.SiteRepository{}
	mockSite.On("GetGlobalConfig").Return(&domain.Site{}, nil)
	mockSite.On("GetThemeConfig").Return(domain.ThemeConfig{})

	return NewManager(g, &models.Store{
		Options: &mockOptions,
		Site:    &mockSite,
	}, &domain.PostData{}, config.Configuration{})
}

// Execute
//
// Executes the templates with functions and returns the resulting
// html or an error if there was a problem executing the template.
func (t *TplTestSuite) Execute(tpl string, data interface{}) (string, error) {
	tt := template.Must(template.New("test").Funcs(t.GetFunctions()).Parse(tpl))

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
		t.Contains(err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	t.Equal(got, b)
}

// Run
//
// Run the template test by executing the tpl given
func (t *TplTestSuite) RunT(tpl string, expect interface{}) {
	b, err := t.Execute(tpl, map[string]string{})
	if err != nil {
		t.Contains(err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")

	t.Equal(got, b)
}


type Tester struct {
	test string
}

func data() Tester {
	return Tester{}
}

func (t *TplTestSuite) TestR() {
	t.RunTWithData(`{{ if . }} Will content be printed here? {{ end }}`, "", data())
}

