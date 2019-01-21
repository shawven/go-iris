package services

import (
	"github.com/kataras/iris/core/errors"
	"my_iris/users/models"
	"my_iris/users/repositories"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (service *UserServiceImpl) GetAll() []models.User {
	return service.userRepository.GetAll()
}

func (service *UserServiceImpl) GetOne(id int64) (models.User, error) {
	if id < 0 {
		return models.User{}, errors.New("id 错误")
	}
	return service.userRepository.GetOne(id), nil
}

func (service *UserServiceImpl) Get(where models.User) []models.User {
	return service.userRepository.Get(where)
}

func (service *UserServiceImpl) Insert(user models.User) (models.User, error) {
	b := service.userRepository.Insert(user)
	if b {
		return user, nil
	}
	return user, errors.New("添加用户失败")
}

func (service *UserServiceImpl) Update(user models.User, id int64) (models.User, error) {
	updated, b := service.userRepository.Update(user, id)
	if b {
		return updated, nil
	}
	return updated, errors.New("更新用户失败")
}

func (service *UserServiceImpl) delete(id int64) (models.User, error) {
	deleted, b := service.userRepository.Delete(id)
	if b {
		return deleted, nil
	}
	return deleted, errors.New("删除用户失败")
}
