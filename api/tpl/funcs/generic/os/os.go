package os

import "os"

// Env
//
// Retrieve an environment variable by key
//
// Example: {{ env "APP_DEBUG" }}
func (ns *Namespace) Env(key string) string {
	return os.Getenv(key)
}

// ExpandEnv
//
// Retrieve an environment variable by key and
// substitute variables in a string.
//
// Example: {{ expandEnv "Welcome to $APP_NAME" }}
func (ns *Namespace) ExpandEnv(str string) string {
	return os.ExpandEnv(str)
}
