package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	sStrs := strings.Split(EStampReadLine(rdr), "")
	tStrs := strings.Split(EStampReadLine(rdr), "")

	resultStrs := make([]string, n)

	result := "No"
	for i := 0; i < n; i++ {
		if sStrs[i] == resultStrs[i] {
			if i == n-1 {
				result = "Yes"
			}
		} else {
			if n-i >= m {
				// 置き換える
				for j := 0; j < m; j++ {
					resultStrs[j+i] = tStrs[j]
				}
				// 置き換えた後もダメか
				if sStrs[i] != resultStrs[i] {
					break
				}
			} else {
				break
			}
		}
	}

	fmt.Println(result)

}

func EStampReadLine(rdr *bufio.Reader) string {
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
