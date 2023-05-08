package service

import (
	"context"
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
)

type registerService struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewRegisterService(userRepository model.UserRepository, timeout time.Duration) model.RegisterService {
	return &registerService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (s *registerService) Save(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.Save(ctx, user)
}

func (s *registerService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.GetByEmail(ctx, email)
}
