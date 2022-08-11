package coordinateModel

import "time"

// CoordinatePostRequest コーデ投稿登録のリクエスト
type CoordinatePostRequest struct {
	Title         string       `form:"title" bson:"title"`
	Detail        string       `form:"detail" bson:"detail"`
	URL           string       `form:"url" bson:"url"`
	Status        string       `form:"status" bson:"status"`
	ShopSettingID string       `form:"shop_setting_id" bson:"shop_setting_id"`
	Gender        string       `form:"gender" bson:"gender"`
	Silhouette    string       `form:"silhouette" bson:"silhouette"`
	Height        int          `form:"height" bson:"height"`
	Weight        int          `form:"weight" bson:"weight"`
	Size          string       `form:"size" bson:"size"`
	Price         int          `form:"price" bson:"price"`
	Category      string       `form:"category" bson:"category"`
	Sale          *SaleRequest `form:"sale" bson:"sale"`
}

// CoordinatePostUpdateRequest コーデ投稿更新のリクエスト
type CoordinatePostUpdateRequest struct {
	ID              string       `form:"_id" bson:"_id"`
	Title           string       `form:"title" bson:"title"`
	Detail          string       `form:"detail" bson:"detail"`
	URL             string       `form:"url" bson:"url"`
	Status          string       `form:"status" bson:"status"`
	ShopSettingID   string       `form:"shop_setting_id" bson:"shop_setting_id"`
	Gender          string       `form:"gender" bson:"gender"`
	Silhouette      string       `form:"silhouette" bson:"silhouette"`
	Height          int          `form:"height" bson:"height"`
	Weight          int          `form:"weight" bson:"weight"`
	Size            string       `form:"size" bson:"size"`
	DeleteImageKeys []int        `form:"delete_image_keys" bson:"delete_image_keys"`
	Price           int          `form:"price" bson:"price"`
	Category        string       `form:"category" bson:"category"`
	Sale            *SaleRequest `form:"sale" bson:"sale"`
}

// CoordinateDeleteRequest コーデ投稿削除のリクエスト
type CoordinateDeleteRequest struct {
	CoordinateID string `json:"coordinate_id" bson:"coordinate_id"`
}

// SaleRequest セール情報のリクエスト
type SaleRequest struct {
	SalePrice int        `form:"sale[sale_price]" bson:"sale_price"`
	StartDate *time.Time `form:"sale[start_date]" bson:"start_date"`
	EndDate   *time.Time `form:"sale[end_date]" bson:"end_date"`
}
