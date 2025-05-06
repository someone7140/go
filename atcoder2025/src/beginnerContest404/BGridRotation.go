package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BGridRotationMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sListList := make([][]string, n)
	tListList := make([][]string, n)

	// sの読み込み
	for i := 0; i < n; i++ {
		sList := strings.Split(BGridRotationRdr(rdr), "")
		sListList[i] = sList
	}

	// tの読み込み
	for i := 0; i < n; i++ {
		tList := strings.Split(BGridRotationRdr(rdr), "")
		tListList[i] = tList
	}

	result := -1

	for i := 0; i < 4; i++ {
		tempResult := 0
		if i != 0 {
			tempResult = tempResult + i
			// 90度ひっくり返す
			newSListList := make([][]string, n)
			for j := 0; j < n; j++ {
				newSListList[j] = make([]string, n)
			}

			for j1 := 0; j1 < n; j1++ {
				for j2 := 0; j2 < n; j2++ {
					newSListList[j2][j1] = sListList[n-1-j1][j2]
				}
			}
			sListList = newSListList
		}

		for j1 := 0; j1 < n; j1++ {
			for j2 := 0; j2 < n; j2++ {
				if sListList[j1][j2] != tListList[j1][j2] {
					tempResult = tempResult + 1
				}
			}
		}
		if result == -1 || result > tempResult {
			result = tempResult
		}
	}
	fmt.Println(result)
}

func BGridRotationRdr(rdr *bufio.Reader) string {
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
