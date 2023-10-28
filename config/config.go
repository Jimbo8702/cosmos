package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DATABASE_TYPE 		string 
	DATABASE_HOST 		string
	DATABASE_PORT 		string
	DATABASE_USER 		string 
	DATABASE_PASS 		string
	DATABASE_NAME 		string 
	DATABASE_SSL_MODE 	string
}

type SessionConfig struct {
	SessionType 		string 
	CookieName 	 		string
	CookieLifetime 		string 
	CookiePersist  		string
	CookieSecure 		string 
	CookieDomain 		string 
}

//change from viper to standard gotenv
type Config struct {
	DatabaseConfig
	SessionConfig
	Renderer 			string
	AppName 			string
	ServerAddress 		string 
	Port 				string 
	RootPath 			string
}

func BuildConfig() (*Config) {
	dbConfig := DatabaseConfig{
		DATABASE_TYPE: 	os.Getenv("DATABASE_TYPE"),
		DATABASE_HOST:	os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:	os.Getenv("DATABASE_PORT"),
		DATABASE_USER: 	os.Getenv("DATABASE_USER"),
		DATABASE_PASS:	os.Getenv("DATABASE_PASS"),
		DATABASE_NAME: 	os.Getenv("DATABASE_NAME"),
		DATABASE_SSL_MODE: os.Getenv("DATABASE_SSL_MODE"),
	}
	sessionConfig := SessionConfig {
		SessionType: os.Getenv("SESSION_TYPE"),
		CookieName: os.Getenv("COOKIE_NAME"),
		CookieLifetime:  os.Getenv("COOKIE_LIFETIME"),
		CookiePersist: os.Getenv("COOKIE_PERSISTS"),
		CookieSecure: os.Getenv("COOKIE_SECURE"),
		CookieDomain: os.Getenv("COOKIE_DOMAIN"),
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Renderer: os.Getenv("RENDERER"),
		DatabaseConfig: dbConfig,
		SessionConfig: sessionConfig,
	}
}

func BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"), 
			os.Getenv("DATABASE_SSL_MODE"), 
		)

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}

	default:
	}

	return dsn
}

//loadconfig reads config from file or env variables
// func LoadConfig() (*Config, error) {
// 	var config *Config
// 	path, err := os.Getwd()
// 	if err != nil {
// 		return nil, err
// 	}
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("env")
// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	config.RootPath = path
// 	return config, nil
// }