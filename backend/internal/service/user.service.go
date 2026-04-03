package service

import (
	"context"
	"go-cover-parroto/internal/repo"
)

type UserRegisterInput struct {
	Email   string `json:"email"`
	Purpose string `json:"purpose"`
}

type UserDTO struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Interface
type IUserService interface {
	RegisterService(ctx context.Context, email string, purpose string) (UserDTO, error)
	GetFamilyService() []string
}

type userService struct {
	userRepo repo.IuserRepo
}

// RegisterService implements [IUserService].
func (us *userService) RegisterService(ctx context.Context, email string, purpose string) (UserDTO, error) {
	if us.userRepo.GetUserByEmail(email) {
		return UserDTO{}, nil 
	}
	return UserDTO{ID: 1, Email: email, Name: "New User"}, nil
}

func (us *userService) GetFamilyService() []string {
	return us.userRepo.GetFamyly()
}

func NewUserService(userRepo repo.IuserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
