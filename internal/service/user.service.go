package service

import "go-familytree/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService{
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoService() string {
	return us.userRepo.GetInfoUser()
}

func (us *UserService) GetFamilyService() []string {
	return us.userRepo.GetFamyly()
}