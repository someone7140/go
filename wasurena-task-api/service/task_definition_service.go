package service

import (
	"context"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/middleware"

	"github.com/rs/xid"
)

func CreateTaskService(ctx context.Context, input model.NewTask) (bool, error) {
	id := xid.New()
	create_data := db.CreateTaskDefinitionParams{
		ID:                      id.String(),
		Title:                   input.Title,
		OwnerUserID:             "dummy_user",
		DisplayFlag:             input.DisplayFlag,
		NotificationFlag:        input.NotificationFlag,
		CategoryID:              input.CategoryID,
		DeadLineCheck:           input.DeadLineCheck,
		DeadLineCheckSubSetting: input.DeadLineCheckSubSetting,
		Detail:                  input.Detail,
	}
	_, err := middleware.GetDbConnection(ctx).CreateTaskDefinition(ctx, create_data)

	if err != nil {
		return false, err
	}
	return true, err
}
