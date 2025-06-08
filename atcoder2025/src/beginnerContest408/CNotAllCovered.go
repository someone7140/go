package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CNotAllCoveredLr struct {
	startCount int
	endCount   int
}

func CNotAllCoveredMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	lrList := make([]CNotAllCoveredLr, n)
	for i := 0; i < m; i++ {
		lrStrList := strings.Split(CNotAllCoveredRdr(rdr), " ")
		l, _ := strconv.Atoi(lrStrList[0])
		r, _ := strconv.Atoi(lrStrList[1])
		l = l - 1
		r = r - 1

		lrList[l].startCount = lrList[l].startCount + 1
		lrList[r].endCount = lrList[r].endCount + 1
	}

	nowCount := 0
	result := 0
	for i := 0; i < n; i++ {
		// 一つ前のエンドのカウントをマイナスする
		if i > 0 {
			nowCount = nowCount - lrList[i-1].endCount
		}
		// スタートのカウントをプラスする
		nowCount = nowCount + lrList[i].startCount

		if i == 0 {
			result = nowCount
		}
		if nowCount < result {
			result = nowCount
		}
	}

	fmt.Println(result)
}

func CNotAllCoveredRdr(rdr *bufio.Reader) string {
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
