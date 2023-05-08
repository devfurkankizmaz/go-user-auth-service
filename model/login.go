package model

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginService interface {
	GetUserByEmail(c context.Context, email string) (User, error)
}
