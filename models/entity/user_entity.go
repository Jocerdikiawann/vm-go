package entity

type User struct {
	Id uint `gorm:"primaryKey;<-:false"`
	Email string
	Password string
	firstName string
	lastName string
	roles []Roles `gorm:"many2many:users_roles"`
	createdAt uint64
	updatedAt uint64
}