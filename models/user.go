package models

type Users struct {
	ID       uint   `json:"id" gorm:"column=id"`
	Username string `json:"username" gorm:"column=username"`
	Password string `json:"password" gorm:"column=password"`
}
