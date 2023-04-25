package services

import (
	"context"
	"errors"

	"github.com/farakh-shahid/CRUD-APP/models"
	"github.com/farakh-shahid/CRUD-APP/utils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserServiceInterface {
	return &UserService{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserService) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserService) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New(constants.NotFound)
	}
	return users, nil
}

func (u *UserService) UpdateUser(userID string, user *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": user}

	result, err := u.usercollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return errors.New(constants.NotFound)
	}

	return nil
}

func (u *UserService) DeleteUser(id *string) error {
	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	result, err := u.usercollection.DeleteOne(u.ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New(constants.NotFound)
	}
	return nil
}
