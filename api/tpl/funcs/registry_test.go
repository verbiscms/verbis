package funcs

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFuncs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(rr)
	g.Request, _ = http.NewRequest("GET", "/get", nil)


	f := Funcs{
		deps: &deps.Deps{},
		ctx: g,
		post: &domain.PostData{},
	}

	v := variables.New(&deps.Deps{}, g, &domain.PostData{})

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