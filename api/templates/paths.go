package templates

// basePath
//
// Returns the base path of the project
func (t *TemplateFunctions) basePath() string {
	return basePath
}

// adminPath
//
// Returns the admin path of the project
func (t *TemplateFunctions) adminPath() string {
	return adminPath
}

// apiPath
//
// Returns the API path of the project
func (t *TemplateFunctions) apiPath() string {
	return apiPath
}

// themePath
//
// Returns the currently active theme path
func (t *TemplateFunctions) themePath() string {
	return themePath
}

// uploadsPath
//
// Returns the uploads path of the project
func (t *TemplateFunctions) uploadsPath() string {
	return uploadsPath
}

// assetsPath
//
// Returns the assets path for the theme
func (t *TemplateFunctions) assetsPath() string {
	return t.themeConfig.AssetsPath
}

// storagePath
//
// Returns the storage path for the project
func (t *TemplateFunctions) storagePath() string {
	return storagePath
}

// templatesPath
//
// Returns the directory where page templates
// are stored.
func (t *TemplateFunctions) templatesPath() string {
	return themePath + t.themeConfig.TemplateDir
}

// layoutsPath
//
// Returns the directory where page layouts
// are stored.
func (t *TemplateFunctions) layoutsPath() string {
	return themePath + t.themeConfig.LayoutDir
}
