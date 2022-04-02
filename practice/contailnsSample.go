package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

// SampleStruct サンプル用の構造体
type SampleStruct struct {
	Key   string
	Value string
}

func contailnsSample() {
	// 数字が含まれているか
	var numberSlice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(funk.Contains(numberSlice, 6))

	// 文字列が含まれているか
	var stringSlice = []string{"a", "b", "c"}
	fmt.Println(funk.Contains(stringSlice, "c"))

	// 構造体のキーが含まれているか
	var structSlice = []SampleStruct{{
		Key:   "key1",
		Value: "value1",
	}, {
		Key:   "key2",
		Value: "value2",
	},
	}
	fmt.Println(funk.Contains(structSlice, func(structSlice SampleStruct) bool {
		return structSlice.Key == "key1"
	}))
}
