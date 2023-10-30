package config

import (
	"fmt"
	"os"
	"strconv"
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
	Debug				bool
}

type RendererConfig struct {
	RendererType string
	ViewPath 	 string
}

type CosmosOptions struct {
	Config 
	SessionConfig
	DatabaseConfig
	RendererConfig
}

func Load() (*Config) {
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

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	return &Config{
		Debug: debug,
		Port: os.Getenv("PORT"),
		Renderer: os.Getenv("RENDERER"),
		DatabaseConfig: dbConfig,
		SessionConfig: sessionConfig,
	}
}

// why is this in the config package?
// I keep it here instead of database because it deals with os.Getenv and i want 
// everything consistant.
// if i remove this and switch it to use the generated config instead of the os.Getenv
// ill move it to another package

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

//build dsn version with config instead of os.getenv
// func BuildDSN(c *config.DatabaseConfig) string {
// 	var dsn string

// 	switch c.DATABASE_TYPE {
// 	case "postgres", "postgresql":
// 		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
// 			c.DATABASE_HOST,
// 			c.DATABASE_PORT,
// 			c.DATABASE_USER,
// 			c.DATABASE_NAME, 
// 			c.DATABASE_SSL_MODE, 
// 		)

// 		if c.DATABASE_PASS != "" {
// 			dsn = fmt.Sprintf("%s password=%s", dsn, c.DATABASE_PASS)
// 		}

// 	default:
// 	}

// 	return dsn
// }


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
