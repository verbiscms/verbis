package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"time"
)

func (t *TplTestSuite) Test_GetData() {

	categoryMock := mocks.CategoryRepository{}
	categoryMock.On("ExistsBySlug", mock.Anything).Return(false)
	t.store.Categories = &categoryMock

	uuid := uuid.New()
	time := time.Now()

	site := domain.Site{}
	siteMock := mocks.SiteRepository{}
	siteMock.On("GetGlobalConfig").Return(&site)

	theme := domain.ThemeConfig{Theme: domain.Theme{}}
	siteMock.On("GetThemeConfig").Return(&theme)

	t.options = domain.Options{
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

	t.post = &domain.PostData{
		Post: domain.Post{
			Id:                1,
			UUID:              uuid,
			Slug:              "/slug",
			Title:             "Verbis",
			Status:            "published",
			Resource:          &resource,
			PageTemplate:      "pagetemplate",
			PageLayout:        "pagelayout",
			CodeInjectionHead: nil,
			CodeInjectionFoot: nil,
			UserId:            0,
			PublishedAt:       &time,
			CreatedAt:         &time,
			UpdatedAt:         &time,
			SeoMeta:           domain.PostOptions{},
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

	t.EqualValues(want, t.GetData())
}
