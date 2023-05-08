package service

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *registerService) Save(c context.Context, payload *model.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(payload.Password)
	payload.Password = hashedPassword

	if err != nil {
		log.Fatal(err.Error())
	}

	newUser := model.User{
		ID:       primitive.NewObjectID(),
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	}

	return s.userRepository.Save(ctx, &newUser)
}

func (s *registerService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.GetByEmail(ctx, email)
}
