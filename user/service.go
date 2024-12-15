package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisteUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
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