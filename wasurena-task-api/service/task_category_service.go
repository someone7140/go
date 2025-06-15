package service

import (
	"context"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// カテゴリーを追加する
func CreateTaskCategoryService(ctx context.Context, input model.CategoryInput) (bool, error) {
	id := xid.New()
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	createData := db.CreateTaskCategoryParams{
		ID:           id.String(),
		Name:         input.Name,
		OwnerUserID:  *userAccountID,
		DisplayOrder: input.DisplayOrder,
	}
	_, err := custom_middleware.GetDbQueries(ctx).CreateTaskCategory(ctx, createData)

	if err != nil {
		return false, err
	}
	return true, err
}

// カテゴリーを取得する
func GetTaskCategoriesService(ctx context.Context) ([]*model.TaskCategoryResponse, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	categories, err := custom_middleware.GetDbQueries(ctx).SelectTaskCategories(ctx, *userAccountID)
	if err != nil {
		return []*model.TaskCategoryResponse{}, err
	}

	responseSlice := []*model.TaskCategoryResponse{}
	for _, category := range categories {
		responseSlice = append(responseSlice, &model.TaskCategoryResponse{
			ID:           category.ID,
			Name:         category.Name,
			DisplayOrder: category.DisplayOrder,
		})
	}

	return responseSlice, err
}

// ID指定でカテゴリーを取得する
func GetTaskCategoryByIDService(ctx context.Context, categoryID string) (*model.TaskCategoryResponse, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	category, err := custom_middleware.GetDbQueries(ctx).SelectTaskCategoryByID(ctx, db.SelectTaskCategoryByIDParams{
		OwnerUserID: *userAccountID,
		ID:          categoryID,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &gqlerror.Error{
				Message: "Can not find taskCategory",
				Extensions: map[string]any{
					"code": 404,
				}}
		} else {
			return nil, err
		}
	}

	response := model.TaskCategoryResponse{
		ID:           category.ID,
		Name:         category.Name,
		DisplayOrder: category.DisplayOrder,
	}

	return &response, err
}

// カテゴリーを更新する
func UpdateTaskCategoryService(ctx context.Context, id string, input model.CategoryInput) (bool, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	updateData := db.UpdateTaskCategoryParams{
		ID:           id,
		Name:         input.Name,
		OwnerUserID:  *userAccountID,
		DisplayOrder: input.DisplayOrder,
	}
	_, err := custom_middleware.GetDbQueries(ctx).UpdateTaskCategory(ctx, updateData)

	if err != nil {
		return false, err
	}
	return true, err
}

// カテゴリーを削除する
func DeleteTaskCategoryService(ctx context.Context, id string) (bool, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	// タスク定義のカテゴリーをnullにする
	updateTaskData := db.UpdateAllTaskCategoryNullParams{
		CategoryID:  &id,
		OwnerUserID: *userAccountID,
	}
	_, err := custom_middleware.GetDbQueries(ctx).UpdateAllTaskCategoryNull(ctx, updateTaskData)
	if err != nil {
		return false, err
	}

	// カテゴリーのレコード削除
	deleteCategoryData := db.DeleteTaskCategoryParams{
		ID:          id,
		OwnerUserID: *userAccountID,
	}
	_, err = custom_middleware.GetDbQueries(ctx).DeleteTaskCategory(ctx, deleteCategoryData)
	if err != nil {
		return false, err
	}

	return true, err
}
