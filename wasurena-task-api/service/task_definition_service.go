package service

import (
	"context"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

func CreateTaskService(ctx context.Context, input model.NewTask) (bool, error) {
	id := xid.New()
	userAccountId := custom_middleware.GeUserAccountId(ctx)
	createData := db.CreateTaskDefinitionParams{
		ID:                      id.String(),
		Title:                   input.Title,
		OwnerUserID:             *userAccountId,
		DisplayFlag:             input.DisplayFlag,
		NotificationFlag:        input.NotificationFlag,
		CategoryID:              input.CategoryID,
		DeadLineCheck:           input.DeadLineCheck,
		DeadLineCheckSubSetting: input.DeadLineCheckSubSetting,
		Detail:                  input.Detail,
	}
	_, err := custom_middleware.GetDbQueries(ctx).CreateTaskDefinition(ctx, createData)

	if err != nil {
		return false, err
	}
	return true, err
}
