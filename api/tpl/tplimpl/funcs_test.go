package tplimpl

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Setup(t *testing.T) *Funcs {
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

	os.Setenv("foo", "bar")

	f := Funcs{
		&deps.Deps{
			Options: domain.Options{
				GeneralLocale: "en-gb",
			},
		},
		&internal.TemplateDeps{
			Context: ctx,
			Post: &domain.PostData{
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
		},
	}

	return &f
}

// Test all internal template function mappings
func TestFuncs(t *testing.T) {
	f := Setup(t)

	v := variables.New(&deps.Deps{}, f.TemplateDeps.Context, f.TemplateDeps.Post)

	for _, ns := range f.getNamespaces() {
		for _, mm := range ns.MethodMappings {
			for _, e := range mm.Examples {
				file, err := template.New("test").Funcs(f.FuncMap()).Parse(e[0])
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

func TestFuncs_FuncMap(t *testing.T) {

	tt := map[string]struct {
		namespaces internal.FuncNamespaces
		want       template.FuncMap
		panics     bool
	}{
		"Success": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"func": {
						Method: nil,
						Name:   "func",
					},
				}},
			},
			template.FuncMap{"func": nil},
			false,
		},
		"Duplicate Func": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"test":    {Method: nil, Name: "replace"},
					"replace": {Method: nil, Name: "replace"},
				}},
			},
			nil,
			true,
		},
		"Duplicate Alias": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"test": {Method: nil, Name: "test", Aliases: []string{"test"}},
				}},
			},
			template.FuncMap{},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := &Funcs{}

			if test.panics {
				assert.Panics(t, func() {
					f.getFuncs(test.namespaces)
				})
				return
			}

			assert.Equal(t, test.want, f.getFuncs(test.namespaces))
		})
	}
}