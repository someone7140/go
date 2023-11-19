package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DElectionQuickReportMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aStrs := strings.Split(DElectionQuickReportReadLine(rdr), " ")
	resultStr := make([]string, m)
	aMap := make(map[int]int)
	topNumber := -1
	for i, aStr := range aStrs {
		a, _ := strconv.Atoi(aStr)

		if i == 0 {
			topNumber = a
			aMap[a] = 1
			resultStr[i] = aStr
		} else {
			// 該当の番号に1票足す
			v, ok := aMap[a]
			if !ok {
				aMap[a] = 1
			} else {
				aMap[a] = v + 1
			}
			updatedCount := aMap[a]
			if a == topNumber {
				resultStr[i] = strconv.FormatInt(int64(topNumber), 10)
			} else {
				//票数の比較
				topCount := aMap[topNumber]
				if updatedCount > topCount {
					topNumber = a
					resultStr[i] = aStr
				} else if updatedCount == topCount {
					if topNumber < a {
						resultStr[i] = strconv.FormatInt(int64(topNumber), 10)
					} else {
						topNumber = a
						resultStr[i] = aStr
					}
				} else {
					resultStr[i] = strconv.FormatInt(int64(topNumber), 10)
				}
			}
		}
	}

	fmt.Println(strings.Join(resultStr, "\n"))

}

func DElectionQuickReportReadLine(rdr *bufio.Reader) string {
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
