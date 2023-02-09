package design

import (
	"context"

	"github.com/Jocerdikiawann/vm-go/models/entity"
	"github.com/Jocerdikiawann/vm-go/models/request"
)

type UserRepository interface {
	CreateOne(ctx context.Context, user request.SignUpRequest) (entity.Users, error)
	FindOneByID(ctx context.Context, id uint) (entity.Users, error)
	FindOneByEmail(ctx context.Context, email string) (entity.Users, error)
	UpdateOne(ctx context.Context, user entity.Users) (entity.Users, error)
	DeleteOne(ctx context.Context, id uint) (bool, error)
}
