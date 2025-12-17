// Пакет config отвечает за загрузку настроек из .env файла.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Файл .env не найден, используем системные переменные")
	}

	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", "orchestra"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func (c *DBConfig) DSN() string {
	dsn := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	args := []interface{}{c.Host, c.Port, c.User, c.Name, c.SSLMode}

	if c.Password != "" {
		dsn += " password=%s"
		args = append(args, c.Password)
	}

	return fmt.Sprintf(dsn, args...)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
