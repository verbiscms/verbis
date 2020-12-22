package templates

import "os"

// env
//
// Retrieve a environment variable by key
func (t *TemplateFunctions) env(key string) string {
	return os.Getenv(key)
}

// expandEnv
//
// Retrieve a environment variable by key and
// substitute variables in a string.
func (t *TemplateFunctions) expandEnv(key string) string {
	return os.ExpandEnv(key)
}
