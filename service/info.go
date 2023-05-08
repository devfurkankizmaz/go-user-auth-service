package service

import (
	"context"
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
)

type infoService struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewInfoService(userRepository model.UserRepository, timeout time.Duration) model.InfoService {
	return &infoService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (s *infoService) GetInfoByID(c context.Context, userID string) (*model.Info, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	user, err := s.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.Info{Name: user.Name}, nil
}
