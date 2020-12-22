package templates

import "os"

func (t *TemplateFunctions) env(key string) string {
	return os.Getenv(key)
}

func (t *TemplateFunctions) expandEnv(key string) string {
	return os.ExpandEnv(key)
}