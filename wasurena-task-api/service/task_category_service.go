package service

import (
	"context"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/middleware"

	"github.com/rs/xid"
)

func CreateCategoryService(ctx context.Context, input model.NewCategory) (bool, error) {
	id := xid.New()
	create_data := db.CreateTaskCategoryParams{
		ID:           id.String(),
		Name:         input.Name,
		OwnerUserID:  "dummy_user",
		DisplayOrder: input.DisplayOrder,
	}
	_, err := middleware.GetDbQueries(ctx).CreateTaskCategory(ctx, create_data)

	if err != nil {
		return false, err
	}
	return true, err
}
