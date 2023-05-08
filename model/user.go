package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserRepository interface {
	Save(c context.Context, user *User) error
	GetByEmail(c context.Context, id string) (User, error)
}
