package models

type User struct {
	ID       int    `json:"id"`
	Name     string `gorm:"size:100" json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// Token    string `json:"token" form:"token"`
}
