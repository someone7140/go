package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ACountingPassesMain() {
	var n, l int
	fmt.Scan(&n, &l)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(ACountingPassesReadLine(rdr), " ")

	result := 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrs[i])
		if a >= l {
			result = result + 1
		}
	}

	fmt.Println(result)

}

func ACountingPassesReadLine(rdr *bufio.Reader) string {
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
