package models

type UserRole struct {
	Id       int32  `gorm:"primary_key" json:"id"` // equals user id
	RoleName string `json:"role_name"`
	RoleType string `json:"role_type"`
}

func (ur *UserRole) TableName() string {
	return "t_user_role"
}