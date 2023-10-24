package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ASameMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(ASameReadLine(rdr), " ")
	result := "Yes"
	a := aStrs[0]
	for i := 1; i < n; i++ {
		if a != aStrs[i] {
			result = "No"
			break
		}
	}

	fmt.Println(result)

}

func ASameReadLine(rdr *bufio.Reader) string {
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
