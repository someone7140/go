package user

import (
	"coda-api/src/service/auth"
	"coda-api/src/service/user"
	"net/http"

	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	userModel "coda-api/src/model/user"

	"github.com/gin-gonic/gin"
)

// CheckRegisterByGoogleLogin Googleログイン認証済みユーザの登録ができるか
func CheckRegisterByGoogleLogin(c *gin.Context) {
	idToken := c.Query("id_token")
	if idToken == "" {
		c.Status(http.StatusBadRequest)
	}
	user, err := auth.GetUserByGoogleIDToken(idToken)
	if user != nil {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrNotFound {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusInternalServerError)
	}
}

// RegisterByGoogleLogin Googleログイン認証済みユーザの登録
func RegisterByGoogleLogin(c *gin.Context) {
	var request userModel.RegisterUserByGoogleLoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	file, fileHeader, err := c.Request.FormFile("user_info[icon_image]")
	response, err := user.RegisterByGoogleLogin(request, file, fileHeader)
	setUserResponse(c, response, err)

}

// CheckRegisterByFacebookLogin Facebookログイン認証済みユーザの登録ができるか
func CheckRegisterByFacebookLogin(c *gin.Context) {
	accessToken := c.Query("access_token")
	if accessToken == "" {
		c.Status(http.StatusBadRequest)
	}
	user, err := auth.GetUserByFacebookAccessToken(accessToken)
	if user != nil {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrNotFound {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusInternalServerError)
	}
}

// RegisterByFacebookLogin Facebookログイン認証済みユーザの登録
func RegisterByFacebookLogin(c *gin.Context) {
	var request userModel.RegisterUserByFacebookLoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	file, fileHeader, err := c.Request.FormFile("user_info[icon_image]")
	response, err := user.RegisterByFacdebookLogin(request, file, fileHeader)
	setUserResponse(c, response, err)
}

// RegisterByEmailAuth メール認証済みユーザの登録
func RegisterByEmailAuth(c *gin.Context) {
	var request userModel.RegisterUserByEmailAuthRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	file, fileHeader, err := c.Request.FormFile("user_info[icon_image]")
	response, err := user.RegisterByEmailAuth(request, file, fileHeader)
	setUserResponse(c, response, err)
}

// UpdateUserInfo ユーザ情報の更新
func UpdateUserInfo(c *gin.Context) {
	var request userModel.UpdateUserInfoRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	file, fileHeader, err := c.Request.FormFile("user_info[icon_image]")
	loginInfo := auth.GetLoginInfoFromContext(c)
	response, err := user.UpdateUserInfo(loginInfo, request, file, fileHeader)
	setUserResponse(c, response, err)
}

// DeleteUser ユーザの削除
func DeleteUser(c *gin.Context) {
	loginInfo := auth.GetLoginInfoFromContext(c)
	err := user.DeleteUser(loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

func setUserResponse(c *gin.Context, response *authModel.AuthUserResponse, err error) {
	if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		token, tokenGetErr := auth.GetJwtTokenUser(*response)
		if tokenGetErr != nil {
			c.Status(http.StatusInternalServerError)
		} else {
			response.Token = token
			c.JSON(http.StatusOK, response)
		}
	}
}
