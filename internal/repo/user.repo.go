package repo

type IuserRepo interface {
	GetUserByEmail(email string) bool
	GetInfoUser() string
	GetFamyly() []string // Keeping the typo for compatibility with existing service calls
}

type userRepo struct {
}

// GetUserByEmail implements [IuserRepo].
func (u *userRepo) GetUserByEmail(email string) bool {
	return true
}

func (u *userRepo) GetInfoUser() string {
	return "Finn"
}

func (u *userRepo) GetFamyly() []string {
	return []string{"Father", "mother", "brother"}
}

func NewUserRepo() IuserRepo {
	return &userRepo{}
}
