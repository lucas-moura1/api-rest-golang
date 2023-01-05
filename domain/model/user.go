package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    Name        string      `json:"name"`
    Email       string      `json:"email"`
    Password    string      `json:"password"`
}

type UserModel struct {
    Name        string      `json:"name" bson:"name"`
    Email       string      `json:"email" bson:"email"`
    Password    string      `json:"password" bson:"password"`
    CreatedAt   time.Time   `json:"created_at" bson:"created_at,omitempty"`
    UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at,omitempty"`
}

type DBUserResponse struct {
    Id          primitive.ObjectID  `json:"id" bson:"_id"`
    Name        string              `json:"name" bson:"name"`
    Email       string              `json:"email" bson:"email"`
    Password    string              `json:"password" bson:"password"`
    CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}

type UserInput struct {
    Name                string      `json:"name" validate:"required,min=4,max=15"`
    Email               string      `json:"email" validate:"required,email"`
    Password            string      `json:"password" validate:"required,min=6,containsany=!@#?*"`
    PasswordConfirm     string      `json:"passwordConfirm" validate:"eqfield=Password"`
}

