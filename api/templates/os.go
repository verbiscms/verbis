package templates

import "os"

// env
//
// Retrieve an environment variable by key
//
// Example: {{ env "APP_DEBUG" }}
func (t *TemplateFunctions) env(key string) string {
	return os.Getenv(key)
}

// expandEnv
//
// Retrieve an environment variable by key and
// substitute variables in a string.
//
//
func (t *TemplateFunctions) expandEnv(str string) string {
	return os.ExpandEnv(str)
}
