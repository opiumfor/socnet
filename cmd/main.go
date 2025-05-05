package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"socnet/handlers"
)

func main() {
	// Чтение переменных окружения для подключения к базе данных
	dbHost := getEnv("DB_HOST", "postgres")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "socnet")
	dbPassword := getEnv("DB_PASSWORD", "socnet")
	dbName := getEnv("DB_NAME", "socnet")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	// Формирование строки подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Подключение к базе данных
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Инициализация маршрутов
	r := gin.Default()

	// Middleware для обработки CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Группы маршрутов
	authGroup := r.Group("/login")
	{
		authGroup.POST("", handlers.Login(db))
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handlers.Register(db))
		userGroup.GET("/get/:id", handlers.GetUser(db))
	}

	// Запуск сервера
	r.Run(":8787")
}

// Вспомогательная функция для получения переменных окружения с fallback-значением
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
