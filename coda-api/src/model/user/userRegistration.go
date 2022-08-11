package userModel

// RegisterUserByGoogleLoginRequest Googleログインによるユーザ登録リクエスト
type RegisterUserByGoogleLoginRequest struct {
	GoogleIDToken string          `form:"google_id_token" bson:"google_id_token"`
	UserInfo      UserInfoRequest `form:"user_info" bson:"user_info"`
}

// RegisterUserByFacebookLoginRequest Facebookログインによるユーザ登録リクエスト
type RegisterUserByFacebookLoginRequest struct {
	FacebookAccessToken string          `form:"facebook_access_token" bson:"facebook_access_token"`
	UserInfo            UserInfoRequest `form:"user_info" bson:"user_info"`
}

// RegisterUserByEmailAuthRequest メール認証によるユーザ登録リクエスト
type RegisterUserByEmailAuthRequest struct {
	UserID   string          `form:"user_id" bson:"user_id"`
	Token    string          `form:"token" bson:"token"`
	UserInfo UserInfoRequest `form:"user_info" bson:"user_info"`
}

// UpdateUserInfoRequest ユーザ更新リクエスト
type UpdateUserInfoRequest struct {
	UserInfo UserInfoRequest `form:"user_info" bson:"user_info"`
}

// UserInfoRequest ユーザ情報のリクエスト内容
type UserInfoRequest struct {
	UserSettingID string   `form:"user_info[user_setting_id]" bson:"user_setting_id"`
	Name          string   `form:"user_info[name]" bson:"name"`
	Gender        string   `form:"user_info[gender]" bson:"gender"`
	BirthDate     string   `form:"user_info[birth_date]" bson:"birth_date"`
	Silhouette    string   `form:"user_info[silhouette]" bson:"silhouette"`
	Height        int      `form:"user_info[height]" bson:"height"`
	Weight        int      `form:"user_info[weight]" bson:"weight"`
	Genres        []string `form:"user_info[genres][]" bson:"genres"`
	Complexes     []string `form:"user_info[complexes][]" bson:"complexes"`
}

// UserRegisterEntity ユーザ情報のEntity
type UserRegisterEntity struct {
	ID            string   `json:"_id" bson:"_id"`
	Email         string   `json:"email" bson:"email"`
	Password      string   `json:"password" bson:"password"`
	FacebookID    string   `json:"facebook_id" bson:"facebook_id"`
	GoogleID      string   `json:"google_id" bson:"google_id"`
	UserSettingID string   `json:"user_setting_id" bson:"user_setting_id"`
	Name          string   `json:"name" bson:"name"`
	Gender        string   `json:"gender" bson:"gender"`
	BirthDate     string   `json:"birth_date" bson:"birth_date"`
	Silhouette    string   `json:"silhouette" bson:"silhouette"`
	Height        int      `json:"height" bson:"height"`
	Weight        int      `json:"weight" bson:"weight"`
	Genres        []string `json:"genres" bson:"genres"`
	Complexes     []string `json:"complexes" bson:"complexes"`
	Categories    []string `json:"categories" bson:"categories"`
	Status        string   `json:"status" bson:"status"`
	UserType      string   `json:"user_type" bson:"user_type"`
	IconURL       string   `json:"icon_url" bson:"icon_url"`
}
