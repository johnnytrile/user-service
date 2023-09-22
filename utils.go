package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func generateToken(username string) string {
	// Ở đây, bạn nên sử dụng một secret key riêng để ký và xác thực JWT.
	// Đảm bảo giữ bí mật và không chia sẻ secret key này với bất kỳ ai.
	// Hãy lưu secret key trong biến môi trường hoặc cách an toàn hơn.
	secretKey := []byte("secret_key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(fmt.Sprintf("failed to generate token: %s", err.Error()))
	}

	return tokenString
}
