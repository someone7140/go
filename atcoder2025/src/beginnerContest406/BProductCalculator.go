package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BProductCalculatorMain() {
	var n, k int
	fmt.Scan(&n, &k)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrList := strings.Split(BProductCalculatorRdr(rdr), " ")

	var result int64 = 1
	for _, aStr := range aStrList {
		resultStr := strconv.FormatInt(int64(result), 10)
		resultLen := len(resultStr)
		aLen := len(aStr)

		tempTotalLen := resultLen + aLen - 1

		if tempTotalLen > k {
			result = 1
		} else {
			a, _ := strconv.ParseInt(aStr, 10, 64)
			tempResult := a * result
			tempResultStr := strconv.FormatInt(int64(tempResult), 10)
			tempResultLen := len(tempResultStr)
			if tempResultLen > k {
				result = 1
			} else {
				result = tempResult
			}
		}
	}

	fmt.Println(result)
}

func BProductCalculatorRdr(rdr *bufio.Reader) string {
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
