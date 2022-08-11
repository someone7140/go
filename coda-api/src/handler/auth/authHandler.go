package auth

import (
	"net/http"

	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	"coda-api/src/service/auth"

	"github.com/gin-gonic/gin"
)

// AuthCheck 認証チェック
func AuthCheck(c *gin.Context) {
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	auth.UpdateLoginInfo(loginInfo)
	c.Status(http.StatusOK)
}

// LoginByGoogle Google認証
func LoginByGoogle(c *gin.Context) {
	var request authModel.GoogleAuthRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	response, err := auth.AuthenticationByGoogle(request.GoogleIDToken)
	setUserResponse(c, response, err)
}

// LoginByFacebook Facebook認証
func LoginByFacebook(c *gin.Context) {
	var request authModel.FacebookAuthRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	response, err := auth.AuthenticationByFacebook(request.FacebookAccessToken)
	setUserResponse(c, response, err)
}

// LoginByEmail Eメール認証
func LoginByEmail(c *gin.Context) {
	var request authModel.EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	response, err := auth.AuthenticationByEmail(request.Email, request.Password)
	setUserResponse(c, response, err)
}

func setUserResponse(c *gin.Context, response *authModel.AuthUserResponse, err error) {
	if err == errorConstants.ErrInternalServer {
		c.Status(http.StatusInternalServerError)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrUnauthorized {
		c.Status(http.StatusUnauthorized)
	} else if err == errorConstants.ErrNotFound {
		c.Status(http.StatusNotFound)
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
