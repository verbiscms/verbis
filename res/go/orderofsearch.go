package res

//type TypeOfPage struct {
//	PageType string
//	Data     interface{}
//}
//
//func (t *TemplateManager) orderOfSearch() TypeOfPage {
//	const op = "Templates.orderOfSearch"
//
//	data := TypeOfPage{
//		PageType: "page",
//		Data:     nil,
//	}
//
//	slug := t.post.Slug
//	slugArr := strings.Split(slug, "/")
//	last := slugArr[len(slugArr)-1]
//
//	theme := t.store.Site.GetThemeConfig()
//
//	if _, ok := theme.Resources[last]; ok {
//		data.PageType = "archive"
//		data.Data = t.post.Post.Resource
//		return data
//	}
//
//	if t.store.Categories.ExistsBySlug(last) {
//
//		cat, err := t.store.Categories.GetBySlug(last)
//		if err != nil {
//			return data
//		}
//
//		parentCat, err := t.store.Categories.GetByID(cat.Id)
//		if err != nil {
//			data.PageType = "category_child_archive"
//			data.Data = cat
//			return data
//		} else {
//			data.PageType = "category_archive"
//			data.Data = parentCat
//			return data
//		}
//	}
//
//	return data
//}
