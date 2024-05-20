package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"system_employee/config"
	"system_employee/model/response"
	"time"
)

func GetDataRole(email string) (*response.RoleResponse, error) {
	db := config.NewDB()
	defer db.Close() // Make sure to close the DB connection

	fmt.Println("params : ", email)
	sql := `SELECT role, email FROM company_employee.user WHERE email = ?`

	rows, err := db.Query(sql, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resultEmail := &response.RoleResponse{}
	if rows.Next() {
		err := rows.Scan(&resultEmail.RoleUser, &resultEmail.Email)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("data register not found")
	}

	return resultEmail, nil
}

func GenerateNewAccessToken(email string) (string, error) {
	configuration := config.New()

	// Set secret key from .env file.
	secret := configuration.Get("JWT_SECRET_KEY")
	fmt.Println("secret : ", secret)
	// Set expires hours count for secret key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRE_HOUR_COUNT"))
	fmt.Println("hoursCount : ", hoursCount)
	// Create a new claims.
	claims := jwt.MapClaims{}
	result, err := GetDataRole(email)
	fmt.Println("result : ", result)
	// Set public claims:
	claims["email"] = email
	claims["role"] = result.RoleUser
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
