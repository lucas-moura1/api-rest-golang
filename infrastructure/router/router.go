package router

import (
	userController "user-api/adapter/controller"

	"github.com/go-chi/chi/v5"
)

func NewRouter(uc userController.UserController) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/user", uc.GetUsers)
    r.Get("/user/{userId}", uc.GetUserById)
    r.Post("/user", uc.CreateUser)
    r.Patch("/user/{userId}", uc.UpdateUser)
    r.Delete("/user/{userId}", uc.DeleteUser)

	return r
}
