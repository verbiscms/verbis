package environment

import (
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	pkgValidate "github.com/go-playground/validator/v10"
	"strconv"

	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var env envMap

type envMap struct {
	AppName 			string `json:"APP_NAME"`
	AppEnv				string `json:"APP_ENV"`
	AppDebug 			string `json:"APP_DEBUG"`
	AppUrl				string `json:"APP_URL" binding:"required"`
	AppPort				string `json:"APP_PORT" binding:"required"`
	DbHost				string `json:"DB_HOST" binding:"required,ip"`
	DbPort				string `json:"DB_PORT" binding:"required"`
	DbDatabase			string `json:"DB_DATABASE" binding:"required"`
	DbUser				string `json:"DB_USERNAME" binding:"required"`
	DbPassword			string `json:"DB_PASSWORD" binding:"required"`
	SparkpostApiKey		string `json:"SPARKPOST_API_KEY"`
	SparkpostUrl		string `json:"SPARKPOST_URL"`
	MailFromAddress		string `json:"MAIL_FROM_ADDRESS"`
	MailFromName		string `json:"MAIL_FROM_NAME"`
}

type Mail struct {
	SparkpostApiKey		string `json:"SPARKPOST_API_KEY"`
	SparkpostUrl		string `json:"SPARKPOST_URL"`
	FromAddress			string `json:"MAIL_FROM_ADDRESS"`
	FromName			string `json:"MAIL_FROM_NAME"`
}

// Load populates environment, loads and validates the environment file.
// Returns errors.INVALID if the env file failed to load.
func Load() error {
	const op = "environment.Load"

	var (
		basePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		envPath = ".env"
	)

	if _, err := os.Stat(basePath + "/.env"); err == nil {
		envPath = basePath + "/.env"
	}

	if err := godotenv.Overload(envPath); err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Could not load the enviromnent file, is there a .env file in the root of the verbis project?", Operation: op, Err: err}
	}

	env = envMap{
		AppEnv:  	os.Getenv("APP_ENV"),
		AppDebug:  	os.Getenv("APP_DEBUG"),
		AppUrl:  	os.Getenv("APP_URL"),
		AppPort:  	os.Getenv("APP_PORT"),
		DbHost:  	os.Getenv("DB_HOST"),
		DbPort:  	os.Getenv("DB_PORT"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbUser: 	os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		SparkpostApiKey: os.Getenv("SPARKPOST_API_KEY"),
		SparkpostUrl: os.Getenv("SPARKPOST_URL"),
		MailFromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		MailFromName: os.Getenv("MAIL_FROM_NAME"),
	}

	return nil
}

// Validate the environment file for missing keys
func Validate() []validation.ValidationError {
	v := validation.New()
	err := v.Package.Struct(env)
	if err != nil {
		validationErrors := err.(pkgValidate.ValidationErrors)
		return v.Process(validationErrors)
	}
	return nil
}

// App - GetAppName
func GetAppName() string {
	return env.AppName
}

// App - GetAppEv
func GetAppEnv() string {
	return env.AppEnv
}

// Database - GetPort
func GetPort() int {
	n, _ := strconv.Atoi(env.AppPort)
	return n
}

// Database - ConnectString
func ConnectString() string {
	return env.DbUser + ":" + env.DbPassword + "@tcp(" + env.DbHost + ":" + env.DbPort + ")/" + env.DbDatabase + "?tls=false&parseTime=true&multiStatements=true"
}

// Database - GetDatabaseName
func GetDatabaseName() string {
	return env.DbDatabase
}

// Mail - GetMailConfiguration
func GetMailConfiguration() Mail {
	return Mail{
		FromAddress: env.MailFromAddress,
		FromName: env.MailFromName,
	}
}

// Env - IsProduction
func IsProduction() bool {
	return env.AppEnv == "production" || env.AppEnv == "prod"
}

// Env - IsDevelopment
func IsDevelopment() bool {
	return env.AppEnv != "production" && env.AppEnv != "prod"
}

// Env - IsDebug
func IsDebug() bool {
	return env.AppDebug != "false"
}


