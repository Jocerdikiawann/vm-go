package app

import (
	"fmt"

	"github.com/Jocerdikiawann/vm-go/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(host, port, dbname, username, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	db.AutoMigrate(&entity.Users{}, &entity.Roles{})
	return db, err
}
