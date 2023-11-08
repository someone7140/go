package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CNumberPlaceMain() {
	var aSliceSlice = make([][]int, 9)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < 9; i++ {
		aStrs := strings.Split(CNumberPlaceReadLine(rdr), " ")
		var aSlice = make([]int, 9)
		for j := 0; j < 9; j++ {
			a, _ := strconv.Atoi(aStrs[j])
			aSlice[j] = a
		}
		aSliceSlice[i] = aSlice
	}

	result := true

	// 行のチェック
	for i := 0; i < 9; i++ {
		aSlice := aSliceSlice[i]
		checkMap := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
		for j := 0; j < 9; j++ {
			v, ok := checkMap[aSlice[j]]
			if ok {
				delete(checkMap, v)
			}
		}
		if len(checkMap) != 0 {
			result = false
			break
		}
	}
	if !result {
		fmt.Println("No")
		return
	}

	// 列のチェック
	for i := 0; i < 9; i++ {
		checkMap := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
		for j := 0; j < 9; j++ {
			v, ok := checkMap[aSliceSlice[j][i]]
			if ok {
				delete(checkMap, v)
			}
		}
		if len(checkMap) != 0 {
			result = false
			break
		}
	}
	if !result {
		fmt.Println("No")
		return
	}

	// 正方形のチェック
	for i := 0; i <= 6; i = i + 3 {
		for j := 0; j <= 6; j = j + 3 {
			checkMap := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
			// 起点から9マス
			for i2 := 0; i2 <= 2; i2++ {
				for j2 := 0; j2 <= 2; j2++ {
					v, ok := checkMap[aSliceSlice[i+i2][j+j2]]
					if ok {
						delete(checkMap, v)
					}
				}
			}
			if len(checkMap) != 0 {
				result = false
				break
			}
		}

	}
	if !result {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}

func CNumberPlaceReadLine(rdr *bufio.Reader) string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
