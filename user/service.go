package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisteUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	CheckEmailInput(input CheckEmailInput) (bool, error)
	UploadAvatar(id int, filleLocation string) (User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) RegisteUser(input RegisterUserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err!= nil {
        return User{}, err
    }

	user := User{
		Name: input.Name,
		Email: input.Email,
		Occupation: input.Occupation,
        PasswordHash: string(passwordHash),
		Role: "user",
	}

	newUser, newErr := s.r.Save(user)
	if newErr!= nil {
        return User{}, newErr
    }

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	user, err := s.r.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err!= nil {
        return user, err
    }

	return user, nil
}

func (s *service) CheckEmailInput(input CheckEmailInput) (bool, error) {
	user, err := s.r.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) UploadAvatar(id int, fileLocation string) (User, error) {
	user, err := s.r.FindByID(id)
	if err != nil {
		return user, nil
	}

	user.AvatarFileName = fileLocation

	userUpdated, err := s.r.Update(user)
	if err != nil {
		return userUpdated, err
	}

	return userUpdated, nil
}