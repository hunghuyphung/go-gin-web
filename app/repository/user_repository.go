package repository

import (
	"go-gin-web/app/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Save(user *entity.User) (entity.User, error)
	DeleteById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil
	}
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindById(id int) (entity.User, error) {
	var user entity.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u UserRepositoryImpl) Save(user *entity.User) (entity.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}

func (u UserRepositoryImpl) DeleteById(id int) error {
	err := u.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
