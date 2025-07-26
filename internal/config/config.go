package config

import (
	"log"
	"os"
	"strconv"
)

var cfg *Config

type Config struct {
	AccessSecret  string
	RefreshSecret string
	Port          int
	DBPath        string
}

func InitConfig() {
	cfg = &Config{
		AccessSecret:  GetEnvString("JWT_ACCESS_SECRET", ""),
		RefreshSecret: GetEnvString("JWT_REFRESH_SECRET", ""),
		Port:          GetEnvInt("PORT", 8080),
		DBPath:        GetEnvString("DB_PATH", "./data.db"),
	}
}

func Get() *Config {
	if cfg == nil {
		log.Fatal("Config not initialized. Call config.InitConfig() first")
	}
	return cfg
}

func GetEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
