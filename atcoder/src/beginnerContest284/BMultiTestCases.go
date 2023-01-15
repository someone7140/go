package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BMultiTestCasesMain() {
	var t int
	fmt.Scan(&t)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	resultArray := make([]int, t)

	for i := 0; i < t; i++ {
		result := 0
		n, _ := strconv.Atoi(BMultiTestCasesReadLine(rdr))
		aStrArray := strings.Split(BMultiTestCasesReadLine(rdr), " ")
		for j := 0; j < n; j++ {
			a, _ := strconv.Atoi(aStrArray[j])
			if a%2 != 0 {
				result = result + 1
			}
		}
		resultArray[i] = result
	}

	for i := 0; i < t; i++ {
		fmt.Println(resultArray[i])
	}

}

func BMultiTestCasesReadLine(rdr *bufio.Reader) string {
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
