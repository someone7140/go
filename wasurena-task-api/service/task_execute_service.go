package service

import (
	"context"
	"time"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

func CreateTaskExecuteService(ctx context.Context, input model.NewTaskExecute) (bool, error) {
	id := xid.New()
	userAccountId := custom_middleware.GeUserAccountId(ctx)
	// 現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return false, err
	}
	now := time.Now().In(jst)

	createData := db.CreateTaskExecuteParams{
		ID:               id.String(),
		TaskDefinitionID: input.TaskDefinitionID,
		ExecuteUserID:    *userAccountId,
		ExecuteDateTime:  now,
		Memo:             input.Memo,
	}
	_, err = custom_middleware.GetDbQueries(ctx).CreateTaskExecute(ctx, createData)

	if err != nil {
		return false, err
	}
	return true, err
}
