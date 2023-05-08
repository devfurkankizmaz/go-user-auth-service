package repository

import (
	"context"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	db         mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) model.UserRepository {
	return &userRepository{
		db:         db,
		collection: collection,
	}
}

func (r *userRepository) Save(c context.Context, user *model.User) error {
	collection := r.db.Collection(r.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (r *userRepository) GetByID(c context.Context, id string) (model.User, error) {
	collection := r.db.Collection(r.collection)

	var user model.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (r *userRepository) GetByEmail(c context.Context, email string) (model.User, error) {
	collection := r.db.Collection(r.collection)
	var user model.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}
