package service

import (
	"context"
	"slices"
	"time"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/domain"
	"wasurena-task-api/graph/model"

	"github.com/rs/xid"
)

// タスク定義を追加する
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

// タスク定義を取得する
func GetTaskDefinitionService(ctx context.Context) ([]*model.TaskDefinitionResponse, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)
	selectResults, err := custom_middleware.GetDbQueries(ctx).SelectTaskDefinitionList(ctx, *userAccountId)

	responseSlice := []*model.TaskDefinitionResponse{}
	for _, task := range selectResults {
		responseSlice = append(responseSlice, &model.TaskDefinitionResponse{
			ID:                      task.ID,
			Title:                   task.Title,
			DisplayFlag:             task.DisplayFlag,
			NotificationFlag:        task.NotificationFlag,
			CategoryID:              task.CategoryID,
			DeadLineCheck:           task.DeadLineCheck,
			DeadLineCheckSubSetting: task.DeadLineCheckSubSetting,
			Detail:                  task.Detail,
			CategoryName:            task.CategoryName,
		})
	}

	return responseSlice, err
}

// タスク定義を削除
func DeleteTaskDefinitionService(ctx context.Context, id string) (bool, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)

	// 実行履歴のレコード削除
	deleteTaskExecuteData := db.DeleteTaskExecuteByDefinitionIdParams{
		TaskDefinitionID: id,
		OwnerUserID:      *userAccountId,
	}
	_, err := custom_middleware.GetDbQueries(ctx).DeleteTaskExecuteByDefinitionId(ctx, deleteTaskExecuteData)
	if err != nil {
		return false, err
	}

	// タスク定義のレコード削除
	deleteTaskDefinitionData := db.DeleteTaskDefinitionParams{
		ID:          id,
		OwnerUserID: *userAccountId,
	}
	_, err = custom_middleware.GetDbQueries(ctx).DeleteTaskDefinition(ctx, deleteTaskDefinitionData)
	if err != nil {
		return false, err
	}

	return true, err
}

// タスクのチェック対象一覧を取得
func GetTaskCheckDisplayListService(ctx context.Context) ([]*model.TaskCheckDisplayResponse, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)

	// DBからレコード取得
	selectResults, err := custom_middleware.GetDbQueries(ctx).SelectTaskCheckDisplayList(ctx, *userAccountId)
	if err != nil {
		return []*model.TaskCheckDisplayResponse{}, err
	}

	// 現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return []*model.TaskCheckDisplayResponse{}, err
	}
	now := time.Now().In(jst)

	// 期限が超過してるタスクをフィルタリング
	checkList := []domain.TaskDeadLineCheckTarget{}
	for _, checkTask := range selectResults {
		if checkTask.DeadLineCheck != nil {
			checkList = append(checkList, domain.TaskDeadLineCheckTarget{
				TaskID:                  checkTask.ID,
				TaskTitle:               checkTask.Title,
				DeadLineCheck:           *checkTask.DeadLineCheck,
				DeadLineCheckSubSetting: checkTask.DeadLineCheckSubSetting,
				OwnerUserID:             checkTask.OwnerUserID,
				LatestExecDateTime:      checkTask.LatestExecDateTime,
			})
		}
	}
	deadLineCheck := domain.TaskDeadLineCheck{
		CheckTargetList: checkList,
		NowDateTime:     now,
	}
	exceededList := deadLineCheck.GetExceededTaskList()

	// 期限超過の判定をしてレスポンスのリストを作る
	exceededResponseSlice := []*model.TaskCheckDisplayResponse{}
	notExceededResponseSlice := []*model.TaskCheckDisplayResponse{}
	for _, check := range selectResults {
		isExceededContains := slices.ContainsFunc(exceededList, func(e domain.TaskDeadLineCheckTarget) bool {
			return e.TaskID == check.ID
		})

		response := model.TaskCheckDisplayResponse{
			ID:                      check.ID,
			Title:                   check.Title,
			DisplayFlag:             check.DisplayFlag,
			NotificationFlag:        check.NotificationFlag,
			CategoryID:              check.CategoryID,
			CategoryName:            check.CategoryName,
			DeadLineCheck:           check.DeadLineCheck,
			DeadLineCheckSubSetting: check.DeadLineCheckSubSetting,
			Detail:                  check.Detail,
			LatestExecDateTime:      getLatestExecDateTimeForDisplay(check.LatestExecDateTime, jst),
			IsExceedDeadLine:        isExceededContains,
		}
		// 期限超過してるかで入れる配列を判定
		if isExceededContains {
			exceededResponseSlice = append(exceededResponseSlice, &response)
		} else {
			notExceededResponseSlice = append(notExceededResponseSlice, &response)
		}
	}

	return append(exceededResponseSlice, notExceededResponseSlice...), err
}

// 表示用の日付を取得
func getLatestExecDateTimeForDisplay(targetTime time.Time, locale *time.Location) *time.Time {

	// 2020年より前の日付であれば実施無しとみなす
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, locale)
	diffHour := initialTime.Sub(targetTime).Hours()
	if diffHour > 0 {
		return nil
	}

	return &targetTime
}
