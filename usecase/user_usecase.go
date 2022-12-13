package userUseCase

import (
	"time"
	"user-api/adapter/repository"
	models "user-api/domain/model"
	pw "user-api/utils"
)


type userUseCase struct {
    userRepository repository.UserRepository
}

type UserUseCase interface {
    GetAll() ([]models.DBUserResponse, error)
    GetById(id string) (*models.DBUserResponse, error)
    Create(u models.User) (*models.DBUserResponse, error)
    Update(id string, u models.User) (int, error)
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
    user, err := uuc.userRepository.FindUserById(id)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func(uuc *userUseCase) Create(u models.User) (*models.DBUserResponse, error) {
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

func(uuc *userUseCase) Update(id string, u models.User) (int, error) {
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
    result, err := uuc.userRepository.Delete(id)

    if err != nil {
        return 0, err
    }

    return result, nil
}

