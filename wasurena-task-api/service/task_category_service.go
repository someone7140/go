package service

import (
	"context"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

func CreateCategoryService(ctx context.Context, input model.NewCategory) (bool, error) {
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
