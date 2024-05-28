package user

import "gorm.io/gorm"

type Repository interface {
	CreateRepository(user User) (User, error)
	FindByEmailRepository(email string) (User, error)
	FindByIDRepository(ID int) (User, error)
	UpdateImageRepository(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateRepository(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (r *repository) FindByEmailRepository(email string) (User, error) {
	var keyUser User
	err := r.db.Where("email = ?", email).Find(&keyUser).Error
	if err != nil {
		return keyUser, err
	} else {
		return keyUser, nil
	}
}

func (r *repository) FindByIDRepository(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateImageRepository(user User) (User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
