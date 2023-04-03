package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BCatMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BCatReadLine(rdr)

	fmt.Println(strings.ReplaceAll(s, "na", "nya"))
}

func BCatReadLine(rdr *bufio.Reader) string {
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
