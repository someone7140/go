package service

import (
	"context"
	"time"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

// タスクの実行履歴を追加する
func CreateTaskExecuteService(ctx context.Context, input model.NewTaskExecute) (bool, error) {
	id := xid.New()
	userAccountID := custom_middleware.GeUserAccountID(ctx)
	// 現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return false, err
	}
	now := time.Now().In(jst)

	createData := db.CreateTaskExecuteParams{
		ID:               id.String(),
		TaskDefinitionID: input.TaskDefinitionID,
		ExecuteUserID:    *userAccountID,
		ExecuteDateTime:  now,
		Memo:             input.Memo,
	}
	_, err = custom_middleware.GetDbQueries(ctx).CreateTaskExecute(ctx, createData)

	if err != nil {
		return false, err
	}
	return true, err
}

// タスクの定義を指定してタスクの実行履歴を取得する
func GetTaskExecuteListByDefinitionIDService(ctx context.Context, taskDefinitionID string) ([]*model.TaskExecuteResponse, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)
	responseSlice := []*model.TaskExecuteResponse{}

	taskExecutes, err := custom_middleware.GetDbQueries(ctx).SelectTaskExecuteByDefinitionId(ctx, db.SelectTaskExecuteByDefinitionIdParams{
		ExecuteUserID:    *userAccountID,
		TaskDefinitionID: taskDefinitionID,
	})
	if err != nil {
		return responseSlice, err
	}

	for _, taskExec := range taskExecutes {
		responseSlice = append(responseSlice, &model.TaskExecuteResponse{
			ID:               taskExec.ID,
			TaskDefinitionID: taskExec.TaskDefinitionID,
			ExecuteDateTime:  taskExec.ExecuteDateTime,
			Memo:             taskExec.Memo,
		})
	}

	return responseSlice, err
}

// タスクの実行履歴を削除する
func DeleteTaskExecuteService(ctx context.Context, taskExecuteID string) (bool, error) {
	userAccountID := custom_middleware.GeUserAccountID(ctx)

	// 実行履歴のレコード削除
	deleteTaskExecuteData := db.DeleteTaskExecuteByIdParams{
		ID:            taskExecuteID,
		ExecuteUserID: *userAccountID,
	}
	_, err := custom_middleware.GetDbQueries(ctx).DeleteTaskExecuteById(ctx, deleteTaskExecuteData)
	if err != nil {
		return false, err
	}

	return true, err
}
