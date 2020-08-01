package sql

type User struct {
	// gorm.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
