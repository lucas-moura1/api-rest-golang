package repository

import (
	"context"
	"log"
	models "user-api/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
    ctx     context.Context
    db      *mongo.Collection
}

type UserRepository interface {
    FindAll() ([]models.DBUserResponse, error)
    FindUserById(id string) (*models.DBUserResponse, error)
    Save(u models.UserModel) (*models.DBUserResponse, error)
    Update(id string, u models.UserModel) (int, error)
    Delete(id string) (int, error)
}

func NewUserRepository(ctx context.Context, db *mongo.Collection) UserRepository {
	return &userRepository{ctx, db}
}

func(ur *userRepository) FindAll() ([]models.DBUserResponse, error) {
    cursor, err := ur.db.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Print("Error: ", err)
	}

    var users []models.DBUserResponse
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Print("Error: ", err)
	}

    return users, nil
}

func(ur *userRepository) FindUserById(id string) (*models.DBUserResponse, error) {
    oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.DBUserResponse

	query := bson.M{"_id": oid}
	err := ur.db.FindOne(ur.ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func(ur *userRepository) Save(u models.UserModel) (*models.DBUserResponse, error) {

    res, err := ur.db.InsertOne(ur.ctx, u)
    if err != nil {
		log.Print("Error: ", err)
        return nil, err
	}

    var newUser *models.DBUserResponse
    query := bson.M{"_id": res.InsertedID}

    err = ur.db.FindOne(ur.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func(ur *userRepository) Update(id string, u models.UserModel) (int, error) {
    oid, _ := primitive.ObjectIDFromHex(id)

    query := bson.M{"_id": oid}
    update := bson.M{"$set": u}

    result, err := ur.db.UpdateOne(ur.ctx, query, update)
    if err != nil {
        log.Print("Error: ", err)
        return 0, nil
    }

    return int(result.MatchedCount), err
}

func(ur *userRepository) Delete(id string) (int, error) {
    oid, _ := primitive.ObjectIDFromHex(id)

    query := bson.M{"_id": oid}
	result, err := ur.db.DeleteOne(ur.ctx, query)
    if err != nil {
        return 0, err
    }

    return int(result.DeletedCount), nil
}
