package service

import (
	"context"
	"fmt"
	"os"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"

	"github.com/jackc/pgx/v5"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

// LINEで通知アカウントを登録された場合
func FollowEvent(
	ctx context.Context,
	lineID string,
	bot *linebot.Client,
	event webhook.FollowEvent) error {
	// 該当のIDですでに会員登録済みか
	user, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountByLineId(ctx, lineID)
	if err != nil {
		if err != pgx.ErrNoRows {
			return err
		} else {
			// まだ会員登録してない場合はその旨をメッセージ
			notRegisteredMessage := fmt.Sprintf(
				"会員登録されてません。一度ブロックをして頂いた後、こちらより会員登録をして頂き再度友達追加をお願いします。:%s%s",
				os.Getenv("FRONTEND_DOMAIN"),
				os.Getenv("LINE_AUTH_REGISTER_REDIRECT_PATH"))
			_, err = bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(notRegisteredMessage),
			).Do()
			return err
		}

	}

	// ユーザのレコードの通知フラグを更新
	_, err = custom_middleware.GetDbQueries(ctx).UpdateUserAccountLineBotFollow(
		ctx, db.UpdateUserAccountLineBotFollowParams{
			ID:              user.ID,
			IsLineBotFollow: true,
		})
	if err != nil {
		return err
	}

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("通知登録をしました"),
	).Do()

	return err
}

// LINEで通知アカウント登録を解除された場合
func UnFollowEvent(
	ctx context.Context,
	lineID string,
	bot *linebot.Client,
	event webhook.UnfollowEvent) error {
	// 該当のIDですでに会員登録済みか
	user, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountByLineId(ctx, lineID)
	if err != nil {
		if err != pgx.ErrNoRows {
			return err
		} else {
			// まだ会員登録してない場合はスキップ
			return nil
		}

	}

	// ユーザのレコードの通知フラグを更新
	_, err = custom_middleware.GetDbQueries(ctx).UpdateUserAccountLineBotFollow(
		ctx, db.UpdateUserAccountLineBotFollowParams{
			ID:              user.ID,
			IsLineBotFollow: false,
		})
	if err != nil {
		return err
	}
	// タスクのレコードの通知フラグを全てOFFで更新
	_, err = custom_middleware.GetDbQueries(ctx).UpdateAllTaskNotificationFlagByUser(
		ctx, db.UpdateAllTaskNotificationFlagByUserParams{
			OwnerUserID:      user.ID,
			NotificationFlag: false,
		})
	return err
}
