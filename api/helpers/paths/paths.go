package paths

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"os"
)

// Base path of project
func Base() string {
	basePath := ""

	basePath, err := os.Getwd()
	if err != nil {
		return ""
	}

	return basePath
}

// BaseCheck environment is passable to run Terminal
func BaseCheck() error {
	const op = "paths.BaseCheck"
	basePath := Base()

	if !files.Exists(basePath + "/.env") {
		return fmt.Errorf("Could not locate the .env file in the current directory")
	}

	if !files.DirectoryExists(basePath + "/admin") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis admin folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath + "/admin")}
	}

	if !files.DirectoryExists(basePath + "/storage") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis storage folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath + "/storage")}
	}

	if !files.DirectoryExists(basePath + "/config") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis config folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath + "/config")}
	}

	if !files.DirectoryExists(basePath + "/storage") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis storage folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath + "/storage")}
	}

	return nil
}

// Admin path of project
func Admin() string {
	return Base() + "/admin"
}

// API path of project
func Api() string {
	return Base() + "/api"
}

// Migration is the Database migration path
func Migration() string {
	if api.SuperAdmin {
		return Api() + "/database/migrations"
	} else {
		return Api() + "/database"
	}
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

// PublicUploads Path
func PublicUploads() string {
	return config.Media.UploadPath
}

// Templates path
func Templates() string {
	// TODO - Make dynamic based on config
	return Theme() + "/templates"
}

// Web (Verbis specific)
func Web() string {
	return Api() + "/web"
}