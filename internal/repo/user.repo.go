package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo{
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "Finn"
}

func (ur *UserRepo) GetFamyly() []string {
	return []string{"Father", "mother", "brother"}
}