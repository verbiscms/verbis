package paths

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/environment"
	"os"
	"path/filepath"
	"runtime"
)

// Base path of project
func Base() string {
	path := ""
	if environment.Env.AppEnv == "production" || environment.Env.AppEnv == "prod" {
		path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		_, b, _, _ := runtime.Caller(0)
		path = filepath.Join(filepath.Dir(b), "../../..")
	}
	return path
}

// Admin path of project
func Admin() string {
	return Base() + "/admin"
}

// API path of project
func Api() string {
	return Base() + "/api"
}

// Database migration path
func Migration() string {
	return Api() + "/database/migrations"
}

// Theme path
func Theme() string {
	return Base() + "/theme"
}

// Storage path
func Storage() string {
	return Base() + "/storage"
}

// Storage path
func Uploads() string {
	return Storage() + "/uploads"
}

// Public Uploads Path
func PublicUploads() string {
	return config.Media.UploadPath
}

// Templates path
func Templates() string {
	// TODO - Make dynamic based on config
	return Theme() + "/templates"
}