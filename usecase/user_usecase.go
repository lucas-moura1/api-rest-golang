package userUseCase

import (
	"time"
	"user-api/adapter/repository"
	models "user-api/domain/model"
	validator "user-api/domain/validation"
	pw "user-api/utils"
)


type userUseCase struct {
    userRepository repository.UserRepository
}

type UserUseCase interface {
    GetAll() ([]models.DBUserResponse, error)
    GetById(id string) (*models.DBUserResponse, error)
    Create(u models.UserInput) (*models.DBUserResponse, error)
    Update(id string, u models.UserInput) (int, error)
    DeleteById(id string) (int, error)
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{ur}
}

func(uuc userUseCase) GetAll() ([]models.DBUserResponse, error) {
    users, err := uuc.userRepository.FindAll()

    if err != nil {
        return nil, err
    }

    return users, nil
}

func(uuc userUseCase) GetById(id string) (*models.DBUserResponse, error) {
    err := validator.ValidateUserId(id)
    if err != nil {
        return nil, err
    }

    user, err := uuc.userRepository.FindUserById(id)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func(uuc *userUseCase) Create(u models.UserInput) (*models.DBUserResponse, error) {
    err := validator.ValidateUser(u)
    if err != nil {
        return nil, err
    }

    user := models.UserModel{
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Name: u.Name,
        Email: u.Email,
    }

    hashedPassword, err := pw.HashPassword(u.Password)
    if err != nil {
        return nil, err
    }

    user.Password = hashedPassword
    newUser, err := uuc.userRepository.Save(user)

    if err != nil {
        return nil, err
    }

    return newUser, nil
}

func(uuc *userUseCase) Update(id string, u models.UserInput) (int, error) {
    err := validator.ValidateUserId(id)
    if err != nil {
        return 0, err
    }

    err = validator.ValidateUser(u)
    if err != nil {
        return 0, err
    }

    hashedPassword, err := pw.HashPassword(u.Password)
    if err != nil {
        return 0, err
    }

    updateUser := models.UserModel{
        Name: u.Name,
        Email: u.Email,
        Password: hashedPassword,
        UpdatedAt: time.Now(),
    }

    result, err := uuc.userRepository.Update(id, updateUser)

    if err != nil {
        return 0, err
    }

    return result, nil
}

func(uuc userUseCase) DeleteById(id string) (int, error) {
    err := validator.ValidateUserId(id)
    if err != nil {
        return 0, err
    }

    result, err := uuc.userRepository.Delete(id)

    if err != nil {
        return 0, err
    }

    return result, nil
}

