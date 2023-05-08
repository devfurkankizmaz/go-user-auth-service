package model

import "context"

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type RegisterService interface {
	Save(c context.Context, user *RegisterRequest) error
	GetUserByEmail(c context.Context, email string) (User, error)
}
