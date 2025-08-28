package entity

type UserRole struct {
	UserID uint64 `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	RoleID uint64 `json:"role_id" gorm:"column:role_id"`
	Role   Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	DefaultAttribute
}

func (u UserRole) TableName() string {
	return "user_roles"
}
