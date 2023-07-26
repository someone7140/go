package controller

import (
	"follow-check/src/service"
	"syscall/js"
)

// インスタグラムのチェック
func ExecInstagramCheck(this js.Value, args []js.Value) any {
	result, err := service.CheckInstagramFollow(args[0].String(), args[1].String())

	if err != nil {
		return map[string]any{"successFlag": false, "err": err.Error()}
	}
	return map[string]any{"successFlag": true, "text": result.Text}
}
