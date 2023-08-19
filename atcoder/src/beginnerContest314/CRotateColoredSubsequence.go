package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CRotateColoredSubsequenceMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CRotateColoredSubsequenceReadLine(rdr)

	sArrayMap := map[int][]string{}
	sArrayFromIndexMap := map[int]int{}
	sArrayNowIndexMap := map[int]int{}

	for i, cMoji := range strings.Split(CRotateColoredSubsequenceReadLine(rdr), " ") {
		c, _ := strconv.Atoi(cMoji)

		sMoji := s[i : i+1]
		sArray, ok1 := sArrayMap[c]
		if ok1 {
			sArrayMap[c] = append(sArray, sMoji)
		} else {
			sArrayMap[c] = []string{sMoji}
		}

		sArrayFromIndexMap[i] = c

		_, ok2 := sArrayNowIndexMap[c]
		if !ok2 {
			sArrayNowIndexMap[c] = 0
		}

	}

	// 右シフト
	for k, v := range sArrayMap {
		lenV := len(v)
		if lenV > 1 {
			newV := append([]string{v[lenV-1]}, v[0:lenV-1]...)
			sArrayMap[k] = newV
		}
	}

	// 結果出力
	resultArray := make([]string, n)
	for i := 0; i < n; i++ {
		c := sArrayFromIndexMap[i]
		sArray := sArrayMap[c]
		sArrayNowIndex := sArrayNowIndexMap[c]
		resultArray[i] = sArray[sArrayNowIndex]
		sArrayNowIndexMap[c] = sArrayNowIndex + 1
	}

	fmt.Println(strings.Join(resultArray, ""))
}

func CRotateColoredSubsequenceReadLine(rdr *bufio.Reader) string {
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
