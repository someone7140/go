package authModel

// EmailLoginRequest メールログインのリクエスト
type EmailLoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
