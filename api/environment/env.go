// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package environment

import (
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	pkgValidate "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Env struct {
	AppEnv          string `json:"APP_ENV"`
	AppDebug        string `json:"APP_DEBUG"`
	AppPort         string `json:"APP_PORT" binding:"required"`
	DbHost          string `json:"DB_HOST" binding:"required"`
	DbPort          string `json:"DB_PORT" binding:"required"`
	DbDatabase      string `json:"DB_DATABASE" binding:"required"`
	DbUser          string `json:"DB_USERNAME" binding:"required"`
	DbPassword      string `json:"DB_PASSWORD" binding:"required"`
	SparkpostApiKey string `json:"SPARKPOST_API_KEY"`
	SparkpostUrl    string `json:"SPARKPOST_URL"`
	MailFromAddress string `json:"MAIL_FROM_ADDRESS"`
	MailFromName    string `json:"MAIL_FROM_NAME"`
}

type Mail struct {
	SparkpostApiKey string `json:"SPARKPOST_API_KEY"`
	SparkpostUrl    string `json:"SPARKPOST_URL"`
	FromAddress     string `json:"MAIL_FROM_ADDRESS"`
	FromName        string `json:"MAIL_FROM_NAME"`
}

var (
	// The absolute path of the Verbis project, where the
	// .env is stored.
	basePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	// The environment file extension.
	envExt      = ".env"
)

// Load
//
// Load populates environment, loads and validates the environment file.
// Returns errors.INVALID if the env file failed to load.
func Load() (*Env, error) {
	const op = "environment.Load"

	err := godotenv.Overload(basePath + "/" + envExt)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Could not load the .env file", Operation: op, Err: err}
	}

	return &Env{
		AppEnv:          os.Getenv("APP_ENV"),
		AppDebug:        os.Getenv("APP_DEBUG"),
		AppPort:         os.Getenv("APP_PORT"),
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          os.Getenv("DB_PORT"),
		DbDatabase:      os.Getenv("DB_DATABASE"),
		DbUser:          os.Getenv("DB_USERNAME"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		SparkpostApiKey: os.Getenv("SPARKPOST_API_KEY"),
		SparkpostUrl:    os.Getenv("SPARKPOST_URL"),
		MailFromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		MailFromName:    os.Getenv("MAIL_FROM_NAME"),
	}, nil
}

// Validate
//
// the environment file for missing keys
func (e *Env) Validate() validation.Errors {
	v := validation.New()
	err := v.Package.Struct(e)
	if err != nil {
		validationErrors := err.(pkgValidate.ValidationErrors)
		return v.Process(validationErrors)
	}
	return nil
}

// Port
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
//
func (e *Env) Port() int {
	const op = "Env.Port"
	port, err := strconv.Atoi(e.AppPort)
	if err != nil {
		log.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast app port to int using port 5000", Operation: op, Err: err}).Error()
		return 5000
	}
	return port
}

// ConnectString
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
//
func (e *Env) ConnectString() string {
	return e.DbUser + ":" + e.DbPassword + "@tcp(" + e.DbHost + ":" + e.DbPort + ")/" + e.DbDatabase + "?tls=false&parseTime=true&multiStatements=true"
}

// MailConfig
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
//
func (e *Env) MailConfig() Mail {
	return Mail{
		FromAddress:     e.MailFromAddress,
		FromName:        e.MailFromName,
		SparkpostApiKey: e.SparkpostApiKey,
		SparkpostUrl:    e.SparkpostUrl,
	}
}

// IsProduction
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
//
func (e *Env) IsProduction() bool {
	return e.AppEnv == "production" || e.AppEnv == "prod"
}

// IsDebug
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
//
func (e *Env) IsDebug() bool {
	return e.AppDebug != "false"
}
