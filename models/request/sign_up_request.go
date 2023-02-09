package request

import "github.com/Jocerdikiawann/vm-go/models/entity"

type SignUpRequest struct {
	Id        uint
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt uint64
	UpdatedAt uint64
}

func (signUp *SignUpRequest) ToEntity() entity.Users {
	return entity.Users{
		Id:        signUp.Id,
		Email:     signUp.Email,
		Password:  signUp.Password,
		FirstName: signUp.FirstName,
		LastName:  signUp.LastName,
		CreatedAt: signUp.CreatedAt,
		UpdatedAt: signUp.UpdatedAt,
	}
}
