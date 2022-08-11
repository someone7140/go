package userRepository

import (
	"coda-api/src/util"
	"context"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	authModel "coda-api/src/model/auth"
	userModel "coda-api/src/model/user"
)

var selectProjection = bson.M{
	"_id":             1,
	"email":           1,
	"password":        1,
	"facebook_id":     1,
	"google_id":       1,
	"status":          1,
	"user_type":       1,
	"user_setting_id": 1,
	"name":            1,
	"categories":      1,
	"icon_url":        1,
	"login_count":     1,
	"email_auth":      1,
	"password_reset":  1,
}

// GetUserEntityByUserID userIDでのユーザ取得
func GetUserEntityByUserID(userID string, authMethod string) (*authModel.AuthUserEntity, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "_id", Value: userID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfo(queryResult[0], authMethod)
}

// GetUserEntityByUserSettingID userSettingIDでのユーザ取得
func GetUserEntityByUserSettingID(userSettingID string, authMethod string) (*authModel.AuthUserEntity, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "user_setting_id", Value: userSettingID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfo(queryResult[0], authMethod)
}

// GetUserEntityByEmail emailでのユーザ取得
func GetUserEntityByEmail(email string) (*authModel.AuthUserEntity, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "email", Value: email}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfo(queryResult[0], "email")
}

// GetUserEntityByGoogleID googleIDでのユーザ取得
func GetUserEntityByGoogleID(googleID string) (*authModel.AuthUserEntity, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "google_id", Value: googleID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfo(queryResult[0], "googleID")
}

// GetUserEntityByFacebookID facebookIDでのユーザ取得
func GetUserEntityByFacebookID(facebookID string) (*authModel.AuthUserEntity, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "facebook_id", Value: facebookID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfo(queryResult[0], "facebookID")
}

// CheckUserSettingIDDuplicate UserSettingIdが重複していないかチェック
func CheckUserSettingIDDuplicate(userSettingID string) bool {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "user_setting_id", Value: userSettingID}}
		err := col.FindOne(context.Background(), filter).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return false
	}
	if queryResult == nil {
		return true
	}
	return false
}

// decodeUserInfo bson.Mからユーザの構造体にdecode
func decodeUserInfo(input primitive.M, authMethod string) (*authModel.AuthUserEntity, error) {
	id, errGet := dproxy.New(input).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	email, errGet := dproxy.New(input).M("email").String()
	password, errGet := dproxy.New(input).M("password").String()
	if len(email) > 0 && errGet != nil {
		return nil, errGet
	}
	facebookID, errGet := dproxy.New(input).M("facebook_id").String()
	googleID, errGet := dproxy.New(input).M("google_id").String()
	status, errGet := dproxy.New(input).M("status").String()
	if errGet != nil {
		return nil, errGet
	}
	userType, errGet := dproxy.New(input).M("user_type").String()
	if errGet != nil {
		return nil, errGet
	}
	userSettingID, errGet := dproxy.New(input).M("user_setting_id").String()
	if errGet != nil {
		return nil, errGet
	}
	name, errGet := dproxy.New(input).M("name").String()
	if errGet != nil {
		return nil, errGet
	}
	var categories []string
	categoriesInput := input["categories"].(primitive.A)
	for i := 0; i < reflect.ValueOf(categoriesInput).Len(); i++ {
		category, errGet := dproxy.New(categoriesInput[i]).String()
		if errGet != nil {
			return nil, errGet
		}
		categories = append(categories, category)
	}
	iconURL, errGet := dproxy.New(input).M("icon_url").String()
	loginCount, errGet := dproxy.New(input).M("login_count").Int64()
	if errGet != nil {
		return nil, errGet
	}
	var emailAuthEntity *authModel.EmailAuthEntity
	emailAuthEmail, _ := dproxy.New(input).M("email_auth").M("email").String()
	if emailAuthEmail != "" {
		token, errGet := dproxy.New(input).M("email_auth").M("token").String()
		if errGet != nil {
			return nil, errGet
		}
		e := input["email_auth"].(primitive.M)
		expireDateInt := int64(e["expire_date"].(primitive.DateTime))
		emailAuthEntity = &authModel.EmailAuthEntity{
			Email:      emailAuthEmail,
			Token:      token,
			ExpireDate: util.GetMongoDateIntToTime(expireDateInt),
		}

	} else {
		emailAuthEntity = nil
	}

	var passwordResetEntity *authModel.PasswordResetEntity
	passwordResetToken, _ := dproxy.New(input).M("password_reset").M("token").String()
	if passwordResetToken != "" {
		p := input["password_reset"].(primitive.M)
		expireDateInt := int64(p["expire_date"].(primitive.DateTime))
		passwordResetEntity = &authModel.PasswordResetEntity{
			Token:      passwordResetToken,
			ExpireDate: util.GetMongoDateIntToTime(expireDateInt),
		}

	} else {
		passwordResetEntity = nil
	}

	return &authModel.AuthUserEntity{
		ID:            id,
		Email:         email,
		Password:      password,
		FacebookID:    facebookID,
		GoogleID:      googleID,
		Status:        status,
		UserType:      userType,
		UserSettingID: userSettingID,
		Name:          name,
		Categories:    categories,
		IconURL:       iconURL,
		LoginCount:    loginCount,
		AuthMethod:    authMethod,
		EmailAuth:     emailAuthEntity,
		PasswordReset: passwordResetEntity,
	}, nil

}

// GetUserEntityByUserIDForDetail userIDでの詳細用ユーザ取得
func GetUserEntityByUserIDForDetail(userID string) (*userModel.UserDetailInfoResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "_id", Value: userID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(
			bson.M{
				"_id":             1,
				"user_setting_id": 1,
				"name":            1,
				"icon_url":        1,
				"attribute":       1,
			},
		)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeUserInfoDetail(queryResult[0])
}

// GetRegistrationUserCount 指定した日付までに登録された会員数
func GetRegistrationUserCount(untilDate time.Time) (int64, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStatusStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "active"}},
		}}
		matchRegistrationDateStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "registration_date", Value: bson.D{
				{Key: "$lte", Value: untilDate}},
			}}}}
		countStage := bson.D{{Key: "$count", Value: "user_count"}}

		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStatusStage, matchRegistrationDateStage, countStage,
			})
		if err != nil {
			return nil, err
		}
		var docs []bson.M
		for cur.Next(context.Background()) {
			var doc bson.M
			if err = cur.Decode(&doc); err != nil {
				return nil, err
			}
			docs = append(docs, doc)
		}
		return docs, nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return 0, err
	}
	if reflect.ValueOf(queryResult).Len() == 0 {
		return 0, nil
	}
	count, err := dproxy.New(queryResult[0]).M("user_count").Int64()
	if err != nil {
		return 0, err
	}
	internalCount, _ := strconv.ParseInt(os.Getenv("INTERNAL_USER_COUNT"), 10, 64)
	return count - internalCount, nil
}

// GetLoginUserCount 指定した日付までに登録された会員数
func GetLoginUserCount(startDate time.Time, endDate time.Time) (int64, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchTypeStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "user_type", Value: "normal"}},
		}}
		matchStartDateStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "login_date", Value: bson.D{
				{Key: "$gte", Value: startDate}},
			}}}}
		matchEndDateStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "login_date", Value: bson.D{
				{Key: "$lte", Value: endDate}},
			}}}}
		countStage := bson.D{{Key: "$count", Value: "user_count"}}

		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchTypeStage, matchStartDateStage, matchEndDateStage, countStage,
			})
		if err != nil {
			return nil, err
		}
		var docs []bson.M
		for cur.Next(context.Background()) {
			var doc bson.M
			if err = cur.Decode(&doc); err != nil {
				return nil, err
			}
			docs = append(docs, doc)
		}
		return docs, nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return 0, err
	}
	if reflect.ValueOf(queryResult).Len() == 0 {
		return 0, nil
	}
	count, err := dproxy.New(queryResult[0]).M("user_count").Int64()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// decodeUserInfo bson.Mからユーザ詳細の構造体にdecode
func decodeUserInfoDetail(input primitive.M) (*userModel.UserDetailInfoResponse, error) {
	id, errGet := dproxy.New(input).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	userSettingID, errGet := dproxy.New(input).M("user_setting_id").String()
	if errGet != nil {
		return nil, errGet
	}
	name, errGet := dproxy.New(input).M("name").String()
	if errGet != nil {
		return nil, errGet
	}
	iconURL, errGet := dproxy.New(input).M("icon_url").String()
	gender, _ := dproxy.New(input).M("attribute").M("gender").String()
	birthDate, _ := dproxy.New(input).M("attribute").M("birth_date").String()
	silhouette, _ := dproxy.New(input).M("attribute").M("silhouette").String()
	height, _ := dproxy.New(input).M("attribute").M("height").Int64()
	weight, _ := dproxy.New(input).M("attribute").M("weight").Int64()

	attributeInput := input["attribute"].(primitive.M)

	var genres []string
	if attributeInput["genres"] != nil {
		genresInput := attributeInput["genres"].(primitive.A)
		for i := 0; i < reflect.ValueOf(genresInput).Len(); i++ {
			genre, errGet := dproxy.New(genresInput[i]).String()
			if errGet != nil {
				return nil, errGet
			}
			genres = append(genres, genre)
		}
	}

	var complexes []string
	if attributeInput["complexes"] != nil {
		complexessInput := attributeInput["complexes"].(primitive.A)
		for i := 0; i < reflect.ValueOf(complexessInput).Len(); i++ {
			complex, errGet := dproxy.New(complexessInput[i]).String()
			if errGet != nil {
				return nil, errGet
			}
			complexes = append(complexes, complex)
		}
	}

	return &userModel.UserDetailInfoResponse{
		ID:            id,
		UserSettingID: userSettingID,
		Name:          name,
		IconURL:       iconURL,
		Gender:        gender,
		BirthDate:     birthDate,
		Silhouette:    silhouette,
		Height:        int(height),
		Weight:        int(weight),
		Genres:        genres,
		Complexes:     complexes,
	}, nil
}
