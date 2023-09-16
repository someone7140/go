package userUseCase

import (
	"errors"
	placeNote "placeNote/src/gen/proto"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"
	"placeNote/src/repository"

	"github.com/bufbuild/connect-go"
)

const UserAccountJwtPropertyName = "userAccountId"

type userAccountId int

var UserAccountIdContextKey userAccountId

type userAccountToken int

var UserAccountTokenContextKey userAccountToken

// NewRegistrationUserAccountUseCase ユーザアカウントの新規登録
func NewRegistrationUserAccountUseCase(req *connect.Request[placeNote.RegisterUserAccountRequest]) (*connect.Response[placeNote.UserAccountResponse], *connect.Error) {
	reqMsg := req.Msg

	// userSettingIdの重複チェック
	_, err := repository.GetUserAccountByUserSettingIdRepository(reqMsg.UserSettingId)
	// ユーザが取得できている（すでにidが登録済み）の場合はエラー
	if err == nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("CodeAlreadyExists"))
	} else if err.Code() != connect.CodeNotFound {
		return nil, err
	}

	uid, err := placeNoteUtil.GenerateUUID()
	if err != nil {
		return nil, err
	}

	entity := modelDb.UserAccountsEntity{
		ID:            uid,
		UserSettingId: reqMsg.UserSettingId,
		Name:          reqMsg.Name,
		AuthMethod:    reqMsg.AuthMethod,
	}

	if reqMsg.AuthMethod == placeNote.AuthMethod_GOOGLE {
		// tokenからgmailを取得
		gmail, err := GetGmailFromToken(reqMsg.AuthToken)
		if err != nil {
			return nil, err
		}
		// gmailの重複チェック
		_, err = repository.GetUserAccountByGmailRepository(gmail)
		// ユーザが取得できている（すでにgmailが登録済み）の場合はエラー
		if err == nil {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("GmailAlreadyExists"))
		} else if err.Code() != connect.CodeNotFound {
			return nil, err
		}
		entity.Gmail = gmail
	}
	if reqMsg.AuthMethod == placeNote.AuthMethod_EMAIL {
		// emailの重複チェック
	}

	// 登録処理
	err = repository.NewRegistrationUserAccountRepository(entity)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("can not register"))
	}
	// idをトークン化
	idToken, err := GetJwtTokenAuth(UserAccountJwtPropertyName, entity.ID, GetJwtAuthMonthExpire(3))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&placeNote.UserAccountResponse{
		Token:         idToken,
		AuthMethod:    entity.AuthMethod,
		UserSettingId: entity.UserSettingId,
		Name:          entity.Name,
	}), nil

}

// GetUserAccountFromUserAccountIdUseCase ユーザアカウントIDからユーザを取得
func GetUserAccountFromUserAccountIdUseCase(userAccountId string, token string) (*connect.Response[placeNote.UserAccountResponse], *connect.Error) {
	//userAccountの取得
	userAccount, err := repository.GetUserAccountByUserAccountIdRepository(userAccountId)
	// ユーザが取得できている（すでにidが登録済み）の場合はエラー
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&placeNote.UserAccountResponse{
		Token:         token,
		AuthMethod:    userAccount.AuthMethod,
		UserSettingId: userAccount.UserSettingId,
		Name:          userAccount.Name,
	}), nil

}
