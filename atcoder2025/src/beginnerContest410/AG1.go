package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AG1Main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrList := strings.Split(AG1Rdr(rdr), " ")
	k, _ := strconv.Atoi(AG1Rdr(rdr))

	result := 0
	for _, aStr := range aStrList {
		a, _ := strconv.Atoi(aStr)
		if a >= k {
			result = result + 1
		}
	}

	fmt.Println(result)
}

func AG1Rdr(rdr *bufio.Reader) string {
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
