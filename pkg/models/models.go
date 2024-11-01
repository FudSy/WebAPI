package models

type Users struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role string `json:"role"`
}