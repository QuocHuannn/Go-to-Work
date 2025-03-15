package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct holds all application configuration
type Config struct {
	App    AppSettings
	SMTP   SMTPConfig
	DB     DBConfig
	Redis  RedisConfig
	Server ServerConfig
	JWT    JWTConfig
}

// AppSettings holds application-specific settings
type AppSettings struct {
	Name         string
	Environment  string
	LogLevel     string
	IsProduction bool
	IsTest       bool
}

// SMTPConfig holds email configuration
type SMTPConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	FromName  string
	FromEmail string
}

// DBConfig holds database configuration
type DBConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port           int
	Mode           string
	ReadTimeout    int
	WriteTimeout   int
	RequestTimeout int
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret    string
	ExpiresIn int // hours
}

// Global configuration instance
var Cfg Config

// LoadConfig loads configuration from environment variables
// and initializes the global configuration
func LoadConfig() {
	loadEnvFiles()

	Cfg = Config{
		App: AppSettings{
			Name:         getEnv("APP_NAME", "CRM System"),
			Environment:  getEnv("APP_ENV", "development"),
			LogLevel:     getEnv("LOG_LEVEL", "info"),
			IsProduction: getEnv("APP_ENV", "development") == "production",
			IsTest:       getEnv("APP_ENV", "development") == "test",
		},
		SMTP: SMTPConfig{
			Host:      getEnv("SMTP_HOST", "smtp.gmail.com"),
			Port:      getEnv("SMTP_PORT", "587"),
			Username:  getEnv("SMTP_USERNAME", ""),
			Password:  getEnv("SMTP_PASSWORD", ""),
			FromName:  getEnv("SMTP_FROM_NAME", "CRM System"),
			FromEmail: getEnv("SMTP_FROM_EMAIL", ""),
		},
		DB: DBConfig{
			Host:         getEnv("DB_HOST", "localhost"),
			Port:         getEnvAsInt("DB_PORT", 3306),
			Username:     getEnv("DB_USERNAME", "root"),
			Password:     getEnv("DB_PASSWORD", ""),
			DBName:       getEnv("DB_NAME", "crm-database"),
			MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Server: ServerConfig{
			Port:           getEnvAsInt("PORT", 8080),
			Mode:           getEnv("GIN_MODE", "debug"),
			ReadTimeout:    getEnvAsInt("SERVER_READ_TIMEOUT", 10),
			WriteTimeout:   getEnvAsInt("SERVER_WRITE_TIMEOUT", 10),
			RequestTimeout: getEnvAsInt("SERVER_REQUEST_TIMEOUT", 10),
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "your-secret-key"),
			ExpiresIn: getEnvAsInt("JWT_EXPIRES_IN", 24),
		},
	}

	// If SMTP username is provided but FromEmail is not, use username as FromEmail
	if Cfg.SMTP.FromEmail == "" && Cfg.SMTP.Username != "" {
		Cfg.SMTP.FromEmail = Cfg.SMTP.Username
	}

	// Print configuration for debugging (excluding sensitive info)
	if !Cfg.App.IsProduction {
		logConfig()
	}
}

// loadEnvFiles loads environment variables from .env files
// It tries multiple locations to find the .env file
func loadEnvFiles() {
	// List of possible .env file locations
	envFiles := []string{
		".env",                                   // Current directory
		".env." + os.Getenv("APP_ENV"),           // Environment-specific .env
		"../.env",                                // Parent directory
		filepath.Join("..", ".env"),              // Parent directory (alternative path format)
		filepath.Join(os.Getenv("HOME"), ".env"), // Home directory
	}

	// Load the first .env file found
	for _, envFile := range envFiles {
		if _, err := os.Stat(envFile); err == nil {
			err := godotenv.Load(envFile)
			if err == nil {
				log.Printf("Loaded environment from %s", envFile)
				return
			}
		}
	}

	log.Println("No .env file found. Using default environment variables.")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt gets an environment variable as an integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid integer value for %s: %s. Using default: %d", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

// logConfig prints the loaded configuration (excluding sensitive information)
func logConfig() {
	fmt.Println("=== Application Configuration ===")
	fmt.Printf("App Name: %s\n", Cfg.App.Name)
	fmt.Printf("Environment: %s\n", Cfg.App.Environment)
	fmt.Printf("SMTP Host: %s:%s\n", Cfg.SMTP.Host, Cfg.SMTP.Port)
	fmt.Printf("SMTP Username: %s\n", Cfg.SMTP.Username)
	fmt.Printf("Database: %s@%s:%d/%s\n", Cfg.DB.Username, Cfg.DB.Host, Cfg.DB.Port, Cfg.DB.DBName)
	fmt.Printf("Redis: %s:%d\n", Cfg.Redis.Host, Cfg.Redis.Port)
	fmt.Printf("Server Port: %d\n", Cfg.Server.Port)
	fmt.Printf("Server Mode: %s\n", Cfg.Server.Mode)
	fmt.Println("===============================")
}
