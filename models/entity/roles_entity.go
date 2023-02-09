package entity

type Roles struct {
	Id       uint   `gorm:"primaryKey;auto_increment"`
	RoleName string `gorm:"column:roleName"`
}
