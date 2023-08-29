package placeNoteUtil

import (
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

// GenerateUUID UUIDを生成して取得
func GenerateUUID() (string, *connect.Error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", connect.NewError(connect.CodeInternal, err)
	}
	return uid.String(), nil
}
