package userdashboard

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/gocrypto"
)

func CreateUser(db *sql.DB, email string, password string) (int, error) {
	// Hash password
	hash, err := gocrypto.HashPassword(password)
	if err != nil {
		return 0, err
	}

	// Insert user into database
	_, err = db.Exec(`
		INSERT INTO users (email, password) VALUES ($1, $2)
		RETURNING id;
	`, email, hash)
	if err != nil {
		return 0, err
	}

	// Get user ID
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		return 0, errors.New("failed to get user ID")
	}

	return userID, nil
}

func VerifyUserPassword(db *sql.DB, email string, password string) (bool, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		return false, errors.New("user not found")
	}

	// Check password hash
	match, err := gocrypto.ComparePassword(password, hashedPassword)
	if err != nil {
		return false, err
	}

	return match, nil
}

func GenerateJWT(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  "user-dashboard",
		"sub":  user.UserID,
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(token string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func CleanInput(input string) string {
	input = strings.TrimSpace(input)
	return strings.ToLower(input)
}

func CleanInt(input string) int {
	parsed, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return 0
	}

	return parsed
}

func CleanFloat(input string) float64 {
	parsed, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Println(err)
		return 0
	}

	return parsed
}

func CheckEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "", errors.New("environment variable not set")
	}

	return value, nil
}