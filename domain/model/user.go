package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    Name        string      `json:"name" bson:"name"`
    Email       string      `json:"email" bson:"email"`
    Password    string      `json:"password" bson:"password"`
}

type UserModel struct {
    Name        string      `json:"name" bson:"name"`
    Email       string      `json:"email" bson:"email"`
    Password    string      `json:"password" bson:"password"`
    CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}

type DBUserResponse struct {
    Id          primitive.ObjectID  `json:"id" bson:"_id"`
    Name        string      `json:"name" bson:"name"`
    Email       string      `json:"email" bson:"email"`
    Password    string      `json:"password" bson:"password"`
    CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}

type UserInput struct {
    User
    PasswordConfirm    string    `json:"password"`
}

