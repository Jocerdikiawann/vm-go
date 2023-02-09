package repository

import (
	"context"

	"github.com/Jocerdikiawann/vm-go/models/entity"
	"github.com/Jocerdikiawann/vm-go/models/request"
	"github.com/Jocerdikiawann/vm-go/repository/design"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) design.UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (repository *UserRepositoryImpl) CreateOne(ctx context.Context, user request.SignUpRequest) (entity.Users, error) {
	request := user.ToEntity()
	result := repository.Db.WithContext(ctx).
		Create(&request)

	return request, result.Error
}

func (repository *UserRepositoryImpl) FindOneByID(ctx context.Context, id uint) (entity.Users, error) {
	var user entity.Users
	repository.Db.WithContext(ctx).
		Where("id=?", id).
		Preload("Roles").
		First(&user)
	return user, nil
}

func (repository *UserRepositoryImpl) FindOneByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users
	repository.Db.WithContext(ctx).
		Where("email=?", email).
		Preload("Roles").
		First(&user)
	return user, nil
}

func (repository *UserRepositoryImpl) UpdateOne(ctx context.Context, user entity.Users) (entity.Users, error) {
	return entity.Users{}, nil
}

func (repository *UserRepositoryImpl) DeleteOne(ctx context.Context, id uint) (bool, error) {
	return false, nil
}
