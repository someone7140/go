package authModel

import "time"

// AuthUserEntity 認証用ユーザのEntity
type AuthUserEntity struct {
	ID            string               `json:"_id" bson:"_id"`
	Email         string               `json:"email" bson:"email"`
	Password      string               `json:"password" bson:"password"`
	FacebookID    string               `json:"facebook_id" bson:"facebook_id"`
	GoogleID      string               `json:"google_id" bson:"google_id"`
	Status        string               `json:"status" bson:"status"`
	UserType      string               `json:"user_type" bson:"user_type"`
	UserSettingID string               `json:"user_setting_id" bson:"user_setting_id"`
	Name          string               `json:"name" bson:"name"`
	Categories    []string             `json:"categories" bson:"categories"`
	IconURL       string               `json:"icon_url" bson:"icon_url"`
	LoginCount    int64                `json:"login_count" bson:"login_count"`
	AuthMethod    string               `json:"auth_method" bson:"auth_method"`
	EmailAuth     *EmailAuthEntity     `json:"email_auth" bson:"email_auth"`
	PasswordReset *PasswordResetEntity `json:"password_reset" bson:"password_reset"`
}

// EmailAuthEntity Eメール認証用のオブジェクト
type EmailAuthEntity struct {
	Email      string    `json:"email" bson:"email"`
	Token      string    `json:"token" bson:"token"`
	ExpireDate time.Time `json:"expire_date" bson:"expire_date"`
}

// PasswordResetEntity パスワードリセット登録のオブジェクト
type PasswordResetEntity struct {
	Token      string    `json:"token" bson:"token"`
	ExpireDate time.Time `json:"expire_date" bson:"expire_date"`
}

// AuthUserResponse 認証後に返すユーザのレスポンス
type AuthUserResponse struct {
	ID            string   `json:"_id" bson:"_id"`
	UserType      string   `json:"user_type" bson:"user_type"`
	UserSettingID string   `json:"user_setting_id" bson:"user_setting_id"`
	Name          string   `json:"name" bson:"name"`
	Categories    []string `json:"categories" bson:"categories"`
	IconURL       string   `json:"icon_url" bson:"icon_url"`
	AuthMethod    string   `json:"auth_method" bson:"auth_method"`
	Token         string   `json:"token" bson:"token"`
	Exp           int64    `json:"exp" bson:"exp"`
}

// GoogleAuthRequest Google認証のリクエスト
type GoogleAuthRequest struct {
	GoogleIDToken string `json:"google_id_token" bson:"google_id_token"`
}

// FacebookAuthRequest Facebook認証のリクエスト
type FacebookAuthRequest struct {
	FacebookAccessToken string `json:"facebook_access_token" bson:"facebook_access_token"`
}

// EmailAuthRequest メール認証のリクエスト
type EmailAuthRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// EmailAuthTokenRequest メール認証のトークン認証リクエスト
type EmailAuthTokenRequest struct {
	UserID   string `json:"user_id" bson:"user_id"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}

// PasswordChangeRequest パスワード変更のリクエスト
type PasswordChangeRequest struct {
	Password string `json:"password" bson:"password"`
}

// PasswordResetRegisterRequest パスワードリセット登録のリクエスト
type PasswordResetRegisterRequest struct {
	Email string `json:"email" bson:"email"`
}

// PasswordResetUpdateRequest パスワードリセットによろ更新のリクエスト
type PasswordResetUpdateRequest struct {
	UserID   string `json:"user_id" bson:"user_id"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}
