package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nC121321Main int

func C121321Main() {
	fmt.Scan(&nC121321Main)
	if nC121321Main == 1 {
		fmt.Println("1")
	} else {
		sArray := []string{"1"}
		fmt.Println(calcC121321Main(sArray, 1))
	}

}

func calcC121321Main(sArray []string, index int) string {
	addSlice := []string{strconv.Itoa(index + 1)}
	newSlice := append(append(sArray, addSlice...), sArray...)
	if index == (nC121321Main - 1) {
		return strings.Join(newSlice, " ")
	} else {
		return calcC121321Main(newSlice, index+1)
	}
}
