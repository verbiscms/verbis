package templates

// Retrieve the assets path for the theme
func (t *TemplateFunctions) assetsPath() string {
	//return config.Theme.AssetsPath
	//
	return ""
}

// Retrieve the assets path for the theme
func (t *TemplateFunctions) storagePath() string {
	// TODO: Make dynamic?
	return "/storage"
}
