package entity

type Role struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	DefaultAttribute
}

func (r Role) TableName() string {
	return "roles"
}
