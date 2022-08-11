package auth

import (
	errorConstants "coda-api/src/constants"
	authRepository "coda-api/src/repository/auth"
	userRepository "coda-api/src/repository/user"
	"coda-api/src/util"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthEmailRegister メール認証登録
func AuthEmailRegister(email string, password string) error {
	if email == "" || password == "" {
		return errorConstants.ErrBadRequest
	}
	// emailでユーザ取得
	userEntity, _ := userRepository.GetUserEntityByEmail(email)
	if userEntity != nil && userEntity.Status != "confirming" {
		return errorConstants.ErrForbidden
	}
	// 確認中でユーザが取得できた場合は削除
	if userEntity != nil {
		err := userRepository.DeleteUser(userEntity.ID)
		if err != nil {
			return err
		}
	}
	// トークンとユーザIDを発行して仮登録
	token, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	userID, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	err = authRepository.EmailAuthRegister(userID, email, string(hashedPw), token)
	if err != nil {
		return err
	}
	// メール送信
	url := os.Getenv("VIEW_DOMAIN") + "/auth/emailAuth?user_id=" + userID + "&token=" + token
	plainTextContent := "こちらのURLより会員登録をしてください。有効期限は1日です。\r\n" + url
	htmlContent := "<a href='" + url + "'>こちら</a>" + "のURLより会員登録をしてください。有効期限は1日です。"
	return util.SendMailBySendGrid(email, "Coda会員登録", plainTextContent, htmlContent)
}

// AuthEmailToken メール認証トークンの認証
func AuthEmailToken(userID string, password string, token string) error {
	if userID == "" || password == "" || token == "" {
		return errorConstants.ErrBadRequest
	}
	// IDでユーザ取得
	userEntity, _ := userRepository.GetUserEntityByUserID(userID, "email")
	if userEntity == nil || userEntity.EmailAuth == nil {
		return errorConstants.ErrForbidden
	}
	// 有効期限
	now := time.Now().UTC()
	if now.After(userEntity.EmailAuth.ExpireDate) {
		return errorConstants.ErrForbidden
	}
	// tokenの比較
	if token != userEntity.EmailAuth.Token {
		return errorConstants.ErrUnauthorized
	}
	// passwordの比較
	err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password))
	if err != nil {
		return errorConstants.ErrUnauthorized
	}
	return nil
}

// ChangePassword パスワード変更
func ChangePassword(password string, userID string) error {
	if password == "" {
		return errorConstants.ErrBadRequest
	}
	// パスワードの生成
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	// passwordの更新
	err = authRepository.ChangePassword(string(hashedPw), userID)
	if err != nil {
		return errorConstants.ErrInternalServer
	}
	return nil
}

// RegisterPasswordReset パスワードリセット登録
func RegisterPasswordReset(email string) error {
	if email == "" {
		return errorConstants.ErrBadRequest
	}
	// メールアドレスでユーザ取得
	userEntity, err := userRepository.GetUserEntityByEmail(email)
	if err != nil {
		return err
	}
	if userEntity == nil || userEntity.Status != "active" {
		return errorConstants.ErrBadRequest
	}
	// パスワードリセット情報があれば削除
	if userEntity.PasswordReset != nil {
		authRepository.DeletePasswordReset(userEntity.ID)
	}
	token, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	// パスワードリセット情報の更新
	err = authRepository.AddPasswordReset(token, userEntity.ID)
	if err != nil {
		return err
	}
	// メール送信
	url := os.Getenv("VIEW_DOMAIN") + "/auth/passwordResetAuth?user_id=" + userEntity.ID + "&token=" + token
	plainTextContent := "こちらのURLよりパスワードリセットを行ってください。有効期限は1日です。\r\n" + url
	htmlContent := "<a href='" + url + "'>こちら</a>" + "のURLよりパスワードリセットを行ってください。有効期限は1日です。"
	return util.SendMailBySendGrid(email, "Codaパスワードリセット", plainTextContent, htmlContent)
}

// PasswordResetUpdate パスワードリセットによる更新
func PasswordResetUpdate(userID string, token string, password string) error {
	if userID == "" || token == "" || password == "" {
		return errorConstants.ErrBadRequest
	}
	// ユーザIDでユーザ取得
	userEntity, err := userRepository.GetUserEntityByUserID(userID, "email")
	if err != nil {
		return err
	}
	if userEntity == nil || userEntity.Status != "active" {
		return errorConstants.ErrBadRequest
	}
	// パスワードリセット情報のチェック
	if userEntity.PasswordReset == nil ||
		userEntity.PasswordReset.ExpireDate.Before(time.Now().UTC()) ||
		token != userEntity.PasswordReset.Token {
		return errorConstants.ErrForbidden
	}
	// パスワードリセット情報の削除
	authRepository.DeletePasswordReset(userEntity.ID)
	// パスワードの生成
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	// passwordの更新
	err = authRepository.ChangePassword(string(hashedPw), userID)
	if err != nil {
		return errorConstants.ErrInternalServer
	}
	return nil
}
