package service

import (
	"context"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

func CreateTaskCategoryService(ctx context.Context, input model.NewCategory) (bool, error) {
	id := xid.New()
	userAccountId := custom_middleware.GeUserAccountId(ctx)

	createData := db.CreateTaskCategoryParams{
		ID:           id.String(),
		Name:         input.Name,
		OwnerUserID:  *userAccountId,
		DisplayOrder: input.DisplayOrder,
	}
	_, err := custom_middleware.GetDbQueries(ctx).CreateTaskCategory(ctx, createData)

	if err != nil {
		return false, err
	}
	return true, err
}

func GetTaskCategoriesService(ctx context.Context) ([]*model.TaskCategoryResponse, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)

	categories, err := custom_middleware.GetDbQueries(ctx).SelectTaskCategories(ctx, *userAccountId)
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

func DeleteTaskCategoryService(ctx context.Context, id string) (bool, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)

	// タスク定義のカテゴリーをnullにする
	updateTaskData := db.UpdateAllTaskCategoryNullParams{
		CategoryID:  &id,
		OwnerUserID: *userAccountId,
	}
	_, err := custom_middleware.GetDbQueries(ctx).UpdateAllTaskCategoryNull(ctx, updateTaskData)
	if err != nil {
		return false, err
	}

	// カテゴリーのレコード削除
	deleteCategoryData := db.DeleteTaskCategoryParams{
		ID:          id,
		OwnerUserID: *userAccountId,
	}
	_, err = custom_middleware.GetDbQueries(ctx).DeleteTaskCategory(ctx, deleteCategoryData)
	if err != nil {
		return false, err
	}

	return true, err
}
