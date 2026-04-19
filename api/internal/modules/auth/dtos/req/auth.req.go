package req

type SyncReq struct {
	FirebaseToken string `json:"firebase_token" binding:"required"`
}

type GetTokenReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
