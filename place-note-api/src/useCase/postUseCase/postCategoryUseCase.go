package postUseCase

import (
	placeNote "placeNote/src/gen/proto"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"
	"placeNote/src/repository"

	"github.com/bufbuild/connect-go"
)

// AddPostCategory 投稿カテゴリーの追加
func AddPostCategory(req *placeNote.AddPostCategoryRequest, userAccountId string) (*string, *connect.Error) {

	uid, err := placeNoteUtil.GenerateUUID()
	if err != nil {
		return nil, err
	}
	// 登録処理
	err = repository.AddPostCategoryRepository(modelDb.PostCategoriesEntity{
		ID:                  uid,
		Name:                req.Name,
		CreateUserAccountId: userAccountId,
		ParentCategoryId:    req.ParentId,
		Memo:                req.Memo,
		DisplayOrder:        req.DisplayOrder,
	})
	if err != nil {
		return nil, err
	}

	return &uid, err
}

// UpdatePostCategory 投稿カテゴリーの更新
func UpdatePostCategory(req *placeNote.UpdatePostCategoryRequest, userAccountId string) *connect.Error {

	// 更新処理
	err := repository.UpdatePostCategoryRepository(modelDb.PostCategoriesEntity{
		ID:                  req.Id,
		Name:                req.Name,
		CreateUserAccountId: userAccountId,
		ParentCategoryId:    req.ParentId,
		Memo:                req.Memo,
		DisplayOrder:        req.DisplayOrder,
	})
	if err != nil {
		return err
	}

	return err
}

// DeletePostCategory 投稿カテゴリーの削除
func DeletePostCategory(id string, userAccountId string) *connect.Error {

	// 削除処理
	err := repository.DeletePostCategoryRepository(id, userAccountId)
	if err != nil {
		return err
	}

	return err
}

// GetMyPostCategoryList 自分の投稿カテゴリー一覧を取得
func GetMyPostCategoryList(userAccountId string) (*connect.Response[placeNote.GetPostCategoryListResponse], *connect.Error) {

	categoryEntities, err := repository.GetPostCategoryListByUserAccountIdRepository(userAccountId)
	if err != nil {
		return nil, err
	}
	var categories []*placeNote.PostCategory
	for _, entity := range categoryEntities {
		categories = append(categories, &placeNote.PostCategory{
			Id:           entity.ID,
			Name:         entity.Name,
			ParentId:     entity.ParentCategoryId,
			Memo:         entity.Memo,
			DisplayOrder: entity.DisplayOrder,
		})
	}

	return connect.NewResponse(&placeNote.GetPostCategoryListResponse{
		CategoryList: categories,
	}), nil
}
