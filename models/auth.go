package models

// Auth des
type Auth struct {
    ID int `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

// CheckAuth des
func CheckAuth(username, password string) bool {
    var auth Auth
    db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
    if auth.ID > 0 {
        return true
    }

    return false
}