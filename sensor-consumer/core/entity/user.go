package entity

type User struct {
	DefaultAttribute
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
