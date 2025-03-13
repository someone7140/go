package service

import (
	"context"
	"time"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/middleware"

	"github.com/rs/xid"
)

func CreateTaskExecuteService(ctx context.Context, input model.NewTaskExecute) (bool, error) {
	id := xid.New()
	// 現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return false, err
	}
	now := time.Now().In(jst)

	create_data := db.CreateTaskExecuteParams{
		ID:               id.String(),
		TaskDefinitionID: input.TaskDefinitionID,
		ExecuteUserID:    "dummy_user",
		ExecuteDateTime:  now,
		Memo:             input.Memo,
	}
	_, err = middleware.GetDbQueries(ctx).CreateTaskExecute(ctx, create_data)

	if err != nil {
		return false, err
	}
	return true, err
}
