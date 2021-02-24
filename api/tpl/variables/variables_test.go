package variables

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData(t *testing.T) {
	td := TemplateData{
		Site:    domain.Site{},
		Theme:   domain.ThemeConfig{},
		Post:    domain.PostDatum{},
		Options: Options{},
	}

	got := Data(&deps.Deps{
		Site:    domain.Site{},
		Theme:   &domain.ThemeConfig{},
		Options: &domain.Options{},
	}, &gin.Context{}, &domain.PostDatum{})

	assert.Equal(t, td, got)
}
