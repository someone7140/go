package topImageModel

import (
	"time"

	coordinateModel "coda-api/src/model/coordinate"
)

// TopImageResponse トップ画像用のresponse用struct
type TopImageResponse struct {
	RecentPosts   []TopImage `json:"recent_posts" bson:"recent_posts"`
	SaleOnlyPosts []TopImage `json:"sale_only_posts" bson:"sale_only_posts"`
}

// TopImage トップ画像のstruct
type TopImage struct {
	ID                 string                        `json:"_id" bson:"_id"`
	Category           string                        `json:"category" bson:"category"`
	Title              string                        `json:"title" bson:"title"`
	ImageUrl           string                        `json:"image_url" bson:"image_url"`
	PostDate           time.Time                     `json:"post_date" bson:"post_date"`
	ShopName           string                        `json:"shop_name" bson:"shop_name"`
	ShopSettingID      string                        `json:"shop_setting_id" bson:"shop_setting_id"`
	ItemType           string                        `json:"item_type" bson:"item_type"`
	PostUserID         string                        `json:"post_user_id" bson:"post_user_id"`
	PostUserSettingID  string                        `json:"post_user_setting_id" bson:"post_user_setting_id"`
	PostUserName       string                        `json:"post_user_name" bson:"post_user_name"`
	FavoritedFlg       bool                          `json:"favorited_flg" bson:"favorited_flg"`
	CoordinateCategory string                        `json:"coordinate_category" bson:"coordinate_category"`
	Price              int                           `json:"price" bson:"price"`
	Sale               *coordinateModel.SaleResponse `json:"sale" bson:"sale"`
}
