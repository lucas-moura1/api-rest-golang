package userController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	models "user-api/domain/model"
	uuc "user-api/usecase"

	"github.com/go-chi/chi/v5"
)

type userGetUsers struct {
    userCaseUse uuc.UserUseCase;
}

type UserController interface {
    GetUsers(w http.ResponseWriter, r *http.Request)
    GetUserById(w http.ResponseWriter, r *http.Request)
    CreateUser(w http.ResponseWriter, r *http.Request)
    UpdateUser(w http.ResponseWriter, r *http.Request)
    DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUserController(uuc uuc.UserUseCase) UserController {
    return &userGetUsers{uuc}
}

func(ugu *userGetUsers) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := ugu.userCaseUse.GetAll()

    if err != nil {
        log.Print("Deu ruim :", err)
    }

    responseJson, err := json.Marshal(users)

    if err != nil {
        log.Print("Error: ", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJson)
}

func(ugu *userGetUsers) GetUserById(w http.ResponseWriter, r *http.Request) {
    userId := chi.URLParam(r, "userId")
    if userId == "" {
        w.WriteHeader(409)
        w.Write([]byte("Insert userId"))
    }

    user, err := ugu.userCaseUse.GetById(userId)

    if err != nil {
        if strings.Contains(err.Error(), "no documents") {
            w.WriteHeader(404)
        }
        log.Print("Error: ", err)
        w.WriteHeader(400)
    }

    responseJson, err := json.Marshal(user)

    if err != nil {
        log.Print("Error: ", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJson)
}

func(ugu *userGetUsers) CreateUser(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Print(err)
        w.WriteHeader(409)
        w.Write([]byte( "Error to readind data"))
	}

    var newUser models.User
	json.Unmarshal(body, &newUser)

    user, err := ugu.userCaseUse.Create(newUser)
    if err != nil {
        log.Print("Error: ", err)
    }

    responseJson, err := json.Marshal(user)
    if err != nil {
        log.Print("Error: ", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJson)
}

func(ugu *userGetUsers) UpdateUser(w http.ResponseWriter, r *http.Request) {
    userId := chi.URLParam(r, "userId")
    if userId == "" {
        w.WriteHeader(409)
        w.Write([]byte("Insert userId"))
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Print(err)
        w.WriteHeader(409)
        w.Write([]byte( "Error to readind data"))
	}

    var updateUser models.User
	json.Unmarshal(body, &updateUser)

    result, err := ugu.userCaseUse.Update(userId, updateUser)
    if err != nil {
        log.Fatal("Error: ", err)
    }

    type updatedResponse struct{ UpdatedDocument int `json:"updatedDocument"`}
    responseJson, err := json.Marshal(updatedResponse{UpdatedDocument: result})
    if err != nil {
        log.Print("Error: ", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJson)
}

func(ugu *userGetUsers) DeleteUser(w http.ResponseWriter, r *http.Request) {
    userId := chi.URLParam(r, "userId")
    if userId == "" {
        w.WriteHeader(409)
        w.Write([]byte("Insert userId"))
    }

    result, err := ugu.userCaseUse.DeleteById(userId)

    if err != nil {
        log.Print("Error: ", err)
        w.WriteHeader(400)
    }

    type deletedResponse struct{ DeletedDocument int `json:"deletedDocument"`}
    responseJson, err := json.Marshal(deletedResponse{DeletedDocument: result})
    if err != nil {
        log.Print("Error: ", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJson)
}
