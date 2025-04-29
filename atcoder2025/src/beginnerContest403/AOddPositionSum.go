package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AOddPositionSumMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(AOddPositionSumRdr(rdr), " ")
	result := 0

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a, _ := strconv.Atoi(aStrs[i])
			result = result + a
		}
	}
	fmt.Println(result)
}

func AOddPositionSumRdr(rdr *bufio.Reader) string {
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
