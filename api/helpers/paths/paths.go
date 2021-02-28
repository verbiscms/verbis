// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"os"
	"path/filepath"
)

type Paths struct {
	Base      string
	Admin     string
	API       string
	Uploads   string
	Migration string
	Storage   string
	Web       string
	Forms     string
}

const (
	Admin   = "/admin"
	Api     = "/api"
	Storage = "/storage"
	Web     = Api + "/web"
	Uploads = Storage + "/uploads"
)

func Get() Paths {
	base := base()
	return Paths{
		Base:      base,
		Admin:     base + Admin,
		API:       base + Api,
		Migration: base + Api + migration(),
		Uploads:   base + Uploads,
		Storage:   base + Storage,
		Web:       base + Web,
	}
}

// Base path of project
func base() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// BaseCheck environment is passable to run Terminal
func BaseCheck() error {
	const op = "paths.BaseCheck"
	basePath := base()

	if !files.Exists(basePath + "/.env") {
		return fmt.Errorf("Could not locate the .env file in the current directory")
	}

	if !files.DirectoryExists(basePath + "/admin") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis admin folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath+"/admin")}
	}

	if !files.DirectoryExists(basePath + "/storage") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis storage folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath+"/storage")}
	}

	if !files.DirectoryExists(basePath + "/storage") {
		return &errors.Error{Code: errors.INVALID, Message: "Could not locate the Verbis storage folder in the current directory", Operation: op, Err: fmt.Errorf("%s does not exist", basePath+"/storage")}
	}

	return nil
}

//// Admin path of project
//func Admin() string {
//	return Base() + "/admin"
//}
//
//// API path of project
//func Api() string {
//	return Base() + "/api"
//}

// Migration is the Database migration path
func migration() string {
	if api.SuperAdmin {
		return "/database/migrations"
	} else {
		return "/database"
	}
}

// Theme path
//func Theme() string {
//	return Base() + "/theme"
//}

//// Storage path
//func Storage() string {
//	return Base() + "/storage"
//}

// Storage path
//func Uploads() string {
//	return Storage() + "/uploads"
//}

// Web (Verbis specific)
//func Web() string {
//	return Api() + "/web"
//}

//func Forms() string {
//	return Storage() + "/forms"
//}
