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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// newTestSuite
//
// Sets up up a TemplateFunctions with gin read
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

// execute
//
// Executes the templates with functions and returns the resulting
// html or an error if there was a problem executing the template.
func execute(tf *TemplateFunctions, tpl string, data interface{}) (string, error) {
	tt := template.Must(template.New("test").Funcs(tf.GetFunctions()).Parse(tpl))

	var b bytes.Buffer
	err := tt.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// runtv
//
// Run the template test by executing the tpl given
// with data.
func runtv(t *testing.T, tf *TemplateFunctions, tpl string, expect interface{}, data interface{}) {
	b, err := execute(tf, tpl, data)
	if err != nil {
		assert.Contains(t, err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")
	assert.Equal(t, got, b)
}

// runt
//
// Run the template test by executing the tpl given
func runt(t *testing.T, tf *TemplateFunctions, tpl string, expect interface{}) {
	b, err := execute(tf, tpl, map[string]string{})
	if err != nil {
		fmt.Println(err)
		assert.Contains(t, err.Error(), expect.(string))
		return
	}

	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")
	assert.Equal(t, got, b)
}

func Test_GetData(t *testing.T) {
	f := newTestSuite()

	categoryMock := mocks.CategoryRepository{}
	categoryMock.On("ExistsBySlug", mock.Anything).Return(false)
	f.store.Categories = &categoryMock

	uuid := uuid.New()
	time := time.Now()

	site := domain.Site{}
	siteMock := mocks.SiteRepository{}
	siteMock.On("GetGlobalConfig").Return(&site)

	theme := domain.ThemeConfig{Theme: domain.Theme{}}
	siteMock.On("GetThemeConfig").Return(&theme)

	f.options = domain.Options{
		ContactEmail:     "email",
		ContactTelephone: "phone",
		ContactAddress:   "address",
		SocialFacebook:   "facebook",
		SocialTwitter:    "twitter",
		SocialInstagram:  "instagram",
		SocialLinkedIn:   "linkedin",
		SocialYoutube:    "youtube",
		SocialPinterest:  "pinterest",
	}

	author := &domain.PostAuthor{}
	category := &domain.PostCategory{}
	resource := "resource"

	f.post = &domain.PostData{
		Post: domain.Post{
			Id:                1,
			UUID:              uuid,
			Slug:              "/slug",
			Title:             "Verbis",
			Status:            "published",
			Resource:          &resource,
			PageTemplate:      "pagetemplate",
			PageLayout:        "pagelayout",
			Fields:            nil,
			CodeInjectionHead: nil,
			CodeInjectionFoot: nil,
			UserId:            0,
			PublishedAt:       &time,
			CreatedAt:         &time,
			UpdatedAt:         &time,
			SeoMeta:           domain.PostSeoMeta{},
		},
		Author:   author,
		Category: category,
	}

	want := map[string]interface{}{
		"Type": TypeOfPage{
			PageType: "page",
			Data:     nil,
		},
		"Site":  &site,
		"Theme": theme.Theme,
		"Post": map[string]interface{}{
			"Id":           1,
			"UUID":         uuid,
			"Slug":         "/slug",
			"Title":        "Verbis",
			"Status":       "published",
			"Resource":     &resource,
			"PageTemplate": "pagetemplate",
			"PageLayout":   "pagelayout",
			"PublishedAt":  &time,
			"UpdatedAt":    &time,
			"CreatedAt":    &time,
			"Author":       author,
			"Category":     category,
		},
		"Options": map[string]interface{}{
			"Social": map[string]interface{}{
				"Facebook":  "facebook",
				"Twitter":   "twitter",
				"Youtube":   "youtube",
				"LinkedIn":  "linkedin",
				"Instagram": "instagram",
				"Pintrest":  "pinterest",
			},
			"Contact": map[string]interface{}{
				"Email":     "email",
				"Telephone": "phone",
				"Address":   "address",
			},
		},
	}

	assert.EqualValues(t, want, f.GetData())
}
