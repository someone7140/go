package service

import (
	"math/rand"
	"time"

	"xorm.io/xorm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewAuthenticationUserService(dbEngine *xorm.Engine) *AuthenticationUserService {
	return &AuthenticationUserService{
		dbEngine: dbEngine,
	}
}

func NewGeographicPointService(dbEngine *xorm.Engine) *GeographicPointService {
	return &GeographicPointService{
		dbEngine: dbEngine,
	}
}
