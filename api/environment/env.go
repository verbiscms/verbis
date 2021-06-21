// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package environment

import (
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	pkgValidate "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"strconv"

	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

// Env defines the environment variables set in the .env file.
//nolint
type Env struct {

	// Prod, production or dev.
	AppEnv string `json:"APP_ENV"`

	// If Verbis is in debug mode (true, false).
	AppDebug string `json:"APP_DEBUG"`

	// The port the server should listen to.
	AppPort string `json:"APP_PORT" binding:"required"`

	// The database port.
	DbDriver string `json:"DB_DRIVER" binding:"required"`

	// The database host (IP) for the store.
	DbHost string `json:"DB_HOST" binding:"required"`

	// The database port for the store.
	DbPort string `json:"DB_PORT" binding:"required"`

	// The database name.
	DbDatabase string `json:"DB_DATABASE" binding:"required"`

	// The database user name.
	DbUser string `json:"DB_USERNAME" binding:"required"`

	// The database port.
	DbPassword string `json:"DB_PASSWORD" binding:"required"`

	// The database port.
	DbSchema string `json:"DB_SCHEMA"`

	// The database port.
	MailDriver string `json:"MAIL_DRIVER"`

	// The mailing from address.
	MailFromAddress string `json:"MAIL_FROM_ADDRESS"`

	// The mailing from name.
	MailFromName string `json:"MAIL_FROM_NAME"`

	// The API key for Sparkpost.
	SparkpostAPIKey string `json:"SPARKPOST_API_KEY"`

	// The url for Sparkpost (could be EU).
	SparkpostURL string `json:"SPARKPOST_URL"`

	// The API key for MailGun.
	MailGunAPIKey string `json:"MAILGUN_API_KEY"`

	// The url for MailGun.
	MailGunURL string `json:"MAILGUN_URL"`

	// The domain for MailGun.
	MailGunDomain string `json:"MAILGUN_DOMAIN"`

	// The API key for SendGrid.
	SendGridAPIKey string `json:"SENDGRID_API_KEY"`
}

var (
	// The absolute path of the Verbis project, where the
	// .env is stored.
	basePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	// The environment file extension.
	envExt = ".env"
)

const (
	// DefaultPort is the default port Verbis should sit
	// on when none is defined.
	DefaultPort = 5000
)

// Load
//
// Load populates environment, loads and validates the
// environment file.
//
// Returns errors.INVALID if the env file failed to load.
func Load() (*Env, error) {
	const op = "environment.Load"

	err := godotenv.Load(basePath + "/" + envExt)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Could not load the .env file", Operation: op, Err: err}
	}

	return &Env{
		AppEnv:          os.Getenv("APP_ENV"),
		AppDebug:        os.Getenv("APP_DEBUG"),
		AppPort:         os.Getenv("APP_PORT"),
		DbDriver:        os.Getenv("DB_DRIVER"),
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          os.Getenv("DB_PORT"),
		DbDatabase:      os.Getenv("DB_DATABASE"),
		DbUser:          os.Getenv("DB_USERNAME"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		DbSchema:        os.Getenv("DB_SCHEMA"),
		MailDriver:      os.Getenv("MAIL_DRIVER"),
		MailFromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		MailFromName:    os.Getenv("MAIL_FROM_NAME"),
		SparkpostAPIKey: os.Getenv("SPARKPOST_API_KEY"),
		SparkpostURL:    os.Getenv("SPARKPOST_URL"),
		MailGunAPIKey:   os.Getenv("MAILGUN_API_KEY"),
		MailGunURL:      os.Getenv("MAILGUN_URL"),
		MailGunDomain:   os.Getenv("MAILGUN_DOMAIN"),
		SendGridAPIKey:  os.Getenv("SENDGRID_API_KEY"),
	}, nil
}

// Validate
//
// Validates the environment file for missing keys, if
// there are no validation errors, nil will be
// returned.
func (e *Env) Validate() validation.Errors {
	v := validation.New()
	err := v.Package.Struct(e)
	if err != nil {
		validationErrors := err.(pkgValidate.ValidationErrors)
		return v.Process(validationErrors)
	}
	return nil
}

// Set
//
// Accepts a key, value pair and writes to the .env
// file when installing.
func (e *Env) Set(key string, value interface{}) error {
	const op = "Env.Set"

	val, err := cast.ToStringE(value)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error casting value to string", Operation: op, Err: err}
	}

	path := basePath + "/" + envExt
	env, err := godotenv.Read(path)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error reading env file with the path " + path, Operation: op, Err: err}
	}

	env[key] = val

	err = godotenv.Write(env, path)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error writing env file with the path " + path, Operation: op, Err: err}
	}

	return nil
}

// Port
//
// Returns the env port as an integer, if the strconv
// produced an error, the default port of 5000 will
// be used,
func (e *Env) Port() int {
	const op = "Env.Port"
	port, err := strconv.Atoi(e.AppPort)
	if err != nil {
		log.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast app port to int using port 5000", Operation: op, Err: err}).Error()
		return DefaultPort
	}
	return port
}

// IsProduction
//
// If the application is set to production.
func (e *Env) IsProduction() bool {
	return e.AppEnv == "production" || e.AppEnv == "prod"
}

// IsDebug
//
// If the application is set to debug.
func (e *Env) IsDebug() bool {
	return e.AppDebug != "false"
}
