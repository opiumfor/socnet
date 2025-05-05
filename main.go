package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"socnet/handlers"
)

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "host=localhost port=5444 user=socnet password=socnet dbname=socnet sslmode=disable")
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
