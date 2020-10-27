package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
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
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get", nil)

	p := &domain.Post{}
	if len(args) == 1 {
		data := []byte(args[0])
		p = &domain.Post{
			Fields: (*json.RawMessage)(&data),
		}
	}

	return NewFunctions(g, &models.Store{}, p)
}

// runt - Run the template test by executing the tpl give.
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