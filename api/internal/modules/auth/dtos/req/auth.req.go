package req

type SyncReq struct {
	FirebaseToken string `json:"firebase_token" binding:"required"`
}
