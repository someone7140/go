package auth

import (
	"net/http"

	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	"coda-api/src/service/auth"

	"github.com/gin-gonic/gin"
)

// AuthEmailRegister 認証用のメール登録
func AuthEmailRegister(c *gin.Context) {
	var request authModel.EmailAuthRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = auth.AuthEmailRegister(request.Email, request.Password)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}

}

// AuthEmailToken メール認証用のトークン
func AuthEmailToken(c *gin.Context) {
	var request authModel.EmailAuthTokenRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = auth.AuthEmailToken(request.UserID, request.Password, request.Token)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrUnauthorized {
		c.Status(http.StatusUnauthorized)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}

}

// ChangePassword パスワード変更
func ChangePassword(c *gin.Context) {
	var request authModel.PasswordChangeRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	err = auth.ChangePassword(request.Password, loginInfo.ID)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}

}

// RegisterPasswordReset パスワードリセット登録
func RegisterPasswordReset(c *gin.Context) {
	var request authModel.PasswordResetRegisterRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = auth.RegisterPasswordReset(request.Email)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}

}

// PasswordResetUpdate パスワードリセットによる更新
func PasswordResetUpdate(c *gin.Context) {
	var request authModel.PasswordResetUpdateRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = auth.PasswordResetUpdate(request.UserID, request.Token, request.Password)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}

}
