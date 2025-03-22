package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CVarietySplitEasyMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aStrs := strings.Split(CVarietySplitEasyRdr(rdr), " ")
	aArray := make([]int, n)
	map1 := map[int]int{}
	map2 := map[int]int{}
	countSyurui1 := 0
	countSyurui2 := 0
	for i, aMoji := range aStrs {
		a, _ := strconv.Atoi(aMoji)
		aArray[i] = a

		if i == 0 {
			map1[a] = 1
			countSyurui1 = 1
		} else {
			val, ok := map2[a]
			if ok {
				map2[a] = val + 1
			} else {
				map2[a] = 1
				countSyurui2 = countSyurui2 + 1
			}

		}
	}

	result := countSyurui1 + countSyurui2
	for i := 1; i < n-1; i++ {
		a := aArray[i]
		// 2を削除
		val2, ok2 := map2[a]
		if ok2 {
			map2[a] = val2 - 1
			if val2 < 2 {
				countSyurui2 = countSyurui2 - 1
			}
		}

		// 1に追加
		val1, ok1 := map1[a]
		if ok1 {
			map1[a] = val1 + 1
		} else {
			map1[a] = 1
			countSyurui1 = countSyurui1 + 1
		}

		if result < (countSyurui1 + countSyurui2) {
			result = countSyurui1 + countSyurui2
		}
	}

	fmt.Println(result)
}

func CVarietySplitEasyRdr(rdr *bufio.Reader) string {
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
