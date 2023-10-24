package modelDb

// PostCategoriesEntity postCategoriesのEntity
type PostCategoriesEntity struct {
	ID                  string  `json:"id" bson:"_id"`
	Name                string  `json:"name" bson:"name"`
	CreateUserAccountId string  `json:"createUserAccountId" bson:"create_user_account_id"`
	ParentCategoryId    *string `json:"parentCategoryId" bson:"parent_category_id"`
	Memo                *string `json:"memo" bson:"memo"`
	DisplayOrder        *int32  `json:"displayOrder" bson:"display_order"`
}
