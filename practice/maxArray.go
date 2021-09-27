package main

import (
	"fmt"

	"github.com/adam-lavrik/go-imath/ix"
	"github.com/thoas/go-funk"
)

func mainMaxArray() {
	sampleArray := []int{1, 6, 2, 10, 9, 14, 100, 5}
	// 愚直にループを回す
	max := -9999999
	for _, s := range sampleArray {
		if max < s {
			max = s
		}
	}
	fmt.Println(max)

	// go-imathのライブラリを使う
	max = -9999999
	max = ix.MaxSlice(sampleArray)
	fmt.Println(max)

	//go-funkのライブラリを使う
	max = -9999999
	max = funk.MaxInt(sampleArray)
	fmt.Println(max)
}
