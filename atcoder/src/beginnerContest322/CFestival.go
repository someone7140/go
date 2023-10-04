package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CFestivalMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrList := strings.Split(CFestivalReadLine(rdr), " ")
	nowMIndex := 0
	nowMValue, _ := strconv.Atoi(aStrList[0])
	var resultSlice []string

	for i := 0; i < n; i++ {
		tempResult := i + 1
		if tempResult < nowMValue {
			tempResult = nowMValue - tempResult
			resultSlice = append(resultSlice, strconv.FormatInt(int64(tempResult), 10))
		} else {
			resultSlice = append(resultSlice, "0")
			nowMIndex = nowMIndex + 1
			if nowMIndex < m {
				tempMValue, _ := strconv.Atoi(aStrList[nowMIndex])
				nowMValue = tempMValue
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CFestivalReadLine(rdr *bufio.Reader) string {
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
