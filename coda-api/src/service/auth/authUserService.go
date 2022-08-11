package auth

import (
	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	authRepository "coda-api/src/repository/auth"
	userRepository "coda-api/src/repository/user"
	"coda-api/src/util"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/oauth2/v2"
)

var secretKey = "zPVT2TjgkKjfDJTz9P3CLY9p9HyZx8DrezQ4iyWdca3pC9XZrE6sR6f42m3Ef7x9uVCEwKV5nMFrBi7HENgYAt7"

// GetJwtTokenUser ユーザ認証用のJwtトークンを生成
func GetJwtTokenUser(auth authModel.AuthUserResponse) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"id":            auth.ID,
		"userType":      auth.UserType,
		"userSettingID": auth.UserSettingID,
		"name":          auth.Name,
		"categories":    strings.Join(auth.Categories, ","),
		"iconURL":       auth.IconURL,
		"authMethod":    auth.AuthMethod,
		"exp":           GetAuthExpire(),
	}
	return token.SignedString([]byte(secretKey))
}

// GetLoginInfoInfoFromJwtToken jwtトークンからログイン情報を取得
func GetLoginInfoInfoFromJwtToken(c *gin.Context) (*authModel.AuthUserResponse, error) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})
	if token == nil || err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	// expの判定
	expTime := time.Unix(int64(claims["exp"].(float64)), 0).UTC()
	if time.Now().UTC().After(expTime) {
		return nil, errorConstants.ErrForbidden
	}
	authInfo := authModel.AuthUserResponse{
		ID:            claims["id"].(string),
		UserType:      claims["userType"].(string),
		UserSettingID: claims["userSettingID"].(string),
		Name:          claims["name"].(string),
		Categories:    strings.Split(claims["categories"].(string), ","),
		IconURL:       claims["iconURL"].(string),
		AuthMethod:    claims["authMethod"].(string),
	}
	return &authInfo, nil
}

// SetLoginInfoJSONToContext ログイン情報のJSONをコンテキストにセット
func SetLoginInfoJSONToContext(c *gin.Context, loginInfo *authModel.AuthUserResponse) error {
	loginUserJSONByte, err := json.Marshal(loginInfo)
	if err != nil {
		return err
	}
	c.Set("loginUserJson", string(loginUserJSONByte))
	return nil
}

// GetLoginInfoFromContext Contextからログイン情報を取得
func GetLoginInfoFromContext(c *gin.Context) *authModel.AuthUserResponse {
	var loginInfo authModel.AuthUserResponse
	err := json.Unmarshal([]byte(c.GetString("loginUserJson")), &loginInfo)
	if err != nil {
		return nil
	}
	return &loginInfo
}

// GetAuthExpire 認証期限の取得
func GetAuthExpire() int64 {
	return time.Now().UTC().AddDate(0, 3, 0).Unix()
}

// AuthenticationByEmail Eメール認証
func AuthenticationByEmail(email string, password string) (*authModel.AuthUserResponse, error) {
	if email == "" || password == "" {
		return nil, errorConstants.ErrBadRequest
	}
	entity, err := userRepository.GetUserEntityByEmail(email)
	if err != nil {
		return nil, errorConstants.ErrInternalServer
	}
	if entity == nil {
		return nil, errorConstants.ErrNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(password))
	if err != nil {
		return nil, errorConstants.ErrUnauthorized
	}
	if entity.Status != "active" {
		return nil, errorConstants.ErrForbidden
	}
	// ログイン情報の更新
	authRepository.UpdateLoginInfo(entity.ID, entity.LoginCount+1)

	return GetResponseFromEntity(entity), nil
}

// AuthenticationByGoogle Google認証
func AuthenticationByGoogle(idToken string) (*authModel.AuthUserResponse, error) {
	if idToken == "" {
		return nil, errorConstants.ErrBadRequest
	}
	gmail := GetGmailFromGoogleIDToken(idToken)
	if gmail == "" {
		return nil, errorConstants.ErrForbidden
	}
	userEntity, err := userRepository.GetUserEntityByGoogleID(gmail)
	if err != nil {
		return nil, errorConstants.ErrInternalServer
	}
	if userEntity == nil {
		return nil, errorConstants.ErrNotFound
	}
	if userEntity.Status != "active" {
		return nil, errorConstants.ErrForbidden
	}
	// ログイン情報の更新
	authRepository.UpdateLoginInfo(userEntity.ID, userEntity.LoginCount+1)

	return GetResponseFromEntity(userEntity), nil
}

// AuthenticationByFacebook Facebook認証
func AuthenticationByFacebook(accessToken string) (*authModel.AuthUserResponse, error) {
	if accessToken == "" {
		return nil, errorConstants.ErrBadRequest
	}
	facebookID := GetIDFromFacebookAccessToken(accessToken)
	if facebookID == "" {
		return nil, errorConstants.ErrForbidden
	}
	userEntity, err := userRepository.GetUserEntityByFacebookID(facebookID)
	if err != nil {
		return nil, errorConstants.ErrInternalServer
	}
	if userEntity == nil {
		return nil, errorConstants.ErrNotFound
	}
	if userEntity.Status != "active" {
		return nil, errorConstants.ErrForbidden
	}
	// ログイン情報の更新
	authRepository.UpdateLoginInfo(userEntity.ID, userEntity.LoginCount+1)

	return GetResponseFromEntity(userEntity), nil
}

// GetUserByGoogleIDToken GoogleのIDトークンによるユーザの取得
func GetUserByGoogleIDToken(idToken string) (*authModel.AuthUserResponse, error) {
	gmail := GetGmailFromGoogleIDToken(idToken)
	if gmail == "" {
		return nil, errorConstants.ErrForbidden
	}
	userEntity, err := userRepository.GetUserEntityByGoogleID(gmail)
	if err != nil {
		return nil, errorConstants.ErrInternalServer
	}

	if userEntity == nil {
		return nil, errorConstants.ErrNotFound
	}

	return GetResponseFromEntity(userEntity), nil
}

// GetGmailFromGoogleIDToken GoogleのIDトークンが認証できたらgmailアドレスを返す
func GetGmailFromGoogleIDToken(idToken string) string {
	httpClient := &http.Client{}
	oauth2Service, _ := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return ""
	}
	return tokenInfo.Email
}

// GetUserByFacebookAccessToken FacebookのAccessトークンによるユーザの取得
func GetUserByFacebookAccessToken(accessToken string) (*authModel.AuthUserResponse, error) {
	facebookID := GetIDFromFacebookAccessToken(accessToken)
	if facebookID == "" {
		return nil, errorConstants.ErrForbidden
	}
	userEntity, err := userRepository.GetUserEntityByFacebookID(facebookID)
	if err != nil {
		return nil, errorConstants.ErrInternalServer
	}

	if userEntity == nil {
		return nil, errorConstants.ErrNotFound
	}

	return GetResponseFromEntity(userEntity), nil
}

// GetIDFromFacebookAccessToken Facebookのアクセストークンが認証できたらIDを返す
func GetIDFromFacebookAccessToken(accessToken string) string {
	url := os.Getenv("FACEBOOK_API_DOMAIN") + "debug_token?input_token=" + accessToken + "&access_token=" +
		os.Getenv("FACEBOOK_LOGIN_APP_ID") + "|" + os.Getenv("FACEBOOK_LOGIN_APP_SECFRET")
	resultStr, err := util.SendGetHTTPRequest(url)
	if err != nil {
		return ""
	}
	var responseInterface interface{}
	err = json.Unmarshal([]byte(resultStr), &responseInterface)
	if err != nil {
		return ""
	}
	id, err := dproxy.New(responseInterface).M("data").M("user_id").String()
	if err != nil {
		return ""
	}
	return id
}

// UpdateLoginInfo ログインの情報を更新
func UpdateLoginInfo(authUser *authModel.AuthUserResponse) error {
	userEntity, err := userRepository.GetUserEntityByUserSettingID(authUser.UserSettingID, "")
	if err == nil {
		// ログイン情報の更新
		err = authRepository.UpdateLoginInfo(userEntity.ID, userEntity.LoginCount+1)
	}

	return err
}

// GetResponseFromEntity Entityからレスポンスとする認証情報を返す
func GetResponseFromEntity(entity *authModel.AuthUserEntity) *authModel.AuthUserResponse {
	return &authModel.AuthUserResponse{
		ID:            entity.ID,
		UserType:      entity.UserType,
		UserSettingID: entity.UserSettingID,
		Name:          entity.Name,
		Categories:    entity.Categories,
		IconURL:       entity.IconURL,
		AuthMethod:    entity.AuthMethod,
		Exp:           GetAuthExpire(),
	}

}
