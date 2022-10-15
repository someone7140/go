package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AIntegerSumMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := AIntegerSumRdr(rdr)
	aStrArray := strings.Split(aLine, " ")

	result := 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		result = result + a
	}

	fmt.Println(result)

}

func AIntegerSumRdr(rdr *bufio.Reader) string {
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
