package templates

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/fields"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"html/template"
	"strings"
	"time"
)

type TemplateManager struct {
	gin          *gin.Context
	post         *domain.PostData
	site         *domain.Site
	store        *models.Store
	options      domain.Options
	config       config.Configuration
	themeConfig  domain.ThemeConfig
	fieldService fields.FieldService
	token        string
}

type TypeOfPage struct {
	PageType string
	Data     interface{}
}

// Obtain all of the paths in to variables
// for use with testing
var (
	themePath   = paths.Theme()
	basePath    = paths.Base()
	adminPath   = paths.Admin()
	apiPath     = paths.Api()
	storagePath = paths.Storage()
	uploadsPath = paths.Uploads()
)

// NewManager - Construct
func NewManager(g *gin.Context, s *models.Store, p *domain.PostData, c config.Configuration) *TemplateManager {
	return &TemplateManager{
		gin:         g,
		post:        p,
		site:        s.Site.GetGlobalConfig(),
		store:       s,
		options:     s.Options.GetStruct(),
		themeConfig: s.Site.GetThemeConfig(),
		fieldService: fields.NewService(s, *p),
		config: c,
	}
}

// Get all template functions
func (t *TemplateManager) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		// Attributes
		"body": t.body,
		"lang": t.lang,
		// Auth
		"auth":  t.auth,
		"admin": t.admin,
		// Categories
		"category":         t.getCategory,
		"categoryByName":   t.getCategoryByName,
		"categoryByParent": t.getCategoryByParent,
		"categories":       t.getCategories,
		// Cast
		"toBool":     cast.ToBool,
		"toTime":     cast.ToTime,
		"toDuration": cast.ToDuration,
		"toString":   cast.ToString,
		"toInt":      cast.ToInt,
		"toInt8":     cast.ToInt8,
		"toInt16":    cast.ToInt16,
		"toInt32":    cast.ToInt32,
		"toInt64":    cast.ToInt64,
		"toUInt":     cast.ToUint,
		"toUInt8":    cast.ToUint8,
		"toUInt16":   cast.ToUint16,
		"toUInt32":   cast.ToUint32,
		"toUInt64":   cast.ToUint64,
		"toFloat32":  cast.ToFloat32,
		"toFloat64":  cast.ToFloat64,
		// Date & Time
		"now":            time.Now,
		"date":           t.date,
		"dateInZone":     t.dateInZone,
		"ago":            t.ago,
		"duration":       t.duration,
		"htmlDate":       t.htmlDate,
		"htmlDateInZone": t.htmlDateInZone,
		// Dict
		"dict": t.dict,
		// Fields
		//"field":    t.fieldService.GetField,
		//"fields":   t.fieldService.GetFields,
		//"layout":   t.fieldService.GetLayout,
		//"layouts":  t.fieldService.GetLayouts,
		//"repeater": t.fieldService.GetRepeater,
		//"flexible": t.fieldService.GetFlexible,
		// Header & Footer
		"verbisHead": t.header,
		"verbisFoot": t.footer,
		"metaTitle":  t.metaTitle,
		// Math
		"add":      t.add,
		"subtract": t.subtract,
		"divide":   t.divide,
		"multiply": t.multiply,
		"mod":      t.modulus,
		"round":    t.round,
		"ceil":     t.ceil,
		"floor":    t.floor,
		"min":      t.min,
		"max":      t.max,
		// Media
		"media": t.getMedia,
		// OS
		"env":       t.env,
		"expandEnv": t.expandEnv,
		// Partials
		"partial": t.partial,
		// Posts
		"post":           t.getPost,
		"posts":          t.getPosts,
		"paginationPage": t.getPaginationPage,
		// Paths
		"basePath":      t.basePath,
		"adminPath":     t.adminPath,
		"apiPath":       t.apiPath,
		"themePath":     t.themePath,
		"uploadsPath":   t.uploadsPath,
		"assetsPath":    t.assetsPath,
		"storagePath":   t.storagePath,
		"templatesPath": t.templatesPath,
		"layoutsPath":   t.layoutsPath,
		// Rand
		"randInt":      t.randInt,
		"randFloat":    t.randFloat,
		"randAlpha":    t.randAlpha,
		"randAlphaNum": t.randAlphaNum,
		// Reflect
		"kindIs":     t.kindIs,
		"kindOf":     t.kindOf,
		"typeOf":     t.typeOf,
		"typeIs":     t.typeIs,
		"typeIsLike": t.typeIsLike,
		// Regex"
		"regexMatch":             t.regexMatch,
		"regexFindAll":           t.regexFindAll,
		"regexFind":              t.regexFind,
		"regexReplaceAll":        t.regexReplaceAll,
		"regexReplaceAllLiteral": t.regexReplaceAllLiteral,
		"regexSplit":             t.regexSplit,
		"regexQuoteMeta":         t.regexQuoteMeta,
		// Safe
		"safeHTML":     t.safeHTML,
		"safeHTMLAttr": t.safeHTMLAttr,
		"safeCSS":      t.safeCSS,
		"safeJS":       t.safeJS,
		"safeJSStr":    t.safeJSStr,
		"safeURL":      t.safeUrl,
		// Slice
		"slice":   t.slice,
		"append":  t.append,
		"prepend": t.prepend,
		"first":   t.first,
		"last":    t.last,
		"reverse": t.reverse,
		// Strings
		"trim":     strings.TrimSpace,
		"upper":    strings.ToUpper,
		"lower":    strings.ToLower,
		"title":    strings.Title,
		"replace":  t.replace,
		"substr":   t.substr,
		"trunc":    t.trunc,
		"ellipsis": t.ellipsis,
		// URL
		"baseUrl": t.getBaseURL,
		"scheme":  t.getScheme,
		"host":    t.getHost,
		"fullUrl": t.getFullURL,
		"url":     t.getURL,
		"query":   t.getQueryParams,
		// Users
		"user":  t.getUser,
		"users": t.getUsers,
		// Util
		"len":     t.len,
		"explode": t.explode,
		"implode": t.implode,
	}

	return funcMap
}

// GetData - Returns all the necessary data for template usage.
func (t *TemplateManager) GetData() map[string]interface{} {

	theme := t.store.Site.GetThemeConfig()

	data := map[string]interface{}{
		"Type":  t.orderOfSearch(),
		"Site":  t.store.Site.GetGlobalConfig(),
		"Theme": theme.Theme,
		//"Token": csrf.GetToken(t.gin),
		"Post": map[string]interface{}{
			"Id":           t.post.Id,
			"UUID":         t.post.UUID,
			"Slug":         t.post.Slug,
			"Title":        t.post.Title,
			"Status":       t.post.Status,
			"Resource":     t.post.Resource,
			"PageTemplate": t.post.PageTemplate,
			"PageLayout":   t.post.PageLayout,
			"PublishedAt":  t.post.PublishedAt,
			"UpdatedAt":    t.post.UpdatedAt,
			"CreatedAt":    t.post.CreatedAt,
			"Author":       t.post.Author,
			"Category":     t.post.Category,
		},
		"Options": map[string]interface{}{
			"Social": map[string]interface{}{
				"Facebook":  t.options.SocialFacebook,
				"Twitter":   t.options.SocialTwitter,
				"Youtube":   t.options.SocialYoutube,
				"LinkedIn":  t.options.SocialLinkedIn,
				"Instagram": t.options.SocialInstagram,
				"Pintrest":  t.options.SocialPinterest,
			},
			"Contact": map[string]interface{}{
				"Email":     t.options.ContactEmail,
				"Telephone": t.options.ContactTelephone,
				"Address":   t.options.ContactAddress,
			},
		},
	}

	return data
}

func (t *TemplateManager) orderOfSearch() TypeOfPage {
	const op = "Templates.orderOfSearch"

	data := TypeOfPage{
		PageType: "page",
		Data:     nil,
	}

	slug := t.post.Slug
	slugArr := strings.Split(slug, "/")
	last := slugArr[len(slugArr)-1]

	theme := t.store.Site.GetThemeConfig()

	if _, ok := theme.Resources[last]; ok {
		data.PageType = "archive"
		data.Data = t.post.Post.Resource
		return data
	}

	if t.store.Categories.ExistsBySlug(last) {

		cat, err := t.store.Categories.GetBySlug(last)
		if err != nil {
			return data
		}

		parentCat, err := t.store.Categories.GetById(cat.Id)
		if err != nil {
			data.PageType = "category_child_archive"
			data.Data = cat
			return data
		} else {
			data.PageType = "category_archive"
			data.Data = parentCat
			return data
		}
	}

	return data
}
