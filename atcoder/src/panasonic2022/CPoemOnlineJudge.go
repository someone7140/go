package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CPoemOnlineJudgeMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	mapSt := map[string]int{}
	maxPoint := -10000
	result := -10000
	for i := 0; i < n; i++ {
		stLine := CPoemOnlineJudgeReadLine(rdr)
		stStrArray := strings.Split(stLine, " ")
		s := stStrArray[0]
		t, _ := strconv.Atoi(stStrArray[1])

		_, ok := mapSt[s]
		if !ok {
			mapSt[s] = 1
			if maxPoint < t {
				maxPoint = t
				result = i + 1
			}
		}
	}
	fmt.Println(result)
}

func CPoemOnlineJudgeReadLine(rdr *bufio.Reader) string {
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
