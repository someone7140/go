package modelDb

// PostPlacesEntity postPlacesのEntity
type PostPlacesEntity struct {
	ID                  string     `json:"id" bson:"_id"`
	Name                string     `json:"name" bson:"name"`
	CreateUserAccountId string     `json:"createUserAccountId" bson:"create_user_account_id"`
	Address             *string    `json:"address" bson:"address"`
	LonLat              *[]float64 `json:"lonLat" bson:"lon_lat"`
	PrefectureCode      *string    `json:"prefectureCode" bson:"prefecture_code"`
	CategoryIdList      *[]string  `json:"categoryIdList" bson:"category_id_list"`
	Detail              *string    `json:"detail" bson:"detail"`
	UrlList             *[]string  `json:"urlList" bson:"url_list"`
}
