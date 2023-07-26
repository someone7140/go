package main

import (
	"follow-check/src/controller"
	"syscall/js"
)

func main() {
	js.Global().Set("checkInstagram", js.FuncOf(controller.ExecInstagramCheck))
	select {}
}

/*
func main() {
	result, err := service.CheckInstagramFollow("testUserId", "testPassword")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
*/
