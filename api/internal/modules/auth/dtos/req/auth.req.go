package req

type UpdateProfileReq struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}
