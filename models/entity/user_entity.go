package entity

type Users struct {
	Id        uint    `gorm:"primaryKey;auto_increment;column:id"`
	Email     string  `gorm:"column:email"`
	Password  string  `gorm:"column:password"`
	FirstName string  `gorm:"column:firstName"`
	LastName  string  `gorm:"column:lastName"`
	Roles     []Roles `gorm:"many2many:user_roles;"`
	CreatedAt uint64  `gorm:"column:createdAt"`
	UpdatedAt uint64  `gorm:"column:updatedAt"`
}

// func (u *Users) AfterCreate(tx *gorm.DB) (err error) {
// 	roles := []Roles{
// 		{1, "Staff"},
// 	}
// 	tx.Model(&u).Where("id IN ?", u.Id).Association("Roles").Append(roles)
// 	return
// }
