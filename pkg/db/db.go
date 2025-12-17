// Пакет db предоставляет пул подключений к базе данных PostgreSQL.
package db

import (
	"database/sql"
	"fmt"
	"log"

	// Драйвер PostgreSQL. Пустой импорт (_), потому что он регистрирует себя в database/sql.
	_ "github.com/lib/pq"
)

// Connect создаёт и возвращает пул подключений к PostgreSQL.
// dsn — строка подключения (Data Source Name).
func Connect(dsn string) *sql.DB {
	// Открываем подключение к PostgreSQL (драйвер = "postgres")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Не удалось подготовить подключение к БД: %v", err)
	}

	// Проверяем реальное подключение
	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Настраиваем пул соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	fmt.Println("✅ Успешное подключение к базе данных")
	return db
}
