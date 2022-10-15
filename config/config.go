package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName                  string
	AppPort                  string
	LogLevel                 string
	Environment              string
	JWTSecret                string
	RedisAddress             string
	DBUsername               string
	DBPassword               string
	DBHost                   string
	DBPort                   int
	DBName                   string
	MinioEndpoint            string
	MinioAccessKey           string
	MinioSecretKey           string
	MinioRegion              string
	MinioBucket              string
	DBMaxConnections         int
	DBMaxIdleConnections     int
	DBMaxLifetimeConnections int
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Warning("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetString("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		// JWTSecret:      GetString("JWT_SECRET"),
		// RedisAddress:   GetString("REDIS_ADDRESS"),
		DBUsername:               GetString("DB_USERNAME"),
		DBPassword:               GetString("DB_PASSWORD"),
		DBHost:                   GetString("DB_HOST"),
		DBPort:                   GetInt("DB_PORT"),
		DBName:                   GetString("DB_NAME"),
		DBMaxConnections:         GetInt("DB_MAX_CONNECTIONS"),
		DBMaxIdleConnections:     GetInt("DB_MAX_IDLE_CONNECTIONS"),
		DBMaxLifetimeConnections: GetInt("DB_MAX_LIFETIME_CONNECTIONS"),
		// MinioEndpoint:  GetString("MINIO_ENDPOINT"),
		// MinioAccessKey: GetString("MINIO_ACCESS_KEY"),
		// MinioSecretKey: GetString("MINIO_SECRET_KEY"),
		// MinioRegion:    GetString("MINIO_REGION"),
		// MinioBucket:    GetString("MINIO_BUCKET"),
	}

	return appConfig
}
