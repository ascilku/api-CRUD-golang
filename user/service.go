package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateService(creatUser CreateUser) (User, error)
	AuthUser(auth AuthUser) (User, error)
	CheckByEmail(checkEmail CheckEmailUser) (bool, error)
	UpdateImage(userID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateService(creatUser CreateUser) (User, error) {

	findByEmailRepository, err := s.repository.FindByEmailRepository(creatUser.Email)
	if err != nil {
		return findByEmailRepository, err
	} else {
		if findByEmailRepository.ID != 0 {
			return findByEmailRepository, errors.New("data sudah ada")
		} else {
			var keyUser User
			keyUser.Name = creatUser.Name
			keyUser.Email = creatUser.Email
			pass, err := bcrypt.GenerateFromPassword([]byte(creatUser.Password), bcrypt.MinCost)
			if err != nil {
				return keyUser, err
			} else {
				keyUser.Password = string(pass)
				createRepository, err := s.repository.CreateRepository(keyUser)
				if err != nil {
					return createRepository, err
				} else {
					return keyUser, nil
				}
			}
		}
	}
}

func (s *service) AuthUser(auth AuthUser) (User, error) {
	email := auth.Email
	user, err := s.repository.FindByEmailRepository(email)
	if err != nil {
		return User{}, err
	}

	if user.ID == 0 {
		return User{}, errors.New("tidak ada data")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		return User{}, errors.New("kata sandi tidak cocok")
	}

	return user, nil
}

func (s *service) CheckByEmail(checkEmail CheckEmailUser) (bool, error) {
	findByEmail, err := s.repository.FindByEmailRepository(checkEmail.Email)
	if err != nil {
		return false, err
	} else {
		if findByEmail.ID == 0 {
			return false, nil
		} else {
			return true, nil
		}
	}
}

func (s *service) UpdateImage(userID int, fileLocation string) (User, error) {
	findByID, err := s.repository.FindByIDRepository(userID)
	if err != nil {
		return findByID, err
	} else {
		if findByID.ID == 0 {
			return findByID, errors.New("not empty user")
		} else {
			findByID.File = fileLocation
			UpdateImage, err := s.repository.UpdateImageRepository(findByID)
			if err != nil {
				return UpdateImage, err
			} else {
				return UpdateImage, nil
			}
		}
	}
}

func (s *service) GetUserByID(ID int) (User, error) {
	findByID, err := s.repository.FindByIDRepository(ID)
	if err != nil {
		return findByID, err
	}

	if findByID.ID == 0 {
		return findByID, errors.New("no user found on with that ID")
	}

	return findByID, nil
}
