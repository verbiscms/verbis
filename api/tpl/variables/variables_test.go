package variables

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/services/site"
	"testing"
)

func TestData(t *testing.T) {
	td := TemplateData{
		Site:    domain.Site{},
		Theme:   domain.ThemeConfig{},
		Post:    domain.PostDatum{},
		Options: Options{},
	}

	mockSite := &mocks.Repository{}
	mockSite.On("Global").Return(domain.Site{})

	got := Data(&deps.Deps{
		Site:    mockSite,
		Config:  &domain.ThemeConfig{},
		Options: &domain.Options{},
	}, &gin.Context{}, &domain.PostDatum{})

	assert.Equal(t, td, got)
}
