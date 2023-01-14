package entity

type Roles struct {
	Id       uint `gorm:"primaryKey;<-:false"`
	RoleName string
}
