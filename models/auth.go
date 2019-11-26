package models

import (
	"golang.org/x/crypto/bcrypt"
)

// Auth des
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth des
func CheckAuth(username, password string) bool {
	var auth Auth

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// fmt.Println(string(hashedPassword))

	db.Select("id,password").Where(Auth{Username: username}).First(&auth)
	if bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)) == nil {
		return true
	}

	return false
}
