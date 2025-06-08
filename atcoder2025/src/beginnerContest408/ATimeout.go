package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ATimeoutMain() {
	var n, s int
	fmt.Scan(&n, &s)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	tStrList := strings.Split(ATimeoutRdr(rdr), " ")
	result := "Yes"

	before := 0
	for _, tStr := range tStrList {
		t, _ := strconv.Atoi(tStr)
		if t-before > s {
			result = "No"
			break
		}
		before = t
	}

	fmt.Println(result)
}

func ATimeoutRdr(rdr *bufio.Reader) string {
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
