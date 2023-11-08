package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DGoodTupleProblemMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var aSlice = make([]int, m)
	aStrs := strings.Split(DGoodTupleProblemReadLine(rdr), " ")
	for j := 0; j < m; j++ {
		a, _ := strconv.Atoi(aStrs[j])
		aSlice[j] = a
	}

	var bSlice = make([]int, m)
	bStrs := strings.Split(DGoodTupleProblemReadLine(rdr), " ")
	for j := 0; j < m; j++ {
		b, _ := strconv.Atoi(bStrs[j])
		bSlice[j] = b
	}

	abMap := map[int]map[int]struct{}{}

	for j := 0; j < m; j++ {
		a := aSlice[j]
		b := bSlice[j]
		// aに追加
		aMapV, ok1 := abMap[a]
		if !ok1 {
			temp := map[int]struct{}{}
			temp[b] = struct{}{}
			abMap[a] = temp
		} else {
			aMapV[b] = struct{}{}
			abMap[a] = aMapV
		}
		// bに追加
		bMapV, ok2 := abMap[b]
		if !ok2 {
			temp := map[int]struct{}{}
			temp[a] = struct{}{}
			abMap[b] = temp
		} else {
			bMapV[a] = struct{}{}
			abMap[b] = bMapV
		}
	}

	result := "Yes"
	resultMap := map[int]int{}
	alreadyMap := map[int]struct{}{}

	var funcLoop func(key int)
	funcLoop = func(key int) {
		resultValue, _ := resultMap[key]
		values, ok := abMap[key]
		if ok && result == "Yes" {
			for key2, _ := range values {
				resultValue2, ok2 := resultMap[key2]
				if !ok2 {
					if resultValue == 0 {
						resultMap[key2] = 1
					} else {
						resultMap[key2] = 0
					}
					funcLoop(key2)
				} else {
					if resultValue2 == resultValue {
						result = "No"
						break
					}
				}
				_, ok3 := alreadyMap[key2]
				if !ok3 {
					alreadyMap[key2] = struct{}{}
					funcLoop(key2)
				}
			}
		}
	}

	startFlag := false

	for key, _ := range abMap {
		if !startFlag {
			resultMap[key] = 0
			startFlag = true
		}
		_, ok := resultMap[key]
		if !ok {
			resultMap[key] = 0
		}
		_, ok3 := alreadyMap[key]
		if !ok3 {
			alreadyMap[key] = struct{}{}
			funcLoop(key)
		}

		if result == "No" {
			break
		}
	}

	fmt.Println(result)
}

func DGoodTupleProblemReadLine(rdr *bufio.Reader) string {
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
