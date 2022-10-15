package main

import (
	"fmt"
)

func ReplaceSampleMain() {
	result := fmt.Sprintf("%[1]s一%[2]s　%[1]s二%[2]s　%[1]s三%[2]s", "第", "著者")
	fmt.Println(result)
	fmt.Printf("%[1]s一%[2]s　%[1]s二%[2]s　%[1]s三%[2]s", "第", "著者")
}
