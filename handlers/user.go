package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthdate  string `json:"birthdate"`
	Gender     string `json:"gender"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Вставка пользователя в базу данных
		rows := db.QueryRow(`
            INSERT INTO users (first_name, second_name, birthdate, gender, biography, city, password)
            VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
        `, req.FirstName, req.SecondName, req.Birthdate, req.Gender, req.Biography, req.City, req.Password)

		var id int
		err := rows.Scan(&id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read newly created user's id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": id})
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Получение данных пользователя из базы данных
		var user struct {
			ID         string `json:"id"`
			FirstName  string `json:"first_name"`
			SecondName string `json:"second_name"`
			Birthdate  string `json:"birthdate"`
			Gender     string `json:"gender"`
			Biography  string `json:"biography"`
			City       string `json:"city"`
		}
		err := db.QueryRow(`
            SELECT id, first_name, second_name, birthdate, gender, biography, city
            FROM users
            WHERE id = $1
        `, userID).Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Birthdate, &user.Gender, &user.Biography, &user.City)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
