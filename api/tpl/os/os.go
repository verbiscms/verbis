package os

import "os"

// env
//
// Retrieve an environment variable by key
//
// Example: {{ env "APP_DEBUG" }}
func (ns *Namespace) env(key string) string {
	return os.Getenv(key)
}

// expandEnv
//
// Retrieve an environment variable by key and
// substitute variables in a string.
//
// {{ expandEnv "Welcome to $APP_NAME" }}
func (ns *Namespace) expandEnv(str string) string {
	return os.ExpandEnv(str)
}
