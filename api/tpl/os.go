package tpl

import "os"

// env
//
// Retrieve an environment variable by key
//
// Example: {{ env "APP_DEBUG" }}
func (t *TemplateManager) env(key string) string {
	return os.Getenv(key)
}

// expandEnv
//
// Retrieve an environment variable by key and
// substitute variables in a string.
//
// {{ expandEnv "Welcome to $APP_NAME" }}
func (t *TemplateManager) expandEnv(str string) string {
	return os.ExpandEnv(str)
}
