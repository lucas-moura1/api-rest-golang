package implementation

import (
	"context"
	c "user-api/adapter/controller"
	r "user-api/adapter/repository"

	"go.mongodb.org/mongo-driver/mongo"

	"user-api/infrastructure/datastore"
	uc "user-api/usecase"
)

func GetImplementation(ctx context.Context, dataBaseName string, dataBaseDomain string) c.UserController {

    var db *mongo.Client = datastore.ConnectDB(dataBaseDomain, dataBaseName)

    var collection = datastore.GetCollection(db, dataBaseName, "users")

    var userRepository = r.NewUserRepository(ctx, collection)

    var userUseCase = uc.NewUserUseCase(userRepository)

    var userController = c.NewUserController(userUseCase)

    return userController
}







