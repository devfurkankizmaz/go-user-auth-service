package model

import "context"

type RegisterRequest struct {
	Name            string `json:"name" binding:"required,min=2,max=30"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

type RegisterService interface {
	Save(c context.Context, payload *RegisterRequest) error
	GetUserByEmail(c context.Context, email string) (User, error)
}
