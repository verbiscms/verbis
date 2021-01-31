package funcs

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFuncs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ctx.Request, _ = http.NewRequest("GET", "/page", nil)
	engine.Use(location.Default())

	engine.GET("/page", func(g *gin.Context) {
		ctx = g
		return
	})

	req, err := http.NewRequest("GET", "http://verbiscms.com/page?page=2&foo=bar", nil)
	assert.NoError(t, err)
	engine.ServeHTTP(rr, req)

	f := Funcs{
		deps: &deps.Deps{
			Options: domain.Options{
				GeneralLocale: "en-gb",
			},
		},
		ctx: ctx,
		post: &domain.PostData{
			Post: domain.Post{
				Id:           1,
				Slug:         "/page",
				Title:        "My Verbis Page",
				Status:       "published",
				Resource:     nil,
				PageTemplate: "single",
				PageLayout:   "main",
				UserId:       1,
			},
			Fields: []domain.PostField{
				{PostId: 1, Type: "text", Name: "text", Key: "", OriginalValue: "Hello World!"},
			},
		},
	}

	os.Setenv("foo", "bar")

	v := variables.New(&deps.Deps{}, ctx, f.post)

	for _, ns := range f.getNamespaces() {
		for _, mm := range ns.MethodMappings {
			for _, e := range mm.Examples {
				file, err := template.New("test").Funcs(f.Map()).Parse(e[0])
				if err != nil {
					t.Errorf("test failed for %s: %s", mm.Name, err.Error())
					continue
				}

				var tpl bytes.Buffer
				err = file.Execute(&tpl, v.Get())
				if err != nil {
					t.Error(err)
				}

				assert.Equal(t, e[1], html.UnescapeString(tpl.String()))
			}
		}
	}
}
