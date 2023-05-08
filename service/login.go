package service

import (
	"context"
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
)

type loginService struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewLoginService(userRepository model.UserRepository, timeout time.Duration) model.LoginService {
	return &loginService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}
