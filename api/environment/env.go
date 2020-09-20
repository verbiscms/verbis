package environment

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var Env Map

type Map struct {
	AppName 			string
	AppEnv				string
	AppDebug 			string
	AppUrl				string
	DbHost				string
	DbPort				string
	DbDatabase			string
	DbUser				string
	DbPassword			string
	RedisHost			string
	RedisPassword		string
	RedisPort			string
	SparkpostApiKey		string
	SparkpostUrl		string
	MailFromAddress		string
	MailFromName		string
}

// Populate environment, loads and validates the environment file.
func Load() error {

	var (
		basePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		envPath = ".env"
	)

	if _, err := os.Stat(basePath + "/.env"); err == nil {
		envPath = basePath + "/.env"
	}

	// Init ENV
	if err := godotenv.Overload(envPath); err != nil {
		return fmt.Errorf("cannot load environment file: %w", err)
	}

	Env = Map{
		AppName: 	os.Getenv("APP_NAME"),
		AppEnv:  	os.Getenv("APP_ENV"),
		AppDebug:  	os.Getenv("APP_DEBUG"),
		AppUrl:  	os.Getenv("APP_URL"),
		DbHost:  	os.Getenv("DB_HOST"),
		DbPort:  	os.Getenv("DB_PORT"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbUser: 	os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisPort: os.Getenv("REDIS_PORT"),
		SparkpostApiKey: os.Getenv("SPARKPOST_API_KEY"),
		SparkpostUrl: os.Getenv("SPARKPOST_URL"),
		MailFromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		MailFromName: os.Getenv("MAIL_FROM_NAME"),
	}

	if err := validate(&Env); err != nil {
		fmt.Errorf("error validating environment file: %w", err)
	}

	return nil
}


// Validate the environment file for missing keys
func validate(e *Map) error {

	// Cast struct to map
	var envMap map[string]interface{}
	inrec, _ := json.Marshal(e)
	json.Unmarshal(inrec, &envMap)

	for k, v := range envMap {
		if v == "" {
			return fmt.Errorf(k + " is missing from the .environment file")
		}
	}

	return nil
}

// Get the database connection string.
func ConnectString() string {
	return Env.DbUser + ":" + Env.DbPassword + "@tcp(" + Env.DbHost + ":" + Env.DbPort + ")/" + Env.DbDatabase + "?tls=false&parseTime=true&multiStatements=true"
}

// Is production
func IsProduction() bool {
	return Env.AppEnv == "production" || Env.AppEnv == "prod"
}

// Is in debug mode
func IsDebug() bool {
	return Env.AppDebug == "true"
}
