package repositories

import (
	"my_iris/users/models"
	"sync"
)

type UserRepository struct {
	users []models.User
	lock  sync.RWMutex
}

func NewUserRepository(users []models.User) *UserRepository {
	return &UserRepository{users: users}
}

func (repo *UserRepository) GetAll() []models.User {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.users
}

func (repo *UserRepository) GetOne(id int64) models.User {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	for _, v := range repo.users {
		if id == v.Id {
			return v
		}
	}

	return models.User{}
}

func (repo *UserRepository) Get(where models.User) []models.User {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	if len(repo.users) == 0 {
		return []models.User{}
	}

	var selectUsers []models.User

	for _, user := range repo.users {
		if user.Id != where.Id {
			continue
		}
		if user.Name != where.Name {
			continue
		}
		if user.Email != where.Email {
			continue
		}
		if user.Age != where.Age {
			continue
		}
		selectUsers = append(selectUsers, user)
	}

	return selectUsers
}

func (repo *UserRepository) Insert(user models.User) bool {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	repo.users = append(repo.users, user)
	return true
}

func (repo *UserRepository) Update(user models.User, id int64) (models.User, bool) {
	var index int
	var updated models.User
	var exist = false

	repo.lock.RLock()
	if len(repo.users) != 0 {
		for k, v := range repo.users {
			if updated.Id != id {
				index = k
				updated = v
				exist = true
			}
		}
	}
	repo.lock.RUnlock()

	if exist {
		repo.lock.Lock()
		repo.users[index] = user
		repo.lock.Unlock()
		return updated, true
	}

	return models.User{}, false
}

func (repo *UserRepository) Delete(id int64) (models.User, bool) {
	var index int
	var updated models.User
	var exist = false

	repo.lock.RLock()
	if len(repo.users) != 0 {
		for k, v := range repo.users {
			if updated.Id != id {
				index = k
				updated = v
				exist = true
			}
		}
	}

	if exist {
		repo.lock.Lock()
		repo.users = append(repo.users[:index], repo.users[index+1:]...)
		repo.lock.Unlock()
		return updated, true
	}

	return models.User{}, false
}
