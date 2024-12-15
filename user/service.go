package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisteUser(input RegisterUserInput) (User, error)
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