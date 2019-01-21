package services

import "my_iris/users/models"

type UserService interface {

	/**
	 * 获取所有用户
	 */
	GetAll() []models.User

	/**
	 * 获取一个用户
	 */
	GetOne(id int64) (models.User, error)

	/**
	 * 获取符合条件的用户
	 */
	Get(where models.User) []models.User

	/**
	 * 获取所有用户
	 */
	Insert(user models.User) (models.User, error)

	/**
	 * 更新用户信息
	 */
	Update(user models.User, id int64) (models.User, error)

	/**
	 * 删除用户
	 */
	Delete(id int64) (models.User, error)
}
