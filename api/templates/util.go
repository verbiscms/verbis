package templates

// Get all fields for template
func (t *TemplateFunctions) getFullUrl() string {
	return t.gin.Request.Host + t.gin.Request.URL.String()
}
