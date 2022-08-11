package userModel

// UserDetailInfoResponse ユーザの詳細情報レスポンス
type UserDetailInfoResponse struct {
	ID            string   `json:"_id" bson:"_id"`
	UserSettingID string   `json:"user_setting_id" bson:"user_setting_id"`
	Name          string   `json:"name" bson:"name"`
	IconURL       string   `json:"icon_url" bson:"icon_url"`
	Gender        string   `json:"gender" bson:"gender"`
	BirthDate     string   `json:"birth_date" bson:"birth_date"`
	Silhouette    string   `json:"silhouette" bson:"silhouette"`
	Height        int      `json:"height" bson:"height"`
	Weight        int      `json:"weight" bson:"weight"`
	Genres        []string `json:"genres" bson:"genres"`
	Complexes     []string `json:"complexes" bson:"complexes"`
}
