package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AConflictMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	tStrs := strings.Split(AConflictRdr(rdr), "")
	aStrs := strings.Split(AConflictRdr(rdr), "")
	result := "No"
	for i := 0; i < n; i++ {
		if tStrs[i] == "o" && aStrs[i] == "o" {
			result = "Yes"
		}
	}

	fmt.Println(result)
}

func AConflictRdr(rdr *bufio.Reader) string {
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
