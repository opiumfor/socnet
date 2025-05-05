package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

		// Хэширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Вставка пользователя в базу данных
		var userID int
		err = db.QueryRow(`
            INSERT INTO users (first_name, second_name, birthdate, gender, biography, city, password)
            VALUES ($1, $2, $3, $4, $5, $6, $7)
            RETURNING id
        `, req.FirstName, req.SecondName, req.Birthdate, req.Gender, req.Biography, req.City, string(hashedPassword)).Scan(&userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": userID})
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Получение данных пользователя из базы данных
		var user struct {
			ID         int    `json:"id"`
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

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
