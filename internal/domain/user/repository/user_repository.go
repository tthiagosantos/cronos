package repository

import (
	"context"
	"cronos/internal/domain/user/entity"
	"cronos/internal/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	collection := database.GetMongoDB().Collection("users")
	ctx := context.Background()
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Save(user *entity.User) error {
	collection := database.GetMongoDB().Collection("users")

	ctx := context.Background()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(user)
		log.Println("Error:", err)
		return err
	}

	return nil
}
