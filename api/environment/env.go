// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package environment

import (
	"bytes"
	_ "embed"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"html/template"
	"io/ioutil"
	"strconv"

	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

// Env defines the environment variables set in the .env file.
type Env struct {
	// Prod, production or dev.
	AppEnv string `json:"APP_ENV"`
	// If Verbis is in debug mode (true, false).
	AppDebug string `json:"APP_DEBUG"`
	// The port the server should listen to.
	AppPort string `json:"APP_PORT" binding:"required"`
	// The database port.
	DbDriver string `json:"DB_DRIVER" binding:"required"` //nolint
	// The database host (IP) for the store.
	DbHost string `json:"DB_HOST" binding:"required"` //nolint
	// The database port for the store.
	DbPort string `json:"DB_PORT" binding:"required"` //nolint
	// The database name.
	DbDatabase string `json:"DB_DATABASE" binding:"required"` //nolint
	// The database user name.
	DbUser string `json:"DB_USERNAME" binding:"required"` //nolint
	// The database port.
	DbPassword string `json:"DB_PASSWORD" binding:"required"` //nolint
	// The database port.
	DbSchema string `json:"DB_SCHEMA"` //nolint
	// The cache driver to use.
	CacheDriver string `json:"CACHE_DRIVER"`
	// Redis IP address.
	RedisAddress string `json:"REDIS_ADDRESS"`
	// The password for Redis.
	RedisPassword string `json:"REDIS_PASSWORD"`
	// The database to use for Redis.
	RedisDb string `json:"REDIS_DB"` //nolint
	// The IP addresses to use for MemCached.
	MemCachedHosts string `json:"MEMCACHED_HOSTS"`
	// The access key for AWS storage.
	AWSAccessKey string `json:"STORAGE_AWS_ACCESS_KEY"`
	// The access secret for AWS storage.
	AWSSecret string `json:"STORAGE_AWS_SECRET"`
	// The JSON file for GCP storage.
	GCPJson string `json:"STORAGE_GCP_JSON_FILE"`
	// The Project ID for GCP storage.
	GCPProjectID string `json:"STORAGE_GCP_PROJECT_ID"`
	// The account details for Azure storage.
	AzureAccount string `json:"STORAGE_AZURE_ACCOUNT"`
	// The account details for Azure storage.
	AzureKey string `json:"STORAGE_AZURE_KEY"`
	// The mail driver to use.
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
	// basePath is the absolute path of the Verbis project,
	// where the .env is stored.
	basePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	// EnvExtension is the environment file extension.
	EnvExtension = ".env"
	// The default environment when no file is loaded.
	defaults = Env{
		AppEnv:   "dev",
		AppDebug: "false",
		AppPort:  strconv.Itoa(DefaultPort),
	}
	// The default tpl file for the environment (used when
	// installing).
	//go:embed env.txt
	envTpl string
)

const (
	// DefaultPort is the default port Verbis should sit
	// on when none is defined.
	DefaultPort = 5000
)

// Load populates environment, loads and validates the
// environment file.
// Returns errors.INVALID if the env file failed to load.
func Load() (*Env, error) {
	const op = "Environment.Load"
	err := godotenv.Load(filepath.Join(basePath, EnvExtension))
	if err != nil {
		return &defaults, &errors.Error{Code: errors.INVALID, Message: "Could not load the .env file", Operation: op, Err: err}
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
		CacheDriver:     os.Getenv("CACHE_DRIVER"),
		RedisAddress:    os.Getenv("REDIS_ADDRESS"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
		RedisDb:         os.Getenv("REDIS_DB"),
		MemCachedHosts:  os.Getenv("MEMCACHED_HOSTS"),
		AWSAccessKey:    os.Getenv("STORAGE_AWS_ACCESS_KEY"),
		AWSSecret:       os.Getenv("STORAGE_AWS_SECRET"),
		GCPJson:         os.Getenv("STORAGE_GCP_JSON"),
		GCPProjectID:    os.Getenv("STORAGE_GCP_PROJECT_ID"),
		AzureAccount:    os.Getenv("STORAGE_AZURE_ACCOUNT"),
		AzureKey:        os.Getenv("STORAGE_AZURE_KEY"),
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

// Validate the environment file for missing keys, if
// there are no validation errors, nil will be
// returned.
func (e *Env) Validate() validation.Errors {
	err := validation.Validator().Struct(e)
	if err != nil {
		return validation.Process(err)
	}
	return nil
}

// write is the env write function.
var write = godotenv.Write

// Set accepts a key, value pair and writes to the .env
// file when installing.
func (e *Env) Set(key string, value interface{}) error {
	const op = "Env.Set"

	val, err := cast.ToStringE(value)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error casting value to string", Operation: op, Err: err}
	}

	path := filepath.Join(basePath, EnvExtension)
	env, err := godotenv.Read(path)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error reading env file with the path " + path, Operation: op, Err: err}
	}

	env[key] = val

	err = write(env, path)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error writing env file with the path " + path, Operation: op, Err: err}
	}

	return nil
}

// newTpl is an alias for template.New
var newTpl = template.New

// Install Parses the the env and creates a
// new .env file in the root directory.
func (e *Env) Install() error {
	const op = "Env.Install"

	e.AppPort = cast.ToString(DefaultPort)

	tp, err := newTpl("").Parse(envTpl)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error parsing env tpl", Operation: op, Err: err}
	}

	buf := bytes.Buffer{}
	err = tp.ExecuteTemplate(&buf, "", e)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing env tpl", Operation: op, Err: err}
	}

	err = ioutil.WriteFile(filepath.Join(basePath, EnvExtension), buf.Bytes(), os.ModePerm)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error writing env file", Operation: op, Err: err}
	}

	return nil
}

// Port returns the env port as an integer, if the strconv
// produced an error, the default port of 5000 will
// be used,
func (e *Env) Port() int {
	const op = "Env.Port"
	if e.AppPort == "" {
		return DefaultPort
	}
	port, err := strconv.Atoi(e.AppPort)
	if err != nil {
		log.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast app port to int using port 5000", Operation: op, Err: err}).Error()
		return DefaultPort
	}
	return port
}

// IsProduction returns true if the application is set to
// production.
func (e *Env) IsProduction() bool {
	return e.AppEnv == "production" || e.AppEnv == "prod"
}

// IsDebug returns true if the application is set to
// debug.
func (e *Env) IsDebug() bool {
	return e.AppDebug != "false"
}
