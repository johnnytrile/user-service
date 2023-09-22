package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Import utils.go

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "myuser"),
		getEnv("DB_PASSWORD", "mypassword"),
		getEnv("DB_NAME", "mydb"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	user = User{Username: username, Password: password}
	db.Create(&user)

	fmt.Fprintf(w, "Registered user: %s\n", username)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	db.First(&user, "username = ? AND password = ?", username, password)

	if user.ID == 0 {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Tạo JWT token
	token := generateToken(user.Username)

	fmt.Fprintf(w, "Logged in as: %s\nToken: %s\n", username, token)
}
