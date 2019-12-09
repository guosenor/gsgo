package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Auth des
type Auth struct {
	Model
	Username string `gorm:"type:varchar(45);unique_index"json:"username"`
	Password string `gorm:"type:varchar(64);"json:"password"`
}

// CheckAuth des
func CheckAuth(username, password string) bool {
	var auth Auth

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// fmt.Println(string(hashedPassword))

	db.Select("id,password").Where(Auth{Username: username}).First(&auth)
	if bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)) == nil {
		fmt.Println("CheckAuth:true")
		return true
	}
	fmt.Println("CheckAuth:false")

	return false
}
