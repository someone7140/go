package cron

import (
	"coda-api/src/service/analytics"
	"coda-api/src/service/instagramBatch"
	"context"
	"fmt"
	"net/http"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/gin-gonic/gin"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

// StartPostGatherHandler Instagram投稿の収集開始Handler
func StartPostGatherHandler(c *gin.Context) {
	StartCronHandler(c, "/executePostGatherBatch")
}

// StartAccessAnalysisHandler アクセス情報の収集開始Handler
func StartAccessAnalysisHandler(c *gin.Context) {
	StartCronHandler(c, "/executeAccessAnalysisBatch")
}

// StartCronHandler Cronの開始Handler
func StartCronHandler(c *gin.Context, processUri string) {
	if c.Request.Header.Get("X-Appengine-Cron") != "true" {
		c.Status(http.StatusBadRequest)
	}
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	locationID := os.Getenv("QUEUE_LOCATION_ID")
	queueID := "coda-queue"
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queueID)

	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#AppEngineHttpRequest
			MessageType: &taskspb.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &taskspb.AppEngineHttpRequest{
					HttpMethod:  taskspb.HttpMethod_POST,
					RelativeUri: processUri,
				},
			},
		},
	}
	_, err = client.CreateTask(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

// ExecutePostGatherBatch インスタ投稿収集のバッチ実行Handler
func ExecutePostGatherBatch(c *gin.Context) {
	// バッチのヘッダー確認
	t, ok := c.Request.Header["X-Appengine-Taskname"]
	if !ok || len(t[0]) == 0 {
		c.Status(http.StatusBadRequest)
	}
	_, ok = c.Request.Header["X-Appengine-Queuename"]
	if !ok {
		c.Status(http.StatusBadRequest)
	}
	// 実処理の実行
	err := instagramBatch.GatherInstagramPostService()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// ExecuteAccessAnalysis アクセス情報収集の実行Handler
func ExecuteAccessAnalysisBatch(c *gin.Context) {
	// バッチのヘッダー確認
	t, ok := c.Request.Header["X-Appengine-Taskname"]
	if !ok || len(t[0]) == 0 {
		c.Status(http.StatusBadRequest)
	}
	_, ok = c.Request.Header["X-Appengine-Queuename"]
	if !ok {
		c.Status(http.StatusBadRequest)
	}
	// 実処理の実行
	err := analytics.AccessAnalysisBatchService()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
