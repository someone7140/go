package util

import (
	"github.com/google/uuid"
)

// GenerateUUID UUIDを生成して取得
func GenerateUUID() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}
	return uid.String(), nil
}
