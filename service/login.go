package service

import (
	"context"
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/utils"
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

func (ls *loginService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, ls.contextTimeout)
	defer cancel()
	return ls.userRepository.GetByEmail(ctx, email)
}

func (ls *loginService) GenAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return utils.GenAccessToken(user, secret, expiry)
}

func (ls *loginService) GenRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.GenRefreshToken(user, secret, expiry)
}
