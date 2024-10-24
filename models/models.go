package models



type User struct {
	ID uint64 `json:"id" gorm:"primaryKey"`
	LastName string `json:"lastName"`
	FirstName string `json:"firstName"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role string `json:"role"`
}
type UserData struct {
	LastName string `json:"lastName"`
	FirstName string `json:"firstName"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
