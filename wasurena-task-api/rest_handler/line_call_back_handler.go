package rest_handler

import (
	"os"
	"wasurena-task-api/service"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func LineCallBackHander(c echo.Context) error {
	// bot作成
	bot, err := linebot.New(
		os.Getenv("LINE_MESSAGE_SECRET"),
		os.Getenv("LINE_MESSAGE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return err
	}
	// コールバックのリクエスト
	callbackRequest, err := webhook.ParseRequest(os.Getenv("LINE_MESSAGE_SECRET"), c.Request())
	if err != nil {
		return err
	}

	for _, event := range callbackRequest.Events {
		// イベントの種類
		switch e := event.(type) {
		case webhook.FollowEvent:
			source := e.Source
			// ソースタイプがユーザーの場合のみ
			switch s := source.(type) {
			case webhook.UserSource:
				// 友達登録された時の処理
				err = service.FollowEvent(
					c.Request().Context(),
					s.UserId,
					bot,
					e,
				)
				return err
			}
		case webhook.UnfollowEvent:
			source := e.Source
			// ソースタイプがユーザーの場合のみ
			switch s := source.(type) {
			case webhook.UserSource:
				// 友達登録された時の処理
				err = service.UnFollowEvent(
					c.Request().Context(),
					s.UserId,
					bot,
					e,
				)
				return err
			}
		}
	}

	return nil
}
